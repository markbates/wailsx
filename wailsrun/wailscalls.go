//go:build !dev && !desktop && !production && !wails

// when not built with wails, the stubs are used
package wailsrun

// snippet: BrowserOpenURL
import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

func BrowserOpenURL(ctx context.Context, url string) error {
	return ErrNotAvailable("BrowserOpenURL")
}

// snippet: BrowserOpenURL

func ClipboardGetText(ctx context.Context) (string, error) {
	return "", ErrNotAvailable("ClipboardGetText")
}

func ClipboardSetText(ctx context.Context, text string) error {
	return ErrNotAvailable("ClipboardSetText")
}

func EventsEmit(ctx context.Context, event string, data ...any) error {
	return ErrNotAvailable("EventsEmit")
}

func EventsOff(ctx context.Context, event string, additional ...string) error {
	return ErrNotAvailable("EventsOff")
}

func EventsOffAll(ctx context.Context) error {
	return ErrNotAvailable("EventsOffAll")
}

func EventsOn(ctx context.Context, event string, callback CallbackFn) (CancelFn, error) {
	return nil, ErrNotAvailable("EventsOn")
}

func EventsOnMultiple(ctx context.Context, event string, callback CallbackFn, counter int) (CancelFn, error) {
	return nil, ErrNotAvailable("EventsOnMultiple")
}

func EventsOnce(ctx context.Context, event string, callback CallbackFn) (CancelFn, error) {
	return nil, ErrNotAvailable("EventsOnce")
}

func Hide(ctx context.Context) error {
	return ErrNotAvailable("Hide")
}

func LogDebug(ctx context.Context, message string) error {
	return ErrNotAvailable("LogDebug")
}

func LogDebugf(ctx context.Context, format string, args ...any) error {
	return ErrNotAvailable("LogDebugf")
}

func LogError(ctx context.Context, message string) error {
	return ErrNotAvailable("LogError")
}

func LogErrorf(ctx context.Context, format string, args ...any) error {
	return ErrNotAvailable("LogErrorf")
}

func LogFatal(ctx context.Context, message string) error {
	return ErrNotAvailable("LogFatal")
}

func LogFatalf(ctx context.Context, format string, args ...any) error {
	return ErrNotAvailable("LogFatalf")
}

func LogInfo(ctx context.Context, message string) error {
	return ErrNotAvailable("LogInfo")
}

func LogInfof(ctx context.Context, format string, args ...any) error {
	return ErrNotAvailable("LogInfof")
}

func LogPrint(ctx context.Context, message string) error {
	return ErrNotAvailable("LogPrint")
}

func LogPrintf(ctx context.Context, format string, args ...any) error {
	return ErrNotAvailable("LogPrintf")
}

func LogSetLogLevel(ctx context.Context, level logger.LogLevel) error {
	return ErrNotAvailable("LogSetLogLevel")
}

func LogTrace(ctx context.Context, message string) error {
	return ErrNotAvailable("LogTrace")
}

func LogTracef(ctx context.Context, format string, args ...any) error {
	return ErrNotAvailable("LogTracef")
}

func LogWarning(ctx context.Context, message string) error {
	return ErrNotAvailable("LogWarning")
}

func LogWarningf(ctx context.Context, format string, args ...any) error {
	return ErrNotAvailable("LogWarningf")
}

func MenuSetApplicationMenu(ctx context.Context, menu *menu.Menu) error {
	return ErrNotAvailable("MenuSetApplicationMenu")
}

func MenuUpdateApplicationMenu(ctx context.Context) error {
	return ErrNotAvailable("MenuUpdateApplicationMenu")
}

func MessageDialog(ctx context.Context, dialogOptions MessageDialogOptions) (string, error) {
	return "", ErrNotAvailable("MessageDialog")
}

func OpenDirectoryDialog(ctx context.Context, dialogOptions OpenDialogOptions) (string, error) {
	return "", ErrNotAvailable("OpenDirectoryDialog")
}

func OpenFileDialog(ctx context.Context, dialogOptions OpenDialogOptions) (string, error) {
	return "", ErrNotAvailable("OpenFileDialog")
}

func OpenMultipleFilesDialog(ctx context.Context, dialogOptions OpenDialogOptions) ([]string, error) {
	return nil, ErrNotAvailable("OpenMultipleFilesDialog")
}

