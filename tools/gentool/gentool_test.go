package main

import (
	"testing"
)

func Test(t *testing.T) {
	// db, err := connectDB(DBType("mysql"), "jcptydlm:jcptydlm8@tcp(rm-uf6q11rs8xp84t5dcdo.mysql.rds.aliyuncs.com:3306)/new_media_content?charset=utf8mb4&parseTime=true&loc=Local&timeout=500ms&readTimeout=500ms&writeTimeout=500ms")
	// if err != nil {
	// 	log.Fatalln("connect db server fail:", err)
	// }

	// g := gen.NewGenerator(gen.Config{
	// 	OutPath:      "./query",
	// 	ModelPkgPath: "./entity",
	// })

	// // set tinyint type
	// dataMap := map[string]func(detailType string) (dataType string){
	// 	"tinyint": func(detailType string) (dataType string) {

	// 		return "int8"
	// 	},
	// }
	// g.WithDataTypeMap(dataMap)

	// g.UseDB(db)

	// models, err := genModels(g, db, []string{"new_media_book"})
	// if err != nil {
	// 	log.Fatalln("get tables info fail:", err)
	// }

	// // if !config.OnlyModel {
	// g.ApplyBasic(models...)
	// // }
	// // g.ApplyInterface(func(dynamic.Querier) {}, models...)
	// g.ApplyInterface([]string{"../../example/dynamic", "Querier"}, models...)
	// g.Execute()
}
