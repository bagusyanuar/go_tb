package member

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
	"github.com/gin-gonic/gin"
)

type CityHandler struct {
	CityService usecaseMember.CityService
}

func NewCityHandler(cityService usecaseMember.CityService) CityHandler {
	return CityHandler{CityService: cityService}
}

func (handler *CityHandler) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/member")
	{
		api.GET("/city", handler.GetData)
	}
}

func (handler *CityHandler) GetData(c *gin.Context) {
	q := c.Query("q")
	data, err := handler.CityService.GetData(q)
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
