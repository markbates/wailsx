package eventx

import (
	"context"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/stretchr/testify/require"
)

func Test_Manager_Once(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const event = "event:test"

	em := NewManager()

	var counter int
	var canceled bool

	fn := func(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error) {
		counter = 1

		return func() error {
			canceled = true
			return nil
		}, nil
	}
	em.OnceFn = fn

	cancelFn, err := em.EventsOnce(ctx, event, func(data ...any) error {
		return nil
	})

	r.NoError(err)
	r.NotNil(cancelFn)
	r.Equal(1, counter)

	r.NoError(cancelFn())
	r.True(canceled)
}
