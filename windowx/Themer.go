package windowx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

var _ ThemeManager = Themer{}

func NewNOOPThemer() Themer {
	return Themer{
		WindowSetBackgroundColourFn:   func(ctx context.Context, R, G, B, A uint8) error { return nil },
		WindowSetDarkThemeFn:          func(ctx context.Context) error { return nil },
		WindowSetLightThemeFn:         func(ctx context.Context) error { return nil },
		WindowSetSystemDefaultThemeFn: func(ctx context.Context) error { return nil },
	}
}

type Themer struct {
	WindowSetBackgroundColourFn   func(ctx context.Context, R, G, B, A uint8) error
	WindowSetDarkThemeFn          func(ctx context.Context) error
	WindowSetLightThemeFn         func(ctx context.Context) error
	WindowSetSystemDefaultThemeFn func(ctx context.Context) error

	data ThemeData
}

func (th Themer) WindowSetDarkTheme(ctx context.Context) error {
	return safe.Run(func() error {
		if th.WindowSetDarkThemeFn == nil {
			return wailsrun.WindowSetDarkTheme(ctx)
		}
		return th.WindowSetDarkThemeFn(ctx)
	})
}

func (th Themer) WindowSetLightTheme(ctx context.Context) error {
	return safe.Run(func() error {
		if th.WindowSetLightThemeFn == nil {
			return wailsrun.WindowSetLightTheme(ctx)
		}

		return th.WindowSetLightThemeFn(ctx)
	})
}

func (th Themer) WindowSetSystemDefaultTheme(ctx context.Context) error {
	return safe.Run(func() error {
		if th.WindowSetSystemDefaultThemeFn == nil {
			return wailsrun.WindowSetSystemDefaultTheme(ctx)
		}

		return th.WindowSetSystemDefaultThemeFn(ctx)
	})
}

func (th Themer) WindowSetBackgroundColour(ctx context.Context, R uint8, G uint8, B uint8, A uint8) error {
	return safe.Run(func() error {
		if th.WindowSetBackgroundColourFn == nil {
			return wailsrun.WindowSetBackgroundColour(ctx, R, G, B, A)
		}

		return th.WindowSetBackgroundColourFn(ctx, R, G, B, A)
	})
}
