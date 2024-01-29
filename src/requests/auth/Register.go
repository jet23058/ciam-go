package requests_atuh

type UserRegister struct {
	Email string `json:"email" validate:"required" db:"email" example:"test@gmail.com"`
	User  User   `json:"user" validate:"required"`
}
