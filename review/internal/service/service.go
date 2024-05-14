package couponService

import (
	"coupon_service/internal/service/entity"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type Repository interface {
	FindByCode(string) (*entity.Coupon, error)
	Save(entity.Coupon) error
}

type CouponService struct {
	repo Repository
}

func New(repo Repository) CouponService {
	return CouponService{
		repo: repo,
	}
}

// Update the given basket with the provided coupon code
func (s CouponService) UpdateBasket(basket entity.Basket, code string) (*entity.Basket, error) {
	b := &basket

	// Retrieve the coupon from the repository
	coupon, err := s.repo.FindByCode(code)
	if err != nil {
		return nil, err
	}

	// Check if the basket value is positive
	if b.Value > 0 {
		// Apply discount and mark the application as successful
		b.AppliedDiscount = coupon.Discount
		b.ApplicationSuccessful = true
	} else if b.Value < 0 {
		// Return an error if trying to apply discount to a negative value
		return nil, fmt.Errorf("tried to apply discount to negative value")
	}

	// Return the modified basket
	return b, nil
}

// Insert a new coupon into the system with the given discount
func (s CouponService) Insert(discount int, code string, minBasketValue int) (string, error) {
	coupon := entity.Coupon{
		Discount:       discount,
		Code:           code,
		MinBasketValue: minBasketValue,
		ID:             uuid.NewString(),
	}

	// Save the coupon to the repository
	// TODO: fix Save and add tests
	if err := s.repo.Save(coupon); err != nil {
		return "", err
	}
	return coupon.ID, nil
}

// Retrieves a list of coupons based on the provided codes
func (s CouponService) GetByCodes(codes []string) ([]entity.Coupon, error) {
	var (
		coupons []entity.Coupon
		errs    []string
	)

	for idx, code := range codes {
		coupon, err := s.repo.FindByCode(code)
		if err != nil {
			errs = append(errs, fmt.Sprintf("code: %s, index: %d", code, idx))
			continue
		}
		coupons = append(coupons, *coupon)
	}

	if len(errs) > 0 {
		return coupons, fmt.Errorf("errors occurred while fetching coupons: %s", strings.Join(errs, "; "))
	}

	return coupons, nil
}
