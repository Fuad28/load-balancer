package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/Fuad28/load-balancer/utils"
)

type Server struct {
	healthCheckPath string
	Address         *url.URL
	ServerMux       http.Handler
	Proxy           *httputil.ReverseProxy
}

func (s *Server) IsAlive(randomDown bool) bool {

	// randomly claim servers are not alive to simulate cases where server goes down.
	if SimulateServerDown() {
		return false
	}

	res, err := http.Get(s.healthCheckPath)

	if err != nil {
		log.Fatal(err)
		return false
	}

	if res.StatusCode == 200 {
		return true
	}

	return false

}

func (s *Server) Serve(w http.ResponseWriter, req *http.Request) {
	s.Proxy.ServeHTTP(w, req)

}

func NewDevServer(address string) *Server {
	serverUrl, err := url.Parse("http://" + address)

	utils.OnErrorPanic(err, "Invalid server address")

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Response from server: %v\n", address)

	})

	return &Server{
		healthCheckPath: serverUrl.String(),
		Address:         serverUrl,
		ServerMux:       mux,
		Proxy:           NewSingleHostReverseProxy(serverUrl),
	}
}

func NewProdServer(healthCheckPath string, address string) *Server {
	serverUrl, err := url.Parse(address)

	utils.OnErrorPanic(err, "Invalid server address")

	return &Server{
		healthCheckPath: serverUrl.JoinPath(healthCheckPath).String(),
		Address:         serverUrl,
		Proxy:           NewSingleHostReverseProxy(serverUrl),
	}
}

func SimulateServerDown() bool {

	trueFalseArr := []bool{true, false}
	idx := rand.Intn(len(trueFalseArr))

	return trueFalseArr[idx]
}

func singleJoiningSlash(a, b string) string {
	aslash := strings.HasSuffix(a, "/")
	bslash := strings.HasPrefix(b, "/")
	switch {
	case aslash && bslash:
		return a + b[1:]
	case !aslash && !bslash:
		return a + "/" + b
	}
	return a + b
}

// httputil implementation of NewSingleHostReverseProxy doesn't overwrite req.Host
// If left unmodified, the req.Host defaults to the load balance host.
func NewSingleHostReverseProxy(target *url.URL) *httputil.ReverseProxy {
	targetQuery := target.RawQuery
	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.Host = target.Host
		req.URL.Path = singleJoiningSlash(target.Path, req.URL.Path)
		if targetQuery == "" || req.URL.RawQuery == "" {
			req.URL.RawQuery = targetQuery + req.URL.RawQuery
		} else {
			req.URL.RawQuery = targetQuery + "&" + req.URL.RawQuery
		}
	}
	return &httputil.ReverseProxy{Director: director}
}
