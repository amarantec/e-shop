package ordermodel

import "time"

type OrderDetailsResponse struct {
	OrderDate  time.Time
	TotalPrice float64
	Products   []OrderDetailsProductResponse
}
