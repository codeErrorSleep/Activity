# create table, init data
goose -dir storage/mysql/migrations mysql "root:root@tcp(localhost:3306)/activity?parseTime=true" up