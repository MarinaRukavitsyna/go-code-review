package main

import (
	"coupon_service/internal/api"
	"coupon_service/internal/config"
	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/service"
	"log"
	"time"
)

var (
	cfg  = config.New()
	repo = memdb.New()
)

func main() {
	srv := service.New(repo)
	con := api.New(cfg.API, srv)
	con.Start()
	defer con.Close()

	log.Printf("Starting Coupon service")
	<-time.After(1 * time.Hour * 24 * 365)
	log.Println("Coupon service server alive for a year, closing")
}
