# 限流服务提供接口 HTTP API
- [限流接口](#限流接口)
- [字段定义](#字段定义)

### 错误码定义

| 错误码 | 说明           |
|-----| -------------- |
| 0   | 成功           |


### 限流接口
- 路由: http://cros-rate-limit-service/api/v1/rate-limit
- 请求方法: GET
- 参数:

  | 字段  | 说明  |
  |-----| -----------|
  | service   | 服务名 |
  | limit_key | 限流唯一key |
  | limit_method |限流方式 |
```shell
curl http://cros-rate-limit-service/api/v1/rate-limit?service=anti-fraud&limit_key=apply&limit_method=
```
- 返回值
```json
{
  "code": 0, // 错误码
  "data": {
    "service": "anti-fraud",
    "limit_key": "apply",
    "limit_result": "accept" // 通过 accept 拒绝 reject
  },
  "msg": ""
}
```

### 限流配置  
- 路由: http://cros-rate-limit-service/api/v1/rate-limit-config
- 请求方法: POST
  - 参数:

    | 字段  | 说明  | 类型
    |-----| -----------|------- |
    | service   | 服务名 | str |
    | limit_key | 限流唯一key | str |
    | limit_algorithm | 限流算法 | str :PeriodLimit or TokenLimit|
    | period | 窗口大小 | int |
    | quota | 请求上限 | int |
    | rate | 每秒生产速率 | int |
    | burst | 桶容量 | int |
  - 返回值:
```json
{
  "code": 0, // 错误码
  "data": {
    "config_status": "success"
  },
  "msg": ""
}
```
### 请求限流
- 路由 http://cros-rate-limit-service/api/v1/rate-limit
- 请求方法: GET
- 参数

    | 字段 | 描述 | 类型
    | ---- | ----- | ----- |
    | service_name | 服务名 | str
    | limit_key| 限流key | str
- 返回值:
```json
{
  "code": 0,   // 错误码
  "data": {}, // 
  "msg": ""
}
```