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
sqlcmd:
	docker run --rm -v "%cd%:/src" -w /src kjconroy/sqlc generate
mock:
	mockgen -package mockdb -destination db/mock/store.go authentication-service/db/sqlc Store
proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
        --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
        proto/*.proto
test:
	go test -v -cover ./...
.PHONY: createdb dropdb new_migration migrateup migratedown server sqlcmd mock