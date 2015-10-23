package main


import (
	"fmt"
	"strconv"
	"os"
	"net"
)

// NewWorker creates, and returns a new Worker object. Its only argument
// is a channel that the worker can add itself to whenever it is done its
// work.
func NewWorker(id int, workerQueue chan chan WorkRequest) Worker {
	// Create, and return the worker.
	worker := Worker{
		ID:          id,
		Work:        make(chan WorkRequest),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool)}

	return worker
}

type Worker struct {
	ID          int
	Work        chan WorkRequest
	WorkerQueue chan chan WorkRequest
	QuitChan    chan bool
}

// This function "starts" the worker by starting a goroutine, that is
// an infinite "for-select" loop.
func (worker Worker) Start() {
	go func() {
		for {
			// Add ourselves into the worker queue.
			worker.WorkerQueue <- worker.Work

			select {
				case work := <-worker.Work:
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
//					fmt.Println(response.Status)

					if (work.Method == "GET") {
						response.Body = file
					}

					var ext = GetExtByFileName(filename)
					response.Headers.Add("Content-Type", GetHeaderByExt(ext))
					response.Headers.Add("Content-Length", strconv.Itoa(len(file)))
					_, writeErr := work.Connection.Write(responseToByte(response))
					checkError(writeErr)
					work.Connection.Close()

				case <-worker.QuitChan:
				// We have been asked to stop.
					fmt.Printf("worker%d stopping\n", worker.ID)
					return
			}
		}
	}()
}

// Stop tells the worker to stop listening for work requests.
//
// Note that the worker will only stop *after* it has finished its work.
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

func checkError (err error) {
	if (err != nil) {
		fmt.Println(err)
		os.Exit(1)
	}
}

func writeAndClose (response Response, connection net.Conn) {
	_, writeErr := connection.Write(responseToByte(response))
	checkError(writeErr)
	connection.Close()
}