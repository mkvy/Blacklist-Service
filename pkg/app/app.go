package app

import (
	"github.com/mkvy/BlacklistTestTask/pkg/config"
	"github.com/mkvy/BlacklistTestTask/pkg/controller"
	"github.com/mkvy/BlacklistTestTask/pkg/database"
	"github.com/mkvy/BlacklistTestTask/pkg/repo"
	"github.com/mkvy/BlacklistTestTask/pkg/server"
	"github.com/mkvy/BlacklistTestTask/pkg/service"
	"log"
	"os"
	"os/signal"
)

func Run() {
	cfg := config.GetConfig()
	dbConn, err := database.NewDBConn(*cfg)
	if err != nil {
		log.Panic(err)
	}
	db, err := dbConn.GetDB()
	if err != nil {
		log.Panic(err)
	}
	repository := repo.NewDBBlacklistRepo(db)
	svc := service.NewBlacklistSvcImpl(repository)
	handler := controller.NewController(svc)
	s := server.NewServer(*cfg, handler)
	go s.Start()
	log.Println("http: Server is running")
	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, os.Interrupt, os.Kill)
	<-sigTerm
	s.Stop()
	dbConn.DBClose()
	log.Println("application finished")
}
