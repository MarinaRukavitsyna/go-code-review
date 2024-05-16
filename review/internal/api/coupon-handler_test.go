package api

import (
	"bytes"
	"coupon_service/internal/service/entity"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestParsePayload(t *testing.T) {
	tests := []struct {
		name          string
		jsonData      string
		payloadType   any
		expected      any
		expectedError bool
	}{
		// Tests for RequestPayload
		{
			name:        "Valid RequestPayload",
			jsonData:    `{"code": "SuperDiscount", "basket": {"id": "basket1", "value": 100, "appliedDiscount": 10, "applicationSuccessful": true}}`,
			payloadType: &RequestPayload{},
			expected: &RequestPayload{
				Code: "SuperDiscount",
				Basket: entity.Basket{
					ID:                    "basket1",
					Value:                 100,
					AppliedDiscount:       10,
					ApplicationSuccessful: true,
				},
			},
			expectedError: false,
		},
		{
			name:          "Invalid RequestPayload - Invalid JSON",
			jsonData:      `{"code" "invalid"}`,
			payloadType:   &RequestPayload{},
			expected:      &RequestPayload{},
			expectedError: true,
		},

		// Tests for CouponPayload
		{
			name:        "Valid CouponPayload",
			jsonData:    `{"discount": 10, "code": "SuperDiscount", "minBasketValue": 50}`,
			payloadType: &CouponPayload{},
			expected: &CouponPayload{
				Discount:       10,
				Code:           "SuperDiscount",
				MinBasketValue: 50,
			},
			expectedError: false,
		},
		{
			name:          "Invalid CouponPayload - - Invalid JSON",
			jsonData:      `{"discount" 10, "code": "SuperDiscount"}`,
			payloadType:   &CouponPayload{},
			expected:      &CouponPayload{},
			expectedError: true,
		},

		// Tests for CouponCodePayload
		{
			name:        "Valid CouponCodePayload",
			jsonData:    `{"codes":["SuperDiscount", "MegaDiscount", "FinalSale"]}`,
			payloadType: &CouponCodePayload{},
			expected: &CouponCodePayload{
				Codes: []string{"SuperDiscount", "MegaDiscount", "FinalSale"},
			},
			expectedError: false,
		},
		{
			name:          "Invalid CouponCodePayload - - Invalid JSON",
			jsonData:      `{"codes":["SuperDiscount",] "MegaDiscount", "FinalSale"]}`,
			payloadType:   &CouponCodePayload{},
			expected:      &CouponCodePayload{},
			expectedError: true,
		},

		// Tests for invalid JSON
		{
			name:          "Invalid JSON",
			jsonData:      `{"codes":"invalid"}`,
			payloadType:   &CouponCodePayload{},
			expected:      &CouponCodePayload{},
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up the Gin context
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			c.Request, _ = http.NewRequest(http.MethodPost, "/", bytes.NewBufferString(tt.jsonData))
			c.Request.Header.Set("Content-Type", "application/json")

			// Call the function to test
			err := parsePayload(c, tt.payloadType)

			// Check for expected errors
			if tt.expectedError {
				assert.Error(t, err)
			} else {
				// No error expected
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, tt.payloadType)
			}
		})
	}
}
