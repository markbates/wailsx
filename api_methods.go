package wailsx

import (
	"context"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

func (api *API) ClipboardGetText(ctx context.Context) (string, error) {
	if api == nil {
		return wailsrun.ClipboardGetText(ctx)
	}

	return api.ClipboardManager.ClipboardGetText(ctx)
}

func (api *API) ClipboardSetText(ctx context.Context, text string) error {
	if api == nil {
		return wailsrun.ClipboardSetText(ctx, text)
	}

	return api.ClipboardManager.ClipboardSetText(ctx, text)
}

func (api *API) EventsEmit(ctx context.Context, event string, data ...any) error {
	if api == nil {
		return wailsrun.EventsEmit(ctx, event, data...)
	}

	return api.EventManager.EventsEmit(ctx, event, data...)
}

func (api *API) EventsOff(ctx context.Context, event string, additional ...string) error {
	if api == nil {
		return wailsrun.EventsOff(ctx, event, additional...)
	}

	return api.EventManager.EventsOff(ctx, event, additional...)
}

func (api *API) EventsOffAll(ctx context.Context) error {
	if api == nil {
		return wailsrun.EventsOffAll(ctx)
	}

	return api.EventManager.EventsOffAll(ctx)
}

func (api *API) EventsOn(ctx context.Context, eventName string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error) {
	if api == nil {
		return wailsrun.EventsOn(ctx, eventName, callback)
	}

	return api.EventManager.EventsOn(ctx, eventName, callback)
}

func (api *API) EventsOnMultiple(ctx context.Context, eventName string, callback wailsrun.CallbackFn, counter int) (wailsrun.CancelFn, error) {
	if api == nil {
		return wailsrun.EventsOnMultiple(ctx, eventName, callback, counter)
	}

	return api.EventManager.EventsOnMultiple(ctx, eventName, callback, counter)
}

func (api *API) EventsOnce(ctx context.Context, eventName string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error) {
	if api == nil {
		return wailsrun.EventsOnce(ctx, eventName, callback)
	}

	return api.EventManager.EventsOnce(ctx, eventName, callback)
}

func (api *API) Hide(ctx context.Context) error {
	if api == nil {
		return wailsrun.Hide(ctx)
	}

	return api.WindowManager.Hide(ctx)
}

func (api *API) LogDebug(ctx context.Context, message string) error {
	if api == nil {
		return wailsrun.LogDebug(ctx, message)
	}

	return api.WailsLogger.LogDebug(ctx, message)
}

func (api *API) LogDebugf(ctx context.Context, format string, args ...any) error {
	if api == nil {
		return wailsrun.LogDebugf(ctx, format, args...)
	}

	return api.WailsLogger.LogDebugf(ctx, format, args...)
}

func (api *API) LogError(ctx context.Context, message string) error {
	if api == nil {
		return wailsrun.LogError(ctx, message)
	}

	return api.WailsLogger.LogError(ctx, message)
}

func (api *API) LogErrorf(ctx context.Context, format string, args ...any) error {
	if api == nil {
		return wailsrun.LogErrorf(ctx, format, args...)
	}

	return api.WailsLogger.LogErrorf(ctx, format, args...)
}

func (api *API) LogFatal(ctx context.Context, message string) error {
	if api == nil {
		return wailsrun.LogFatal(ctx, message)
	}

	return api.WailsLogger.LogFatal(ctx, message)
}

func (api *API) LogFatalf(ctx context.Context, format string, args ...any) error {
	if api == nil {
		return wailsrun.LogFatalf(ctx, format, args...)
	}

	return api.WailsLogger.LogFatalf(ctx, format, args...)
}

func (api *API) LogInfo(ctx context.Context, message string) error {
	if api == nil {
		return wailsrun.LogInfo(ctx, message)
	}

	return api.WailsLogger.LogInfo(ctx, message)
}

func (api *API) LogInfof(ctx context.Context, format string, args ...any) error {
	if api == nil {
		return wailsrun.LogInfof(ctx, format, args...)
	}

	return api.WailsLogger.LogInfof(ctx, format, args...)
}

func (api *API) LogPrint(ctx context.Context, message string) error {
	if api == nil {
		return wailsrun.LogPrint(ctx, message)
	}

	return api.WailsLogger.LogPrint(ctx, message)
}

func (api *API) LogPrintf(ctx context.Context, format string, args ...any) error {
	if api == nil {
		return wailsrun.LogPrintf(ctx, format, args...)
	}

	return api.WailsLogger.LogPrintf(ctx, format, args...)
}

func (api *API) LogSetLogLevel(ctx context.Context, level logger.LogLevel) error {
	if api == nil {
		return wailsrun.LogSetLogLevel(ctx, level)
	}

	return api.WailsLogger.LogSetLogLevel(ctx, level)
}

func (api *API) LogTrace(ctx context.Context, message string) error {
	if api == nil {
		return wailsrun.LogTrace(ctx, message)
	}

	return api.WailsLogger.LogTrace(ctx, message)
}

func (api *API) LogTracef(ctx context.Context, format string, args ...any) error {
	if api == nil {
		return wailsrun.LogTracef(ctx, format, args...)
	}

	return api.WailsLogger.LogTracef(ctx, format, args...)
}

func (api *API) LogWarning(ctx context.Context, message string) error {
	if api == nil {
		return wailsrun.LogWarning(ctx, message)
	}

	return api.WailsLogger.LogWarning(ctx, message)
}

func (api *API) LogWarningf(ctx context.Context, format string, args ...any) error {
	if api == nil {
		return wailsrun.LogWarningf(ctx, format, args...)
	}

	return api.WailsLogger.LogWarningf(ctx, format, args...)
}

func (api *API) MenuSetApplicationMenu(ctx context.Context, menu *menu.Menu) error {
	if api == nil {
		return wailsrun.MenuSetApplicationMenu(ctx, menu)
	}

	return api.MenuManager.MenuSetApplicationMenu(ctx, menu)
}

func (api *API) MenuUpdateApplicationMenu(ctx context.Context) error {
	if api == nil {
		return wailsrun.MenuUpdateApplicationMenu(ctx)
	}

	return api.MenuManager.MenuUpdateApplicationMenu(ctx)
}

func (api *API) MessageDialog(ctx context.Context, opts wailsrun.MessageDialogOptions) (string, error) {
	if api == nil {
		return wailsrun.MessageDialog(ctx, opts)
	}

	return api.DialogManager.MessageDialog(ctx, opts)
}

func (api *API) OpenDirectoryDialog(ctx context.Context, opts wailsrun.OpenDialogOptions) (string, error) {
	if api == nil {
		return wailsrun.OpenDirectoryDialog(ctx, opts)
	}

	return api.DialogManager.OpenDirectoryDialog(ctx, opts)
}

func (api *API) OpenFileDialog(ctx context.Context, opts wailsrun.OpenDialogOptions) (string, error) {
	if api == nil {
		return wailsrun.OpenFileDialog(ctx, opts)
	}

	return api.DialogManager.OpenFileDialog(ctx, opts)
}

func (api *API) OpenMultipleFilesDialog(ctx context.Context, opts wailsrun.OpenDialogOptions) ([]string, error) {
	if api == nil {
		return wailsrun.OpenMultipleFilesDialog(ctx, opts)
	}

	return api.DialogManager.OpenMultipleFilesDialog(ctx, opts)
}

func (api *API) SaveFileDialog(ctx context.Context, opts wailsrun.SaveDialogOptions) (string, error) {
	if api == nil {
		return wailsrun.SaveFileDialog(ctx, opts)
	}

	return api.DialogManager.SaveFileDialog(ctx, opts)
}

func (api *API) Show(ctx context.Context) error {
	if api == nil {
		return wailsrun.Show(ctx)
	}

	return api.WindowManager.Show(ctx)
}

func (api *API) WindowCenter(ctx context.Context) error {
	if api == nil {
		return wailsrun.WindowCenter(ctx)
	}

	return api.WindowManager.WindowCenter(ctx)
}

func (api *API) WindowExecJS(ctx context.Context, js string) error {
	if api == nil {
		return wailsrun.WindowExecJS(ctx, js)
	}

	return api.WindowManager.WindowExecJS(ctx, js)
}

func (api *API) WindowFullscreen(ctx context.Context) error {
	if api == nil {
		return wailsrun.WindowFullscreen(ctx)
	}

	return api.WindowManager.WindowFullscreen(ctx)
}

func (api *API) WindowGetPosition(ctx context.Context) (int, int, error) {
	if api == nil {
		return wailsrun.WindowGetPosition(ctx)
	}

	return api.WindowManager.WindowGetPosition(ctx)
}

func (api *API) WindowGetSize(ctx context.Context) (int, int, error) {
	if api == nil {
		return wailsrun.WindowGetSize(ctx)
	}

	return api.WindowManager.WindowGetSize(ctx)
}

func (api *API) WindowHide(ctx context.Context) error {
	if api == nil {
		return wailsrun.WindowHide(ctx)
	}

	return api.WindowManager.WindowHide(ctx)
}

func (api *API) WindowIsFullscreen(ctx context.Context) (bool, error) {
	if api == nil {
		return wailsrun.WindowIsFullscreen(ctx)
	}

	return api.WindowManager.WindowIsFullscreen(ctx)
}

func (api *API) WindowIsMaximised(ctx context.Context) (bool, error) {
	if api == nil {
		return wailsrun.WindowIsMaximised(ctx)
	}

	return api.WindowManager.WindowIsMaximised(ctx)
}

func (api *API) WindowIsMinimised(ctx context.Context) (bool, error) {
	if api == nil {
		return wailsrun.WindowIsMinimised(ctx)
	}

	return api.WindowManager.WindowIsMinimised(ctx)
}

func (api *API) WindowIsNormal(ctx context.Context) (bool, error) {
	if api == nil {
		return wailsrun.WindowIsNormal(ctx)
	}

	return api.WindowManager.WindowIsNormal(ctx)
}

func (api *API) WindowMaximise(ctx context.Context) error {
	if api == nil {
		return wailsrun.WindowMaximise(ctx)
	}

	return api.WindowManager.WindowMaximise(ctx)
}

func (api *API) WindowMinimise(ctx context.Context) error {
	if api == nil {
		return wailsrun.WindowMinimise(ctx)
	}

	return api.WindowManager.WindowMinimise(ctx)
}

func (api *API) WindowPrint(ctx context.Context) error {
	if api == nil {
		return wailsrun.WindowPrint(ctx)
	}

	return api.WindowManager.WindowPrint(ctx)
}

func (api *API) WindowReload(ctx context.Context) error {
	if api == nil {
		return wailsrun.WindowReload(ctx)
	}

	return api.WindowManager.WindowReload(ctx)
}

func (api *API) WindowReloadApp(ctx context.Context) error {
	if api == nil {
		return wailsrun.WindowReloadApp(ctx)
	}

	return api.WindowManager.WindowReloadApp(ctx)
}

func (api *API) WindowSetAlwaysOnTop(ctx context.Context, b bool) error {
	if api == nil {
		return wailsrun.WindowSetAlwaysOnTop(ctx, b)
	}

	return api.WindowManager.WindowSetAlwaysOnTop(ctx, b)
}

func (api *API) WindowSetBackgroundColour(ctx context.Context, R, G, B, A uint8) error {
	if api == nil {
		return wailsrun.WindowSetBackgroundColour(ctx, R, G, B, A)
	}

	return api.WindowManager.WindowSetBackgroundColour(ctx, R, G, B, A)
}

func (api *API) WindowSetDarkTheme(ctx context.Context) error {
	if api == nil {
		return wailsrun.WindowSetDarkTheme(ctx)
	}

	return api.WindowManager.WindowSetDarkTheme(ctx)
}

func (api *API) WindowSetLightTheme(ctx context.Context) error {
	if api == nil {
		return wailsrun.WindowSetLightTheme(ctx)
	}

	return api.WindowManager.WindowSetLightTheme(ctx)
}

func (api *API) WindowSetMaxSize(ctx context.Context, width int, height int) error {
	if api == nil {
		return wailsrun.WindowSetMaxSize(ctx, width, height)
	}

	return api.WindowManager.WindowSetMaxSize(ctx, width, height)
}

func (api *API) WindowSetMinSize(ctx context.Context, width int, height int) error {
	if api == nil {
		return wailsrun.WindowSetMinSize(ctx, width, height)
	}

	return api.WindowManager.WindowSetMinSize(ctx, width, height)
}

func (api *API) WindowSetPosition(ctx context.Context, x int, y int) error {
	if api == nil {
		return wailsrun.WindowSetPosition(ctx, x, y)
	}

	return api.WindowManager.WindowSetPosition(ctx, x, y)
}

func (api *API) WindowSetSize(ctx context.Context, width int, height int) error {
	if api == nil {
		return wailsrun.WindowSetSize(ctx, width, height)
	}

	return api.WindowManager.WindowSetSize(ctx, width, height)
}

func (api *API) WindowSetSystemDefaultTheme(ctx context.Context) error {
	if api == nil {
		return wailsrun.WindowSetSystemDefaultTheme(ctx)
	}

	return api.WindowManager.WindowSetSystemDefaultTheme(ctx)
}

func (api *API) WindowSetTitle(ctx context.Context, title string) error {
	if api == nil {
		return wailsrun.WindowSetTitle(ctx, title)
	}

	return api.WindowManager.WindowSetTitle(ctx, title)
}

func (api *API) WindowShow(ctx context.Context) error {
	if api == nil {
		return wailsrun.WindowShow(ctx)
	}

	return api.WindowManager.WindowShow(ctx)
}

func (api *API) WindowToggleMaximise(ctx context.Context) error {
	if api == nil {
		return wailsrun.WindowToggleMaximise(ctx)
	}

	return api.WindowManager.WindowToggleMaximise(ctx)
}

func (api *API) WindowUnfullscreen(ctx context.Context) error {
	if api == nil {
		return wailsrun.WindowUnfullscreen(ctx)
	}

	return api.WindowManager.WindowUnfullscreen(ctx)
}

func (api *API) WindowUnmaximise(ctx context.Context) error {
	if api == nil {
		return wailsrun.WindowUnmaximise(ctx)
	}

	return api.WindowManager.WindowUnmaximise(ctx)
}

func (api *API) WindowUnminimise(ctx context.Context) error {
	if api == nil {
		return wailsrun.WindowUnminimise(ctx)
	}

	return api.WindowManager.WindowUnminimise(ctx)
}

func (api *API) ScreenGetAll(ctx context.Context) ([]wailsrun.Screen, error) {
	if api == nil {
		return wailsrun.ScreenGetAll(ctx)
	}

	return api.WindowManager.ScreenGetAll(ctx)
}
