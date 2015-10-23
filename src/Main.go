package main

import (
	"fmt"
	"net"
	"os"
	"flag"
)

const (
	CONN_HOST = ""
	CONN_PORT = "8000"
	CONN_TYPE = "tcp"
)

var (
	nWorkers = flag.Int("n", 15, "Number of workers to start with")
	workPath = flag.String("r", "/", "Root directory")
	nCPU = flag.Int("c", 1, "Number of CPU cores")
)


func main() {
	flag.Parse()
	StartDispatcher(*nWorkers)
	l, err := net.Listen(CONN_TYPE, ":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	for {
		connextion, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Received message %s -> %s \n", connextion.RemoteAddr(), connextion.LocalAddr())
		Collector(connextion, workPath)
	}
}