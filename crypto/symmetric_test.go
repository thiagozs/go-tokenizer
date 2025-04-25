package crypto

import (
	"testing"

	"github.com/thiagozs/go-tokenizer/config"
)

func TestEncryptDecryptSymmetric(t *testing.T) {
	original := "12345678900"

	cfg, err := config.NewConfig()
	if err != nil {
		t.Fatalf("Erro ao criar configuração: %v", err)
	}

	es, err := NewSymmetric(WithConfig(cfg))
	if err != nil {
		t.Fatalf("Erro ao criar instância de criptografia simétrica: %v", err)
	}

	token, err := es.EncryptSymmetric([]byte(original))
	if err != nil {
		t.Fatalf("Erro ao criptografar: %v", err)
	}

	if token == "" {
		t.Fatal("Token está vazio após criptografia")
	}

	plaintext, err := es.DecryptSymmetric(token)
	if err != nil {
		t.Fatalf("Erro ao descriptografar: %v", err)
	}

	if string(plaintext) != original {
		t.Errorf("Esperava %s, mas recebeu %s", original, string(plaintext))
	}
}

func TestDecryptSymmetricWithInvalidToken(t *testing.T) {
	// Token inválido (não é base64 ou truncado)
	invalidToken := "token_invalido_curto"

	cfg, err := config.NewConfig()
	if err != nil {
		t.Fatalf("Erro ao criar configuração: %v", err)
	}

	es, err := NewSymmetric(WithConfig(cfg))
	if err != nil {
		t.Fatalf("Erro ao criar instância de criptografia simétrica: %v", err)
	}

	_, err = es.DecryptSymmetric(invalidToken)
	if err == nil {
		t.Error("Esperava erro ao descriptografar token inválido, mas não ocorreu")
	}
}

func TestEncryptSymmetricWithEmptyString(t *testing.T) {
	// Testar criptografia com string vazia
	emptyString := ""

	cfg, err := config.NewConfig()
	if err != nil {
		t.Fatalf("Erro ao criar configuração: %v", err)
	}

	es, err := NewSymmetric(WithConfig(cfg))
	if err != nil {
		t.Fatalf("Erro ao criar instância de criptografia simétrica: %v", err)
	}

	token, err := es.EncryptSymmetric([]byte(emptyString))
	if err != nil {
		t.Fatalf("Erro ao criptografar string vazia: %v", err)
	}

	if token == "" {
		t.Fatal("Token está vazio após criptografia de string vazia")
	}

	plaintext, err := es.DecryptSymmetric(token)
	if err != nil {
		t.Fatalf("Erro ao descriptografar token de string vazia: %v", err)
	}

	if string(plaintext) != emptyString {
		t.Errorf("Esperava %s, mas recebeu %s", emptyString, string(plaintext))
	}
}

func TestEncryptSymmetricWithSpecialCharacters(t *testing.T) {
	// Testar criptografia com caracteres especiais
	specialChars := "!@#$%^&*()_+{}|:\"<>?"

	cfg, err := config.NewConfig()
	if err != nil {
		t.Fatalf("Erro ao criar configuração: %v", err)
	}

	es, err := NewSymmetric(WithConfig(cfg))
	if err != nil {
		t.Fatalf("Erro ao criar instância de criptografia simétrica: %v", err)
	}

	token, err := es.EncryptSymmetric([]byte(specialChars))
	if err != nil {
		t.Fatalf("Erro ao criptografar string com caracteres especiais: %v", err)
	}

	if token == "" {
		t.Fatal("Token está vazio após criptografia de string com caracteres especiais")
	}

	plaintext, err := es.DecryptSymmetric(token)
	if err != nil {
		t.Fatalf("Erro ao descriptografar token de string com caracteres especiais: %v", err)
	}

	if string(plaintext) != specialChars {
		t.Errorf("Esperava %s, mas recebeu %s", specialChars, string(plaintext))
	}
}

func TestEncryptSymmetricWithLongString(t *testing.T) {
	// Testar criptografia com string longa
	longString := "a" + string(make([]byte, 10000)) + "b"

	cfg, err := config.NewConfig()
	if err != nil {
		t.Fatalf("Erro ao criar configuração: %v", err)
	}

	es, err := NewSymmetric(WithConfig(cfg))
	if err != nil {
		t.Fatalf("Erro ao criar instância de criptografia simétrica: %v", err)
	}

	token, err := es.EncryptSymmetric([]byte(longString))
	if err != nil {
		t.Fatalf("Erro ao criptografar string longa: %v", err)
	}

	if token == "" {
		t.Fatal("Token está vazio após criptografia de string longa")
	}

	plaintext, err := es.DecryptSymmetric(token)
	if err != nil {
		t.Fatalf("Erro ao descriptografar token de string longa: %v", err)
	}

	if string(plaintext) != longString {
		t.Errorf("Esperava %s, mas recebeu %s", longString, string(plaintext))
	}
}

func TestEncryptSymmetricWithNilInput(t *testing.T) {
	// Testar criptografia com entrada nula
	var nilInput []byte

	cfg, err := config.NewConfig()
	if err != nil {
		t.Fatalf("Erro ao criar configuração: %v", err)
	}

	es, err := NewSymmetric(WithConfig(cfg))
	if err != nil {
		t.Fatalf("Erro ao criar instância de criptografia simétrica: %v", err)
	}

	token, err := es.EncryptSymmetric(nilInput)
	if err != nil {
		t.Fatalf("Erro ao criptografar entrada nula: %v", err)
	}

	if token == "" {
		t.Fatal("Token está vazio após criptografia de entrada nula")
	}

	plaintext, err := es.DecryptSymmetric(token)
	if err != nil {
		t.Fatalf("Erro ao descriptografar token de entrada nula: %v", err)
	}

	if string(plaintext) != "" {
		t.Errorf("Esperava string vazia, mas recebeu %s", string(plaintext))
	}
}
