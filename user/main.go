package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/micro/go-micro/v2"
	"github.com/xing-you-ji/go-container-micro/user/domain/repository"
	service2 "github.com/xing-you-ji/go-container-micro/user/domain/service"
	"github.com/xing-you-ji/go-container-micro/user/handler"
	user "github.com/xing-you-ji/go-container-micro/user/proto/user"

	// 数据库驱动
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// 服务参数设置
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8081"),
	)
	// 初始化服务
	srv.Init()

	// 创建数据库连接
	db, err := gorm.Open("mysql",
		"root:unraveltheworld@tcp(120.79.17.230:3307)/micro?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	// db.SingularTable(true) 让grom转义struct名字的时候不用加上s
	db.SingularTable(true)
	rp := repository.NewUserRepository(db)
	// 初始化表 只执行一次
	rp.InitTable()
	// 创建服务实例
	userDataService := service2.NewUserDataService(rp)

	// 注册Handler
	err = user.RegisterUserHandler(srv.Server(),
		&handler.User{UserDataService: userDataService})
	if err != nil {
		fmt.Println(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
