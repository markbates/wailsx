package wailsx

import (
	"context"
	"encoding/json"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_API_StateData(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	api := NopAPI()

	ctx := context.Background()

	const (
		maxW = 1200
		maxH = 800
		minW = 20
		minH = 30
		w    = 800
		h    = 600
		x    = 100
		y    = 200
	)

	err := api.WindowSetMaxSize(ctx, maxW, maxH)
	r.NoError(err)

	err = api.WindowSetMinSize(ctx, minW, minH)
	r.NoError(err)

	err = api.WindowSetPosition(ctx, x, y)
	r.NoError(err)

	err = api.WindowSetSize(ctx, w, h)
	r.NoError(err)

	err = api.WindowSetBackgroundColour(ctx, 1, 2, 3, 4)
	r.NoError(err)

	err = api.WindowSetDarkTheme(ctx)
	r.NoError(err)

	err = api.WindowMaximise(ctx)
	r.NoError(err)

	const event = "event:test"

	cancel, err := api.EventsOn(ctx, event, func(data ...any) error {
		r.Len(data, 1)
		r.Equal(42, data[0])
		return nil
	})
	r.NoError(err)
	defer cancel()

	err = api.EventsEmit(ctx, event, 42)
	r.NoError(err)

	sd, err := api.StateData(ctx)
	r.NoError(err)

	r.NotNil(sd.Data)
	r.Equal(APIStateDataProviderName, sd.Name)

	ed := sd.Data.Events
	r.NotNil(ed)

	wd := sd.Data.Window
	r.NotNil(wd)

	r.Equal(h, wd.H)
	r.Equal(maxH, wd.MaxH)
	r.Equal(maxW, wd.MaxW)
	r.Equal(minH, wd.MinH)
	r.Equal(minW, wd.MinW)
	r.Equal(w, wd.W)
	r.Equal(x, wd.X)
	r.Equal(y, wd.Y)

	b, err := json.MarshalIndent(sd, "", "  ")
	r.NoError(err)

	act := string(b)
	act = strings.TrimSpace(act)

	// fmt.Println(act)

	// f, err := os.Create("testdata/api.json")
	// r.NoError(err)
	// f.WriteString(act)
	// r.NoError(f.Close())

	b, err = os.ReadFile("testdata/api.json")
	r.NoError(err)

	exp := string(b)
	exp = strings.TrimSpace(exp)

	r.Equal(exp, act)

}
