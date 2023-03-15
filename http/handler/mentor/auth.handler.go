package mentor

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/exception"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		api.POST("/sign-in", handler.SignIn)
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

func (handler *AuthHandler) SignIn(c *gin.Context) {
	var request domain.CreateSignInMentorRequest
	c.BindJSON(&request)
	accessToken, err := handler.AuthMentorService.SignIn(request)
	if err != nil {
		switch err {
		case exception.ErrorPasswordNotMatch:
			c.AbortWithStatusJSON(http.StatusUnauthorized, common.APIResponse{
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
				Data:    nil,
			})
			return
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, common.APIResponse{
				Code:    http.StatusNotFound,
				Message: err.Error(),
				Data:    nil,
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, common.APIResponse{
				Code:    http.StatusInternalServerError,
				Message: err.Error(),
				Data:    nil,
			})
			return
		}
	}
	c.JSON(http.StatusOK, common.APIResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    accessToken,
	})
}
