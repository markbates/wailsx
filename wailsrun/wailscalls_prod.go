//go:build wails || dev || desktop || production

// when built with wails the real api
// is used and the stubs are not
package wailsrun

// snippet: BrowserOpenURL
import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func BrowserOpenURL(ctx context.Context, url string) error {
	runtime.BrowserOpenURL(ctx, url)
	return nil
}

// snippet: BrowserOpenURL

func ClipboardGetText(ctx context.Context) (string, error) {
	return runtime.ClipboardGetText(ctx)
}

func ClipboardSetText(ctx context.Context, text string) error {
	return runtime.ClipboardSetText(ctx, text)
}

func EventsEmit(ctx context.Context, event string, data ...any) error {
	runtime.EventsEmit(ctx, event, data...)
	return nil
}

func EventsOff(ctx context.Context, event string, additional ...string) error {
	runtime.EventsOff(ctx, event, additional...)
	return nil
}

func EventsOffAll(ctx context.Context) error {
	runtime.EventsOffAll(ctx)
	return nil
}

func EventsOn(ctx context.Context, event string, callback CallbackFn) (CancelFn, error) {
	cb := func(data ...any) {
		_ = callback(data...)
	}

	fn := runtime.EventsOn(ctx, event, cb)

	return func() error {
		fn()
		return nil
	}, nil
}

func EventsOnMultiple(ctx context.Context, event string, callback CallbackFn, counter int) (CancelFn, error) {
	cb := func(data ...any) {
		_ = callback(data...)
	}

	fn := runtime.EventsOnMultiple(ctx, event, cb, counter)

	return func() error {
		fn()
		return nil
	}, nil
}

func EventsOnce(ctx context.Context, event string, callback CallbackFn) (CancelFn, error) {
	cb := func(data ...any) {
		_ = callback(data...)
	}

	fn := runtime.EventsOnce(ctx, event, cb)
	return func() error {
		fn()
		return nil
	}, nil
}

func Hide(ctx context.Context) error {
	runtime.Hide(ctx)
	return nil
}

func LogDebug(ctx context.Context, message string) error {
	runtime.LogDebug(ctx, message)
	return nil
}

func LogDebugf(ctx context.Context, format string, args ...any) error {
	runtime.LogDebugf(ctx, format, args...)
	return nil
}

func LogError(ctx context.Context, message string) error {
	runtime.LogError(ctx, message)
	return nil
}

func LogErrorf(ctx context.Context, format string, args ...any) error {
	runtime.LogErrorf(ctx, format, args...)
	return nil
}

func LogFatal(ctx context.Context, message string) error {
	runtime.LogFatal(ctx, message)
	return nil
}

func LogFatalf(ctx context.Context, format string, args ...any) error {
	runtime.LogFatalf(ctx, format, args...)
	return nil
}

func LogInfo(ctx context.Context, message string) error {
	runtime.LogInfo(ctx, message)
	return nil
}

func LogInfof(ctx context.Context, format string, args ...any) error {
	runtime.LogInfof(ctx, format, args...)
	return nil
}

func LogPrint(ctx context.Context, message string) error {
	runtime.LogPrint(ctx, message)
	return nil
}

func LogPrintf(ctx context.Context, format string, args ...any) error {
	runtime.LogPrintf(ctx, format, args...)
	return nil
}

func LogSetLogLevel(ctx context.Context, level logger.LogLevel) error {
	runtime.LogSetLogLevel(ctx, level)
	return nil
}

func LogTrace(ctx context.Context, message string) error {
	runtime.LogTrace(ctx, message)
	return nil
}

func LogTracef(ctx context.Context, format string, args ...any) error {
	runtime.LogTracef(ctx, format, args...)
	return nil
}

func LogWarning(ctx context.Context, message string) error {
	runtime.LogWarning(ctx, message)
	return nil
}

func LogWarningf(ctx context.Context, format string, args ...any) error {
	runtime.LogWarningf(ctx, format, args...)
	return nil
}

func MenuSetApplicationMenu(ctx context.Context, menu *menu.Menu) error {
	runtime.MenuSetApplicationMenu(ctx, menu)
	return nil
}

func MenuUpdateApplicationMenu(ctx context.Context) error {
	runtime.MenuUpdateApplicationMenu(ctx)
	return nil
}

func MessageDialog(ctx context.Context, dialogOptions MessageDialogOptions) (string, error) {
	return runtime.MessageDialog(ctx, dialogOptions)
}

func OpenDirectoryDialog(ctx context.Context, dialogOptions OpenDialogOptions) (string, error) {
	return runtime.OpenDirectoryDialog(ctx, dialogOptions)
}

func OpenFileDialog(ctx context.Context, dialogOptions OpenDialogOptions) (string, error) {
	return runtime.OpenFileDialog(ctx, dialogOptions)
}

func OpenMultipleFilesDialog(ctx context.Context, dialogOptions OpenDialogOptions) ([]string, error) {
	return runtime.OpenMultipleFilesDialog(ctx, dialogOptions)
}

