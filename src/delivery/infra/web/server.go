package web

import (
	"log"
	"net/http"
	"os"

	"github.com/BinMunawir/mashi/src/core/usecases"
	"github.com/BinMunawir/mashi/src/delivery/adaptors"
	"github.com/BinMunawir/mashi/src/delivery/infra/web/rest"
)

func InitializeHTTPServer() {
	initializerUsecases()
	rest.RegisterRoutes()

	HOST := os.Getenv("HOST")
	PORT := os.Getenv("PORT")
	log.Println("Server running on " + HOST + ":" + PORT)
	log.Fatal(http.ListenAndServe(HOST+":"+PORT, nil))
}

func initializerUsecases() {
	htmlRenderer := adaptors.NewGolangHtmlRenderer()
	usecases.Init(htmlRenderer)
}
