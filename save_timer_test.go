package wailsx

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/markbates/wailsx/eventx"
	"github.com/markbates/wailsx/statedata"
	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_SaveTimer_Save_Loop(t *testing.T) {

	t.Parallel()
	r := require.New(t)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	app, err := NopApp("test save timer")
	r.NoError(err)

	em := eventx.NopManager()

	st := SaveTimer{
		Duration: 2 * time.Millisecond,
		Manager:  em,
		DataFn: func(ctx context.Context) (any, error) {
			return app.StateData(ctx)
		},
	}

	s, err := NopApp("save timer")
	r.NoError(err)
	s.SaveFn = func(ctx context.Context) error {
		cancel()
		return nil
	}

	err = st.Save(ctx, s)
	r.NoError(err)

	<-ctx.Done()

	sd, err := em.StateData(ctx)
	r.NoError(err)

	data := sd.Data
	r.Len(data.Emitted, 3)

	evnts := data.Emitted[EvtSaveTimerSaveStarted]
	r.Len(evnts, 1)

	evt := evnts[0]
	r.NotNil(evt)
	r.NotNil(evt.Data)

	msgs := evt.Data
	r.Len(msgs, 1)

	msg := msgs[0]
	r.NotNil(msg)
	_, ok := msg.MsgData().(statedata.Data[AppData])
	r.True(ok, fmt.Sprintf("%T", msg.MsgData()))

	evnts = data.Emitted[EvtSaveTimerSaveFinished]
	r.Len(evnts, 1)
}

func Test_SaveTimer_Save_Once(t *testing.T) {

	t.Parallel()
	r := require.New(t)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	em := eventx.NopManager()

	st := SaveTimer{
		Duration: 0,
		Manager:  em,
	}

	s, err := NewApp("save timer")
	r.NoError(err)
	s.SaveFn = func(ctx context.Context) error {
		cancel()
		return nil
	}

	err = st.Save(ctx, s)
	r.NoError(err)

	<-ctx.Done()

	sd, err := em.StateData(ctx)
	r.NoError(err)

	data := sd.Data
	r.GreaterOrEqual(len(data.Emitted), 2)

	evt := data.Emitted[EvtSaveTimerSaveStarted]
	r.Len(evt, 1)

	evt = data.Emitted[EvtSaveTimerSaveFinished]
	r.Len(evt, 1)
}

func Test_SaveTimer_Save_Error(t *testing.T) {

	t.Parallel()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tcs := []struct {
		name string
		fn   func(context.Context) error
	}{
		{
			name: "error",
			fn: func(ctx context.Context) error {
				cancel()
				return wailstest.ErrTest
			},
		},
		{
			name: "panic error",
			fn: func(ctx context.Context) error {
				panic(wailstest.ErrTest)
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			em := eventx.NopManager()

			st := SaveTimer{
				Duration: 0,
				Manager:  em,
			}

			s, err := NewApp("save timer")
			r.NoError(err)
			s.SaveFn = func(ctx context.Context) error {
				cancel()
				return wailstest.ErrTest
			}

			err = st.Save(ctx, s)
			r.Error(err)

			<-ctx.Done()

			sd, err := em.StateData(ctx)
			r.NoError(err)

			data := sd.Data
			r.Len(data.Emitted, 4)

			evt := data.Emitted[EvtSaveTimerSaveStarted]
			r.Len(evt, 1)

			evt = data.Emitted[EvtSaveTimerSaveFinished]
			r.Len(evt, 1)

		})

	}
}
