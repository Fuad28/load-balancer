package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type LoadBalancer struct {
	Port     int
	LastPort int
	Count    int
	Servers  []*Server
	Config   LoadBalancerConfig
}

func (lb *LoadBalancer) getNextServer() *Server {

	lb.Count++
	server := lb.Servers[lb.Count%len(lb.Servers)]

	if !server.IsAlive() || SimulateDownServer(lb) {
		lb.Count++
		lb.getNextServer()
	}

	return server
}

func portInuse(port int) bool {
	_, err := http.Get("http://localhost:" + strconv.Itoa(port))
	return err == nil
}

func (lb *LoadBalancer) getNextPort() int {

	lb.LastPort++

	for {
		if portInuse(lb.LastPort) {
			lb.LastPort++
		} else {
			return lb.LastPort
		}
	}
}

func (lb *LoadBalancer) Start() {
	waitGroup := sync.WaitGroup{}

	if lb.Config.Env == "dev" {
		lb.StartDemoServers(&waitGroup)
		lb.StartLB(&waitGroup)

		waitGroup.Wait()
	} else {
		lb.StartLB(&waitGroup)
	}

}

func (lb *LoadBalancer) StartDemoServers(waitGroup *sync.WaitGroup) {
	for _, server := range lb.Servers {
		waitGroup.Add(1)

		go func(server *Server) {
			defer waitGroup.Done()
			http.ListenAndServe(server.Address.Host, server.ServerMux)
		}(server)
	}

}

func (lb *LoadBalancer) StartLB(waitGroup *sync.WaitGroup) {
	waitGroup.Add(1)

	go func() {
		defer waitGroup.Done()

		http.HandleFunc("/", lb.Serve)
		http.ListenAndServe(":"+strconv.Itoa(lb.Port), nil)
	}()

}

func (lb *LoadBalancer) Serve(w http.ResponseWriter, req *http.Request) {
	server := lb.getNextServer()

	fmt.Printf("Sending request to server %v\n", server.Address)

	server.Serve(w, req)

}
