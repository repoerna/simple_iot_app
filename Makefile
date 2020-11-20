postgres:
	docker run --name dev-postgres --network="host" -e POSTGRES_PASSWORD=postgres -d -p 54320:5432 -v pgdata:/var/lib/postgresql/data postgres

redis:
	docker run -d --name dev-redis --network="host" -p 6379:6379 redis

createdb:
	docker exec -it dev-postgres createdb --owner=postgres --username=postgres  simple_iot_app

dropdb:
	docker exec -it dev-postgres dropdb --username=postgres  simple_iot_app

createmigrate:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_iot_app?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_iot_app?sslmode=disable" -verbose down

migrateforce:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_iot_app?sslmode=disable" force 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown migrateforce sqlc test server mock
