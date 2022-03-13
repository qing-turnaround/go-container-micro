# Category service 
Go微服务定制容器
名称为 Category 类型 service 

创建初始化模版请使用

```
micro new category --namespace=go.micro --type=service
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.service.category
- Type: service
- Alias: category

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
./category-service
```

Build a docker image
```
make docker
```