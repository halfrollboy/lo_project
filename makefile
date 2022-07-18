dev:
	go run cmd/http-server/main.go
migrations:
	docker run -v /mnt/d/Projects/Golangs/wb-project-l0/http-server/pkg/postgres/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable up
migrations2:
	migrate -path pkg/postgres/migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" up