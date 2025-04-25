package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/thiagozs/go-tokenizer/config"
	"github.com/thiagozs/go-tokenizer/crypto"
	"github.com/thiagozs/go-tokenizer/token"
)

func TestProtectedEndpoint_Valid(t *testing.T) {
	value := "teste"
	ts := time.Now().Unix()

	config, err := config.NewConfig()
	if err != nil {
		t.Fatalf("Erro ao criar configuração: %v", err)
	}

	ti, err := token.NewToken(token.WithConfig(config))
	if err != nil {
		t.Fatalf("Erro ao criar token: %v", err)
	}

	crypto, err := crypto.NewSymmetric(crypto.WithConfig(config))
	if err != nil {
		t.Fatalf("Erro ao criar instância de criptografia simétrica: %v", err)
	}

	hadlr, err := NewHandlers(
		WithConfig(config),
		WithCrypto(crypto),
		WithToken(ti),
	)
	if err != nil {
		t.Fatalf("Erro ao criar manipulador: %v", err)
	}

	hmacToken := ti.GenerateTimedHMAC(value, ts)

	body, _ := json.Marshal(ProtectedRequest{Signature: value})

	req := httptest.NewRequest(http.MethodPost, "/protected", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-HMAC-Token", hmacToken)
	req.Header.Set("X-HMAC-Timestamp", strconv.FormatInt(ts, 10))

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hadlr.ProtectedEndpoint)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Código esperado 200 OK, mas foi %d", rr.Code)
	}

	expected := "✅ Acesso autorizado com HMAC temporizado"
	if rr.Body.String() != expected {
		t.Errorf("Resposta esperada %q, mas foi %q", expected, rr.Body.String())
	}
}

func TestProtectedEndpoint_InvalidHMAC(t *testing.T) {
	value := "teste"
	ts := time.Now().Unix()

	body, _ := json.Marshal(ProtectedRequest{Signature: value})

	req := httptest.NewRequest(http.MethodPost, "/protected", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-HMAC-Token", "hmacinvalido")
	req.Header.Set("X-HMAC-Timestamp", strconv.FormatInt(ts, 10))

	config, err := config.NewConfig()
	if err != nil {
		t.Fatalf("Erro ao criar configuração: %v", err)
	}

	ti, err := token.NewToken(token.WithConfig(config))
	if err != nil {
		t.Fatalf("Erro ao criar token: %v", err)
	}

	crypto, err := crypto.NewSymmetric(crypto.WithConfig(config))
	if err != nil {
		t.Fatalf("Erro ao criar instância de criptografia simétrica: %v", err)
	}

	hadlr, err := NewHandlers(
		WithConfig(config),
		WithCrypto(crypto),
		WithToken(ti),
	)
	if err != nil {
		t.Fatalf("Erro ao criar manipulador: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(hadlr.ProtectedEndpoint)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("Código esperado 401 Unauthorized, mas foi %d", rr.Code)
	}
}

func TestRequireHMAC_Valid(t *testing.T) {
	value := "teste"
	ts := time.Now().Unix()

	config, err := config.NewConfig()
	if err != nil {
		t.Fatalf("Erro ao criar configuração: %v", err)
	}

	ti, err := token.NewToken(token.WithConfig(config))
	if err != nil {
		t.Fatalf("Erro ao criar token: %v", err)
	}

	crypto, err := crypto.NewSymmetric(crypto.WithConfig(config))
	if err != nil {
		t.Fatalf("Erro ao criar instância de criptografia simétrica: %v", err)
	}

	hadlr, err := NewHandlers(
		WithConfig(config),
		WithCrypto(crypto),
		WithToken(ti),
	)
	if err != nil {
		t.Fatalf("Erro ao criar manipulador: %v", err)
	}

	hmacToken := ti.GenerateTimedHMAC(value, ts)

	req := httptest.NewRequest(http.MethodGet, "/middleware", nil)
	req.Header.Set("X-HMAC-Token", hmacToken)
	req.Header.Set("X-HMAC-Signature", value)
	req.Header.Set("X-HMAC-Timestamp", strconv.FormatInt(ts, 10))

	rr := httptest.NewRecorder()

	called := false
	protectedHandler := hadlr.RequireHMAC(func(w http.ResponseWriter, r *http.Request) {
		called = true
		w.WriteHeader(http.StatusOK)
	})

	protectedHandler.ServeHTTP(rr, req)

	if !called {
		t.Error("Handler protegido não foi chamado")
	}

	if rr.Code != http.StatusOK {
		t.Errorf("Código esperado 200 OK, mas foi %d", rr.Code)
	}
}

func TestRequireHMAC_Invalid(t *testing.T) {

	config, err := config.NewConfig()
	if err != nil {
		t.Fatalf("Erro ao criar configuração: %v", err)
	}

	ti, err := token.NewToken(token.WithConfig(config))
	if err != nil {
		t.Fatalf("Erro ao criar token: %v", err)
	}

	crypto, err := crypto.NewSymmetric(crypto.WithConfig(config))
	if err != nil {
		t.Fatalf("Erro ao criar instância de criptografia simétrica: %v", err)
	}

	hadlr, err := NewHandlers(
		WithConfig(config),
		WithCrypto(crypto),
		WithToken(ti),
	)
	if err != nil {
		t.Fatalf("Erro ao criar manipulador: %v", err)
	}

	req := httptest.NewRequest(http.MethodGet, "/middleware", nil)
	req.Header.Set("X-HMAC-Token", "hmacinvalido")
	req.Header.Set("X-HMAC-Signature", "assinaturainvalida")
	req.Header.Set("X-HMAC-Timestamp", strconv.FormatInt(time.Now().Unix(), 10))

	rr := httptest.NewRecorder()

	protectedHandler := hadlr.RequireHMAC(func(w http.ResponseWriter, r *http.Request) {
		t.Error("Handler protegido não deveria ter sido chamado com HMAC inválido")
	})

	protectedHandler.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("Código esperado 401 Unauthorized, mas foi %d", rr.Code)
	}
}
