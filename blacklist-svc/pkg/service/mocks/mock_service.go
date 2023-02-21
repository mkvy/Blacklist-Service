package mocks

import (
	"errors"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/dto"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/models"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/service"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/utils"
	"time"
)

type mockService struct{}

func NewMockService() *service.Service {
	return &service.Service{UserBlackListSvc: &mockService{}}
}

func (s *mockService) Add(data dto.BlacklistRequestDto) (string, error) {
	uniquePhoneNumber := "1155669977"
	uniqueUsername := "test_unique"
	uniqueUsernameWhoBanned := "admin"
	if data.PhoneNumber == uniquePhoneNumber && data.Username == uniqueUsername && data.UsernameWhoBanned == uniqueUsernameWhoBanned {
		return "", utils.ErrAlreadyExists
	}
	return "123", nil
}

func (s *mockService) Delete(id string) error {
	if id == "123" {
		return nil
	}
	if id == "456" {
		return utils.ErrNotFound
	}
	return errors.New("some error")
}

func (s *mockService) GetByUsername(username string) ([]models.BlacklistData, error) {
	if username == "user1" {
		return []models.BlacklistData{
			{
				Id:                "test",
				PhoneNumber:       "78884442233",
				Username:          username,
				BanReason:         "ban reason",
				DateBanned:        time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
				UsernameWhoBanned: "test1",
			},
		}, nil
	}
	return nil, utils.ErrNotFound
}

func (s *mockService) GetByPhoneNumber(phoneNumber string) ([]models.BlacklistData, error) {
	if phoneNumber == "1234567890" {
		return []models.BlacklistData{
			{
				Id:                "test456456456",
				PhoneNumber:       phoneNumber,
				Username:          "test",
				BanReason:         "ban reason",
				DateBanned:        time.Date(2000, 1, 1, 0, 0, 0, 0, time.FixedZone("test", 0)),
				UsernameWhoBanned: "test1",
			},
		}, nil
	}
	return nil, utils.ErrNotFound
}
