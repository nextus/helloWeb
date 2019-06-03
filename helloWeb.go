package main

import (
	"flag"
	"fmt"
)

const (
	defaultVerboseLevel = 3 // Warning level
)

var httpAddr string     // bind address
var httpLocation string // main location URL path

func init() {
	verboseLevel := flag.Uint("verbose", defaultVerboseLevel, "verbose level")
	httpHost := flag.String("host", httpDefaultListenHost, "listening host")
	httpPort := flag.Uint("port", httpDefaultListenPort, "listening port")
	flag.StringVar(&httpLocation, "location", httpDefaultLocation, "listening HTTP location")
	flag.Parse()
	setVerboseLevel(*verboseLevel)
	httpAddr = fmt.Sprintf("%s:%d", *httpHost, *httpPort)
}

func main() {
	// do not start web server until the motd is generated
	startMotd()
	go startWebServer(httpAddr, httpLocation)
	select {}
}
