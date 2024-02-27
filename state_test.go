package wailsx

import (
	"context"
	"errors"
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
		panic(io.EOF)
	}

	err = st.Save(context.Background())
	r.Error(err)

	r.True(errors.Is(err, io.EOF))
}

func Test_State_Startup(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	ctx := context.Background()
	r.Error((&State{}).Startup(ctx))

	const name = "state-startup-test"
	st := newState(t, name)

	st.StartupFn = func(ctx context.Context) error {
		st.Position = NewPosition()
		return nil
	}

	r.NoError(st.Startup(ctx))
	r.Equal(name, st.Name)
	r.NotNil(st.Position)

	st.StartupFn = func(ctx context.Context) error {
		return io.EOF
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
	r.NotNil(st.Position)
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
		panic(io.EOF)
	}

	err = st.Startup(context.Background())
	r.Error(err)

	r.True(errors.Is(err, io.EOF))
}

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
		panic(io.EOF)
	}

	err = st.Shutdown(context.Background())
	r.Error(err)

	r.True(errors.Is(err, io.EOF))
}
