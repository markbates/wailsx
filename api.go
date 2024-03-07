package wailsx

import (
	"context"
	"os"

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
		DialogManager: dialogx.Manager{},
		EventManager:  eventx.NewManager(),
		WailsLogger:   logx.NewLogger(os.Stdout, wailsrun.INFO),
		MenuManager:   menux.Manager{},
		WindowManager: windowx.NewManager(),
	}
}

func NopAPI() *API {
	return &API{
		DialogManager: dialogx.NopManager(),
		EventManager:  eventx.NopManager(),
		WailsLogger:   logx.NewLogger(os.Stdout, wailsrun.INFO),
		MenuManager:   menux.NopManager(),
		WindowManager: windowx.NopManager(),
	}
}

type API struct {
	dialogx.DialogManager
	eventx.EventManager
	logx.WailsLogger
	menux.MenuManager
	windowx.WindowManager
}

func (api *API) StateData(ctx context.Context) (statedata.Data[*APIData], error) {
	sd := statedata.Data[*APIData]{
		Name: APIStateDataProviderName,
	}

	if api == nil {
		return sd, nil
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
