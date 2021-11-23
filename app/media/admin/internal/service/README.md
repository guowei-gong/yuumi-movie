# Service

> 类似 DDD 的 application 层。

实现了 api 定义的服务层，处理 DTO 到 biz 领域实体的转换(DTO -> DO)，同时协同各类 biz 交互，但是不应处理复杂逻辑。

关注 grpc，主要做协调和编排的逻辑。

DO 是一个贫血模型。