# Go容器化微服务

## user模块的开发
1. 快速构建初始代码
* docker pull micro/micro
* docker run --rm -v $(pwd):$(pwd) -w $(pwd) micro/micro new user

2. 编写[user.proto](./user/proto/user/user.proto)，来快速生成代码
* docker pull zhugeqing/protoc
* 提前编写好user目录下的[Makefile](./user/Makefile)
* make proto