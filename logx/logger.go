package logx

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"sync"

	"github.com/markbates/wailsx/wailsrun"
)

var _ WailsLogger = &Logger{}

type Logger struct {
	*slog.Logger `json:"-"`

	level wailsrun.LogLevel
	mu    sync.Mutex
}

func NopLogger() *Logger {
	return NewLogger(io.Discard, wailsrun.INFO)
}

func NewLogger(w io.Writer, level wailsrun.LogLevel) *Logger {
	var ll slog.Level

	switch level {
	case wailsrun.DEBUG:
		ll = slog.LevelDebug
	case wailsrun.INFO:
		ll = slog.LevelInfo
	case wailsrun.WARNING:
		ll = slog.LevelWarn
	case wailsrun.ERROR:
		ll = slog.LevelError
	case wailsrun.FATAL:
		ll = slog.LevelError
	case wailsrun.TRACE:
		ll = slog.LevelDebug
	default:
		ll = slog.LevelInfo
	}

	sl := slog.New(slog.NewTextHandler(w, &slog.HandlerOptions{
		Level: ll,
	}))

	return &Logger{
		Logger: sl,
		level:  level,
	}
}

func (l *Logger) LogDebug(ctx context.Context, message string) error {
	if err := l.init(); err != nil {
		return err
	}

	l.Logger.Debug(message)
	return nil
}

func (l *Logger) LogDebugf(ctx context.Context, format string, args ...any) error {
	if err := l.init(); err != nil {
		return err
	}

	l.Debug(fmt.Sprintf(format, args...))
	return nil
}

func (l *Logger) LogError(ctx context.Context, message string) error {
	if err := l.init(); err != nil {
		return err
	}

	l.Error(message)
	return nil
}

func (l *Logger) LogErrorf(ctx context.Context, format string, args ...any) error {
	if err := l.init(); err != nil {
		return err
	}

	l.Error(fmt.Sprintf(format, args...))
	return nil
}

func (l *Logger) LogFatal(ctx context.Context, message string) error {
	if err := l.init(); err != nil {
		return err
	}

	l.Error(message)
	return nil
}

func (l *Logger) LogFatalf(ctx context.Context, format string, args ...any) error {
	if err := l.init(); err != nil {
		return err
	}

	l.LogErrorf(ctx, format, args...)
	return nil
}

func (l *Logger) LogInfo(ctx context.Context, message string) error {
	if err := l.init(); err != nil {
		return err
	}

	l.Info(message)
	return nil
}

func (l *Logger) LogInfof(ctx context.Context, format string, args ...any) error {
	if err := l.init(); err != nil {
		return err
	}

	l.Info(fmt.Sprintf(format, args...))
	return nil
}

func (l *Logger) LogPrint(ctx context.Context, message string) error {
	if err := l.init(); err != nil {
		return err
	}

	l.Info(message)
	return nil
}

func (l *Logger) LogPrintf(ctx context.Context, format string, args ...any) error {
	if err := l.init(); err != nil {
		return err
	}

	l.Info(fmt.Sprintf(format, args...))
	return nil
}

func (l *Logger) LogSetLogLevel(ctx context.Context, level wailsrun.LogLevel) error {
	if err := l.init(); err != nil {
		return err
	}

	l.Info(fmt.Sprintf("LogSetLogLevel [NOOP]: %d", level))
	return nil
}

func (l *Logger) LogTrace(ctx context.Context, message string) error {
	if err := l.init(); err != nil {
		return err
	}

	l.Info(message)
	return nil
}

func (l *Logger) LogTracef(ctx context.Context, format string, args ...any) error {
	if err := l.init(); err != nil {
		return err
	}

	l.Info(fmt.Sprintf(format, args...))
	return nil
}

func (l *Logger) LogWarning(ctx context.Context, message string) error {
	if err := l.init(); err != nil {
		return err
	}

	l.Warn(message)
	return nil
}

func (l *Logger) LogWarningf(ctx context.Context, format string, args ...any) error {
	if err := l.init(); err != nil {
		return err
	}

	l.Warn(fmt.Sprintf(format, args...))
	return nil
}

func (l *Logger) init() error {
	if l == nil {
		return fmt.Errorf("Logger is nil")
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.Logger == nil {
		l.Logger = NewLogger(os.Stdout, wailsrun.INFO).Logger
	}

	return nil
}
