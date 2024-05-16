package api

import (
	"coupon_service/internal/service/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RequestPayload represents the payload for applying a coupon
type RequestPayload struct {
	Code   string        `json:"code"`
	Basket entity.Basket `json:"basket"`
}

// CouponCodePayload represents the payload for getting coupons by codes
type CouponCodePayload struct {
	Codes []string `json:"codes"`
}

// CouponPayload represents the payload for creating a coupon
type CouponPayload struct {
	Discount       int    `json:"discount"`
	Code           string `json:"code"`
	MinBasketValue int    `json:"minBasketValue"`
}

// It returns an error if the JSON parsing or binding fails
func parsePayload(c *gin.Context, payload any) error {
	if err := c.ShouldBindJSON(payload); err != nil {
		return err
	}
	return nil
}

// ApplyCoupon handles the application of a coupon to a basket
func (a *APIRouter) ApplyCoupon(c *gin.Context) {
	var apiReq RequestPayload
	if err := parsePayload(c, &apiReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	basket, err := a.csrv.UpdateBasket(apiReq.Basket, apiReq.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, basket)
}

// CreateCoupon handles the creation of a new coupon
func (a *APIRouter) CreateCoupon(c *gin.Context) {
	var apiReq CouponPayload
	if err := parsePayload(c, &apiReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := a.csrv.Insert(apiReq.Discount, apiReq.Code, apiReq.MinBasketValue); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// GetCouponsByCodes handles fetching coupons by their codes
func (a *APIRouter) GetCouponsByCodes(c *gin.Context) {
	var apiReq CouponCodePayload
	if err := parsePayload(c, &apiReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	coupons, err := a.csrv.GetByCodes(apiReq.Codes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, coupons)
}
