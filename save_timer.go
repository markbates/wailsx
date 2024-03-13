package wailsx

import (
	"context"
	"fmt"
	"time"

	"github.com/markbates/wailsx/eventx"
)

func NewAppSaveTimer(app *App, d time.Duration) (SaveTimer, error) {
	if app == nil {
		return SaveTimer{}, fmt.Errorf("app is nil")
	}

	return SaveTimer{
		Duration: d,
		Manager:  app.EventManager,
		DataFn: func(ctx context.Context) (any, error) {
			return app.StateData(ctx)
		},
	}, nil
}

type SaveTimer struct {
	Duration      time.Duration                          `json:"duration,omitempty"`       // save duration, if zero, save once and exit
	DisableEvents bool                                   `json:"disable_events,omitempty"` // disable save events
	Manager       eventx.EventManager                    `json:"manager,omitempty"`        // emit save events
	DataFn        func(ctx context.Context) (any, error) `json:"-"`                        // data to emit with save events
}

func (st SaveTimer) Save(ctx context.Context, s Saver) error {
	if s == nil {
		return fmt.Errorf("saver is nil")
	}

	if st.Duration == 0 {
		return st.save(ctx, s)
	}

	tick := time.NewTicker(st.Duration)
	defer tick.Stop()

	for {
		select {
		case <-tick.C:
			err := st.save(ctx, s)
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return nil
		}
	}
}

func (st SaveTimer) emit(ctx context.Context, ev string, data any) error {
	if st.DisableEvents {
		return nil
	}

	return st.Manager.EventsEmit(ctx, ev, data)
}

func (st SaveTimer) save(ctx context.Context, s Saver) (err error) {
	var data any
	if st.DataFn != nil {
		data, err = st.DataFn(ctx)
		if err != nil {
			return err
		}
	}

	err = st.emit(ctx, EvtSaveTimerSaveStarted, data)
	if err != nil {
		return err
	}

	defer st.emit(ctx, EvtSaveTimerSaveFinished, data)
	defer func() {
		if err != nil {
			st.emit(ctx, EvtSaveTimerSaveError, err)
		}
	}()

	err = s.Save(ctx)
	if err != nil {
		return err
	}

	return nil
}
