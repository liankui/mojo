# Mojo

## What can do

1. Code generator
2. Database query language for relational, document and graph database (inspired by the RQL of the Rethink_d_b)
   1. SQL
   2. Graph_q_l
   3. [Cypher Query Language](https://en.wikipedia.org/wiki/Cypher_Query_Language)
3. Script for data process

## How to use mojo
### base command lines
```bash
$ mojo -h
NAME:
   chaosmojo - A new cli application

USAGE:
   chaosmojo [global options] command [command options] [arguments...]

COMMANDS:
   build    build the mojo source files
   create   create the scaffolding code for mojo module
   format   format the mojo code source
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help
```

### mojo create
```bash
$ mojo create -h     
NAME:
   chaosmojo create - create the scaffolding code for mojo module

USAGE:
   chaosmojo create [command options] [arguments...]

OPTIONS:
   --package value, -p value          the mojo package name to compile
   --output value, -o value           the Output to generate (default: .)
   --repo value, -r value             the git repository of NCraft releated code
   --version value, -v value          the version of the package to be created (default: 0.1.0)
   --license value, -l value          the license of the package to be created
   --Author value, -a value           the license of the package to be created (default: YOUR NAME)
   --email value, -e value            the email of the author of the package to be created (default: YOUR EMAIL)
   --organization value, --org value  the organization of the package to be created, for generating java code (default: YOUR ORGANIZATION (like: mojolang.org))
   --run                              run the hello world server immediately (default: false)
   --help, -h                         show help
```
注意：如果不使用--package，这时会生成默认的hello-world项目，结构如下
```bash
$ mojo create
$ tree hello-world
hello-world
├── README.md
├── document
│   └── hello-world
│       └── v1
│           ├── echo.md
│           └── hello_world.md
├── go
│   ├── go.mod
│   ├── go.sum
│   └── pkg
│       └── hello-world
│           └── v1
│               ├── echo.pb.go
│               ├── hello_world.pb.go
│               └── hello_world_grpc.pb.go
├── mojo
│   └── hello-world
│       └── v1
│           ├── echo.mojo
│           └── hello_world.mojo
├── openapi
│   └── hello-world
│       └── v1
│           ├── echo.schema.json
│           └── hello_world.yaml
├── package.mojo
├── protobuf
│   └── hello-world
│       └── v1
│           ├── echo.proto
│           └── hello_world.proto
└── service-go
    ├── README.md
    ├── build
    │   └── hello-world-server-docker
    │       └── Docerfile
    ├── cmd
    │   ├── hello-world-server
    │   │   └── main.go
    │   └── hello-world-unified-server
    │       └── main.go
    ├── configs
    │   ├── logs.yaml
    │   ├── metrics.yaml
    │   ├── sd.yaml
    │   ├── server.yaml
    │   └── tracing.yaml
    ├── go.mod
    ├── go.sum
    ├── internal
    │   ├── hello-world-server
    │   │   └── run.go
    │   └── server
    │       └── run.go
    ├── pkg
    │   └── hello-world-service
    │       ├── handlers
    │       │   ├── handlers.go
    │       │   ├── hooks.go
    │       │   └── middlewares.go
    │       └── svc
    │           ├── endpoints.go
    │           ├── transport_grpc.go
    │           └── transport_http.go
    └── scripts
        └── run.sh

32 directories, 35 files
```
## 实践指南
### 1.指定package生成代码
在任一目录下使用如下命令生成user服务：
```bash
mojo create --package user --output . --repo github.com/xxx/xxx/user --Author xxx
```
目录结构如下：
```bash
tree user
user
├── README.md
├── mojo
│   └── user
└── package.mojo

3 directories, 2 files
```
### 2.在mojo/user目录下构思接口和数据结构

新增mojo/user/v1/user.mojo文件，内容如下
```bash
/// 用户服务
///
/// 用户相关的接口
interface Users {
    /// 新增用户
    @http.post("/v1/users")
    create_user(user: User @1) -> User

    /// 获取用户详情
    @http.get("/v1/users/{id}")
    get_user(id: String @1) -> User

    /// 更新用户信息
    @http.put("/v1/users/{id}")
    update_user(user: User @1) -> User

    /// 删除用户
    @http.delete("/v1/users/{id}")
    delete_user(id: String @1)
}

type User {
    id: String @1
    name: String @2
    create_time: Timestamp @100
    update_time: Timestamp @101
}

注意：
1. interface定义接口信息，type定义数据接口
2. 字段名称小写且为下划线形式，字段类型大写
3. 接口url书写规则采用googleAIP的规范形式
```
### 3.使用mojo命令生成项目代码
在生成user目录下使用命令：
```bash
mojo build --targets=api,service
```
--targets=api参数命令构建api接口，可生成openapi、document、protobuf和go目录文件

--targets=service参数命令构建service-go目录，其中包含http服务相关的代码

构建后目录结构如下：
```bash
tree
.
├── README.md
├── document
│   └── user
│       └── v1
│           ├── user.md
│           └── users.md
├── go
│   ├── go.mod
│   ├── go.sum
│   └── pkg
│       └── user
│           └── v1
│               ├── users.pb.go
│               └── users_grpc.pb.go
├── mojo
│   └── user
│       └── v1
│           └── users.mojo
├── openapi
│   └── user
│       └── v1
│           ├── user.schema.json
│           └── users.yaml
├── package.mojo
├── protobuf
│   └── user
│       └── v1
│           └── users.proto
└── service-go
    ├── README.md
    ├── build
    │   └── users-server-docker
    │       └── Docerfile
    ├── cmd
    │   ├── user-unified-server
    │   │   └── main.go
    │   └── users-server
    │       └── main.go
    ├── configs
    │   ├── logs.yaml
    │   ├── metrics.yaml
    │   ├── sd.yaml
    │   ├── server.yaml
    │   └── tracing.yaml
    ├── go.mod
    ├── go.sum
    ├── internal
    │   ├── server
    │   │   └── run.go
    │   └── users-server
    │       └── run.go
    ├── pkg
    │   └── users-service
    │       ├── handlers
    │       │   ├── handlers.go
    │       │   ├── hooks.go
    │       │   └── middlewares.go
    │       └── svc
    │           ├── endpoints.go
    │           ├── transport_grpc.go
    │           └── transport_http.go
    └── scripts
        └── run.sh

32 directories, 32 files
```

