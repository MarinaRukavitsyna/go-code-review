package memdb

import (
	"coupon_service/internal/service/entity"
	"fmt"
)

type Config struct{}

type repository interface {
	FindByCode(string) (*entity.Coupon, error)
	Insert(entity.Coupon) error
}

// CouponDao is a simple in-memory repository for coupons
// DAO stands for "Data Access Object"
type CouponDao struct {
	entries map[string]entity.Coupon
}

func New() *CouponDao {
	return &CouponDao{
		entries: make(map[string]entity.Coupon),
	}
}

// Retrieve a coupon from the in-memory database based on the provided code
func (dao *CouponDao) FindByCode(code string) (*entity.Coupon, error) {
	coupon, ok := dao.entries[code]
	if !ok {
		return nil, fmt.Errorf("coupon not found")
	}
	return &coupon, nil
}

// Insert a new coupon into the in-memory database
func (dao *CouponDao) Insert(coupon entity.Coupon) error {
	// Check if the entries map is nil
	if dao.entries == nil {
		return fmt.Errorf("entries map is not initialized")
	}

	// Check if the coupon code already exists
	if _, exists := dao.entries[coupon.Code]; exists {
		return fmt.Errorf("coupon with code '%s' already exists", coupon.Code)
	}

	// If no errors, insert the coupon into the entries map
	dao.entries[coupon.Code] = coupon
	return nil
}
