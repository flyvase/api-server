package response

import "harvest/src/domain/value"

type price struct {
	Price    uint  `json:"price"`
	Duration int64 `json:"duration"`
}

func priceFromValue(p *value.Price) *price {
	return &price{
		Price:    p.Price,
		Duration: p.Duration.Milliseconds(),
	}
}
