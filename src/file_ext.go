package main

import "strings"

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
	"empty": "text/html; charset=utf-8",
}

func GetHeaderByExt(ext string) string {
	cType, ok := exts[ext]
	if ok {
		return cType
	} else {
		return exts["empty"]
	}
}

func getHeader (extension string) string {
	content_type, ok := exts[extension]
	if (ok) {
		return content_type
	} else {
		return exts["empty"]
	}
}

func GetExtByFileName(name string) string {
	parts := strings.Split(name, ".")
	return parts[len(parts)-1]
}

func IsDirectory(path string) bool {
	if path[len(path)-1:] == "/" {
		return true
	}
	return false
}