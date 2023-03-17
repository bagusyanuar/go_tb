package admin

import (
	"errors"
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/http/request"
	usecaseAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DistrictHandler struct {
	DistrictService usecaseAdmin.DistrictService
}

func NewDistrictHandler(districtService usecaseAdmin.DistrictService) DistrictHandler {
	return DistrictHandler{DistrictService: districtService}
}

func (handler *DistrictHandler) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/admin")
	{
		api.GET("/district", handler.Find)
		api.POST("/district", handler.Create)
		api.GET("/district/:id", handler.FindByID)
		api.PATCH("/district/:id/patch", handler.Patch)
		api.DELETE("/district/:id/delete", handler.Delete)
	}
}

func (handler *DistrictHandler) Find(c *gin.Context) {
	param := c.Query("q")
	data, err := handler.DistrictService.FindAll(param)
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

func (handler *DistrictHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	data, err := handler.DistrictService.FindByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, common.APIResponse{
				Code:    http.StatusNotFound,
				Message: err.Error(),
				Data:    nil,
			})
			return
		}
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

func (handler *DistrictHandler) Create(c *gin.Context) {
	var request request.CreateDistrictRequest
	c.BindJSON(&request)
	data, err := handler.DistrictService.Create(request)
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

func (handler *DistrictHandler) Patch(c *gin.Context) {
	id := c.Param("id")
	var request request.CreateDistrictRequest
	c.BindJSON(&request)
	_, err := handler.DistrictService.Patch(id, request)
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
		Data:    nil,
	})
}

func (handler *DistrictHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := handler.DistrictService.Delete(id)
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
		Data:    nil,
	})
}
