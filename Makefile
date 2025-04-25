# Variáveis
APP_NAME=hmacgen
CMD_DIR=cmd/hmacgen

# Compila o CLI
build:
	@echo "🔨 Compilando o CLI $(APP_NAME)..."
	@go build -o $(APP_NAME) ./$(CMD_DIR)
	@echo "✅ Compilado: ./$(APP_NAME)"

# Executa todos os testes
test:
	@echo "🧪 Executando testes..."
	@go test ./... -v
	@echo "✅ Testes finalizados."

# Executa o hmacgen CLI com argumentos customizados
run-hmacgen:
	@echo "🚀 Executando $(APP_NAME) com args: $(ARGS)"
	@go run ./$(CMD_DIR) $(ARGS)

# Apaga o binário gerado
clean:
	@echo "🧹 Limpando arquivos gerados..."
	@rm -f $(APP_NAME)
	@echo "✅ Clean finalizado."

# Instala o binário no GOPATH/bin ou GOBIN
install:
	@echo "📦 Instalando o CLI $(APP_NAME)..."
	@go install ./$(CMD_DIR)
	@echo "✅ Instalado em $(shell go env GOPATH)/bin ou $(shell go env GOBIN)"

# Formata o código
fmt:
	@echo "🎨 Formatando código..."
	@go fmt ./...
	@echo "✅ Código formatado."
