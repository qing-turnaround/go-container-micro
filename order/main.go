package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/xing-you-ji/go-container-micro/common"
	"github.com/xing-you-ji/go-container-micro/order/domain/repository"
	service2 "github.com/xing-you-ji/go-container-micro/order/domain/service"

	"github.com/xing-you-ji/go-container-micro/order/handler"
	order "github.com/xing-you-ji/go-container-micro/order/proto/order"
)

var (
	QPS = 1000
)

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}
	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// 链路追踪
	t, io, err := common.NewTracer("go.micro.service.product",
		"127.0.0.1:6831")
	if err != nil {
		log.Fatal(err) // Fatal系列函数会在写入日志信息后调用os.Exit(1)。Panic系列函数会在写入日志信息后panic。
	}
	defer io.Close() // 关闭数据流
	opentracing.SetGlobalTracer(t)

	// 初始化数据库
	mysqlInfo := common.GetMysqlConfigFromConsul(consulConfig, "mysql") // 获取mysql配置，路径不带前缀
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
			mysqlInfo.User, mysqlInfo.Pwd, mysqlInfo.Host, mysqlInfo.Port, mysqlInfo.Database))
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	// db.SingularTable(true) 让gorm转义struct名字的时候不用加上s
	db.SingularTable(true)
	// 初始化表 只执行一次
	// repository.NewOrderRepository(db).InitTable()

	// 实例化 orderDataService
	orderDataService := service2.NewOrderDataService(repository.NewOrderRepository(db))

	// 暴露监控地址
	common.Promethues(9092)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.order"),
		micro.Version("latest"),
		micro.Registry(consulRegistry),
		// 暴露的服务地址
		micro.Address("127.0.0.1:8086"),
		// 添加链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		// 添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
		// 添加监控
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)
	// Initialise service
	service.Init()

	// Register Handler
	order.RegisterOrderHandler(service.Server(),
		&handler.Order{OrderDataService: orderDataService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}

}
