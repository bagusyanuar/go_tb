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
	memberRepository := repository.NewMemberRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)
	subjectRepository := repository.NewSubjectRepository(db)
	gradeRepository := repository.NewGradeRepository(db)
	mentorLevelRepository := repository.NewMentorLevelRepository(db)

	//build up service
	authService := service.NewAuthService(authRepository)
	userService := service.NewUserService(userRepository)
	memberService := service.NewMemberService(memberRepository)
	categoryService := service.NewCategoryService(categoryRepository)
	subjectService := service.NewSubjectService(subjectRepository)
	gradeService := service.NewGradeService(gradeRepository)
	mentorLevelService := service.NewMentorLevelService(mentorLevelRepository)

	//build up http handler
	authController := handler.NewAuthController(authService)
	userController := handler.NewUserController(userService)
	memberController := handler.NewMemberController(memberService)
	categoryController := handler.NewCategoryController(categoryService)
	subjectController := handler.NewSubjectController(subjectService)
	gradeController := handler.NewGradeController(gradeService)
	mentorLevelController := handler.NewMentorLevelController(mentorLevelService)

	//setup route
	authController.Route(route)
	userController.Route(route)
	memberController.Route(route)
	categoryController.Route(route)
	subjectController.Route(route)
	gradeController.Route(route)
	mentorLevelController.Route(route)
}
