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

type MentorLevelHandler struct {
	MentorLevelService usecaseAdmin.MentorLevelAdminService
}

func NewMentorLevelHandler(mentorLevelService usecaseAdmin.MentorLevelAdminService) MentorLevelHandler {
	return MentorLevelHandler{MentorLevelService: mentorLevelService}
}

func (handler *MentorLevelHandler) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/admin")
	{
		api.GET("/mentor-level", handler.Find)
		api.POST("/mentor-level", handler.Create)
		api.GET("/mentor-level/:id", handler.FindByID)
		api.PATCH("/mentor-level/:id/patch", handler.Patch)
		api.DELETE("/mentor-level/:id/delete", handler.Delete)
	}
}

func (handler *MentorLevelHandler) Find(c *gin.Context) {
	param := c.Query("q")
	data, err := handler.MentorLevelService.FindAll(param)
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

func (handler *MentorLevelHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	data, err := handler.MentorLevelService.FindByID(id)
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

func (handler *MentorLevelHandler) Create(c *gin.Context) {
	var request request.CreateMentorLevelRequest
	c.BindJSON(&request)
	data, err := handler.MentorLevelService.Create(request)
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

func (handler *MentorLevelHandler) Patch(c *gin.Context) {
	id := c.Param("id")
	var request request.CreateMentorLevelRequest
	c.BindJSON(&request)
	data, err := handler.MentorLevelService.Patch(id, request)
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

func (handler *MentorLevelHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := handler.MentorLevelService.Delete(id)
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
