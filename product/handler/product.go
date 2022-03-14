package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	product "product/proto/product"
)

type Product struct{}

// 这里生产实现服务方法
 

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Product) Stream(ctx context.Context, req *product.StreamingRequest, stream product.Product_StreamStream) error {
	log.Infof("Received Product.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&product.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}
 
