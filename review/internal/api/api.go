package api

import (
	"context"
	"coupon_service/internal/service/entity"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Defines the interface for the coupon-related services
type CouponService interface {
	UpdateBasket(entity.Basket, string) (*entity.Basket, error)
	Insert(int, string, int) (string, error)
	GetByCodes([]string) ([]entity.Coupon, error)
}

// Hold the configuration settings for the API server
type Config struct {
	Host string
	Port int
}

// Encapsulate the HTTP server, Gin engine, service, and configuration
type APIRouter struct {
	srv  *http.Server
	mux  *gin.Engine
	csrv CouponService
	cfg  Config
}

// Initialize a new APIRouter instance with the given configuration and service
func New(cfg Config, svc CouponService) APIRouter {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	apiRouter := APIRouter{
		mux:  r,
		cfg:  cfg,
		csrv: svc,
	}

	return apiRouter.withServer().withRoutes()
}

// Configure the HTTP server for the APIRouter instance
func (a APIRouter) withServer() APIRouter {
	a.srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", a.cfg.Port),
		Handler: a.mux,
	}
	return a
}

// Set up the API routes and returns the APIRouter instance
func (a APIRouter) withRoutes() APIRouter {
	apiGroup := a.mux.Group("/api")
	apiGroup.POST("/apply", a.ApplyCoupon)
	apiGroup.POST("/create", a.CreateCoupon)
	apiGroup.GET("/coupons", a.GetCouponsByCodes)
	return a
}

// Start the HTTP server
func (a APIRouter) Start() error {
	if err := a.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

// Close gracefully shuts down the HTTP server.
func (a APIRouter) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}
