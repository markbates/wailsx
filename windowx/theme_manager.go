package windowx

import (
	"context"

	"github.com/markbates/wailsx/statedata"
)

type ThemeManager interface {
	WindowSetBackgroundColour(ctx context.Context, R, G, B, A uint8) error
	WindowSetDarkTheme(ctx context.Context) error
	WindowSetLightTheme(ctx context.Context) error
	WindowSetSystemDefaultTheme(ctx context.Context) error
}

type ThemeManagerDataProvider interface {
	ThemeManager
	StateData(ctx context.Context) (statedata.Data[*ThemeData], error)
}
