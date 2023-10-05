DB_URL=postgres://postgres:12345@localhost:5435/blog?sslmode=disable

mig-add:
	migrate create -ext sql -dir db/migrations -seq $(name)

mig-up:
	migrate -path db/migrations -database $(DB_URL) up

mig-down:
	migrate -path db/migrations -database $(DB_URL) down

mig-up1:
	migrate -path db/migrations -database $(DB_URL) up 1

mig-down1:
	migrate -path db/migrations -database $(DB_URL) down 1