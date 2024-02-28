package eventx

import (
	"testing"

	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_Manager_Now(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	em := EventManager{}
	r.NotZero(em.Now())
	r.NotEqual(wailstest.OldTime(), em.Now())

	em.NowFn = wailstest.OldTime
	r.Equal(wailstest.OldTime(), em.Now())

}
