package main

import (
	"bytes"
	"time"
	"fmt"
)

type Headers map[string]string

type Response struct {
	Status  string
	Proto   string
	Body    string
	Headers Headers
}

var EmptyHeaders = Headers{}

func responseToByte(response Response) []byte {
	var result bytes.Buffer
	result.WriteString(response.Proto + " " + response.Status + stringSeparator)
	result.WriteString(response.Headers.ToPlainData() + stringSeparator)
	result.WriteString(response.Body)
	return result.Bytes()
}

func responseToString (response Response) string {
	var result = ""
	result += response.Proto + " " + response.Status + stringSeparator
	result += response.Headers.ToPlainData() + stringSeparator
	result += response.Body
	return result
}

func (h Headers) Add(key string, value string) {
	h[key] = value
}

func (h Headers) Get(key string) string {
	return h[key]
}

func (headers Headers) ToPlainData() string {
	var result bytes.Buffer
	for key, value := range headers {
		_, err := result.WriteString(key + headerSeparator + value + stringSeparator)
		if (err != nil) {
			fmt.Println(err)
		}
	}
	strResult := result.String()
	return strResult
}

func setDefaultHeaders (response Response) Response {
	response.Headers.Add("Server", "echo-server")
	response.Headers.Add("Date", time.Now().Format(time.RFC822))
	response.Headers.Add("Connection", "close")
	return response
}