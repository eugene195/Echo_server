package main


import (
	"fmt"
	"net"
	"errors"
	"strings"
	"net/url"
//	"extension"
)

// A buffered channel that we can send work requests on.
var WorkQueue = make(chan WorkRequest, 100)

func Collector(connection net.Conn, workDir *string) {
	var buf = make([]byte, 1024)
	_, err := connection.Read(buf)
	var requestStr = string(buf)
	if err != nil {
		fmt.Println("Request can not be retrieved")
		return
	}

	work, err := splitRequest(requestStr)
	if err != nil {
		fmt.Println("Request not parsed")
		return
	}
	work.Connection = connection

	// Push the work onto the queue.
	WorkQueue <- *work
//	fmt.Println("Work request queued")
	return
}


func splitRequest(query string) (*WorkRequest, error) {
	parts := strings.Split(query, " ")
	request := new(WorkRequest)
	request.Method = parts[0]
	if (len(parts) < 3) {
		return nil, errors.New("Request not valid")
	}

	uri := strings.Split(parts[1], "?")
//	fmt.Println(parts[1])
	request.Path, _ = url.QueryUnescape(uri[0])

	request.HTTPVersion = parts[2]
	return request, nil
}