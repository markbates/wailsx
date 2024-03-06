package windowx

import (
	"context"

	"github.com/markbates/wailsx/statedata"
)

type StateDataProvider interface {
	StateData(ctx context.Context) (statedata.Data[*WindowData], error)
}
