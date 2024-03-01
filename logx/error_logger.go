package logx

import (
	"context"

	"github.com/markbates/wailsx/wailsrun"
)

type ErrorLoggable interface {
	LogErrorf(ctx context.Context, format string, args ...any) error
	LogError(ctx context.Context, message string) error
}

type ErrorLogger struct {
	LogErrorFn  func(ctx context.Context, message string) error
	LogErrorfFn func(ctx context.Context, format string, args ...any) error
}

func (el ErrorLogger) LogError(ctx context.Context, message string) error {
	if el.LogErrorFn != nil {
		return el.LogErrorFn(ctx, message)
	}

	return wailsrun.LogError(ctx, message)
}

func (el ErrorLogger) LogErrorf(ctx context.Context, format string, args ...any) error {
	if el.LogErrorfFn != nil {
		return el.LogErrorfFn(ctx, format, args...)
	}

	return wailsrun.LogErrorf(ctx, format, args...)
}
