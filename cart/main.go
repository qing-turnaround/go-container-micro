package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	consul2 "github.com/micro/go-plugins/registry/consul/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/xing-you-ji/go-container-micro/cart/domain/repository"
	service2 "github.com/xing-you-ji/go-container-micro/cart/domain/service"
	"github.com/xing-you-ji/go-container-micro/cart/handler"
	cart "github.com/xing-you-ji/go-container-micro/cart/proto/cart"
	"github.com/xing-you-ji/go-container-micro/common"
)

var QPS = 100 // 每秒查询率

func main() {

	// 配置中心
	consulConfig, err := common.GetConsulConfig("127.0.0.1	", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}

	// 注册中心
	consulRegistry := consul2.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// 链路追踪（实例化）
	t, io, err := common.NewTracer("go.micro.service.product",
		"127.0.0.1:6831")
	if err != nil {
		log.Fatal(err) // Fatal系列函数会在写入日志信息后调用os.Exit(1)。Panic系列函数会在写入日志信息后panic。
	}
	defer io.Close()               // 关闭数据流
	opentracing.SetGlobalTracer(t) // 设置全局 Tracer

	// 数据库设置
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

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.cart"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8084"),
		// 添加consul为注册中心
		micro.Registry(consulRegistry),
		// 绑定链路追踪
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
		// 添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
	)

	// Initialise service
	service.Init()

	// Register Handler
	cartDataService := service2.NewCartDataService(repository.NewCartRepository(db))
	cart.RegisterCartHandler(service.Server(),
		&handler.Cart{CartDataService: cartDataService})

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
