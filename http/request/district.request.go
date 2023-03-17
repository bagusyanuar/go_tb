package request

// request to create new district
type CreateDistrictRequest struct {
	CityID string `json:"city_id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
}
