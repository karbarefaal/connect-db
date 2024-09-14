include .env

export $(shell sed 's/=.*//' .env)

run:
	go run main.go

create_db:
	docker run --name postgres-container1 -d  -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) -e POSTGRES_USER=$(POSTGRES_USER) -p $(POSTGRES_PORT):5432 postgres:14
