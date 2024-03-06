package windowx

import (
	"context"

	"github.com/markbates/wailsx/statedata"
)

type WindowData struct {
	*MaximiserData  `json:"maximiser_data,omitempty"`
	*PositionerData `json:"positioner_data,omitempty"`
	*ThemerData     `json:"themer_data,omitempty"`
}

func (wm *Manager) StateData(ctx context.Context) (statedata.StateData[*WindowData], error) {
	sd := statedata.StateData[*WindowData]{
		Name: ManagerStateDataName,
	}
	if wm == nil {
		return sd, nil
	}

	data := &WindowData{}

	if x, ok := wm.Maximiser.(interface {
		StateData(context.Context) (statedata.StateData[*MaximiserData], error)
	}); ok {
		md, err := x.StateData(ctx)
		if err != nil {
			return sd, err
		}
		data.MaximiserData = md.Data
	}

	if x, ok := wm.Positioner.(interface {
		StateData(context.Context) (statedata.StateData[*PositionerData], error)
	}); ok {
		pd, err := x.StateData(ctx)
		if err != nil {
			return sd, err
		}
		data.PositionerData = pd.Data
	}

	if x, ok := wm.Themer.(interface {
		StateData(context.Context) (statedata.StateData[*ThemerData], error)
	}); ok {
		td, err := x.StateData(ctx)
		if err != nil {
			return sd, err
		}
		data.ThemerData = td.Data
	}

	sd.Data = data

	return sd, nil
}
