package main

import (
	"fmt"
	"net"
	"os"
	"flag"
	"runtime"
)

const (
	CONN_HOST = ""
	CONN_PORT = "8000"
	CONN_TYPE = "tcp"
)

var (
	nWorkers = flag.Int("n", 15, "Number of workers to start with")
	workPath = flag.String("r", rootDir, "Root directory")
	nCPU = flag.Int("c", 4, "Number of CPU cores")
	port = flag.String("p", CONN_PORT, "Default port")
)


func main() {
	flag.Parse()
	runtime.GOMAXPROCS(*nCPU)
	config.rootDir = *workPath
	config.port = *port
	config.indexFile = indexFile

	StartDispatcher(*nWorkers)
	l, err := net.Listen(CONN_TYPE, ":" + config.port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + config.port)
	for {
		connextion, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
//		fmt.Printf("Received message %s -> %s \n", connextion.RemoteAddr(), connextion.LocalAddr())
		Collector(connextion, workPath)
	}
}