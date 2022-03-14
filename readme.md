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
* docker pull zhugeqing/micro:latest
* docker run --rm -v $(pwd):$(pwd) -w $(pwd) zhugeqing/micro:latest new category

3. 编写[category.proto](./category/proto/category/category.proto)，来快速生成代码

4. 编写[category/domain/](./category/domain)来完成完善领域模型

5. 编写[Handle](category/handler/category.go)要暴露的服务

6. 编写[main.go](category/main.go)

7. 编写关于配置中心的代码并前面配置中心配置相关配置

### Consul注意事项和一些知识
* Consul 使用会有健康检查，不健康的服务会被主动剔除
* Consul 多集群使用过程中注意数据落盘（将数据放入服务器中）
* 多了解Consul的两个重要协议Gossip（八卦协议）、Raft（选举协议）