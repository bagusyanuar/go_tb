package request

type CreatePricingRequest struct {
	GradeID       string `json:"grade_id"`
	CityID        string `json:"city_id"`
	MentorLevelID string `json:"mentor_level_id"`
	Method        uint   `json:"method"`
	Price         uint   `json:"price"`
}
