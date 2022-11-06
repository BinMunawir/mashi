package rest

import (
	"encoding/json"
	"net/http"

	"github.com/BinMunawir/mashi/src/core/usecases"
	"github.com/go-chi/chi/v5"
)

func RegisterRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/api/reports/skeleton", func(w http.ResponseWriter, r *http.Request) {
		content, _ := usecases.SkeletonReport()
		json, _ := json.Marshal(content)
		w.Write(json)
	})
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Mashi"))
	})

	return router
}
