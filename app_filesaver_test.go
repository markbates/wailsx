package wailsx

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_AppFilesaver_Save(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ctx := context.Background()

	var af *AppFilesaver
	r.Error(af.Save(ctx))

	af = &AppFilesaver{}
	r.Error(af.Save(ctx))

	app, err := NopApp("My App")
	r.NoError(err)

	af.App = app

	fp := filepath.Join(os.TempDir(), "Test_AppFilesaver_Save", "wailsx.json")
	af.Path = fp
	defer os.RemoveAll(fp)

	r.NoError(af.Save(ctx))

	b, err := os.ReadFile(fp)
	r.NoError(err)

	act := string(b)
	act = strings.TrimSpace(act)

	// fmt.Println(act)
	exp := `{"app_name":"My App","api":{"events":{},"window":{"maximiser":{},"positioner":{},"themer":{"background_colour":{}}}}}`

	r.Equal(exp, act)

}

func Test_AppFilesaver(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	uh, err := os.UserHomeDir()
	r.NoError(err)

	uh = filepath.Join(uh, ".config", "wailsx.json")

	tcs := []struct {
		name string
		af   *AppFilesaver
		exp  string
	}{
		{
			name: "nil",
			exp:  uh,
		},
		{
			name: "empty",
			af:   &AppFilesaver{},
			exp:  uh,
		},
		{
			name: "dir",
			af: &AppFilesaver{
				Path: "/tmp",
			},
			exp: filepath.Join("/tmp", "wailsx.json"),
		},
		{
			name: "file",
			af: &AppFilesaver{
				Path: "/tmp/file.json",
			},
			exp: "/tmp/file.json",
		},
		{
			name: "app",
			af: &AppFilesaver{
				Path: "/tmp",
				App: &App{
					Name: "My App",
				},
			},
			exp: filepath.Join("/tmp", "my_app.json"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			act, err := tc.af.filepath()
			r.NoError(err)

			r.Equal(tc.exp, act)
		})
	}

}
