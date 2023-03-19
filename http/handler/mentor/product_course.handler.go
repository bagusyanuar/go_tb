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

type ProductCourseHandler struct {
	ProductCourseService usecaseMentor.ProductCourseService
}

func NewProductCourseHandler(productCourseService usecaseMentor.ProductCourseService) ProductCourseHandler {
	return ProductCourseHandler{ProductCourseService: productCourseService}
}

func (handler *ProductCourseHandler) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/mentor")
	{
		api.GET("/product-course", middleware.Auth, handler.GetMyProductCourse)
		api.POST("/product-course", middleware.Auth, handler.CreateMyProductCourse)
	}
}

func (handler *ProductCourseHandler) CreateMyProductCourse(c *gin.Context) {
	var request request.CreateMyProductCourseRequest
	c.BindJSON(&request)
	authorizedUser := c.MustGet("user").(jwt.MapClaims)
	authorizedUserID := authorizedUser["unique"].(string)
	data, err := handler.ProductCourseService.CreateMyProductCourse(authorizedUserID, request)
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

func (handler *ProductCourseHandler) GetMyProductCourse(c *gin.Context) {

	authorizedUser := c.MustGet("user").(jwt.MapClaims)
	authorizedUserID := authorizedUser["unique"].(string)
	data, err := handler.ProductCourseService.GetMyProductCourse(authorizedUserID)
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
