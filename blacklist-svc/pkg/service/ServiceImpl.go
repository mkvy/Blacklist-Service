package service

import (
	"github.com/google/uuid"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/dto"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/models"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/repo"
	"log"
	"time"
)

type BlacklistSvcImpl struct {
	repo repo.BlacklistRepository
}

func NewBlacklistSvcImpl(repo repo.BlacklistRepository) *BlacklistSvcImpl {
	return &BlacklistSvcImpl{repo}
}

func (b *BlacklistSvcImpl) Add(data dto.BlacklistRequestDto) (string, error) {
	log.Println("BlacklistSvcImpl: adding data")
	id := uuid.New().String()
	entity := models.BlacklistData{
		Id:                id,
		PhoneNumber:       data.PhoneNumber,
		Username:          data.Username,
		BanReason:         data.BanReason,
		DateBanned:        time.Now(),
		UsernameWhoBanned: data.UsernameWhoBanned,
	}
	err := b.repo.Create(entity)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (b BlacklistSvcImpl) Delete(s string) error {
	//TODO implement me
	panic("implement me")
}

func (b BlacklistSvcImpl) GetByID(s string) ([]models.BlacklistData, error) {
	//TODO implement me
	panic("implement me")
}

func (b BlacklistSvcImpl) GetByUsername(s string) ([]models.BlacklistData, error) {
	//TODO implement me
	panic("implement me")
}
