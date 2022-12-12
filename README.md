# sqlgen
通过sql 生成 数据裤操作接口


# 安装
go install github.com/go-leo/sqlgen@latest

# 使用
gentool -h

Usage of gentool:
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

参考
- https://github.com/go-gorm/gen
- https://github.com/anqiansong/sqlgen
- https://goframe.org/pages/viewpage.action?pageId=7296196