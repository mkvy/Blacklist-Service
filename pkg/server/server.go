package server

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/mkvy/BlacklistTestTask/docs"
	_ "github.com/mkvy/BlacklistTestTask/docs"
	"github.com/mkvy/BlacklistTestTask/pkg/auth"
	"github.com/mkvy/BlacklistTestTask/pkg/config"
	"github.com/mkvy/BlacklistTestTask/pkg/controller"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net"
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
func NewServer(cfg config.Config, c *controller.Controller) *Server {
	router := mux.NewRouter()
	auth.SetupGoGuardian()
	docs.SwaggerInfo.Host = net.JoinHostPort(cfg.HttpServer.Host, cfg.HttpServer.Port)
	router.HandleFunc("/api/v1/auth/token", auth.Middleware(http.HandlerFunc(auth.CreateToken))).Methods("GET")
	router.HandleFunc("/api/v1/blacklist", auth.Middleware(http.HandlerFunc(c.AddHandler))).Methods("POST")
	router.HandleFunc("/api/v1/blacklist/", auth.Middleware(http.HandlerFunc(c.AddHandler))).Methods("POST")
	router.HandleFunc("/api/v1/blacklist/{id}", auth.Middleware(http.HandlerFunc(c.DeleteHandler))).Methods("DELETE")
	router.HandleFunc("/api/v1/blacklist", auth.Middleware(http.HandlerFunc(c.GetByUsernameHandler))).Queries("user_name", "{user_name}").Methods("GET")
	router.HandleFunc("/api/v1/blacklist", auth.Middleware(http.HandlerFunc(c.GetByPhoneHandler))).Queries("phone_number", "{phone_number}").Methods("GET")

	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Println("swagger running at /swagger/index.html host: ", docs.SwaggerInfo.Host)
	server := &http.Server{Addr: ":" + cfg.HttpServer.Port, Handler: router}
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
