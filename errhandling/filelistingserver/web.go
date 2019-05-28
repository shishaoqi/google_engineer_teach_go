package main

import (
	"net/http"
	"os"
	"io/ioutil"
)

const path0 = "/list/"

func main() {
	http.HandleFunc(path0,
		func(writer http.ResponseWriter, request *http.Request) {
			path := request.URL.Path[len("/list/"):] //  /list/google_learngo/
			file, err := os.Open(path)
			if err != nil {
				panic(err)
			}
			defer file.Close()

			all, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}

			writer.Write(all)
	})

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		panic(err)
	}
}
