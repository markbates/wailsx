package windowx

import (
	"context"
	"fmt"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

var _ ThemeManagerDataProvider = &Themer{}
var _ RestorableThemeManager = &Themer{}

func NopThemer() *Themer {
	return &Themer{
		WindowSetBackgroundColourFn:   func(ctx context.Context, R, G, B, A uint8) error { return nil },
		WindowSetDarkThemeFn:          func(ctx context.Context) error { return nil },
		WindowSetLightThemeFn:         func(ctx context.Context) error { return nil },
		WindowSetSystemDefaultThemeFn: func(ctx context.Context) error { return nil },
	}
}

type Themer struct {
	WindowSetBackgroundColourFn   func(ctx context.Context, R, G, B, A uint8) error `json:"-"`
	WindowSetDarkThemeFn          func(ctx context.Context) error                   `json:"-"`
	WindowSetLightThemeFn         func(ctx context.Context) error                   `json:"-"`
	WindowSetSystemDefaultThemeFn func(ctx context.Context) error                   `json:"-"`

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

func (th *Themer) StateData(ctx context.Context) (*ThemeData, error) {
	if th == nil {
		return nil, fmt.Errorf("themer is nil")
	}

	return th.data.StateData(ctx)
}

func (th *Themer) RestoreTheme(ctx context.Context, data *ThemeData) error {
	if th == nil {
		return fmt.Errorf("themer is nil")
	}

	if data == nil {
		return fmt.Errorf("data is nil")
	}

	bc := data.BackgroundColour

	err := th.WindowSetBackgroundColourFn(ctx, bc.R, bc.G, bc.B, bc.A)
	if err != nil {
		return err
	}

	switch data.Theme {
	case THEME_SYSTEM:
		return th.WindowSetSystemDefaultTheme(ctx)
	case THEME_DARK:
		return th.WindowSetDarkTheme(ctx)
	case THEME_LIGHT:
		return th.WindowSetLightTheme(ctx)
	}

	return fmt.Errorf("unknown theme: %s", data.Theme)
}
