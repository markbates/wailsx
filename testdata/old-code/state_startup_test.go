package wailsx

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_State_Startup(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	ctx := context.Background()
	r.Error((&State{}).Startup(ctx))

	const name = "state-startup-test"
	st := newState(t, name)

	st.StartupFn = func(ctx context.Context) error {
		st.Layout = NewLayout()
		return nil
	}

	r.NoError(st.Startup(ctx))
	r.Equal(name, st.Name)
	r.NotNil(st.Layout)

	st.StartupFn = func(ctx context.Context) error {
		return wailstest.ErrTest
	}
	r.Error(st.Startup(ctx))
}

func Test_State_Startp_From_File(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	const name = "wailsx-test-load"

	home, err := os.UserHomeDir()
	r.NoError(err)

	fp := filepath.Join(home, ".config", name, "state.json")
	r.NoError(os.RemoveAll(fp))

	st := newState(t, name)

	ctx := context.Background()

	r.NoError(st.Save(ctx))

	r.NoError(st.Startup(context.Background()))

	r.Equal(name, st.Name)
	r.NotNil(st.Layout)
}

func Test_State_StartupPanic(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	const name = "state-startup-panic"

	st := &State{
		Name: name,
		StartupFn: func(ctx context.Context) error {
			panic("startup panic")
		},
	}

	err := st.Startup(context.Background())
	r.Error(err)

	r.Contains(err.Error(), "startup panic")

	st.StartupFn = func(ctx context.Context) error {
		panic(wailstest.ErrTest)
	}

	err = st.Startup(context.Background())
	r.Error(err)

	r.True(errors.Is(err, wailstest.ErrTest))
}

func Test_State_Startup_WithPlugins(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	const name = "state-startup-plugins"

	st := newState(t, name)

	var startup bool

	st.StartupFn = func(ctx context.Context) error {
		startup = true
		return nil
	}

	p1 := &wailstest.StartuperPlugin{}
	p2 := &wailstest.StartuperPlugin{}

	st.Plugins = append(st.Plugins, p1, p2)

	ctx := context.Background()

	r.NoError(st.Startup(ctx))
	r.True(startup)

	r.True(p1.Called)
	r.True(p2.Called)
}

func Test_State_Startup_PluginError(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	const name = "state-startup-plugin-error"

	st := newState(t, name)

	var startup bool

	st.StartupFn = func(ctx context.Context) error {
		startup = true
		return nil
	}

	p1 := &wailstest.StartuperPlugin{Error: true}
	p2 := &wailstest.StartuperPlugin{}

	st.Plugins = append(st.Plugins, p1, p2)

	ctx := context.Background()

	err := st.Startup(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))

	r.True(startup)
	r.True(p1.Called)
	r.False(p2.Called)

}

func Test_State_Startup_Error_Stops_Plugins(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	const name = "state-startup-error-stops-plugins"

	st := newState(t, name)

	var startup bool

	st.StartupFn = func(ctx context.Context) error {
		startup = true
		return wailstest.ErrTest
	}

	p1 := &wailstest.StartuperPlugin{}
	p2 := &wailstest.StartuperPlugin{}

	st.Plugins = append(st.Plugins, p1, p2)

	ctx := context.Background()

	err := st.Startup(ctx)
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))

	r.True(startup)
	r.False(p1.Called)
	r.False(p2.Called)
}
