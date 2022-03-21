# Go容器化微服务

## user模块的开发
1. 快速构建初始代码
* docker pull micro/micro
* docker run --rm -v $(pwd):$(pwd) -w $(pwd) micro/micro new user

2. 编写[user.proto](./user/proto/user/user.proto)，来快速生成代码
* docker pull zhugeqing/protoc
* 提前编写好user目录下的[Makefile](./user/Makefile)
* make proto

3. 编写[user.go](./user/domain/model/user.go)和[user_repository.go](./user/domain/repository/user_repository.go)以及[user_data_service.go](./user/domain/service/user_data_service.go)来对数据库和模型进行操作

4. 编写[Handle](user/handler/user.go)要暴露的服务

5. 编写[main.go](user/main.go)

6. 打包user模块
* make dockerBuild

## 注册配置中心的实现
1. 安装和运行注册中心Consul
* docker pull zhugeqing/consul:latest
* docker run -d -p 6666:6666 zhugeqing/consul:latest

2. 快速构建初始代码
* docker pull zhugeqing/micro:2.93
* docker run --rm -v $(pwd):$(pwd) -w $(pwd) zhugeqing/micro:2.93 new category

3. 编写[category.proto](./category/proto/category/category.proto)来快速生成代码

4. 编写[category/domain/](./category/domain)来完成完善领域模型

5. 编写[Handle](category/handler/category.go)要暴露的服务

6. 编写[main.go](category/main.go)

7. 编写关于配置中心的代码并前面配置中心配置相关配置

### Consul注意事项和一些知识
* Consul 使用会有健康检查，不健康的服务会被主动剔除
* Consul 多集群使用过程中注意数据落盘（将数据放入服务器中）
* 多了解Consul的两个重要协议Gossip（八卦协议）、Raft（选举协议）

## 链路追踪
> 用来监视和诊断基于微服务的分布式系统

### 微服务链路追踪（jaeger）主要特性
* 高扩展性
* 原生支持 opentracing
* 可观察性

### 微服务链路追踪（jaeger）的五个重要的组件
* jaeger-client（客户端库）
* Agent（客户端代理）
* Collector（数据收集处理）
* Data Store（数据存储）
* UI（数据查询和前端界面展示）

### 代码开发
1. 安装和运行微服务链路追踪（jaeger）
* docker pull zhugeqing/jaeger:latest
* docker run -d -name jaeger -p 6831:6831/udp -p 16686:16686 zhugeqing/jaeger
> jaeger端口说明：6831协议为UDP，所属模块为agent，功能为通过兼容性Thrift协议，接收jaeger thrift类型协议；
> 16686协议为HTTP，所属模块为query，功能为客户端前端界面展示接口

2. 快速构建初始代码
* docker run --rm -v $(pwd):$(pwd) -w $(pwd) zhugeqing/micro:2.93 new product

3. 编写[product.proto](./product/proto/product/product.proto)来快速生成代码

4. 编写[product/domain/](./product/domain)来完成完善领域模型

5. 编写[Handle](product/handler/product.go)要暴露的服务

6. 编写[main.go](product/main.go)

7. 编写关于链路追踪的代码

8. 编写客户端[productClient](/product/producClient.go)来测试链路追踪

### 链路追踪一些知识
* 链路追踪数据写入的过程中可以加入kafaka缓冲压力
* 可以通过链路追踪发现是否有服务的循环调用

## 熔断，限流，负载均衡

### 代码开发

1. 快速构建初始代码
* docker run --rm -v $(pwd):$(pwd) -w $(pwd) zhugeqing/micro:2.93 new cart

2. 编写[cart.proto](./cart/proto/cart/cart.proto)来快速生成代码

3. 编写[cart/domain/](./cart/domain)来完成完善领域模型

4. 编写[Handle](cart/handler/cart.go)要暴露的服务

5. 编写[main.go](cart/main.go)

6. 限流（漏桶算法）
* go get github.com/micro/go-pligins/wrapper/ratelimiter/uber/v2

7. 创建Api网关
* docker run --rm -v $(pwd):$(pwd) -w $(pwd) zhugeqing/micro:2.93 new --type=api cartApi
* 编写[cartApi.proto](./cartApi/proto/cartApi/cartApi.proto)来快速生成代码
* 编写[handler](./cartApi/handler/cartApi.go)来暴露api服务
* 编写[main.go](./cartApi/main.go)
> 添加熔断
> 添加负载均衡
> 

## 性能监控和订单模块

### 微服务监控系统Prometheus 
* 开源的监控 & 报警 & 时间序列数据库的组合
* 基本原理是通过HTTP协议周期性抓取被监控组件的状态
* 适合Docker, k8s环境

### 开发

1. 快速构建初始代码
* docker run --rm -v $(pwd):$(pwd) -w $(pwd) zhugeqing/micro:2.93 new order

2. 编写[order.proto](./order/proto/order/order.proto)来快速生成代码

3. 编写[order/domain/](./order/domain)来完成完善领域模型

4. 编写[Handle](order/handler/order.go)要暴露的服务

5. 编写[main.go](order/main.go)

6. 编写[docker-compose中的一些配置文件](docker-compose)


7. 通过docker-compose启动服务，测试prometheus和grafana
* grafana默认登陆账号密码均为admin

8. 通过docker-compose运行容器，启动prometheus和grafana

## 微服务日志系统

### ELK日志系统（Elasticsearch + Logstash + Kibana，Beats）
* Elasticsearch：分布式搜索引擎，提供搜集、分析、存储数据三大功能
* Logstash：日志的收集、分析、过滤日志
* Kibana：提供Web界面，帮助汇总、分析、搜索数据
* Beats：轻量级日志收集处理工具
  * Filebeat：日志文件（收集文件数据）
  
### 开发
1. 快速构建初始代码
* docker run --rm -v $(pwd):$(pwd) -w $(pwd) zhugeqing/micro:2.93 new payment

2. 编写[payment.proto](./payment/proto/payment/payment.proto)来快速生成代码

3. 编写[payment/domain/](./payment/domain)来完成完善领域模型

4. 编写[Handle](payment/handler/payment.go)要暴露的服务

5. 编写[main.go](payment/main.go)

6. 创建Api网关
* docker run --rm -v $(pwd):$(pwd) -w $(pwd) zhugeqing/micro:2.93 new --type=api paymentApi
* 编写[cartApi.proto](./cartApi/proto/cartApi/cartApi.proto)来快速生成代码
* 编写[handler](./cartApi/handler/cartApi.go)来暴露api服务
* 编写[main.go](./cartApi/main.go)

7. 部署日志系统
* 安装filebeat`https://www.elastic.co/guide/en/beats/filebeat/8.1/setup-repositories.html#_yum`

