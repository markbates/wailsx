package logx

import (
	"bytes"
	"context"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/stretchr/testify/require"
)

func Test_Logger(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	const msg = "test message"
	const msgf = "test message %s"
	const arg = "test arg"
	const expf = "test message test arg"

	type errFn func(l *Logger) error

	tcs := []struct {
		name  string
		level wailsrun.LogLevel
		fn    errFn
		exp   string
	}{
		{
			name:  "LogDebug",
			level: wailsrun.DEBUG,
			fn: func(l *Logger) error {
				return l.LogDebug(ctx, msg)
			},
			exp: msg,
		},
		{
			name:  "LogDebugf",
			level: wailsrun.DEBUG,
			fn: func(l *Logger) error {
				return l.LogDebugf(ctx, msgf, arg)
			},
			exp: expf,
		},
		{
			name:  "LogError",
			level: wailsrun.ERROR,
			fn: func(l *Logger) error {
				return l.LogError(ctx, msg)
			},
			exp: msg,
		},
		{
			name:  "LogErrorf",
			level: wailsrun.ERROR,
			fn: func(l *Logger) error {
				return l.LogErrorf(ctx, msgf, arg)
			},
			exp: expf,
		},
		{
			name:  "LogFatal",
			level: wailsrun.FATAL,
			fn: func(l *Logger) error {
				return l.LogFatal(ctx, msg)
			},
			exp: msg,
		},
		{
			name:  "LogFatalf",
			level: wailsrun.FATAL,
			fn: func(l *Logger) error {
				return l.LogFatalf(ctx, msgf, arg)
			},
			exp: expf,
		},
		{
			name:  "LogInfo",
			level: wailsrun.INFO,
			fn: func(l *Logger) error {
				return l.LogInfo(ctx, msg)
			},
			exp: msg,
		},
		{
			name: "LogInfof",
			fn: func(l *Logger) error {
				return l.LogInfof(ctx, msgf, arg)
			},
			exp: expf,
		},
		{
			name:  "LogPrint",
			level: wailsrun.PRINT,
			fn: func(l *Logger) error {
				return l.LogPrint(ctx, msg)
			},
			exp: msg,
		},
		{
			name:  "LogPrintf",
			level: wailsrun.PRINT,
			fn: func(l *Logger) error {
				return l.LogPrintf(ctx, msgf, arg)
			},
			exp: expf,
		},
		{
			name:  "LogTrace",
			level: wailsrun.TRACE,
			fn: func(l *Logger) error {
				return l.LogTrace(ctx, msg)
			},
			exp: msg,
		},
		{
			name:  "LogTracef",
			level: wailsrun.TRACE,
			fn: func(l *Logger) error {
				return l.LogTracef(ctx, msgf, arg)
			},
			exp: expf,
		},
		{
			name:  "LogWarning",
			level: wailsrun.WARNING,
			fn: func(l *Logger) error {
				return l.LogWarning(ctx, msg)
			},
			exp: msg,
		},
		{
			name:  "LogWarningf",
			level: wailsrun.WARNING,
			fn: func(l *Logger) error {
				return l.LogWarningf(ctx, msgf, arg)
			},
			exp: expf,
		},
		{
			name:  "LogSetLogLevel",
			level: wailsrun.INFO,
			fn: func(l *Logger) error {
				return l.LogSetLogLevel(ctx, wailsrun.INFO)
			},
			exp: "LogSetLogLevel [NOOP]",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)
			bb := &bytes.Buffer{}

			l := NewLogger(bb, tc.level)
			err := tc.fn(l)

			r.NoError(err)

			r.Contains(bb.String(), tc.exp)
		})
	}
}
