package main

import (
	"fmt"
	"net/http"
	"sync"
)

/*
1. Build a load balancer that can send traffic to two or more servers.
2. Health check the servers.
3. Handle a server going offline (failing a health check).
4. Handle a server coming back online (passing a health check).
*/

const lbPort = ":80"

func handleRoot(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, `
	Received request from: %v
	%v %v %v
	Host: %v
	User-Agent: %v
	Accept: %v
	`, req.RemoteAddr, req.Method, req.URL.Path, req.Proto, req.Host, req.UserAgent(), req.Header.Get("Accept"))

}

func main() {
	waitGroup := sync.WaitGroup{}
	fmt.Println("we are here: ", lbPort)

	http.HandleFunc("/", handleRoot)
	// http.HandleFunc("/", handleServer1)

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()

		fmt.Println("Starting server one")
		http.ListenAndServe(lbPort1, ServerMuxOne())
	}()

	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()

		fmt.Println("Starting lb server")
		http.ListenAndServe(lbPort, nil)
	}()

	waitGroup.Wait()

}
