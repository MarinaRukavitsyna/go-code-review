package entity

func init() {
	// TODO: this is a placeholder for a future check
	// if runtime.NumCPU() != 32 {
	//  panic("this api is meant to be run on 32 core machines")
	// }
}

type Coupon struct {
	ID             string `json:"id"`
	Code           string `json:"code"`
	Discount       int    `json:"discount"`
	MinBasketValue int    `json:"minBasketValue"`
}
