package wailsx

import (
	"context"
	"testing"
	"time"

	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_SaveTimer_Save_Loop(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	emitter, ec := newEmitter()

	st := SaveTimer{
		Duration: 2 * time.Millisecond,
		Emitter:  emitter,
	}

	s, err := NewState("save timer")
	r.NoError(err)
	s.SaveFn = func(ctx context.Context) error {
		cancel()
		return nil
	}

	err = st.Save(ctx, s)
	r.NoError(err)

	<-ctx.Done()

	r.Equal(len(ec.Events), 2)
	r.Equal(EvtSaveTimerSaveStarted, ec.Events[0].Event)
	r.Equal(EvtSaveTimerSaveFinished, ec.Events[1].Event)

}

func Test_SaveTimer_Save_Once(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	emitter, ec := newEmitter()

	st := SaveTimer{
		Duration: 0,
		Emitter:  emitter,
	}

	s, err := NewState("save timer")
	r.NoError(err)
	s.SaveFn = func(ctx context.Context) error {
		cancel()
		return nil
	}

	err = st.Save(ctx, s)
	r.NoError(err)

	<-ctx.Done()

	r.Equal(2, len(ec.Events))
	r.Equal(EvtSaveTimerSaveStarted, ec.Events[0].Event)
	r.Equal(EvtSaveTimerSaveFinished, ec.Events[1].Event)
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
				return wailstest.ERR
			},
		},
		{
			name: "panic error",
			fn: func(ctx context.Context) error {
				panic(wailstest.ERR)
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			emitter, ec := newEmitter()

			st := SaveTimer{
				Duration: 0,
				Emitter:  emitter,
			}

			s, err := NewState("save timer")
			r.NoError(err)
			s.SaveFn = func(ctx context.Context) error {
				cancel()
				return wailstest.ERR
			}

			err = st.Save(ctx, s)
			r.Error(err)

			<-ctx.Done()

			r.Equal(3, len(ec.Events))
			r.Equal(EvtSaveTimerSaveStarted, ec.Events[0].Event)
			r.Equal(EvtSaveTimerSaveError, ec.Events[1].Event)
		})

	}
}
