GOPATH := $(shell pwd)

all:		
	rm -f ./httpd	
	go build -o ./httpd ./src/*.go
	
