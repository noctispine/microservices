devup:
	docker compose -f dev.docker-compose.yml up -d
devdown:
	docker compose -f dev.docker-compose.yml down
produp:
	docker compose -f prod.docker-compose.yml up -d
proddown:
	docker compose -f prod.docker-compose.yml down
rmdata:
	rm -r ./data/
run:
	nodemon --exec APP_ENV=DEV go run cmd/main.go --signal SIGTERM || exit 1
