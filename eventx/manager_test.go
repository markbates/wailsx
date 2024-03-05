package eventx

import (
	"context"
	"strings"
	"testing"

	"github.com/markbates/wailsx/wailstest"
	"github.com/stretchr/testify/require"
)

func Test_Manager_Now(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	em := Manager{}
	r.NotZero(em.Now())
	r.NotEqual(wailstest.OldTime(), em.Now())

	em.NowFn = wailstest.OldTime
	r.Equal(wailstest.OldTime(), em.Now())
}

func Test_Manager_StateData_JSON(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ctx := context.Background()

	var em *Manager
	_, err := em.StateData(ctx)
	r.Error(err)

	em = &Manager{
		data: EventsData{
			Emitted: map[string][]any{
				"test:event": {"my arg"},
			},
			Caught: map[string][]any{
				"test:event": {"my arg"},
			},
			Callbacks: map[string]*CallbackCounter{
				"test:event": {
					MaxCalls: 5,
					Called:   3,
				},
			},
		},
	}

	sd, err := em.StateData(ctx)
	r.NoError(err)
	r.Equal(EventManagerStateDataName, sd.Name)

	data := sd.Data
	r.Len(data.Emitted, 1)
	r.Len(data.Caught, 1)

	b, err := em.MarshalJSON()
	r.NoError(err)

	act := string(b)
	act = strings.TrimSpace(act)

	exp := `
{
  "callbacks": {
    "test:event": {
      "called": 3,
      "max_calls": 5,
      "off": false
    }
  },
  "emitted": {
    "test:event": [
      "my arg"
    ]
  },
  "caught": {
    "test:event": [
      "my arg"
    ]
  }
}`

	exp = strings.TrimSpace(exp)

	r.Equal(exp, act)
}
