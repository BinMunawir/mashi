package rest

import (
	"fmt"
	"net/http"

	usecases "github.com/BinMunawir/mashi/src/core/usecases"
)

func RegisterRoutes() {
	http.HandleFunc("/api/skeleton_report", func(w http.ResponseWriter, r *http.Request) {
		w.Write(usecases.SkeletonReportHtml())
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Mashi")
	})
}
