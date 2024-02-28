package wailsx

import (
	"context"
	"fmt"

	"github.com/markbates/plugins"
	"golang.org/x/sync/errgroup"
)

func (st *State) Shutdown(ctx context.Context) (err error) {
	if st == nil {
		return nil
	}

	st.mu.RLock()
	fn := st.ShutdownFn
	st.mu.RUnlock()

	// recover from external function call
	defer func() {
		r := recover()
		if r == nil {
			return
		}

		switch t := r.(type) {
		case error:
			err = t
		default:
			err = fmt.Errorf("%v", t)
		}
	}()

	if fn == nil {
		fn = st.Save
	}

	if err := fn(ctx); err != nil {
		return err
	}

	sps := plugins.ByType[Shutdowner](st.Plugins)

	var wg errgroup.Group
	for _, sp := range sps {
		sp := sp
		wg.Go(func() error {
			return sp.Shutdown(ctx)
		})
	}

	if err := wg.Wait(); err != nil {
		return err
	}

	return nil
}
