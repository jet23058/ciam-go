package auth_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	controllers_auth "ciam_exam/src/controllers/auth"

	requests_atuh "ciam_exam/src/requests/auth"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func GetGinContext(res *httptest.ResponseRecorder) *gin.Context {
	c, _ := gin.CreateTestContext(res)
	return c
}
func GetResponseBody(body []byte, data interface{}) {
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
}
func GetRequestBody(payload interface{}) io.ReadCloser {
	jsonbytes, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	return io.NopCloser(bytes.NewBuffer(jsonbytes))
}

func UnmarshalResponseBody(body []byte, data interface{}) {
	if err := json.Unmarshal(body, &data); err != nil {
		panic(err)
	}
}

type SuccessAPIResponse struct {
	Data string `json:"data"`
}

func TestRegisterSuccess(t *testing.T) {
	// // arrange
	// gin.SetMode(gin.TestMode)

	// // config.InitOs()
	// router := gin.Default()

	// request := requests_atuh.UserRegister{
	// 	Email: "test@gmail.com",
	// 	User: requests_atuh.User{
	// 		Username: "test@gmail.com",
	// 		Password: "testPassword",
	// 	},
	// }
	// requestBody, err := json.Marshal(request)
	// assert.NoError(t, err)

	// req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewReader(requestBody))
	// assert.NoError(t, err)
	// req.Header.Set("Content-Type", "application/json")

	// res := httptest.NewRecorder()

	// router.POST("/register", controllers_auth.Register)
	// router.ServeHTTP(res, req)

	// // assert
	// assert.Equal(t, http.StatusOK, res.Code)

	// var resBody SuccessAPIResponse
	// GetResponseBody(res.Body.Bytes(), &resBody)

	// assert.Equal(t, "success", resBody.Data)

	// arrange
	resForRegister := httptest.NewRecorder()
	cForRegister := GetGinContext(resForRegister)

	request := requests_atuh.UserRegister{
		Email: "test@gmail.com",
		User: requests_atuh.User{
			Username: "test@gmail.com",
			Password: "testPassword",
		},
	}

	cForRegister.Request = &http.Request{
		Header: make(http.Header),
		Body:   GetRequestBody(request),
	}

	// action
	controllers_auth.Register(cForRegister)

	var resBody SuccessAPIResponse
	GetResponseBody(resForRegister.Body.Bytes(), &resBody)

	// assert
	assert.Equal(t, http.StatusOK, resForRegister.Code)
}
