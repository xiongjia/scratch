package util

import (
	"log/slog"
	"runtime/debug"
)

type (
	DebugRuntimeInfo struct {
		Deps []*debug.Module
	}
)

var (
	runtimeInfo *DebugRuntimeInfo
)

func init() {
	runtimeInfo = &DebugRuntimeInfo{}
	if bi, ok := debug.ReadBuildInfo(); ok {
		runtimeInfo.Deps = bi.Deps
	}
}

func DebugDumpDeps() {
	for _, depModule := range runtimeInfo.Deps {
		if depModule == nil {
			continue
		}

		slog.Info("Module",
			slog.String("path", depModule.Path),
			slog.String("ver", depModule.Version))
	}
}
