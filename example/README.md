# Mojo
## 编译chaosmojo二进制命令
```bash
../cmd/mojo/build.sh
```
## 实践指南
### 1.指定package生成代码
在任一目录下使用如下命令生成user服务：
```bash
chaosmojo create --package=user --output=. --repo=github.com/liankui/mojo/example/user --Author=cc
```
目录结构如下：
```bash
tree user             
```
```
user
├── README.md
├── mojo
│     └── user
└── package.mojo
```

### 2.在mojo/demo目录下构思接口和数据结构
新增路径mojo/user/v1，并添加mojo/user/v1/user.mojo文件，内容如下
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
```
注意：
1. 使用interface定义接口信息，使用type定义数据接口
2. 字段名称小写且为下划线形式，字段类型大写
3. 接口url书写规则采用[Google AIP](https://google.aip.dev/122)规范形式

### 3.使用mojo命令生成项目代码
在生成user目录下使用命令：
```bash
chaosmojo build --path=user --targets=api,service
```
--targets=api参数命令构建api接口，可生成openapi、document、protobuf和go目录文件

--targets=service参数命令构建service-go目录，其中包含http服务相关的代码

构建后目录结构如下：
```bash
tree user -L 2
```

``` 
user
├── README.md
├── document
│     └── user
├── go
│     ├── go.mod
│     ├── go.sum
│     └── pkg
├── mojo
│     └── user
├── openapi
│     └── user
├── package.mojo
├── protobuf
│     └── user
└── service-go
├── README.md
├── build
├── cmd
├── configs
├── go.mod
├── internal
├── pkg
└── scripts
```
### 4.完善go.mod引用，开始编写业务代码
DONE

