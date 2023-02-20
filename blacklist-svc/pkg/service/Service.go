package service

import (
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/dto"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/models"
)

type UserBlackListSvc interface {
	Add(data dto.BlacklistRequestDto) (string, error)
	Delete(string) error
	GetByPhoneNumber(dto dto.GetByPhoneDto) ([]models.BlacklistData, error)
	GetByUsername(dto dto.GetByUsernameDto) ([]models.BlacklistData, error)
}
