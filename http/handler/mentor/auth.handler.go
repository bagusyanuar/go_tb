package mentor

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthMentorService usecase.AuthMentorService
}

func NewAuthHandler(authMentorService usecase.AuthMentorService) AuthHandler {
	return AuthHandler{AuthMentorService: authMentorService}
}

func (handler *AuthHandler) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/mentor")
	{
		api.POST("/sign-up", handler.SignUp)
	}
}

func (handler *AuthHandler) SignUp(c *gin.Context) {
	var request domain.CreateSignUpMentorRequest
	c.BindJSON(&request)
	accessToken, err := handler.AuthMentorService.SignUp(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, common.APIResponse{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}
	c.JSON(http.StatusOK, common.APIResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    accessToken,
	})
}
