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

type WailsAPI interface {
	dialogx.DialogManager
	eventx.EventManager
	logx.WailsLogger
	menux.MenuManager
	windowx.WindowManager

	APIStateDataProvider
}

var _ WailsAPI = &API{}

func NewAPI() *API {
	return &API{
		DialogManager: dialogx.Manager{},
		EventManager:  eventx.NewManager(),
		WailsLogger:   logx.NewLogger(os.Stdout, wailsrun.INFO),
		MenuManager:   menux.Manager{},
		WindowManager: windowx.NewManager(),
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
		Name: "api",
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
			data.WindowData = wd.Data
		}
	}

	if x, ok := api.EventManager.(eventx.StateDataProvider); ok {
		ed, err := x.StateData(ctx)
		if err != nil {
			return sd, err
		}

		if ed.Data != nil {
			data.EventsData = ed.Data
		}
	}

	return sd, nil
}

type APIData struct {
	*eventx.EventsData
	*windowx.WindowData
}
