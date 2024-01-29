package requests_atuh

type ConfirmForgotPassword struct {
	ConfirmationCode string `json:"confirmation_code" validate:"required" example:"123456"`
	User             User   `json:"user" validate:"required"`
}
