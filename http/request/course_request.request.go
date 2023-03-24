package request

type CreateCourseRequestRequest struct {
	MentorID   string `json:"mentor_id"`
	SubjectID  string `json:"subject_id"`
	GradeID    string `json:"grade_id"`
	DistrictID string `json:"district_id"`
	Method     uint8  `json:"method"`
	Duration   uint8  `json:"duration"`
	Attendees  uint8  `json:"attendees"`
	Encounter  uint8  `json:"encounter"`
	FirstMeet  string `json:"first_meet"`
	Address    string `json:"address"`
	Note       string `json:"note"`
}
