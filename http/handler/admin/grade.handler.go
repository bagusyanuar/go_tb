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

type GradeHandler struct {
	GradeService usecaseAdmin.GradeService
}

func NewGradeHandler(gradeService usecaseAdmin.GradeService) GradeHandler {
	return GradeHandler{GradeService: gradeService}
}

func (handler *GradeHandler) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/admin")
	{
		api.GET("/grade", handler.Find)
		api.POST("/grade", handler.Create)
		api.GET("/grade/:id", handler.FindByID)
		api.PATCH("/grade/:id/patch", handler.Patch)
		api.DELETE("/grade/:id/delete", handler.Delete)
	}
}

func (handler *GradeHandler) Find(c *gin.Context) {
	param := c.Query("q")
	data, err := handler.GradeService.FindAll(param)
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

func (handler *GradeHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	data, err := handler.GradeService.FindByID(id)
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

func (handler *GradeHandler) Create(c *gin.Context) {
	var request request.CreateGradeRequest
	c.BindJSON(&request)
	data, err := handler.GradeService.Create(request)
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

func (handler *GradeHandler) Patch(c *gin.Context) {
	id := c.Param("id")
	var request request.CreateGradeRequest
	c.BindJSON(&request)
	data, err := handler.GradeService.Patch(id, request)
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

func (handler *GradeHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := handler.GradeService.Delete(id)
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
