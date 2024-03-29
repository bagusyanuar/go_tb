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

type SubjectHandler struct {
	SubjectService usecaseAdmin.SubjectService
}

func NewSubjectHandler(subjectService usecaseAdmin.SubjectService) SubjectHandler {
	return SubjectHandler{SubjectService: subjectService}
}

func (handler *SubjectHandler) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/admin")
	{
		api.GET("/subject", handler.Find)
		api.POST("/subject", handler.Create)
		api.GET("/subject/:id", handler.FindByID)
		api.PATCH("/subject/:id/patch", handler.Patch)
		api.DELETE("/subject/:id/delete", handler.Delete)
		api.POST("/subject/:id/grades", handler.AppendGrade)
	}
}

func (handler *SubjectHandler) Find(c *gin.Context) {
	param := c.Query("q")
	data, err := handler.SubjectService.FindAll(param)
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

func (handler *SubjectHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	data, err := handler.SubjectService.FindByID(id)
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

func (handler *SubjectHandler) Create(c *gin.Context) {
	var request request.CreateSubjectRequest
	c.BindJSON(&request)
	data, err := handler.SubjectService.Create(request)
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

func (handler *SubjectHandler) Patch(c *gin.Context) {
	id := c.Param("id")
	var request request.CreateSubjectRequest
	c.BindJSON(&request)
	_, err := handler.SubjectService.Patch(id, request)
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

func (handler *SubjectHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := handler.SubjectService.Delete(id)
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

func (handler *SubjectHandler) AppendGrade(c *gin.Context) {
	id := c.Param("id")
	var request request.CreateSubjectAppendGradeRequest
	c.BindJSON(&request)
	err := handler.SubjectService.AppendGrade(id, request)
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
