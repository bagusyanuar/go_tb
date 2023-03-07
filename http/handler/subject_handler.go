package handler

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/gin-gonic/gin"
)

type SubjectController struct {
	SubjectService usecase.SubjectService
}

func NewSubjectController(subjectService usecase.SubjectService) SubjectController {
	return SubjectController{SubjectService: subjectService}
}

//routing
func (controller *SubjectController) Route(route *gin.Engine) {
	route.GET("/api/subjects", controller.Index)
	route.POST("/api/subjects", controller.Index)
}

func (controller *SubjectController) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		controller.create(c)
		return
	}
	controller.fetch(c)
}

func (controller *SubjectController) fetch(c *gin.Context) {
	param := c.Query("q")
	data, err := controller.SubjectService.Fetch(param)
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

func (controller *SubjectController) create(c *gin.Context) {
	categoryID := c.PostForm("category_id")
	name := c.PostForm("name")

	request := model.CreateSubjectRequest{
		CategoryID: categoryID,
		Name:       name,
	}

	category, err := controller.SubjectService.Create(request)
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
		Data:    category,
	})
}
