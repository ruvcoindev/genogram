.PHONY: help migrate-up migrate-down start stop clean status

help:
	@echo "📋 Доступные команды:"
	@echo "  make migrate-up    - Применить миграции"
	@echo "  make migrate-down  - Откатить миграции"
	@echo "  make start         - Запустить все сервисы"
	@echo "  make stop          - Остановить все сервисы"
	@echo "  make clean         - Очистить данные"
	@echo "  make status        - Статус сервисов"

migrate-up:
	@echo "📦 Применение миграций..."
	@chmod +x scripts/migrate.sh
	@./scripts/migrate.sh

migrate-down:
	@echo "⏪ Откат последней миграции..."
	@PGPASSWORD=$(POSTGRES_PASSWORD) psql -h $(POSTGRES_HOST) -U $(POSTGRES_USER) -d $(POSTGRES_DB) -f scripts/rollback.sql

start:
	@echo "🚀 Запуск системы..."
	@docker-compose up -d
	@echo "✅ Система запущена:"
	@echo "   Frontend: http://localhost"
	@echo "   Backend API: http://localhost:8080"
	@echo "   Neo4j Browser: http://localhost:7474"

stop:
	@echo "⏹️ Остановка системы..."
	@docker-compose down

clean:
	@echo "🧹 Очистка данных..."
	@docker-compose down -v
	@rm -rf ./data/*
	@echo "✅ Очищено"

status:
	@docker-compose ps
