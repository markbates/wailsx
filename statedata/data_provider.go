package statedata

import "context"

type DataProvider[T any] interface {
	StateData(ctx context.Context) (Data[T], error)
}
