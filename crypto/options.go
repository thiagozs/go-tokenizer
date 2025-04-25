package crypto

import "github.com/thiagozs/go-tokenizer/config"

type Options func(*SymmetricParams) error

type SymmetricParams struct {
	Config *config.Config
}

func newSymmetricParams(opts ...Options) (*SymmetricParams, error) {
	params := &SymmetricParams{}
	for _, opt := range opts {
		if err := opt(params); err != nil {
			return nil, err
		}
	}

	return params, nil
}

func WithConfig(cfg *config.Config) Options {
	return func(p *SymmetricParams) error {
		p.Config = cfg
		return nil
	}
}

func (s *SymmetricParams) GetConfig() *config.Config {
	return s.Config
}
