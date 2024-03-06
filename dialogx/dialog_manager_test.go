package dialogx

import (
	"context"
	"errors"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_DialogManager_MessageDialog(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tcs := []struct {
		name string
		fn   func(ctx context.Context, opts MessageDialogOptions) (string, error)
		exp  string
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context, opts MessageDialogOptions) (string, error) {
				return "foo", nil
			},
			exp: "foo",
		},
		{
			name: "no function",
			err:  wailsrun.ErrNotAvailable,
		},
		{
			name: "error",
			fn: func(ctx context.Context, opts MessageDialogOptions) (string, error) {
				return "", wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "panic",
			fn: func(ctx context.Context, opts MessageDialogOptions) (string, error) {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			dm := DialogManager{
				MessageDialogFn: tc.fn,
			}

			fp, err := dm.MessageDialog(ctx, MessageDialogOptions{})

			if tc.err != nil {
				r.Error(err)
				r.True(errors.Is(err, tc.err))
				return
			}

			r.NoError(err)
			r.Equal(tc.exp, fp)
		})
	}

}

func Test_DialogManager_OpenDirectoryDialog(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tcs := []struct {
		name string
		fn   func(ctx context.Context, opts OpenDialogOptions) (string, error)
		exp  string
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context, opts OpenDialogOptions) (string, error) {
				return "foo", nil
			},
			exp: "foo",
		},
		{
			name: "no function",
			err:  wailsrun.ErrNotAvailable,
		},
		{
			name: "error",
			fn: func(ctx context.Context, opts OpenDialogOptions) (string, error) {
				return "", wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "panic",
			fn: func(ctx context.Context, opts OpenDialogOptions) (string, error) {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			dm := DialogManager{
				OpenDirectoryDialogFn: tc.fn,
			}

			fp, err := dm.OpenDirectoryDialog(ctx, OpenDialogOptions{})

			if tc.err != nil {
				r.Error(err)
				r.True(errors.Is(err, tc.err))
				return
			}

			r.NoError(err)
			r.Equal(tc.exp, fp)
		})
	}
}

func Test_DialogManager_OpenFileDialog(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tcs := []struct {
		name string
		fn   func(ctx context.Context, opts OpenDialogOptions) (string, error)
		exp  string
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context, opts OpenDialogOptions) (string, error) {
				return "foo", nil
			},
			exp: "foo",
		},
		{
			name: "no function",
			err:  wailsrun.ErrNotAvailable,
		},
		{
			name: "error",
			fn: func(ctx context.Context, opts OpenDialogOptions) (string, error) {
				return "", wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "panic",
			fn: func(ctx context.Context, opts OpenDialogOptions) (string, error) {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			dm := DialogManager{
				OpenFileDialogFn: tc.fn,
			}

			fp, err := dm.OpenFileDialog(ctx, OpenDialogOptions{})

			if tc.err != nil {
				r.Error(err)
				r.True(errors.Is(err, tc.err))
				return
			}

			r.NoError(err)
			r.Equal(tc.exp, fp)
		})
	}
}

func Test_DialogManager_OpenMultipleFilesDialog(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tcs := []struct {
		name string
		fn   func(ctx context.Context, opts OpenDialogOptions) ([]string, error)
		exp  []string
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context, opts OpenDialogOptions) ([]string, error) {
				return []string{"foo"}, nil
			},
			exp: []string{"foo"},
		},
		{
			name: "no function",
			err:  wailsrun.ErrNotAvailable,
		},
		{
			name: "error",
			fn: func(ctx context.Context, opts OpenDialogOptions) ([]string, error) {
				return nil, wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "panic",
			fn: func(ctx context.Context, opts OpenDialogOptions) ([]string, error) {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			dm := DialogManager{
				OpenMultipleFilesDialogFn: tc.fn,
			}

			fp, err := dm.OpenMultipleFilesDialog(ctx, OpenDialogOptions{})

			if tc.err != nil {
				r.Error(err)
				r.True(errors.Is(err, tc.err))
				return
			}

			r.NoError(err)
			r.Equal(tc.exp, fp)
		})
	}
}

func Test_DialogManager_SaveFileDialog(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tcs := []struct {
		name string
		fn   func(ctx context.Context, opts SaveDialogOptions) (string, error)
		exp  string
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context, opts SaveDialogOptions) (string, error) {
				return "foo", nil
			},
			exp: "foo",
		},
		{
			name: "no function",
			err:  wailsrun.ErrNotAvailable,
		},
		{
			name: "error",
			fn: func(ctx context.Context, opts SaveDialogOptions) (string, error) {
				return "", wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "panic",
			fn: func(ctx context.Context, opts SaveDialogOptions) (string, error) {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			dm := DialogManager{
				SaveFileDialogFn: tc.fn,
			}

			fp, err := dm.SaveFileDialog(ctx, SaveDialogOptions{})

			if tc.err != nil {
				r.Error(err)
				r.True(errors.Is(err, tc.err))
				return
			}

			r.NoError(err)
			r.Equal(tc.exp, fp)
		})
	}
}
