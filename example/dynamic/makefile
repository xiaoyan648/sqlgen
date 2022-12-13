.PHONY: data
data:
	@sqlgen -dsn 'jcptydlm:jcptydlm8@tcp(rm-uf6q11rs8xp84t5dcdo.mysql.rds.aliyuncs.com:3306)/new_media_content?charset=utf8mb4&parseTime=true&loc=Local&timeout=500ms&readTimeout=500ms&writeTimeout=500ms' \
	 -outPath "./query" \
	 -outFile "query.gen.go" \
	 -tables "new_media_book,category" \
	 -fieldSignable \
	 -modelPkgName "./entity" \
	 -dynamicSQLFile "./querys.go"  