package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type ProtectedRequest struct {
	Signature string `json:"sig"`
}

// ProtectedEndpoint valida HMAC via body (usando JSON field "sig")
func (h *Handlers) ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	hmacHeader := r.Header.Get("X-HMAC-Token")
	tsHeader := r.Header.Get("X-HMAC-Timestamp")

	if hmacHeader == "" || tsHeader == "" {
		http.Error(w, "Cabeçalhos obrigatórios ausentes", http.StatusUnauthorized)
		return
	}

	ts, err := strconv.ParseInt(tsHeader, 10, 64)
	if err != nil {
		http.Error(w, "Timestamp inválido", http.StatusBadRequest)
		return
	}

	var req ProtectedRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	if req.Signature == "" {
		http.Error(w, "Campo 'sig' ausente no corpo", http.StatusBadRequest)
		return
	}

	interval, err := strconv.ParseInt(h.Config.GetInterval(), 10, 64)
	if err != nil {
		http.Error(w, "Intervalo inválido", http.StatusInternalServerError)
		return
	}

	if !h.Token.ValidateTimedHMAC(req.Signature, hmacHeader, ts, interval) {
		http.Error(w, "HMAC inválido ou expirado", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("✅ Acesso autorizado com HMAC temporizado"))
}

// RequireHMAC é um middleware que valida HMAC no header antes de chamar o handler
func (h *Handlers) RequireHMAC(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hmacHeader := r.Header.Get("X-HMAC-Token")
		signature := r.Header.Get("X-HMAC-Signature")
		tsHeader := r.Header.Get("X-HMAC-Timestamp")

		if hmacHeader == "" || signature == "" || tsHeader == "" {
			http.Error(w, "Cabeçalhos obrigatórios ausentes", http.StatusUnauthorized)
			return
		}

		ts, err := strconv.ParseInt(tsHeader, 10, 64)
		if err != nil {
			http.Error(w, "Timestamp inválido", http.StatusBadRequest)
			return
		}

		interval, err := strconv.ParseInt(h.Config.GetInterval(), 10, 64)
		if err != nil {
			http.Error(w, "Intervalo inválido", http.StatusInternalServerError)
			return
		}

		if !h.Token.ValidateTimedHMAC(signature, hmacHeader, ts, interval) {
			http.Error(w, "HMAC inválido ou expirado", http.StatusUnauthorized)
			return
		}

		// Continua para o handler real
		next(w, r)
	}
}
