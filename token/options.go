package token

import "github.com/thiagozs/go-tokenizer/config"

type Options func(*TokenParams) error

type TokenParams struct {
	Config *config.Config
}

func newTokenParams(opts ...Options) (*TokenParams, error) {
	params := &TokenParams{}

	for _, opt := range opts {
		if err := opt(params); err != nil {
			return nil, err
		}
	}

	return params, nil
}

func WithConfig(cfg *config.Config) Options {
	return func(p *TokenParams) error {
		p.Config = cfg
		return nil
	}
}

func (t *TokenParams) GetConfig() *config.Config {
	return t.Config
}

func (t *TokenParams) SetConfig(cfg *config.Config) {
	t.Config = cfg
}

func (t *TokenParams) GetSecretKey() string {
	return t.Config.GetSecretKey()
}
