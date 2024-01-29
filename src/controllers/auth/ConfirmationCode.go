package controllers_auth

import (
	"ciam_exam/src/facades"
	"ciam_exam/src/helpers"
	requests_atuh "ciam_exam/src/requests/auth"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/gin-gonic/gin"
)

// @Summary 註冊認證碼 API
// @Tags Auth
// @version 1.0
// @param {array} body requests_atuh.ConfirmationCode false "paramters"
// @produce application/json
// @Success 200 {object} facades.ResponseSuccess "成功回傳值"
// @Router /confirmation-code [post]
func ConfirmationCode(c *gin.Context) {
	var payload requests_atuh.ConfirmationCode

	if err := c.ShouldBindJSON(&payload); err != nil {
		facades.ApiOnError(c, &facades.ApiError{
			StatusCode: http.StatusBadRequest,
			ErrorType:  facades.ERROR_SIGN_IN_PAYLOAD_IS_INVALID,
			Error:      err,
		})
		return
	}

	conf := &aws.Config{Region: aws.String(os.Getenv("COGNITO_REGION"))}
	mySession := session.Must(session.NewSession(conf))
	a := facades.App{
		CognitoClient:   cognito.New(mySession),
		UserPoolID:      os.Getenv("COGNITO_USER_POOL_ID"),
		AppClientID:     os.Getenv("COGNITO_APP_CLIENT_ID"),
		AppClientSecret: os.Getenv("COGNITO_APP_CLIENT_SECRET"),
	}

	user := &cognito.ConfirmSignUpInput{
		ConfirmationCode: aws.String(payload.ConfirmationCode),
		Username:         aws.String(payload.Username),
		ClientId:         aws.String(a.AppClientID),
	}

	secretHash := helpers.ComputeSecretHash(a.AppClientSecret, payload.Username, a.AppClientID)
	user.SecretHash = aws.String(secretHash)

	_, err := a.CognitoClient.ConfirmSignUp(user)

	if err != nil {
		facades.ApiOnError(c, &facades.ApiError{
			StatusCode: http.StatusInternalServerError,
			ErrorType:  facades.ERROR_REGISTER_FAILED,
			Error:      err,
		})
		return
	}

	facades.ApiOnSuccess(c, &facades.ApiSuccess{
		StatusCode: http.StatusOK,
		Data:       "success",
	})
}
