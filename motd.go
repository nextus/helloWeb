package main

import (
	"crypto/sha256"
	"fmt"
	"net"
	"time"
)

var updateTime = time.Hour * 24 // how often generate new motd
var motd string                 // IPC variable to pass through motd from generator to web server gorotine
// TODO: use channels instead of global variable

// get all alive ip addresses from the host and return the first one. Return loopback address if empty
func getServerIP() string {
	// TODO: if we bind to specific address we could return it immediately
	// TODO: clearify what address we should return exactly
	// BUG: returning result unordered right now
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		Warning.Printf("[Motd] Unable to get IP address: %v", err)
	}
	loopbackAddr := "[unknown address]"
	for _, addr := range addrs {
		ipnet, ok := addr.(*net.IPNet)
		if !ok {
			Warning.Printf("[Motd] %v", err)
			continue
		}
		if ipnet.IP.IsLoopback() {
			loopbackAddr = ipnet.IP.String()
		} else {
			return ipnet.IP.String()
		}
	}
	return loopbackAddr
}

// generate new message of the day
func generateMotd() string {
	// init
	dt := time.Now()
	hash := sha256.New()
	// date
	currentDateTime := dt.Format("2006-01-02 15:04:05 MST")
	// hash
	hash.Write([]byte(currentDateTime))
	// ip
	ipAddr := getServerIP()
	return fmt.Sprintf("%s %s %x", currentDateTime, ipAddr, hash.Sum(nil))
}

// motd watchdog
func updateMotd() {
	for {
		time.Sleep(updateTime)
		motd = generateMotd()
	}
}

// init motd server
func startMotd() {
	motd = generateMotd()
	go updateMotd()
}
