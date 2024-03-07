package eventx

import (
	"context"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_Manager_OnMultiple(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	const event = "event:test"

	tcs := []struct {
		name string
		fn   func() error
		err  error
	}{
		{
			name: "with function",
			fn: func() error {
				return nil
			},
		},
		{
			name: "with error",
			fn: func() error {
				return wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "with panic",
			fn: func() error {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
		{
			name: "no function",
			err:  wailsrun.ErrNotAvailable("EventsOnMultiple"),
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			r := require.New(t)

			m := &Manager{}

			if tc.fn != nil {
				m.EventsOnMultipleFn = func(ctx context.Context, name string, callback wailsrun.CallbackFn, counter int) (wailsrun.CancelFn, error) {
					return nil, tc.fn()
				}
			}

			_, err := m.EventsOnMultiple(ctx, event, func(data ...any) error {
				return nil
			}, 1)

			r.Equal(tc.err, err)
		})
	}

}
