package wailsx

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
	"github.com/wailsapp/wails/v2/pkg/menu"
)

func Test_API_BrowserOpenURL(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	exp := wailsrun.ErrNotAvailable("BrowserOpenURL")

	tcs := []struct {
		name string
		fn   func() error
		err  error
	}{
		{
			name: "with function",
			fn:   func() error { return nil },
		},
		{
			name: "with nil function",
			err:  exp,
		},
		{
			name: "with error",
			fn: func() error {
				return wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func() error {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			api := NewAPI()
			r.NotNil(api)

			if tc.fn != nil {
				api.BrowserOpenURLFn = func(ctx context.Context, url string) error {
					return tc.fn()
				}
			}

			err := api.BrowserOpenURL(ctx, "https://example.com")
			r.Equal(tc.err, err)
		})
	}
}

func Test_API_Quit(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	exp := wailsrun.ErrNotAvailable("Quit")

	tcs := []struct {
		name string
		fn   func() error
		err  error
	}{
		{
			name: "with function",
			fn:   func() error { return nil },
		},
		{
			name: "with nil function",
			err:  exp,
		},
		{
			name: "with error",
			fn: func() error {
				return wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func() error {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			api := NewAPI()
			r.NotNil(api)

			if tc.fn != nil {
				api.QuitFn = func(ctx context.Context) error {
					return tc.fn()
				}
			}

			err := api.Quit(ctx)
			r.Equal(tc.err, err)
		})
	}
}

func Test_API_StateData(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	api := NopAPI()

	ctx := context.Background()

	const (
		maxW = 1200
		maxH = 800
		minW = 20
		minH = 30
		w    = 800
		h    = 600
		x    = 100
		y    = 200
	)

	err := api.WindowSetMaxSize(ctx, maxW, maxH)
	r.NoError(err)

	err = api.WindowSetMinSize(ctx, minW, minH)
	r.NoError(err)

	err = api.WindowSetPosition(ctx, x, y)
	r.NoError(err)

	err = api.WindowSetSize(ctx, w, h)
	r.NoError(err)

	err = api.WindowSetBackgroundColour(ctx, 1, 2, 3, 4)
	r.NoError(err)

	err = api.WindowSetDarkTheme(ctx)
	r.NoError(err)

	err = api.WindowMaximise(ctx)
	r.NoError(err)

	const event = "event:test"

	cancel, err := api.EventsOn(ctx, event, func(data ...any) error {
		r.Len(data, 1)
		r.Equal(42, data[0])
		return nil
	})
	r.NoError(err)
	defer cancel()

	err = api.EventsEmit(ctx, event, 42)
	r.NoError(err)

	sd, err := api.StateData(ctx)
	r.NoError(err)

	r.NotNil(sd.Data)
	r.Equal(APIStateDataProviderName, sd.Name)

	ed := sd.Data.Events
	r.NotNil(ed)

	wd := sd.Data.Window
	r.NotNil(wd)

	r.Equal(h, wd.H)
	r.Equal(maxH, wd.MaxH)
	r.Equal(maxW, wd.MaxW)
	r.Equal(minH, wd.MinH)
	r.Equal(minW, wd.MinW)
	r.Equal(w, wd.W)
	r.Equal(x, wd.X)
	r.Equal(y, wd.Y)

	b, err := json.Marshal(sd)
	r.NoError(err)

	act := string(b)
	act = strings.TrimSpace(act)

	// fmt.Println(act)

	// f, err := os.Create("testdata/api.json")
	// r.NoError(err)
	// f.WriteString(act)
	// r.NoError(f.Close())

	b, err = os.ReadFile("testdata/api.json")
	r.NoError(err)

	exp := string(b)
	exp = strings.TrimSpace(exp)

	r.Equal(exp, act)

}

func Test_Nil_API(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	var api *API

	tcs := []struct {
		name string
		fn   func() error
	}{
		{
			name: "BrowserOpenURL",
			fn:   func() error { return api.BrowserOpenURL(ctx, "") },
		},
		{
			name: "ClipboardGetText",
			fn:   func() error { _, err := api.ClipboardGetText(ctx); return err },
		},
		{
			name: "ClipboardSetText",
			fn:   func() error { return api.ClipboardSetText(ctx, "") },
		},
		{
			name: "EventsEmit",
			fn:   func() error { return api.EventsEmit(ctx, "", 1) },
		},
		{
			name: "EventsOff",
			fn:   func() error { return api.EventsOff(ctx, "", "") },
		},
		{
			name: "EventsOffAll",
			fn:   func() error { return api.EventsOffAll(ctx) },
		},
		{
			name: "EventsOn",
			fn:   func() error { _, err := api.EventsOn(ctx, "", nil); return err },
		},
		{
			name: "EventsOnMultiple",
			fn:   func() error { _, err := api.EventsOnMultiple(ctx, "", nil, 1); return err },
		},
		{
			name: "EventsOnce",
			fn:   func() error { _, err := api.EventsOnce(ctx, "", nil); return err },
		},
		{
			name: "Hide",
			fn:   func() error { return api.Hide(ctx) },
		},
		{
			name: "LogDebug",
			fn:   func() error { return api.LogDebug(ctx, "") },
		},
		{
			name: "LogDebugf",
			fn:   func() error { return api.LogDebugf(ctx, "") },
		},
		{
			name: "LogError",
			fn:   func() error { return api.LogError(ctx, "") },
		},
		{
			name: "LogErrorf",
			fn:   func() error { return api.LogErrorf(ctx, "") },
		},
		{
			name: "LogFatal",
			fn:   func() error { return api.LogFatal(ctx, "") },
		},
		{
			name: "LogFatalf",
			fn:   func() error { return api.LogFatalf(ctx, "") },
		},
		{
			name: "LogInfo",
			fn:   func() error { return api.LogInfo(ctx, "") },
		},
		{
			name: "LogInfof",
			fn:   func() error { return api.LogInfof(ctx, "") },
		},
		{
			name: "LogPrint",
			fn:   func() error { return api.LogPrint(ctx, "") },
		},
		{
			name: "LogPrintf",
			fn:   func() error { return api.LogPrintf(ctx, "") },
		},
		{
			name: "LogSetLogLevel",
			fn:   func() error { return api.LogSetLogLevel(ctx, 0) },
		},
		{
			name: "LogTrace",
			fn:   func() error { return api.LogTrace(ctx, "") },
		},
		{
			name: "LogTracef",
			fn:   func() error { return api.LogTracef(ctx, "") },
		},
		{
			name: "LogWarning",
			fn:   func() error { return api.LogWarning(ctx, "") },
		},
		{
			name: "LogWarningf",
			fn:   func() error { return api.LogWarningf(ctx, "") },
		},
		{
			name: "MenuSetApplicationMenu",
			fn:   func() error { return api.MenuSetApplicationMenu(ctx, nil) },
		},
		{
			name: "MenuUpdateApplicationMenu",
			fn:   func() error { return api.MenuUpdateApplicationMenu(ctx) },
		},
		{
			name: "MessageDialog",
			fn:   func() error { _, err := api.MessageDialog(ctx, wailsrun.MessageDialogOptions{}); return err },
		},
		{
			name: "OpenDirectoryDialog",
			fn:   func() error { _, err := api.OpenDirectoryDialog(ctx, wailsrun.OpenDialogOptions{}); return err },
		},
		{
			name: "OpenFileDialog",
			fn:   func() error { _, err := api.OpenFileDialog(ctx, wailsrun.OpenDialogOptions{}); return err },
		},
		{
			name: "OpenMultipleFilesDialog",
			fn:   func() error { _, err := api.OpenMultipleFilesDialog(ctx, wailsrun.OpenDialogOptions{}); return err },
		},
		{
			name: "Quit",
			fn:   func() error { return api.Quit(ctx) },
		},
		{
			name: "SaveFileDialog",
			fn:   func() error { _, err := api.SaveFileDialog(ctx, wailsrun.SaveDialogOptions{}); return err },
		},
		{
			name: "Show",
			fn:   func() error { return api.Show(ctx) },
		},
		{
			name: "WindowCenter",
			fn:   func() error { return api.WindowCenter(ctx) },
		},
		{
			name: "WindowExecJS",
			fn:   func() error { return api.WindowExecJS(ctx, "") },
		},
		{
			name: "WindowFullscreen",
			fn:   func() error { return api.WindowFullscreen(ctx) },
		},
		{
			name: "WindowGetPosition",
			fn:   func() error { _, _, err := api.WindowGetPosition(ctx); return err },
		},
		{
			name: "WindowGetSize",
			fn:   func() error { _, _, err := api.WindowGetSize(ctx); return err },
		},
		{
			name: "WindowHide",
			fn:   func() error { return api.WindowHide(ctx) },
		},
		{
			name: "WindowIsFullscreen",
			fn:   func() error { _, err := api.WindowIsFullscreen(ctx); return err },
		},
		{
			name: "WindowIsMaximised",
			fn:   func() error { _, err := api.WindowIsMaximised(ctx); return err },
		},
		{
			name: "WindowIsMinimised",
			fn:   func() error { _, err := api.WindowIsMinimised(ctx); return err },
		},
		{
			name: "WindowIsNormal",
			fn:   func() error { _, err := api.WindowIsNormal(ctx); return err },
		},
		{
			name: "WindowMaximise",
			fn:   func() error { return api.WindowMaximise(ctx) },
		},
		{
			name: "WindowMinimise",
			fn:   func() error { return api.WindowMinimise(ctx) },
		},
		{
			name: "WindowPrint",
			fn:   func() error { return api.WindowPrint(ctx) },
		},
		{
			name: "WindowReload",
			fn:   func() error { return api.WindowReload(ctx) },
		},
		{
			name: "WindowReloadApp",
			fn:   func() error { return api.WindowReloadApp(ctx) },
		},
		{
			name: "WindowSetAlwaysOnTop",
			fn:   func() error { return api.WindowSetAlwaysOnTop(ctx, false) },
		},
		{
			name: "WindowSetBackgroundColour",
			fn:   func() error { return api.WindowSetBackgroundColour(ctx, 1, 2, 3, 4) },
		},
		{
			name: "WindowSetDarkTheme",
			fn:   func() error { return api.WindowSetDarkTheme(ctx) },
		},
		{
			name: "WindowSetLightTheme",
			fn:   func() error { return api.WindowSetLightTheme(ctx) },
		},
		{
			name: "WindowSetMaxSize",
			fn:   func() error { return api.WindowSetMaxSize(ctx, 1, 2) },
		},
		{
			name: "WindowSetMinSize",
			fn:   func() error { return api.WindowSetMinSize(ctx, 1, 2) },
		},
		{
			name: "WindowSetPosition",
			fn:   func() error { return api.WindowSetPosition(ctx, 1, 2) },
		},
		{
			name: "WindowSetSize",
			fn:   func() error { return api.WindowSetSize(ctx, 1, 2) },
		},
		{
			name: "WindowSetSystemDefaultTheme",
			fn:   func() error { return api.WindowSetSystemDefaultTheme(ctx) },
		},
		{
			name: "WindowSetTitle",
			fn:   func() error { return api.WindowSetTitle(ctx, "") },
		},
		{
			name: "WindowShow",
			fn:   func() error { return api.WindowShow(ctx) },
		},
		{
			name: "WindowToggleMaximise",
			fn:   func() error { return api.WindowToggleMaximise(ctx) },
		},
		{
			name: "WindowUnfullscreen",
			fn:   func() error { return api.WindowUnfullscreen(ctx) },
		},
		{
			name: "WindowUnmaximise",
			fn:   func() error { return api.WindowUnmaximise(ctx) },
		},
		{
			name: "WindowUnminimise",
			fn:   func() error { return api.WindowUnminimise(ctx) },
		},
		{
			name: "ScreenGetAll",
			fn:   func() error { _, err := api.ScreenGetAll(ctx); return err },
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			err := tc.fn()
			r.Error(err)
			exp := wailsrun.ErrNotAvailable(tc.name)
			r.Equal(exp, err)
		})
	}
}

func Test_NopAPI(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	api := NopAPI()
	r.NotNil(api)
	r.NotNil(api.ClipboardManager)
	r.NotNil(api.DialogManager)
	r.NotNil(api.EventManager)
	r.NotNil(api.MenuManager)
	r.NotNil(api.WailsLogger)
	r.NotNil(api.WindowManager)
	r.NotNil(api.BrowserOpenURLFn)
	r.NotNil(api.QuitFn)

	ctx := context.Background()

	tcs := []struct {
		name string
		fn   func() error
	}{
		{
			name: "BrowserOpenURL",
			fn:   func() error { return api.BrowserOpenURL(ctx, "") },
		},
		{
			name: "ClipboardGetText",
			fn:   func() error { _, err := api.ClipboardGetText(ctx); return err },
		},
		{
			name: "ClipboardSetText",
			fn:   func() error { return api.ClipboardSetText(ctx, "") },
		},
		{
			name: "EventsEmit",
			fn:   func() error { return api.EventsEmit(ctx, "", 1) },
		},
		{
			name: "EventsOff",
			fn:   func() error { return api.EventsOff(ctx, "", "") },
		},
		{
			name: "EventsOffAll",
			fn:   func() error { return api.EventsOffAll(ctx) },
		},
		{
			name: "EventsOn",
			fn:   func() error { _, err := api.EventsOn(ctx, "", nil); return err },
		},
		{
			name: "EventsOnMultiple",
			fn:   func() error { _, err := api.EventsOnMultiple(ctx, "", nil, 1); return err },
		},
		{
			name: "EventsOnce",
			fn:   func() error { _, err := api.EventsOnce(ctx, "", nil); return err },
		},
		{
			name: "Hide",
			fn:   func() error { return api.Hide(ctx) },
		},
		{
			name: "LogDebug",
			fn:   func() error { return api.LogDebug(ctx, "") },
		},
		{
			name: "LogDebugf",
			fn:   func() error { return api.LogDebugf(ctx, "") },
		},
		{
			name: "LogError",
			fn:   func() error { return api.LogError(ctx, "") },
		},
		{
			name: "LogErrorf",
			fn:   func() error { return api.LogErrorf(ctx, "") },
		},
		{
			name: "LogFatal",
			fn:   func() error { return api.LogFatal(ctx, "") },
		},
		{
			name: "LogFatalf",
			fn:   func() error { return api.LogFatalf(ctx, "") },
		},
		{
			name: "LogInfo",
			fn:   func() error { return api.LogInfo(ctx, "") },
		},
		{
			name: "LogInfof",
			fn:   func() error { return api.LogInfof(ctx, "") },
		},
		{
			name: "LogPrint",
			fn:   func() error { return api.LogPrint(ctx, "") },
		},
		{
			name: "LogPrintf",
			fn:   func() error { return api.LogPrintf(ctx, "") },
		},
		{
			name: "LogSetLogLevel",
			fn:   func() error { return api.LogSetLogLevel(ctx, 0) },
		},
		{
			name: "LogTrace",
			fn:   func() error { return api.LogTrace(ctx, "") },
		},
		{
			name: "LogTracef",
			fn:   func() error { return api.LogTracef(ctx, "") },
		},
		{
			name: "LogWarning",
			fn:   func() error { return api.LogWarning(ctx, "") },
		},
		{
			name: "LogWarningf",
			fn:   func() error { return api.LogWarningf(ctx, "") },
		},
		{
			name: "MenuSetApplicationMenu",
			fn:   func() error { return api.MenuSetApplicationMenu(ctx, &menu.Menu{}) },
		},
		{
			name: "MenuUpdateApplicationMenu",
			fn:   func() error { return api.MenuUpdateApplicationMenu(ctx) },
		},
		{
			name: "MessageDialog",
			fn:   func() error { _, err := api.MessageDialog(ctx, wailsrun.MessageDialogOptions{}); return err },
		},
		{
			name: "OpenDirectoryDialog",
			fn:   func() error { _, err := api.OpenDirectoryDialog(ctx, wailsrun.OpenDialogOptions{}); return err },
		},
		{
			name: "OpenFileDialog",
			fn:   func() error { _, err := api.OpenFileDialog(ctx, wailsrun.OpenDialogOptions{}); return err },
		},
		{
			name: "OpenMultipleFilesDialog",
			fn:   func() error { _, err := api.OpenMultipleFilesDialog(ctx, wailsrun.OpenDialogOptions{}); return err },
		},
		{
			name: "Quit",
			fn:   func() error { return api.Quit(ctx) },
		},
		{
			name: "SaveFileDialog",
			fn:   func() error { _, err := api.SaveFileDialog(ctx, wailsrun.SaveDialogOptions{}); return err },
		},
		{
			name: "Show",
			fn:   func() error { return api.Show(ctx) },
		},
		{
			name: "WindowCenter",
			fn:   func() error { return api.WindowCenter(ctx) },
		},
		{
			name: "WindowExecJS",
			fn:   func() error { return api.WindowExecJS(ctx, "") },
		},
		{
			name: "WindowFullscreen",
			fn:   func() error { return api.WindowFullscreen(ctx) },
		},
		{
			name: "WindowGetPosition",
			fn:   func() error { _, _, err := api.WindowGetPosition(ctx); return err },
		},
		{
			name: "WindowGetSize",
			fn:   func() error { _, _, err := api.WindowGetSize(ctx); return err },
		},
		{
			name: "WindowHide",
			fn:   func() error { return api.WindowHide(ctx) },
		},
		{
			name: "WindowIsFullscreen",
			fn:   func() error { _, err := api.WindowIsFullscreen(ctx); return err },
		},
		{
			name: "WindowIsMaximised",
			fn:   func() error { _, err := api.WindowIsMaximised(ctx); return err },
		},
		{
			name: "WindowIsMinimised",
			fn:   func() error { _, err := api.WindowIsMinimised(ctx); return err },
		},
		{
			name: "WindowIsNormal",
			fn:   func() error { _, err := api.WindowIsNormal(ctx); return err },
		},
		{
			name: "WindowMaximise",
			fn:   func() error { return api.WindowMaximise(ctx) },
		},
		{
			name: "WindowMinimise",
			fn:   func() error { return api.WindowMinimise(ctx) },
		},
		{
			name: "WindowPrint",
			fn:   func() error { return api.WindowPrint(ctx) },
		},
		{
			name: "WindowReload",
			fn:   func() error { return api.WindowReload(ctx) },
		},
		{
			name: "WindowReloadApp",
			fn:   func() error { return api.WindowReloadApp(ctx) },
		},
		{
			name: "WindowSetAlwaysOnTop",
			fn:   func() error { return api.WindowSetAlwaysOnTop(ctx, false) },
		},
		{
			name: "WindowSetBackgroundColour",
			fn:   func() error { return api.WindowSetBackgroundColour(ctx, 1, 2, 3, 4) },
		},
		{
			name: "WindowSetDarkTheme",
			fn:   func() error { return api.WindowSetDarkTheme(ctx) },
		},
		{
			name: "WindowSetLightTheme",
			fn:   func() error { return api.WindowSetLightTheme(ctx) },
		},
		{
			name: "WindowSetMaxSize",
			fn:   func() error { return api.WindowSetMaxSize(ctx, 1, 2) },
		},
		{
			name: "WindowSetMinSize",
			fn:   func() error { return api.WindowSetMinSize(ctx, 1, 2) },
		},
		{
			name: "WindowSetPosition",
			fn:   func() error { return api.WindowSetPosition(ctx, 1, 2) },
		},
		{
			name: "WindowSetSize",
			fn:   func() error { return api.WindowSetSize(ctx, 1, 2) },
		},
		{
			name: "WindowSetSystemDefaultTheme",
			fn:   func() error { return api.WindowSetSystemDefaultTheme(ctx) },
		},
		{
			name: "WindowSetTitle",
			fn:   func() error { return api.WindowSetTitle(ctx, "") },
		},
		{
			name: "WindowShow",
			fn:   func() error { return api.WindowShow(ctx) },
		},
		{
			name: "WindowToggleMaximise",
			fn:   func() error { return api.WindowToggleMaximise(ctx) },
		},
		{
			name: "WindowUnfullscreen",
			fn:   func() error { return api.WindowUnfullscreen(ctx) },
		},
		{
			name: "WindowUnmaximise",
			fn:   func() error { return api.WindowUnmaximise(ctx) },
		},
		{
			name: "WindowUnminimise",
			fn:   func() error { return api.WindowUnminimise(ctx) },
		},
		{
			name: "ScreenGetAll",
			fn:   func() error { _, err := api.ScreenGetAll(ctx); return err },
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			err := tc.fn()
			r.NoError(err)
		})
	}
}
