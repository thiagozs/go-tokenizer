package token

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/thiagozs/go-tokenizer/config"
)

type Token struct {
	Config *config.Config
}

func NewToken(opts ...Options) (*Token, error) {
	params, err := newTokenParams(opts...)
	if err != nil {
		return nil, err
	}

	token := &Token{
		Config: params.Config,
	}

	return token, nil
}

// Gera HMAC com valor + timestamp
func (t *Token) GenerateTimedHMAC(value string, timestamp int64) string {
	data := fmt.Sprintf("%s|%d", value, timestamp)
	h := hmac.New(sha256.New, []byte(t.Config.GetSecretKey()))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))[:16]
}

// Valida o HMAC com tolerância de tempo (ex: 30s)
func (t *Token) ValidateTimedHMAC(signature string, receivedHMAC string,
	receivedTS int64, toleranceSeconds int64) bool {
	now := time.Now().Unix()

	// Verifica tolerância de tempo
	if receivedTS < now-toleranceSeconds || receivedTS > now+toleranceSeconds {
		return false
	}

	expected := t.GenerateTimedHMAC(signature, receivedTS)

	return hmac.Equal([]byte(expected), []byte(receivedHMAC))
}
