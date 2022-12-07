up:
	docker compose up -d
down:
	docker compose down
migrate:
	rm -r ./data/ & docker compose up -d
run:
	nodemon --exec APP_ENV=DEV go run cmd/main.go --signal SIGTERM || exit 1
