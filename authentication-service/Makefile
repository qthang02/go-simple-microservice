createdb:
	docker exec -it postgres createdb --username=root --owner=root mico-go
dropdb:
	docker exec -it postgres dropdb mico-go
new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/micro-go?sslmode=disable" --verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/micro-go?sslmode=disable" --verbose down
server:
	go run main.go
.PHONY: createdb dropdb new_migration migrateup migratedown server