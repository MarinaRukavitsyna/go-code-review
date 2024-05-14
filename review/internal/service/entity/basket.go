package entity

type Basket struct {
	ID                    string `json:"id"`
	Value                 int    `json:"value"`
	AppliedDiscount       int    `json:"appliedDiscount"`
	ApplicationSuccessful bool   `json:"applicationSuccessful"`
}
