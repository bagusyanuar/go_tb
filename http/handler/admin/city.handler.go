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

type CityHandler struct {
	CityService usecaseAdmin.CityService
}

func NewCityHandler(cityService usecaseAdmin.CityService) CityHandler {
	return CityHandler{CityService: cityService}
}

func (handler *CityHandler) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/admin")
	{
		api.GET("/city", handler.Find)
		api.POST("/city", handler.Create)
		api.GET("/city/:id", handler.FindByID)
		api.PATCH("/city/:id/patch", handler.Patch)
		api.DELETE("/city/:id/delete", handler.Delete)
	}
}

func (handler *CityHandler) Find(c *gin.Context) {
	param := c.Query("q")
	data, err := handler.CityService.FindAll(param)
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

func (handler *CityHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	data, err := handler.CityService.FindByID(id)
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

func (handler *CityHandler) Create(c *gin.Context) {
	var request request.CreateCityRequest
	c.BindJSON(&request)
	_, err := handler.CityService.Create(request)
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
		Data:    request,
	})
}

func (handler *CityHandler) Patch(c *gin.Context) {
	id := c.Param("id")
	var request request.CreateCityRequest
	c.BindJSON(&request)
	_, err := handler.CityService.Patch(id, request)
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

func (handler *CityHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := handler.CityService.Delete(id)
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
