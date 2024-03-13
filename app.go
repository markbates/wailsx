package wailsx

import (
	"context"
	"fmt"
	"sync"

	"github.com/markbates/plugins"
	"github.com/markbates/wailsx/statedata"
	"github.com/markbates/wailsx/wailsrun"
)

var _ wailsrun.API = &App{}
var _ AppStateDataProvider = &App{}

func NewApp(name string, plugins ...plugins.Plugin) (*App, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("name is required")
	}

	app := &App{
		Name:    name,
		API:     NewAPI(),
		Plugins: plugins,
	}

	af := &AppFilesaver{
		App: app,
	}

	app.SaveFn = af.Save

	return app, nil
}

func NopApp(name string, plugins ...plugins.Plugin) (*App, error) {
	app, err := NewApp(name, plugins...)
	if err != nil {
		return nil, err
	}

	app.API = NopAPI()

	app.SaveFn = func(ctx context.Context) error {
		return nil
	}

	app.StartupFn = func(ctx context.Context) error {
		return nil
	}

	app.ShutdownFn = func(ctx context.Context) error {
		return nil
	}

	app.DomReadyFn = func(ctx context.Context) error {
		return nil
	}

	app.BeforeCloseFn = func(ctx context.Context) error {
		return nil
	}

	return app, nil
}

type App struct {
	*API

	Name    string          `json:"name,omitempty"`    // application name
	Plugins plugins.Plugins `json:"plugins,omitempty"` // plugins for the application

	// save function, if nil, save to file in ~/.config/<name>/wailsx.json
	// will call Saver plugins
	SaveFn func(ctx context.Context) error `json:"-"`

	// startup function, if nil, load from file in ~/.config/<name>/wailsx.json
	// will call Startuper plugins
	StartupFn func(ctx context.Context) error `json:"-"`

	// shutdown function, if nil, call Save
	// will call Shutdowner plugins
	ShutdownFn func(ctx context.Context) error `json:"-"`

	// dom ready function, if nil, do nothing
	// will call DomReadyer plugins
	DomReadyFn func(ctx context.Context) error `json:"-"`

	// before close function, if nil, do nothing
	// will call BeforeCloser plugins
	BeforeCloseFn func(ctx context.Context) error `json:"-"`

	mu sync.RWMutex
}

func (app *App) PluginName() string {
	if app == nil {
		return ""
	}
	return fmt.Sprintf("%T: %s", app, app.Name)
}

func (app *App) StateData(ctx context.Context) (statedata.Data[AppData], error) {
	sd := statedata.Data[AppData]{
		Name: AppStateDataProviderName,
	}

	if app == nil {
		return sd, fmt.Errorf("app is nil")
	}

	app.mu.RLock()
	defer app.mu.RUnlock()

	api, err := app.API.StateData(ctx)
	if err != nil {
		return sd, err
	}

	data := AppData{
		AppName: app.Name,
		API:     api.Data,
		Plugins: map[string]any{},
	}

	for _, p := range app.Plugins {
		sdp, ok := p.(PluginDataProvider)
		if !ok {
			continue
		}

		pd, err := sdp.StateData(ctx)
		if err != nil {
			return sd, err
		}

		data.Plugins[p.PluginName()] = pd
	}

	sd.Data = data

	return sd, nil
}
