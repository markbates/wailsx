package statedata

import (
	"context"
	"fmt"
)

var _ DataProvider[any] = Data[any]{}

type Data[T any] struct {
	Name string `json:"name,omitempty"` // name of the data
	Data T      `json:"data,omitempty"` // data for the state
}

func (sd Data[T]) PluginName() string {
	return fmt.Sprintf("%T: %s", sd, sd.Name)
}

func (sd Data[T]) StateData(ctx context.Context) (Data[T], error) {
	return sd, nil
}
