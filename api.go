package wailsx

import (
	"context"
	"os"

	"github.com/markbates/wailsx/eventx"
	"github.com/markbates/wailsx/logx"
	"github.com/markbates/wailsx/statedata"
	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/windowx"
)

type WailsAPI interface {
	eventx.EventManager
	logx.WailsLogger
	windowx.WindowManager
}

var _ WailsAPI = &API{}

func NewAPI() *API {
	return &API{
		EventManager:  eventx.NewManager(),
		WailsLogger:   logx.NewLogger(os.Stdout, wailsrun.INFO),
		WindowManager: windowx.NewManager(),
	}
}

type API struct {
	eventx.EventManager
	logx.WailsLogger
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

	if x, ok := api.WindowManager.(interface {
		StateData(context.Context) (statedata.Data[*windowx.WindowData], error)
	}); ok {
		wd, err := x.StateData(ctx)
		if err != nil {
			return sd, err
		}
		data.WindowData = wd.Data
	}

	if x, ok := api.EventManager.(interface {
		StateData(context.Context) (statedata.Data[*eventx.EventsData], error)
	}); ok {
		ed, err := x.StateData(ctx)
		if err != nil {
			return sd, err
		}
		data.EventsData = ed.Data
	}

	return sd, nil
}

type APIData struct {
	*eventx.EventsData
	*windowx.WindowData
}
