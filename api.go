package wailsx

import (
	"context"
	"fmt"
	"os"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/clipx"
	"github.com/markbates/wailsx/dialogx"
	"github.com/markbates/wailsx/eventx"
	"github.com/markbates/wailsx/logx"
	"github.com/markbates/wailsx/menux"
	"github.com/markbates/wailsx/statedata"
	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/windowx"
)

var _ wailsrun.API = &API{}

func NewAPI() *API {
	return &API{
		ClipboardManager: &clipx.Manager{},
		DialogManager:    dialogx.Manager{},
		EventManager:     eventx.NewManager(),
		MenuManager:      menux.Manager{},
		WailsLogger:      logx.NewLogger(os.Stdout, wailsrun.INFO),
		WindowManager:    windowx.NewManager(),
	}
}

func NopAPI() *API {
	return &API{
		ClipboardManager: clipx.NopManager(),
		DialogManager:    dialogx.NopManager(),
		EventManager:     eventx.NopManager(),
		MenuManager:      menux.NopManager(),
		WailsLogger:      logx.NewLogger(os.Stdout, wailsrun.INFO),
		WindowManager:    windowx.NopManager(),
		BrowserOpenURLFn: func(ctx context.Context, url string) error {
			return nil
		},
		QuitFn: func(ctx context.Context) error {
			return nil
		},
	}
}

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

func (api *API) StateData(ctx context.Context) (statedata.Data[*APIData], error) {
	sd := statedata.Data[*APIData]{
		Name: APIStateDataProviderName,
	}

	if api == nil {
		return sd, fmt.Errorf("api is nil")
	}

	data := &APIData{}

	if x, ok := api.WindowManager.(windowx.StateDataProvider); ok {
		wd, err := x.StateData(ctx)
		if err != nil {
			return sd, err
		}

		if wd.Data != nil {
			data.Window = wd.Data
		}
	}

	if x, ok := api.EventManager.(eventx.StateDataProvider); ok {
		ed, err := x.StateData(ctx)
		if err != nil {
			return sd, err
		}

		if ed.Data != nil {
			data.Events = ed.Data
		}
	}

	sd.Data = data

	return sd, nil
}

func (api *API) BrowserOpenURL(ctx context.Context, url string) error {
	if api == nil {
		return wailsrun.BrowserOpenURL(ctx, url)
	}

	return safe.Run(func() error {
		fn := api.BrowserOpenURLFn
		if fn == nil {
			fn = wailsrun.BrowserOpenURL
		}

		return fn(ctx, url)
	})
}

func (api *API) Quit(ctx context.Context) error {
	if api == nil {
		return wailsrun.Quit(ctx)
	}

	return safe.Run(func() error {
		fn := api.QuitFn
		if fn == nil {
			fn = wailsrun.Quit
		}

		return fn(ctx)
	})
}
