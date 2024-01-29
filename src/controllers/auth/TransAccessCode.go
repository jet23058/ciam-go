package controllers_auth

import (
	"ciam_exam/src/facades"
	"ciam_exam/src/models"
	requests_atuh "ciam_exam/src/requests/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AccessCode struct {
	DB *gorm.DB
}

// @Summary Access Code 轉換為 Access Token API
// @Tags Auth
// @version 1.0
// @param {array} body requests_atuh.TransAccessCode false "paramters"
// @produce application/json
// @Success 200 {object} facades.ResponseSuccess "成功回傳值"
// @Router /trans-access-code [post]
func TransAccessCode(p AccessCode) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload requests_atuh.TransAccessCode

		if err := c.ShouldBindJSON(&payload); err != nil {
			facades.ApiOnError(c, &facades.ApiError{
				StatusCode: http.StatusBadRequest,
				ErrorType:  facades.ERROR_SIGN_IN_PAYLOAD_IS_INVALID,
				Error:      err,
			})
			return
		}

		member := models.Members{
			Id: payload.AccessCode,
		}

		if err := p.DB.First(&member, "id = ? AND is_used = 0", payload.AccessCode).Error; err != nil {
			facades.ApiOnError(c, &facades.ApiError{
				StatusCode: http.StatusInternalServerError,
				ErrorType:  facades.ERROR_ACCESS_CODE_INVALID,
				Error:      err,
			})
			return
		}

		if err := p.DB.Model(&member).Where("id = ?", payload.AccessCode).Updates(models.Members{
			IsUsed: true,
		}).Error; err != nil {
			facades.ApiOnError(c, &facades.ApiError{
				StatusCode: http.StatusServiceUnavailable,
				ErrorType:  facades.ERROR_ACCESS_CODE_INVALID,
				Error:      err,
			})
			return
		}

		c.Header("Authorization", "bearer "+member.AccessToken)
		facades.ApiOnSuccess(c, &facades.ApiSuccess{
			StatusCode: http.StatusOK,
			Data:       "suceess",
		})
	}
}
