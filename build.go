package config

import "strings"

// Build allows adding version number, unique hash from the version-control system and an web description.
type Build struct {
	Environment string `env:"ENV,default=development"`
	Machine     string `env:"MACHINE,default=local"`
	AppName     string `env:"APP_NAME,default=myapp"`
	Version     string `env:"VERSION,default=0.0.0"`
	Hash        string `env:"HASH,default=000000"`
	Desc        string `env:"DESC,default=Copyright 2022"`
}

// IsDevelopment indicates whether the binary is not running in non-optimized development mode, or production optimized build.
func (b *Build) IsDevelopment() bool {
	return strings.EqualFold(b.Environment, "development")
}

// IsRemote indicates whether the binary is not running on a remote server or a local workstation.
func (b *Build) IsRemote() bool {
	return strings.EqualFold(b.Machine, "remote")
}

// IsLocal an alias for IsRemote, indicates whether the binary is running on a local workstation.
func (b *Build) IsLocal() bool {
	return !b.IsRemote()
}
