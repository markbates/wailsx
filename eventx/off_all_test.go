package eventx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Manager_OffAll(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	em := NewManager()

	var called bool
	fn := func(ctx context.Context) error {
		called = true
		return nil
	}

	em.OffAllFn = fn

	ctx := context.Background()
	r.NoError(em.OffAll(ctx))

	r.True(called)

}
