package wailsx

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/markbates/wailsx/eventx"
	"github.com/markbates/wailsx/windowx"
	"github.com/stretchr/testify/require"
)

type basicPlugin string

func (b basicPlugin) PluginName() string {
	return string(b)
}

var _ PluginDataProvider = intProvider(0)

type intProvider int

func (i intProvider) PluginName() string {
	return fmt.Sprintf("%T: %d", i, int(i))
}

func (i intProvider) StateData(ctx context.Context) (any, error) {
	return int(i), nil
}

type stringProvider string

func (s stringProvider) PluginName() string {
	return fmt.Sprintf("%T: %s", s, string(s))
}

func (s stringProvider) StateData(ctx context.Context) (any, error) {
	return string(s), nil
}

func Test_NewApp(t *testing.T) {
	// t.Skip()
	t.Parallel()
	r := require.New(t)

	_, err := NewApp("")
	r.Error(err)

	app, err := NewApp("test")
	r.NoError(err)

	r.Equal("test", app.Name)
	r.Equal("*wailsx.App: test", app.PluginName())

	r.NotNil(app.API)
	r.NotNil(app.ClipboardManager)
	r.NotNil(app.DialogManager)
	r.NotNil(app.EventManager)
	r.NotNil(app.MenuManager)
	r.NotNil(app.SaveFn)
	r.NotNil(app.WailsLogger)
	r.NotNil(app.WindowManager)
	r.Nil(app.BrowserOpenURLFn)
	r.Nil(app.QuitFn)
	r.Nil(app.StartupFn)
	r.Nil(app.ShutdownFn)
	r.Nil(app.DomReadyFn)
	r.Nil(app.BeforeCloseFn)
}

func Test_NopApp(t *testing.T) {
	// t.Skip()
	t.Parallel()
	r := require.New(t)

	_, err := NopApp("")
	r.Error(err)

	app, err := NopApp("test")
	r.NoError(err)

	r.Equal("test", app.Name)
	r.NotNil(app.API)
	r.NotNil(app.ClipboardManager)
	r.NotNil(app.DialogManager)
	r.NotNil(app.EventManager)
	r.NotNil(app.WailsLogger)
	r.NotNil(app.MenuManager)
	r.NotNil(app.WindowManager)
	r.NotNil(app.BrowserOpenURLFn)
	r.NotNil(app.QuitFn)
	r.NotNil(app.SaveFn)
	r.NotNil(app.StartupFn)
	r.NotNil(app.ShutdownFn)
	r.NotNil(app.DomReadyFn)
	r.NotNil(app.BeforeCloseFn)
}

func Test_App_StateData(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ctx := context.Background()

	var app *App
	_, err := app.StateData(ctx)
	r.Error(err)

	app, err = NopApp("test",
		basicPlugin("basic"),
		intProvider(42),
		stringProvider("hello"),
	)
	r.NoError(err)

	sd, err := app.StateData(ctx)
	r.NoError(err)

	api := sd.API
	r.NotNil(api)

	r.NotNil(api.Events)
	r.NotNil(api.Window)

	b, err := json.Marshal(sd)
	r.NoError(err)

	act := string(b)
	// fmt.Println(act)

	// f, err := os.Create("testdata/app1.json")
	// r.NoError(err)
	// f.WriteString(act)
	// r.NoError(f.Close())

	b, err = os.ReadFile("testdata/app1.json")
	r.NoError(err)

	exp := string(b)
	exp = strings.TrimSpace(exp)

	r.Equal(exp, act)

	const evnt = "event:test"
	cancel, err := app.EventsOn(ctx, evnt, func(data ...any) error {
		return nil
	})
	r.NoError(err)
	r.NotNil(cancel)
	defer cancel()

	err = app.EventsEmit(ctx, evnt, "my data")
	r.NoError(err)

	err = app.WindowSetBackgroundColour(ctx, 1, 2, 3, 4)
	r.NoError(err)

	err = app.WindowSetPosition(ctx, 10, 20)
	r.NoError(err)

	err = app.WindowSetSize(ctx, 100, 200)
	r.NoError(err)

	err = app.WindowSetSystemDefaultTheme(ctx)
	r.NoError(err)

	err = app.WindowMaximise(ctx)
	r.NoError(err)

	sd, err = app.StateData(ctx)
	r.NoError(err)

	b, err = json.Marshal(sd)
	r.NoError(err)

	act = string(b)

	// fmt.Println(act)
	// f, err := os.Create("testdata/app2.json")
	// r.NoError(err)
	// f.WriteString(act)
	// r.NoError(f.Close())

	b, err = os.ReadFile("testdata/app2.json")
	r.NoError(err)

	exp = string(b)
	exp = strings.TrimSpace(exp)

	r.Equal(exp, act)
}

func Test_App_RestoreAPP(t *testing.T) {
	// t.Skip()
	t.Parallel()
	r := require.New(t)

	ctx := context.Background()

	var app *App
	err := app.RestoreAPP(ctx, &AppData{})
	r.Error(err)

	app, err = NopApp("test")
	r.NoError(err)
	r.Equal("test", app.Name)

	em := &restoreableEvents{
		EventManager: eventx.NopManager(),
	}
	app.EventManager = em

	wm := &restoreableWindow{
		WindowManager: windowx.NopManager(),
	}
	app.WindowManager = wm

	err = app.RestoreAPP(ctx, &AppData{})
	r.NoError(err)

	data := &AppData{
		AppName: "My App",
		API: &APIData{
			Events: &eventx.EventsData{
				Callbacks: map[string]*eventx.CallbackCounter{
					"event:test": {
						Called: 1,
					},
				},
			},
			Window: &windowx.WindowData{
				ThemeData: &windowx.ThemeData{
					Theme: windowx.THEME_DARK,
				},
			},
		},
	}

	err = app.RestoreAPP(ctx, data)
	r.NoError(err)

	r.Equal("My App", app.Name)
	r.Equal(data.API.Events, em.Data)
	r.Equal(data.API.Window, wm.Data)

}

var _ PluginDataProvider = &restoreablePlugin{}
var _ RestorablePlugin = &restoreablePlugin{}

type restoreablePlugin struct {
	Data int
}

func (r *restoreablePlugin) PluginName() string {
	return fmt.Sprintf("%T", r)
}

func (r *restoreablePlugin) StateData(ctx context.Context) (any, error) {
	if r == nil {
		return nil, fmt.Errorf("plugin is nil")
	}

	return r.Data, nil
}

func (r *restoreablePlugin) RestorePlugin(ctx context.Context, data any) error {
	i, ok := data.(int)
	if !ok {
		return fmt.Errorf("invalid data type: %T", data)
	}
	r.Data = i
	return nil
}

func Test_App_RestoreAPP_Plugins(t *testing.T) {

	t.Parallel()
	r := require.New(t)

	rp := &restoreablePlugin{}

	app, err := NopApp("My App", rp)
	r.NoError(err)
	r.Equal("My App", app.Name)

	data := &AppData{
		AppName: "My App",
		Plugins: map[string]any{
			rp.PluginName(): 42,
		},
	}

	err = app.RestoreAPP(context.Background(), data)
	r.NoError(err)

	r.Equal(42, rp.Data)

}
