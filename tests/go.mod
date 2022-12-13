module github.com/go-leo/sqlgen/tests

go 1.16

require (
	golang.org/x/sys v0.1.0 // indirect
	gorm.io/driver/mysql v1.4.3
	gorm.io/driver/sqlite v1.4.2
	github.com/go-leo/sqlgen v0.3.16
	gorm.io/gorm v1.24.0
	gorm.io/hints v1.1.1 // indirect
	gorm.io/plugin/dbresolver v1.3.0
)

replace github.com/go-leo/sqlgen => ../
