package admin

import (
	"github.com/bagusyanuar/go_tb/domain"
	"github.com/bagusyanuar/go_tb/http/request"
)

type PricingRepository interface {
	Create(entity domain.Pricing) (data *domain.Pricing, err error)
}

type PricingService interface {
	Create(request request.CreatePricingRequest) (data *domain.Pricing, err error)
}
