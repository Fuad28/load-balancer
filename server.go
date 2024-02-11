package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server struct {
	healthCheckPath string
	Address         *url.URL
	ServerMux       http.Handler
	Proxy           *httputil.ReverseProxy
}

func (s *Server) IsAlive() bool {
	_, err := http.Get(s.healthCheckPath)

	if err != nil {
		log.Fatal(err)
		return false
	}

	return true

}

func (s *Server) Serve(w http.ResponseWriter, req *http.Request) {
	s.Proxy.ServeHTTP(w, req)

}

func NewServer(healthCheckPath string, address string) *Server {
	serverUrl, err := url.Parse("http://" + address)

	OnErrorPanic(err, "Invalid server address")

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Response from server: %v\n", address)

	})

	return &Server{
		healthCheckPath: serverUrl.String(),
		Address:         serverUrl,
		ServerMux:       mux,
		Proxy:           httputil.NewSingleHostReverseProxy(serverUrl),
	}
}
