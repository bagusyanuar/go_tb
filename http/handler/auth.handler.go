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
	route.POST("/api/auth/member/sign-up", controller.SignUpMember)
	route.POST("/api/auth/mentor/sign-up", controller.SignUpMentor)
	route.POST("/api/auth/mentor/sign-in", controller.SignInMentor)
}

func (controller *AuthController) SignUpMember(c *gin.Context) {
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

func (controller *AuthController) SignUpMentor(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	address := c.PostForm("address")
	gender := c.PostForm("gender")
	var vPassword *string
	if password == "" {
		vPassword = nil
	} else {
		vPassword = &password
	}

	request := model.CreateAuthMentorRequest{
		Email:         email,
		Username:      username,
		Password:      vPassword,
		Name:          name,
		Phone:         phone,
		Address:       address,
		Gender:        gender,
		MentorLevelID: "1dde9eec-6ed0-4e52-bc81-061ad9762efb",
	}

	accessToken, err := controller.AuthService.SignUpMentor(request)
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

func (controller *AuthController) SignInMentor(c *gin.Context) {
	password := c.PostForm("password")
	email := c.PostForm("email")

	var vPassword *string
	if password == "" {
		vPassword = nil
	} else {
		vPassword = &password
	}

	request := model.CreateMentorSignInRequest{
		Email:    email,
		Password: vPassword,
	}

	accessToken, err := controller.AuthService.SignInMentor(request)
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
