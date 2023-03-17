package request

type CreateSubjectRequest struct {
	CategoryID string `json:"category_id"`
	Name       string `json:"name"`
}

type CreateSubjectAppendGradeRequest struct {
	GradeIDS []string `json:"grade_ids"`
}
