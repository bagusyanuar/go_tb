package admin

import "github.com/bagusyanuar/go_tb/usecase"

type CityHandler struct {
	CityAdminService usecase.CityAdminService
}

func NewCityHandler(cityAdminService usecase.CityAdminService) CityHandler {
	return CityHandler{CityAdminService: cityAdminService}
}
