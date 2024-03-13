package wailsx

import (
	"context"

	"github.com/markbates/wailsx/statedata"
)

type PluginDataProvider interface {
	StateData(ctx context.Context) (statedata.Data[any], error)
}
