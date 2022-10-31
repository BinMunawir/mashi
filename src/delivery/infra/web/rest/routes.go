package rest

import (
	"fmt"
	"net/http"

	usecases "github.com/BinMunawir/mashi/src/core/usecases/skeleton"
	"github.com/BinMunawir/mashi/src/delivery/adaptors"
)

func RegisterRoutes() {
	http.HandleFunc("/api/skeleton_report", func(w http.ResponseWriter, r *http.Request) {
		data := usecases.SkeletonReport()
		html := adaptors.HTMLRender(data)
		w.Write(html)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Mashi")
	})
}
