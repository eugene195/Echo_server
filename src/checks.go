package main

import (
	"strings"
	"io/ioutil"
)

func methodCheck (checkedMethod string) (string, string) {
	var errorString = NOT_ALLOWED
	if (checkedMethod != "POST") {
		return checkedMethod, ""
	}
	return "", errorString
}

func urlSecurityCheck (checkedUrl string) (string, string) {
	var errorString = FORBIDDEN
	if (strings.Contains(checkedUrl, "../")) {
		return "", errorString
	}
	return checkedUrl, ""
}

func fileNotFoundCheck (path string) (string, string, string) {
	var errorString = NOT_FOUND
	file, err := ioutil.ReadFile(config.rootDir + path)
	if err == nil {
		return string(file[:]), "", config.rootDir + path
	}
	file, _ = ioutil.ReadFile(config.rootDir + "/httptest/" + notFoundFile)
	return string(file[:]), errorString, config.rootDir + "/httptest/" + notFoundFile
}

func indexFileCheck (path string) (string, string, string) {
	var errorString = FORBIDDEN
	if (IsDirectory(path)) {
		file, err := ioutil.ReadFile(config.rootDir + path + indexFile)
		if err != nil {
			return "", errorString, ""
		}
		return string(file[:]), "", config.rootDir + path + indexFile
	}
	return "", "", ""
}