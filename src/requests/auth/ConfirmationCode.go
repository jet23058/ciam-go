package requests_atuh

type ConfirmationCode struct {
	Username         string `json:"username" example:"test@gmail.com"`
	ConfirmationCode string `json:"confirmation_code" example:"123456"`
}
