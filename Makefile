postgres:
	docker run --name postgres16 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=admin123 -p 5432:5432 -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root red_pocket_game

dropdb:
	docker exec -it postgres16 dropdb --username=root --owner=root red_pocket_game

migrateup:
	migrate -path db/migration -database "postgresql://root:admin123@localhost:5432/red_pocket_game?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:admin123@localhost:5432/red_pocket_game?sslmode=disable" -verbose down

migratecreate:
	migrate create -ext sql -dir db/migration -seq $(name)

migrateforce:
	migrate -path db/migration -database "postgresql://root:admin123@localhost:5432/red_pocket_game?sslmode=disable" force $(version)

sqlc:
	sqlc generate

server:
	go run main.go

test:
	go test -cover ./db/sqlc/ ./handlers

mock:
	mockgen -package mockdb -destination db/mock/store.go Food_Shop_Server/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown migratecreate migrateforce sqlc server mock test

