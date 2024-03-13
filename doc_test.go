package wailsx

import (
	"context"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/stretchr/testify/require"
)

// snippet: nil-api
func Test_Nil_API_Call(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var api *API

	ctx := context.Background()

	err := api.Show(ctx)
	r.Error(err)

	exp := wailsrun.ErrNotAvailable("Show")
	r.Equal(exp, err)
}

// snippet: nil-api

// snippet: nop-api
func Test_Nop_API_Call(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	api := NopAPI()

	ctx := context.Background()

	err := api.Show(ctx)
	r.NoError(err)
}

// snippet: nop-api
