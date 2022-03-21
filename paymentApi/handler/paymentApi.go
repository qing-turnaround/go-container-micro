package handler

import (
	"context"
	"encoding/json"
	log "github.com/micro/go-micro/v2/logger"

	"paymentApi/client"
	"github.com/micro/go-micro/v2/errors"
	api "github.com/micro/go-micro/v2/api/proto"
	paymentApi "path/to/service/proto/paymentApi"
)

type PaymentApi struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// PaymentApi.Call is called by the API as /paymentApi/call with post body {"name": "foo"}
func (e *PaymentApi) Call(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Info("Received PaymentApi.Call request")

	// extract the client from the context
	paymentApiClient, ok := client.PaymentApiFromContext(ctx)
	if !ok {
		return errors.InternalServerError("go.micro.api.paymentApi.paymentApi.call", "paymentApi client not found")
	}

	// make request
	response, err := paymentApiClient.Call(ctx, &paymentApi.Request{
		Name: extractValue(req.Post["name"]),
	})
	if err != nil {
		return errors.InternalServerError("go.micro.api.paymentApi.paymentApi.call", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}
