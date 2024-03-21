package eventx

import (
	"context"
)

type StateDataProvider interface {
	StateData(ctx context.Context) (*EventsData, error)
}
