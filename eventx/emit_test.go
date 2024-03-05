package eventx

import (
	"context"
	"errors"
	"testing"

	"github.com/markbates/wailsx/msgx"
	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

// . "github.com/markbates/wailsx/eventx"

func Test_Manager_Emit(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	const event = "test:event"

	tcs := []struct {
		name string
		fn   func(ctx context.Context, event string, args ...any) error
		err  error
	}{
		{
			name: "no error",
			fn: func(ctx context.Context, event string, args ...any) error {
				return nil
			},
		},
		{
			name: "returns an error",
			fn: func(ctx context.Context, event string, args ...any) error {
				return wailstest.ErrTest
			},
			err: wailstest.ErrTest,
		},
		{
			name: "panics",
			fn: func(ctx context.Context, event string, args ...any) error {
				panic(wailstest.ErrTest)
			},
			err: wailstest.ErrTest,
		},
		{
			name: "nil function",
			err:  wailsrun.ErrNotAvailable,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			em := Manager{
				EventsEmitFn: tc.fn,
			}

			err := em.EventsEmit(ctx, event, 42)

			if tc.err != nil {
				r.Error(err)
				r.True(errors.Is(err, tc.err))
				return
			}

			r.NoError(err)

			sd, err := em.StateData(ctx)
			r.NoError(err)

			ed := sd.Data
			r.Len(ed.Emitted, 2)
			r.Len(ed.Caught, 0)
		})
	}
}

func Test_Manager_Emit_Args(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	const event = "test:event"

	now := wailstest.NowTime()

	tcs := []struct {
		name string
		args []any
		exp  msgx.Messenger
	}{
		{
			name: "no args",
			exp: msgx.Message{
				Event: event,
				Text:  event,
				Time:  now,
				Data:  event,
			},
		},
		{
			name: "string",
			args: []any{"A"},
			exp: msgx.Message{
				Event: event,
				Text:  "A",
				Time:  now,
				Data:  "A",
			},
		},
		{
			name: "error",
			args: []any{wailstest.ErrTest},
			exp: msgx.ErrorMessage{
				Err: wailstest.ErrTest,
				Message: msgx.Message{
					Event: event,
					Text:  wailstest.ErrTest.Error(),
					Time:  now,
					Data:  wailstest.ErrTest,
				},
			},
		},
		{
			name: "Messenger",
			args: []any{msgx.Message{
				Event: event,
				Text:  "B",
				Time:  now,
				Data:  "C",
			}},
			exp: msgx.Message{
				Event: event,
				Text:  "B",
				Time:  now,
				Data:  "C",
			},
		},
		{
			name: "int",
			args: []any{1},
			exp: msgx.Message{
				Event: event,
				Data:  1,
				Time:  now,
			},
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)
			em := newTestManager()
			em.EventsEmitFn = func(ctx context.Context, event string, args ...any) error {
				return nil
			}

			err := em.EventsEmit(ctx, event, tc.args...)
			r.NoError(err)

			sd, err := em.StateData(ctx)
			r.NoError(err)

			ed := sd.Data
			r.Len(ed.Emitted, 1)
			r.Len(ed.Emitted[event], 1)

			evnt := ed.Emitted[event][0]

			msg, ok := evnt.(msgx.Messenger)
			r.True(ok, "evnt is not a Messenger", evnt)

			r.Equal(tc.exp.MsgEvent(), msg.MsgEvent())
			r.Equal(tc.exp.MsgText(), msg.MsgText())
			r.Equal(tc.exp.MsgTime(), msg.MsgTime())
			r.Equal(tc.exp.MsgData(), msg.MsgData())

			// args := ed.Emitted[event][0]
		})
	}
}

// func Test_EventManager_Emit(t *testing.T) {
// 	t.Parallel()
// 	r := require.New(t)
// 	ctx := context.Background()

// 	t.Run("error tests", func(t *testing.T) {
// 		etcs := []struct {
// 			name string
// 			fn   func(ctx context.Context, event string, args ...any) error
// 		}{
// 			{
// 				name: "returns an error",
// 				fn: func(ctx context.Context, event string, args ...any) error {
// 					return wailstest.ErrTest
// 				},
// 			},
// 			{
// 				name: "panics",
// 				fn: func(ctx context.Context, event string, args ...any) error {
// 					panic(wailstest.ErrTest)
// 				},
// 			},
// 		}

// 		for _, etc := range etcs {
// 			t.Run(etc.name, func(t *testing.T) {
// 				r := require.New(t)

// 				em := Manager{
// 					EventsEmitFn: etc.fn,
// 				}

// 				err := em.EventsEmit(ctx, "test", "A")
// 				r.Error(err)
// 				r.True(errors.Is(err, wailstest.ErrTest))
// 			})
// 		}
// 	})

// 	t.Run("marshal tests", func(t *testing.T) {
// 		t.Run("marshal error", func(t *testing.T) {
// 			r := require.New(t)

// 			em := newEventManager()

// 			const name = "test"
// 			err := em.EventsEmit(ctx, name, wailstest.ErrTest)
// 			r.NoError(err)

// 			sd, err := em.StateData(context.Background())
// 			r.NoError(err)

// 			ec := sd.Data
// 			r.Len(ec.Emitted, 1)

// 			ev := ec.Emitted[name][0]
// 			r.Len(ev, 1)

