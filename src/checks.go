package main

import (
	"status"
	"strings"
	"io/ioutil"
)

func methodCheck (checkedMethod string) (string, string) {
	var errorString = status.NOT_ALLOWED
	if (checkedMethod != "POST") {
		return checkedMethod, ""
	}
	return "", errorString
}

func urlSecurityCheck (checkedUrl string) (string, string) {
	var errorString = status.FORBIDDEN
	if (!strings.Contains(checkedUrl, "../")) {
		return checkedUrl, ""
	}
	return "", errorString
}

func fileNotFoundCheck (path string) (string, string, string) {
	var errorString = status.NOT_FOUND
	file, err := ioutil.ReadFile(rootDir + path)
	if err == nil {
		return string(file[:]), "", rootDir + path
	}
	file, _ = ioutil.ReadFile(rootDir + "/httptest/" + notFoundFile)
	return string(file[:]), errorString, rootDir + "/httptest/" + notFoundFile
}

func indexFileCheck (path string) (string, string, string) {
	var errorString = status.FORBIDDEN
	if (IsDirectory(path)) {
		file, err := ioutil.ReadFile(rootDir + path + indexFile)
		if err != nil {
			return "", errorString, ""
		}
		return string(file[:]), "", rootDir + path + indexFile
	}
	return "", "", ""
}