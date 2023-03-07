package builder

import (
	"github.com/bagusyanuar/go_tb/http/handler"
	"github.com/bagusyanuar/go_tb/repository"
	"github.com/bagusyanuar/go_tb/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Build(route *gin.Engine, db *gorm.DB) {

	//build up repository
	authRepository := repository.NewAuthRepository(db)
	userRepository := repository.NewUserRepository(db)
	MemberRepository := repository.NewMemberRepository(db)

	//build up service
	authService := service.NewAuthService(authRepository)
	userService := service.NewUserService(userRepository)
	memberService := service.NewMemberService(MemberRepository)

	//build up http handler
	authController := handler.NewAuthController(authService)
	userController := handler.NewUserController(userService)
	memberController := handler.NewMemberController(memberService)

	//setup route
	authController.Route(route)
	userController.Route(route)
	memberController.Route(route)
}
