package dto

type CreateUserInput struct {
	Username string `json:"usernamea" biding:"required,min=1"`
	Email    string `json:"email" biding:"required,email"`
	Password string `json:"password" biding:"required,min=8"`
}
