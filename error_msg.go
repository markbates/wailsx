package wailsx

import (
	"encoding/json"
	"fmt"
	"time"
)

type ErrorMessenger interface {
	MsgError() error
	MsgEvent() string
	MsgText() string
	MsgTime() time.Time
	MsgData() any
}

var _ ErrorMessenger = ErrorMessage{}

type ErrorMessage struct {
	Message
	Err error
}

func (ee ErrorMessage) MsgError() error {
	return ee.Err
}

func (ee ErrorMessage) MarshalJSON() ([]byte, error) {
	mm, err := ee.JSONMap()
	if err != nil {
		return nil, err
	}

	return json.Marshal(mm)
}

func (ee ErrorMessage) JSONMap() (map[string]any, error) {
	mm, err := ee.Message.JSONMap()
	if err != nil {
		return nil, err
	}

	if ee.Err == nil {
		return nil, fmt.Errorf("error is required: %+v", ee)
	}

	mm["error"] = ee.Err.Error()
	if len(ee.Text) == 0 {
		mm["text"] = ee.Err.Error()
	}

	switch t := ee.Err.(type) {
	case json.Marshaler:
		mm["error"] = t
	default:
		b, err := json.Marshal(ee.Err)
		if err != nil {
			return nil, err
		}

		s := string(b)
		if len(s) > 0 && s != "{}" {
			mm["error"] = s
		}
	}

	return mm, nil
}
