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

func Test_State_Shutdown_WithPlugins(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	const name = "state-shutdown-plugins"

	st := newState(t, name)

	var shutdown bool

	st.ShutdownFn = func(ctx context.Context) error {
		shutdown = true
		return nil
	}

	p1 := &wailstest.ShutdownerPlugin{}
	p2 := &wailstest.ShutdownerPlugin{}

	st.Plugins = append(st.Plugins, p1, p2)

	ctx := context.Background()

	r.NoError(st.Shutdown(ctx))
	r.True(shutdown)

	r.True(p1.Called)
	r.True(p2.Called)
}

func Test_State_Shutdown_PluginError(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	const name = "state-shutdown-plugin-error"

	st := newState(t, name)

	var shutdown bool

	st.ShutdownFn = func(ctx context.Context) error {
		shutdown = true
		return nil
	}

	p1 := &wailstest.ShutdownerPlugin{Error: true}
	p2 := &wailstest.ShutdownerPlugin{}

	st.Plugins = append(st.Plugins, p1, p2)

	ctx := context.Background()

	err := st.Shutdown(ctx)
	r.Error(err)

	r.True(errors.Is(err, wailstest.ERR))

	r.True(shutdown)
	r.True(p1.Called)
}

func Test_State_Shutdown_Error_Stops_Plugins(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	const name = "state-shutdown-error-stops-plugins"

	st := newState(t, name)

	var shutdown bool

	st.ShutdownFn = func(ctx context.Context) error {
		shutdown = true
		return wailstest.ERR
	}

	p1 := &wailstest.ShutdownerPlugin{}
	p2 := &wailstest.ShutdownerPlugin{}

	st.Plugins = append(st.Plugins, p1, p2)

	ctx := context.Background()

	err := st.Shutdown(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ERR))

	r.True(shutdown)
	r.False(p1.Called)
	r.False(p2.Called)

}
