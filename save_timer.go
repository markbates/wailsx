package wailsx

import (
	"context"
	"fmt"
	"time"
)

type SaveTimer struct {
	Duration      time.Duration // save duration, if zero, save once and exit
	DisableEvents bool          // disable save events
	Emitter       Emitter       // emit save events
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

	return st.Emitter.Emit(ctx, ev, data)
}

func (st SaveTimer) save(ctx context.Context, s Saver) (err error) {
	err = st.emit(ctx, EvtSaveTimerSaveStarted, nil)
	if err != nil {
		return err
	}

	defer st.emit(ctx, EvtSaveTimerSaveFinished, nil)
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
