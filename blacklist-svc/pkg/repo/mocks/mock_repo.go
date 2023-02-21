package mocks

import (
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/models"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/repo"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/utils"
	"time"
)

type mockRepo struct{}

func NewMockRepository() *repo.Repository {
	return &repo.Repository{BlacklistRepository: &mockRepo{}}
}

func (m *mockRepo) Create(data models.BlacklistData) error {
	uniquePhoneNumber := "1155669977"
	uniqueUsername := "test_unique"
	uniqueUsernameWhoBanned := "admin"
	if data.PhoneNumber == uniquePhoneNumber && data.Username == uniqueUsername && data.UsernameWhoBanned == uniqueUsernameWhoBanned {
		return utils.ErrAlreadyExists
	}
	return nil
}

func (m *mockRepo) Delete(id string) error {
	idDelete := "123e4567-e89b-12d3-a456-426614174000"
	if id != idDelete {
		return utils.ErrNotFound
	}
	return nil
}

func (m *mockRepo) GetByPhoneNumber(phone string) ([]models.BlacklistData, error) {
	testphone := "78785551111"
	if phone != testphone {
		return nil, utils.ErrNotFound
	}
	return []models.BlacklistData{
		{
			Id:                "test",
			PhoneNumber:       testphone,
			Username:          "test",
			BanReason:         "ban reason",
			DateBanned:        time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			UsernameWhoBanned: "test1",
		},
	}, nil
}

func (m *mockRepo) GetByUsername(username string) ([]models.BlacklistData, error) {
	test_username := "user1"
	if username != test_username {
		return nil, utils.ErrNotFound
	}
	return []models.BlacklistData{
		{
			Id:                "test",
			PhoneNumber:       "78884442233",
			Username:          test_username,
			BanReason:         "ban reason",
			DateBanned:        time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			UsernameWhoBanned: "test1",
		},
	}, nil
}
