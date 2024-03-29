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

type ProvinceHandler struct {
	ProvinceService usecaseAdmin.ProvinceService
}

func NewProvinceHandler(provinceService usecaseAdmin.ProvinceService) ProvinceHandler {
	return ProvinceHandler{ProvinceService: provinceService}
}

func (handler *ProvinceHandler) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/admin")
	{
		api.GET("/province", handler.Find)
		api.POST("/province", handler.Create)
		api.GET("/province/:id", handler.FindByID)
		api.PATCH("/province/:id/patch", handler.Patch)
		api.DELETE("/province/:id/delete", handler.Delete)
	}
}

func (handler *ProvinceHandler) Find(c *gin.Context) {
	param := c.Query("q")
	data, err := handler.ProvinceService.FindAll(param)
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

func (handler *ProvinceHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	data, err := handler.ProvinceService.FindByID(id)
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

func (handler *ProvinceHandler) Create(c *gin.Context) {
	var request request.CreateProvinceRequest
	c.BindJSON(&request)
	_, err := handler.ProvinceService.Create(request)
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

func (handler *ProvinceHandler) Patch(c *gin.Context) {
	id := c.Param("id")
	var request request.CreateProvinceRequest
	c.BindJSON(&request)
	_, err := handler.ProvinceService.Patch(id, request)
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

func (handler *ProvinceHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := handler.ProvinceService.Delete(id)
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
