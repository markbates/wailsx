package wailsx

import (
	"path/filepath"
	"testing"

	"github.com/markbates/plugins"
	"github.com/stretchr/testify/require"
)

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
