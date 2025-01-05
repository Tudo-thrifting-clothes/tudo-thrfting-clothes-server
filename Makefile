# app name
APP_NAME = server

run:
	go run ./cmd/${APP_NAME}/

up: 
	docker-compose up -d