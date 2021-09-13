package main

import (
	"net/http"
	"./filelisting"
	"os"
	"log"
)

type appHandler3 func(writer http.ResponseWriter, request *http.Request) error

func errWrapper3(handler appHandler3) func(http.ResponseWriter, *http.Request){
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			log.Printf("Error handling request: %s", err.Error())
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusBadRequest)
				return
			}

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

type userError interface {
	error
	Message() string
}

const path3 = "/"

func main() {
	http.HandleFunc(path3, errWrapper3(filelisting.HandleFileList3))

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
