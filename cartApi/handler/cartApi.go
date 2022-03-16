package handler

import (
	"context"
	"encoding/json"
	log "github.com/micro/go-micro/v2/logger"

	"cartApi/client"
	"github.com/micro/go-micro/v2/errors"
	api "github.com/micro/go-micro/v2/api/proto"
	cartApi "path/to/service/proto/cartApi"
)

type CartApi struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// CartApi.Call is called by the API as /cartApi/call with post body {"name": "foo"}
func (e *CartApi) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("Received CartApi.Call request")

	// extract the client from the context
	cartApiClient, ok := client.CartApiFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.cartApi.cartApi.call", "cartApi client not found")
	}

	// make request
	response, err := cartApiClient.Call(ctx, &cartApi.Request{
		Name: extractValue(req.Post["name"]),
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.cartApi.cartApi.call", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
