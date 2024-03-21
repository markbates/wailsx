package eventx

import (
	"context"
	"errors"
	"testing"

	"github.com/markbates/wailsx/eventx/msgx"
	"github.com/markbates/wailsx/wailsrun"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

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
			err:  wailsrun.ErrNotAvailable("EventsEmit"),
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

			ed := sd
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

			ed, err := em.StateData(ctx)
			r.NoError(err)

			r.Len(ed.Emitted, 1)
			r.Len(ed.Emitted[event], 1)

			evnt := ed.Emitted[event][0]

			msgs := evnt.Data
			r.Len(msgs, 1)

			msg := msgs[0]
			r.Equal(tc.exp.MsgEvent(), msg.MsgEvent())
			r.Equal(tc.exp.MsgText(), msg.MsgText())
			r.Equal(tc.exp.MsgTime(), msg.MsgTime())
			r.Equal(tc.exp.MsgData(), msg.MsgData())

		})
	}
}
