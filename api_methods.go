package wailsx

// func (api *API) BrowserOpenURL(ctx context.Context, url string) error {
// 	if api == nil {
// 		return fmt.Errorf("api is nil")
// 	}

// }

// func (api *API) ClipboardGetText(ctx context.Context) (string, error) {
// 	return "", ErrNotAvailable("ClipboardGetText")
// }

// func (api *API) ClipboardSetText(ctx context.Context, text string) error {
// 	return ErrNotAvailable("ClipboardSetText")
// }

// func (api *API) EventsEmit(ctx context.Context, event string, data ...any) error {
// 	return ErrNotAvailable("EventsEmit")
// }

// func (api *API) EventsOff(ctx context.Context, event string, additional ...string) error {
// 	return ErrNotAvailable("EventsOff")
// }

// func (api *API) EventsOffAll(ctx context.Context) error {
// 	return ErrNotAvailable("EventsOffAll")
// }

// func (api *API) EventsOn(ctx context.Context, eventName string, callback CallbackFn) (CancelFn, error) {
// 	return nil, ErrNotAvailable("EventsOn")
// }

// func (api *API) EventsOnMultiple(ctx context.Context, eventName string, callback CallbackFn, counter int) (CancelFn, error) {
// 	return nil, ErrNotAvailable("EventsOnMultiple")
// }

// func (api *API) EventsOnce(ctx context.Context, eventName string, callback CallbackFn) (CancelFn, error) {
// 	return nil, ErrNotAvailable("EventsOnce")
// }

// func (api *API) Hide(ctx context.Context) error {
// 	return ErrNotAvailable("Hide")
// }

// func (api *API) LogDebug(ctx context.Context, message string) error {
// 	return ErrNotAvailable("LogDebug")
// }

// func (api *API) LogDebugf(ctx context.Context, format string, args ...any) error {
// 	return ErrNotAvailable("LogDebugf")
// }

// func (api *API) LogError(ctx context.Context, message string) error {
// 	return ErrNotAvailable("LogError")
// }

// func (api *API) LogErrorf(ctx context.Context, format string, args ...any) error {
// 	return ErrNotAvailable("LogErrorf")
// }

// func (api *API) LogFatal(ctx context.Context, message string) error {
// 	return ErrNotAvailable("LogFatal")
// }

// func (api *API) LogFatalf(ctx context.Context, format string, args ...any) error {
// 	return ErrNotAvailable("LogFatalf")
// }

// func (api *API) LogInfo(ctx context.Context, message string) error {
// 	return ErrNotAvailable("LogInfo")
// }

// func (api *API) LogInfof(ctx context.Context, format string, args ...any) error {
// 	return ErrNotAvailable("LogInfof")
// }

// func (api *API) LogPrint(ctx context.Context, message string) error {
// 	return ErrNotAvailable("LogPrint")
// }

// func (api *API) LogPrintf(ctx context.Context, format string, args ...any) error {
// 	return ErrNotAvailable("LogPrintf")
// }

// func (api *API) LogSetLogLevel(ctx context.Context, level logger.LogLevel) error {
// 	return ErrNotAvailable("LogSetLogLevel")
// }

// func (api *API) LogTrace(ctx context.Context, message string) error {
// 	return ErrNotAvailable("LogTrace")
// }

// func (api *API) LogTracef(ctx context.Context, format string, args ...any) error {
// 	return ErrNotAvailable("LogTracef")
// }

// func (api *API) LogWarning(ctx context.Context, message string) error {
// 	return ErrNotAvailable("LogWarning")
// }

// func (api *API) LogWarningf(ctx context.Context, format string, args ...any) error {
// 	return ErrNotAvailable("LogWarning")
// }

// func (api *API) MenuSetApplicationMenu(ctx context.Context, menu *menu.Menu) error {
// 	return ErrNotAvailable("MenuSetApplicationMenu")
// }

// func (api *API) MenuUpdateApplicationMenu(ctx context.Context) error {
// 	return ErrNotAvailable("MenuUpdateApplicationMenu")
// }

// func (api *API) MessageDialog(ctx context.Context, dialogOptions MessageDialogOptions) (string, error) {
// 	return "", ErrNotAvailable("MessageDialog")
// }

// func (api *API) OpenDirectoryDialog(ctx context.Context, dialogOptions OpenDialogOptions) (string, error) {
// 	return "", ErrNotAvailable("OpenDirectoryDialog")
// }

// func (api *API) OpenFileDialog(ctx context.Context, dialogOptions OpenDialogOptions) (string, error) {
// 	return "", ErrNotAvailable("OpenFileDialog")
// }

// func (api *API) OpenMultipleFilesDialog(ctx context.Context, dialogOptions OpenDialogOptions) ([]string, error) {
// 	return nil, ErrNotAvailable("OpenMultipleFilesDialog")
// }

// func (api *API) Quit(ctx context.Context) error {
// 	return ErrNotAvailable("Quit")
// }

