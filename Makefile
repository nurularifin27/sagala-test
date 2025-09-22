# Variables
# This variable is optional, but it's good practice to define the migration tool and directory.
MIGRATE_TOOL = migrate
MIGRATION_DIR = migrations

# Help target to display available commands
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make build             - Compile the Go application"
	@echo "  make run               - Run the compiled application"
	@echo "  make clean             - Remove the executable"
	@echo "  make migration/create  - Create a new database migration"
	@echo "  make migration/up      - Apply all new migrations"
	@echo "                           Usage: make migration/up dsn=\"sqlite3://./data.db\""
	@echo "  make migration/down    - Revert the last applied migration"
	@echo "                           Usage: make migration/down dsn=\"sqlite3://./data.db\""
	@echo "  make migration/force   - Force set a database version (useful for fixing dirty states)"
	@echo "                           Usage: make migration/force dsn=\"sqlite3://./data.db\" version=N"
	@echo "  make migration/version - Show the current database version"
	@echo "                           Usage: make migration/version dsn=\"sqlite3://./data.db\""

# Application targets
.PHONY: build
build:
	go build -o sagala main.go

.PHONY: run
run: build
	./sagala

.PHONY: clean
clean:
	rm sagala

# Database Migration targets using 'go-migrate/migrate'
# To create a new migration file, run:
# make migration/create name="add_users_table"
.PHONY: migration/create
migration/create:
ifndef name
	$(error "name is required. Usage: make migration/create name=\"your_migration_name\"")
endif
	@echo "Creating new migration files in $(MIGRATION_DIR)/..."
	$(MIGRATE_TOOL) create -ext sql -dir $(MIGRATION_DIR) $(name)

# To apply all new migrations, run:
# make migration/up dsn="your_database_dsn"
.PHONY: migration/up
migration/up:
ifndef dsn
	$(error "dsn is required. Usage: make migration/up dsn=\"your_database_dsn\"")
endif
	@echo "Applying migrations..."
	$(MIGRATE_TOOL) -database "$(dsn)" -path $(MIGRATION_DIR) up

# To revert the last applied migration, run:
# make migration/down dsn="your_database_dsn"
.PHONY: migration/down
migration/down:
ifndef dsn
	$(error "dsn is required. Usage: make migration/down dsn=\"your_database_dsn\"")
endif
	@echo "Reverting last migration..."
	$(MIGRATE_TOOL) -database "$(dsn)" -path $(MIGRATION_DIR) down

# To force set a database version, run:
# make migration/force dsn="your_database_dsn" version=N
.PHONY: migration/force
migration/force:
ifndef dsn
	$(error "dsn is required. Usage: make migration/force dsn=\"your_database_dsn\" version=N")
endif
ifndef version
	$(error "version is required. Usage: make migration/force dsn=\"your_database_dsn\" version=N")
endif
	@echo "Forcing database version to $(version)..."
	$(MIGRATE_TOOL) -database "$(dsn)" -path $(MIGRATION_DIR) force $(version)

# To show the current database version, run:
# make migration/version dsn="your_database_dsn"
.PHONY: migration/version
migration/version:
ifndef dsn
	$(error "dsn is required. Usage: make migration/version dsn=\"your_database_dsn\"")
endif
	@echo "Checking database version..."
	$(MIGRATE_TOOL) -database "$(dsn)" -path $(MIGRATION_DIR) version
