# autocode 
### 一个自动生成后端代码的工具，小巧精致，功能齐全
#### 基于gin,sqlx等工具

#### 生成代码使用示例
```go
  autocode.InitDB("root", "123456", "localhost", "3306", "demo")
	b := autocode.ProBasic{ModName: "testAuto", TblName: "tbl_product", Name: "Product"}
	b.Start()
```

#### 项目启动示例
```go
	db.Init("root", "123456", "localhost", "3306", "demo")
	r := gin.Default()
	router.Register(r)
	r.Run(":8080")
```
> 注意：需要将router包下生成的router注册到，register.go 文件中，该模块方可使用
