//go:build dev || desktop || production

// when built with wails the real api
// is used and the stubs are not
package wailsrun

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/menu"
	wailsrun "github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm/logger"
)

func BrowserOpenURL(ctx context.Context, url string) error {
	wailsrun.BrowserOpenURL(ctx, url)
	return nil
}

func ClipboardGetText(ctx context.Context) (string, error) {
	return wailsrun.ClipboardGetText(ctx)
}

func ClipboardSetText(ctx context.Context, text string) error {
	return wailsrun.ClipboardSetText(ctx, text)
}

func EventsEmit(ctx context.Context, event string, data ...any) error {
	wailsrun.EventsEmit(ctx, event, data...)
	return nil
}

func EventsOff(ctx context.Context, event string, additional ...string) error {
	wailsrun.EventsOff(ctx, event, additional...)
	return nil
}

func EventsOffAll(ctx context.Context) error {
	wailsrun.EventsOffAll(ctx)
	return nil
}

func EventsOn(ctx context.Context, eventName string, callback CallbackFn) (CancelFn, error) {
	fn := wailsrun.EventsOn(ctx, event, callback)
	return func() error {
		fn()
		return nil
	}
}

func EventsOnMultiple(ctx context.Context, eventName string, callback CallbackFn, counter int) (CancelFn, error) {
	fn := wailsrun.EventsOnMultiple(ctx, event, data...)
	return func() error {
		fn()
		return nil
	}
}

func EventsOnce(ctx context.Context, eventName string, callback CallbackFn) (CancelFn, error) {
	fn := wailsrun.EventsOnce(ctx, event, data)
	return func() error {
		fn()
		return nil
	}
}

func Hide(ctx context.Context) error {
	wailsrun.Hide(ctx)
	return nil
}

func LogDebug(ctx context.Context, message string) error {
	wailsrun.LogDebug(ctx, message)
	return nil
}

func LogDebugf(ctx context.Context, format string, args ...any) error {
	wailsrun.LogDebugf(ctx, format, args...)
	return nil
}

func LogError(ctx context.Context, message string) error {
	wailsrun.LogError(ctx, message)
	return nil
}

func LogErrorf(ctx context.Context, format string, args ...any) error {
	wailsrun.LogErrorf(ctx, format, args...)
	return nil
}

func LogFatal(ctx context.Context, message string) error {
	wailsrun.LogFatal(ctx, message)
	return nil
}

func LogFatalf(ctx context.Context, format string, args ...any) error {
	wailsrun.LogFatalf(ctx, format, args...)
	return nil
}

func LogInfo(ctx context.Context, message string) error {
	wailsrun.LogInfo(ctx, message)
	return nil
}

func LogInfof(ctx context.Context, format string, args ...any) error {
	wailsrun.LogInfof(ctx, format, args...)
	return nil
}

func LogPrint(ctx context.Context, message string) error {
	wailsrun.LogPrint(ctx, message)
	return nil
}

func LogPrintf(ctx context.Context, format string, args ...any) error {
	wailsrun.LogPrintf(ctx, format, args...)
	return nil
}

func LogSetLogLevel(ctx context.Context, level logger.LogLevel) error {
	wailsrun.LogSetLogLevel(ctx, level)
	return nil
}

func LogTrace(ctx context.Context, message string) error {
	wailsrun.LogTrace(ctx, message)
	return nil
}

func LogTracef(ctx context.Context, format string, args ...any) error {
	wailsrun.LogTracef(ctx, format, args...)
	return nil
}

func LogWarning(ctx context.Context, message string) error {
	wailsrun.LogWarning(ctx, message)
	return nil
}

func LogWarningf(ctx context.Context, format string, args ...any) error {
	wailsrun.LogWarningf(ctx, format, args...)
	return nil
}

func MenuSetApplicationMenu(ctx context.Context, menu *menu.Menu) error {
	wailsrun.MenuSetApplicationMenu(ctx, menu)
	return nil
}

func MenuUpdateApplicationMenu(ctx context.Context) error {
	wailsrun.MenuUpdateApplicationMenu(ctx)
	return nil
}

func MessageDialog(ctx context.Context, dialogOptions MessageDialogOptions) (string, error) {
	return wailsrun.MessageDialog(ctx, dialogOptions)
}

func OpenDirectoryDialog(ctx context.Context, dialogOptions OpenDialogOptions) (string, error) {
	return wailsrun.OpenDirectoryDialog(ctx, dialogOptions)
}

func OpenFileDialog(ctx context.Context, dialogOptions OpenDialogOptions) (string, error) {
	return wailsrun.OpenFileDialog(ctx, dialogOptions)
}

