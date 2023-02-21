package service

import (
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/dto"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/models"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/repo"
)

type UserBlackListSvc interface {
	Add(data dto.BlacklistRequestDto) (string, error)
	Delete(string) error
	GetByPhoneNumber(string) ([]models.BlacklistData, error)
	GetByUsername(string) ([]models.BlacklistData, error)
}

type Service struct {
	UserBlackListSvc
}

func NewService(repo *repo.Repository) *Service {
	return &Service{NewBlacklistSvcImpl(*repo)}
}
