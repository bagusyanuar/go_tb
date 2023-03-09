package handler

import (
	"net/http"

	"github.com/bagusyanuar/go_tb/model"
	"github.com/bagusyanuar/go_tb/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MemberController struct {
	MemberService usecase.MemberService
}

func NewMemberController(memberService usecase.MemberService) MemberController {
	return MemberController{
		MemberService: memberService,
	}
}

//routing
func (controller *MemberController) Route(route *gin.Engine) {
	route.GET("/api/members", controller.Index)
	route.POST("/api/members", controller.Index)
}

func (controller *MemberController) Index(c *gin.Context) {
	if c.Request.Method == "POST" {
		controller.create(c)
		return
	}
	controller.findAll(c)
}

func (controller *MemberController) findAll(c *gin.Context) {
	data, err := controller.MemberService.Fetch()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    data,
	})
}

func (controller *MemberController) create(c *gin.Context) {
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	address := c.PostForm("address")

	userUUID, _ := uuid.Parse("19b9ad6c-6cac-460a-82a8-e255f15dce9f")
	request := model.CreateMemberRequest{
		UserID:  userUUID,
		Name:    name,
		Phone:   phone,
		Address: address,
	}
	member, err := controller.MemberService.Create(request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    member,
	})
}
