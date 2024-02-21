.SILENT:
run:
	docker compose up --build

migrate:
	goose -dir migrations postgres "postgres://postgres:postgres@localhost:5432/zip?sslmode=disable" up