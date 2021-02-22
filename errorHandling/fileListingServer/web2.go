package main

import (
	"github.com/shishao/hello/errorHandling/fileListingServer/filelisting"
	"net/http"
	"os"
	"log"
)

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request){
	return func(writer http.ResponseWriter, request *http.Request) {
		err := handler(writer, request)
		if err != nil {
			log.Printf("Error handling request: %s", err.Error())
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

const path2 = "/list/"

func main() {
	http.HandleFunc(path2, errWrapper(filelisting.HandleFileList))

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		panic(err)
	}
}
