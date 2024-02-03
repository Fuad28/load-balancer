package main

import (
	"fmt"
	"strings"
)

/*
1. Build a load balancer that can send traffic to two or more servers.
2. Health check the servers.
3. Handle a server going offline (failing a health check).
4. Handle a server coming back online (passing a health check).
*/

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

	// waitGroup := sync.WaitGroup{}
	// fmt.Println("we are here: ", lbPort)

	// http.HandleFunc("/", handleRoot)

	// waitGroup.Add(1)
	// go func() {
	// 	defer waitGroup.Done()

	// 	fmt.Println("Starting server one")
	// 	http.ListenAndServe(lbPort1, ServerMuxOne())
	// }()

	// waitGroup.Add(1)
	// go func() {
	// 	defer waitGroup.Done()

	// 	fmt.Println("Starting lb server")
	// 	http.ListenAndServe(lbPort, nil)
	// }()

	// waitGroup.Wait()

}

var config LoadBalancerConfig

func setupLoadBalancer() {

	err := LoadFile[LoadBalancerConfig]("config.json", &config)

	if err != nil {
		panic("Couldn't load config file")
	}

	ENV := strings.ToLower(config.Env)
	var lb *LoadBalancer

	if ENV == "dev" {
		lb = devSetup()
	} else if ENV == "prod" {
		lb = ProdSetup()
	} else {
		panic("Invalid ENV value")
	}

	// start lb server

}

type LoadBalancer struct {
	Port     int
	lastPort int
	Count    int
	Servers  []*Server
	// Config T
}

func devSetup() *LoadBalancer {
	fmt.Println("DEV SETUP")

	baseUrl := "http://localhost:"
	lb := LoadBalancer{
		Port:     config.Port,
		Count:    0,
		lastPort: 8000,
	}

	for i := 0; i <= config.NoOfServers; i++ {
		nextPort := lb.getNextPort()
		address := baseUrl + string(nextPort)
		healthCheckPath := address + "/"
		server := NewServer(healthCheckPath, address)
		lb.Servers = append(lb.Servers, server)
	}

	return &lb

}

func ProdSetup() *LoadBalancer {
	fmt.Println("PROD SETUP")

	lb := LoadBalancer{
		Port:  config.Port,
		Count: 0,
	}

	for s := range config.Servers {
		server := NewServer(s.healthCheckPath, s.address)
		lb.Servers = append(lb.Servers, server)
	}

	return &lb
}
