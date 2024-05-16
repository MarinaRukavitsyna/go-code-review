package api

import (
	"coupon_service/internal/service/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RequestPayload struct {
	Code   string        `json:"code"`
	Basket entity.Basket `json:"basket"`
}

type CouponCodePayload struct {
	Codes []string `json:"codes"`
}

type CouponPayload struct {
	Discount       int    `json:"discount"`
	Code           string `json:"code"`
	MinBasketValue int    `json:"minBasketValue"`
}

// ApplyCoupon handles the application of a coupon to a basket
func (a *API) ApplyCoupon(c *gin.Context) {
	apiReq, err := parseAppRequest(c)
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
	apiReq, err := parseCouponCodesRequest(c)
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
func parseAppRequest(c *gin.Context) (RequestPayload, error) {
	var apiReq RequestPayload
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return apiReq, err
	}
	return apiReq, nil
}

// Helper function to parse CouponRequest from the context
func parseCouponRequest(c *gin.Context) (CouponPayload, error) {
	var apiReq CouponPayload
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return apiReq, err
	}
	return apiReq, nil
}

// Helper function to parse CouponRequest by codes from the context
func parseCouponCodesRequest(c *gin.Context) (CouponCodePayload, error) {
	var apiReq CouponCodePayload
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		return apiReq, err
	}
	return apiReq, nil
}
