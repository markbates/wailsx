
# <a id="heading-0"></a><toc-level>1.1</toc-level> - Wailsx - A Testable, Idiomatic Wrapper for Wails

I love [Wails](https://wails.io) and have been using it to create some great applications. Unfortunately, the Wails `v2` API is not very testable or idiomatic. Wailsx is a wrapper around the Wails API that makes it easier to test and use in a more idiomatic way.

<toc>

* [<toc-level>1.1</toc-level> - Wailsx - A Testable, Idiomatic Wrapper for Wails](#heading-0)
* [<toc-level>1.2</toc-level> - Runtime and API](#heading-1)


* [<toc-level>1.2.1</toc-level> - The API Interface](#heading-2)


* [<toc-level>1.2.1.1</toc-level> - Error Handling](#heading-3)

* [<toc-level>1.2.2</toc-level> - Protecting Wails API Calls](#heading-4)


* [<toc-level>1.2.2.2</toc-level> - Testing Invalid Wails API Calls](#heading-5)


* [<toc-level>1.3</toc-level> - Clipboard](#heading-6)


* [<toc-level>1.3.1</toc-level> - The `ClipboardManager` Interface](#heading-7)

* [<toc-level>1.4</toc-level> - Dialogs](#heading-8)


* [<toc-level>1.4.1</toc-level> - The `DialogManager` Interface](#heading-9)

* [<toc-level>1.5</toc-level> - Events](#heading-10)


* [<toc-level>1.5.1</toc-level> - The `EventManager` Interface](#heading-11)

* [<toc-level>1.6</toc-level> - Messages](#heading-12)


* [<toc-level>1.6.1</toc-level> - The `Messenger` Interface](#heading-13)
* [<toc-level>1.6.2</toc-level> - The `ErrorMessenger` Interface](#heading-14)

* [<toc-level>1.7</toc-level> - Logging](#heading-15)


* [<toc-level>1.7.1</toc-level> - The `WailsLogger` Interface](#heading-16)

* [<toc-level>1.8</toc-level> - Menus](#heading-17)


* [<toc-level>1.8.1</toc-level> - The `MenuManager` Interface](#heading-18)

* [<toc-level>1.9</toc-level> - State Data](#heading-19)


* [<toc-level>1.9.1</toc-level> - The `DataProvider` Interface](#heading-20)

* [<toc-level>1.10</toc-level> - Window Management](#heading-21)


* [<toc-level>1.10.1</toc-level> - The `WindowManager` Interface](#heading-22)
* [<toc-level>1.10.2</toc-level> - The `MaximiseManager` Interface](#heading-23)
* [<toc-level>1.10.3</toc-level> - The `PositionManager` Interface](#heading-24)
* [<toc-level>1.10.4</toc-level> - The `ReloadManager` Interface](#heading-25)
* [<toc-level>1.10.5</toc-level> - The `ThemeManager` Interface](#heading-26)
* [<toc-level>1.10.6</toc-level> - The `Toggler` Interface](#heading-27)

</toc>

---

# <a id="heading-1"></a><toc-level>1.2</toc-level> - Runtime and API

## <a id="heading-2"></a><toc-level>1.2.1</toc-level> - The API Interface

The [`wailsx.API`](https://pkg.go.dev/github.com/markbates/wailsx#API) interface, [Listing 1.1](#listing-1-1) declares an idiomatic interface for the [`github.com/wailsapp/wails/v2/pkg/runtime`](https://pkg.go.dev/github.com/wailsapp/wails/v2/pkg/runtime) package.

<a id="listing-1-1"></a>


```shell
$ go doc github.com/markbates/wailsx/wailsrun.API

package wailsrun // import "github.com/markbates/wailsx/wailsrun"

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

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.1:_ The [`wailsx.API`](https://pkg.go.dev/github.com/markbates/wailsx#API) interface*


### <a id="heading-3"></a><toc-level>1.2.1.1</toc-level> - Error Handling

In particular, the [`wailsx.API`](https://pkg.go.dev/github.com/markbates/wailsx#API) improves error handling by adding error returns to the methods that previously returned nothing. For example, the [`github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu`](https://pkg.go.dev/github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu), [Listing 1.2](#listing-1-2), method now returns an error, [Listing 1.3](#listing-1-3).

<a id="listing-1-2"></a>


```shell
$ go doc github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu

package runtime // import "github.com/wailsapp/wails/v2/pkg/runtime"

func MenuSetApplicationMenu(ctx context.Context, menu *menu.Menu)

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.2:_ The [`github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu`](https://pkg.go.dev/github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu) method*


<a id="listing-1-3"></a>


```shell
$ go doc github.com/markbates/wailsx/wailsrun.API.MenuSetApplicationMenu

package wailsrun // import "github.com/markbates/wailsx/wailsrun"

type API interface {
	MenuSetApplicationMenu(ctx context.Context, menu *menu.Menu) error
}

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.3:_ The [`github.com/markbates/wailsx/wailsrun.API.MenuSetApplicationMenu`](https://pkg.go.dev/github.com/markbates/wailsx/wailsrun.API.MenuSetApplicationMenu) method*


## <a id="heading-4"></a><toc-level>1.2.2</toc-level> - Protecting Wails API Calls

Wailsx uses Go build tags to protect the Wails API calls from being called in a production environment. The [`wailsrun.API`](https://pkg.go.dev/github.com/markbates/wailsx/wailsrun#API) interface, [Listing 1.1](#listing-1-1), is implemented in two different files, [Listing 1.4](#listing-1-4) and [Listing 1.5](#listing-1-5).

The `wailsrun/wailscalls_prod.go` file, [Listing 1.4](#listing-1-4), is only built when any of the following builds are provided: `wails || dev || desktop || production`. This file contains the actual Wails API calls and most returned errors are `nil`.

<a id="listing-1-4"></a>


```go
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
```
> *source: wailsrun/wailscalls_prod.go:BrowserOpenURL*


> *_Listing 1.4:_ Production Wails API calls: `wailsrun/wailscalls_prod.go`*


In all other environments, such as testing, the `wailsrun/wailscalls.go` file, [Listing 1.5](#listing-1-5), is built in all environments and contains the Wailsx API calls. The Wailsx API calls are then used to call the Wails API calls in the `development` environment.

<a id="listing-1-5"></a>


```go
import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

func BrowserOpenURL(ctx context.Context, url string) error {
	return ErrNotAvailable("BrowserOpenURL")
}
```
> *source: wailsrun/wailscalls.go:BrowserOpenURL*


> *_Listing 1.5:_ Stubbed Wails API calls: `wailsrun/wailscalls.go`*


In these environments all of the Wails API calls will return the `ErrNotAvailable` error, [Listing 1.6](#listing-1-6).

<a id="listing-1-6"></a>


```shell
$ go doc github.com/markbates/wailsx/wailsrun.ErrNotAvailable

package wailsrun // import "github.com/markbates/wailsx/wailsrun"

type ErrNotAvailable string

func (e ErrNotAvailable) Error() string

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.6:_ The `ErrNotAvailable` error*


### <a id="heading-5"></a><toc-level>1.2.2.2</toc-level> - Testing Invalid Wails API Calls

With the help of Go build tags, any direct calls made to the Wails API, _outside_ of a running Wails application, will return the [`github.com/markbates/wailsx/wailsrun.ErrNotAvailable`](https://pkg.go.dev/github.com/markbates/wailsx/wailsrun.ErrNotAvailable) error. This allows for testing of the Wails API calls in a non-Wails environment.

In the test seen in [Listing 1.7](#listing-1-7) we are making a direct call to the Wails API and checking the error returned. The test passes when the error returned is `ErrNotAvailable`.

<a id="listing-1-7"></a>


```go
package wailsrun_test

import (
	"context"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/stretchr/testify/require"
)

func Test_ErrNotAvailable(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	ctx := context.Background()

	err := wailsrun.BrowserOpenURL(ctx, "https://example.com")
	r.Error(err)

	exp := wailsrun.ErrNotAvailable("BrowserOpenURL")
	r.Equal(exp, err)
}
```
> *source: wailsrun/api_calls_test.go*


> *_Listing 1.7:_ Testing the `BrowserOpenURL` method*


When running the tests outside of a Wails application, the `BrowserOpenURL` method will return the `ErrNotAvailable` error, [Listing 1.6](#listing-1-6).

<a id="listing-1-8"></a>


```shell
$ go test -v -run Test_ErrNotAvailable

testing: warning: no tests to run
PASS
ok  	github.com/markbates/wailsx	0.003s

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.8:_ Testing the `BrowserOpenURL` method output.*


If the tests are run in a Wails application, using one of the known build tags, the `BrowserOpenURL` method will call the actual Wails API method, [Listing 1.4](#listing-1-4). The result is a call to [`log.Fatal`](https://pkg.go.dev/log.Fatal) because we don't have a valid Wails context.

<a id="listing-1-9"></a>


```shell
$ go test -v -run Test_ErrNotAvailable -tags wails

testing: warning: no tests to run
PASS
ok  	github.com/markbates/wailsx	0.003s

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.9:_ Testing the `BrowserOpenURL` method output in `production`.*


---

# <a id="heading-6"></a><toc-level>1.3</toc-level> - Clipboard

## <a id="heading-7"></a><toc-level>1.3.1</toc-level> - The <code>ClipboardManager</code> Interface

<a id="figure-1-1"></a>


```shell
$ go doc github.com/markbates/wailsx/clipx.ClipboardManager

package clipx // import "github.com/markbates/wailsx/clipx"

type ClipboardManager interface {
	ClipboardGetText(ctx context.Context) (string, error)
	ClipboardSetText(ctx context.Context, text string) error
}

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Figure 1.1:_ The [`github.com/markbates/wailsx/clipx.ClipboardManager`](https://pkg.go.dev/github.com/markbates/wailsx/clipx.ClipboardManager) interface*


---

# <a id="heading-8"></a><toc-level>1.4</toc-level> - Dialogs

## <a id="heading-9"></a><toc-level>1.4.1</toc-level> - The <code>DialogManager</code> Interface

<a id="listing-1-10"></a>


```shell
$ go doc github.com/markbates/wailsx/dialogx.DialogManager

package dialogx // import "github.com/markbates/wailsx/dialogx"

type DialogManager interface {
	MessageDialog(ctx context.Context, opts MessageDialogOptions) (string, error)
	OpenDirectoryDialog(ctx context.Context, opts OpenDialogOptions) (string, error)
	OpenFileDialog(ctx context.Context, opts OpenDialogOptions) (string, error)
	OpenMultipleFilesDialog(ctx context.Context, opts OpenDialogOptions) ([]string, error)
	SaveFileDialog(ctx context.Context, opts SaveDialogOptions) (string, error)
}

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.10:_ The [`github.com/markbates/wailsx/dialogx.DialogManager`](https://pkg.go.dev/github.com/markbates/wailsx/dialogx.DialogManager) interface*


---

# <a id="heading-10"></a><toc-level>1.5</toc-level> - Events

## <a id="heading-11"></a><toc-level>1.5.1</toc-level> - The <code>EventManager</code> Interface

<a id="listing-1-11"></a>


```shell
$ go doc github.com/markbates/wailsx/eventx.EventManager

package eventx // import "github.com/markbates/wailsx/eventx"

type EventManager interface {
	EventsEmit(ctx context.Context, event string, args ...any) (err error)
	EventsOff(ctx context.Context, name string, additional ...string) error
	EventsOffAll(ctx context.Context) error
	EventsOn(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error)
	EventsOnMultiple(ctx context.Context, name string, callback wailsrun.CallbackFn, counter int) (wailsrun.CancelFn, error)
	EventsOnce(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error)
}

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.11:_ The [`github.com/markbates/wailsx/eventx.EventManager`](https://pkg.go.dev/github.com/markbates/wailsx/eventx.EventManager) interface*


---

# <a id="heading-12"></a><toc-level>1.6</toc-level> - Messages

## <a id="heading-13"></a><toc-level>1.6.1</toc-level> - The <code>Messenger</code> Interface

<a id="listing-1-12"></a>


```shell
$ go doc github.com/markbates/wailsx/eventx/msgx.Messenger

package msgx // import "github.com/markbates/wailsx/eventx/msgx"

type Messenger interface {
	MsgEvent() string
	MsgText() string
	MsgTime() time.Time
	MsgData() any
}

func NewMessage(event string, now time.Time, arg any) Messenger

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.12:_ The [`github.com/markbates/wailsx/eventx/msgx.Messenger`](https://pkg.go.dev/github.com/markbates/wailsx/eventx/msgx.Messenger) interface*


## <a id="heading-14"></a><toc-level>1.6.2</toc-level> - The <code>ErrorMessenger</code> Interface

<a id="listing-1-13"></a>


```shell
$ go doc github.com/markbates/wailsx/eventx/msgx.ErrorMessenger

package msgx // import "github.com/markbates/wailsx/eventx/msgx"

type ErrorMessenger interface {
	Messenger
	MsgError() error
}

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.13:_ The [`github.com/markbates/wailsx/eventx/msgx.ErrorMessenger`](https://pkg.go.dev/github.com/markbates/wailsx/eventx/msgx.ErrorMessenger) interface*


---

# <a id="heading-15"></a><toc-level>1.7</toc-level> - Logging

## <a id="heading-16"></a><toc-level>1.7.1</toc-level> - The <code>WailsLogger</code> Interface

<a id="listing-1-14"></a>


```shell
$ go doc github.com/markbates/wailsx/logx.WailsLogger

package logx // import "github.com/markbates/wailsx/logx"

type WailsLogger interface {
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
	LogSetLogLevel(ctx context.Context, level wailsrun.LogLevel) error
	LogTrace(ctx context.Context, message string) error
	LogTracef(ctx context.Context, format string, args ...any) error
	LogWarning(ctx context.Context, message string) error
	LogWarningf(ctx context.Context, format string, args ...any) error
}

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.14:_ The [`github.com/markbates/wailsx/logx.WailsLogger`](https://pkg.go.dev/github.com/markbates/wailsx/logx.WailsLogger) interface*


---

# <a id="heading-17"></a><toc-level>1.8</toc-level> - Menus

## <a id="heading-18"></a><toc-level>1.8.1</toc-level> - The <code>MenuManager</code> Interface

<a id="listing-1-15"></a>


```shell
$ go doc github.com/markbates/wailsx/menux.MenuManager

package menux // import "github.com/markbates/wailsx/menux"

type MenuManager interface {
	MenuSetApplicationMenu(ctx context.Context, menu *menu.Menu) error
	MenuUpdateApplicationMenu(ctx context.Context) error
}

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.15:_ The [`github.com/markbates/wailsx/menux.MenuManager`](https://pkg.go.dev/github.com/markbates/wailsx/menux.MenuManager) interface*


---

# <a id="heading-19"></a><toc-level>1.9</toc-level> - State Data

## <a id="heading-20"></a><toc-level>1.9.1</toc-level> - The <code>DataProvider</code> Interface

<a id="listing-1-16"></a>


```shell
$ go doc github.com/markbates/wailsx/statedata.DataProvider

package statedata // import "github.com/markbates/wailsx/statedata"

type DataProvider[T any] interface {
	StateData(ctx context.Context) (Data[T], error)
}

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.16:_ The [`github.com/markbates/wailsx/statedata.DataProvider`](https://pkg.go.dev/github.com/markbates/wailsx/statedata.DataProvider) interface*


---

# <a id="heading-21"></a><toc-level>1.10</toc-level> - Window Management

## <a id="heading-22"></a><toc-level>1.10.1</toc-level> - The <code>WindowManager</code> Interface

<a id="listing-1-17"></a>


```shell
$ go doc github.com/markbates/wailsx/windowx.WindowManager

package windowx // import "github.com/markbates/wailsx/windowx"

type WindowManager interface {
	MaximiseManager
	PositionManager
	ReloadManager
	ThemeManager
	Toggler

	ScreenGetAll(ctx context.Context) ([]Screen, error)
	WindowExecJS(ctx context.Context, js string) error
	WindowPrint(ctx context.Context) error
	WindowSetAlwaysOnTop(ctx context.Context, b bool) error
	WindowSetTitle(ctx context.Context, title string) error
}

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.17:_ The [`github.com/markbates/wailsx/windowx.WindowManager`](https://pkg.go.dev/github.com/markbates/wailsx/windowx.WindowManager) interface*


## <a id="heading-23"></a><toc-level>1.10.2</toc-level> - The <code>MaximiseManager</code> Interface

<a id="listing-1-18"></a>


```shell
$ go doc github.com/markbates/wailsx/windowx.MaximiseManager

package windowx // import "github.com/markbates/wailsx/windowx"

type MaximiseManager interface {
	WindowFullscreen(ctx context.Context) error
	WindowIsFullscreen(ctx context.Context) (bool, error)
	WindowIsMaximised(ctx context.Context) (bool, error)
	WindowIsMinimised(ctx context.Context) (bool, error)
	WindowIsNormal(ctx context.Context) (bool, error)
	WindowMaximise(ctx context.Context) error
	WindowMinimise(ctx context.Context) error
	WindowToggleMaximise(ctx context.Context) error
	WindowUnfullscreen(ctx context.Context) error
	WindowUnmaximise(ctx context.Context) error
	WindowUnminimise(ctx context.Context) error
}

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.18:_ The [`github.com/markbates/wailsx/windowx.MaximiseManager`](https://pkg.go.dev/github.com/markbates/wailsx/windowx.MaximiseManager) interface*


## <a id="heading-24"></a><toc-level>1.10.3</toc-level> - The <code>PositionManager</code> Interface

<a id="listing-1-19"></a>


```shell
$ go doc github.com/markbates/wailsx/windowx.PositionManager

package windowx // import "github.com/markbates/wailsx/windowx"

type PositionManager interface {
	WindowCenter(ctx context.Context) error
	WindowGetPosition(ctx context.Context) (int, int, error)
	WindowGetSize(ctx context.Context) (int, int, error)
	WindowSetMaxSize(ctx context.Context, width int, height int) error
	WindowSetMinSize(ctx context.Context, width int, height int) error
	WindowSetPosition(ctx context.Context, x int, y int) error
	WindowSetSize(ctx context.Context, width int, height int) error
}

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.19:_ The [`github.com/markbates/wailsx/windowx.PositionManager`](https://pkg.go.dev/github.com/markbates/wailsx/windowx.PositionManager) interface*


## <a id="heading-25"></a><toc-level>1.10.4</toc-level> - The <code>ReloadManager</code> Interface

<a id="listing-1-20"></a>


```shell
$ go doc github.com/markbates/wailsx/windowx.ReloadManager

package windowx // import "github.com/markbates/wailsx/windowx"

type ReloadManager interface {
	WindowReload(ctx context.Context) error
	WindowReloadApp(ctx context.Context) error
}

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.20:_ The [`github.com/markbates/wailsx/windowx.ReloadManager`](https://pkg.go.dev/github.com/markbates/wailsx/windowx.ReloadManager) interface*


## <a id="heading-26"></a><toc-level>1.10.5</toc-level> - The <code>ThemeManager</code> Interface

<a id="listing-1-21"></a>


```shell
$ go doc github.com/markbates/wailsx/windowx.ThemeManager

package windowx // import "github.com/markbates/wailsx/windowx"

type ThemeManager interface {
	WindowSetBackgroundColour(ctx context.Context, R, G, B, A uint8) error
	WindowSetDarkTheme(ctx context.Context) error
	WindowSetLightTheme(ctx context.Context) error
	WindowSetSystemDefaultTheme(ctx context.Context) error
}

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.21:_ The [`github.com/markbates/wailsx/windowx.ThemeManager`](https://pkg.go.dev/github.com/markbates/wailsx/windowx.ThemeManager) interface*


## <a id="heading-27"></a><toc-level>1.10.6</toc-level> - The <code>Toggler</code> Interface

<a id="listing-1-22"></a>


```shell
$ go doc github.com/markbates/wailsx/windowx.Toggler

package windowx // import "github.com/markbates/wailsx/windowx"

type Toggler interface {
	Hide(ctx context.Context) error
	Show(ctx context.Context) error
	WindowHide(ctx context.Context) error
	WindowShow(ctx context.Context) error
}

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.22:_ The [`github.com/markbates/wailsx/windowx.Toggler`](https://pkg.go.dev/github.com/markbates/wailsx/windowx.Toggler) interface*


