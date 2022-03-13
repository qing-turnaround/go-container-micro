package main

import (
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2"
	"category/handler"
	"category/subscriber"

	category "category/proto/category"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	category.RegisterCategoryHandler(service.Server(), new(handler.Category))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.service.category", service.Server(), new(subscriber.Category))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
