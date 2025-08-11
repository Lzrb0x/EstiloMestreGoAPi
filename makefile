# Variáveis
APP_NAME=estiloMestreGO
BUILD_DIR=./bin
MAIN_PATH=./cmd/api/main.go
DOCKER_IMAGE=estilo-mestre-go

# Comandos padrão
.DEFAULT_GOAL := help

# Cores para output
RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[1;33m
BLUE=\033[0;34m
NC=\033[0m # No Color

## help: Mostra esta mensagem de ajuda
.PHONY: help
help:
	@echo "Comandos disponíveis:"
	@echo ""
	@grep -E '^##.*:' $(MAKEFILE_LIST) | sed 's/## //' | awk -F: '{printf "  $(BLUE)%-15s$(NC) %s\n", $$1, $$2}'

## build: Compila a aplicação
.PHONY: build
build:
	@echo "$(YELLOW)Compilando aplicação...$(NC)"
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PATH)
	@echo "$(GREEN)✓ Aplicação compilada em $(BUILD_DIR)/$(APP_NAME)$(NC)"

## run: Executa a aplicação
.PHONY: run
run:
	@echo "$(YELLOW)Executando aplicação...$(NC)"
	@go run $(MAIN_PATH)

## dev: Executa a aplicação em modo desenvolvimento com hot reload
.PHONY: dev
dev:
	@echo "$(YELLOW)Iniciando modo desenvolvimento...$(NC)"
	@if command -v air > /dev/null; then \
		air; \
	else \
		echo "$(RED)Air não instalado. Instalando...$(NC)"; \
		go install github.com/air-verse/air@latest; \
		echo "$(GREEN)✓ Air instalado com sucesso$(NC)"; \
		air; \
	fi

## test: Executa todos os testes
.PHONY: test
test:
	@echo "$(YELLOW)Executando testes...$(NC)"
	@go test -v ./...

## test-coverage: Executa testes com relatório de cobertura
.PHONY: test-coverage
test-coverage:
	@echo "$(YELLOW)Executando testes com cobertura...$(NC)"
	@go test -v -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "$(GREEN)✓ Relatório de cobertura gerado em coverage.html$(NC)"

## lint: Executa o linter
.PHONY: lint
lint:
	@echo "$(YELLOW)Executando linter...$(NC)"
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "$(RED)golangci-lint não instalado. Para instalar:$(NC)"; \
		echo "curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.2"; \
	fi

## format: Formata o código
.PHONY: format
format:
	@echo "$(YELLOW)Formatando código...$(NC)"
	@go fmt ./...
	@echo "$(GREEN)✓ Código formatado$(NC)"

## mod-tidy: Limpa as dependências
.PHONY: mod-tidy
mod-tidy:
	@echo "$(YELLOW)Limpando dependências...$(NC)"
	@go mod tidy
	@echo "$(GREEN)✓ Dependências limpas$(NC)"

## mod-vendor: Cria pasta vendor com dependências
.PHONY: mod-vendor
mod-vendor:
	@echo "$(YELLOW)Criando pasta vendor...$(NC)"
	@go mod vendor
	@echo "$(GREEN)✓ Pasta vendor criada$(NC)"

## swagger: Gera documentação Swagger
.PHONY: swagger
swagger:
	@echo "$(YELLOW)Gerando documentação Swagger...$(NC)"
	@if command -v swag > /dev/null; then \
		swag init -g $(MAIN_PATH); \
		echo "$(GREEN)✓ Documentação Swagger gerada$(NC)"; \
	else \
		echo "$(RED)Swag não instalado. Para instalar:$(NC)"; \
		echo "go install github.com/swaggo/swag/cmd/swag@latest"; \
	fi

## clean: Remove arquivos de build e temporários
.PHONY: clean
clean:
	@echo "$(YELLOW)Limpando arquivos...$(NC)"
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@rm -rf vendor/
	@echo "$(GREEN)✓ Arquivos limpos$(NC)"

## docker-build: Constrói imagem Docker
.PHONY: docker-build
docker-build:
	@echo "$(YELLOW)Construindo imagem Docker...$(NC)"
	@docker build -t $(DOCKER_IMAGE):latest .
	@echo "$(GREEN)✓ Imagem Docker construída$(NC)"

## docker-run: Executa container Docker
.PHONY: docker-run
docker-run:
	@echo "$(YELLOW)Executando container Docker...$(NC)"
	@docker run -p 8080:8080 $(DOCKER_IMAGE):latest

## docker-compose-up: Sobe serviços com docker-compose
.PHONY: docker-compose-up
docker-compose-up:
	@echo "$(YELLOW)Subindo serviços com docker-compose...$(NC)"
	@docker-compose up -d

## docker-compose-down: Para serviços do docker-compose
.PHONY: docker-compose-down
docker-compose-down:
	@echo "$(YELLOW)Parando serviços do docker-compose...$(NC)"
	@docker-compose down

## install-tools: Instala ferramentas de desenvolvimento
.PHONY: install-tools
install-tools:
	@echo "$(YELLOW)Instalando ferramentas de desenvolvimento...$(NC)"
	@go install github.com/air-verse/air@latest
	@go install github.com/swaggo/swag/cmd/swag@latest
	@echo "$(GREEN)✓ Ferramentas instaladas$(NC)"

## setup: Configura ambiente de desenvolvimento
.PHONY: setup
setup: install-tools mod-tidy swagger
	@echo "$(GREEN)✓ Ambiente configurado com sucesso!$(NC)"

## all: Executa formato, lint, test e build
.PHONY: all
all: format lint test build
	@echo "$(GREEN)✓ Pipeline completo executado$(NC)"
