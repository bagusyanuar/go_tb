package request

//request to create new city
type CreateCityRequest struct {
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
}
