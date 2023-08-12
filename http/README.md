# nunu-layout-basic — 基础布局

## 功能
- **Gin**: https://github.com/gin-gonic/gin
- **Gorm**: https://github.com/go-gorm/gorm
- **Wire**: https://github.com/google/wire
- **Viper**: https://github.com/spf13/viper
- **Zap**: https://github.com/uber-go/zap
- **Golang-jwt**: https://github.com/golang-jwt/jwt
- **Go-redis**: https://github.com/go-redis/redis
- More...

## 目录结构
```
.
├── config
│   ├── config.go
│   ├── log.go
│   └── load.go
├── internal
│   ├── middleware
│   │   └── cors.go
│   ├── models
│   │   └── user.go
│   ├── dao
│   │   ├── db
│   │   │  └── mysql.go
│   │   └── user.go
│   ├── pkg
│   │   ├── result
│   │   │  └── result.go
│   ├── user
│   │   ├── http
│   │   ├── user.go
│   │   └── impl
├── pkg
├── LICENSE
├── README.md
├── README_zh.md
├── go.mod
└── go.sum

```

这是一个经典的Golang 项目的目录结构，包含以下目录：
- config: 存放应用程序的配置文件。
  - config.go: 配置文件对应的go的结构体。
  - load.go: 加载配置文件的代码。

- internal: 存放应用程序的内部代码。
  - middleware: 存放中间件代码。
    - cors.go: 跨域资源共享中间件。
  - models: 存放数据模型代码。
    - user.go: 用户数据模型。
  - dao: 存放数据访问代码。
    - db: 和数据库连接相关
      - mysql.go: 初始化数据库连接。
    - user.go: 用户数据访问接口的实现。
  - pkg: 存放app中用到的公共包，但不想被外部引用的。  
    * result
      * result.go    返回结果使用的包，定义各种状态码    
  - user: userApp对应的包
    - http: 
      - user.go: 相当于controller层
    - impl: 接口实现层相当于java中的impl层
      - user.go: 接口实现
      - user_test.go: 接口测试
    - user.go: 接口定义相当于service层
  - router
    - router.go 路由注册的地方

- pkg: 存放应用程序的公共包。
- storage: 存放应用程序的存储文件。
- go.mod: Go模块文件。
- go.sum: Go模块的依赖版本文件。