package member

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/http/middleware"
	"github.com/bagusyanuar/go_tb/http/request"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type CourseRequestHandler struct {
	CourseRequestService usecaseMember.CourseRequestService
}

func NewCourseRequestHandler(courseRequestService usecaseMember.CourseRequestService) CourseRequestHandler {
	return CourseRequestHandler{CourseRequestService: courseRequestService}
}

func (handler *CourseRequestHandler) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/member")
	{
		api.POST("/course-request", middleware.Auth, handler.Send)
	}
}

func (handler *CourseRequestHandler) Send(c *gin.Context) {
	var request request.CreateCourseRequestRequest
	c.BindJSON(&request)
	authorizedUser := c.MustGet("user").(jwt.MapClaims)
	authorizedUserID := authorizedUser["unique"].(string)
	data, err := handler.CourseRequestService.Send(authorizedUserID, request)
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
