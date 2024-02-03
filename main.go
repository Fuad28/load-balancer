package main

import (
	"fmt"
	"strconv"
)

// const lbPort = ":80"

// func handleRoot(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, `
// 	Received request from: %v
// 	%v %v %v
// 	Host: %v
// 	User-Agent: %v
// 	Accept: %v
// 	`, req.RemoteAddr, req.Method, req.URL.Path, req.Proto, req.Host, req.UserAgent(), req.Header.Get("Accept"))

// }

func main() {
	lb := setupLoadBalancer()
	lb.Start()

}

var config LoadBalancerConfig

func setupLoadBalancer() *LoadBalancer {

	_ = config.LoadConfig()

	lb := LoadBalancer{
		Port:     config.Port,
		Count:    0,
		LastPort: 8000,
		Config:   config,
	}

	if config.Env == "dev" {
		devSetup(&lb)
	} else if config.Env == "prod" {
		ProdSetup(&lb)
	} else {
		panic("Invalid ENV value")
	}

	return &lb

}

func devSetup(lb *LoadBalancer) {
	baseUrl := "http://localhost:"

	for i := 0; i <= lb.Config.NoOfServers; i++ {
		nextPort := lb.getNextPort()
		address := baseUrl + strconv.Itoa(nextPort)
		healthCheckPath := address + "/"
		server := NewServer(healthCheckPath, address)
		lb.Servers = append(lb.Servers, server)
	}

}

func ProdSetup(lb *LoadBalancer) {
	fmt.Println("PROD SETUP")

	for _, s := range lb.Config.Servers {
		server := NewServer(s.HealthCheckPath, s.Address)
		lb.Servers = append(lb.Servers, server)
	}
}
