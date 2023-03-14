package service

import (
	"github.com/go-playground/validator/v10"
	"github.com/mkvy/BlacklistTestTask/internal/dto"
	"github.com/mkvy/BlacklistTestTask/internal/models"
	"github.com/mkvy/BlacklistTestTask/internal/repo/mocks"
	"github.com/mkvy/BlacklistTestTask/pkg/utils"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	svc := NewBlacklistSvcImpl(mocks.NewMockRepository())
	data := dto.BlacklistRequestDto{
		PhoneNumber:       "1234567890",
		Username:          "test_user",
		BanReason:         "reason",
		UsernameWhoBanned: "admin",
	}
	id, err := svc.Add(data)
	if err != nil {
		t.Errorf("Add failed with error: %v", err)
	}
	if !assert.IsType(t, "", id) {
		t.Errorf("Add returned an invalid ID type: %v", reflect.TypeOf(id).Kind())
	}
}

func TestAdd_ValidationError(t *testing.T) {
	svc := NewBlacklistSvcImpl(mocks.NewMockRepository())
	data := dto.BlacklistRequestDto{
		PhoneNumber:       "1234567890",
		Username:          "test_user",
		BanReason:         "reason",
		UsernameWhoBanned: "",
	}
	_, err := svc.Add(data)
	if err == nil {
		t.Error("Add should have returned an error for missing UsernameWhoBanned")
	}
	if _, ok := err.(validator.ValidationErrors); !ok {
		t.Errorf("Add should have returned a validation error, but got: %v", err)
	}
}

func TestAdd_RepoUniqueError(t *testing.T) {
	svc := NewBlacklistSvcImpl(mocks.NewMockRepository())
	uniquePhoneNumber := "1155669977"
	uniqueUsername := "test_unique"
	uniqueUsernameWhoBanned := "admin"
	data := dto.BlacklistRequestDto{
		PhoneNumber:       uniquePhoneNumber,
		Username:          uniqueUsername,
		BanReason:         "reason",
		UsernameWhoBanned: uniqueUsernameWhoBanned,
	}
	expectedErr := utils.ErrAlreadyExists
	_, err := svc.Add(data)
	if err != expectedErr {
		t.Errorf("Add should have returned repo error: %v, but got: %v", expectedErr, err)
	}
}

func TestDelete(t *testing.T) {
	svc := NewBlacklistSvcImpl(mocks.NewMockRepository())
	id := "123e4567-e89b-12d3-a456-426614174000"
	err := svc.Delete(id)
	if err != nil {
		t.Errorf("Delete failed with error: %v", err)
	}
}

func TestDelete_NotFound(t *testing.T) {
	svc := NewBlacklistSvcImpl(mocks.NewMockRepository())
	id := "222e4567-e89b-12d3-a456-426614174000"
	err := svc.Delete(id)
	assert.ErrorIs(t, err, utils.ErrNotFound)
}

func TestBlacklistSvcImpl_GetByPhoneNumber(t *testing.T) {
	svc := NewBlacklistSvcImpl(mocks.NewMockRepository())
	testphone := "78785551111"
	data := []models.BlacklistData{
		{
			Id:                "test",
			PhoneNumber:       testphone,
			Username:          "test",
			BanReason:         "ban reason",
			DateBanned:        time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			UsernameWhoBanned: "test1",
		},
	}
	val, err := svc.GetByPhoneNumber(testphone)
	if err != nil {
		t.Errorf("GetByPhoneNumber failed with error: %v", err)
	}
	assert.Equal(t, data, val)
}

func TestBlacklistSvcImpl_GetByPhoneNumber_NotFound(t *testing.T) {
	svc := NewBlacklistSvcImpl(mocks.NewMockRepository())
	testphone := "78785551112"
	_, err := svc.GetByPhoneNumber(testphone)
	assert.Equal(t, utils.ErrNotFound, err)
}

func TestBlacklistSvcImpl_GetByUsername(t *testing.T) {
	svc := NewBlacklistSvcImpl(mocks.NewMockRepository())
	testUsername := "user1"
	data := []models.BlacklistData{
		{
			Id:                "test",
			PhoneNumber:       "78884442233",
			Username:          testUsername,
			BanReason:         "ban reason",
			DateBanned:        time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			UsernameWhoBanned: "test1",
		},
	}
	val, err := svc.GetByUsername(testUsername)
	if err != nil {
		t.Errorf("GetByUsername failed with error: %v", err)
	}
	assert.Equal(t, data, val)
}

func TestBlacklistSvcImpl_GetByUsername_NotFound(t *testing.T) {
	svc := NewBlacklistSvcImpl(mocks.NewMockRepository())
	testUsername := "test3"
	_, err := svc.GetByPhoneNumber(testUsername)
	assert.Equal(t, utils.ErrNotFound, err)
}
