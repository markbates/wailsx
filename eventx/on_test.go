package eventx_test

import (
	"context"
	"errors"
	"testing"

	. "github.com/markbates/wailsx/eventx"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_EventManager_On(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	em, ec := newEventManager()
	_ = ec

	const evt = "event:test"

	em.OnFn = func(ctx context.Context, name string, callback CallbackFn) (CancelFn, error) {
		if name != evt {
			return nil, wailstest.ErrTest
		}

		if err := callback(); err != nil {
			return nil, err
		}

		return func() error {
			return nil
		}, nil
	}

	tcs := []struct {
		name string
		cb   CallbackFn
		err  bool
	}{
		{
			name: "no error",
			cb:   func(data ...any) error { return nil },
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := em.On(context.Background(), evt, tc.cb)

			if !tc.err {
				r.NoError(err)
				return
			}

			r.Error(err)
			r.True(errors.Is(err, wailstest.ErrTest))
		})
	}

}
