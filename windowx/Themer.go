package windowx

import (
	"context"
	"fmt"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/statedata"
	"github.com/markbates/wailsx/wailsrun"
)

var _ ThemeManagerDataProvider = &Themer{}

func NopThemer() *Themer {
	return &Themer{
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

func (th *Themer) WindowSetDarkTheme(ctx context.Context) error {
	if th == nil {
		return wailsrun.WindowSetDarkTheme(ctx)
	}

	return safe.Run(func() error {
		fn := th.WindowSetDarkThemeFn
		if fn == nil {
			fn = wailsrun.WindowSetDarkTheme
		}

		if err := fn(ctx); err != nil {
			return err
		}

		return th.data.SetDarkTheme()
	})
}

func (th *Themer) WindowSetLightTheme(ctx context.Context) error {
	if th == nil {
		return wailsrun.WindowSetLightTheme(ctx)
	}

	return safe.Run(func() error {
		if th.WindowSetLightThemeFn == nil {
			return wailsrun.WindowSetLightTheme(ctx)
		}

		err := th.WindowSetLightThemeFn(ctx)
		if err != nil {
			return err
		}

		return th.data.SetLightTheme()
	})
}

func (th *Themer) WindowSetSystemDefaultTheme(ctx context.Context) error {
	if th == nil {
		return wailsrun.WindowSetSystemDefaultTheme(ctx)
	}

	return safe.Run(func() error {
		if th.WindowSetSystemDefaultThemeFn == nil {
			return wailsrun.WindowSetSystemDefaultTheme(ctx)
		}

		err := th.WindowSetSystemDefaultThemeFn(ctx)
		if err != nil {
			return err
		}

		return th.data.SetSystemTheme()
	})
}

func (th *Themer) WindowSetBackgroundColour(ctx context.Context, R uint8, G uint8, B uint8, A uint8) error {
	if th == nil {
		return wailsrun.WindowSetBackgroundColour(ctx, R, G, B, A)
	}

	return safe.Run(func() error {
		if th.WindowSetBackgroundColourFn == nil {
			return wailsrun.WindowSetBackgroundColour(ctx, R, G, B, A)
		}

		err := th.WindowSetBackgroundColourFn(ctx, R, G, B, A)
		if err != nil {
			return err
		}

		return th.data.SetBackgroundColour(R, G, B, A)
	})
}

func (th *Themer) StateData(ctx context.Context) (statedata.Data[*ThemeData], error) {
	if th == nil {
		return statedata.Data[*ThemeData]{}, fmt.Errorf("themer is nil")
	}
	return th.data.StateData(ctx)
}
