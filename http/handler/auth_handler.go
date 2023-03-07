package handler

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService usecase.AuthService
}

func NewAuthController(authService usecase.AuthService) AuthController {
	return AuthController{
		AuthService: authService,
	}
}

func (controller *AuthController) Route(route *gin.Engine) {
	route.POST("/api/sign-up", controller.SignUp)
}

func (controller *AuthController) SignUp(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	address := c.PostForm("address")
	var vPassword *string
	if password == "" {
		vPassword = nil
	} else {
		vPassword = &password
	}

	request := model.CreateAuthMemberRequest{
		Email:    email,
		Username: username,
		Password: vPassword,
		Name:     name,
		Phone:    phone,
		Address:  address,
	}

	accessToken, err := controller.AuthService.SignUpMember(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    accessToken,
	})
}
