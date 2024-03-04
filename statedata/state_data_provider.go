package statedata

import "context"

type StateDataProvider[T any] interface {
	StateData(ctx context.Context) (StateData[T], error)
}
