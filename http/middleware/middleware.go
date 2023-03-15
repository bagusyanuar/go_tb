package middleware

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	authorization := c.Request.Header.Get("Authorization")

	jwtClaim, err := common.ClaimToken(authorization)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, common.APIResponse{
			Code:    http.StatusUnauthorized,
			Data:    nil,
			Message: err.Error(),
		})
		return
	}

	c.Set("user", jwtClaim)
	c.Next()
}

func Mentor(c *gin.Context) {

}
