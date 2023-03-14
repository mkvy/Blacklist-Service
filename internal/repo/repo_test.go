package repo

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mkvy/BlacklistTestTask/internal/models"
	"github.com/mkvy/BlacklistTestTask/pkg/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDBBlacklistRepo_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database connection: %s", err)
	}
	defer db.Close()
	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewDBBlacklistRepo(dbx)
	data := models.BlacklistData{
		Id:                uuid.New().String(),
		PhoneNumber:       "555-555-5555",
		Username:          "testuser",
		BanReason:         "test ban reason",
		DateBanned:        time.Now(),
		UsernameWhoBanned: "admin",
	}
	mock.ExpectExec("INSERT INTO blacklist").
		WithArgs(data.Id, data.PhoneNumber, data.Username, data.BanReason, data.DateBanned, data.UsernameWhoBanned).
		WillReturnResult(sqlmock.NewResult(1, 1))

	if err := repo.Create(data); err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %s", err)
	}
}

func TestDBBlacklistRepo_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database connection: %s", err)
	}
	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewDBBlacklistRepo(dbx)

	id := uuid.New().String()

	mock.ExpectExec("delete from blacklist").
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	if err := repo.Delete(id); err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("unfulfilled expectations: %s", err)
	}
}

func TestDBBlacklistRepo_Delete_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database connection: %s", err)
	}
	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewDBBlacklistRepo(dbx)

	id := uuid.New().String()

	mock.ExpectExec("delete from blacklist").
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err = repo.Delete(id)
	if err != nil {
		assert.Equal(t, err, utils.ErrNotFound)
		return
	}
}

func TestDBBlacklistRepo_GetByPhoneNumber(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database connection: %s", err)
	}
	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewDBBlacklistRepo(dbx)

	phone := "79891234545"

	rows := sqlmock.NewRows([]string{"id", "phone_number", "user_name", "ban_reason", "date_banned", "username_who_banned"}).
		AddRow(uuid.New().String(), phone, "testuser", "test ban reason", time.Now(), "admin")

	mock.ExpectQuery("select (.*) from blacklist").
		WithArgs(phone).
		WillReturnRows(rows)

	result, err := repo.GetByPhoneNumber(phone)

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if len(result) != 1 {
		t.Errorf("expected 1 result, got %d", len(result))
	}
}

func TestDBBlacklistRepo_GetByPhoneNumber_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database connection: %s", err)
	}
	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewDBBlacklistRepo(dbx)

	phone := "testuser"

	rows := sqlmock.NewRows([]string{"id", "phone_number", "user_name", "ban_reason", "date_banned", "username_who_banned"})

	mock.ExpectQuery("select (.*) from blacklist").
		WithArgs(phone).WillReturnRows(rows)

	_, err = repo.GetByPhoneNumber(phone)

	if err != nil {
		assert.Equal(t, err, utils.ErrNotFound)
	}
}

func TestDBBlacklistRepo_GetByUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database connection: %s", err)
	}
	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewDBBlacklistRepo(dbx)

	username := "testuser"

	rows := sqlmock.NewRows([]string{"id", "phone_number", "user_name", "ban_reason", "date_banned", "username_who_banned"}).
		AddRow(uuid.New().String(), "79891234545", username, "test ban reason", time.Now(), "admin")

	mock.ExpectQuery("select (.*) from blacklist").
		WithArgs(username).
		WillReturnRows(rows)

	result, err := repo.GetByUsername(username)

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if len(result) != 1 {
		t.Errorf("expected 1 result, got %d", len(result))
	}
}

func TestDBBlacklistRepo_GetByUsername_NotFound(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database connection: %s", err)
	}
	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")
	repo := NewDBBlacklistRepo(dbx)

	username := "testuser"

	rows := sqlmock.NewRows([]string{"id", "phone_number", "user_name", "ban_reason", "date_banned", "username_who_banned"})

	mock.ExpectQuery("select (.*) from blacklist").
		WithArgs(username).WillReturnRows(rows)

	_, err = repo.GetByUsername(username)

	if err != nil {
		assert.Equal(t, err, utils.ErrNotFound)
	}
}
