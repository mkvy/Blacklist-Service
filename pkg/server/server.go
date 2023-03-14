package server

import (
	"context"
	_ "github.com/mkvy/BlacklistTestTask/docs"
	"github.com/mkvy/BlacklistTestTask/internal/config"
	"log"
	"net/http"
)

type Server struct {
	srv  *http.Server
	addr string
}

// @title Blacklist API
// @version 1.0
// @description This is a service that manages blacklist of users
// @contact.name Paul Mkvts
// @BasePath /api/v1
// @securityDefinitions.basic BasicAuth
func NewServer(cfg config.Config, handler http.Handler) *Server {
	server := &http.Server{Addr: ":" + cfg.HttpServer.Port, Handler: handler}
	return &Server{server, cfg.HttpServer.Port}
}
func (s *Server) Start() {
	log.Println("starting server at port: ", s.addr)
	if err := s.srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func (s *Server) Stop() {
	s.srv.Shutdown(context.Background())
}
