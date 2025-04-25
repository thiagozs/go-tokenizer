# 🔐 Go Tokenizer API

API para **tokenização**, **destokenização** e **geração de HMAC temporizado**, desenvolvida em **Golang** com foco em segurança, autenticação por assinatura (HMAC) e proteção de dados sensíveis.

---

## ✨ Funcionalidades

- Gerar **tokens criptografados** a partir de valores sensíveis
- **Reverter tokens** para valores originais
- **Gerar HMACs** seguros e temporizados para proteger requisições
- **Validar requisições** autenticadas via HMAC
- Compatível com uso de HMAC baseado em **timestamp**, prevenindo **replay attacks**

---

## 🚀 Endpoints

| Método | Rota                  | Descrição |
|:------:|:----------------------:|:--------- |
| POST   | `/tokenize`             | Tokeniza um valor sensível |
| POST   | `/detokenize`           | Reverte um token para o valor original |
| POST   | `/genhmac`              | Gera um HMAC baseado em valor + timestamp |
| POST   | `/selftest-protected`   | Endpoint protegido que valida HMAC via Headers (modo middleware) |

---

## 🛠 Como rodar localmente

### Pré-requisitos

- Go 1.20 ou superior
- VS Code (opcional, mas recomendado)
- Extensão [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)

### Rodando a aplicação

```bash
# Clone o projeto
git clone https://github.com/thiagozs/go-tokenizer.git

# Acesse o diretório
cd go-tokenizer

# Instale dependências
go mod tidy

# Rode a aplicação
go run main.go
```

A aplicação estará disponível em:

```
http://localhost:8880
```

---

## 📦 Como testar via `.http` no VS Code

1. Instale a extensão **REST Client** (`humao.rest-client`)
2. Abra o arquivo `testes.http` fornecido
3. Execute as requisições individualmente:
   - Primeiro `POST /genhmac`
   - Depois `POST /selftest-protected`
   - Depois `POST /tokenize`
   - Depois `POST /detokenize`
4. Observe as respostas diretamente no painel lateral do VS Code

---

## 🧠 Observações Importantes

- Os tokens gerados são criptografados com **XChaCha20-Poly1305** e protegidos por **Scrypt** para derivação da chave.
- Os HMACs são baseados em **SHA-256** com chave secreta fixa (pode ser movida para variável de ambiente em produção).
- Todos os HMACs possuem **validação de timestamp** com tolerância configurável (padrão: 30 segundos).

---

## 📜 Exemplo de Fluxo de Uso (Resumo)

1. Cliente chama `POST /genhmac` para obter um token HMAC + timestamp.
2. Cliente utiliza esses valores (`hmac`, `timestamp`, `sig`) para acessar endpoints protegidos (`/protected`, `/tokenize`, `/detokenize`).
3. Servidor valida assinatura HMAC + validade de tempo para permitir ou negar a requisição.

---

## 📜 Licença

Este projeto está sob a licença MIT.  
Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

2025, Desenvolvido por [Thiago Zilli Sarmento](https://github.com/thiagozs) ❤️
