package config

import (
	"sync"

	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"

	"practice/lib/environment"
)

type RedisConfig struct {
	DB   int    `env:"DB"`
	Host string `env:"HOST"`
}

type DBConfig struct {
	URL                string `env:"URL"`
	MaxIdleConnections int    `env:"MAX_IDLE_CONNECTIONS"`
	MaxOpenConnections int    `env:"MAX_OPEN_CONNECTIONS"`
}

type Config struct {
	Env   environment.Environment `env:"ENV"`
	Port  int                     `env:"PORT"`
	Redis RedisConfig             `envPrefix:"REDIS_"`
	DB    DBConfig                `envPrefix:"DATABASE_"`
}

var (
	Cfg        *Config
	configOnce sync.Once
)

func parseConfig() (*Config, error) {
	// https://github.com/spf13/viper use this if complicated config is required
	// makes all fields required if default is not defined
	err := godotenv.Load("/Users/ayush/simpl/employee-directory/development.env")
	if err != nil {
		return nil, err
	}

	opts := env.Options{RequiredIfNoDef: true}

	cfg := &Config{}
	if err := env.Parse(cfg, opts); err != nil {
		return nil, err
	}
	return cfg, nil
}

func NewConfig() *Config {
	configOnce.Do(func() {
		var err error
		Cfg, err = parseConfig()
		if err != nil {
			panic(err)
		}
	})

	return Cfg
}
