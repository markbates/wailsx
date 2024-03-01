// when not built with wails, the stubs are used
package wailsrun

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
	wailsrun "github.com/wailsapp/wails/v2/pkg/runtime"
)

const (
	ErrNotAvailable = es("wails api calls are not available in this environment")
)

type es string

func (e es) Error() string {
	return string(e)
}

func BrowserOpenURL(ctx context.Context, url string) error {
	return ErrNotAvailable
}

func ClipboardGetText(ctx context.Context) (string, error) {
	return "", ErrNotAvailable
}

func ClipboardSetText(ctx context.Context, text string) error {
	return ErrNotAvailable
}

func EventsEmit(ctx context.Context, event string, data ...any) error {
	return ErrNotAvailable
}

func EventsOff(ctx context.Context, event string, additional ...string) error {
	return ErrNotAvailable
}

func EventsOffAll(ctx context.Context) error {
	return ErrNotAvailable
}

func EventsOn(ctx context.Context, eventName string, callback CallbackFn) (CancelFn, error) {
	return nil, ErrNotAvailable
}

func EventsOnMultiple(ctx context.Context, eventName string, callback CallbackFn, counter int) (CancelFn, error) {
	return nil, ErrNotAvailable
}

func EventsOnce(ctx context.Context, eventName string, callback CallbackFn) (CancelFn, error) {
	return nil, ErrNotAvailable
}

func Hide(ctx context.Context) error {
	return ErrNotAvailable
}

func LogDebug(ctx context.Context, message string) error {
	return ErrNotAvailable
}

func LogDebugf(ctx context.Context, format string, args ...any) error {
	return ErrNotAvailable
}

func LogError(ctx context.Context, message string) error {
	return ErrNotAvailable
}

func LogErrorf(ctx context.Context, format string, args ...any) error {
	return ErrNotAvailable
}

func LogFatal(ctx context.Context, message string) error {
	return ErrNotAvailable
}

func LogFatalf(ctx context.Context, format string, args ...any) error {
	return ErrNotAvailable
}

func LogInfo(ctx context.Context, message string) error {
	return ErrNotAvailable
}

func LogInfof(ctx context.Context, format string, args ...any) error {
	return ErrNotAvailable
}

func LogPrint(ctx context.Context, message string) error {
	return ErrNotAvailable
}

func LogPrintf(ctx context.Context, format string, args ...any) error {
	return ErrNotAvailable
}

func LogSetLogLevel(ctx context.Context, level logger.LogLevel) error {
	return ErrNotAvailable
}

func LogTrace(ctx context.Context, message string) error {
	return ErrNotAvailable
}

func LogTracef(ctx context.Context, format string, args ...any) error {
	return ErrNotAvailable
}

func LogWarning(ctx context.Context, message string) error {
	return ErrNotAvailable
}

func LogWarningf(ctx context.Context, format string, args ...any) error {
	return ErrNotAvailable
}

func MenuSetApplicationMenu(ctx context.Context, menu *menu.Menu) error {
	return ErrNotAvailable
}

func MenuUpdateApplicationMenu(ctx context.Context) error {
	return ErrNotAvailable
}

func MessageDialog(ctx context.Context, dialogOptions MessageDialogOptions) (string, error) {
	return "", ErrNotAvailable
}

func OpenDirectoryDialog(ctx context.Context, dialogOptions OpenDialogOptions) (string, error) {
	return "", ErrNotAvailable
}

func OpenFileDialog(ctx context.Context, dialogOptions OpenDialogOptions) (string, error) {
	return "", ErrNotAvailable
}

func OpenMultipleFilesDialog(ctx context.Context, dialogOptions OpenDialogOptions) ([]string, error) {
	return nil, ErrNotAvailable
}

func Quit(ctx context.Context) error {
	return ErrNotAvailable
}

func SaveFileDialog(ctx context.Context, dialogOptions SaveDialogOptions) (string, error) {
	return "", ErrNotAvailable
}

func Show(ctx context.Context) error {
	return ErrNotAvailable
}

func WindowCenter(ctx context.Context) error {
	return ErrNotAvailable
}

func WindowExecJS(ctx context.Context, js string) error {
	return ErrNotAvailable
}

func WindowFullscreen(ctx context.Context) error {
	return ErrNotAvailable
}

func WindowGetPosition(ctx context.Context) (int, int, error) {
	return 0, 0, ErrNotAvailable
}

func WindowGetSize(ctx context.Context) (int, int, error) {
	return 0, 0, ErrNotAvailable
}

func WindowHide(ctx context.Context) error {
	return ErrNotAvailable
}

func WindowIsFullscreen(ctx context.Context) (bool, error) {
	return false, ErrNotAvailable
}

func WindowIsMaximised(ctx context.Context) bool {
	return wailsrun.WindowIsMaximised(ctx)
}

func WindowIsMinimised(ctx context.Context) bool {
	return wailsrun.WindowIsMinimised(ctx)
}

func WindowIsNormal(ctx context.Context) bool {
	return wailsrun.WindowIsNormal(ctx)
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
