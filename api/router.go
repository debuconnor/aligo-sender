package api

import (
	"log"
	"net/http"
)

type request func(http.ResponseWriter, *http.Request)

func HandleRequest(path string, req request) {
	http.HandleFunc(path, req)
}

func Listen() {
	log.Fatal(http.ListenAndServe(":8090", nil))
}
