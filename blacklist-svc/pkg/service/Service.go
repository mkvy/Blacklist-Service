package service

import (
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/dto"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/models"
)

type UserBlackListSvc interface {
	Add(data dto.BlacklistRequestDto) (string, error)
	Delete(string) error
	GetByID(string) ([]models.BlacklistData, error)
	GetByUsername(string) ([]models.BlacklistData, error)
}
