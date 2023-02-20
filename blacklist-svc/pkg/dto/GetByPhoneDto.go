package dto

type GetByPhoneDto struct {
	PhoneNumber string `json:"phone_number" validate:"required,min=8,max=15"`
}
