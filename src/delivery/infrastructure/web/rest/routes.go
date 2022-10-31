package rest

import (
	"fmt"
	"net/http"

	usecases "github.com/BinMunawir/mashi/src/core/usecases/general"
	"github.com/BinMunawir/mashi/src/delivery/adaptors"
)

func RegisterRoutes() {
	http.HandleFunc("/api/welcome_report", func(w http.ResponseWriter, r *http.Request) {
		data := usecases.GenerateWelcomReport()
		html := adaptors.HTMLRender(data)
		w.Write(html)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Mashi")
	})
}
