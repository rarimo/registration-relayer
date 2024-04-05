package handlers

import (
	"context"
	"net/http"

	"github.com/rarimo/registration-relayer/internal/config"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	relayerConfigCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxRelayerConfig(cfg *config.RelayerConfig) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, relayerConfigCtxKey, cfg)
	}
}

func RelayerConfig(r *http.Request) *config.RelayerConfig {
	return r.Context().Value(relayerConfigCtxKey).(*config.RelayerConfig)
}
