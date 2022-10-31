package web

import (
	"log"
	"net/http"
	"os"

	"github.com/BinMunawir/mashi/src/delivery/infra/web/rest"
)

func InitializeHTTPServer() {
	rest.RegisterRoutes()
	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
	log.Println("Server running on " + HOST + ":" + PORT)
	log.Fatal(http.ListenAndServe(HOST+":"+PORT, nil))
}
