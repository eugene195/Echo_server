package main
import (
"net"
"strings"
	"net/url"
	"errors"
)

func (p *Pool) Start() {
	go func() {
		for {
			select {
			case connection, ok := <-p.tasks:
				if (!ok) {
					checkError(errors.New("Error during connection receival in worker"))
				}
				work, err := splitRequest(getRequest(connection))
				if err != nil {
					Debug(err.Error())
				} else {
					work.Connection = connection
					ProcessRequest(work)
				}
			case <-p.kill:
				return
			}
		}
	}()
}

func getRequest (conn net.Conn) (string) {
	var buf = make([]byte, 1024)
	_, err := conn.Read(buf)
	var requestStr = string(buf)
	if err != nil {
		Debug("Request can not be retrieved")
	}
	return requestStr
}


func splitRequest(query string) (*UnitOfWork, error) {
	parts := strings.Split(query, " ")
	request := new(UnitOfWork)
	request.Method = parts[0]
	if (len(parts) < 3) {
		return nil, errors.New("Request not valid")
	}
	uri := strings.Split(parts[1], "?")
	request.Path, _ = url.QueryUnescape(uri[0])
	request.HTTPVersion = parts[2]
	return request, nil
}