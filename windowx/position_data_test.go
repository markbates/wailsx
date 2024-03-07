package windowx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_PositionData_IsCentered(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	var pd *PositionData
	r.Error(pd.SetCentered())

	pd = &PositionData{}
	r.NoError(pd.SetCentered())
	r.True(pd.IsCentered)
}

func Test_PositionData_SetPosition(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	var pd *PositionData
	r.Error(pd.SetPosition(0, 0))

	pd = &PositionData{}
	r.NoError(pd.SetPosition(0, 0))
	r.Equal(0, pd.X)
	r.Equal(0, pd.Y)

	r.Error(pd.SetPosition(-1, 0))
	r.Error(pd.SetPosition(0, -1))
}

func Test_PositionData_SetSize(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	var pd *PositionData
	r.Error(pd.SetSize(0, 0))

	pd = &PositionData{}
	r.NoError(pd.SetSize(0, 0))
	r.Equal(0, pd.W)
	r.Equal(0, pd.H)

	r.Error(pd.SetSize(-1, 0))
	r.Error(pd.SetSize(0, -1))
}

func Test_PositionData_SetMaxSize(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	var pd *PositionData
	r.Error(pd.SetMaxSize(0, 0))

	pd = &PositionData{}
	r.NoError(pd.SetMaxSize(0, 0))
	r.Equal(0, pd.MaxW)
	r.Equal(0, pd.MaxH)

	r.Error(pd.SetMaxSize(-1, 0))
	r.Error(pd.SetMaxSize(0, -1))
}

func Test_PositionData_SetMinSize(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	var pd *PositionData
	r.Error(pd.SetMinSize(0, 0))

	pd = &PositionData{}
	r.NoError(pd.SetMinSize(0, 0))
	r.Equal(0, pd.MinW)
	r.Equal(0, pd.MinH)

	r.Error(pd.SetMinSize(-1, 0))
	r.Error(pd.SetMinSize(0, -1))
}

func Test_PositionData_PluginName(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	exp := "*windowx.PositionData"
	pd := &PositionData{}
	r.Equal(exp, pd.PluginName())

}
