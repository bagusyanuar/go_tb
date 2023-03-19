package request

type CreateMyProductCourseRequest struct {
	ID        string `json:"id"`
	SubjectID string `json:"subject_id"`
	GradeID   string `json:"grade_id"`
	Method    []int  `json:"method"`
}
