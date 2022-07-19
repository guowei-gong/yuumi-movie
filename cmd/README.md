# cmd
项目主干。分为如下 4 类服务类型。

```
cmd
├── admin
├── interface
├── job
├── service
└── task
```

- `admin`: 区别于 `service`，更多是面向运营侧的服务，通常数据权限更高，单独隔离是为了带来更好的代码级别安全。
- `interface`: 对外的 `BFF` 服务，接受来自用户的请求，例如暴露 `HTTP/gRPC` 接口。
- `service`: 纯对内的服务，仅接受来自内部其他服务或者网关的请求。
- `job`: 流式任务处理的服务，上游一般依赖 `message broker`，例如 `Kafka`、订阅 `MySQL binlog`。
- `task`: 定时任务，类似 `cronjob`，如果任务简单，可以撸简单一点，直接干一个 `main.go`。

