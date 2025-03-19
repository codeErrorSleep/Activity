# init project data
GOOSE_CMD = goose -dir db/migrations mysql "root:root@tcp(localhost:3306)/db_q_goods_center?parseTime=true"

.PHONY: migrate-up
migrate-up:
	$(GOOSE_CMD) up

.PHONY: migrate-down
migrate-down:
	$(GOOSE_CMD) down

.PHONY: migrate-status
migrate-status:
	$(GOOSE_CMD) status