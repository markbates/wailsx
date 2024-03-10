package wailsrun_test

import (
	"context"
	"testing"

	"github.com/markbates/wailsx/wailsrun"
	"github.com/stretchr/testify/require"
)

// snippet: err-not-available
func Test_ErrNotAvailable(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	ctx := context.Background()

	err := wailsrun.BrowserOpenURL(ctx, "https://example.com")
	r.Error(err)

	exp := wailsrun.ErrNotAvailable("BrowserOpenURL")
	r.Equal(exp, err)
}

// snippet: err-not-available
