//go:build !dev && !desktop && !production && !wails

// when not built with wails, the stubs are used
package windowx

import (
	"context"

	"github.com/markbates/wailsx/statedata"
)

func (mm MaximiserManager) StateData(ctx context.Context) (statedata.StateData[MaximiserData], error) {
	return mm.data.StateData(ctx)
}
