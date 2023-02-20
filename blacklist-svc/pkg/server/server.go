package server

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/config"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/controller"
	"log"
	"net/http"
)

type Server struct {
	srv  *http.Server
	addr string
}

func NewServer(cfg config.Config, c *controller.Controller) *Server {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/test/", c.AddHandler).Methods("POST")
	router.HandleFunc("/api/v1/test/{id}", c.DeleteHandler).Methods("DELETE")
	router.HandleFunc("/api/v1/testGet/", c.GetByUsernameHandler).Methods("POST")
	router.HandleFunc("/api/v1/testGetPhone/", c.GetByPhoneHandler).Methods("POST")
	server := &http.Server{Addr: ":" + cfg.HttpServer.Port, Handler: router}
	return &Server{server, cfg.HttpServer.Port}
}
func (s *Server) Start() {
	if err := s.srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func (s *Server) Stop() {
	s.srv.Shutdown(context.Background())
}
