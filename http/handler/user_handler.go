package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService domain.UserService
}

func NewUserController(userService domain.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (controller *UserController) Create(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	roles, _ := json.Marshal([]string{"member"})

	request := model.CreateUserRequest{
		Username: username,
		Email:    email,
		Password: &password,
		Roles:    roles,
	}

	user, err := controller.UserService.Create(request)
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
		"data":    user,
	})
}
