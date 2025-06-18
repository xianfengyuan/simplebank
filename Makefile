DB_URL=postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable

createdb:
	createdb -U root -O root simple_bank

dropdb:
	dropdb -U root simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb dropdb migrateup migratedown sqlc test

