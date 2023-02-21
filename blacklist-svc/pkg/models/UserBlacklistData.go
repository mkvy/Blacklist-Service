package models

import (
	"time"
)

type BlacklistData struct {
	Id                string    `json:"id,omitempty" db:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	PhoneNumber       string    `json:"phone_number" db:"phone_number" validate:"required,min=8,max=15" example:"79990004422"`
	Username          string    `json:"user_name" db:"user_name" validate:"required,min=1" example:"Test Testov"`
	BanReason         string    `json:"ban_reason" db:"ban_reason" validate:"required,min=1" example:"Ban Reason"`
	DateBanned        time.Time `json:"date_banned,omitempty" db:"date_banned" validate:"required" db:"date_banned" example:"2023-01-15 13:24:49"`
	UsernameWhoBanned string    `json:"username_who_banned" db:"username_who_banned" validate:"required,min=1" example:"SomeUser"`
}
