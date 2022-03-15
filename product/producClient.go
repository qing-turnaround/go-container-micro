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
	"github.com/xing-you-ji/go-container-micro/common"
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
	t, io, err := common.NewTracer("go.micro.service.client",
		"127.0.0.1:6831")
	if err != nil {
		log.Fatal(err) // Fatal系列函数会在写入日志信息后调用os.Exit(1)。Panic系列函数会在写入日志信息后panic。
	}
	defer io.Close() // 关闭数据流
	opentracing.SetGlobalTracer(t)

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.service.client"),
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
		ProductName:        "彩虹海",
		ProductSku:         "xing-you-ji",
		ProductPrice:       999,
		ProductDescription: "zhugeqing",
		ProductCategoryId:  1,
		ProductImage: []*product.ProductImage{
			{
				ImageName: "444",
				ImageCode: "444",
				ImageUrl:  "zhu4geqing.top",
			},
		},
		ProductSize: []*product.ProductSize{
			{
				SizeName: "b4ig",
				SizeCode: "4",
			},
		},
		ProductSeo: &product.ProductSeo{
			SeoTitle:       "one4 piece",
			SeoKeywords:    "one4 piece",
			SeoDescription: "seo4",
			SeoCode:        "seo4",
		},
	}

	response, err := productService.AddProduct(context.TODO(), productAdd)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(response)

}
