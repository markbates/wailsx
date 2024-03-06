package msgx

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func assertJSON(t testing.TB, fp string, data any) {
	t.Helper()

	r := require.New(t)

	b, err := json.MarshalIndent(data, "", "  ")
	r.NoError(err)

	act := string(b)
	act = strings.TrimSpace(act)

	// fmt.Println(act)

	fp = filepath.Join("testdata", fp+".json")

	os.MkdirAll(filepath.Dir(fp), 0755)
	f, err := os.Create(fp)
	r.NoError(err)
	f.Write([]byte(act))
	r.NoError(f.Close())

	b, err = os.ReadFile(fp)
	r.NoError(err)

	exp := string(b)
	exp = strings.TrimSpace(exp)

	r.Equal(exp, act)
}

func Test_NewMessage(t *testing.T) {
	t.Parallel()

	now := wailstest.NowTime()
	name := "events:test"

	tcs := []struct {
		name string
		arg  any
		exp  Messenger
	}{
		{
			name: "string",
			arg:  "hello",
			exp: Message{
				Event: name,
				Time:  now,
				Text:  "hello",
				Data:  "hello",
			},
		},
		{
			name: "error",
			arg:  wailstest.ErrTest,
			exp: ErrorMessage{
				Err: wailstest.ErrTest,
				Message: Message{
					Event: name,
					Time:  now,
					Text:  wailstest.ErrTest.Error(),
					Data:  wailstest.ErrTest,
				},
			},
		},
		{
			name: "Messenger",
			arg:  NewMessage(name, now, "hello"),
			exp:  NewMessage(name, now, "hello"),
		},
		{
			name: "any",
			arg:  123,
			exp: Message{
				Event: name,
				Time:  now,
				Data:  123,
			},
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {

			r := require.New(t)

			msg := NewMessage(name, now, tc.arg)

			r.Equal(tc.exp, msg)
		})
	}

}

func Test_NewMessages(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	name := "events:test"
	now := wailstest.NowTime()
	args := []any{1}

	msgs := NewMessages(name, now, args...)
	r.Len(msgs, 1)

	msg, ok := msgs[0].(Message)
	r.True(ok)

	r.Equal(name, msg.Event)
	r.Equal(now, msg.Time)
	r.Equal(args[0], msg.Data)
}
