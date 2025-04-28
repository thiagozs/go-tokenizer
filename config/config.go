package config

import "github.com/caarlos0/env"

type Config struct {
	Passphrase string `json:"passphrase" env:"TKZ_PASSPHRASE" envDefault:"default-passphrase"`
	Salt       string `json:"salt" env:"TKZ_SALT" envDefault:"default-salt"`
	SecretKey  string `json:"secret_key" env:"TKZ_SECRET_KEY" envDefault:"default-secret-key"`
	Port       string `json:"port" env:"TKZ_PORT" envDefault:"8880"`
	Host       string `json:"host" env:"TKZ_HOST" envDefault:"localhost"`
	TolSecs    string `json:"timeout" env:"TKZ_TOLERANCE_SECS" envDefault:"30"` // segundos
	Interval   string `json:"interval" env:"TKZ_INTERVAL" envDefault:"30"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func (c *Config) GetPassphrase() string {
	return c.Passphrase
}

func (c *Config) GetSalt() string {
	return c.Salt
}

func (c *Config) GetPort() string {
	return c.Port
}

func (c *Config) GetHost() string {
	return c.Host
}

func (c *Config) GetTolSec() string {
	return c.TolSecs
}

func (c *Config) GetSecretKey() string {
	return c.SecretKey
}

func (c *Config) GetInterval() string {
	return c.Interval
}

func (c *Config) SetPassphrase(passphrase string) {
	c.Passphrase = passphrase
}

func (c *Config) SetSalt(salt string) {
	c.Salt = salt
}

func (c *Config) SetPort(port string) {
	c.Port = port
}

func (c *Config) SetHost(host string) {
	c.Host = host
}

func (c *Config) SetTolSecs(tolSecs string) {
	c.TolSecs = tolSecs
}

func (c *Config) SetSecretKey(secretKey string) {
	c.SecretKey = secretKey
}

func (c *Config) GetConfig() *Config {
	return c
}

func (c *Config) SetConfig(cfg *Config) {
	c.Passphrase = cfg.Passphrase
	c.Salt = cfg.Salt
	c.Port = cfg.Port
	c.Host = cfg.Host
	c.TolSecs = cfg.TolSecs
	c.SecretKey = cfg.SecretKey
	c.Interval = cfg.Interval
	c.Passphrase = cfg.Passphrase
}