// 			msg, ok := ev.(msgx.ErrorMessage)
// 			r.True(ok, "ec.Events[0] is not an ErrorMessage", ev)
// 			r.True(errors.Is(msg.Err, wailstest.ErrTest))
// 			r.Equal("test", msg.Event)
// 			r.Equal(wailstest.ErrTest.Error(), msg.Text)
// 			r.Equal(wailstest.NowTime(), msg.Time)
// 		})

// 		t.Run("marshal string", func(t *testing.T) {
// 			r := require.New(t)

// 			em, ec := newEventManager()

// 			const name = "test"
// 			err := em.EventsEmit(ctx, name, "A")
// 			r.NoError(err)

// 			r.Len(ec.Events, 1)

// 			ev := ec.Events[0]
// 			r.Equal(name, ev.Event)
// 			r.Len(ev.Args, 1)

// 			msg, ok := ev.Args[0].(msgx.Message)
// 			r.True(ok, "ec.Events[0] is not a Message", ec.Events[0])
// 			r.Equal("A", msg.Text)
// 			r.Equal(wailstest.NowTime(), msg.Time)
// 		})

// 		t.Run("marshal Messenger", func(t *testing.T) {
// 			r := require.New(t)

// 			em, ec := newEventManager()

// 			const name = "test"

// 			msg := msgx.Message{
// 				Event: name,
// 				Text:  "B",
// 				Time:  wailstest.NowTime(),
// 				Data:  "C",
// 			}

// 			err := em.EventsEmit(ctx, name, msg)
// 			r.NoError(err)

// 			r.Len(ec.Events, 1)

// 			ev := ec.Events[0]
// 			r.Equal(name, ev.Event)
// 			r.Len(ev.Args, 1)

// 			am, ok := ev.Args[0].(msgx.Messenger)
// 			r.True(ok, "ec.Events[0] is not a Message", ec.Events[0])
// 			r.Equal(name, am.MsgEvent())
// 			r.Equal("B", am.MsgText())
// 			r.Equal(wailstest.NowTime(), am.MsgTime())
// 			r.Equal("C", am.MsgData())
// 		})

// 		t.Run("marshal any", func(t *testing.T) {
// 			r := require.New(t)

// 			em, ec := newEventManager()

// 			const name = "test"

// 			err := em.EventsEmit(ctx, name, 1)
// 			r.NoError(err)

// 			r.Len(ec.Events, 1)

// 			ev := ec.Events[0]
// 			r.Equal(name, ev.Event)
// 			r.Len(ev.Args, 1)

// 			msg, ok := ev.Args[0].(msgx.Message)
// 			r.True(ok, "ec.Events[0] is not a Message", ec.Events[0])
// 			r.Equal(name, msg.Event)
// 			r.Equal(1, msg.Data)
// 			r.Equal(wailstest.NowTime(), msg.Time)
// 		})
// 	})

// 	em, ec := newEventManager()

// 	em.EventsEmitFn = func(ctx context.Context, event string, args ...any) error {
// 		return wailstest.ErrTest
// 	}

// 	err := em.EventsEmit(ctx, "test", "A")
// 	r.Error(err)

// 	r.Len(ec.Events, 0)

// 	em.EventsEmitFn = ec.Emit

// 	err = em.EventsEmit(ctx, "test", "A")
// 	r.NoError(err)

// 	err = em.EventsEmit(ctx, "test", "B")
// 	r.NoError(err)

// 	err = em.EventsEmit(ctx, "test", map[string]string{"1": "2"})
// 	r.NoError(err)

// 	r.Len(ec.Events, 3)

// 	ev := ec.Events[0]
// 	r.Equal("test", ev.Event)
// 	r.Len(ev.Args, 1)

// 	am, ok := ev.Args[0].(msgx.Messenger)
// 	r.True(ok, "ev.Args[0] is not a Message", ev.Args[0])
// 	r.Equal("A", am.MsgText())

// 	ev = ec.Events[1]
// 	r.Equal("test", ev.Event)
// 	r.Len(ev.Args, 1)

// 	am, ok = ev.Args[0].(msgx.Messenger)
// 	r.True(ok, "ev.Args[0] is not a Message", ev.Args[0])
// 	r.Equal("B", am.MsgText())
// }

// func Test_EventManager_Emit_error(t *testing.T) {
// 	t.Parallel()

// 	r := require.New(t)

// 	em, ec := newEventManager()

// 	ctx := context.Background()
// 	em.EventsEmitFn = func(ctx context.Context, event string, args ...any) error {
// 		return wailstest.ErrTest
// 	}

// 	err := em.EventsEmit(ctx, "test", "A")
// 	r.Error(err)
// 	r.True(errors.Is(err, wailstest.ErrTest))

// 	r.Len(ec.Events, 0)
// }

// func Test_EventManager_Emit_Panic(t *testing.T) {
// 	t.Parallel()

// 	r := require.New(t)

// 	em, ec := newEventManager()

// 	ctx := context.Background()
// 	em.EventsEmitFn = func(ctx context.Context, event string, args ...any) error {
// 		panic(wailstest.ErrTest)
// 	}

// 	err := em.EventsEmit(ctx, "test", "A")
// 	r.Error(err)
// 	r.True(errors.Is(err, wailstest.ErrTest))

// 	r.Len(ec.Events, 0)
// }
