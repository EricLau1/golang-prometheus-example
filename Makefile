run:
	@go run main.go

ps:
	@docker images
	@docker ps

up:
	@docker-compose up -d

down:
	@docker-compose down

clear:
	- @docker rm -f grafana
	- @docker rm -f prometheus
	- @docker rm -f golang-prometheus-example
	- @docker volume prune
	- @docker rmi -f golang-prometheus-example

clients:
	@go run cmd/clients/main.go

jobs:
	@go run cmd/jobs/main.go