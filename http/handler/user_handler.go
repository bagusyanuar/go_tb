package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService usecase.UserService
}

func NewUserController(userService usecase.UserService) UserController {
	return UserController{
		UserService: userService,
	}
}

func (controller *UserController) Route(route *gin.Engine) {
	route.GET("/api/users", controller.Index)
	route.POST("/api/users", controller.Index)
}

func (controller *UserController) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		controller.create(c)
		return
	}
	controller.findAll(c)
}

func (controller *UserController) findAll(c *gin.Context) {
	data, err := controller.UserService.Fetch()
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
		"data":    data,
	})
}

func (controller *UserController) create(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")

	roles, _ := json.Marshal([]string{"member"})

	var vPassword *string
	if password == "" {
		vPassword = nil
	} else {
		vPassword = &password
	}
	request := model.CreateUserRequest{
		Username: username,
		Email:    email,
		Password: vPassword,
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
