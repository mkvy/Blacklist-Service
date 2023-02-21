package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/dto"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/models"
	"github.com/mkvy/BlacklistTestTask/blacklist-svc/pkg/service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestAddHandler(t *testing.T) {
	body := &dto.BlacklistRequestDto{
		Username:          "testuser",
		PhoneNumber:       "1234567890",
		BanReason:         "reason",
		UsernameWhoBanned: "user1",
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/api/v1/test/", bytes.NewReader(bodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	ctrl := NewController(mocks.NewMockService())
	router.HandleFunc("/api/v1/test/", ctrl.AddHandler).Methods("POST")
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
	id := "123"
	ans := map[string]string{}
	err = json.NewDecoder(rr.Body).Decode(&ans)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, id, ans["id"])
}

func TestAddHandler_UniqueViolation(t *testing.T) {
	uniquePhoneNumber := "1155669977"
	uniqueUsername := "test_unique"
	uniqueUsernameWhoBanned := "admin"
	body := &dto.BlacklistRequestDto{
		Username:          uniqueUsername,
		PhoneNumber:       uniquePhoneNumber,
		BanReason:         "reason",
		UsernameWhoBanned: uniqueUsernameWhoBanned,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/api/v1/test/", bytes.NewReader(bodyBytes))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	ctrl := NewController(mocks.NewMockService())
	router.HandleFunc("/api/v1/test/", ctrl.AddHandler).Methods("POST")
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusConflict {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusConflict)
	}
}

func TestAddHandler_ValidationError(t *testing.T) {
	body := &dto.BlacklistRequestDto{
		Username:          "1",
		PhoneNumber:       "1",
		BanReason:         "reason",
		UsernameWhoBanned: "1",
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		t.Fatal(err)
	}
	req, err := http.NewRequest("POST", "/api/v1/test/", bytes.NewReader(bodyBytes))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	ctrl := NewController(mocks.NewMockService())
	router.HandleFunc("/api/v1/test/", ctrl.AddHandler).Methods("POST")
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}
}

func TestDeleteHandler(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/api/v1/test/123", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	controller := NewController(mocks.NewMockService())
	router.HandleFunc("/api/v1/test/{id}", controller.DeleteHandler).Methods("DELETE")
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNoContent)
	}
	req, err = http.NewRequest("DELETE", "/api/v1/test/456", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
	req, err = http.NewRequest("DELETE", "/api/v1/test/789", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusInternalServerError)
	}
}

func TestController_GetByPhoneHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/test?phone_number=1234567890", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	controller := NewController(mocks.NewMockService())
	router.HandleFunc("/api/v1/test", controller.GetByPhoneHandler).Queries("phone_number", "{phone_number}").Methods("GET")
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	ans := []models.BlacklistData{}
	err = json.NewDecoder(rr.Body).Decode(&ans)
	expected := []models.BlacklistData{
		{
			Id:                "test456456456",
			PhoneNumber:       "1234567890",
			Username:          "test",
			BanReason:         "ban reason",
			DateBanned:        time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			UsernameWhoBanned: "test1",
		},
	}
	assert.Equal(t, expected, ans)
	req, err = http.NewRequest("GET", "/api/v1/test?phone_number=1232227890", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}

func TestController_GetByUsernameHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/test?user_name=user1", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	controller := NewController(mocks.NewMockService())
	router.HandleFunc("/api/v1/test", controller.GetByUsernameHandler).Queries("user_name", "{user_name}").Methods("GET")
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	ans := []models.BlacklistData{}
	err = json.NewDecoder(rr.Body).Decode(&ans)
	expected := []models.BlacklistData{
		{
			Id:                "test",
			PhoneNumber:       "78884442233",
			Username:          "user1",
			BanReason:         "ban reason",
			DateBanned:        time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			UsernameWhoBanned: "test1",
		},
	}
	assert.Equal(t, expected, ans)
	req, err = http.NewRequest("GET", "/api/v1/test?user_name=aw1ddd", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr = httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
}
