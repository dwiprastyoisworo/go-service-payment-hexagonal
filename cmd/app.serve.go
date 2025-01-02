package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/transport/http"
	gorHdl "github.com/gorilla/handlers"
	"go-service-payment-hexagonal/lib/config"
)

// start service
func startService(cfg *config.Config) {
	_, err := initMongodb(cfg)
	if err != nil {
		log.Fatal("failed to close redis client: %v", err)
	}

	httpOpts := []http.ServerOption{
		http.Timeout(cfg.Apps.Timeout),
		http.Address(cfg.Apps.Address),
		http.Filter(
			gorHdl.CORS(
				gorHdl.AllowedOrigins([]string{"*"}),
				gorHdl.AllowedHeaders([]string{"Content-Type", "Authorization"}),
				gorHdl.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "DELETE"}),
			),
		),
		http.Middleware(
			metadata.Server(),
			logging.Server(log.GetLogger()),
		),
		http.Logger(log.GetLogger()),
	}

	httpServer := http.NewServer(
		httpOpts...,
	)

	server := kratos.New(
		kratos.Name(cfg.Apps.Name),
		kratos.Server(
			httpServer,
		),
	)

	if err := server.Run(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
