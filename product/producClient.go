package main

import (
	"context"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/xing-you-ji/go-container-micro/product/common"
	product "github.com/xing-you-ji/go-container-micro/product/proto/product"
)

func main() {

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

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.product"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8083"),
		// 添加consul为注册中心
		micro.Registry(consulRegistry),
		// 绑定链路追踪
		// 客户端绑定Client，服务端使用handler
		micro.WrapClient(opentracing2.NewClientWrapper(opentracing.GlobalTracer())),
	)

	productService := product.NewProductService("go.micro.service.product", service.Client())

	productAdd := &product.ProductInfo{
		ProductName:        "蕾贝卡",
		ProductSku:         "cap",
		ProductPrice:       999,
		ProductDescription: "哈哈哈哈",
		ProductCategoryId:  1,
		ProductImage: []*product.ProductImage{
			{
				ImageName: "zhugeqing-image",
				ImageCode: "666",
				ImageUrl:  "zhugeqing.top",
			},
		},
		ProductSize: []*product.ProductSize{
			{
				SizeName: "大",
				SizeCode: "1",
			},
		},
		ProductSeo: &product.ProductSeo{
			SeoTitle:       "海贼王蕾贝卡",
			SeoKeywords:    "蕾贝卡",
			SeoDescription: "seo",
			SeoCode:        "seo",
		},
	}

	response, err := productService.AddProduct(context.TODO(), productAdd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)

}
