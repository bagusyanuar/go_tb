package mentor

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/http/middleware"
	mentorUsecase "github.com/bagusyanuar/go_tb/usecase/mentor"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	ProfileService mentorUsecase.ProfileService
}

func NewProfileHandler(profileService mentorUsecase.ProfileService) ProfileHandler {
	return ProfileHandler{ProfileService: profileService}
}

func (handler *ProfileHandler) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/mentor")
	{
		api.GET("/profile", middleware.Auth, handler.GetProfile)
		// api.POST("/sign-up", handler.SignUp)
	}
}

func (handler *ProfileHandler) GetProfile(c *gin.Context) {
	authorizedUser := c.MustGet("user").(jwt.MapClaims)
	authorizedUserID := authorizedUser["unique"].(string)
	user, err := handler.ProfileService.GetProfile(authorizedUserID)
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
		Data:    user,
	})
}
