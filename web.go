package config

import "time"

// Web allows adding Web HTTP server related configuration.
type Web struct {
	Host            string        `env:"HOST"`
	Port            int           `env:"PORT,default=80"`
	ReadTimeout     time.Duration `env:"READ_TIMEOUT,default=5s"`
	WriteTimeout    time.Duration `env:"WRITE_TIMEOUT,default=10s"`
	IdleTimeout     time.Duration `env:"IDLE_TIMEOUT,default=120s"`
	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT,default=20s"`
}
