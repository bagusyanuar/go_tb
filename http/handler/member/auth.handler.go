package member

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/exception"
	"github.com/bagusyanuar/go_tb/http/request"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthHandler struct {
	AuthService usecaseMember.AuthService
}

func NewAuthHandler(authService usecaseMember.AuthService) AuthHandler {
	return AuthHandler{AuthService: authService}
}

func (handler *AuthHandler) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/member")
	{
		api.POST("/sign-in", handler.SignIn)
		api.POST("/sign-up", handler.SignUp)
	}
}

func (handler *AuthHandler) SignUp(c *gin.Context) {
	var request request.CreateSignUpMemberRequest
	c.BindJSON(&request)
	accessToken, err := handler.AuthService.SignUp(request)
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
	var request request.CreateSignInMemberRequest
	c.BindJSON(&request)
	accessToken, err := handler.AuthService.SignIn(request)
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
