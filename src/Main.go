package main

import (
	"fmt"
	"net"
	"flag"
	"runtime"
	"errors"
)

const (
	CONN_HOST = ""
	CONN_PORT = "8000"
	CONN_TYPE = "tcp"
)

var (
	nWorkers = flag.Int("n", 20, "Number of workers to start with")
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

	l, err := net.Listen(CONN_TYPE, ":" + config.port)
	if err != nil {
		checkError(errors.New("Error listening"))
	}
	defer l.Close()
	fmt.Println("Listening on " + CONN_HOST + ":" + config.port)

	pool := NewPool(*nWorkers)

	for {
		connection, err := l.Accept()
		if err != nil {
			checkError(errors.New("Error accepting"))
		}
		pool.Exec(connection)
	}
}
