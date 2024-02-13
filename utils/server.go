package utils

import (
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

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