func Quit(ctx context.Context) error {
	return ErrNotAvailable("Quit")
}

func SaveFileDialog(ctx context.Context, dialogOptions SaveDialogOptions) (string, error) {
	return "", ErrNotAvailable("SaveFileDialog")
}

func Show(ctx context.Context) error {
	return ErrNotAvailable("Show")
}

func WindowCenter(ctx context.Context) error {
	return ErrNotAvailable("WindowCenter")
}

func WindowExecJS(ctx context.Context, js string) error {
	return ErrNotAvailable("WindowExecJS")
}

func WindowFullscreen(ctx context.Context) error {
	return ErrNotAvailable("WindowFullscreen")
}

func WindowGetPosition(ctx context.Context) (int, int, error) {
	return 0, 0, ErrNotAvailable("WindowGetPosition")
}

func WindowGetSize(ctx context.Context) (int, int, error) {
	return 0, 0, ErrNotAvailable("WindowGetSize")
}

func WindowHide(ctx context.Context) error {
	return ErrNotAvailable("WindowHide")
}

func WindowIsFullscreen(ctx context.Context) (bool, error) {
	return false, ErrNotAvailable("WindowIsFullscreen")
}

func WindowIsMaximised(ctx context.Context) (bool, error) {
	return false, ErrNotAvailable("WindowIsMaximised")
}

func WindowIsMinimised(ctx context.Context) (bool, error) {
	return false, ErrNotAvailable("WindowIsMinimised")
}

func WindowIsNormal(ctx context.Context) (bool, error) {
	return false, ErrNotAvailable("WindowIsNormal")
}

func WindowMaximise(ctx context.Context) error {
	return ErrNotAvailable("WindowMaximise")
}

func WindowMinimise(ctx context.Context) error {
	return ErrNotAvailable("WindowMinimise")
}

func WindowPrint(ctx context.Context) error {
	return ErrNotAvailable("WindowPrint")
}

func WindowReload(ctx context.Context) error {
	return ErrNotAvailable("WindowReload")
}

func WindowReloadApp(ctx context.Context) error {
	return ErrNotAvailable("WindowReloadApp")
}

func WindowSetAlwaysOnTop(ctx context.Context, b bool) error {
	return ErrNotAvailable("WindowSetAlwaysOnTop")
}

func WindowSetBackgroundColour(ctx context.Context, R, G, B, A uint8) error {
	return ErrNotAvailable("WindowSetBackgroundColour")
}

func WindowSetDarkTheme(ctx context.Context) error {
	return ErrNotAvailable("WindowSetDarkTheme")
}

func WindowSetLightTheme(ctx context.Context) error {
	return ErrNotAvailable("WindowSetLightTheme")
}

func WindowSetMaxSize(ctx context.Context, width int, height int) error {
	return ErrNotAvailable("WindowSetMaxSize")
}

func WindowSetMinSize(ctx context.Context, width int, height int) error {
	return ErrNotAvailable("WindowSetMinSize")
}

func WindowSetPosition(ctx context.Context, x int, y int) error {
	return ErrNotAvailable("WindowSetPosition")
}

func WindowSetSize(ctx context.Context, width int, height int) error {
	return ErrNotAvailable("WindowSetSize")
}

func WindowSetSystemDefaultTheme(ctx context.Context) error {
	return ErrNotAvailable("WindowSetSystemDefaultTheme")
}

func WindowSetTitle(ctx context.Context, title string) error {
	return ErrNotAvailable("WindowSetTitle")
}

func WindowShow(ctx context.Context) error {
	return ErrNotAvailable("WindowShow")
}

func WindowToggleMaximise(ctx context.Context) error {
	return ErrNotAvailable("WindowToggleMaximise")
}

func WindowUnfullscreen(ctx context.Context) error {
	return ErrNotAvailable("WindowUnfullscreen")
}

func WindowUnmaximise(ctx context.Context) error {
	return ErrNotAvailable("WindowUnmaximise")
}

func WindowUnminimise(ctx context.Context) error {
	return ErrNotAvailable("WindowUnminimise")
}

func ScreenGetAll(ctx context.Context) ([]Screen, error) {
	return nil, ErrNotAvailable("ScreenGetAll")
}
