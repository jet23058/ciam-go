package responses

type AuthenticationResult struct {
	AccessToken       string `example:"eyJxxxxxxx"`
	ExpiresIn         int    `example:"3600"`
	IdToken           string `example:"eyJxxxxxxx"`
	NewDeviceMetadata string
	RefreshToken      string `example:"eyJxxxxxxx"`
	TokenType         string `example:"bearer"`
}

type ResponseLogin struct {
	AuthenticationResult AuthenticationResult
	ChallengeName        string
	Session              string
}
