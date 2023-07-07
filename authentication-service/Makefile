createdb:
	docker exec -it postgres createdb --username=root --owner=root micro_go
dropdb:
	docker exec -it postgres dropdb micro_go
new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)
migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/micro_go?sslmode=disable" --verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/micro_go?sslmode=disable" --verbose down
server:
	go run main.go
.PHONY: createdb dropdb new_migration migrateup migratedown server