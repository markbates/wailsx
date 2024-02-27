package wailsx

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_State_Save(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	var saved bool

	st := &State{
		Name:     "test",
		Position: NewPosition(),
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
	fp := filepath.Join(home, ".config", "wailsx-test", "state.json")
	r.NoError(os.RemoveAll(fp))

	st, err := NewState("wailsx-test")
	r.NoError(err)

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

func Test_State_Load(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	ctx := context.Background()
	r.Error((&State{}).Load(ctx))

	st := &State{
		Name: "test",
	}

	st.LoadFn = func(ctx context.Context) error {
		st.Position = NewPosition()
		return nil
	}

	r.NoError(st.Load(ctx))
	r.Equal("test", st.Name)
	r.NotNil(st.Position)

	st.LoadFn = func(ctx context.Context) error {
		return io.EOF
	}
	r.Error(st.Load(ctx))
}

func Test_State_Load_From_File(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	const name = "wailsx-test-load"

	home, err := os.UserHomeDir()
	r.NoError(err)
	fp := filepath.Join(home, ".config", name, "state.json")
	r.NoError(os.RemoveAll(fp))

	st, err := NewState(name)
	r.NoError(err)

	ctx := context.Background()

	r.NoError(st.Save(ctx))

	r.NoError(st.Load(context.Background()))

	r.Equal(name, st.Name)
	r.NotNil(st.Position)
}
