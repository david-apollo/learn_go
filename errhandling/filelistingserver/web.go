package main

import (
	"learn_go/errhandling/filelistingserver/filelisting"
	"net/http"
)


type appHandler func(writer http.ResponseWriter, request *http.Request) error


func errWrapper(handler appHandler) func (writer http.ResponseWriter, request *http.Request)  {
	
} {
	
}


func main() {
	http.HandleFunc("/list/", filelisting.HandleFileList)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}