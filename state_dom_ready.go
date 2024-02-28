package wailsx

import (
	"context"

	"github.com/markbates/plugins"
)

func (st *State) DomReady(ctx context.Context) error {
	if st.DomReadyFn != nil {
		if err := st.DomReadyFn(ctx); err != nil {
			return err
		}
	}

	drs := plugins.ByType[DomReadyer](st.Plugins)
	for _, dr := range drs {
		if err := dr.DomReady(ctx); err != nil {
			return err
		}
	}

	return nil
}
