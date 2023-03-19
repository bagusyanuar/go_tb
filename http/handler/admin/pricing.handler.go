package admin

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/http/middleware"
	"github.com/bagusyanuar/go_tb/http/request"
	usecaseAdmin "github.com/bagusyanuar/go_tb/usecase/admin"
	"github.com/gin-gonic/gin"
)

type PricingHandler struct {
	PricingService usecaseAdmin.PricingService
}

func NewPricingHandler(pricingService usecaseAdmin.PricingService) PricingHandler {
	return PricingHandler{PricingService: pricingService}
}

func (handler *PricingHandler) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/admin")
	{
		api.POST("/pricing", middleware.Auth, handler.Create)
	}
}

func (handler *PricingHandler) Create(c *gin.Context) {
	var request request.CreatePricingRequest
	c.BindJSON(&request)
	data, err := handler.PricingService.Create(request)
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
