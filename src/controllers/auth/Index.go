package controllers_auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	// callbackUrl := c.Query("callback_url")

	// if callbackUrl == "" {
	// 	facades.ApiOnError(c, &facades.ApiError{
	// 		StatusCode: http.StatusBadRequest,
	// 		ErrorType:  facades.ERROR_SIGN_IN_PAYLOAD_IS_INVALID,
	// 	})
	// 	return
	// }

	c.HTML(http.StatusOK, "login.go.tmpl", gin.H{})
}
