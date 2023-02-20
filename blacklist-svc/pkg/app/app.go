package app

import (
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/config"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/controller"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/database"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/repo"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/server"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/service"
	"log"
	"os"
	"os/signal"
)

func Run() {
	cfg := config.NewConfigFromFile()
	dbConn, err := database.NewDBConn(cfg)
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
	s := server.NewServer(cfg, handler)
	go s.Start()
	log.Println("Http Server is running")
	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, os.Interrupt, os.Kill)
	<-sigTerm
	s.Stop()
	db.Close()
	log.Println("finished")
}
