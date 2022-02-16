package public_fetch

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

func (s *Server) routes() {
	s.router.Use(cors.New(cors.Options{
		AllowedOrigins:     []string{"*"},
		AllowedHeaders:     s.Cfg.Cors.AllowedHeaders,
		AllowCredentials:   s.Cfg.Cors.AllowCredentials,
		Debug:              s.Cfg.Cors.Debug,
		MaxAge:             s.Cfg.Cors.MaxAge,
		OptionsPassthrough: s.Cfg.Cors.OptionsPassthrough,
	}).Handler)

	s.router.Route("/v1", func(router chi.Router) {
		router.Route("/fetch", func(router chi.Router) {
			router.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "It Works")
			})
			router.Get("/resources", s.GetResources)
			router.Get("/resources/admin", s.GetResourcesAdmin)
			router.Get("/check-token", s.CheckToken)
		})
	})
}
