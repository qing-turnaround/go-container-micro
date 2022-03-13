package handler

import (
	"context"

	log "github.com/micro/go-micro/v2/logger"

	category "category/proto/category"
)

type Category struct{}

// 这里生产实现服务方法
 

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Category) Stream(ctx context.Context, req *category.StreamingRequest, stream category.Category_StreamStream) error {
	log.Infof("Received Category.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&category.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}
 
