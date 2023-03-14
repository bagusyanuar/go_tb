package builder

import (
	"github.com/bagusyanuar/go_tb/http/handler"
	adminHandler "github.com/bagusyanuar/go_tb/http/handler/admin"
	mentorHandler "github.com/bagusyanuar/go_tb/http/handler/mentor"
	"github.com/bagusyanuar/go_tb/repository"
	"github.com/bagusyanuar/go_tb/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Build(route *gin.Engine, db *gorm.DB) {

	//build up repository

	//admin
	categoryAdminRepository := repository.NewCategoryAdminRepository(db)
	provinceAdminRepository := repository.NewProvinceAdminRepository(db)
	cityAdminRepository := repository.NewCityAdminRepository(db)
	districtAdminRepository := repository.NewDistrictAdminRepository(db)
	gradeAdminRepository := repository.NewGradeAdminRepository(db)
	mentorLevelAdminRepository := repository.NewMentorLevelAdminRepository(db)
	subjectAdminRepository := repository.NewSubjectAdminRepository(db)

	//member
	authRepository := repository.NewAuthRepository(db)
	userRepository := repository.NewUserRepository(db)
	memberRepository := repository.NewMemberRepository(db)
	mentorRepository := repository.NewMentorRepository(db)
	categoryRepository := repository.NewCategoryRepository(db)
	subjectRepository := repository.NewSubjectRepository(db)
	gradeRepository := repository.NewGradeRepository(db)
	mentorLevelRepository := repository.NewMentorLevelRepository(db)
	provinceRepository := repository.NewProvinceRepository(db)
	cityRepository := repository.NewCityRepository(db)
	districtRepository := repository.NewDistrictRepository(db)
	productCourseRepository := repository.NewProductCourseRepository(db)

	//mentor
	authMentorRepository := repository.NewAuthMentorRepository(db)

	//build up service

	//admin
	categoryAdminService := service.NewCategoryAdminService(categoryAdminRepository)
	provinceAdminService := service.NewProvinceAdminService(provinceAdminRepository)
	cityAdminService := service.NewCityAdminService(cityAdminRepository)
	districtAdminService := service.NewDistrictAdminService(districtAdminRepository)
	gradeAdminService := service.NewGradeAdminService(gradeAdminRepository)
	mentorLevelAdminService := service.NewMentorLevelAdminService(mentorLevelAdminRepository)
	subjectAdminService := service.NewSubjectAdminService(subjectAdminRepository)

	//member
	authService := service.NewAuthService(authRepository)
	userService := service.NewUserService(userRepository)
	memberService := service.NewMemberService(memberRepository)
	categoryService := service.NewCategoryService(categoryRepository)
	subjectService := service.NewSubjectService(subjectRepository)
	gradeService := service.NewGradeService(gradeRepository)
	mentorLevelService := service.NewMentorLevelService(mentorLevelRepository)
	provinceService := service.NewProvinceRepository(provinceRepository)
	cityService := service.NewCityService(cityRepository)
	districtService := service.NewDistrictService(districtRepository)
	productCourseService := service.NewProductCourseService(productCourseRepository, mentorRepository, subjectRepository, gradeRepository)

	//mentor
	authMentorService := service.NewAuthMentorService(authMentorRepository, mentorLevelAdminRepository)

	//build up http handler

	//admin
	categoryAdminHandler := adminHandler.NewCategoryHandler(categoryAdminService)
	provinceAdminHandler := adminHandler.NewProvinceHandler(provinceAdminService)
	cityAdminHandler := adminHandler.NewCityHandler(cityAdminService)
	districtAdminHandler := adminHandler.NewDistrictHandler(districtAdminService)
	gradeAdminHandler := adminHandler.NewGradeHandler(gradeAdminService)
	mentorLevelAdminHandler := adminHandler.NewMentorLevelHandler(mentorLevelAdminService)
	subjectAdminHandler := adminHandler.NewSubjectHandler(subjectAdminService)

	//member
	authController := handler.NewAuthController(authService)
	userController := handler.NewUserController(userService)
	memberController := handler.NewMemberController(memberService)
	categoryController := handler.NewCategoryController(categoryService)
	subjectController := handler.NewSubjectController(subjectService)
	gradeController := handler.NewGradeController(gradeService)
	mentorLevelController := handler.NewMentorLevelController(mentorLevelService)
	provinceController := handler.NewProvinceController(provinceService)
	cityController := handler.NewCityController(cityService)
	districtController := handler.NewDistrictController(districtService)
	productCourseController := handler.NewProductCourseController(productCourseService)

	//mentor
	authMentorController := mentorHandler.NewAuthHandler(authMentorService)
	//setup route

	//admin
	categoryAdminHandler.RegisterRoute(route)
	provinceAdminHandler.RegisterRoute(route)
	cityAdminHandler.RegisterRoute(route)
	districtAdminHandler.RegisterRoute(route)
	gradeAdminHandler.RegisterRoute(route)
	mentorLevelAdminHandler.RegisterRoute(route)
	subjectAdminHandler.RegisterRoute(route)

	//member
	authController.Route(route)
	userController.Route(route)
	memberController.Route(route)
	categoryController.Route(route)
	subjectController.Route(route)
	gradeController.Route(route)
	mentorLevelController.Route(route)
	provinceController.Route(route)
	cityController.Route(route)
	districtController.Route(route)
	productCourseController.Route(route)

	//mentor
	authMentorController.RegisterRoute(route)
}
