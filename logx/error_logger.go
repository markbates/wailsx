package logx

import (
	"context"

	wailsrun "github.com/wailsapp/wails/v2/pkg/runtime"
)

type ErrorLoggable interface {
	LogErrorf(ctx context.Context, format string, args ...any)
	LogError(ctx context.Context, message string)
}

type ErrorLogger struct {
	LogErrorFn  func(ctx context.Context, message string)
	LogErrorfFn func(ctx context.Context, format string, args ...any)
}

func (el ErrorLogger) LogError(ctx context.Context, message string) {
	if el.LogErrorFn != nil {
		el.LogErrorFn(ctx, message)
		return
	}

	wailsrun.LogError(ctx, message)
}

func (el ErrorLogger) LogErrorf(ctx context.Context, format string, args ...any) {
	if el.LogErrorfFn != nil {
		el.LogErrorfFn(ctx, format, args...)
		return
	}

	wailsrun.LogErrorf(ctx, format, args...)
}
