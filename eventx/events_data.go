package eventx

import (
	"context"
	"fmt"
	"sync"

	"github.com/markbates/wailsx/statedata"
	"github.com/markbates/wailsx/wailsrun"
)

var _ statedata.StateDataProvider[*EventsData] = &EventsData{}

type EventsData struct {
	Callbacks map[string]*CallbackCounter `json:"callbacks"`
	Emitted   map[string][]any            `json:"emitted"` // emitted events
	Caught    map[string][]any            `json:"caught"`  // caught events

	mu sync.Mutex
}

func (ev *EventsData) EmitEvent(event string, data ...any) error {
	if err := ev.init(); err != nil {
		return err
	}

	ev.mu.Lock()
	defer ev.mu.Unlock()

	ev.Emitted[event] = append(ev.Emitted[event], data...)

	cc, ok := ev.Callbacks[event]
	if !ok {
		return nil
	}

	b, err := cc.Catch(data...)
	if err != nil {
		return err
	}

	if b {
		ev.Caught[event] = append(ev.Caught[event], data...)
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

func (ev *EventsData) AddCallback(event string, cb wailsrun.CallbackFn, max int) error {
	if err := ev.init(); err != nil {
		return err
	}

	ev.mu.Lock()
	defer ev.mu.Unlock()

	cc := &CallbackCounter{
		Callback: cb,
		MaxCalls: max,
	}

	ev.Callbacks[event] = cc
	return nil
}

func (ev *EventsData) StateData(ctx context.Context) (statedata.StateData[*EventsData], error) {
	sd := statedata.StateData[*EventsData]{
		Name: EventManagerStateDataName,
		Data: ev,
	}

	if err := ev.init(); err != nil {
		return sd, err
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
		ev.Emitted = map[string][]any{}
	}

	if ev.Caught == nil {
		ev.Caught = map[string][]any{}
	}
	return nil
}