func Quit(ctx context.Context) error {
	runtime.Quit(ctx)
	return nil
}

func SaveFileDialog(ctx context.Context, dialogOptions SaveDialogOptions) (string, error) {
	return runtime.SaveFileDialog(ctx, dialogOptions)
}

func Show(ctx context.Context) error {
	runtime.Show(ctx)
	return nil
}

func WindowCenter(ctx context.Context) error {
	runtime.WindowCenter(ctx)
	return nil
}

func WindowExecJS(ctx context.Context, js string) error {
	runtime.WindowExecJS(ctx, js)
	return nil
}

func WindowFullscreen(ctx context.Context) error {
	runtime.WindowFullscreen(ctx)
	return nil
}

func WindowGetPosition(ctx context.Context) (int, int, error) {
	x, y := runtime.WindowGetPosition(ctx)
	return x, y, nil
}

func WindowGetSize(ctx context.Context) (int, int, error) {
	w, h := runtime.WindowGetSize(ctx)
	return w, h, nil
}

func WindowHide(ctx context.Context) error {
	runtime.WindowHide(ctx)
	return nil
}

func WindowIsFullscreen(ctx context.Context) (bool, error) {
	return runtime.WindowIsFullscreen(ctx), nil
}

func WindowIsMaximised(ctx context.Context) (bool, error) {
	return runtime.WindowIsMaximised(ctx), nil
}

func WindowIsMinimised(ctx context.Context) (bool, error) {
	return runtime.WindowIsMinimised(ctx), nil
}

func WindowIsNormal(ctx context.Context) (bool, error) {
	return runtime.WindowIsNormal(ctx), nil
}

func WindowMaximise(ctx context.Context) error {
	runtime.WindowMaximise(ctx)
	return nil
}

func WindowMinimise(ctx context.Context) error {
	runtime.WindowMinimise(ctx)
	return nil
}

func WindowPrint(ctx context.Context) error {
	runtime.WindowPrint(ctx)
	return nil
}

func WindowReload(ctx context.Context) error {
	runtime.WindowReload(ctx)
	return nil
}

func WindowReloadApp(ctx context.Context) error {
	runtime.WindowReloadApp(ctx)
	return nil
}

func WindowSetAlwaysOnTop(ctx context.Context, b bool) error {
	runtime.WindowSetAlwaysOnTop(ctx, b)
	return nil
}

func WindowSetBackgroundColour(ctx context.Context, R, G, B, A uint8) error {
	runtime.WindowSetBackgroundColour(ctx, R, G, B, A)
	return nil
}

func WindowSetDarkTheme(ctx context.Context) error {
	runtime.WindowSetDarkTheme(ctx)
	return nil
}

func WindowSetLightTheme(ctx context.Context) error {
	runtime.WindowSetLightTheme(ctx)
	return nil
}

func WindowSetMaxSize(ctx context.Context, width int, height int) error {
	runtime.WindowSetMaxSize(ctx, width, height)
	return nil
}

func WindowSetMinSize(ctx context.Context, width int, height int) error {
	runtime.WindowSetMinSize(ctx, width, height)
	return nil
}

func WindowSetPosition(ctx context.Context, x int, y int) error {
	runtime.WindowSetPosition(ctx, x, y)
	return nil
}

func WindowSetSize(ctx context.Context, width int, height int) error {
	runtime.WindowSetSize(ctx, width, height)
	return nil
}

func WindowSetSystemDefaultTheme(ctx context.Context) error {
	runtime.WindowSetSystemDefaultTheme(ctx)
	return nil
}

func WindowSetTitle(ctx context.Context, title string) error {
	runtime.WindowSetTitle(ctx, title)
	return nil
}

func WindowShow(ctx context.Context) error {
	runtime.WindowShow(ctx)
	return nil
}

func WindowToggleMaximise(ctx context.Context) error {
	runtime.WindowToggleMaximise(ctx)
	return nil
}

func WindowUnfullscreen(ctx context.Context) error {
	runtime.WindowUnfullscreen(ctx)
	return nil
}

func WindowUnmaximise(ctx context.Context) error {
	runtime.WindowUnmaximise(ctx)
	return nil
}

func WindowUnminimise(ctx context.Context) error {
	runtime.WindowUnminimise(ctx)
	return nil
}

func ScreenGetAll(ctx context.Context) ([]Screen, error) {
	fss, err := runtime.ScreenGetAll(ctx)
	if err != nil {
		return nil, err
	}

	screens := make([]Screen, len(fss))
	for i, fs := range fss {
		screens[i] = Screen{
			IsCurrent: fs.IsCurrent,
			IsPrimary: fs.IsPrimary,
			Size: ScreenSize{
				Width:  fs.Size.Width,
				Height: fs.Size.Height,
			},
			PhysicalSize: ScreenSize{
				Width:  fs.PhysicalSize.Width,
				Height: fs.PhysicalSize.Height,
			},
		}
	}

	return screens, nil
}
