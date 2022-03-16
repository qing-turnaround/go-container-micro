# CartApi api 
Go微服务定制容器
名称为 CartApi 类型 api 

创建初始化模版请使用

```
micro new cartApi --namespace=go.micro --type=api
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.api.cartApi
- Type: api
- Alias: cartApi

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./cartApi-api
```

Build a docker image
```
make docker
```