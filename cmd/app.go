package main

import "go-service-payment-hexagonal/lib/config"

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	startService(cfg)
}
