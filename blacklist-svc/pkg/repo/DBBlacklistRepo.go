package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/models"
)

type DBBlacklistRepo struct {
	db *sqlx.DB
}

func NewDBBlacklistRepo(dbConn *sqlx.DB) *DBBlacklistRepo {
	return &DBBlacklistRepo{db: dbConn}
}

func (D DBBlacklistRepo) Create(data models.BlacklistData) error {
	//TODO implement me
	panic("implement me")
}

func (D DBBlacklistRepo) Delete(s string) error {
	//TODO implement me
	panic("implement me")
}

func (D DBBlacklistRepo) GetByID(s string) ([]models.BlacklistData, error) {
	//TODO implement me
	panic("implement me")
}

func (D DBBlacklistRepo) GetByUsername(s string) ([]models.BlacklistData, error) {
	//TODO implement me
	panic("implement me")
}
