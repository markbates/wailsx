package eventx

import (
	"context"
	"testing"

	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_EventsData_EmitEvent(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	const event = "test:event"

	ev := &EventsData{}

	err := ev.AddCallback(event, func(data ...any) error {
		return nil
	}, 0)

	r.NoError(err)

	err = ev.EmitEvent(event, wailstest.NowTime(), "test")
	r.NoError(err)

	r.Len(ev.Emitted[event], 1)
	r.Len(ev.Caught[event], 1)
}

func Test_EventsData_DisableStateData(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	const event = "test:event"

	ev := &EventsData{
		DisableStateData: true,
	}

	err := ev.AddCallback(event, func(data ...any) error {
		return nil
	}, 0)

	r.NoError(err)

	err = ev.EmitEvent(event, wailstest.NowTime(), "test")

	r.NoError(err)

	r.Len(ev.Emitted[event], 0)
	r.Len(ev.Caught[event], 0)

	sd, err := ev.StateData(context.Background())
	r.NoError(err)
	r.Nil(sd)

}
