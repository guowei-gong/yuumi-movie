# /api

`API` 协议定义目录，它就像字典。

存放 `protobuf` 文件以及生成的 `go` 文件。

```
api
├── admin
├── interface
└── service
```

- `admin`: 区别于 `service`，更多是面向运营侧的服务，通常数据权限更高，单独隔离是为了带来更好的代码级别安全。
- `interface`: 对外的 `BFF` 服务，接受来自用户的请求，例如暴露 `HTTP/gRPC` 接口。
- `service`: 纯对内的服务，仅接受来自内部其他服务或者网关的请求。
