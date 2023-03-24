package member

import (
	"fmt"
	"time"

	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
	usecaseMember "github.com/bagusyanuar/go_tb/usecase/member"
	"github.com/google/uuid"
)

type implementCourseRequestService struct {
	CourseRequestRepository usecaseMember.CourseRequestRepository
}

// Send implements member.CourseRequestService
func (service *implementCourseRequestService) Send(authorizedID string, request request.CreateCourseRequestRequest) (data *domain.CourseRequest, err error) {
	studentID, err := uuid.Parse(authorizedID)
	if err != nil {
		return nil, err
	}

	mentorID, err := uuid.Parse(request.MentorID)
	if err != nil {
		return nil, err
	}

	subjectID, err := uuid.Parse(request.SubjectID)
	if err != nil {
		return nil, err
	}

	gradeID, err := uuid.Parse(request.GradeID)
	if err != nil {
		return nil, err
	}

	districtID := new(uuid.UUID)
	referencePrefix := "TP02"
	if request.Method == 1 {
		tmpDistrictID, err := uuid.Parse(request.DistrictID)
		if err != nil {
			return nil, err
		}
		districtID = &tmpDistrictID
		referencePrefix = "TP01"
	}

	now := time.Now()
	referenceSuffix := now.Format("20060102150405")
	referenceID := fmt.Sprintf("%s/%s", referencePrefix, referenceSuffix)

	firstMeet, err := time.Parse("2006-01-02 15:04:05", request.FirstMeet)
	if err != nil {
		return nil, err
	}

	entity := domain.CourseRequest{
		StudentID:   studentID,
		MentorID:    mentorID,
		SubjectID:   subjectID,
		GradeID:     gradeID,
		DistrictID:  districtID,
		ReferenceID: referenceID,
		Method:      request.Method,
		Duration:    request.Duration,
		Attendees:   request.Attendees,
		Encounter:   request.Encounter,
		FirstMeet:   firstMeet,
		Address:     request.Address,
		Note:        request.Note,
	}
	return service.CourseRequestRepository.Send(entity)
}

func NewCourseRequestService(courseRequestRepository usecaseMember.CourseRequestRepository) usecaseMember.CourseRequestService {
	return &implementCourseRequestService{
		CourseRequestRepository: courseRequestRepository,
	}
}
