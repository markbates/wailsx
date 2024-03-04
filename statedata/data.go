package statedata

import (
	"context"
	"fmt"
)

var _ StateDataProvider[any] = StateData[any]{}

type StateData[T any] struct {
	Name string `json:"name,omitempty"` // name of the data
	Data T      `json:"data,omitempty"` // data for the state
}

func (sd StateData[T]) PluginName() string {
	return fmt.Sprintf("%T: %s", sd, sd.Name)
}

func (sd StateData[T]) StateData(ctx context.Context) (StateData[T], error) {
	return sd, nil
}
