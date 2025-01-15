package model

type SignupRequest struct {
	Name     string `validate:"required,min=5,max=20"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}
