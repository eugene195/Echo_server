package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"bytes"
	"flag"
)

const (
	CONN_HOST = ""
	CONN_PORT = "8000"
	CONN_TYPE = "tcp"
)

var (
	nWorkers = flag.Int("n", 10, "Number of workers to start with")
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

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	// Read the incoming connection into the buffer.
	reqLen, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	// Builds the message.
	message := "Hi, I received your message! It was "
	message += strconv.Itoa(reqLen)
	message += " bytes long and that's what it said: \""
	n := bytes.Index(buf, []byte{0})
	message += string(buf[:n-1])
	message += "\" ! Honestly I have no clue about what to do with your messages, so Bye Bye!\n"

	// Write the message in the connection channel.
	conn.Write([]byte(message));
	// Close the connection when you're done with it.
	conn.Close()
}