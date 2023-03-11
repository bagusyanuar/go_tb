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

type CategoryHandler struct {
	CategoryAdminService usecase.CategoryAdminService
}

func NewCategoryHandler(categoryAdminService usecase.CategoryAdminService) CategoryHandler {
	return CategoryHandler{CategoryAdminService: categoryAdminService}
}

func (handler *CategoryHandler) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/admin")
	{
		api.GET("/category", handler.Find)
		api.POST("/category", handler.Create)
		api.GET("/category/:id", handler.FindByID)
		api.PATCH("/category/:id/patch", handler.Patch)
		api.DELETE("/category/:id/delete", handler.Delete)
	}
}
func (handler *CategoryHandler) Find(c *gin.Context) {
	param := c.Query("q")
	data, err := handler.CategoryAdminService.FindAll(param)
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

func (handler *CategoryHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	data, err := handler.CategoryAdminService.FindByID(id)
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

func (handler *CategoryHandler) Create(c *gin.Context) {
	var request domain.CreateCategoryRequest
	c.BindJSON(&request)
	_, err := handler.CategoryAdminService.Create(request)
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

func (handler *CategoryHandler) Patch(c *gin.Context) {
	id := c.Param("id")
	var request domain.CreateCategoryRequest
	c.BindJSON(&request)
	_, err := handler.CategoryAdminService.Patch(id, request)
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

func (handler *CategoryHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := handler.CategoryAdminService.Delete(id)
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
