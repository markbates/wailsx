package wailsx

import (
	"context"

	"github.com/markbates/wailsx/statedata"
)

type APIStateDataProvider interface {
	StateData(ctx context.Context) (statedata.Data[*APIData], error)
}
