package requests_atuh

type AuthLogin struct {
	Account  string `binding:"min=3,max=100" json:"account" example:"test@gmail.com"`
	Password string `binding:"min=3,max=100" json:"password" example:"testPassword"`
}
