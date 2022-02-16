package public

import (
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
		router.Route("/auth", func(router chi.Router) {
			router.Post("/login", s.Login)
			router.Get("/check-token", s.CheckToken)
			router.Post("/upsert-user", s.UpsertUser)
		})
	})
}
