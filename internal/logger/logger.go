package logger

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

func New(isDebug bool) *slog.Logger {
	logLevel := slog.LevelInfo
	if isDebug {
		logLevel = slog.LevelDebug
	}

	logger := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel, ReplaceAttr: replaceAttr})

	return slog.New(logger)
}

func replaceAttr(groups []string, a slog.Attr) slog.Attr {
	switch a.Value.Kind() {

	case slog.KindAny:
		switch v := a.Value.Any().(type) {
		case error:
			a.Value = formatError(v)
		}
	}

	return a
}

func formatError(err error) slog.Value {
	var groupValues []slog.Attr

	groupValues = append(groupValues, slog.String("msg", err.Error()))

	type StackTracer interface {
		StackTrace() errors.StackTrace
	}

	var st StackTracer
	for err := err; err != nil; err = errors.Unwrap(err) {
		if x, ok := err.(StackTracer); ok {
			st = x
		}
	}

	if st != nil {
		groupValues = append(groupValues,
			slog.Any("trace", traceLines(st.StackTrace())),
		)
	}

	return slog.GroupValue(groupValues...)
}

func traceLines(frames errors.StackTrace) []string {
	traceLines := make([]string, len(frames))

	var skipped int
	skipping := true
	for i := len(frames) - 1; i >= 0; i-- {
		pc := uintptr(frames[i]) - 1
		fn := runtime.FuncForPC(pc)
		if fn == nil {
			traceLines[i] = "unknown"
			skipping = false
			continue
		}

		name := fn.Name()

		if skipping && strings.HasPrefix(name, "runtime.") {
			skipped++
			continue
		} else {
			skipping = false
		}

		filename, lineNr := fn.FileLine(pc)

		traceLines[i] = fmt.Sprintf("%s %s:%d", name, filename, lineNr)
	}

	return traceLines[:len(traceLines)-skipped]
}
