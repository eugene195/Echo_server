package main

var (
	httpMethods = []string{
		"OPTIONS", "GET", "HEAD", "POST",
		"PUT", "PATCH", "DELETE", "TRACE", "CONNECT",
	}
	supportedMethods = []string{"GET", "HEAD"}
)


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