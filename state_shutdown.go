package wailsx

import (
	"context"

	"github.com/markbates/plugins"
	"github.com/markbates/wailsx/internal/safe"
)

func (st *State) Shutdown(ctx context.Context) (err error) {
	if st == nil {
		return nil
	}

	st.mu.RLock()
	fn := st.ShutdownFn
	st.mu.RUnlock()

	if fn == nil {
		fn = st.Save
	}

	err = safe.Run(func() error {
		return fn(ctx)
	})

	if err != nil {
		return err
	}

	sps := plugins.ByType[Shutdowner](st.Plugins)

	var wg safe.Group

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
