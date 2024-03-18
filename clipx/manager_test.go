package clipx

import (
	"context"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_Manager_StateData(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	m := NopManager()

	ctx := context.Background()

	exp := "hello"

	r.NoError(m.ClipboardSetText(ctx, exp))

	act, err := m.ClipboardGetText(ctx)
	r.NoError(err)

	r.Equal(exp, act)

	sd, err := m.StateData(ctx)
	r.NoError(err)

	r.Equal(ClipboardManagerStateDataProviderName, sd.Name)
	r.Equal(exp, sd.Data)

	m = nil

	sd, err = m.StateData(ctx)
	r.NoError(err)
	r.Equal(ClipboardManagerStateDataProviderName, sd.Name)
	r.Empty(sd.Data)
}

func Test_Manager_ClipboardGetText(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	const exp = "hello"

	tcs := []struct {
		name string
		fn   func(ctx context.Context) (string, error)
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context) (string, error) {
				return exp, nil
			},
		},
		{
			name: "with error",
			fn: func(ctx context.Context) (string, error) {
				return "", wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func(ctx context.Context) (string, error) {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
		{
			name: "no function",
			err:  wailsrun.ErrNotAvailable("ClipboardGetText"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			m := &Manager{
				ClipboardGetTextFn: tc.fn,
			}

			act, err := m.ClipboardGetText(ctx)

			if tc.err != nil {
				r.Error(err)
				r.Equal(tc.err, err)
				return
			}

			r.NoError(err)
			r.Equal(exp, act)
		})
	}

}

func Test_Manager_ClipboardSetText(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	const exp = "hello"

	tcs := []struct {
		name string
		fn   func(ctx context.Context, text string) error
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context, text string) error {
				return nil
			},
		},
		{
			name: "with error",
			fn: func(ctx context.Context, text string) error {
				return wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func(ctx context.Context, text string) error {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
		{
			name: "no function",
			err:  wailsrun.ErrNotAvailable("ClipboardSetText"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			m := &Manager{
				ClipboardSetTextFn: tc.fn,
			}

			err := m.ClipboardSetText(ctx, exp)

			if tc.err != nil {
				r.Error(err)
				r.Equal(tc.err, err)
				return
			}

			r.NoError(err)
			r.Equal(exp, m.Content)
		})
	}
}

func Test_Nil_Manager(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ctx := context.Background()

	var m *Manager

	_, err := m.ClipboardGetText(ctx)
	r.Error(err)

	exp := wailsrun.ErrNotAvailable("ClipboardGetText")
	r.Equal(exp, err)

	err = m.ClipboardSetText(ctx, "hello")
	r.Error(err)

	exp = wailsrun.ErrNotAvailable("ClipboardSetText")
	r.Equal(exp, err)
}

func Test_Manager_RestoreClipboard(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ctx := context.Background()

	const exp = "hello"

	var m *Manager
	err := m.RestoreClipboard(ctx, exp)
	r.Error(err)

	exr := wailsrun.ErrNotAvailable("ClipboardSetText")
	r.Equal(exr, err)

	m = &Manager{}
	err = m.RestoreClipboard(ctx, exp)
	r.Error(err)

	r.Equal(exr, err)

	var act string
	m.ClipboardSetTextFn = func(ctx context.Context, text string) error {
		act = text
		return nil
	}

	err = m.RestoreClipboard(ctx, exp)
	r.NoError(err)
	r.Equal(exp, act)
}
