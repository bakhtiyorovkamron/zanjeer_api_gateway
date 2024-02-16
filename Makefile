make swag-init:
	swag init -g api/router.go  -o api/docs

make migrate:
	migrate -database "postgres://postgres:password@mypostgres:5432/example?sslmode=disable" -path ./migrations up

make migrate-down:
	migrate -database "postgres://postgres:password@mypostgres:5432/example?sslmode=disable" -path ./migrations down
