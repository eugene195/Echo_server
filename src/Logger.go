package main
import (
	"fmt"
	"os"
	"runtime"
	"log"
)

func checkError(err error) {
	if (err != nil) {
		fmt.Println("Fatal error happened, turning off: ")
		fmt.Println(err)
		os.Exit(1)
	}
}

func Debug(format string, a ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	info := fmt.Sprintf(format, a...)
	log.Printf("[cgl] debug %s:%d %v", file, line, info)
}