package member

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
	"github.com/gin-gonic/gin"
)

type SubjectHandler struct {
	SubjectService usecaseMember.SubjectService
}

func NewSubjectHandler(subjectService usecaseMember.SubjectService) SubjectHandler {
	return SubjectHandler{SubjectService: subjectService}
}

func (handler *SubjectHandler) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/member")
	{
		api.GET("/subject", handler.GetData)
		api.GET("/subject/:id", handler.GetDataByID)
	}
}

func (handler *SubjectHandler) GetData(c *gin.Context) {
	q := c.Query("q")
	data, err := handler.SubjectService.GetData(q)
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

func (handler *SubjectHandler) GetDataByID(c *gin.Context) {
	id := c.Param("id")
	data, err := handler.SubjectService.GetDataByID(id)
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
