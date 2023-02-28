package service

import (
	"github.com/mkvy/BlacklistTestTask/pkg/dto"
	"github.com/mkvy/BlacklistTestTask/pkg/models"
)

type UserBlackListSvc interface {
	Add(data dto.BlacklistRequestDto) (string, error)
	Delete(string) error
	GetByPhoneNumber(string) ([]models.BlacklistData, error)
	GetByUsername(string) ([]models.BlacklistData, error)
}
