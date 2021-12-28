本项目的目录很大程度地参考了基于
[GoFrame](https://github.com/gogf/gf)
框架的
[Focus](https://github.com/gogf/focus#%E4%BA%94%E7%9B%AE%E5%BD%95%E8%AF%B4%E6%98%8E)
项目。

一、目录结构
======
```
├── app
│   ├── dao
│   ├── model
│   ├── shared
│   └── system
│       ├── admin
│       │   └── internal
│       └── index
│           └── internal
│               ├── controller
│               ├── define
│               └── service
├── cmd
│   ├── admin
│   │   └── main.go
│   ├── index
│   │   └── main.go
│   └── gen
│       └── main.go
├── config
│   ├── dev.toml
│   ├── prod.toml
│   └── init.go
├── pkg
│   ├── kafka
│   ├── elasticsearch
│   └── ...
├── test
├── util
└── go.mod
```
二、目录说明
======

|目录/文件名称   | 说明 | 描述
|---|---|---
|`app`           | 应用程序 | 存放所有的应用程序文件的目录
| &nbsp; &nbsp; -`dao`        | 数据访问   | 数据库的访问操作，仅包含最基础的数据库CURD方法
| &nbsp; &nbsp; -`model`      | 数据模型   | 存放系统数据模型/实体结构定义
| &nbsp; &nbsp; -`system`     | 系统模块   | 可包含多个子系统，不同子系统之间资源相互隔离
| &nbsp; &nbsp; &nbsp; &nbsp; -`index`    | 前台子系统 |
| &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; -`internal`       | 内部模块 | 子系统内部模块，仅供子系统内部调用，防止子系统外部调用
| &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; -`controller`     | 请求入口 | 子系统接外部请求的入口/接口层
| &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; -`define`  | 结构定义 | 子系统的输入、输出数据结构定义
| &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; -`service` | 业务逻辑 | 子系统业务逻辑封装，实现特定业务需求
| &nbsp; &nbsp; &nbsp; &nbsp; -`admin`    | 后台子系统 |
|`cmd`           | 程序入口   | 存放启动子系统程序或辅助程序的入口文件`main.go`
|`config`        | 配置管理   | 存放配置文件
|`pkg`        | 公共模块   |
|`go.mod`        | 依赖管理   | 存放包管理的描述文件

三、开发步骤
======
以 account model 和 index system 为例进行说明
1. 定义业务实体。在app/model目录下新建account.go文件，并设计account的结构体
2. 生成对应的dao，service 和 controller文件。根据需要修改cmd/gen/main.go中的 system, projectName 和 entity，entity的修改即导入步骤1中定义的account的结构体；运行main.go
3. 注册api路由。在app/system/index/router.go中添加`web.Router("/api/account", controller.Account)`
4. 业务开发。编写`app/system/index/internal/controller/account.go`和`app/system/index/internal/service/account.go`

四、实体定义
======
```go
type Account struct {
	internal.Base
	Title  string `orm:"size(32)"`
	Type
}
```

# 五、dao层的使用
1. 增
```go
func (s *accountService) Create() error {
	data := &struct {
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}{
		Email:    "example@qq.com",
		Phone:    "13282113456",
		Password: "123456",
	}
	// 或者
	data = map[string]interface{}{
		"email":"example@qq.com",
		"phone":    "13282113456",
		"password": "123456",
    }
	if _, err := dao.Account.SetData(data).Insert(); err != nil {
		return err
	}
	return nil

	// INSERT INTO `account` 
	// (`deleted`, `create_time`, `update_time`, `email`, `phone`, `password`) 
	// VALUES (
	//    `false`, 
	//    `2021-12-27 22:54:25.911764 +0800 CST`, 
	//    `2021-12-27 22:54:25.911766 +0800 CST`, 
	//    `example@qq.com`, `13282113456`, `123456`,
	// )
}
```

2. 删
```go
// DELETE FROM account WHERE id = 1;
func (s *accountService) Delete() error {
	if _, err := dao.Account.Filter(dao.Account.Fields.Id.Exact(), 1).Delete(); err != nil {
		return err
	}
	return nil
}
```
3. 改
```go
func (s *accountService) Create() error {
	data := &struct {
		Password string `json:"password"`
	}{
		Password: "654321",
	}
	// 或者
	data = map[string]interface{}{
		"password": "654321",
    }
	
	// email = 'example@qq.com' OR phone = '13282113456'
	cond := dao.NewCond().
		Or(dao.Account.Fields.Email.Exact(), "example@qq.com").
		Or(dao.Account.Fields.Phone.Exact(), "13282113456")
	
	// deleted = true
	cond2 := dao.NewCond().
		And(dao.Account.Fields.Deleted.Exact(), false)
	
	// deleted = true AND ( email = 'example@qq.com' OR phone = '13282113456' )
	cond3 := cond2.AndCond(cond)

    // UPDATE `account` T0 
	// SET T0.`password` = `654321`, 
	//     T0.`update_time` = `2021-12-27 22:58:19.098462 +0800 CST m=+247.673773035` 
	// WHERE T0.`deleted` = `false` AND ( T0.`email` = `example@qq.com` OR T0.`phone` = `13282113456` )
	if _, err := dao.Account.SetCond(cond3).SetData(data).Update(); err != nil {
		return err
	}
	return nil
}
```
4. 查
```go
func (s *accountService) Get() (interface{}, error) {
	var cont  []struct {
		Phone string `json:"phone"`
		Password string `json:"password"`
		CreateTime string `json:"create_time"`
	}

	// SELECT T0.`phone`, T0.`password`, T0.`create_time` 
	// FROM `account` T0 
	// WHERE T0.`phone` LIKE `%132%`
	if _, err := dao.Account.Filter(dao.Account.Fields.Phone.Contains(), "132").Rows(&cont); err != nil {
		return nil, err
	}
	return cont, nil
}
```
联表
```go
func (s *readerService) Get() (interface{}, error) {
	var cont []struct{
		Name string `json:"name"`
		Gender string `json:"gender"`
		Email string `json:"email" alias:"t1"`
		AccountPassword string `json:"account_password" alias:"t1" field:"password"`
	}
	
	// SELECT T0.name, T0.gender, 
	//        T1.email, T1.password account_password 
	// FROM reader T0 LEFT JOIN account T1 ON T0.account_id = T1.id 
	// WHERE T0.id in (`1`,`2`,`3`) AND T1.email > ``  LIMIT 10
	readerIds := []int64{1,2,3}
	dao.Reader.
		LeftJoin(dao.Account, dao.AliasT1).On(dao.Reader.Fields.AccountId).
		Where(dao.Reader.Fields.Id.Alias().In(len(readerIds)), readerIds).
		And(dao.Account.Fields.Email.AliasT1().Greater(), "").
		Limit(10).
		Rows(&cont)
	return cont, nil
}

```