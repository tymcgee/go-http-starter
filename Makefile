build: generate
	go build

.PHONY: run
run: generate
	go run .

.PHONY: create-migration
# invoke with "make create-migration name=name-of-migration"
create-migration:
	migrate create -ext sql -dir dao/migrations -seq $(name)

.PHONY: apply-migration
apply-migration:
	migrate -database "sqlite://data.db" -path ./dao/migrations up

.PHONY: generate
generate:
	sqlc generate

.PHONY: install
install:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
# add whatever database drivers you want here
	go install -tags 'sqlite,mysql,postgres,pgx5' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
