package metric

import "log/slog"

type (
	Engine struct {
	}
)

func NewEngine() (*Engine, error) {
	return &Engine{}, nil
}

func (*Engine) Run() {
	slog.Debug("Engine running")
}
