package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// 数据库驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"github.com/xing-you-ji/go-container-micro/category/common"
	"github.com/xing-you-ji/go-container-micro/category/domain/repository"
	service2 "github.com/xing-you-ji/go-container-micro/category/domain/service"
	"github.com/xing-you-ji/go-container-micro/category/handler"
	category "github.com/xing-you-ji/go-container-micro/category/proto/category"
)

func main() {
	// 配置中心
	consulConfig, err := common.GetConsulConfig("120.79.17.230", 8500, "/micro/config")
	if err != nil {
		log.Error(err)
	}

	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		// 设置地址和暴露的端口
		micro.Address("127.0.0.1:8082"),
		// 添加consul为注册中心
		micro.Registry(consulRegistry),
	)

	// 获取mysql配置，路径不带前缀
	mysqlInfo := common.GetMysqlConfigFromConsul(consulConfig, "mysql")
	fmt.Println(mysqlInfo)
	// root:unraveltheworld@tcp(120.79.17.230:3307)/micro?charset=utf8mb4&parseTime=true&loc=Local
	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local",
			mysqlInfo.User, mysqlInfo.Pwd, mysqlInfo.Host, mysqlInfo.Port, mysqlInfo.Database))
	if err != nil {
		log.Error(err)
	}
	defer db.Close()
	// db.SingularTable(true) 让grom转义struct名字的时候不用加上s
	db.SingularTable(true)
	rp := repository.NewCategoryRepository(db)
	// 初始化表 只执行一次
	rp.InitTable()

	// Initialise service
	service.Init()

	// 创建 CategoryDataService
	categoryDataService := service2.NewCategoryDataService(repository.NewCategoryRepository(db))
	// Register Handler
	err = category.RegisterCategoryHandler(service.Server(), &handler.Category{
		categoryDataService,
	})
	if err != nil {
		log.Error(err)
	}

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
