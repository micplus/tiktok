# tiktok

补上字节青训营项目摸的鱼_(:з)∠)_

所有接口均正确运行

## 存储

使用MySQL存储视频、用户、评论内容表

用户点赞视频、关注博主是频繁发生的，每次+1开销大，故用Redis存储用户-视频点赞关系、用户关注关系、当前在线ID

## 架构

客户端--前端服务器（对客户端提供HTTP服务，发起RPC）--后端服务器（对前端服务器提供RPC服务）--内存+外存

二者分离从而使客户端不可见到实际的服务调用

后端服务器内部按照Controllers/Services的形式组织

TODO：让Service单独运行，Controller转化为API网关，增加服务注册/发现的守护进程，实现微服务架构

## Problems

problems when coding:

1. when to init():

在init函数中，os.Getenv("KEY")取不到需要的变量
 
2. 函数外的变量初始化时机的问题

见gopl 2.3变量一节：包级变量在main函数入口前完成初始化；局部变量在实际执行到的位置初始化。

```go
package database

// import ...

var DB *sqlx.DB

func Setup() {
  var err error
  DB, err = sqlx.Connect("mysql", config.DSN)
  // if err
}
```

```go
package some

import "database"

var db = database.DB

func SomeMapper() {
  // ...
  rows, err := db.Exec(query, args...)  // Segment Fault，说明这里的db是空的
  // ...
}
```

3. rpc

rpc.Register作用在DefaultServer上，创建新Server则在创建的Server上注册

```go
srv := rpc.NewServer()
srv.Register(new(some.Func))
```

4. rpc RegisterName

观察函数返回的error信息，可知RegisterName(string, any)接收的服务名称是唯一的，且string域指定的是前缀名，具体的服务名称取决于传入类型上的方法

5. 路由坑

postman对/feed和/feed/发请求效果是一样的，但路由设为"/feed",在浏览器、客户端中发/feed/请求就会307
