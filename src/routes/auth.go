package routes

import (
	"ciam_exam/src/config"
	controllers_auth "ciam_exam/src/controllers/auth"

	"github.com/gin-gonic/gin"
)

func UseAuth(r *gin.Engine) {

	db := config.GetDB()

	r.GET("/", controllers_auth.Index)
	r.POST("/register", controllers_auth.Register)
	r.POST("/confirmation-code", controllers_auth.ConfirmationCode)
	r.POST("/login", controllers_auth.Login(controllers_auth.SaveTokens{
		DB: db,
	}))
	r.POST("/trans-access-code", controllers_auth.TransAccessCode(controllers_auth.AccessCode{
		DB: db,
	}))
	r.POST("/forget-password", controllers_auth.ForgotPassword)
	r.POST("/confirm-forget-password", controllers_auth.ConfirmForgotPassword)
	// r.POST("/logout", RevokeToken) // GET https://mydomain.auth.us-east-1.amazoncognito.com/logout?client_id=ad398u21ijw3s9w3939&logout_uri=https://myclient/logout
}
