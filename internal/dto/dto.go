package dto

type BlacklistRequestDto struct {
	PhoneNumber       string `json:"phone_number" validate:"required,min=8,max=15" example:"79990004422"`
	Username          string `json:"user_name" validate:"required,min=1" example:"Test Testov"`
	BanReason         string `json:"ban_reason" validate:"required,min=1" example:"Ban Reason"`
	UsernameWhoBanned string `json:"username_who_banned" validate:"required,min=1" example:"SomeUser"`
}

type ResponseId struct {
	Id string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
}

type Token struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6ImZ1Y2NjY2NjY2NjY2NjY2siLCJpYXQiOjE1MTYyMzkwMjJ9.0xdPqR3zyab1VTKuhaZnS_PzPT2Q5no2IasmlVek1rE"`
}
