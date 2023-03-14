package handler

import (
	"github.com/gorilla/mux"
	"github.com/mkvy/BlacklistTestTask/docs"
	"github.com/mkvy/BlacklistTestTask/internal/auth"
	"github.com/mkvy/BlacklistTestTask/internal/controller"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"
	"net/http"
)

// @title Blacklist API
// @version 1.0
// @description This is a service that manages blacklist of users
// @contact.name Paul Mkvts
// @BasePath /api/v1
// @securityDefinitions.basic BasicAuth
func NewHandler(c *controller.Controller) *mux.Router {
	router := mux.NewRouter()
	auth.SetupGoGuardian()
	router.HandleFunc("/api/v1/auth/token", auth.Middleware(http.HandlerFunc(auth.CreateToken))).Methods("GET")
	router.HandleFunc("/api/v1/blacklist", auth.Middleware(http.HandlerFunc(c.AddHandler))).Methods("POST")
	router.HandleFunc("/api/v1/blacklist/", auth.Middleware(http.HandlerFunc(c.AddHandler))).Methods("POST")
	router.HandleFunc("/api/v1/blacklist/{id}", auth.Middleware(http.HandlerFunc(c.DeleteHandler))).Methods("DELETE")
	router.HandleFunc("/api/v1/blacklist", auth.Middleware(http.HandlerFunc(c.GetByUsernameHandler))).Queries("user_name", "{user_name}").Methods("GET")
	router.HandleFunc("/api/v1/blacklist", auth.Middleware(http.HandlerFunc(c.GetByPhoneHandler))).Queries("phone_number", "{phone_number}").Methods("GET")
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	log.Println("swagger running at /swagger/index.html host: ", docs.SwaggerInfo.Host)
	return router
}
