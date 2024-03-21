package windowx

import (
	"context"
)

type StateDataProvider interface {
	StateData(ctx context.Context) (*WindowData, error)
}
