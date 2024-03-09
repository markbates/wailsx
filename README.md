
# Wailsx - A Testable, Idiomatic Wrapper for Wails

I love [Wails](https://wails.io) and have been using it to create some great applications. Unfortunately, the Wails `v2` API is not very testable or idiomatic. Wailsx is a wrapper around the Wails API that makes it easier to test and use in a more idiomatic way.

## The API Interface

The [`wailsx.API`](https://pkg.go.dev/github.com/markbates/wailsx#API) interface, [Listing 1.1](#listing-1-1) declares an idiomatic interface for the [`github.com/wailsapp/wails/v2/pkg/runtime`](https://pkg.go.dev/github.com/wailsapp/wails/v2/pkg/runtime) package.

<figure id="listing-1-1" type="listing">

<pre><code class="language-go" language="go" src="wailsrun/api.go">package wailsrun

import (
	&#34;context&#34;

	&#34;github.com/wailsapp/wails/v2/pkg/logger&#34;
	&#34;github.com/wailsapp/wails/v2/pkg/menu&#34;
)

type API interface {
	BrowserOpenURL(ctx context.Context, url string) error
	ClipboardGetText(ctx context.Context) (string, error)
	ClipboardSetText(ctx context.Context, text string) error
	EventsEmit(ctx context.Context, event string, data ...any) error
	EventsOff(ctx context.Context, event string, additional ...string) error
	EventsOffAll(ctx context.Context) error
	EventsOn(ctx context.Context, eventName string, callback CallbackFn) (CancelFn, error)
	EventsOnMultiple(ctx context.Context, eventName string, callback CallbackFn, counter int) (CancelFn, error)
	EventsOnce(ctx context.Context, eventName string, callback CallbackFn) (CancelFn, error)
	Hide(ctx context.Context) error
	LogDebug(ctx context.Context, message string) error
	LogDebugf(ctx context.Context, format string, args ...any) error
	LogError(ctx context.Context, message string) error
	LogErrorf(ctx context.Context, format string, args ...any) error
	LogFatal(ctx context.Context, message string) error
	LogFatalf(ctx context.Context, format string, args ...any) error
	LogInfo(ctx context.Context, message string) error
	LogInfof(ctx context.Context, format string, args ...any) error
	LogPrint(ctx context.Context, message string) error
	LogPrintf(ctx context.Context, format string, args ...any) error
	LogSetLogLevel(ctx context.Context, level logger.LogLevel) error
	LogTrace(ctx context.Context, message string) error
	LogTracef(ctx context.Context, format string, args ...any) error
	LogWarning(ctx context.Context, message string) error
	LogWarningf(ctx context.Context, format string, args ...any) error
	MenuSetApplicationMenu(ctx context.Context, menu *menu.Menu) error
	MenuUpdateApplicationMenu(ctx context.Context) error
	MessageDialog(ctx context.Context, dialogOptions MessageDialogOptions) (string, error)
	OpenDirectoryDialog(ctx context.Context, dialogOptions OpenDialogOptions) (string, error)
	OpenFileDialog(ctx context.Context, dialogOptions OpenDialogOptions) (string, error)
	OpenMultipleFilesDialog(ctx context.Context, dialogOptions OpenDialogOptions) ([]string, error)
	Quit(ctx context.Context) error
	SaveFileDialog(ctx context.Context, dialogOptions SaveDialogOptions) (string, error)
	Show(ctx context.Context) error
	WindowCenter(ctx context.Context) error
	WindowExecJS(ctx context.Context, js string) error
	WindowFullscreen(ctx context.Context) error
	WindowGetPosition(ctx context.Context) (int, int, error)
	WindowGetSize(ctx context.Context) (int, int, error)
	WindowHide(ctx context.Context) error
	WindowIsFullscreen(ctx context.Context) (bool, error)
	WindowIsMaximised(ctx context.Context) (bool, error)
	WindowIsMinimised(ctx context.Context) (bool, error)
	WindowIsNormal(ctx context.Context) (bool, error)
	WindowMaximise(ctx context.Context) error
	WindowMinimise(ctx context.Context) error
	WindowPrint(ctx context.Context) error
	WindowReload(ctx context.Context) error
	WindowReloadApp(ctx context.Context) error
	WindowSetAlwaysOnTop(ctx context.Context, b bool) error
	WindowSetBackgroundColour(ctx context.Context, R, G, B, A uint8) error
	WindowSetDarkTheme(ctx context.Context) error
	WindowSetLightTheme(ctx context.Context) error
	WindowSetMaxSize(ctx context.Context, width int, height int) error
	WindowSetMinSize(ctx context.Context, width int, height int) error
	WindowSetPosition(ctx context.Context, x int, y int) error
	WindowSetSize(ctx context.Context, width int, height int) error
	WindowSetSystemDefaultTheme(ctx context.Context) error
	WindowSetTitle(ctx context.Context, title string) error
	WindowShow(ctx context.Context) error
	WindowToggleMaximise(ctx context.Context) error
	WindowUnfullscreen(ctx context.Context) error
	WindowUnmaximise(ctx context.Context) error
	WindowUnminimise(ctx context.Context) error
	ScreenGetAll(ctx context.Context) ([]Screen, error)
}
</code></pre>

<figcaption><em class="figure-name">Listing 1.1:</em> The <a for="github.com/markbates/wailsx#API" href="https://pkg.go.dev/github.com/markbates/wailsx#API" target="_blank"><code>wailsx.API</code></a> interface</figcaption>
</figure>

### Error Handling

In particular, the [`wailsx.API`](https://pkg.go.dev/github.com/markbates/wailsx#API) improves error handling by adding error returns to the methods that previously returned nothing. For example, the [`github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu`](https://pkg.go.dev/github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu), [Listing 1.2](#listing-1-2), method now returns an error, [Listing 1.3](#listing-1-3).

<figure id="listing-1-2" type="listing">

<cmd data-go-version="go1.22.1" doc="github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu" exec="go doc github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu"><pre><code class="language-shell" language="shell">$ go doc github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu

package runtime // import &#34;github.com/wailsapp/wails/v2/pkg/runtime&#34;

func MenuSetApplicationMenu(ctx context.Context, menu *menu.Menu)

--------------------------------------------------------------------------------
Go Version: go1.22.1
</code></pre></cmd>

<figcaption><em class="figure-name">Listing 1.2:</em> The <a for="github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu" href="https://pkg.go.dev/github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu" target="_blank"><code>github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu</code></a> method</figcaption>
</figure>

<figure id="listing-1-3" type="listing">

<cmd data-go-version="go1.22.1" doc="github.com/markbates/wailsx/wailsrun.API.MenuSetApplicationMenu" exec="go doc github.com/markbates/wailsx/wailsrun.API.MenuSetApplicationMenu"><pre><code class="language-shell" language="shell">$ go doc github.com/markbates/wailsx/wailsrun.API.MenuSetApplicationMenu

package wailsrun // import &#34;github.com/markbates/wailsx/wailsrun&#34;

type API interface {
	MenuSetApplicationMenu(ctx context.Context, menu *menu.Menu) error
}

--------------------------------------------------------------------------------
Go Version: go1.22.1
</code></pre></cmd>

<figcaption><em class="figure-name">Listing 1.3:</em> The <a for="github.com/markbates/wailsx/wailsrun.API.MenuSetApplicationMenu" href="https://pkg.go.dev/github.com/markbates/wailsx/wailsrun.API.MenuSetApplicationMenu" target="_blank"><code>github.com/markbates/wailsx/wailsrun.API.MenuSetApplicationMenu</code></a> method</figcaption>
</figure>

## Protecting Wails API Calls

