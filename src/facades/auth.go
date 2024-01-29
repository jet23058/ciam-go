package facades

import (
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type App struct {
	CognitoClient   *cognito.CognitoIdentityProvider
	UserPoolID      string
	AppClientID     string
	AppClientSecret string
	Token           string
}

type User struct {
	// Username is the username decided by the user
	// at signup time. This field is not required but it could
	// be useful to have
	Username string `json:"username" validate:"required"`

	// Password is the password decided by the user
	// at signup time. This field is required and no signup
	// can work without this.
	// To create a secure password, contraints on this field are
	// it must contain an uppercase and lowercase letter,
	// a special symbol and a number.
	Password string `json:"password" validate:"required"`
}

type UserRegister struct {
	Email string `json:"email" validate:"required"`
	User  User   `json:"user" validate:"required"`
}

type Response struct {
	Error error `json:"error"`
}
