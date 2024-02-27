package wailsx

import (
	"context"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Emitter_Emit(t *testing.T) {
	t.Parallel()
	r := require.New(t)
	ctx := context.Background()

	t.Run("error tests", func(t *testing.T) {
		etcs := []struct {
			name string
			fn   func(ctx context.Context, event string, args ...any) error
		}{
			{
				name: "returns an error",
				fn: func(ctx context.Context, event string, args ...any) error {
					return io.EOF
				},
			},
			{
				name: "panics",
				fn: func(ctx context.Context, event string, args ...any) error {
					panic(io.EOF)
				},
			},
		}

		for _, etc := range etcs {
			t.Run(etc.name, func(t *testing.T) {
				r := require.New(t)

				em := Emitter{
					EmitFn: etc.fn,
				}

				err := em.Emit(ctx, "test", "A")
				r.Error(err)
				r.True(errors.Is(err, io.EOF))
			})
		}
	})

	t.Run("marshal tests", func(t *testing.T) {
		t.Run("marshal error", func(t *testing.T) {
			r := require.New(t)

			em, ec := newEmitter()

			const name = "test"
			err := em.Emit(ctx, name, io.EOF)
			r.NoError(err)

			r.Len(ec.Events, 1)

			ev := ec.Events[0]
			r.Equal(name, ev.Event)
			r.Len(ev.Args, 1)

			msg, ok := ev.Args[0].(ErrorMessage)
			r.True(ok, "ec.Events[0] is not an ErrorMessage", ec.Events[0])
			r.True(errors.Is(msg.Err, io.EOF))
			r.Equal("test", msg.Event)
			r.Equal(io.EOF.Error(), msg.Text)
			r.Equal(nowTime(), msg.Time)
		})

		t.Run("marshal string", func(t *testing.T) {
			r := require.New(t)

			em, ec := newEmitter()

			const name = "test"
			err := em.Emit(ctx, name, "A")
			r.NoError(err)

			r.Len(ec.Events, 1)

			ev := ec.Events[0]
			r.Equal(name, ev.Event)
			r.Len(ev.Args, 1)

			msg, ok := ev.Args[0].(Message)
			r.True(ok, "ec.Events[0] is not a Message", ec.Events[0])
			r.Equal("A", msg.Text)
			r.Equal(nowTime(), msg.Time)
		})

		t.Run("marshal Messenger", func(t *testing.T) {
			r := require.New(t)

			em, ec := newEmitter()

			const name = "test"

			msg := Message{
				Event: name,
				Text:  "B",
				Time:  nowTime(),
				Data:  "C",
			}

			err := em.Emit(ctx, name, msg)
			r.NoError(err)

			r.Len(ec.Events, 1)

			ev := ec.Events[0]
			r.Equal(name, ev.Event)
			r.Len(ev.Args, 1)

			am, ok := ev.Args[0].(Messenger)
			r.True(ok, "ec.Events[0] is not a Message", ec.Events[0])
			r.Equal(name, am.MsgEvent())
			r.Equal("B", am.MsgText())
			r.Equal(nowTime(), am.MsgTime())
			r.Equal("C", am.MsgData())
		})

		t.Run("marshal any", func(t *testing.T) {
			r := require.New(t)

			em, ec := newEmitter()

			const name = "test"

			err := em.Emit(ctx, name, 1)
			r.NoError(err)

			r.Len(ec.Events, 1)

			ev := ec.Events[0]
			r.Equal(name, ev.Event)
			r.Len(ev.Args, 1)

			msg, ok := ev.Args[0].(Message)
			r.True(ok, "ec.Events[0] is not a Message", ec.Events[0])
			r.Equal(name, msg.Event)
			r.Equal(1, msg.Data)
			r.Equal(nowTime(), msg.Time)
		})
	})

	em, ec := newEmitter()

	em.EmitFn = func(ctx context.Context, event string, args ...any) error {
		return io.EOF
	}

	err := em.Emit(ctx, "test", "A")
	r.Error(err)

	r.Len(ec.Events, 0)

	em.EmitFn = ec.Emit

	err = em.Emit(ctx, "test", "A")
	r.NoError(err)

	err = em.Emit(ctx, "test", "B")
	r.NoError(err)

	err = em.Emit(ctx, "test", map[string]string{"1": "2"})
	r.NoError(err)

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

func Test_Emitter_Emit_error(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	em, ec := newEmitter()

	ctx := context.Background()
	em.EmitFn = func(ctx context.Context, event string, args ...any) error {
		return io.EOF
	}

	em.Emit(ctx, "test", "A")

	r.Len(ec.Events, 0)
}
