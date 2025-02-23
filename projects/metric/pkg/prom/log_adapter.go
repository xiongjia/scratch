package prom

import (
	"fmt"
	"log/slog"
	"strings"

	kitlog "github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

const (
	LogAdapterLevelError = "error"
	LogAdapterLevelWarn  = "warn"
	LogAdapterLevelInfo  = "info"
	LogAdapterLevelDebug = "debug"

	kitLogMessageKey = "msg"
)

type (
	LogAdapterHandler interface {
		Append(entry *LogEntry)
	}

	LogAdapter struct {
		handler LogAdapterHandler
	}

	LogAdapterField struct {
		Key string
		Val any
	}

	LogEntry struct {
		Level  string
		Msg    string
		Fields []LogAdapterField
	}

	slogAdapter struct{}
)

func (slogAdapter) Append(e *LogEntry) {
	switch e.Level {
	case LogAdapterLevelError:
		slog.Error(e.MergeMessage())
	case LogAdapterLevelWarn:
		slog.Warn(e.MergeMessage())
	case LogAdapterLevelInfo:
		slog.Info(e.MergeMessage())
	case LogAdapterLevelDebug:
		slog.Debug(e.MergeMessage())
	}
}

func convertLogEntry(keyvals ...interface{}) *LogEntry {
	klen := len(keyvals)
	if klen == 0 || klen%2 != 0 {
		return nil
	}
	l := level.InfoValue()
	msg := ""
	fields := make([]LogAdapterField, 0)
	for i := 0; i < len(keyvals); i += 2 {
		if keyvals[i] == level.Key() {
			l = keyvals[i+1].(level.Value)
			continue
		}
		if keyvals[i] == kitLogMessageKey {
			msg = keyvals[i+1].(string)
			continue
		}
		k := keyvals[i].(string)
		fields = append(fields, LogAdapterField{Key: k, Val: keyvals[i+1]})
	}

	return &LogEntry{Level: l.String(), Msg: msg, Fields: fields}
}

func NewSLogAdapterHandler() LogAdapterHandler {
	return &slogAdapter{}
}

func NewSLogAdapter() kitlog.Logger {
	return NewLogAdapter(NewSLogAdapterHandler())
}

func NewLogAdapter(h LogAdapterHandler) kitlog.Logger {
	return &LogAdapter{handler: h}
}

func (a *LogAdapter) Log(keyvals ...interface{}) error {
	if a.handler == nil {
		return nil
	}
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	ent := convertLogEntry(keyvals...)
	if ent == nil {
		return nil
	}
	a.handler.Append(ent)
	return nil
}

func (e *LogEntry) MergeMessage() (rt string) {
	rt = ""
	var b strings.Builder
	defer func() {
		if r := recover(); r != nil {
			return
		}
	}()
	_, _ = b.WriteString(e.Msg)
	_, _ = b.WriteString(" - ")
	for _, f := range e.Fields {
		if f.Key == "" || f.Val == nil {
			continue
		}
		v := fmt.Sprintf("%v", f.Val)
		b.WriteString("{")
		b.WriteString(f.Key)
		b.WriteString(" - ")
		b.WriteString(v)
		b.WriteString("}")
	}
	return b.String()
}
