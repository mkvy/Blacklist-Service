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

func (b *BlacklistSvcImpl) Delete(id string) error {
	log.Println("BlacklistSvcImpl: deleting data with id: ", id)
	err := b.repo.Delete(id)
	if err != nil {
		log.Println("BlacklistSvcImpl: error after repo call, ", err)
		return err
	}
	return nil
}

func (b *BlacklistSvcImpl) GetByPhoneNumber(dto dto.GetByPhoneDto) ([]models.BlacklistData, error) {
	log.Println("BlacklistSvcImpl: GetByPhoneNumber data with: ", dto.PhoneNumber)
	data, err := b.repo.GetByPhoneNumber(dto.PhoneNumber)
	if err != nil {
		log.Println("BlacklistSvcImpl: error after repo call, ", err)
		return nil, err
	}
	return data, nil
}

func (b *BlacklistSvcImpl) GetByUsername(dto dto.GetByUsernameDto) ([]models.BlacklistData, error) {
	log.Println("BlacklistSvcImpl: GetByUsernameDto data with: ", dto.Username)
	data, err := b.repo.GetByUsername(dto.Username)
	if err != nil {
		log.Println("BlacklistSvcImpl: error after repo call, ", err)
		return nil, err
	}
	return data, nil
}
