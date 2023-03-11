package handler

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	"github.com/bagusyanuar/go_tb/http/middleware"
	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ProductCourseController struct {
	ProductCourseService usecase.ProductCourseService
}

func NewProductCourseController(productCourseService usecase.ProductCourseService) ProductCourseController {
	return ProductCourseController{ProductCourseService: productCourseService}
}

//routing
func (controller *ProductCourseController) Route(route *gin.Engine) {
	route.GET("/api/product-course", middleware.Auth, controller.Index)
	route.POST("/api/product-course", middleware.Auth, controller.Index)
}

func (controller *ProductCourseController) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		controller.create(c)
		return
	}
	controller.fetch(c)
}

func (controller *ProductCourseController) fetch(c *gin.Context) {
	// user := c.MustGet("user").(jwt.MapClaims)
	param := c.Query("q")
	data, err := controller.ProductCourseService.Fetch(param)
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

func (controller *ProductCourseController) create(c *gin.Context) {
	user := c.MustGet("user").(jwt.MapClaims)
	subjectID := c.PostForm("subject_id")
	gradeID := c.PostForm("grade_id")
	methods := []int{1, 2}

	userID := user["unique"].(string)

	request := model.CreateProductCourseRequest{
		UserID:    userID,
		SubjectID: subjectID,
		GradeID:   gradeID,
		Method:    methods,
	}

	_, err := controller.ProductCourseService.Create(request)
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
