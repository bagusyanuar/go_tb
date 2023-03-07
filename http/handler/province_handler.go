package handler

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/gin-gonic/gin"
)

type ProvinceController struct {
	ProvinceService usecase.ProvinceService
}

func NewProvinceController(provinceService usecase.ProvinceService) ProvinceController {
	return ProvinceController{ProvinceService: provinceService}
}

//routing
func (controller *ProvinceController) Route(route *gin.Engine) {
	route.GET("/api/province", controller.Index)
	route.POST("/api/province", controller.Index)
}

func (controller *ProvinceController) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		controller.create(c)
		return
	}
	controller.fetch(c)
}

func (controller *ProvinceController) fetch(c *gin.Context) {
	param := c.Query("q")
	data, err := controller.ProvinceService.Fetch(param)
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

func (controller *ProvinceController) create(c *gin.Context) {
	name := c.PostForm("name")

	request := model.CreateProvinceRequest{
		Name: name,
	}

	grade, err := controller.ProvinceService.Create(request)
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
