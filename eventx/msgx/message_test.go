package msgx

import (
	"encoding/json"
	"path/filepath"
	"testing"
	"time"

	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_Message(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	now := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	msg := Message{
		Event: "my event",
		Text:  "my text",
		Time:  now,
		Data:  "my data",
	}

	r.Equal("my event", msg.MsgEvent())
	r.Equal("my text", msg.MsgText())
	r.Equal(now, msg.MsgTime())
	r.Equal("my data", msg.MsgData())

}

func Test_Message_MarshalJSON(t *testing.T) {
	t.Parallel()

	ot := wailstest.OldTime()

	tcs := []struct {
		name string
		msg  Message
		err  bool
	}{
		{
			name: "empty",
			msg:  Message{},
			err:  true,
		},
		{
			name: "no_time",
			msg: Message{
				Event: "no time",
			},
		},
		{
			name: "with_time",
			msg: Message{
				Event: "with time",
				Time:  ot,
			},
		},
		{
			name: "full",
			msg: Message{
				Event: "full",
				Text:  "my text",
				Time:  ot,
				Data:  "my data",
			},
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			r := require.New(t)

			tc.msg.nowFn = wailstest.NowTime

			if tc.err {
				_, err := json.Marshal(tc.msg)
				r.Error(err)
				return
			}

			assertJSON(t, filepath.Join("messages", tc.name), tc.msg)
		})
	}
}
