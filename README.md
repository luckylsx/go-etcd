## go etcd 使用示例：

### 注册：
```
points := []string{"http://127.0.0.1:2379"}
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
registry, err := registry.New(points, WithNameSpace("microservice"), registry.WithCtx(ctx))
instance := registry.ServiceInstance{
    ID:   "1",
    Name: "test",
}
err = registry.Register(instance)
if err != nil {
    t.Error(err)
}
cancel()
```

### 查询：
```
points := []string{"http://127.0.0.1:2379"}
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
registry, err := registry.New(points, WithNameSpace("microservice"), registry.WithCtx(ctx))
instance := registry.ServiceInstance{
    ID:   "1",
    Name: "test",
}
if err != nil {
    log.fatal(err)
}
res, err := registry.GetService(instance)
if err != nil {
    log.fatal(err)
}
cancel()
log.printf("%+v", res)
```