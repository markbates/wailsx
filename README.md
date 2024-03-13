
# <a id="heading-0"></a><toc-level>1.1</toc-level> - Wailsx - A Testable, Idiomatic Wrapper for Wails

I love [Wails](https://wails.io) and have been using it to create some great applications. Unfortunately, the Wails `v2` API is not very testable or idiomatic. Wailsx is a wrapper around the Wails API that makes it easier to test and use in a more idiomatic way.

<toc>

* [<toc-level>1.1</toc-level> - Wailsx - A Testable, Idiomatic Wrapper for Wails](#heading-0)


* [<toc-level>1.1.1</toc-level> - Installation](#heading-1)

* [<toc-level>1.2</toc-level> - Runtime and API](#heading-2)


* [<toc-level>1.2.1</toc-level> - The API Interface](#heading-3)


* [<toc-level>1.2.1.1</toc-level> - Error Handling](#heading-4)

* [<toc-level>1.2.2</toc-level> - Protecting Wails API Calls](#heading-5)


* [<toc-level>1.2.2.2</toc-level> - Testing Invalid Wails API Calls](#heading-6)


* [<toc-level>1.3</toc-level> - Clipboard](#heading-7)


* [<toc-level>1.3.1</toc-level> - The `ClipboardManager` Interface](#heading-8)

* [<toc-level>1.4</toc-level> - Dialogs](#heading-9)


* [<toc-level>1.4.1</toc-level> - The `DialogManager` Interface](#heading-10)

* [<toc-level>1.5</toc-level> - Events](#heading-11)


* [<toc-level>1.5.1</toc-level> - The `EventManager` Interface](#heading-12)
* [<toc-level>1.5.2</toc-level> - The `Manager` Type](#heading-13)


* [<toc-level>1.5.2.1</toc-level> - Creating a New Manager](#heading-14)

* [<toc-level>1.5.3</toc-level> - The `CallbackFn` Type](#heading-15)
* [<toc-level>1.5.4</toc-level> - The `CancelFn` Type](#heading-16)

* [<toc-level>1.6</toc-level> - Messages](#heading-17)


* [<toc-level>1.6.1</toc-level> - The `Messenger` Interface](#heading-18)
* [<toc-level>1.6.2</toc-level> - The `ErrorMessenger` Interface](#heading-19)

* [<toc-level>1.7</toc-level> - Logging](#heading-20)


* [<toc-level>1.7.1</toc-level> - The `WailsLogger` Interface](#heading-21)

* [<toc-level>1.8</toc-level> - Menus](#heading-22)


* [<toc-level>1.8.1</toc-level> - The `MenuManager` Interface](#heading-23)

* [<toc-level>1.9</toc-level> - State Data](#heading-24)


* [<toc-level>1.9.1</toc-level> - The `DataProvider` Interface](#heading-25)
* [<toc-level>1.9.2</toc-level> - The `Data` Interface](#heading-26)

* [<toc-level>1.10</toc-level> - Window Management](#heading-27)


* [<toc-level>1.10.1</toc-level> - The `WindowManager` Interface](#heading-28)
* [<toc-level>1.10.2</toc-level> - The `MaximiseManager` Interface](#heading-29)
* [<toc-level>1.10.3</toc-level> - The `PositionManager` Interface](#heading-30)
* [<toc-level>1.10.4</toc-level> - The `ReloadManager` Interface](#heading-31)
* [<toc-level>1.10.5</toc-level> - The `ThemeManager` Interface](#heading-32)
* [<toc-level>1.10.6</toc-level> - The `Toggler` Interface](#heading-33)

* [<toc-level>1.11</toc-level> - Using the API](#heading-34)


* [<toc-level>1.11.1</toc-level> - The `API` type](#heading-35)
* [<toc-level>1.11.2</toc-level> - `Nil` API Calls](#heading-36)
* [<toc-level>1.11.3</toc-level> - `Nop` API Calls](#heading-37)

</toc>

## <a id="heading-1"></a><toc-level>1.1.1</toc-level> - Installation

Wailsx is a Go module and can be installed with `go get`.

<a id="figure-1-1"></a>


```bash
go get github.com/markbates/wailsx

```

> *_Figure 1.1:_ Installing `github.com/markbates/wailsx` with `go get`.*


Once imported, you can use the `wailsx` package in your application.

<a id="figure-1-2"></a>


```shell
$ go doc -short github.com/markbates/wailsx.NewAPI

func NewAPI() *API
    NewAPI returns a new API with all the functions, and interfaces, set to
    their default implementations.

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Figure 1.2:_ The [`wailsx.NewAPI`](https://pkg.go.dev/github.com/markbates/wailsx#NewAPI) function*


---

# <a id="heading-2"></a><toc-level>1.2</toc-level> - Runtime and API

## <a id="heading-3"></a><toc-level>1.2.1</toc-level> - The API Interface

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
	EventsOn(ctx context.Context, event string, callback CallbackFn) (CancelFn, error)
	EventsOnMultiple(ctx context.Context, event string, callback CallbackFn, counter int) (CancelFn, error)
	EventsOnce(ctx context.Context, event string, callback CallbackFn) (CancelFn, error)
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


### <a id="heading-4"></a><toc-level>1.2.1.1</toc-level> - Error Handling

In particular, the [`wailsx.API`](https://pkg.go.dev/github.com/markbates/wailsx#API) improves error handling by adding error returns to the methods that previously returned nothing. For example, the [`runtime.MenuSetApplicationMenu`](https://pkg.go.dev/github.com/wailsapp/wails/v2/pkg/runtime#MenuSetApplicationMenu), [Listing 1.2](#listing-1-2), method now returns an error, [Listing 1.3](#listing-1-3).

<a id="listing-1-2"></a>


```shell
$ go doc github.com/wailsapp/wails/v2/pkg/runtime.MenuSetApplicationMenu

package runtime // import "github.com/wailsapp/wails/v2/pkg/runtime"

func MenuSetApplicationMenu(ctx context.Context, menu *menu.Menu)

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.2:_ The [`runtime.MenuSetApplicationMenu`](https://pkg.go.dev/github.com/wailsapp/wails/v2/pkg/runtime#MenuSetApplicationMenu) method*


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

> *_Listing 1.3:_ The [`wailsrun.API.MenuSetApplicationMenu`](https://pkg.go.dev/github.com/markbates/wailsx/wailsrun#API.MenuSetApplicationMenu) method*


## <a id="heading-5"></a><toc-level>1.2.2</toc-level> - Protecting Wails API Calls

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

> *_Listing 1.6:_ The [`wailsrun.ErrNotAvailable`](https://pkg.go.dev/github.com/markbates/wailsx/wailsrun#ErrNotAvailable) error*


### <a id="heading-6"></a><toc-level>1.2.2.2</toc-level> - Testing Invalid Wails API Calls

With the help of Go build tags, any direct calls made to the Wails API, _outside_ of a running Wails application, will return the [`wailsrun.ErrNotAvailable`](https://pkg.go.dev/github.com/markbates/wailsx/wailsrun#ErrNotAvailable) error. This allows for testing of the Wails API calls in a non-Wails environment.

In the test seen in [Listing 1.7](#listing-1-7) we are making a direct call to the Wails API and checking the error returned. The test passes when the error returned is `ErrNotAvailable`.

<a id="listing-1-7"></a>


```go
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
> *source: wailsrun/api_calls_test.go:err-not-available*


> *_Listing 1.7:_ Testing the [`wailsrun.ErrNotAvailable`](https://pkg.go.dev/github.com/markbates/wailsx/wailsrun#ErrNotAvailable) method*


When running the tests outside of a Wails application, the [`wailsrun.ErrNotAvailable`](https://pkg.go.dev/github.com/markbates/wailsx/wailsrun#ErrNotAvailable) method will return the `ErrNotAvailable` error, [Listing 1.6](#listing-1-6).

<a id="listing-1-8"></a>


```shell
$ go test -v -run Test_ErrNotAvailable

testing: warning: no tests to run
PASS
ok  	github.com/markbates/wailsx	0.003s

go: downloading github.com/wailsapp/wails/v2 v2.8.0
go: downloading github.com/markbates/safe v1.1.0
go: downloading github.com/stretchr/testify v1.9.0
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading github.com/pmezard/go-difflib v1.0.0
go: downloading gopkg.in/yaml.v3 v3.0.1
go: downloading github.com/leaanthony/slicer v1.6.0
go: downloading github.com/leaanthony/u v1.1.1

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.8:_ Testing the [`wailsrun.ErrNotAvailable`](https://pkg.go.dev/github.com/markbates/wailsx/wailsrun#ErrNotAvailable) method output.*


If the tests are run in a Wails application, using one of the known build tags, the [`wailsrun.ErrNotAvailable`](https://pkg.go.dev/github.com/markbates/wailsx/wailsrun#ErrNotAvailable) method will call the actual Wails API method, [Listing 1.4](#listing-1-4). The result is a call to [`log.Fatal`](https://pkg.go.dev/log#Fatal) because we don't have a valid Wails context.

<a id="listing-1-9"></a>


```shell
$ go test -v -run Test_ErrNotAvailable -tags wails

testing: warning: no tests to run
PASS
ok  	github.com/markbates/wailsx	0.004s

go: downloading github.com/markbates/safe v1.1.0
go: downloading github.com/wailsapp/wails/v2 v2.8.0
go: downloading github.com/stretchr/testify v1.9.0
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading github.com/pmezard/go-difflib v1.0.0
go: downloading gopkg.in/yaml.v3 v3.0.1
go: downloading github.com/leaanthony/u v1.1.1
go: downloading github.com/leaanthony/slicer v1.6.0

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.9:_ Testing the [`wailsrun.ErrNotAvailable`](https://pkg.go.dev/github.com/markbates/wailsx/wailsrun#ErrNotAvailable) method output in `production`.*


---

# <a id="heading-7"></a><toc-level>1.3</toc-level> - Clipboard

## <a id="heading-8"></a><toc-level>1.3.1</toc-level> - The <code>ClipboardManager</code> Interface

<a id="figure-1-3"></a>


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

> *_Figure 1.3:_ The [`clipx.ClipboardManager`](https://pkg.go.dev/github.com/markbates/wailsx/clipx#ClipboardManager) interface*


---

# <a id="heading-9"></a><toc-level>1.4</toc-level> - Dialogs

## <a id="heading-10"></a><toc-level>1.4.1</toc-level> - The <code>DialogManager</code> Interface

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

> *_Listing 1.10:_ The [`dialogx.DialogManager`](https://pkg.go.dev/github.com/markbates/wailsx/dialogx#DialogManager) interface*


---

# <a id="heading-11"></a><toc-level>1.5</toc-level> - Events

## <a id="heading-12"></a><toc-level>1.5.1</toc-level> - The <code>EventManager</code> Interface

<a id="listing-1-11"></a>


```shell
$ go doc github.com/markbates/wailsx/eventx.EventManager

package eventx // import "github.com/markbates/wailsx/eventx"

type EventManager interface {
	EventsEmit(ctx context.Context, event string, args ...any) (err error)
	EventsOff(ctx context.Context, name string, additional ...string) error
	EventsOffAll(ctx context.Context) error
	EventsOn(ctx context.Context, name string, callback CallbackFn) (CancelFn, error)
	EventsOnMultiple(ctx context.Context, name string, callback CallbackFn, counter int) (CancelFn, error)
	EventsOnce(ctx context.Context, name string, callback CallbackFn) (CancelFn, error)
}

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.11:_ The [`eventx.EventManager`](https://pkg.go.dev/github.com/markbates/wailsx/eventx#EventManager) interface*


## <a id="heading-13"></a><toc-level>1.5.2</toc-level> - The <code>Manager</code> Type

<a id="listing-1-12"></a>


```shell
$ go doc github.com/markbates/wailsx/eventx.Manager

package eventx // import "github.com/markbates/wailsx/eventx"

type Manager struct {
	DisableWildcardEmits bool
	DisableStateData     bool

	EventsEmitFn       func(ctx context.Context, name string, data ...any) error
	EventsOffAllFn     func(ctx context.Context) error
	EventsOffFn        func(ctx context.Context, name string, additional ...string) error
	EventsOnFn         func(ctx context.Context, name string, callback CallbackFn) (CancelFn, error)
	EventsOnMultipleFn func(ctx context.Context, name string, callback CallbackFn, counter int) (CancelFn, error)
	EventsOnceFn       func(ctx context.Context, name string, callback CallbackFn) (CancelFn, error)

	NowFn func() time.Time

	// Has unexported fields.
}

func NewManager() *Manager
func NopManager() *Manager
func (em *Manager) EventsEmit(ctx context.Context, event string, args ...any) (err error)
func (em *Manager) EventsOff(ctx context.Context, name string, additional ...string) error
func (em *Manager) EventsOffAll(ctx context.Context) error
func (em *Manager) EventsOn(ctx context.Context, name string, callback CallbackFn) (CancelFn, error)
func (em *Manager) EventsOnMultiple(ctx context.Context, name string, callback CallbackFn, counter int) (CancelFn, error)
func (em *Manager) EventsOnce(ctx context.Context, name string, callback CallbackFn) (CancelFn, error)
func (em *Manager) MarshalJSON() ([]byte, error)
func (em *Manager) Now() time.Time
func (em *Manager) PluginName() string
func (em *Manager) StateData(ctx context.Context) (statedata.Data[*EventsData], error)
func (em *Manager) WithPlugins(fn plugins.FeederFn) error

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.12:_ The [`eventx.Manager`](https://pkg.go.dev/github.com/markbates/wailsx/eventx#Manager) type*


### <a id="heading-14"></a><toc-level>1.5.2.1</toc-level> - Creating a New Manager

<a id="listing-1-13"></a>


```shell
$ go doc github.com/markbates/wailsx/eventx.NewManager

package eventx // import "github.com/markbates/wailsx/eventx"

func NewManager() *Manager

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.13:_ The [`eventx.NewManager`](https://pkg.go.dev/github.com/markbates/wailsx/eventx#NewManager) function*


<a id="listing-1-14"></a>


```shell
$ go doc github.com/markbates/wailsx/eventx.NopManager

package eventx // import "github.com/markbates/wailsx/eventx"

func NopManager() *Manager
    NopManager returns a new Manager with all the functions set to no-ops This
    is useful for testing. The NowFn is set to wailstest.NowTime

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.14:_ The [`eventx.NopManager`](https://pkg.go.dev/github.com/markbates/wailsx/eventx#NopManager) function*


## <a id="heading-15"></a><toc-level>1.5.3</toc-level> - The <code>CallbackFn</code> Type

<a id="listing-1-15"></a>


```shell
$ go doc github.com/markbates/wailsx/wailsrun.CallbackFn

package wailsrun // import "github.com/markbates/wailsx/wailsrun"

type CallbackFn func(data ...any) error

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.15:_ The [`eventx.CallbackFn`](https://pkg.go.dev/github.com/markbates/wailsx/eventx#CallbackFn) type*


## <a id="heading-16"></a><toc-level>1.5.4</toc-level> - The <code>CancelFn</code> Type

<a id="listing-1-16"></a>


```shell
$ go doc github.com/markbates/wailsx/wailsrun.CancelFn

package wailsrun // import "github.com/markbates/wailsx/wailsrun"

type CancelFn func() error

func EventsOn(ctx context.Context, event string, callback CallbackFn) (CancelFn, error)
func EventsOnMultiple(ctx context.Context, event string, callback CallbackFn, counter int) (CancelFn, error)
func EventsOnce(ctx context.Context, event string, callback CallbackFn) (CancelFn, error)

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.16:_ The [`eventx.CancelFn`](https://pkg.go.dev/github.com/markbates/wailsx/eventx#CancelFn) type*


---

# <a id="heading-17"></a><toc-level>1.6</toc-level> - Messages

## <a id="heading-18"></a><toc-level>1.6.1</toc-level> - The <code>Messenger</code> Interface

<a id="listing-1-17"></a>


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

> *_Listing 1.17:_ The [`msgx.Messenger`](https://pkg.go.dev/github.com/markbates/wailsx/eventx/msgx#Messenger) interface*


## <a id="heading-19"></a><toc-level>1.6.2</toc-level> - The <code>ErrorMessenger</code> Interface

<a id="listing-1-18"></a>


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

> *_Listing 1.18:_ The [`msgx.ErrorMessenger`](https://pkg.go.dev/github.com/markbates/wailsx/eventx/msgx#ErrorMessenger) interface*


---

# <a id="heading-20"></a><toc-level>1.7</toc-level> - Logging

## <a id="heading-21"></a><toc-level>1.7.1</toc-level> - The <code>WailsLogger</code> Interface

<a id="listing-1-19"></a>


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

> *_Listing 1.19:_ The [`logx.WailsLogger`](https://pkg.go.dev/github.com/markbates/wailsx/logx#WailsLogger) interface*


---

# <a id="heading-22"></a><toc-level>1.8</toc-level> - Menus

## <a id="heading-23"></a><toc-level>1.8.1</toc-level> - The <code>MenuManager</code> Interface

<a id="listing-1-20"></a>


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

> *_Listing 1.20:_ The [`menux.MenuManager`](https://pkg.go.dev/github.com/markbates/wailsx/menux#MenuManager) interface*


---

# <a id="heading-24"></a><toc-level>1.9</toc-level> - State Data

## <a id="heading-25"></a><toc-level>1.9.1</toc-level> - The <code>DataProvider</code> Interface

<a id="listing-1-21"></a>


```shell
$ go doc github.com/markbates/wailsx/statedata.DataProvider

package statedata // import "github.com/markbates/wailsx/statedata"

type DataProvider[T any] interface {
	StateData(ctx context.Context) (Data[T], error)
}

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.21:_ The [`statedata.DataProvider`](https://pkg.go.dev/github.com/markbates/wailsx/statedata#DataProvider) interface*


## <a id="heading-26"></a><toc-level>1.9.2</toc-level> - The <code>Data</code> Interface

<a id="listing-1-22"></a>


```shell
$ go doc github.com/markbates/wailsx/statedata.Data

package statedata // import "github.com/markbates/wailsx/statedata"

type Data[T any] struct {
	Name string `json:"name,omitempty"` // name of the data
	Data T      `json:"data,omitempty"` // data for the state
}

func (sd Data[T]) PluginName() string
func (sd Data[T]) StateData(ctx context.Context) (Data[T], error)

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Listing 1.22:_ The [`statedata.Data`](https://pkg.go.dev/github.com/markbates/wailsx/statedata#Data) interface*


---

# <a id="heading-27"></a><toc-level>1.10</toc-level> - Window Management

## <a id="heading-28"></a><toc-level>1.10.1</toc-level> - The <code>WindowManager</code> Interface

<a id="listing-1-23"></a>


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

> *_Listing 1.23:_ The [`windowx.WindowManager`](https://pkg.go.dev/github.com/markbates/wailsx/windowx#WindowManager) interface*


## <a id="heading-29"></a><toc-level>1.10.2</toc-level> - The <code>MaximiseManager</code> Interface

<a id="listing-1-24"></a>


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

> *_Listing 1.24:_ The [`windowx.MaximiseManager`](https://pkg.go.dev/github.com/markbates/wailsx/windowx#MaximiseManager) interface*


## <a id="heading-30"></a><toc-level>1.10.3</toc-level> - The <code>PositionManager</code> Interface

<a id="listing-1-25"></a>


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

> *_Listing 1.25:_ The [`windowx.PositionManager`](https://pkg.go.dev/github.com/markbates/wailsx/windowx#PositionManager) interface*


## <a id="heading-31"></a><toc-level>1.10.4</toc-level> - The <code>ReloadManager</code> Interface

<a id="listing-1-26"></a>


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

> *_Listing 1.26:_ The [`windowx.ReloadManager`](https://pkg.go.dev/github.com/markbates/wailsx/windowx#ReloadManager) interface*


## <a id="heading-32"></a><toc-level>1.10.5</toc-level> - The <code>ThemeManager</code> Interface

<a id="listing-1-27"></a>


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

> *_Listing 1.27:_ The [`windowx.ThemeManager`](https://pkg.go.dev/github.com/markbates/wailsx/windowx#ThemeManager) interface*


## <a id="heading-33"></a><toc-level>1.10.6</toc-level> - The <code>Toggler</code> Interface

<a id="listing-1-28"></a>


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

> *_Listing 1.28:_ The [`windowx.Toggler`](https://pkg.go.dev/github.com/markbates/wailsx/windowx#Toggler) interface*


---

# <a id="heading-34"></a><toc-level>1.11</toc-level> - Using the API

While the [`github.com/markbates/wailsx/wailsrun`](https://pkg.go.dev/github.com/markbates/wailsx/wailsrun) package can be used directly, it is recommended to use the [`github.com/markbates/wailsx.API`](https://pkg.go.dev/github.com/markbates/wailsx.API) type to create a testable, idiomatic wrapper around the Wails API.

## <a id="heading-35"></a><toc-level>1.11.1</toc-level> - The <code>API</code> type

The [`wailsx.API`](https://pkg.go.dev/github.com/markbates/wailsx#API) type is a wrapper around the [`github.com/markbates/wailsx/wailsrun`](https://pkg.go.dev/github.com/markbates/wailsx/wailsrun) package. By default, if the [`wailsx.API`](https://pkg.go.dev/github.com/markbates/wailsx#API) type is `nil`, or `zero` (i.e. `&API{}`), all methods will be will be mapped directly to the `wailsrun` package. This allows you to use the [`wailsx.API`](https://pkg.go.dev/github.com/markbates/wailsx#API) type in your application without having to worry about the [`wailsx.API`](https://pkg.go.dev/github.com/markbates/wailsx#API) being `nil`.

<a id="figure-1-4"></a>


```shell
$ go doc -short github.com/markbates/wailsx.API

type API struct {
	clipx.ClipboardManager
	dialogx.DialogManager
	eventx.EventManager
	logx.WailsLogger
	menux.MenuManager
	windowx.WindowManager

	BrowserOpenURLFn func(ctx context.Context, url string) error
	QuitFn           func(ctx context.Context) error
}

func NewAPI() *API
func NopAPI() *API
func (api *API) BrowserOpenURL(ctx context.Context, url string) error
func (api *API) ClipboardGetText(ctx context.Context) (string, error)
func (api *API) ClipboardSetText(ctx context.Context, text string) error
func (api *API) EventsEmit(ctx context.Context, event string, data ...any) error
func (api *API) EventsOff(ctx context.Context, event string, additional ...string) error
func (api *API) EventsOffAll(ctx context.Context) error
func (api *API) EventsOn(ctx context.Context, event string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error)
func (api *API) EventsOnMultiple(ctx context.Context, event string, callback wailsrun.CallbackFn, counter int) (wailsrun.CancelFn, error)
func (api *API) EventsOnce(ctx context.Context, event string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error)
func (api *API) Hide(ctx context.Context) error
func (api *API) LogDebug(ctx context.Context, message string) error
func (api *API) LogDebugf(ctx context.Context, format string, args ...any) error
func (api *API) LogError(ctx context.Context, message string) error
func (api *API) LogErrorf(ctx context.Context, format string, args ...any) error
func (api *API) LogFatal(ctx context.Context, message string) error
func (api *API) LogFatalf(ctx context.Context, format string, args ...any) error
func (api *API) LogInfo(ctx context.Context, message string) error
func (api *API) LogInfof(ctx context.Context, format string, args ...any) error
func (api *API) LogPrint(ctx context.Context, message string) error
func (api *API) LogPrintf(ctx context.Context, format string, args ...any) error
func (api *API) LogSetLogLevel(ctx context.Context, level logger.LogLevel) error
func (api *API) LogTrace(ctx context.Context, message string) error
func (api *API) LogTracef(ctx context.Context, format string, args ...any) error
func (api *API) LogWarning(ctx context.Context, message string) error
func (api *API) LogWarningf(ctx context.Context, format string, args ...any) error
func (api *API) MenuSetApplicationMenu(ctx context.Context, menu *menu.Menu) error
func (api *API) MenuUpdateApplicationMenu(ctx context.Context) error
func (api *API) MessageDialog(ctx context.Context, opts wailsrun.MessageDialogOptions) (string, error)
func (api *API) OpenDirectoryDialog(ctx context.Context, opts wailsrun.OpenDialogOptions) (string, error)
func (api *API) OpenFileDialog(ctx context.Context, opts wailsrun.OpenDialogOptions) (string, error)
func (api *API) OpenMultipleFilesDialog(ctx context.Context, opts wailsrun.OpenDialogOptions) ([]string, error)
func (api *API) Quit(ctx context.Context) error
func (api *API) SaveFileDialog(ctx context.Context, opts wailsrun.SaveDialogOptions) (string, error)
func (api *API) ScreenGetAll(ctx context.Context) ([]wailsrun.Screen, error)
func (api *API) Show(ctx context.Context) error
func (api *API) StateData(ctx context.Context) (statedata.Data[*APIData], error)
func (api *API) WindowCenter(ctx context.Context) error
func (api *API) WindowExecJS(ctx context.Context, js string) error
func (api *API) WindowFullscreen(ctx context.Context) error
func (api *API) WindowGetPosition(ctx context.Context) (int, int, error)
func (api *API) WindowGetSize(ctx context.Context) (int, int, error)
func (api *API) WindowHide(ctx context.Context) error
func (api *API) WindowIsFullscreen(ctx context.Context) (bool, error)
func (api *API) WindowIsMaximised(ctx context.Context) (bool, error)
func (api *API) WindowIsMinimised(ctx context.Context) (bool, error)
func (api *API) WindowIsNormal(ctx context.Context) (bool, error)
func (api *API) WindowMaximise(ctx context.Context) error
func (api *API) WindowMinimise(ctx context.Context) error
func (api *API) WindowPrint(ctx context.Context) error
func (api *API) WindowReload(ctx context.Context) error
func (api *API) WindowReloadApp(ctx context.Context) error
func (api *API) WindowSetAlwaysOnTop(ctx context.Context, b bool) error
func (api *API) WindowSetBackgroundColour(ctx context.Context, R, G, B, A uint8) error
func (api *API) WindowSetDarkTheme(ctx context.Context) error
func (api *API) WindowSetLightTheme(ctx context.Context) error
func (api *API) WindowSetMaxSize(ctx context.Context, width int, height int) error
func (api *API) WindowSetMinSize(ctx context.Context, width int, height int) error
func (api *API) WindowSetPosition(ctx context.Context, x int, y int) error
func (api *API) WindowSetSize(ctx context.Context, width int, height int) error
func (api *API) WindowSetSystemDefaultTheme(ctx context.Context) error
func (api *API) WindowSetTitle(ctx context.Context, title string) error
func (api *API) WindowShow(ctx context.Context) error
func (api *API) WindowToggleMaximise(ctx context.Context) error
func (api *API) WindowUnfullscreen(ctx context.Context) error
func (api *API) WindowUnmaximise(ctx context.Context) error
func (api *API) WindowUnminimise(ctx context.Context) error

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Figure 1.4:_ The [`wailsx.API`](https://pkg.go.dev/github.com/markbates/wailsx#API) type*


The [`wailsx.NewAPI`](https://pkg.go.dev/github.com/markbates/wailsx#NewAPI) function can be used to create a new [`wailsx.API`](https://pkg.go.dev/github.com/markbates/wailsx#API) type. This function will populate the [`wailsx.API`](https://pkg.go.dev/github.com/markbates/wailsx#API) type with implementations of its embedded interfaces. For example, using [`eventx.NewManager`](https://pkg.go.dev/github.com/markbates/wailsx/eventx#NewManager) to create a new [`eventx.Manager`](https://pkg.go.dev/github.com/markbates/wailsx/eventx#Manager) that will fill the needed [`eventx.EventManager`](https://pkg.go.dev/github.com/markbates/wailsx/eventx#EventManager) in the [`wailsx.API`](https://pkg.go.dev/github.com/markbates/wailsx#API) type.

<a id="figure-1-5"></a>


```shell
$ go doc -short github.com/markbates/wailsx.NewAPI

func NewAPI() *API
    NewAPI returns a new API with all the functions, and interfaces, set to
    their default implementations.

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Figure 1.5:_ The [`wailsx.NewAPI`](https://pkg.go.dev/github.com/markbates/wailsx#NewAPI) function*


## <a id="heading-36"></a><toc-level>1.11.2</toc-level> - <code>Nil</code> API Calls

<a id="figure-1-6"></a>


```go
func Test_Nil_API_Call(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var api *API

	ctx := context.Background()

	err := api.Show(ctx)
	r.Error(err)

	exp := wailsrun.ErrNotAvailable("Show")
	r.Equal(exp, err)
}
```
> *source: doc_test.go:nil-api*


> *_Figure 1.6:_ Calling methods on a `nil` [`wailsx.API`](https://pkg.go.dev/github.com/markbates/wailsx#API).*


<a id="figure-1-7"></a>


```shell
$ go test -v -run Test_Nil_API_Call

=== RUN   Test_Nil_API_Call
=== PAUSE Test_Nil_API_Call
=== CONT  Test_Nil_API_Call
--- PASS: Test_Nil_API_Call (0.00s)
PASS
ok  	github.com/markbates/wailsx	0.004s

go: downloading github.com/markbates/safe v1.1.0
go: downloading github.com/wailsapp/wails/v2 v2.8.0
go: downloading github.com/stretchr/testify v1.9.0
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading github.com/pmezard/go-difflib v1.0.0
go: downloading gopkg.in/yaml.v3 v3.0.1
go: downloading github.com/leaanthony/slicer v1.6.0
go: downloading github.com/leaanthony/u v1.1.1

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Figure 1.7:_ Running the test for calling methods on a `nil` [`wailsx.API`](https://pkg.go.dev/github.com/markbates/wailsx#API).*


<a id="figure-1-8"></a>


```shell
$ go test -v -run Test_Nil_API_Call -tags wails

=== RUN   Test_Nil_API_Call
=== PAUSE Test_Nil_API_Call
=== CONT  Test_Nil_API_Call
2024/03/13 03:49:45 cannot call 'github.com/wailsapp/wails/v2/pkg/runtime.Show': An invalid context was passed. This method requires the specific context given in the lifecycle hooks:
https://wails.io/docs/reference/runtime/intro
exit status 1
FAIL	github.com/markbates/wailsx	0.004s

go: downloading github.com/markbates/safe v1.1.0
go: downloading github.com/wailsapp/wails/v2 v2.8.0
go: downloading github.com/stretchr/testify v1.9.0
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading github.com/pmezard/go-difflib v1.0.0
go: downloading gopkg.in/yaml.v3 v3.0.1
go: downloading github.com/leaanthony/slicer v1.6.0
go: downloading github.com/leaanthony/u v1.1.1

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Figure 1.8:_ Running the test for calling methods on a `nil` [`wailsx.API`](https://pkg.go.dev/github.com/markbates/wailsx#API) in production mode.*


## <a id="heading-37"></a><toc-level>1.11.3</toc-level> - <code>Nop</code> API Calls

<a id="figure-1-9"></a>


```shell
$ go doc -short github.com/markbates/wailsx.NopAPI

func NopAPI() *API
    NopAPI returns a new API with all the functions, and interfaces, set to
    no-ops. This is useful for testing.

--------------------------------------------------------------------------------
Go Version: go1.22.0

```

> *_Figure 1.9:_ The [`wailsx.NopAPI`](https://pkg.go.dev/github.com/markbates/wailsx#NopAPI) function*


<a id="figure-1-10"></a>


```go
func Test_Nop_API_Call(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	api := NopAPI()

	ctx := context.Background()

	err := api.Show(ctx)
	r.NoError(err)
}
```
> *source: doc_test.go:nop-api*


> *_Figure 1.10:_ Calling methods on a `nop` [`wailsx.API`](https://pkg.go.dev/github.com/markbates/wailsx#API).*


