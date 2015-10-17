package main

import "net"

type WorkRequest struct {
	// request method (i.e GET, POST ...)
	Method string

	// path to the resource
	Path string

	// Protocl version
	HTTPVersion string

	Connection net.Conn
}