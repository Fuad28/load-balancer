package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

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
	if utils.SimulateServerDown() {
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
		Proxy:           utils.NewSingleHostReverseProxy(serverUrl),
	}
}

func NewProdServer(healthCheckPath string, address string) *Server {
	serverUrl, err := url.Parse(address)

	utils.OnErrorPanic(err, "Invalid server address")

	return &Server{
		healthCheckPath: serverUrl.JoinPath(healthCheckPath).String(),
		Address:         serverUrl,
		Proxy:           utils.NewSingleHostReverseProxy(serverUrl),
	}
}
