package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	category "category/proto"
)

type Category struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Category) Call(ctx context.Context, req *category.Request, rsp *category.Response) error {
	log.Info("Received Category.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

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

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Category) PingPong(ctx context.Context, stream category.Category_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&category.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
