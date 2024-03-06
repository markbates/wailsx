package msgx

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

type complexError struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Err  error  `json:"err,omitempty"`
}

func (ce complexError) Error() string {
	if ce.Err == nil {
		return "nil error"
	}
	return fmt.Sprintf("complex error: %s", ce.Err.Error())
}

type marshalableError struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Err  error  `json:"err,omitempty"`
}

func (me marshalableError) Error() string {
	if me.Err == nil {
		return "nil error"
	}
	return fmt.Sprintf("marshalable error: %s", me.Err.Error())
}

func (me marshalableError) MarshalJSON() ([]byte, error) {
	if me.Err == nil {
		return nil, fmt.Errorf("error is required: %+v", me)
	}

	return json.Marshal(map[string]any{
		"id":   me.ID,
		"name": me.Name,
		"err":  me.Err.Error(),
	})
}

func Test_ErrorMessage(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	now := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	msg := ErrorMessage{
		Err: wailstest.ErrTest,
		Message: Message{
			Event: "my event",
			Text:  "my text",
			Time:  now,
			Data:  "my data",
		},
	}

	r.Equal("my event", msg.MsgEvent())
	r.Equal("my text", msg.MsgText())
	r.Equal(now, msg.MsgTime())
	r.Equal("my data", msg.MsgData())
	r.Equal(wailstest.ErrTest, msg.MsgError())
}

func Test_ErrorMessage_MarshalJSON(t *testing.T) {
	t.Parallel()

	ot := wailstest.OldTime()

	tcs := []struct {
		name string
		msg  ErrorMessage
		exp  string
		err  bool
	}{
		{
			name: "empty",
			msg:  ErrorMessage{},
			err:  true,
		},
		{
			name: "no_error",
			msg: ErrorMessage{
				Message: Message{
					Event: "no error",
				},
			},
			err: true,
		},
		{
			name: "no_time",
			msg: ErrorMessage{
				Err: wailstest.ErrTest,
				Message: Message{
					Event: "no time",
				},
			},
		},
		{
			name: "with_time",
			msg: ErrorMessage{
				Err: wailstest.ErrTest,
				Message: Message{
					Event: "with time",
					Time:  ot,
				},
			},
		},
		{
			name: "full",
			msg: ErrorMessage{
				Err: wailstest.ErrTest,
				Message: Message{
					Event: "full",
					Text:  "my text",
					Time:  ot,
					Data:  "my data",
				},
			},
		},
		{
			name: "complex_error",
			msg: ErrorMessage{
				Message: Message{
					Event: "complex error",
					Text:  "my text",
					Time:  ot,
					Data:  "my data",
				},
				Err: complexError{
					ID:   1,
					Name: "my name",
					Err:  wailstest.ErrTest,
				},
			},
		},
		{
			name: "marshalable_error",
			msg: ErrorMessage{
				Message: Message{
					Event: "marshalable error",
					Text:  "my text",
					Time:  ot,
					Data:  "my data",
				},
				Err: marshalableError{
					ID:   1,
					Name: "my name",
					Err:  wailstest.ErrTest,
				},
			},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			tc.msg.nowFn = wailstest.NowTime

			if tc.err {
				_, err := json.Marshal(tc.msg)
				r.Error(err)
				return
			}

			assertJSON(t, filepath.Join("error_msg", tc.name), tc.msg)
		})
	}

}
