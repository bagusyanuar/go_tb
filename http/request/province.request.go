package request

//request to create new province
type CreateProvinceRequest struct {
	Name string `json:"name"`
	Code string `json:"code"`
}
