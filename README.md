TODO: Element.MD(): "toc"
TODO: Element.MD(): "level"
TODO: Element.MD(): "level"
TODO: Element.MD(): "level"
TODO: Element.MD(): "level"
TODO: Element.MD(): "level"

# <a id="heading-0"></a>Wailsx - A Testable, Idiomatic Wrapper for Wails

I love [Wails](https://wails.io) and have been using it to create some great applications. Unfortunately, the Wails `v2` API is not very testable or idiomatic. Wailsx is a wrapper around the Wails API that makes it easier to test and use in a more idiomatic way.

<toc>
* [<level>1.1</level> Wailsx - A Testable, Idiomatic Wrapper for Wails](#heading-0)


* [<level>1.1.1</level> The API Interface](#heading-1)


* [<level>1.1.1.1</level> Error Handling](#heading-2)

* [<level>1.1.2</level> Protecting Wails API Calls](#heading-3)


* [<level>1.1.2.2</level> Testing Invalid Wails API Calls](#heading-4)


</toc>

## <a id="heading-1"></a>The API Interface

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

### <a id="heading-2"></a>Error Handling

In particular, the [`wailsx.API`](https://pkg.go.dev/github.com/markbates/wailsx#API) improves error handling by adding error returns to the methods that previously returned nothing. For example, the [`github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu`](https://pkg.go.dev/github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu), [Listing 1.2](#listing-1-2), method now returns an error, [Listing 1.3](#listing-1-3).

<figure id="listing-1-2" type="listing">

<cmd data-go-version="go1.22.0" doc="github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu" exec="go doc github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu"><pre><code class="language-shell" language="shell">$ go doc github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu

package runtime // import &#34;github.com/wailsapp/wails/v2/pkg/runtime&#34;

func MenuSetApplicationMenu(ctx context.Context, menu *menu.Menu)

--------------------------------------------------------------------------------
Go Version: go1.22.0
</code></pre></cmd>

<figcaption><em class="figure-name">Listing 1.2:</em> The <a for="github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu" href="https://pkg.go.dev/github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu" target="_blank"><code>github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu</code></a> method</figcaption>
</figure>

<figure id="listing-1-3" type="listing">

<cmd data-go-version="go1.22.0" doc="github.com/markbates/wailsx/wailsrun.API.MenuSetApplicationMenu" exec="go doc github.com/markbates/wailsx/wailsrun.API.MenuSetApplicationMenu"><pre><code class="language-shell" language="shell">$ go doc github.com/markbates/wailsx/wailsrun.API.MenuSetApplicationMenu

package wailsrun // import &#34;github.com/markbates/wailsx/wailsrun&#34;

type API interface {
	MenuSetApplicationMenu(ctx context.Context, menu *menu.Menu) error
}

--------------------------------------------------------------------------------
Go Version: go1.22.0
</code></pre></cmd>

<figcaption><em class="figure-name">Listing 1.3:</em> The <a for="github.com/markbates/wailsx/wailsrun.API.MenuSetApplicationMenu" href="https://pkg.go.dev/github.com/markbates/wailsx/wailsrun.API.MenuSetApplicationMenu" target="_blank"><code>github.com/markbates/wailsx/wailsrun.API.MenuSetApplicationMenu</code></a> method</figcaption>
</figure>

## <a id="heading-3"></a>Protecting Wails API Calls

Wailsx uses Go build tags to protect the Wails API calls from being called in a production environment. The [`wailsrun.API`](https://pkg.go.dev/github.com/markbates/wailsx/wailsrun#API) interface, [Listing 1.1](#listing-1-1), is implemented in two different files, [Listing 1.4](#listing-1-4) and [Listing 1.5](#listing-1-5).

The `wailsrun/wailscalls_prod.go` file, [Listing 1.4](#listing-1-4), is only built when any of the following builds are provided: `wails || dev || desktop || production`. This file contains the actual Wails API calls and most returned errors are `nil`.

<figure id="listing-1-4" type="listing">

<pre><code class="language-go" language="go" snippet="BrowserOpenURL" src="wailsrun/wailscalls_prod.go">import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func BrowserOpenURL(ctx context.Context, url string) error {
	runtime.BrowserOpenURL(ctx, url)
	return nil
}</code></pre>

<figcaption><em class="figure-name">Listing 1.4:</em> Production Wails API calls: <code>wailsrun/wailscalls_prod.go</code></figcaption>
</figure>

In all other environments, such as testing, the `wailsrun/wailscalls.go` file, [Listing 1.5](#listing-1-5), is built in all environments and contains the Wailsx API calls. The Wailsx API calls are then used to call the Wails API calls in the `development` environment.

<figure id="listing-1-5" type="listing">

<pre><code class="language-go" language="go" snippet="BrowserOpenURL" src="wailsrun/wailscalls.go">import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

func BrowserOpenURL(ctx context.Context, url string) error {
	return ErrNotAvailable("BrowserOpenURL")
}</code></pre>

<figcaption><em class="figure-name">Listing 1.5:</em> Stubbed Wails API calls: <code>wailsrun/wailscalls.go</code></figcaption>
</figure>

In these environments all of the Wails API calls will return the `ErrNotAvailable` error, [Listing 1.6](#listing-1-6).

<figure id="listing-1-6" type="listing">

<cmd data-go-version="go1.22.0" doc="github.com/markbates/wailsx/wailsrun.ErrNotAvailable" exec="go doc github.com/markbates/wailsx/wailsrun.ErrNotAvailable"><pre><code class="language-shell" language="shell">$ go doc github.com/markbates/wailsx/wailsrun.ErrNotAvailable

package wailsrun // import &#34;github.com/markbates/wailsx/wailsrun&#34;

type ErrNotAvailable string

func (e ErrNotAvailable) Error() string

--------------------------------------------------------------------------------
Go Version: go1.22.0
</code></pre></cmd>

<figcaption><em class="figure-name">Listing 1.6:</em> The <code>ErrNotAvailable</code> error</figcaption>
</figure>

### <a id="heading-4"></a>Testing Invalid Wails API Calls

With the help of Go build tags, any direct calls made to the Wails API, _outside_ of a running Wails application, will return the [`github.com/markbates/wailsx/wailsrun.ErrNotAvailable`](https://pkg.go.dev/github.com/markbates/wailsx/wailsrun.ErrNotAvailable) error. This allows for testing of the Wails API calls in a non-Wails environment.

In the test seen in [Listing 1.7](#listing-1-7) we are making a direct call to the Wails API and checking the error returned. The test passes when the error returned is `ErrNotAvailable`.

<figure id="listing-1-7" type="listing">

<pre><code class="language-go" language="go" src="docs/examples/api_calls/api_calls_test.go">package demo

import (
	&#34;context&#34;
	&#34;testing&#34;

	&#34;github.com/markbates/wailsx/wailsrun&#34;
	&#34;github.com/stretchr/testify/require&#34;
)

func Test_ErrNotAvailable(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	ctx := context.Background()

	err := wailsrun.BrowserOpenURL(ctx, &#34;https://example.com&#34;)
	r.Error(err)

	exp := wailsrun.ErrNotAvailable(&#34;BrowserOpenURL&#34;)
	r.Equal(exp, err)
}
</code></pre>

<figcaption><em class="figure-name">Listing 1.7:</em> Testing the <code>BrowserOpenURL</code> method</figcaption>
</figure>

When running the tests outside of a Wails application, the `BrowserOpenURL` method will return the `ErrNotAvailable` error, [Listing 1.6](#listing-1-6).

<figure id="listing-1-8" type="listing">

<cmd data-go-version="go1.22.0" exec="go test -v" src="docs/examples/api_calls" test="-v"><pre><code class="language-shell" language="shell">$ go test -v

=== RUN   Test_ErrNotAvailable
=== PAUSE Test_ErrNotAvailable
=== CONT  Test_ErrNotAvailable
--- PASS: Test_ErrNotAvailable (0.00s)
PASS
ok  	demo	0.003s

go: downloading github.com/markbates/wailsx v0.0.0-20240308181724-83fae333e653
go: downloading github.com/stretchr/testify v1.9.0

--------------------------------------------------------------------------------
Go Version: go1.22.0
</code></pre></cmd>

<figcaption><em class="figure-name">Listing 1.8:</em> Testing the <code>BrowserOpenURL</code> method output.</figcaption>
</figure>

If the tests are run in a Wails application, using one of the known build tags, the `BrowserOpenURL` method will call the actual Wails API method, [Listing 1.4](#listing-1-4). The result is a call to [`log.Fatal`](https://pkg.go.dev/log.Fatal) because we don't have a valid Wails context.

<figure id="listing-1-9" type="listing">

<cmd data-go-version="go1.22.0" exec="go test -v -tags wails" exit="1" src="docs/examples/api_calls" test="-v -tags wails"><pre><code class="language-shell" language="shell">$ go test -v -tags wails

=== RUN   Test_ErrNotAvailable
=== PAUSE Test_ErrNotAvailable
=== CONT  Test_ErrNotAvailable
2024/03/09 04:33:21 cannot call &#39;github.com/wailsapp/wails/v2/pkg/runtime.BrowserOpenURL&#39;: An invalid context was passed. This method requires the specific context given in the lifecycle hooks:
https://wails.io/docs/reference/runtime/intro
exit status 1
FAIL	demo	0.003s

go: downloading github.com/markbates/wailsx v0.0.0-20240308181724-83fae333e653
go: downloading github.com/stretchr/testify v1.9.0

--------------------------------------------------------------------------------
Go Version: go1.22.0
</code></pre></cmd>

<figcaption><em class="figure-name">Listing 1.9:</em> Testing the <code>BrowserOpenURL</code> method output in <code>production</code>.</figcaption>
</figure>

