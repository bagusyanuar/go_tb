package member

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
	"github.com/gin-gonic/gin"
)

type ProductCourseHandler struct {
	ProductCourseService usecaseMember.ProductCourseService
}

func NewProductCourseHandler(productCourseService usecaseMember.ProductCourseService) ProductCourseHandler {
	return ProductCourseHandler{ProductCourseService: productCourseService}
}

func (handler *ProductCourseHandler) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/member")
	{
		api.GET("/product-course", handler.GetData)
		api.GET("/product-course/check", handler.Check)
	}
}

func (handler *ProductCourseHandler) Check(c *gin.Context) {
	data, err := handler.ProductCourseService.Check()
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
func (handler *ProductCourseHandler) GetData(c *gin.Context) {
	subjectID := c.Query("subject")
	gradeID := c.Query("grade")
	cityID := c.Query("city")
	data, err := handler.ProductCourseService.GetData(subjectID, gradeID, cityID)
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
