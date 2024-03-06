package msgx

import (
	"testing"

	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_Messengers_Any(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	event := "event 1"
	text := "text 1"
	now := wailstest.NowTime()
	data := "data 1"

	msgs := Messengers{
		Message{
			Event: event,
			Text:  text,
			Time:  now,
			Data:  data,
		},
	}

	r.Len(msgs, 1)

	all := msgs.Any()
	r.Len(all, 1)

	msg, ok := all[0].(Message)
	r.True(ok)

	r.Equal(event, msg.Event)
	r.Equal(text, msg.Text)
	r.Equal(now, msg.Time)
	r.Equal(data, msg.Data)

}
