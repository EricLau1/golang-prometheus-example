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