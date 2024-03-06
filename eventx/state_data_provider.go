package eventx

import (
	"context"

	"github.com/markbates/wailsx/statedata"
)

type StateDataProvider interface {
	StateData(ctx context.Context) (statedata.Data[*EventsData], error)
}
