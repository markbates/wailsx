package wailsx

import (
	"context"
	"testing"

	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_State_Save_WithPlugins(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	st, err := NewState("test")
	r.NoError(err)

	var stSaved bool
	st.SaveFn = func(ctx context.Context) error {
		stSaved = true
		return nil
	}

	p1 := &wailstest.SaverPlugin{}
	p2 := &wailstest.SaverPlugin{}
	st.Plugins = append(st.Plugins, p1, p2)

	ctx := context.Background()

	err = st.Save(ctx)
	r.NoError(err)

	r.True(stSaved)
	r.True(p1.Saved)
	r.True(p2.Saved)
}
