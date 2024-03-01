package eventxtest

import (
	"context"
	"fmt"
	"sync"

	"github.com/markbates/wailsx/eventx"
	"github.com/markbates/wailsx/wailsrun"
)

var _ eventx.EventManager = &Manager{}

type Manager struct {
	Callbacks map[string]*CallbackCounter
	Emitted   map[string][]any // emitted events
	Caught    map[string][]any // caught events

	mu sync.Mutex
}

func NewManager() (*Manager, error) {
	m := &Manager{}
	if err := m.init(); err != nil {
		return nil, err
	}
	return m, nil
}

func (ev *Manager) Emit(ctx context.Context, event string, data ...any) error {
	if err := ev.init(); err != nil {
		return err
	}

	ev.mu.Lock()
	defer ev.mu.Unlock()

	ev.Emitted[event] = append(ev.Emitted[event], data...)

	cb, ok := ev.Callbacks[event]
	if !ok {
		return nil
	}

	ev.Caught[event] = append(ev.Caught[event], data...)

	if err := cb.Call(data...); err != nil {
		return err
	}

	ev.Callbacks[event] = cb

	return nil
}

func (ev *Manager) Off(ctx context.Context, event string, additional ...string) error {
	if err := ev.init(); err != nil {
		return err
	}

	ev.mu.Lock()
	defer ev.mu.Unlock()

	evts := append([]string{event}, additional...)
	for _, evt := range evts {
		if cc, ok := ev.Callbacks[evt]; ok {
			cc.Off = true
			ev.Callbacks[evt] = cc
		}
	}

	return nil
}

func (ev *Manager) OffAll(ctx context.Context) error {
	if err := ev.init(); err != nil {
		return err
	}

	ev.mu.Lock()
	defer ev.mu.Unlock()

	for event, cc := range ev.Callbacks {
		cc.Off = true
		ev.Callbacks[event] = cc
	}

	return nil
}

func (ev *Manager) On(ctx context.Context, event string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error) {
	if ev == nil {
		return nil, fmt.Errorf("event manager is nil")
	}

	ev.mu.Lock()
	defer ev.mu.Unlock()

	ev.Callbacks[event] = &CallbackCounter{
		Callback: callback,
	}

	fn := func() error {
		return ev.Off(ctx, event)
	}

	return fn, nil
}

func (ev *Manager) OnMultiple(ctx context.Context, event string, callback wailsrun.CallbackFn, counter int) (wailsrun.CancelFn, error) {
	if ev == nil {
		return nil, fmt.Errorf("event manager is nil")
	}

	ev.mu.Lock()
	defer ev.mu.Unlock()

	ev.Callbacks[event] = &CallbackCounter{
		Callback: callback,
		MaxCalls: counter,
	}

	fn := func() error {
		return ev.Off(ctx, event)
	}

	return fn, nil
}

func (ev *Manager) Once(ctx context.Context, event string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error) {
	return ev.OnMultiple(ctx, event, callback, 1)
}

func (ev *Manager) init() error {
	if ev == nil {
		return fmt.Errorf("event manager is nil")
	}

	ev.mu.Lock()
	defer ev.mu.Unlock()

	if ev.Callbacks == nil {
		ev.Callbacks = map[string]*CallbackCounter{}
	}

	if ev.Emitted == nil {
		ev.Emitted = map[string][]any{}
	}

	if ev.Caught == nil {
		ev.Caught = map[string][]any{}
	}

	return nil
}
