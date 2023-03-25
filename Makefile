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
	go test -v -cover ./...

sqlc:
	sqlc generate
.PHONY: postgres createdb dropdb migratedown migrateup sqlc