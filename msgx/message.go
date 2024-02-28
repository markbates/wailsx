package msgx

import (
	"encoding/json"
	"fmt"
	"time"
)

var _ Messenger = Message{}

type Message struct {
	Event string
	Text  string
	Time  time.Time
	Data  any

	nowFn func() time.Time // for testing
}

func (ee Message) MarshalJSON() ([]byte, error) {
	mm, err := ee.JSONMap()
	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(mm, "", "  ")
}

func (ee Message) MsgEvent() string {
	return ee.Event
}

func (ee Message) MsgText() string {
	return ee.Text
}

func (ee Message) MsgTime() time.Time {
	return ee.Time
}

func (ee Message) MsgData() any {
	return ee.Data
}

func (ee Message) JSONMap() (map[string]any, error) {
	if len(ee.Event) == 0 {
		return nil, fmt.Errorf("event is required: %+v", ee)
	}

	if ee.Time.IsZero() {
		fn := ee.nowFn
		if fn == nil {
			fn = time.Now
		}
		ee.Time = fn()
	}

	m := map[string]any{
		"event": ee.Event,
		"text":  ee.Text,
		"time":  ee.Time,
		"data":  ee.Data,
	}

	return m, nil
}
