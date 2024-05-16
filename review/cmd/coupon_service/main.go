package main

import (
	"coupon_service/internal/api"
	"coupon_service/internal/config"
	"coupon_service/internal/repository/memdb"
	couponService "coupon_service/internal/service"
	"log"
	"time"
)

func main() {
	// Load the application configuration
	cfg := config.New()

	// Initialize data
	data := memdb.New()

	// Initialize the coupon service
	couponSvc := couponService.New(data)

	// Initialize the API server
	apiServer := api.New(cfg.APIConfig, couponSvc)

	// Start the API server
	err := apiServer.Start()
	if err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}

	defer apiServer.Close()

	log.Printf("Coupon service started at %s", time.Now().Format("2006-01-02 15:04:05"))
	<-time.After(1 * time.Hour * 24 * 365)
	log.Println("Coupon service has been running for a year. Shutting down gracefully...")
}
