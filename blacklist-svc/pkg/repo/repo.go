package repo

import "github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/models"

type BlacklistRepository interface {
	Create(data models.BlacklistData) error
	Delete(string) error
	GetByPhoneNumber(string) ([]models.BlacklistData, error)
	GetByUsername(string) ([]models.BlacklistData, error)
}
