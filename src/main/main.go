package main

import (
	"flag"
	"os"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listens of the specified ip address")
	flag.Parse()

	if isHost {
		// go run main-go -listen <ip>
		connIP := os.Args[2]
		runHost(connIP)
	} else {
		// go run main.go <ip>
		connIP := os.Args[1]
		runGuest(connIp)
	}
}

func runHost(ip string) {

}

func runGuest(ip string) {

}
