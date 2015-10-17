package main

import (
	"strings"
	"net/url"
)

const (
	OK           string = "200 OK"
	NOT_FOUND    string = "404 NOT FOUND"
	ERROR        string = "500 INTERNAL SERVER ERROR"
	BAD_REQUEST  string = "400 BAD REQUEST"
	NOT_ALLOWED  string = "405 METHOD NOT ALLOWED"
	FORBIDDEN    string = "403 FORBIDDEN"
	DEFAULT_FILE string = "/index.html"
	FILE_404     string = "/404.html"
	HTTP_VERSION string = "1.1"
)

func GetStatusLine(status string) string {
	return "HTTP/" + HTTP_VERSION + " " + status + "\r\n"
}

var (
	httpMethods = []string{
		"OPTIONS", "GET", "HEAD", "POST",
		"PUT", "PATCH", "DELETE", "TRACE", "CONNECT",
	}
	supportedMethods = []string{"GET", "HEAD"}
)

var exts = map[string]string{
	"txt":  "application/text",
	"html": "text/html",
	"json": "application/json",
	"jpg":  "image/jpeg",
	"jpeg": "image/jpeg",
	"png":  "image/png",
	"js":   "text/javascript",
	"css":  "text/css",
	"gif":  "image/gif",
	"swf":  "application/x-shockwave-flash",
}

func GetHeaderByExt(ext string) string {
	cType, ok := exts[ext]
	if ok {
		return ("Content-Type: " + cType)
	} else {
		return "Content-Type: text/html; charset=utf-8"
	}
}

func GetExtByFileName(name string) string {
	parts := strings.Split(name, ".")
	return parts[len(parts)-1]
}

func isHttpMethod(method string) bool {
	for _, v := range httpMethods {
		if method == v {
			return true
		}
	}
	return false
}

func isSupportedMethod(method string) bool {
	for _, v := range supportedMethods {
		if method == v {
			return true
		}
	}
	return false
}

func IsDirectory(path string) bool {
	if path[len(path)-1:] == "/" {
		return true
	}
	return false
}

func ParseQueryString(query string) map[string]string {
	parts := strings.Split(query, " ")

	uri := strings.Split(parts[1], "?")

	path, _ := url.QueryUnescape(uri[0])

	return map[string]string{
		"method": parts[0],
		"path":   path,
	}
}