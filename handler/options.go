package handler

import (
	"github.com/thiagozs/go-tokenizer/config"
	"github.com/thiagozs/go-tokenizer/crypto"
	"github.com/thiagozs/go-tokenizer/token"
)

type Options func(*HandlerParams) error

type HandlerParams struct {
	Config *config.Config
	Crypto *crypto.Symmetric
	Token  *token.Token
}

func newHandlerParams(opts ...Options) (*HandlerParams, error) {
	params := &HandlerParams{}

	for _, opt := range opts {
		if err := opt(params); err != nil {
			return nil, err
		}
	}

	return params, nil
}

func WithConfig(cfg *config.Config) Options {
	return func(p *HandlerParams) error {
		p.Config = cfg
		return nil
	}
}

func WithCrypto(crypto *crypto.Symmetric) Options {
	return func(p *HandlerParams) error {
		p.Crypto = crypto
		return nil
	}
}

func WithToken(token *token.Token) Options {
	return func(p *HandlerParams) error {
		p.Token = token
		return nil
	}
}

func (h *HandlerParams) GetConfig() *config.Config {
	return h.Config
}

func (h *HandlerParams) GetCrypto() *crypto.Symmetric {
	return h.Crypto
}

func (h *HandlerParams) GetToken() *token.Token {
	return h.Token
}

func (h *HandlerParams) SetConfig(cfg *config.Config) {
	h.Config = cfg
}

func (h *HandlerParams) SetCrypto(crypto *crypto.Symmetric) {
	h.Crypto = crypto
}

func (h *HandlerParams) SetToken(token *token.Token) {
	h.Token = token
}

func (h *HandlerParams) GetPassphrase() string {
	return h.Config.GetPassphrase()
}

func (h *HandlerParams) GetSalt() string {
	return h.Config.GetSalt()
}

func (h *HandlerParams) GetPort() string {
	return h.Config.GetPort()
}

func (h *HandlerParams) GetHost() string {
	return h.Config.GetHost()
}

func (h *HandlerParams) GetTolSec() string {
	return h.Config.GetTolSec()
}

func (h *HandlerParams) SetPassphrase(passphrase string) {
	h.Config.SetPassphrase(passphrase)
}

func (h *HandlerParams) SetSalt(salt string) {
	h.Config.SetSalt(salt)
}

func (h *HandlerParams) SetPort(port string) {
	h.Config.SetPort(port)
}

func (h *HandlerParams) SetHost(host string) {
	h.Config.SetHost(host)
}

func (h *HandlerParams) SetTolSecs(timeout string) {
	h.Config.SetTolSecs(timeout)
}

func (h *HandlerParams) GetHandlerParams() *HandlerParams {
	return h
}

func (h *HandlerParams) SetHandlerParams(params *HandlerParams) {
	h.Config = params.Config
	h.Crypto = params.Crypto
	h.Token = params.Token
}
