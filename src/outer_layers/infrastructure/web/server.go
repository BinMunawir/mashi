package web

import (
	"log"
	"net/http"
	"os"
)

func InitializeHTTPServer() {
	RegisterRoutes()
	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
	log.Println("Server running on " + HOST + ":" + PORT)
	log.Fatal(http.ListenAndServe(HOST+":"+PORT, nil))
}
