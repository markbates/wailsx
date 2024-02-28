package wailsx

import (
	"context"

	"github.com/markbates/plugins"
	"github.com/markbates/wailsx/internal/safe"
)

func (st *State) DomReady(ctx context.Context) error {
	if st.DomReadyFn != nil {
		err := safe.Run(func() error {
			return st.DomReadyFn(ctx)
		})

		if err != nil {
			return err
		}
	}

	var wg safe.Group

	drs := plugins.ByType[DomReadyer](st.Plugins)
	for _, dr := range drs {
		dr := dr
		wg.Go(func() error {
			return dr.DomReady(ctx)
		})
	}

	if err := wg.Wait(); err != nil {
		return err
	}

	return nil
}
