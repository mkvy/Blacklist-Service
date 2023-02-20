package app

import (
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/config"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/controller"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/database"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/repo"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/service"
	"log"
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
	_ = controller.NewController(svc)
}
