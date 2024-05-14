package memdb

import (
	"coupon_service/internal/service/entity"
	"testing"
)

func TestNew(t *testing.T) {
	dao := New()

	// Check if the entries map is initialized
	if dao.entries == nil {
		t.Errorf("entries map is not initialized")
	}

	// Check if the entries map is empty
	if len(dao.entries) != 0 {
		t.Errorf("entries map is not empty")
	}
}

func TestFindByCode(t *testing.T) {
	dao := New()

	// Create a test coupon
	testCoupon := entity.Coupon{
		Code: "TESTCODE",
	}

	// Insert the test coupon into the in-memory database
	err := dao.Insert(testCoupon)
	if err != nil {
		t.Errorf("failed to insert test coupon: %v", err)
	}

	// Retrieve the coupon by code
	coupon, err := dao.FindByCode(testCoupon.Code)
	if err != nil {
		t.Errorf("failed to retrieve coupon by code: %v", err)
	}

	// Check if the retrieved coupon matches the test coupon
	if coupon == nil || *coupon != testCoupon {
		t.Errorf("retrieved coupon does not match the test coupon")
	}
}

func TestInsert(t *testing.T) {
	dao := New()

	// Create a test coupon
	testCoupon := entity.Coupon{
		Code: "TESTCODE",
	}

	// Insert the test coupon into the in-memory database
	err := dao.Insert(testCoupon)
	if err != nil {
		t.Errorf("failed to insert test coupon: %v", err)
	}

	// Check if the coupon exists in the entries map
	if _, exists := dao.entries[testCoupon.Code]; !exists {
		t.Errorf("coupon with code '%s' was not inserted", testCoupon.Code)
	}
}
