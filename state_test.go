package wailsx

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/markbates/plugins"
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
		return wailstest.ERR
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
		panic(wailstest.ERR)
	}

	err = st.Startup(context.Background())
	r.Error(err)

	r.True(errors.Is(err, wailstest.ERR))
}

func Test_State_MarshalJSON(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name    string
		plugins plugins.Plugins
		err     bool
	}{
		{
			name: "simple",
		},
		{
			name: "with_plugins",
			plugins: plugins.Plugins{
				stringData("hello"),
				stringData("world"),
				StateData{
					Name: "foo",
					Data: map[string]int{
						"one": 1,
						"two": 2,
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			st := &State{
				Name:    "test",
				Plugins: tc.plugins,
			}

			if tc.err {
				_, err := st.MarshalJSON()
				r.Error(err)
				return
			}

			assertJSON(t, filepath.Join("state", "marshal", tc.name), st)
		})
	}

}
