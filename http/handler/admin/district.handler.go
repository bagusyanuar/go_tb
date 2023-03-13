package admin

import (
	"errors"
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DistrictHandler struct {
	DistrictAdminService usecase.DistrictAdminService
}

func NewDistrictHandler(districtAdminService usecase.DistrictAdminService) DistrictHandler {
	return DistrictHandler{DistrictAdminService: districtAdminService}
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
	data, err := handler.DistrictAdminService.FindAll(param)
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
	data, err := handler.DistrictAdminService.FindByID(id)
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
	var request domain.CreateDistrictRequest
	c.BindJSON(&request)
	_, err := handler.DistrictAdminService.Create(request)
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

func (handler *DistrictHandler) Patch(c *gin.Context) {
	id := c.Param("id")
	var request domain.CreateDistrictRequest
	c.BindJSON(&request)
	_, err := handler.DistrictAdminService.Patch(id, request)
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
	err := handler.DistrictAdminService.Delete(id)
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