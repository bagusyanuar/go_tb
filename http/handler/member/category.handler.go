package member

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/common"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	CategoryService usecaseMember.CategoryService
}

func NewCategoryHandler(categoryService usecaseMember.CategoryService) CategoryHandler {
	return CategoryHandler{CategoryService: categoryService}
}

func (handler *CategoryHandler) RegisterRoute(route *gin.Engine) {
	api := route.Group("/api/member")
	{
		api.GET("/category", handler.GetData)
	}
}

func (handler *CategoryHandler) GetData(c *gin.Context) {
	q := c.Query("q")
	data, err := handler.CategoryService.GetData(q)
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