createdb:
	createdb -U xianfengyuan -O xianfengyuan simple_bank

dropdb:
	dropdb -U xianfengyuan simple_bank

migrateup:
	migrate -path db/migration -database "postgres://xianfengyuan@localhost/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://xianfengyuan@localhost/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

.PHONY: createdb dropdb migrateup migratedown sqlc
