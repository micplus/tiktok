# tiktok

补上字节青训营项目摸的鱼_(:з)∠)_

所有接口均正确运行

Client <-> 

## Problems

problems when coding:

1. when to init():

在init函数中，os.Getenv("KEY")取不到需要的变量
 
2. 函数外的变量在包间传递的问题

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
