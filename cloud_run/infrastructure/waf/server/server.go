package server

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"line-halo-stats/cloud_run/infrastructure/waf/config"
	"line-halo-stats/cloud_run/infrastructure/waf/handler"
	"line-halo-stats/cloud_run/infrastructure/waf/router"
	"log"
	"net/http"
	"os"
)

type server struct {
	config *config.ServerConfig
	r      chi.Router
	port   string
}

func NewServer() *server {
	return &server{}
}

func (s *server) Run() {
	s.config = s.GetServerConfig()
	s.r = chi.NewRouter()

	s.Start()
}

func (s *server) Start() {
	s.setMiddleware()
	s.setCors()
	s.setRoutes()
	s.setPort()

	// Start server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", s.port), s.r))
}

func (s *server) GetServerConfig() *config.ServerConfig {
	return config.NewServerConfig()
}

func (s *server) setMiddleware() {
	s.r.Use(middleware.RequestID)
	s.r.Use(middleware.Logger)
	s.r.Use(middleware.Recoverer)
	s.r.Use(middleware.URLFormat)
	s.r.Use(middleware.SetHeader("Cache-Control", "no-store"))
	s.r.Use(middleware.SetHeader("Strict-Transport-Security", "max-age=2592000"))
	s.r.Use(render.SetContentType(render.ContentTypeJSON))
}

func (s *server) setCors() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	s.r.Use(c.Handler)
}

func (s *server) setRoutes() {
	h := router.InitializePlayerStatsHandler(s.config)
	s.r.Route(handler.APIPrefix, func(r chi.Router) {
		r.Route(handler.PlayerStatsBasePath, h.PlayerStats)
	})
}

func (s *server) setPort() {
	// Choose port
	s.port = os.Getenv("PORT")
	if s.port == "" {
		s.port = "8080"
		log.Printf("Defaulting to port %s\n", s.port)
	}
	log.Printf("Listening on port %s\n", s.port)
}
