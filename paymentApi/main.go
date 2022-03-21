package main

import (
	log "github.com/micro/go-micro/v2/logger"

	"github.com/micro/go-micro/v2"
	"paymentApi/handler"
	"paymentApi/client"

	paymentApi "paymentApi/proto/paymentApi"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.paymentApi"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the PaymentApi service client
		micro.WrapHandler(client.PaymentApiWrapper(service)),
	)

	// Register Handler
	paymentApi.RegisterPaymentApiHandler(service.Server(), new(handler.PaymentApi))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
