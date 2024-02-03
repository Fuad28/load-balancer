package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
create server based on a given port
request server's ping url to check if alive
default handler for each server

*/

type Server struct {
	healthCheckPath string
	// Port            int
	Address   string
	ServerMux http.Handler
}

// type LoadBalancerDevConfig struct {

// }

type LoadBalancer struct {
	Port     int
	lastPort int
	Count    int
	Servers  []*Server
	// Config T
}

func (lb *LoadBalancer) getNextServer() *Server {

	lb.Count++
	server := lb.Servers[lb.Count%len(lb.Servers)]

	if !server.IsAlive() {
		lb.Count++
		lb.getNextServer()
	}

	return server
}

func portInuse(port int) bool {
	_, err := http.Get("http://localhost:" + string(port))

	if err != nil {
		return true
	}

	return false
}

func (lb *LoadBalancer) getNextPort() int {

	lb.lastPort++

	for {
		if portInuse(lb.lastPort) {
			lb.lastPort++
		} else {
			return lb.lastPort
		}
	}
}

func NewServer(healthCheckPath string, address string) *Server {

	mux := http.NewServeMux()

	mux.HandleFunc(healthCheckPath, handleServer1)
	mux.HandleFunc("/", handleServer1)

	return &Server{
		// Port:            port,
		healthCheckPath: healthCheckPath,
		Address:         address,
		ServerMux:       mux,
	}
}

func (s *Server) IsAlive() bool {
	res, err := http.Get(s.healthCheckPath)

	if res.StatusCode == 200 {
		return true
	}

	if err != nil {
		log.Fatal(err)
		return false
	}

	return false

}

const lbPort1 = ":81"

func handleServer1(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Response from server 1")

}

func ServerMuxOne() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handleServer1)

	return mux
}
