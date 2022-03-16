package main

import (
	log "github.com/micro/go-micro/v2/logger"

	"github.com/micro/go-micro/v2"
	"cartApi/handler"
	"cartApi/client"

	cartApi "cartApi/proto/cartApi"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.cartApi"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the CartApi service client
		micro.WrapHandler(client.CartApiWrapper(service)),
	)

	// Register Handler
	cartApi.RegisterCartApiHandler(service.Server(), new(handler.CartApi))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
