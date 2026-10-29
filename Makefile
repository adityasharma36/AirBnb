
MIGRATIONS_FOLDER=db/migrations
DB_URL=root:1209@tcp(127.0.0.1:3306)/auth_dev

migrate-create:
	goose -dir $(MIGRATIONS_FOLDER) create $(name) sql

migrate-up:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" up

migrate-down:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" down

migrate-status:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" status

migrate-reset:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" reset
