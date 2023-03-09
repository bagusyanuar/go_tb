package handler

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	CategoryService usecase.CategoryService
}

func NewCategoryController(categoryService usecase.CategoryService) CategoryController {
	return CategoryController{CategoryService: categoryService}
}

//routing
func (controller *CategoryController) Route(route *gin.Engine) {
	route.GET("/api/categories", controller.Index)
	route.POST("/api/categories", controller.Index)
}

func (controller *CategoryController) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		controller.create(c)
		return
	}
	controller.fetch(c)
}

func (controller *CategoryController) fetch(c *gin.Context) {
	param := c.Query("q")
	data, err := controller.CategoryService.Fetch(param)
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

func (controller *CategoryController) create(c *gin.Context) {
	name := c.PostForm("name")

	request := model.CreateCategoryRequest{
		Name: name,
	}

	category, err := controller.CategoryService.Create(request)
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
