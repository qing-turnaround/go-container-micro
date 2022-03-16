package client

import (
	"context"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/server"
	cartApi "path/to/service/proto/cartApi"
)

type cartApiKey struct {}

// FromContext retrieves the client from the Context
func CartApiFromContext(ctx context.Context) (cartApi.CartApiService, bool) {
	c, ok := ctx.Value(cartApiKey{}).(cartApi.CartApiService)
	return c, ok
}

// Client returns a wrapper for the CartApiClient
func CartApiWrapper(service micro.Service) server.HandlerWrapper {
	client := cartApi.NewCartApiService("go.micro.service.template", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, cartApiKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}
