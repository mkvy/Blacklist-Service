package models

import (
	"time"
)

type BlacklistData struct {
	Id                string    `json:"id,omitempty" db:"id"`
	PhoneNumber       string    `json:"phone_number" validate:"required,min=8,max=15" db:"phone_number"`
	Username          string    `json:"user_name" validate:"required,min=1" db:"user_name"`
	BanReason         string    `json:"ban_reason" validate:"required,min=1" db:"ban_reason"`
	DateBanned        time.Time `json:"date_banned,omitempty" validate:"required" db:"date_banned"`
	UsernameWhoBanned string    `json:"username_who_banned" validate:"required,min=1" db:"username_who_banned"`
}
