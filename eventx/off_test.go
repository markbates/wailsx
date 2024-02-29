package eventx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Manager_Off(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	em := NewManager()

	const event = "event:test"

	var act string
	fn := func(ctx context.Context, name string, additional ...string) error {
		act = name
		return nil
	}

	em.OffFn = fn

	ctx := context.Background()
	r.NoError(em.Off(ctx, event))

	r.Equal(event, act)
}
