package windowx

import (
	"context"

	"github.com/markbates/wailsx/statedata"
)

type WindowData struct {
	*MaximiserData  `json:"maximiser_data,omitempty"`
	*PositionerData `json:"positioner_data,omitempty"`
	*ThemeData      `json:"themer_data,omitempty"`
}

func (wm *Manager) StateData(ctx context.Context) (statedata.Data[*WindowData], error) {
	sd := statedata.Data[*WindowData]{
		Name: ManagerStateDataName,
	}
	if wm == nil {
		return sd, nil
	}

	data := &WindowData{}

	if x, ok := wm.MaximiseManager.(interface {
		StateData(context.Context) (statedata.Data[*MaximiserData], error)
	}); ok {
		md, err := x.StateData(ctx)
		if err != nil {
			return sd, err
		}
		data.MaximiserData = md.Data
	}

	if x, ok := wm.PositionerManager.(interface {
		StateData(context.Context) (statedata.Data[*PositionerData], error)
	}); ok {
		pd, err := x.StateData(ctx)
		if err != nil {
			return sd, err
		}
		data.PositionerData = pd.Data
	}

	if x, ok := wm.ThemeManager.(interface {
		StateData(context.Context) (statedata.Data[*ThemeData], error)
	}); ok {
		td, err := x.StateData(ctx)
		if err != nil {
			return sd, err
		}
		data.ThemeData = td.Data
	}

	sd.Data = data

	return sd, nil
}
