package requests_atuh

type ForgotPassword struct {
	Username string `json:"username" validate:"required" example:"test@gmail.com"`
}
