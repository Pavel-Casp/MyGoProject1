# Определяем переменные
SWAGGER_CMD=swag
SWAGGER_INIT_CMD=$(SWAGGER_CMD) init -g cmd/api_server/main.go
GO=go
BINARY_NAME=mygoproject
PORT=8080

.PHONY: help
help: ## Показать помощь
	@echo "Использование:"
	@echo "  make <цель>"
	@echo ""
	@echo "Цели:"
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: swagger
swagger: ## Генерировать Swagger документацию
	@echo "Генерация Swagger документации..."
	@$(SWAGGER_INIT_CMD)
	@echo "Документация сгенерирована в папке docs"

.PHONY: build
build: swagger ## Собрать проект (включая генерацию документации)
	@echo "Сборка проекта..."
	@$(GO) build -o $(BINARY_NAME) ./cmd/api_server

.PHONY: run
run: swagger ## Запустить сервер (с генерацией документации)
	@echo "Запуск сервера на порту $(PORT)..."
	@$(GO) run ./cmd/api_server

.PHONY: clean
clean: ## Очистить сгенерированные файлы
	@echo "Очистка..."
	@rm -f $(BINARY_NAME)
	@rm -rf docs

.PHONY: test
test: ## Запустить тесты
	@echo "Запуск тестов..."
	@$(GO) test ./...

.PHONY: install-deps
install-deps: ## Установить зависимости
	@echo "Установка зависимостей..."
	@go install github.com/swaggo/swag/cmd/swag@latest
	@$(GO) mod download