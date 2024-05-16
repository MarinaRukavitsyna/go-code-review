package api

import (
	"bytes"
	"coupon_service/internal/service/entity"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type MockCouponService struct{}

func (m MockCouponService) UpdateBasket(basket entity.Basket, code string) (*entity.Basket, error) {
	return &basket, nil
}

func (m MockCouponService) Insert(discount int, code string, minBasketValue int) (string, error) {
	return "mockID", nil
}

func (m MockCouponService) GetByCodes(codes []string) ([]entity.Coupon, error) {
	return []entity.Coupon{{Discount: 10, Code: "SuperDiscount", MinBasketValue: 50}}, nil
}

func TestAPIRouter_ApplyCoupon(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cfg := Config{Port: 8080}
	router := New(cfg, MockCouponService{})

	reqBody, _ := json.Marshal(map[string]any{
		"code": "test_code",
		"basket": map[string]interface{}{
			"id":                    "test_id",
			"value":                 100,
			"appliedDiscount":       0,
			"applicationSuccessful": false,
		},
	})
	req, err := http.NewRequest("POST", "/api/apply", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestAPIRouter_CreateCoupon(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cfg := Config{Port: 8080}
	router := New(cfg, MockCouponService{})

	reqBody, _ := json.Marshal(map[string]any{
		"discount":       10,
		"code":           "test_code",
		"minBasketValue": 50,
	})
	req, err := http.NewRequest("POST", "/api/create", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}

func TestAPIRouter_GetCouponsByCodes(t *testing.T) {
	gin.SetMode(gin.TestMode)

	cfg := Config{Port: 8080}
	router := New(cfg, MockCouponService{})

	reqBody, _ := json.Marshal(map[string]any{
		"code": "SuperDiscount",
	})
	req, err := http.NewRequest("GET", "/api/coupons", bytes.NewReader(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	router.mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}
}
