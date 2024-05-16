package entity

type Coupon struct {
	ID             string `json:"id"`
	Code           string `json:"code"`
	Discount       int    `json:"discount"`
	MinBasketValue int    `json:"minBasketValue"`
}
