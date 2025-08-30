.PHONY: dev-up dev-down dev-logs sqlc migrate-up migrate-down

dev-up:
	docker compose up -d --build

dev-down:
	docker compose down -v

dev-logs:
	docker compose logs -f --tail=100

sqlc:
	@echo "(opcional) ejecuta sqlc generate si lo configuras"

migrate-up:
	@echo "(opcional) usa golang-migrate para subir migraciones"

migrate-down:
	@echo "(opcional) usa golang-migrate para bajar migraciones"

