# 首次上传

> 大家想运行http模块的话，在http/etc修改一下配置文件，改成自己的
> 然后运行是`go run main.go`


## 参考配置

http/etc/config.yaml

```yaml
server:
  http:
    name: "backend"
    mode: "dev"
    host: "0.0.0.0" # 局域网内部允许访问
    port: "3000"
  interaction:
    host: "0.0.0.0"
    port: "3001"
  relation:
    host: "0.0.0.0"
    port: "3002"
#之后的微服务会再添加一些服务

# log配置可以先不管，目前还没配置项目的log
log:
  level: "debug"
  infoFilename: "./log/info.log"
  errFilename: "./log/error.log"
  max_size: 200
  max_age: 30
  max_backups: 7

mysql:
  host: "127.0.0.1"
  port: "3306"
  user: "root"
  password: "123456"
  dbname: "tiktok"
  max_idle_conns: 10
  max_open_conns: 30

#redis:
#  host: "127.0.0.1"
#  port: 6379
#  password: "123456"
#  db: 0

oss:
  ALI_AK: "××××××××××××××××××××××××××××"
  ALI_SK: "××××××××××××××××××××××××××××"
  ALI_OSS_ENDPOINT: "oss-cn-beijing.aliyuncs.com"
  ALI_BUCKET_NAME: "××××××××××××××××"
  PLAY_URL_PREFIX: "××××××××××"
```