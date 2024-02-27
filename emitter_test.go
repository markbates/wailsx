package wailsx

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Emitter_Emit(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	em, ec := newEmitter()

	ctx := context.Background()
	em.Emit(ctx, "test", "A")
	em.Emit(ctx, "test", "B")
	em.Emit(ctx, "test", map[string]string{"1": "2"})

	r.Len(ec.Events, 3)

	ev := ec.Events[0]
	r.Equal("test", ev.Event)
	r.Len(ev.Args, 1)

	am, ok := ev.Args[0].(Messenger)
	r.True(ok, "ev.Args[0] is not a Message", ev.Args[0])
	r.Equal("A", am.MsgText())

	ev = ec.Events[1]
	r.Equal("test", ev.Event)
	r.Len(ev.Args, 1)

	am, ok = ev.Args[0].(Messenger)
	r.True(ok, "ev.Args[0] is not a Message", ev.Args[0])
	r.Equal("B", am.MsgText())
}
