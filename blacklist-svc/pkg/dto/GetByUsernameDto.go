package dto

type GetByUsernameDto struct {
	Username string `json:"user_name" validate:"required,min=1"`
}
