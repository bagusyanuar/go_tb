package handler

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/gin-gonic/gin"
)

type GradeController struct {
	GradeService usecase.GradeService
}

func NewGradeController(gradeService usecase.GradeService) GradeController {
	return GradeController{GradeService: gradeService}
}

//routing
func (controller *GradeController) Route(route *gin.Engine) {
	route.GET("/api/grade", controller.Index)
	route.POST("/api/grade", controller.Index)
}

func (controller *GradeController) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		controller.create(c)
		return
	}
	controller.fetch(c)
}

func (controller *GradeController) fetch(c *gin.Context) {
	param := c.Query("q")
	data, err := controller.GradeService.Fetch(param)
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

func (controller *GradeController) create(c *gin.Context) {
	name := c.PostForm("name")

	request := model.CreateGradeRequest{
		Name: name,
	}

	grade, err := controller.GradeService.Create(request)
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
		Data:    grade,
	})
}
