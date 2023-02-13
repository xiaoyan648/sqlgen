# sqlgen
根据数据库表信息生成结构体和基础操作方法，是 `gorm.io/gen` 包的增强包；
在 `gorm.io/gen` 的基础上添加了如下功能

- [x] 代码生成命令行工具 支持动态语句
- [x] 支持外部事务


# 安装
go install github.com/go-leo/sqlgen@latest

# 使用

## 案例
参考 `./example`


## 更多
```
sqlgen -h

Usage of sqlgen:
 -db string
       input mysql or postgres or sqlite or sqlserver. consult[https://gorm.io/docs/connecting_to_the_database.html] (default "mysql")
 -dsn string
       consult[https://gorm.io/docs/connecting_to_the_database.html]
 -fieldNullable
       generate with pointer when field is nullable
 -fieldWithIndexTag
       generate field with gorm index tag
 -fieldWithTypeTag
       generate field with gorm column type tag
 -modelPkgName string
       generated model code's package name
 -outFile string
       query code file name, default: gen.go
 -outPath string
       specify a directory for output (default "./dao/query")
 -tables string
       enter the required data table or leave it blank
 -onlyModel
       only generate models (without query file)
 -withUnitTest
       generate unit test for query code
 -fieldSignable
       detect integer field's unsigned type, adjust generated data type
```

# 参考
- https://github.com/go-gorm/gen
- https://github.com/anqiansong/sqlgen
- https://goframe.org/pages/viewpage.action?pageId=7296196
  
# TODO

- json tag 支持开关控制
- 支持模型分表逻辑