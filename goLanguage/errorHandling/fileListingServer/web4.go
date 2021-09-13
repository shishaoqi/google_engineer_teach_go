package main

import (
	"shishaoGo/goLanguage/errorHandling/fileListingServer/filelisting"
	"log"
	"net/http"
	"os"
)

type appHandler4 func(writer http.ResponseWriter, request *http.Request) error

func errWrapper3(handler appHandler4) func(http.ResponseWriter, *http.Request){
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

const path4 = "/list/"

func main() {
	http.HandleFunc(path4, errWrapper3(filelisting.HandleFileList))

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
