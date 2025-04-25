package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/thiagozs/go-tokenizer/config"
	"github.com/thiagozs/go-tokenizer/crypto"
	"github.com/thiagozs/go-tokenizer/token"
)

type TokenRequest struct {
	Value string `json:"value"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type DecryptRequest struct {
	Token string `json:"token"`
}

type DecryptResponse struct {
	Value string `json:"value"`
}

type HMACRequest struct {
	Signature string `json:"sig"`
	Timestamp int64  `json:"ts,omitempty"` // opcional
}

type HMACResponse struct {
	HMAC      string `json:"hmac"`
	Timestamp int64  `json:"timestamp"`
}

type Handlers struct {
	Crypto *crypto.Symmetric
	Config *config.Config
	Token  *token.Token
}

func NewHandlers(opts ...Options) (*Handlers, error) {
	params, err := newHandlerParams(opts...)
	if err != nil {
		return nil, err
	}

	handlers := &Handlers{
		Crypto: params.GetCrypto(),
		Config: params.GetConfig(),
		Token:  params.GetToken(),
	}

	return handlers, nil
}

func (h *Handlers) Tokenize(w http.ResponseWriter, r *http.Request) {
	var req TokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	token, err := h.Crypto.EncryptSymmetric([]byte(req.Value))
	if err != nil {
		http.Error(w, "Erro ao tokenizar", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(TokenResponse{Token: token})
}

func (h *Handlers) Detokenize(w http.ResponseWriter, r *http.Request) {
	var req DecryptRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	value, err := h.Crypto.DecryptSymmetric(req.Token)
	if err != nil {
		http.Error(w, "Erro ao descriptografar", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(DecryptResponse{Value: string(value)})
}

func (h *Handlers) GenerateHMACOnline(w http.ResponseWriter, r *http.Request) {
	var req HMACRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if req.Signature == "" {
		http.Error(w, "Campo 'sig' é obrigatório", http.StatusBadRequest)
		return
	}

	ts := req.Timestamp
	if ts == 0 {
		ts = time.Now().Unix()
	}

	token := h.Token.GenerateTimedHMAC(req.Signature, ts)

	resp := HMACResponse{
		HMAC:      token,
		Timestamp: ts,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
