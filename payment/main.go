package main

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/xing-you-ji/go-container-micro/common"
	"github.com/xing-you-ji/go-container-micro/payment/domain/repository"
	server2 "github.com/xing-you-ji/go-container-micro/payment/domain/service"
	"github.com/xing-you-ji/go-container-micro/payment/handler"
	payment "github.com/xing-you-ji/go-container-micro/payment/proto/payment"
	"go.uber.org/zap"
)

func main() {
	// 初始化zap
	common.ZapInit()
	// 把缓存区的日志追加到日志
	defer zap.L().Sync()

	// 配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		zap.L().Error("Get consul config error: ", zap.Error(err))
	}

	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer("go.micro.payment", "127.0.0.1:6831")
	if err != nil {
		zap.L().Error("NewTracer error: ", zap.Error(err))
	}
	defer io.Close()

	// 数据库设置
	mysqlInfo := common.GetMysqlConfigFromConsul(consulConfig, "mysql")
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
			mysqlInfo.User, mysqlInfo.Pwd, mysqlInfo.Host, mysqlInfo.Port, mysqlInfo.Database))
	defer db.Close()
	// db.SingularTable(true) 让gorm转义struct名字的时候不用加上s
	db.SingularTable(true)
	// 初始化表 只执行一次
	// repository.NewPaymentRepository(db).InitTable()
	zap.L().Error("测试呀", zap.Error(errors.New("我试试zap")))

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.payment"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8087"),
		// 注册中心
		micro.Registry(consulRegistry),
		// 链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(t)),
		// 限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(1000)),
		// 加载监控
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)

	// Initialise service
	service.Init()

	paymentDataService := server2.NewPaymentDataService(repository.NewPaymentRepository(db))
	// Register Handler
	payment.RegisterPaymentHandler(service.Server(), &handler.Payment{
		PaymentDataService: paymentDataService,
	})

	// Run service
	if err := service.Run(); err != nil {
		zap.L().Fatal("service run error: ", zap.Error(err))
	}
}
