package rest

import (
	"net/http"

	"github.com/BinMunawir/mashi/src/core/usecases"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/api/skeleton_report.html", func(w http.ResponseWriter, r *http.Request) {
		w.Write(usecases.SkeletonReportHtml())
	})
	router.Get("/api/skeleton_report.json", func(w http.ResponseWriter, r *http.Request) {
		w.Write(usecases.SkeletonReportJson())
	})
	router.Get("/api/skeleton_report", func(w http.ResponseWriter, r *http.Request) {
		w.Write(usecases.SkeletonReportJson())
	})
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Mashi"))
	})

	return router
}