func OpenMultipleFilesDialog(ctx context.Context, dialogOptions OpenDialogOptions) ([]string, error) {
	return wailsrun.OpenMultipleFilesDialog(ctx, dialogOptions)
}

func Quit(ctx context.Context) error {
	wailsrun.Quit(ctx)
	return nil
}

func SaveFileDialog(ctx context.Context, dialogOptions SaveDialogOptions) (string, error) {
	return wailsrun.SaveFileDialog(ctx, dialogOptions)
}

func Show(ctx context.Context) error {
	wailsrun.Show(ctx)
	return nil
}

func WindowCenter(ctx context.Context) error {
	wailsrun.WindowCenter(ctx)
	return nil
}

func WindowExecJS(ctx context.Context, js string) error {
	wailsrun.WindowExecJS(ctx, js)
	return nil
}

func WindowFullscreen(ctx context.Context) error {
	wailsrun.WindowFullscreen(ctx)
	return nil
}

func WindowGetPosition(ctx context.Context) (int, int, error) {
	x, y := wailsrun.WindowGetPosition(ctx)
	return x, y, nil
}

func WindowGetSize(ctx context.Context) (int, int, error) {
	w, h := wailsrun.WindowGetSize(ctx)
	return w, h, nil
}

func WindowHide(ctx context.Context) error {
	wailsrun.WindowHide(ctx)
	return nil
}

func WindowIsFullscreen(ctx context.Context) (bool, error) {
	return wailsrun.WindowIsFullscreen(ctx), nil
}

func WindowIsFullscreen(ctx context.Context) (bool, error) {
	return wailsrun.WindowIsMaximised(ctx), nil
}

func WindowIsFullscreen(ctx context.Context) (bool, error) {
	return wailsrun.WindowIsMinimised(ctx), nil
}

func WindowIsFullscreen(ctx context.Context) (bool, error) {
	return wailsrun.WindowIsNormal(ctx), nil
}

func WindowMaximise(ctx context.Context) error {
	wailsrun.WindowMaximise(ctx)
	return nil
}

func WindowMinimise(ctx context.Context) error {
	wailsrun.WindowMinimise(ctx)
	return nil
}

func WindowPrint(ctx context.Context) error {
	wailsrun.WindowPrint(ctx)
	return nil
}

func WindowReload(ctx context.Context) error {
	wailsrun.WindowReload(ctx)
	return nil
}

func WindowReloadApp(ctx context.Context) error {
	wailsrun.WindowReloadApp(ctx)
	return nil
}

func WindowSetAlwaysOnTop(ctx context.Context, b bool) error {
	wailsrun.WindowSetAlwaysOnTop(ctx, b)
	return nil
}

func WindowSetBackgroundColour(ctx context.Context, R, G, B, A uint8) error {
	wailsrun.WindowSetBackgroundColour(ctx, R, G, B, A)
	return nil
}

func WindowSetDarkTheme(ctx context.Context) error {
	wailsrun.WindowSetDarkTheme(ctx)
	return nil
}

func WindowSetLightTheme(ctx context.Context) error {
	wailsrun.WindowSetLightTheme(ctx)
	return nil
}

func WindowSetMaxSize(ctx context.Context, width int, height int) error {
	wailsrun.WindowSetMaxSize(ctx, width, height)
	return nil
}

func WindowSetMinSize(ctx context.Context, width int, height int) error {
	wailsrun.WindowSetMinSize(ctx, width, height)
	return nil
}

func WindowSetPosition(ctx context.Context, x int, y int) error {
	wailsrun.WindowSetPosition(ctx, x, y)
	return nil
}

func WindowSetSize(ctx context.Context, width int, height int) error {
	wailsrun.WindowSetSize(ctx, width, height)
	return nil
}

func WindowSetSystemDefaultTheme(ctx context.Context) error {
	wailsrun.WindowSetSystemDefaultTheme(ctx)
	return nil
}

func WindowSetTitle(ctx context.Context, title string) error {
	wailsrun.WindowSetTitle(ctx, title)
	return nil
}

func WindowShow(ctx context.Context) error {
	wailsrun.WindowShow(ctx)
	return nil
}

func WindowToggleMaximise(ctx context.Context) error {
	wailsrun.WindowToggleMaximise(ctx)
	return nil
}

func WindowUnfullscreen(ctx context.Context) error {
	wailsrun.WindowUnfullscreen(ctx)
	return nil
}

func WindowUnmaximise(ctx context.Context) error {
	wailsrun.WindowUnmaximise(ctx)
	return nil
}

func WindowUnminimise(ctx context.Context) error {
	wailsrun.WindowUnminimise(ctx)
	return nil
}
