package database

import (
	"errors"
	"github.com/jmoiron/sqlx"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/config"
	"log"
	"sync"
)

type DBConn struct {
	db *sqlx.DB
}

func NewDBConn(cfg config.Config) (*DBConn, error) {
	once := sync.Once{}
	var err error
	var db *sqlx.DB
	once.Do(func() {
		log.Println("Creating DB connection in Blacklist DB with driver " + cfg.Database.DriverName)
		connStr := "user=" + cfg.Database.Username + " password=" + cfg.Database.Password + " dbname=" + cfg.Database.DBname + " sslmode=disable"
		db, err = sqlx.Open(cfg.Database.DriverName, connStr)
		if err != nil {
			log.Println("Error while connecting to db")
			log.Println(err)
			return
		}
		if db == nil {
			log.Println("Error with database")
			err = errors.New("error with database")
			return
		}
	})
	return &DBConn{db: db}, err
}

func (d *DBConn) GetDB() (*sqlx.DB, error) {
	if d.db == nil {
		return nil, errors.New("no database connection")
	}
	return d.db, nil
}

func (d *DBConn) DBClose() {
	err := d.db.Close()
	if err != nil {
		log.Println("error closing connection with database")
	}
	log.Println("DB connection closed")
	return
}
