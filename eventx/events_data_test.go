package eventx

import (
	"testing"

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

	err = ev.EmitEvent(event, "test")
	r.NoError(err)

	r.Len(ev.Emitted[event], 1)
	r.Len(ev.Caught[event], 1)
}
