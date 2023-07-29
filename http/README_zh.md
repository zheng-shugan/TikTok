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
├── cmd
│   ├── root.go
│   └── start.go
├── config
│   ├── config.go
│   ├── log.go
│   └── load.go
├── internal
│   ├── all
│   │   └── impl.go
│   ├── middleware
│   │   └── cors.go
│   ├── models
│   │   └── user.go
│   ├── dao
│   │   ├── db
│   │   │  ├── mysql.go
│   │   │  └── user.go
│   │   └── script
│   │      └── gorm_gen.go
│   ├── pkg
│   │   ├── app.go
│   │   └── ioc.go
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

- cmd: 存放应用程序的入口点，包括主函数和依赖注入的代码。
  - root.go: 根命令行
  - start.go: 启动相关命令

- config: 存放应用程序的配置文件。
  - config.go: 配置文件对应的go的结构体。
  - load.go: 加载配置文件的代码。

- internal: 存放应用程序的内部代码。
  - middleware: 存放中间件代码。
    - cors.go: 跨域资源共享中间件。
  - all: 把每个包中的init函数放在此处匿名实现
    - impl.go
  - models: 存放数据模型代码。
    - user.go: 用户数据模型。
  - dao: 存放数据访问代码。
    -db: 和数据库连接相关 
      - mysql.go: 初始化数据库连接。
      - user.go: 用户数据访问接口的实现。
  - pkg: 存放app中用到的公共包，但不想被外部引用的。
    - ioc.go: ioc的逻辑在这里面
    - apps.go: 定义每个app的名字
  - user: userApp对应的包
    - http: 希望外部访问的，暴露url的地方
      - http.go: 把app对应的http注册到ioc之中
      - user.go: 相当于controller层
    - impl: 接口实现层相当于java中的impl层
      - user.go: 接口实现
      - user_test.go: 接口测试
    - user.go: 接口定义相当于service层
- pkg: 存放应用程序的公共包。
- storage: 存放应用程序的存储文件。
- go.mod: Go模块文件。
- go.sum: Go模块的依赖版本文件。

## 要求
要使用Nunu，您需要在系统上安装以下软件：

* Golang 1.16或更高版本
* Git



### 安装

您可以通过以下命令安装Nunu：

```bash
go install github.com/go-nunu/nunu@latest
```


### 创建新项目

您可以使用以下命令创建一个新的Golang项目：

```bash
nunu new projectName
```
默认拉取github源，你也可以使用国内加速仓库
```
// 使用基础模板
nunu new projectName -r https://gitee.com/go-nunu/nunu-layout-basic.git
// 使用高级模板
nunu new projectName -r https://gitee.com/go-nunu/nunu-layout-advanced.git
```

此命令将创建一个名为`projectName`的目录，并在其中生成一个优雅的Golang项目结构。

### 创建组件

您可以使用以下命令为项目创建handler、service和dao等组件：

```bash
nunu create handler user
nunu create service user
nunu create repository user
nunu create model user
```
或
```
nunu create all user
```
这些命令将分别创建一个名为`UserHandler`、`UserService`、`UserDao`和`UserModel`的组件，并将它们放置在正确的目录中。

### 启动项目

您可以使用以下命令快速启动项目：

```bash
nunu run
```

此命令将启动您的Golang项目，并支持文件更新热重启。

### 编译wire.go

您可以使用以下命令快速编译`wire.go`：

```bash
nunu wire
```

此命令将编译您的`wire.go`文件，并生成所需的依赖项。

## 贡献

如果您发现任何问题或有任何改进意见，请随时提出问题或提交拉取请求。我们非常欢迎您的贡献！

## 许可证

Nunu是根据MIT许可证发布的。有关更多信息，请参见[LICENSE](LICENSE)文件。