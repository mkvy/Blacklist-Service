package models

import "time"

type BlacklistData struct {
	Id                string    `json:"id,omitempty"`
	PhoneNumber       string    `json:"phone_number" validate:"required,min=8,max=15"`
	Username          string    `json:"user_name" validate:"required,min=1"`
	BanReason         string    `json:"ban_reason" validate:"required,min=1"`
	DateBanned        time.Time `json:"date_banned,omitempty" validate:"required"`
	UsernameWhoBanned string    `json:"username_who_banned" validate:"required,min=1"`
}
