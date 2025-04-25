package crypto

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/thiagozs/go-tokenizer/config"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/scrypt"
)

type Symmetric struct {
	Config *config.Config
}

func NewSymmetric(opts ...Options) (*Symmetric, error) {
	params, err := newSymmetricParams(opts...)
	if err != nil {
		return &Symmetric{}, fmt.Errorf("erro ao criar parâmetros simétricos: %v", err)
	}

	return &Symmetric{
		Config: params.GetConfig(),
	}, nil
}

func (s *Symmetric) deriveKey(pass []byte) ([]byte, error) {
	return scrypt.Key(pass, []byte(s.Config.Salt), 1<<15, 8, 1, chacha20poly1305.KeySize)
}

func (s *Symmetric) EncryptSymmetric(plaintext []byte) (string, error) {
	key, err := s.deriveKey([]byte(s.Config.Passphrase))
	if err != nil {
		return "", fmt.Errorf("erro ao derivar chave: %w", err)
	}

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return "", fmt.Errorf("erro ao criar AEAD: %w", err)
	}

	nonce := make([]byte, chacha20poly1305.NonceSizeX)
	if _, err := rand.Read(nonce); err != nil {
		return "", fmt.Errorf("erro ao gerar nonce: %w", err)
	}

	cipher := aead.Seal(nil, nonce, plaintext, nil)
	data := append(nonce, cipher...)

	return base64.RawURLEncoding.EncodeToString(data), nil
}

func (s *Symmetric) DecryptSymmetric(encoded string) ([]byte, error) {
	data, err := base64.RawURLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, err
	}

	key, err := s.deriveKey([]byte(s.Config.Passphrase))
	if err != nil {
		return nil, fmt.Errorf("erro ao derivar chave: %w", err)
	}

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar AEAD: %w", err)
	}

	if len(data) < chacha20poly1305.NonceSizeX {
		return nil, fmt.Errorf("tamanho inválido do token: %d", len(data))
	}

	nonce := data[:chacha20poly1305.NonceSizeX]
	ciphertext := data[chacha20poly1305.NonceSizeX:]

	return aead.Open(nil, nonce, ciphertext, nil)
}
