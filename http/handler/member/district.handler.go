package member

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
	"github.com/gin-gonic/gin"
)

type DistrictHandler struct {
	DistrictService usecaseMember.DistrictService
}

func NewDistrictHandler(districtService usecaseMember.DistrictService) DistrictHandler {
	return DistrictHandler{DistrictService: districtService}
}

func (handler *DistrictHandler) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/member")
	{
		api.GET("/district", handler.GetData)
	}
}

func (handler *DistrictHandler) GetData(c *gin.Context) {
	q := c.Query("q")
	data, err := handler.DistrictService.GetData(q)
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
