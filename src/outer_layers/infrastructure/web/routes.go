package web

import (
	"fmt"
	"net/http"

	"github.com/BinMunawir/mashi/src/inner_layers/usecases/general"
	"github.com/BinMunawir/mashi/src/outer_layers/adaptors"
)

func RegisterRoutes() {
	http.HandleFunc("/api/welcome_report", func(w http.ResponseWriter, r *http.Request) {
		data := general.GenerateWelcomReport()
		html := adaptors.HTMLRender(data)
		w.Write(html)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Mashi")
	})
}
