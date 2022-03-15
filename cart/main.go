package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"cart/handler"
	"cart/subscriber"

	cart "cart/proto/cart"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.cart"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	cart.RegisterCartHandler(service.Server(), new(handler.Cart))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.cart", service.Server(), new(subscriber.Cart))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
