package main

import (
	"context"
	"net"
	"net/http"

	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	roundrobin "github.com/micro/go-plugins/wrapper/select/roundrobin/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/xing-you-ji/go-container-micro/common"
	payment "github.com/xing-you-ji/go-container-micro/payment/proto/payment"
	"github.com/xing-you-ji/go-container-micro/paymentApi/handler"
	"go.uber.org/zap"

	"github.com/micro/go-micro/v2"

	paymentApi "github.com/xing-you-ji/go-container-micro/paymentApi/proto/paymentApi"
)

func main() {
	common.ZapInit()

	// 注册中心
	consulRegistry := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer("go.micro.api.paymen", "127.0.0.0:6831")
	defer io.Close()
	if err != nil {
		zap.L().Error("common.NewTracer error", zap.Error(err))
	}
	opentracing.SetGlobalTracer(t)

	// 熔断器
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()

	// 启动监听
	go func() {
		err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", "9192"), hystrixStreamHandler)
		if err != nil {
			log.Error(err)
		}
	}()

	// 监控
	common.Promethues(9092)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.paymentApi"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8088"),
		// 注册中心
		micro.Registry(consulRegistry),
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
		// 负载均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)

	// Initialise service
	service.Init()

	paymentService := payment.NewPaymentService("go.micro.server.payment",
		service.Client())

	// Register Handler
	if err := paymentApi.RegisterPaymentApiHandler(service.Server(),
		&handler.PaymentApi{PaymentService: paymentService}); err != nil {
		zap.L().Error("", zap.Error(err))

	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		// run 正常执行
		zap.L().Info(req.Service() + "." + req.Endpoint())
		return c.Client.Call(ctx, req, rsp, opts...)
	}, func(err error) error {
		zap.L().Error("", zap.Error(err))
		return err
	})
}

func NewClientHystrixWrapper() client.Wrapper {
	return func(i client.Client) client.Client {
		return &clientWrapper{i}
	}
}
