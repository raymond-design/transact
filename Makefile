postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13-alpine

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root transact_service

dropdb:
	docker exec -it postgres13 dropdb transact_service

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/transact_service?sslmode=disable" --verbose up

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/transact_service?sslmode=disable" --verbose down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
