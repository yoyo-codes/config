package config

// PostgresDB holds a GCP project configurations.
type PostgresDB struct {
	User         string `env:"DB_USER,default=postgres"`
	Password     string `env:"DB_PASS,default=postgres"`
	Host         string `env:"DB_PASS,default=localhost:5432"`
	Name         string `env:"DB_PASS,default=postgres"`
	MaxIdleConns int    `env:"DB_MAX_IDLE_CONNS,default=0"`
	MaxOpenConns int    `env:"DB_MAX_OPEN_CONNS,default=0"`
}
