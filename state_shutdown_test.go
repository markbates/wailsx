package wailsx

import (
	"context"
	"errors"
	"testing"

	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_State_Shutdown(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	const name = "state-shutdown-test"

	st := newState(t, name)

	var shutdown bool

	st.ShutdownFn = func(ctx context.Context) error {
		shutdown = true
		return nil
	}

	ctx := context.Background()

	r.NoError(st.Shutdown(ctx))
	r.True(shutdown)
}

func Test_State_ShutdownPanic(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	const name = "state-shutdown-panic"

	st := &State{
		Name: name,
		ShutdownFn: func(ctx context.Context) error {
			panic("shutdown panic")
		},
	}

	err := st.Shutdown(context.Background())
	r.Error(err)

	r.Contains(err.Error(), "shutdown panic")

	st.ShutdownFn = func(ctx context.Context) error {
		panic(wailstest.ERR)
	}

	err = st.Shutdown(context.Background())
	r.Error(err)

	r.True(errors.Is(err, wailstest.ERR))
}
