postgres:
createdb:
	docker run --name postgres-db-master-class -e POSTGRES_USER=myuser -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres
	docker exec -it postgres-db-master-class createdb --username=myuser --owner=myuser simple_bank
dropdb:
	docker exec -it postgres-db-master-class dropdb --username=myuser simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://myuser:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migration -database "postgresql://myuser:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://myuser:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://myuser:mysecretpassword@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/devspace/simplebank/db/sqlc Store
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock migrateup1 migratedown1