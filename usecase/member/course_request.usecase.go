package member

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
)

type CourseRequestRepository interface {
	Send(entity domain.CourseRequest) (data *domain.CourseRequest, err error)
}

type CourseRequestService interface {
	Send(authorizedID string, request request.CreateCourseRequestRequest) (data *domain.CourseRequest, err error)
}
