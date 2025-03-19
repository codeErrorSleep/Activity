# create table, init data
goose -dir db/migrations mysql "root:root@tcp(localhost:3306)/db_q_goods_center?parseTime=true" up