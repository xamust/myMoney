package config

type DB struct {
	Postgres Postgres `yaml:"postgres"`
}

type Postgres struct {
	Host     string `env-required:"true" yaml:"host" env:"DB_PG_HOST"`
	Port     string `env-required:"true" yaml:"port" env:"DB_PG_PORT"`
	User     string `env-required:"true" yaml:"user" env:"DB_PG_USER"`
	Password string `env-required:"true" yaml:"password" env:"DB_PG_PASSWORD"`
	Database string `env-required:"true" yaml:"database" env:"DB_PG_DATABASE"`
	SSL      string `env-required:"true" yaml:"ssl" env:"DB_PG_SSL"`
	Debug    bool   `env-required:"true" yaml:"debug" env:"DB_PG_DEBUG"`
	MaxIdle  int    `env-required:"true" yaml:"max_idle" env:"DB_PG_MAX_IDLE"`
	MaxOpen  int    `env-required:"true" yaml:"max_open" env:"DB_PG_MAX_OPEN"`
}
