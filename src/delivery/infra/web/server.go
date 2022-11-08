package web

import (
	"log"
	"net/http"

	"github.com/BinMunawir/mashi/src/delivery/infra/configs"
	"github.com/BinMunawir/mashi/src/delivery/infra/web/rest"
)

func InitializeHTTPServer() {
	initializerUsecases()
	router := rest.RegisterRoutes()

	HOST := configs.HOST
	PORT := configs.PORT
	log.Println("Server running on " + HOST + ":" + PORT)
	log.Fatal(http.ListenAndServe(HOST+":"+PORT, router))
}

func initializerUsecases() {
	// usecases.Init()
}
