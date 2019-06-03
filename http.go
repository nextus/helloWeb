package main

import (
	"fmt"
	"net/http"
)

const (
	httpDefaultListenHost = ""
	httpDefaultListenPort = 8091
	httpDefaultLocation   = "/"
	httpRobotsLocation    = "/robots.txt"
	// TODO: put robots.txt out from the source code
	httpRobotsTxt = `User-agent: *
Disallow: /
`
)

func startWebServer(addr, location string) {
	mux := http.NewServeMux()
	mux.HandleFunc(location, publicLocationHandler)
	mux.HandleFunc(httpRobotsLocation, robotsLocationHandler)
	srv := &http.Server{Addr: addr, Handler: mux}
	Warning.Printf("[HTTP server] Listening %s", addr)
	if err := srv.ListenAndServe(); err != nil {
		Fatal.Fatalf("[HTTP server] Critical error: %v", err)
	}
}

func publicLocationHandler(w http.ResponseWriter, r *http.Request) {
	Info.Printf("[HTTP server] %s %s %s", r.RemoteAddr, r.Method, r.RequestURI)
	fmt.Fprintf(w, motd)
}

func robotsLocationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, httpRobotsTxt)
}
