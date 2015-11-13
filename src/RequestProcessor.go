package main
import (
	"strconv"
	"net"
)

func ProcessRequest (work *UnitOfWork) {
	response := Response{
		Status: OK,
		Proto: httpProto,
		Body: "",
		Headers: make(Headers) }

	setDefaultHeaders(response)
	var file, filename, result = "", "", ""
	_, result = methodCheck(work.Method)
	if (result != "") {
		response.Status = result
		writeAndClose(response, work.Connection)
		return
	}

	_, result = urlSecurityCheck(work.Path)
	if (result != "") {
		response.Status = result
		writeAndClose(response, work.Connection)
		return
	}

	file, result, filename = indexFileCheck(work.Path)
	if (result != "") {
		response.Status = result
		_, writeErr := work.Connection.Write(responseToByte(response))
		work.Connection.Close()
		checkError(writeErr)
		return
	}

	if (file == "") {
		file, result, filename = fileNotFoundCheck(work.Path)
	}

	if (result != "") {
		response.Status = result
	}

	if (work.Method == "GET") {
		response.Body = file
	}

	var ext = GetExtByFileName(filename)
	response.Headers.Add("Content-Type", GetHeaderByExt(ext))
	response.Headers.Add("Content-Length", strconv.Itoa(len(file)))
	_, writeErr := work.Connection.Write(responseToByte(response))
	checkError(writeErr)
	work.Connection.Close()
}

func writeAndClose (response Response, connection net.Conn) {
	_, writeErr := connection.Write(responseToByte(response))
	checkError(writeErr)
	connection.Close()
}