package facades

import (
	"github.com/gin-gonic/gin"
)

func ApiOnSuccess(c *gin.Context, res *ApiSuccess) {
	c.JSON(res.StatusCode, gin.H{
		"data": res.Data,
	})
}

func ApiOnError(c *gin.Context, res *ApiError) {
	c.JSON(res.StatusCode, res)
}
