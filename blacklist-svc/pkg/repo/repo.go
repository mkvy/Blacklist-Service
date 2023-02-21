package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/models"
)

type BlacklistRepository interface {
	Create(data models.BlacklistData) error
	Delete(string) error
	GetByPhoneNumber(string) ([]models.BlacklistData, error)
	GetByUsername(string) ([]models.BlacklistData, error)
}

type Repository struct {
	BlacklistRepository
}

func NewRepository(conn *sqlx.DB) *Repository {
	return &Repository{NewDBBlacklistRepo(conn)}
}
