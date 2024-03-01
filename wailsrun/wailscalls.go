//go:build !dev && !desktop && !production && !wails

// when not built with wails, the stubs are used
package wailsrun

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

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

func WindowIsMaximised(ctx context.Context) (bool, error) {
	return false, ErrNotAvailable
}

func WindowIsMinimised(ctx context.Context) (bool, error) {
	return false, ErrNotAvailable
}

func WindowIsNormal(ctx context.Context) (bool, error) {
	return false, ErrNotAvailable
}

func WindowMaximise(ctx context.Context) error {
	return ErrNotAvailable
}

func WindowMinimise(ctx context.Context) error {
	return ErrNotAvailable
}

func WindowPrint(ctx context.Context) error {
	return ErrNotAvailable
}

func WindowReload(ctx context.Context) error {
	return ErrNotAvailable
}

func WindowReloadApp(ctx context.Context) error {
	return ErrNotAvailable
}

func WindowSetAlwaysOnTop(ctx context.Context, b bool) error {
	return ErrNotAvailable
}

func WindowSetBackgroundColour(ctx context.Context, R, G, B, A uint8) error {
	return ErrNotAvailable
}

func WindowSetDarkTheme(ctx context.Context) error {
	return ErrNotAvailable
}

func WindowSetLightTheme(ctx context.Context) error {
	return ErrNotAvailable
}

func WindowSetMaxSize(ctx context.Context, width int, height int) error {
	return ErrNotAvailable
}

func WindowSetMinSize(ctx context.Context, width int, height int) error {
	return ErrNotAvailable
}

func WindowSetPosition(ctx context.Context, x int, y int) error {
	return ErrNotAvailable
}

func WindowSetSize(ctx context.Context, width int, height int) error {
	return ErrNotAvailable
}

func WindowSetSystemDefaultTheme(ctx context.Context) error {
	return ErrNotAvailable
}

func WindowSetTitle(ctx context.Context, title string) error {
	return ErrNotAvailable
}

func WindowShow(ctx context.Context) error {
	return ErrNotAvailable
}

func WindowToggleMaximise(ctx context.Context) error {
	return ErrNotAvailable
}

func WindowUnfullscreen(ctx context.Context) error {
	return ErrNotAvailable
}

func WindowUnmaximise(ctx context.Context) error {
	return ErrNotAvailable
}

func WindowUnminimise(ctx context.Context) error {
	return ErrNotAvailable
}

func ScreenGetAll(ctx context.Context) ([]Screen, error) {
	return nil, ErrNotAvailable
}
