package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/mkvy/BlacklistTestTask/internal/dto"
	"github.com/mkvy/BlacklistTestTask/internal/models"
	"github.com/mkvy/BlacklistTestTask/internal/repo"
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
	validate := validator.New()
	err := validate.Struct(entity)
	if err != nil {
		log.Println("BlacklistSvcImpl err while validating entity")
		return "", err
	}
	err = b.repo.Create(entity)
	if err != nil {
		log.Println("BlacklistSvcImpl err while create entity")
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

func (b *BlacklistSvcImpl) GetByPhoneNumber(phone string) ([]models.BlacklistData, error) {
	log.Println("BlacklistSvcImpl: GetByPhoneNumber data with: ", phone)
	data, err := b.repo.GetByPhoneNumber(phone)
	if err != nil {
		log.Println("BlacklistSvcImpl: error after repo call, ", err)
		return nil, err
	}
	return data, nil
}

func (b *BlacklistSvcImpl) GetByUsername(username string) ([]models.BlacklistData, error) {
	log.Println("BlacklistSvcImpl: GetByUsernameDto data with: ", username)
	data, err := b.repo.GetByUsername(username)
	if err != nil {
		log.Println("BlacklistSvcImpl: error after repo call, ", err)
		return nil, err
	}
	return data, nil
}
