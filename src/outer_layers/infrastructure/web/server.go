package web

import (
	"log"
	"net/http"
)

func InitializeHTTPServer() {
	log.Println("Server running on localhost:8080")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
