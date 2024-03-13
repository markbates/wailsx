package eventx

import (
	"context"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_Manager_EventsOn(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tcs := []struct {
		name string
		fn   func(ctx context.Context, name string, callback CallbackFn) (CancelFn, error)
		err  error
	}{
		{
			name: "with function",
			fn: func(ctx context.Context, name string, callback CallbackFn) (CancelFn, error) {
				return nil, nil
			},
		},
		{
			name: "with error",
			fn: func(ctx context.Context, name string, callback CallbackFn) (CancelFn, error) {
				return nil, wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func(ctx context.Context, name string, callback CallbackFn) (CancelFn, error) {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
		{
			name: "no function",
			err:  wailsrun.ErrNotAvailable("EventsOn"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			r := require.New(t)

			m := &Manager{
				EventsOnFn: tc.fn,
			}

			_, err := m.EventsOn(ctx, "test", func(data ...any) error {
				return nil
			})

			if tc.err != nil {
				r.Error(err)
				r.Equal(tc.err, err)
				return
			}

			r.NoError(err)
		})
	}

}