// func (api *API) SaveFileDialog(ctx context.Context, dialogOptions SaveDialogOptions) (string, error) {
// 	return "", ErrNotAvailable("SaveFileDialog")
// }

// func (api *API) Show(ctx context.Context) error {
// 	return ErrNotAvailable("Show")
// }

// func (api *API) WindowCenter(ctx context.Context) error {
// 	return ErrNotAvailable("WindowCenter")
// }

// func (api *API) WindowExecJS(ctx context.Context, js string) error {
// 	return ErrNotAvailable("WindowExecJS")
// }

// func (api *API) WindowFullscreen(ctx context.Context) error {
// 	return ErrNotAvailable("WindowFullscreen")
// }

// func (api *API) WindowGetPosition(ctx context.Context) (int, int, error) {
// 	return 0, 0, ErrNotAvailable("WindowGetPosition")
// }

// func (api *API) WindowGetSize(ctx context.Context) (int, int, error) {
// 	return 0, 0, ErrNotAvailable("WindowGetSize")
// }

// func (api *API) WindowHide(ctx context.Context) error {
// 	return ErrNotAvailable("WindowHide")
// }

// func (api *API) WindowIsFullscreen(ctx context.Context) (bool, error) {
// 	return false, ErrNotAvailable("WindowIsFullscreen")
// }

// func (api *API) WindowIsMaximised(ctx context.Context) (bool, error) {
// 	return false, ErrNotAvailable("WindowIsMaximised")
// }

// func (api *API) WindowIsMinimised(ctx context.Context) (bool, error) {
// 	return false, ErrNotAvailable("WindowIsMinimised")
// }

// func (api *API) WindowIsNormal(ctx context.Context) (bool, error) {
// 	return false, ErrNotAvailable("WindowIsNormal")
// }

// func (api *API) WindowMaximise(ctx context.Context) error {
// 	return ErrNotAvailable("WindowMaximise")
// }

// func (api *API) WindowMinimise(ctx context.Context) error {
// 	return ErrNotAvailable("WindowMinimise")
// }

// func (api *API) WindowPrint(ctx context.Context) error {
// 	return ErrNotAvailable("WindowPrint")
// }

// func (api *API) WindowReload(ctx context.Context) error {
// 	return ErrNotAvailable("WindowReload")
// }

// func (api *API) WindowReloadApp(ctx context.Context) error {
// 	return ErrNotAvailable("WindowReloadApp")
// }

// func (api *API) WindowSetAlwaysOnTop(ctx context.Context, b bool) error {
// 	return ErrNotAvailable("WindowSetAlwaysOnTop")
// }

// func (api *API) WindowSetBackgroundColour(ctx context.Context, R, G, B, A uint8) error {
// 	return ErrNotAvailable("WindowSetBackgroundColour")
// }

// func (api *API) WindowSetDarkTheme(ctx context.Context) error {
// 	return ErrNotAvailable("WindowSetDarkTheme")
// }

// func (api *API) WindowSetLightTheme(ctx context.Context) error {
// 	return ErrNotAvailable("WindowSetLightTheme")
// }

// func (api *API) WindowSetMaxSize(ctx context.Context, width int, height int) error {
// 	return ErrNotAvailable("WindowSetMaxSize")
// }

// func (api *API) WindowSetMinSize(ctx context.Context, width int, height int) error {
// 	return ErrNotAvailable("WindowSetMinSize")
// }

// func (api *API) WindowSetPosition(ctx context.Context, x int, y int) error {
// 	return ErrNotAvailable("WindowSetPosition")
// }

// func (api *API) WindowSetSize(ctx context.Context, width int, height int) error {
// 	return ErrNotAvailable("WindowSetSize")
// }

// func (api *API) WindowSetSystemDefaultTheme(ctx context.Context) error {
// 	return ErrNotAvailable("WindowSetSystemDefaultTheme")
// }

// func (api *API) WindowSetTitle(ctx context.Context, title string) error {
// 	return ErrNotAvailable("WindowSetTitle")
// }

// func (api *API) WindowShow(ctx context.Context) error {
// 	return ErrNotAvailable("WindowShow")
// }

// func (api *API) WindowToggleMaximise(ctx context.Context) error {
// 	return ErrNotAvailable("WindowToggleMaximise")
// }

// func (api *API) WindowUnfullscreen(ctx context.Context) error {
// 	return ErrNotAvailable("WindowUnfullscreen")
// }

// func (api *API) WindowUnmaximise(ctx context.Context) error {
// 	return ErrNotAvailable("WindowUnmaximise")
// }

// func (api *API) WindowUnminimise(ctx context.Context) error {
// 	return ErrNotAvailable("WindowUnminimise")
// }

// func (api *API) ScreenGetAll(ctx context.Context) ([]Screen, error) {
// 	return nil, ErrNotAvailable("ScreenGetAll")
// }
