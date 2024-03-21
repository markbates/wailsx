package statedata

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Data(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	ctx := context.Background()

	ds := Data[string]{Data: "test"}
	r.Equal("statedata.Data[string]", ds.PluginName())

	sds, err := ds.StateData(ctx)
	r.NoError(err)
	r.Equal(ds, sds)

	di := Data[int]{Data: 1}
	r.Equal("statedata.Data[int]", di.PluginName())

	sdi, err := di.StateData(ctx)
	r.NoError(err)
	r.Equal(di, sdi)
}
