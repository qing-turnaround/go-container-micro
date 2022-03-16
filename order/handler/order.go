package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	order "order/proto/order"
)

type Order struct{}

// 这里生产实现服务方法
 

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Order) Stream(ctx context.Context, req *order.StreamingRequest, stream order.Order_StreamStream) error {
	log.Infof("Received Order.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&order.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}
 
