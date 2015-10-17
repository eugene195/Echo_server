package main


import (
	"fmt"
	"net"
//	"errors"
	"strings"
	"net/url"
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

	work, err := parseRequest(requestStr)
	if err != nil {
		fmt.Println("Request not parsed")
		return
	}
	work.Connection = connection

	// Push the work onto the queue.
	WorkQueue <- *work
	fmt.Println("Work request queued")
	return
}


func parseRequest(query string) (*WorkRequest, error) {
	parts := strings.Split(query, " ")
	request := new(WorkRequest)

	uri := strings.Split(parts[1], "?")

	path, _ := url.QueryUnescape(uri[0])

	request.Method = parts[0]
	if parts[1] == "/" {
		request.Path = "index.html"
	} else {
		request.Path = path
	}
	request.HTTPVersion = parts[2]
	return request, nil
}