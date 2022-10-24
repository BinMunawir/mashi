package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BinMunawir/mashi/src/inner_layers/usecases/general"
)

func RegisterRoutes() {
	http.HandleFunc("/api/welcome_report", func(w http.ResponseWriter, r *http.Request) {
		data := general.GenerateWelcomReport()
		json.NewEncoder(w).Encode(data)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Mashi")
	})
}
