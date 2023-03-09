package handler

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/gin-gonic/gin"
)

type MentorLevelController struct {
	MentorLevelService usecase.MentorLevelService
}

func NewMentorLevelController(mentorLevelService usecase.MentorLevelService) MentorLevelController {
	return MentorLevelController{MentorLevelService: mentorLevelService}
}

//routing
func (controller *MentorLevelController) Route(route *gin.Engine) {
	route.GET("/api/mentor-level", controller.Index)
	route.POST("/api/mentor-level", controller.Index)
}

func (controller *MentorLevelController) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		controller.create(c)
		return
	}
	controller.fetch(c)
}

func (controller *MentorLevelController) fetch(c *gin.Context) {
	param := c.Query("q")
	data, err := controller.MentorLevelService.Fetch(param)
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
		Data:    data,
	})
}

func (controller *MentorLevelController) create(c *gin.Context) {
	name := c.PostForm("name")

	request := model.CreateMentorLevelRequest{
		Name: name,
	}

	mentorLevel, err := controller.MentorLevelService.Create(request)
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
		Data:    mentorLevel,
	})
}
