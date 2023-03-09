package handler

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/gin-gonic/gin"
)

type DistrictController struct {
	DistrictService usecase.DistrictService
}

func NewDistrictController(districtService usecase.DistrictService) DistrictController {
	return DistrictController{DistrictService: districtService}
}

//routing
func (controller *DistrictController) Route(route *gin.Engine) {
	route.GET("/api/district", controller.Index)
	route.POST("/api/district", controller.Index)
}

func (controller *DistrictController) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		controller.create(c)
		return
	}
	controller.fetch(c)
}

func (controller *DistrictController) fetch(c *gin.Context) {
	param := c.Query("q")
	data, err := controller.DistrictService.Fetch(param)
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

func (controller *DistrictController) create(c *gin.Context) {
	name := c.PostForm("name")
	cityID := c.PostForm("city_id")

	request := model.CreateDistrictRequest{
		Name:   name,
		CityID: cityID,
	}

	grade, err := controller.DistrictService.Create(request)
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
