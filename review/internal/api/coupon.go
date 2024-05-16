package api

import (
	"coupon_service/internal/api/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ApplyCoupon handles the application of a coupon to a basket
func (a *API) ApplyCoupon(c *gin.Context) {
	apiReq, err := parseApplicationRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	basket, err := a.svc.UpdateBasket(apiReq.Basket, apiReq.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, basket)
}

// CreateCoupon handles the creation of a new coupon
func (a *API) CreateCoupon(c *gin.Context) {
	apiReq, err := parseCouponRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := a.svc.Insert(apiReq.Discount, apiReq.Code, apiReq.MinBasketValue); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// GetCouponsByCodes handles fetching coupons by their codes
func (a *API) GetCouponsByCodes(c *gin.Context) {
	apiReq, err := parseCouponRequestByCodes(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	coupons, err := a.svc.GetByCodes(apiReq.Codes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, coupons)
}

// Helper function to parse ApplicationRequest from the context
func parseApplicationRequest(c *gin.Context) (entity.ApplicationRequest, error) {
	var apiReq entity.ApplicationRequest
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return apiReq, err
	}
	return apiReq, nil
}

// Helper function to parse CouponRequest from the context
func parseCouponRequest(c *gin.Context) (entity.Coupon, error) {
	var apiReq entity.Coupon
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return apiReq, err
	}
	return apiReq, nil
}

// Helper function to parse CouponRequest by codes from the context
func parseCouponRequestByCodes(c *gin.Context) (entity.CouponRequest, error) {
	var apiReq entity.CouponRequest
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return apiReq, err
	}
	return apiReq, nil
}
