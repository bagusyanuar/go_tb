package admin

import (
	"errors"
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/middleware"
	usecasAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryHandler struct {
	CategoryService usecasAdmin.CategoryService
}

func NewCategoryHandler(categoryService usecasAdmin.CategoryService) CategoryHandler {
	return CategoryHandler{CategoryService: categoryService}
}

func (handler *CategoryHandler) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/admin")
	{
		api.GET("/category", middleware.Auth, handler.Find)
		api.POST("/category", middleware.Auth, handler.Create)
		api.GET("/category/:id", middleware.Auth, handler.FindByID)
		api.PATCH("/category/:id/patch", middleware.Auth, handler.Patch)
		api.DELETE("/category/:id/delete", middleware.Auth, handler.Delete)
	}
}

func (handler *CategoryHandler) Find(c *gin.Context) {
	param := c.Query("q")
	data, err := handler.CategoryService.FindAll(param)
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
	data, err := handler.CategoryService.FindByID(id)
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
	_, err := handler.CategoryService.Create(request)
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
	_, err := handler.CategoryService.Patch(id, request)
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
	err := handler.CategoryService.Delete(id)
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
