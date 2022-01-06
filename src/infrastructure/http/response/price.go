package response

import "api-server/src/domain/value"

type price struct {
	Price    uint32 `json:"price"`
	Duration int64  `json:"duration"`
}

func priceFromValue(p value.Price) price {
	return price{
		Price:    p.Price,
		Duration: p.Duration.Milliseconds(),
	}
}
