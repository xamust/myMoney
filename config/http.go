package config

type HTTP struct {
	Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
}
