package models

type RegisterRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Email       string `json:"email" validate:"required,email"`
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required,min=6"`
	Name        string `json:"name" validate:"required"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required"`
}
