package menux

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ErrNilMenu(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	r.Equal("menu is nil", ErrNilMenu.Error())
}
