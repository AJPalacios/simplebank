postgres:
	docker run --name postgres-db-master-class -e POSTGRES_USER=myuser -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres
createdb:
	docker exec -it postgres-db-master-class createdb --username=myuser --owner=myuser simple_bank
dropdb:
	docker exec -it postgres-db-master-class dropdb --username=myuser simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://myuser:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://myuser:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server