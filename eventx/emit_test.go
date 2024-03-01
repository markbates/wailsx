package eventx_test

import (
	"context"
	"errors"
	"testing"

	. "github.com/markbates/wailsx/eventx"
	"github.com/markbates/wailsx/msgx"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_EventManager_Emit(t *testing.T) {
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
					return wailstest.ErrTest
				},
			},
			{
				name: "panics",
				fn: func(ctx context.Context, event string, args ...any) error {
					panic(wailstest.ErrTest)
				},
			},
		}

		for _, etc := range etcs {
			t.Run(etc.name, func(t *testing.T) {
				r := require.New(t)

				em := Manager{
					EmitFn: etc.fn,
				}

				err := em.EventsEmit(ctx, "test", "A")
				r.Error(err)
				r.True(errors.Is(err, wailstest.ErrTest))
			})
		}
	})

	t.Run("marshal tests", func(t *testing.T) {
		t.Run("marshal error", func(t *testing.T) {
			r := require.New(t)

			em, ec := newEventManager()

			const name = "test"
			err := em.EventsEmit(ctx, name, wailstest.ErrTest)
			r.NoError(err)

			r.Len(ec.Events, 1)

			ev := ec.Events[0]
			r.Equal(name, ev.Event)
			r.Len(ev.Args, 1)

			msg, ok := ev.Args[0].(msgx.ErrorMessage)
			r.True(ok, "ec.Events[0] is not an ErrorMessage", ec.Events[0])
			r.True(errors.Is(msg.Err, wailstest.ErrTest))
			r.Equal("test", msg.Event)
			r.Equal(wailstest.ErrTest.Error(), msg.Text)
			r.Equal(wailstest.NowTime(), msg.Time)
		})

		t.Run("marshal string", func(t *testing.T) {
			r := require.New(t)

			em, ec := newEventManager()

			const name = "test"
			err := em.EventsEmit(ctx, name, "A")
			r.NoError(err)

			r.Len(ec.Events, 1)

			ev := ec.Events[0]
			r.Equal(name, ev.Event)
			r.Len(ev.Args, 1)

			msg, ok := ev.Args[0].(msgx.Message)
			r.True(ok, "ec.Events[0] is not a Message", ec.Events[0])
			r.Equal("A", msg.Text)
			r.Equal(wailstest.NowTime(), msg.Time)
		})

		t.Run("marshal Messenger", func(t *testing.T) {
			r := require.New(t)

			em, ec := newEventManager()

			const name = "test"

			msg := msgx.Message{
				Event: name,
				Text:  "B",
				Time:  wailstest.NowTime(),
				Data:  "C",
			}

			err := em.EventsEmit(ctx, name, msg)
			r.NoError(err)

			r.Len(ec.Events, 1)

			ev := ec.Events[0]
			r.Equal(name, ev.Event)
			r.Len(ev.Args, 1)

			am, ok := ev.Args[0].(msgx.Messenger)
			r.True(ok, "ec.Events[0] is not a Message", ec.Events[0])
			r.Equal(name, am.MsgEvent())
			r.Equal("B", am.MsgText())
			r.Equal(wailstest.NowTime(), am.MsgTime())
			r.Equal("C", am.MsgData())
		})

		t.Run("marshal any", func(t *testing.T) {
			r := require.New(t)

			em, ec := newEventManager()

			const name = "test"

			err := em.EventsEmit(ctx, name, 1)
			r.NoError(err)

			r.Len(ec.Events, 1)

			ev := ec.Events[0]
			r.Equal(name, ev.Event)
			r.Len(ev.Args, 1)

			msg, ok := ev.Args[0].(msgx.Message)
			r.True(ok, "ec.Events[0] is not a Message", ec.Events[0])
			r.Equal(name, msg.Event)
			r.Equal(1, msg.Data)
			r.Equal(wailstest.NowTime(), msg.Time)
		})
	})

	em, ec := newEventManager()

	em.EmitFn = func(ctx context.Context, event string, args ...any) error {
		return wailstest.ErrTest
	}

	err := em.EventsEmit(ctx, "test", "A")
	r.Error(err)

	r.Len(ec.Events, 0)

	em.EmitFn = ec.Emit

	err = em.EventsEmit(ctx, "test", "A")
	r.NoError(err)

	err = em.EventsEmit(ctx, "test", "B")
	r.NoError(err)

	err = em.EventsEmit(ctx, "test", map[string]string{"1": "2"})
	r.NoError(err)

	r.Len(ec.Events, 3)

	ev := ec.Events[0]
	r.Equal("test", ev.Event)
	r.Len(ev.Args, 1)

	am, ok := ev.Args[0].(msgx.Messenger)
	r.True(ok, "ev.Args[0] is not a Message", ev.Args[0])
	r.Equal("A", am.MsgText())

	ev = ec.Events[1]
	r.Equal("test", ev.Event)
	r.Len(ev.Args, 1)

	am, ok = ev.Args[0].(msgx.Messenger)
	r.True(ok, "ev.Args[0] is not a Message", ev.Args[0])
	r.Equal("B", am.MsgText())
}

func Test_EventManager_Emit_error(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	em, ec := newEventManager()

	ctx := context.Background()
	em.EmitFn = func(ctx context.Context, event string, args ...any) error {
		return wailstest.ErrTest
	}

	err := em.EventsEmit(ctx, "test", "A")
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))

	r.Len(ec.Events, 0)
}

func Test_EventManager_Emit_Panic(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	em, ec := newEventManager()

	ctx := context.Background()
	em.EmitFn = func(ctx context.Context, event string, args ...any) error {
		panic(wailstest.ErrTest)
	}

	err := em.EventsEmit(ctx, "test", "A")
	r.Error(err)
	r.True(errors.Is(err, wailstest.ErrTest))

	r.Len(ec.Events, 0)
}
