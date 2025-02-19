package config

type DB struct {
	Postgres Postgres `yaml:"postgres"`
}

type Postgres struct {
	PoolMax int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
	URL     string `env-required:"true"                 env:"PG_URL"`
}
