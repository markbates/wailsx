package wailsx

import (
	"context"

	"github.com/markbates/wailsx/statedata"
)

const APIStateDataProviderName = "api"

type APIStateDataProvider interface {
	StateData(ctx context.Context) (statedata.Data[*APIData], error)
}
