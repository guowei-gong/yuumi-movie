#/app

app 目录，存放每个微服务的具体代码。

app 的命名是一个全局唯一的名称。按照 `业务 + 服务`，例如：`video.service`、`courier.job`。

其实可以考虑 DNS 的三段式命名，变为 `业务 + 服务 + 子服务`，这无疑增加了更多的工作量，但是让可用性提高了。