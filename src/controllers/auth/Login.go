package controllers_auth

import (
	"ciam_exam/src/facades"
	"ciam_exam/src/helpers"
	"ciam_exam/src/models"
	requests_atuh "ciam_exam/src/requests/auth"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SaveTokens struct {
	DB *gorm.DB
}

// @Summary 登入 API
// @Tags Auth
// @version 1.0
// @param {array} body requests_atuh.AuthLogin false "paramters"
// @produce application/json
// @Success 200 {object} responses.ResponseLogin "成功回傳值"
// @Router /login [post]
func Login(p SaveTokens) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload requests_atuh.AuthLogin

		if err := c.ShouldBindJSON(&payload); err != nil {
			facades.ApiOnError(c, &facades.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  facades.ERROR_SIGN_IN_PAYLOAD_IS_INVALID,
				Error:      err,
			})
		}

		conf := &aws.Config{Region: aws.String(os.Getenv("COGNITO_REGION"))}
		mySession := session.Must(session.NewSession(conf))

		a := facades.App{
			CognitoClient:   cognito.New(mySession),
			UserPoolID:      os.Getenv("COGNITO_USER_POOL_ID"),
			AppClientID:     os.Getenv("COGNITO_APP_CLIENT_ID"),
			AppClientSecret: os.Getenv("COGNITO_APP_CLIENT_SECRET"),
		}

		u := models.User{
			Account:  payload.Account,
			Password: payload.Password,
		}

		params := map[string]*string{
			"USERNAME": aws.String(u.Account),
			"PASSWORD": aws.String(u.Password),
		}

		secretHash := helpers.ComputeSecretHash(a.AppClientSecret, u.Account, a.AppClientID)
		params["SECRET_HASH"] = aws.String(secretHash)

		authTry := &cognito.InitiateAuthInput{
			AuthFlow: aws.String("USER_PASSWORD_AUTH"),
			AuthParameters: map[string]*string{
				"USERNAME":    aws.String(*params["USERNAME"]),
				"PASSWORD":    aws.String(*params["PASSWORD"]),
				"SECRET_HASH": aws.String(*params["SECRET_HASH"]),
			},
			ClientId: aws.String(a.AppClientID),
		}
		authResp, err := a.CognitoClient.InitiateAuth(authTry)

		if err != nil {
			facades.ApiOnError(c, &facades.ApiError{
				StatusCode: http.StatusInternalServerError,
				ErrorType:  facades.ERROR_SIGN_IN_PAYLOAD_IS_INVALID,
				Error:      err,
			})
		}

		tokens := models.Members{
			Id:            uuid.New(),
			AccessToken:   *authResp.AuthenticationResult.AccessToken,
			IdToken:       *authResp.AuthenticationResult.IdToken,
			RefresthToken: *authResp.AuthenticationResult.RefreshToken,
		}

		if err := p.DB.Create(&tokens).Error; err != nil {
			facades.ApiOnError(c, &facades.ApiError{
				StatusCode: http.StatusInternalServerError,
				ErrorType:  facades.ERROR_SIGN_IN_PAYLOAD_IS_INVALID,
				Error:      err,
			})
		}

		facades.ApiOnSuccess(c, &facades.ApiSuccess{
			StatusCode: http.StatusOK,
			Data:       authResp,
		})
	}
}
