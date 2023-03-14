package app

import (
	"github.com/mkvy/BlacklistTestTask/internal/config"
	"github.com/mkvy/BlacklistTestTask/internal/controller"
	"github.com/mkvy/BlacklistTestTask/internal/handler"
	"github.com/mkvy/BlacklistTestTask/internal/repo"
	"github.com/mkvy/BlacklistTestTask/internal/service"
	"github.com/mkvy/BlacklistTestTask/pkg/database"
	"github.com/mkvy/BlacklistTestTask/pkg/server"
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
	ctrl := controller.NewController(svc)
	handle := handler.NewHandler(ctrl)
	s := server.NewServer(*cfg, handle)
	go s.Start()
	log.Println("http: Server is running")
	sigTerm := make(chan os.Signal, 1)
	signal.Notify(sigTerm, os.Interrupt, os.Kill)
	<-sigTerm
	s.Stop()
	dbConn.DBClose()
	log.Println("application finished")
}
