package prom

import (
	"context"
	"time"

	kitlog "github.com/go-kit/log"
)

type (
	EngineOptions struct {
		Logger LogAdapterHandler

		Disable bool
		DBPath  string

		QuerierMaxMaxSamples int
		QuerierTimeout       time.Duration
	}

	Engine struct {
		logger kitlog.Logger

		ctx     context.Context
		disable bool

		promDiscovery *PromDiscovery
		promStorage   *PromStorage
		promQuerier   *PromQuerier
		promScrape    *PromScrape
		promEngApi    *EngineAPI
	}
)

func (e *Engine) GetQuerier() *PromQuerier {
	return e.promQuerier
}

func (e *Engine) GetStorage() *PromStorage {
	return e.promStorage
}
