package mentor

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/http/middleware"
	usecaseMentor "github.com/bagusyanuar/go_tb/usecase/mentor"
	"github.com/gin-gonic/gin"
)

type SubjectHandler struct {
	SubjectService usecaseMentor.SubjectService
}

func NewSubjectHandler(subjectService usecaseMentor.SubjectService) SubjectHandler {
	return SubjectHandler{SubjectService: subjectService}
}

func (handler *SubjectHandler) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/mentor")
	{
		api.GET("/subject", middleware.Auth, handler.GetData)
	}
}

func (handler *SubjectHandler) GetData(c *gin.Context) {
	param := c.Query("q")
	data, err := handler.SubjectService.GetData(param)
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
