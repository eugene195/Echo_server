package main


import (
	"fmt"
	"net"
	"bufio"
	"errors"
	"strings"
	"net/textproto"
)

// A buffered channel that we can send work requests on.
var WorkQueue = make(chan WorkRequest, 100)



func Collector(conn net.Conn) {
	// Now, we take the delay, and the person's name, and make a WorkRequest out of them.
	work, err := parseRequest(bufio.NewReader(conn), conn)

	if err != nil {
		fmt.Println("Request not parsed")
		return
	}

	// Push the work onto the queue.
	WorkQueue <- *work
	fmt.Println("Work request queued")

	// And let the user know their work request was created.
	return
}


func parseRequest(reader *bufio.Reader, conn net.Conn) (*WorkRequest, error) {

	var r = textproto.NewReader(reader)

	// create new request object
	request := new(WorkRequest)

	methodLine, _ := r.ReadLine()
	methodLineElements := strings.Split(methodLine, " ")

	if (len(methodLineElements) != 3) {
		return request, errors.New("Invalid request")
	}

	request.Method = methodLineElements[0]

	if methodLineElements[1] == "/" {
		request.Path = "index.html"
	} else {
		request.Path = methodLineElements[1]
	}
	request.HTTPVersion = methodLineElements[2]
	request.Connection = conn
	return request, nil
}