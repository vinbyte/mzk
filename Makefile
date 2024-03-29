include .env

DB_HOST=$(DATABASE.WRITE.HOST)
# due to we run a migration via docker container, we need to change the localhost to host.docker.internal to connect to database in local
# ref : https://docs.docker.com/desktop/networking/#i-want-to-connect-from-a-container-to-a-service-on-the-host
ifeq ($(DB_HOST), localhost)
   DB_HOST = host.docker.internal
endif
DB_URL=postgres://$(DATABASE.WRITE.USER):$(DATABASE.WRITE.PASSWORD)@$(DB_HOST):$(DATABASE.WRITE.PORT)/$(DATABASE.WRITE.NAME)?sslmode=$(DATABASE.WRITE.SSLMODE)
MIGRATION_NAME ?= $(shell bash -c 'read -p "Migration name: " name; echo $$name')

migrate-new:
	@docker run --rm -it --network=host -v ./migrations:/db/migrations ghcr.io/amacneil/dbmate new $(MIGRATION_NAME) 

migrate-up:
	@docker run --rm -it --network=host -v ./migrations:/db/migrations ghcr.io/amacneil/dbmate -u $(DB_URL) up

migrate-down:
	@docker run --rm -it --network=host -v ./migrations:/db/migrations ghcr.io/amacneil/dbmate -u $(DB_URL) down

generate:
	go generate ./...

dev: generate
	go run .

compose-up:
	@docker-compose up

compose-down:
	@docker-compose down