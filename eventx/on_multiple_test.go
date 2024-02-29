package eventx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Manager_OnMultiple(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const event = "event:test"

	em := NewManager()

	var counter int
	var canceled bool

	fn := func(ctx context.Context, name string, callback CallbackFn, n int) (CancelFn, error) {
		counter = n

		return func() error {
			canceled = true
			return nil
		}, nil
	}
	em.OnMultipleFn = fn

	cancelFn, err := em.OnMultiple(ctx, event, func(data ...any) error {
		return nil
	}, 42)

	r.NoError(err)
	r.NotNil(cancelFn)
	r.Equal(42, counter)

	r.NoError(cancelFn())
	r.True(canceled)
}
