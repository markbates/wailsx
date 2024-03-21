package eventx

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/markbates/wailsx/statedata"
)

var _ statedata.DataProvider[*EventsData] = &EventsData{}

type EventsData struct {
	DisableStateData     bool `json:"disableStateData,omitempty"`
	DisableWildcardEmits bool `json:"disableWildcardEmits,omitempty"`

	Callbacks map[string]*CallbackCounter `json:"callbacks,omitempty"`
	Emitted   map[string][]Event          `json:"emitted,omitempty"` // emitted events
	Caught    map[string][]Event          `json:"caught,omitempty"`  // caught events

	mu sync.RWMutex
}

func (ev *EventsData) EmitEvent(event string, now time.Time, data ...any) error {
	if err := ev.init(); err != nil {
		return err
	}

	if ev.DisableStateData {
		return nil
	}

	if now.IsZero() {
		now = time.Now()
	}

	ev.mu.Lock()
	defer ev.mu.Unlock()

	envt, err := NewEvent(event, now, data...)
	if err != nil {
		return err
	}

	ev.Emitted[event] = append(ev.Emitted[event], envt)

	cc, ok := ev.Callbacks[event]
	if !ok {
		return nil
	}

	b, err := cc.Catch(data...)
	if err != nil {
		return err
	}

	if b {
		ev.Caught[event] = append(ev.Caught[event], envt)
	}

	return nil
}

func (ev *EventsData) CallbacksOffAll() error {
	if err := ev.init(); err != nil {
		return err
	}

	ev.mu.Lock()
	defer ev.mu.Unlock()

	for _, cc := range ev.Callbacks {
		cc.Off = true
	}
	return nil
}

func (ev *EventsData) CallbacksOff(events ...string) error {
	if err := ev.init(); err != nil {
		return err
	}

	ev.mu.Lock()
	defer ev.mu.Unlock()

	for _, event := range events {
		if cc, ok := ev.Callbacks[event]; ok {
			cc.Off = true
		}
	}
	return nil
}

func (ev *EventsData) AddCallback(event string, cb CallbackFn, max int) error {
	if err := ev.init(); err != nil {
		return err
	}

	ev.mu.Lock()
	defer ev.mu.Unlock()

	cc := &CallbackCounter{
		MaxCalls: max,
	}

	ev.Callbacks[event] = cc
	return nil
}

func (ev *EventsData) StateData(ctx context.Context) (*EventsData, error) {
	if err := ev.init(); err != nil {
		return nil, err
	}

	ev.mu.RLock()
	defer ev.mu.RUnlock()

	if ev.DisableStateData {
		return nil, nil
	}

	sd := &EventsData{
		DisableStateData:     ev.DisableStateData,
		DisableWildcardEmits: ev.DisableWildcardEmits,
		Callbacks:            ev.Callbacks,
		Emitted:              ev.Emitted,
		Caught:               ev.Caught,
	}

	return sd, nil
}

func (ev *EventsData) init() error {
	if ev == nil {
		return fmt.Errorf("events data is nil")
	}

	ev.mu.Lock()
	defer ev.mu.Unlock()

	if ev.Callbacks == nil {
		ev.Callbacks = map[string]*CallbackCounter{}
	}

	if ev.Emitted == nil {
		ev.Emitted = map[string][]Event{}
	}

	if ev.Caught == nil {
		ev.Caught = map[string][]Event{}
	}

	return nil
}
