package builder

import (
	adminHandler "github.com/bagusyanuar/go_tb/http/handler/admin"
	memberHandler "github.com/bagusyanuar/go_tb/http/handler/member"
	mentorHandler "github.com/bagusyanuar/go_tb/http/handler/mentor"
	usecaseAdminRepository "github.com/bagusyanuar/go_tb/repository/admin"
	usecaseMemberRepository "github.com/bagusyanuar/go_tb/repository/member"
	usecaseMentorRepository "github.com/bagusyanuar/go_tb/repository/mentor"
	usecaseAdminService "github.com/bagusyanuar/go_tb/service/admin"
	usecaseMemberService "github.com/bagusyanuar/go_tb/service/member"
	usecaseMentorService "github.com/bagusyanuar/go_tb/service/mentor"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Build(route *gin.Engine, db *gorm.DB) {

	//build up repository

	// admin schema
	authAdminRepository := usecaseAdminRepository.NewAuthRepository(db)
	authAdminService := usecaseAdminService.NewAuthService(authAdminRepository)
	authAdminHandler := adminHandler.NewAuthHandler(authAdminService)
	authAdminHandler.RegisterRoute(route)

	categoryAdminRepository := usecaseAdminRepository.NewCategoryRepository(db)
	categoryAdminService := usecaseAdminService.NewCategoryService(categoryAdminRepository)
	categoryAdminHandler := adminHandler.NewCategoryHandler(categoryAdminService)
	categoryAdminHandler.RegisterRoute(route)

	provinceAdminRepository := usecaseAdminRepository.NewProvinceRepository(db)
	provinceAdminService := usecaseAdminService.NewProvinceService(provinceAdminRepository)
	provinceAdminHandler := adminHandler.NewProvinceHandler(provinceAdminService)
	provinceAdminHandler.RegisterRoute(route)

	cityAdminRepository := usecaseAdminRepository.NewCityRepository(db)
	cityAdminService := usecaseAdminService.NewCityService(cityAdminRepository)
	cityAdminHandler := adminHandler.NewCityHandler(cityAdminService)
	cityAdminHandler.RegisterRoute(route)

	districtAdminRepository := usecaseAdminRepository.NewDistrictRepository(db)
	districtAdminService := usecaseAdminService.NewDistrictService(districtAdminRepository)
	districtAdminHandler := adminHandler.NewDistrictHandler(districtAdminService)
	districtAdminHandler.RegisterRoute(route)

	gradeAdminRepository := usecaseAdminRepository.NewGradeRepository(db)
	gradeAdminService := usecaseAdminService.NewGradeService(gradeAdminRepository)
	gradeAdminHandler := adminHandler.NewGradeHandler(gradeAdminService)
	gradeAdminHandler.RegisterRoute(route)

	mentorLevelAdminRepository := usecaseAdminRepository.NewMentorLevelRepository(db)
	mentorLevelAdminService := usecaseAdminService.NewMentorLevelService(mentorLevelAdminRepository)
	mentorLevelHandler := adminHandler.NewMentorLevelHandler(mentorLevelAdminService)
	mentorLevelHandler.RegisterRoute(route)

	subjectAdminRepository := usecaseAdminRepository.NewSubjectRepository(db)
	subjectAdminService := usecaseAdminService.NewSubjectService(subjectAdminRepository)
	subjectAdminHandler := adminHandler.NewSubjectHandler(subjectAdminService)
	subjectAdminHandler.RegisterRoute(route)

	pricingAdminRepository := usecaseAdminRepository.NewPricingRepository(db)
	pricingAdminSevice := usecaseAdminService.NewPricingservice(pricingAdminRepository)
	pricingAdminHandler := adminHandler.NewPricingHandler(pricingAdminSevice)
	pricingAdminHandler.RegisterRoute(route)

	// mentor schema
	mentorLevelMentorRepository := usecaseMentorRepository.NewMentorLevelRepository(db)

	gradeMentorRepository := usecaseMentorRepository.NewGradeRepository(db)

	authMentorRepository := usecaseMentorRepository.NewAuthRepository(db)
	authMentorService := usecaseMentorService.NewAuthService(authMentorRepository, mentorLevelMentorRepository)
	authMentorHandler := mentorHandler.NewAuthHandler(authMentorService)
	authMentorHandler.RegisterRoute(route)

	profileMentorRepository := usecaseMentorRepository.NewProfileRepository(db)
	profileMentorService := usecaseMentorService.NewProfileService(profileMentorRepository)
	proifleMentorHandler := mentorHandler.NewProfileHandler(profileMentorService)
	proifleMentorHandler.RegisterRoute(route)

	subjectMentorRepository := usecaseMentorRepository.NewSubjectRepository(db)
	subjectMentorService := usecaseMentorService.NewSubjectService(subjectMentorRepository)
	subjectMentorHandler := mentorHandler.NewSubjectHandler(subjectMentorService)
	subjectMentorHandler.RegisterRoute(route)

	productCourseMentorRepository := usecaseMentorRepository.NewProductCourseRepository(db)
	productCourseMentorService := usecaseMentorService.NewProductCourseService(productCourseMentorRepository, profileMentorRepository, subjectMentorRepository, gradeMentorRepository)
	productCourseMentorHandler := mentorHandler.NewProductCourseHandler(productCourseMentorService)
	productCourseMentorHandler.RegisterRoute(route)

	areaMentorRepository := usecaseMentorRepository.NewAreaRepository(db)
	areaMentorService := usecaseMentorService.NewAreaService(areaMentorRepository)
	areaMentorHandler := mentorHandler.NewAreaHandler(areaMentorService)
	areaMentorHandler.RegisterRoute(route)

	//member schema
	authMemberRepository := usecaseMemberRepository.NewAuthRepository(db)
	authMemberService := usecaseMemberService.NewAuthService(authMemberRepository)
	authMemberHandler := memberHandler.NewAuthHandler(authMemberService)
	authMemberHandler.RegisterRoute(route)

	categoryMemberRepository := usecaseMemberRepository.NewCategoryRepository(db)
	categoryMemberService := usecaseMemberService.NewCategoryService(categoryMemberRepository)
	categoryMemberHandler := memberHandler.NewCategoryHandler(categoryMemberService)
	categoryMemberHandler.RegisterRoute(route)

	cityMemberRepository := usecaseMemberRepository.NewCityRepository(db)
	cityMemberService := usecaseMemberService.NewCityService(cityMemberRepository)
	cityMemberHandler := memberHandler.NewCityHandler(cityMemberService)
	cityMemberHandler.RegisterRoute(route)

	districtMemberRepository := usecaseMemberRepository.NewDistrictRepository(db)
	districtMemberService := usecaseMemberService.NewDistrictService(districtMemberRepository)
	districtMemberHandler := memberHandler.NewDistrictHandler(districtMemberService)
	districtMemberHandler.RegisterRoute(route)

	subjectMemberRepository := usecaseMemberRepository.NewSubjectRepository(db)
	subjectMemberService := usecaseMemberService.NewSubjectService(subjectMemberRepository)
	subjectMemberHandler := memberHandler.NewSubjectHandler(subjectMemberService)
	subjectMemberHandler.RegisterRoute(route)

	productCourseMemberRepository := usecaseMemberRepository.NewProductCourseRepository(db)
	productCourseMemberService := usecaseMemberService.NewProductCourseService(productCourseMemberRepository)
	productCourseMemberHandler := memberHandler.NewProductCourseHandler(productCourseMemberService)
	productCourseMemberHandler.RegisterRoute(route)

	courseRequestMemberRepository := usecaseMemberRepository.NewCourseRequestRepository(db)
	courseRequestMemberService := usecaseMemberService.NewCourseRequestService(courseRequestMemberRepository)
	courseRequestMemberHandler := memberHandler.NewCourseRequestHandler(courseRequestMemberService)
	courseRequestMemberHandler.RegisterRoute(route)

}
