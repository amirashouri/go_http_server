network:
	docker network create go-http-server-network

createdb:
	docker exec -it postgres createdb --username=root --owner=root go_http_server

postgres:
	docker run --name postgres --network go-http-server-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

dropdb:
	docker exec -it postgres dropdb go_http_server

sqlc:
	sqlc generate

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

server:
	go run main.go

.PHONY: createdb postgres dropdb sqlc server new_migration