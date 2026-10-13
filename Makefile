MIGRATION_FOLDER=db/migrations
DB_URL=root:1209@tcp(127.0.0.1:3306)/auth_dev

# Create a new migration
# Usage: gmake migrate-create name="create_session_table"
migrate-create:
	goose -dir $(MIGRATIONS_FOLDER) create $(name) sql

# Run all pending migrations
# Usage: gmake migrate-up
migrate-up:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" up

# Rollback the last migration
# Usage: gmake migrate-down
migrate-down:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" down

# Rollback all migrations and reset database
# Usage: gmake migrate-reset
migrate-reset:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" reset

# Show current migration status
# Usage: gmake migrate-status
migrate-status:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" status

# Redo the last migration (down then up)
# Usage: gmake migrate-redo
migrate-redo:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" redo

# Run migrations up to a specific version
# Usage: gmake migrate-to version=20250713105734
migrate-to:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" up-to $(version)

# Rollback to a specific migration version
# Usage: gmake migrate-down-to version=20250713105734
migrate-down-to:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" down-to $(version)

# Force a specific migration version
# Usage: gmake migrate-force version=20250713105734
migrate-force:
	goose -dir $(MIGRATIONS_FOLDER) mysql "$(DB_URL)" force $(version)