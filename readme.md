# Go容器化微服务

## user模块的开发
1. 快速构建初始代码
* docker pull micro/micro
* docker run --rm -v $(pwd):$(pwd) -w $(pwd) micro/micro new user

2. 编写[user/user.proto](./user/proto/user/user.proto)，来快速生成代码
* docker pull zhugeqing/protoc
* 提前编写好user目录下的[Makefile](./user/Makefile)
* make proto

3. 编写[user.go](./user/domain/model/user.go)和[user_repository.go](./user/domain/repository/user_repository.go)来对数据库和模型进行操作

4. 编写[Handle](user/handler/user.go)要暴露的服务

5. 打包user模块
* make dockerBuild

## 注册配置中心的实现
1. 安装和运行注册中心Consul
* docker pull zhugeqing/consul:latest
* docker run 