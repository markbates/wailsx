package wailsx

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_State_Save(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	var saved bool

	st := &State{
		Name:   "test",
		Layout: NewLayout(),
		SaveFn: func(ctx context.Context) error {
			saved = true
			return nil
		},
	}

	ctx := context.Background()

	r.NoError(st.Save(ctx))
	r.True(saved)

	st = &State{}
	r.Error(st.Save(ctx))
}

func Test_State_SaveToFile(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	home, err := os.UserHomeDir()
	r.NoError(err)

	const name = "wailsx-test"

	fp := filepath.Join(home, ".config", name, "state.json")
	r.NoError(os.RemoveAll(fp))

	st := newState(t, name)

	r.NoError(st.Save(context.Background()))

	b, err := os.ReadFile(fp)
	r.NoError(err)

	act := string(b)
	act = strings.TrimSpace(act)

	// os.MkdirAll("testdata/state", 0755)
	// f, err := os.Create("testdata/state/save.json")
	// r.NoError(err)
	// f.Write([]byte(act))
	// r.NoError(f.Close())

	b, err = os.ReadFile("testdata/state/save.json")
	r.NoError(err)

	exp := string(b)
	r.Equal(exp, act)
}

func Test_State_SavePanic(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	st := &State{
		Name: "test",
		SaveFn: func(ctx context.Context) error {
			panic("save panic")
		},
	}

	err := st.Save(context.Background())
	r.Error(err)

	r.Contains(err.Error(), "save panic")

	st.SaveFn = func(ctx context.Context) error {
		panic(wailstest.ERR)
	}

	err = st.Save(context.Background())
	r.Error(err)

	r.True(errors.Is(err, wailstest.ERR))
}

func Test_State_Save_WithPlugins(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	st, err := NewState("test")
	r.NoError(err)

	var stSaved bool
	st.SaveFn = func(ctx context.Context) error {
		stSaved = true
		return nil
	}

	p1 := &wailstest.SaverPlugin{}
	p2 := &wailstest.SaverPlugin{}
	st.Plugins = append(st.Plugins, p1, p2)

	ctx := context.Background()

	err = st.Save(ctx)
	r.NoError(err)

	r.True(stSaved)
	r.True(p1.Saved)
	r.True(p2.Saved)
}
