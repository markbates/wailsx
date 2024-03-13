package wailsx

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/markbates/wailsx/statedata"
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

func (i intProvider) StateData(ctx context.Context) (statedata.Data[any], error) {
	sd := statedata.Data[any]{
		Name: "int-data",
		Data: int(i),
	}

	return sd, nil
}

type stringProvider string

func (s stringProvider) PluginName() string {
	return fmt.Sprintf("%T: %s", s, string(s))
}

func (s stringProvider) StateData(ctx context.Context) (statedata.Data[any], error) {
	sd := statedata.Data[any]{
		Name: "string-data",
		Data: string(s),
	}

	return sd, nil
}

func Test_NewApp(t *testing.T) {
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
	r.NotNil(app.WailsLogger)
	r.NotNil(app.MenuManager)
	r.NotNil(app.WindowManager)
	r.Nil(app.BrowserOpenURLFn)
	r.Nil(app.QuitFn)
	r.Nil(app.SaveFn)
	r.Nil(app.StartupFn)
	r.Nil(app.ShutdownFn)
	r.Nil(app.DomReadyFn)
	r.Nil(app.BeforeCloseFn)
}

func Test_NopApp(t *testing.T) {
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

	r.Equal(AppStateDataProviderName, sd.Name)

	api := sd.Data.API
	r.NotNil(api)

	r.NotNil(api.Events)
	r.NotNil(api.Window)

	b, err := json.Marshal(sd)
	r.NoError(err)

	act := string(b)
	// fmt.Println(act)

	exp := `{"name":"app","data":{"app_name":"test","api":{"events":{},"window":{"maximiser":{},"positioner":{},"themer":{"background_colour":{}}}},"plugins":{"wailsx.intProvider: 42":{"name":"int-data","data":42},"wailsx.stringProvider: hello":{"name":"string-data","data":"hello"}}}}`

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
	exp = `{"name":"app","data":{"app_name":"test","api":{"events":{"callbacks":{"event:test":{"called":1,"max_calls":0,"off":false}},"emitted":{"*":[{"name":"*","data":[{"data":"my data","event":"event:test","text":"my data","time":"2024-01-01T00:00:00Z"}],"emitted_at":"2024-01-01T00:00:00Z"}],"event:test":[{"name":"event:test","data":[{"data":"my data","event":"event:test","text":"my data","time":"2024-01-01T00:00:00Z"}],"emitted_at":"2024-01-01T00:00:00Z"}]},"caught":{"event:test":[{"name":"event:test","data":[{"data":"my data","event":"event:test","text":"my data","time":"2024-01-01T00:00:00Z"}],"emitted_at":"2024-01-01T00:00:00Z"}]}},"window":{"maximiser":{"is_maximised":true},"positioner":{"x":10,"y":20,"w":100,"h":200},"themer":{"background_colour":{"r":1,"g":2,"b":3,"a":4},"is_system_theme":true}}},"plugins":{"wailsx.intProvider: 42":{"name":"int-data","data":42},"wailsx.stringProvider: hello":{"name":"string-data","data":"hello"}}}}`

	r.Equal(exp, act)
}
