# Mojo

## 编译chaosmojo二进制命令

```bash
../cmd/mojo/build.sh
```

## 实践指南

### 1.指定package生成代码

在任一目录下使用如下命令生成user服务：

```bash
chaosmojo create --package=auth --output=. --repo=github.com/liankui/mojo/example/auth --version=v1.0 --license=MIT --Author=cc --email=cc@email.com --organization=cc.org
```

目录结构如下：

```bash
tree auth          
```

```
auth
├── README.md
├── mojo
│     └── auth
└── package.mojo
```

### 2.在mojo/user目录下构思接口和数据结构

添加文件，内容如下

```bash
cat <<EOF > ./auth/mojo/auth/user.mojo
/// user info
type User {
    id: String @1 //< id
    
    name: String @2 //< user name
    
    /// create time
    create_time: Timestamp @100
    
    /// update time
    update_time: Timestamp @101
}
EOF
```
添加文件，内容如下
```bash
mkdir -pv ./auth/mojo/auth/v1
cat <<EOF > ./auth/mojo/auth/v1/users.mojo
/// 用户服务
///
/// 用户相关的接口
@entity("Users")
interface Users {
    /// 新增用户 
    @http.post("auth/v1/users")
    create_user(user: User @1) -> User
    
    /// 获取用户详情
    @http.get("auth/v1/users/{id}")
    get_user(id: String @1) -> User
    
    /// 更新用户信息
    @http.put("auth/v1/users/{id}")
    update_user(user: User @1) -> User
    
    /// 删除用户
    @http.delete("auth/v1/users/{id}")
    delete_user(id: String @1)
}
EOF
```

注意：

1. interface定义接口信息，type定义数据结构
2. 字段的名称使用下划线小写，字段类型使用大驼峰命名
3. 接口url书写规则采用[Google AIP](https://google.aip.dev/122)规范形式

目录结构如下：

```bash
tree auth          
```

```
auth
├── README.md
├── mojo
│    └── auth
│         ├── user.mojo
│         └── v1
│             └── users.mojo
└── package.mojo

```

### 3.使用mojo命令生成项目代码

在生成user目录下使用命令：

```bash
chaosmojo build --path=./auth --targets=api,service
```

--targets=api参数命令构建api接口，可生成openapi、document、protobuf和go目录文件

--targets=service参数命令构建service-go目录，其中包含http服务相关的代码

构建后目录结构如下：

```bash
tree auth -L 2
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

### 5.generate again

```bash
rm -rf auth/document auth/go auth/openapi auth/protobuf auth/service-go
```

```bash
../cmd/mojo/build.sh
chaosmojo build --path=./auth --targets=api,service
```

```bash
chaosmojo build --path=./auth --targets=api,service
```
