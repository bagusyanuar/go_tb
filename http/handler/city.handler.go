package handler

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/gin-gonic/gin"
)

type CityController struct {
	CityService usecase.CityService
}

func NewCityController(cityService usecase.CityService) CityController {
	return CityController{CityService: cityService}
}

//routing
func (controller *CityController) Route(route *gin.Engine) {
	route.GET("/api/city", controller.Index)
	route.POST("/api/city", controller.Index)
}

func (controller *CityController) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		controller.create(c)
		return
	}
	controller.fetch(c)
}

func (controller *CityController) fetch(c *gin.Context) {
	param := c.Query("q")
	data, err := controller.CityService.Fetch(param)
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

func (controller *CityController) create(c *gin.Context) {
	name := c.PostForm("name")
	provinceID := c.PostForm("province_id")

	request := model.CreateCityRequest{
		Name:       name,
		ProvinceID: provinceID,
	}

	grade, err := controller.CityService.Create(request)
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
