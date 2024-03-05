package eventx

import (
	"context"
	"errors"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_Manager_EventsOffAll(t *testing.T) {
	t.Parallel()

	events := []string{"test:event1", "test:event2"}

	cb := func(data ...any) error {
		return nil
	}

	tcs := []struct {
		name string
		fn   func(ctx context.Context) error
		err  error
	}{
		{
			name: "no error",
			fn: func(ctx context.Context) error {
				return nil
			},
		},
		{
			name: "returns an error",
			fn: func(ctx context.Context) error {
				return wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "panics",
			fn: func(ctx context.Context) error {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
		{
			name: "nil function",
			err:  wailsrun.ErrNotAvailable,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			em := newTestManager()
			em.EventsOffAllFn = tc.fn

			for _, event := range events {
				em.data.AddCallback(event, cb, 0)
			}

			err := em.EventsOffAll(context.Background())
			if tc.err != nil {
				r.Equal(tc.err, err)
				errors.Is(err, tc.err)
				return
			}

			r.NoError(err)

			for _, event := range events {
				ce, ok := em.data.Callbacks[event]
				r.True(ok)
				r.True(ce.Off)
			}

		})
	}

	// r := require.New(t)

	// em := NewManager()

	// var called bool
	// fn := func(ctx context.Context) error {
	// 	called = true
	// 	return nil
	// }

	// em.EventsOffAllFn = fn

	// ctx := context.Background()
	// r.NoError(em.EventsOffAll(ctx))

	// r.True(called)

}
