package windowx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

type ThemeManager struct {
	WindowSetSystemDefaultThemeFn func(ctx context.Context) error
	WindowSetDarkThemeFn          func(ctx context.Context) error
	WindowSetLightThemeFn         func(ctx context.Context) error
}

func (th ThemeManager) WindowSetDarkTheme(ctx context.Context) error {
	return safe.Run(func() error {
		if th.WindowSetDarkThemeFn == nil {
			return wailsrun.WindowSetDarkTheme(ctx)
		}
		return th.WindowSetDarkThemeFn(ctx)
	})
}

func (th ThemeManager) WindowSetLightTheme(ctx context.Context) error {
	return safe.Run(func() error {
		if th.WindowSetLightThemeFn == nil {
			return wailsrun.WindowSetLightTheme(ctx)
		}

		return th.WindowSetLightThemeFn(ctx)
	})
}

func (th ThemeManager) WindowSetSystemDefaultTheme(ctx context.Context) error {
	return safe.Run(func() error {
		if th.WindowSetSystemDefaultThemeFn == nil {
			return wailsrun.WindowSetSystemDefaultTheme(ctx)
		}

		return th.WindowSetSystemDefaultThemeFn(ctx)
	})
}
