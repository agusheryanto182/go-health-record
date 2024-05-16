include .env
export

.phony: build migrate rollback

build:
	GOOS=linux GOARCH=amd64 go build -o main_AgusHeryanto182 main.go

migrate:
	migrate -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -path db/migrations -verbose up

rollback:
	migrate -database "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -path db/migrations -verbose drop