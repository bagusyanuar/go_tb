package mentor

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/http/middleware"
	"github.com/bagusyanuar/go_tb/http/request"
	usecaseMentor "github.com/bagusyanuar/go_tb/usecase/mentor"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AreaHandler struct {
	AreaService usecaseMentor.AreaService
}

func NewAreaHandler(areaService usecaseMentor.AreaService) AreaHandler {
	return AreaHandler{AreaService: areaService}
}

func (handler *AreaHandler) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/mentor")
	{
		api.POST("/area", middleware.Auth, handler.CreateMyArea)
	}
}

func (handler *AreaHandler) CreateMyArea(c *gin.Context) {
	var request request.CreateMentorAreaRequest
	c.BindJSON(&request)
	authorizedUser := c.MustGet("user").(jwt.MapClaims)
	authorizedUserID := authorizedUser["unique"].(string)
	err := handler.AreaService.AppendAreas(authorizedUserID, request)
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
