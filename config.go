package config

import (
	"context"
	"github.com/sethvargo/go-envconfig"
)

func Parse(mc interface{}) error {
	ctx := context.Background()
	return envconfig.Process(ctx, mc)
}
