package wailsx

import (
	"context"

	"github.com/markbates/wailsx/statedata"
)

const AppStateDataProviderName = "app"

type AppStateDataProvider interface {
	StateData(ctx context.Context) (statedata.Data[AppData], error)
}
