postgres:
	docker run --name postgress -p 5432:5432 -e POSTGRES_USER=root -e  POSTGRES_PASSWORD=secret -d postgres
createdb:
	docker exec -it postgress createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it postgress dropdb --username=root --owner=root simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

test:
	go test -timeout 30000000ms -v -cover ./...

sqlc:
	sqlc generate

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/ArsalanKm/simple_bank/db/sqlc Store
	
.PHONY: postgres createdb dropdb migratedown migrateup sqlc server mock