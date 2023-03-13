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

type MentorLevelHandler struct {
	MentorLevelAdminService usecase.MentorLevelAdminService
}

func NewMentorLevelHandler(mentorLevelAdminService usecase.MentorLevelAdminService) MentorLevelHandler {
	return MentorLevelHandler{MentorLevelAdminService: mentorLevelAdminService}
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
	data, err := handler.MentorLevelAdminService.FindAll(param)
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
	data, err := handler.MentorLevelAdminService.FindByID(id)
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
	var request domain.CreateMentorLevelRequest
	c.BindJSON(&request)
	data, err := handler.MentorLevelAdminService.Create(request)
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
	var request domain.CreateMentorLevelRequest
	c.BindJSON(&request)
	data, err := handler.MentorLevelAdminService.Patch(id, request)
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
	err := handler.MentorLevelAdminService.Delete(id)
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
