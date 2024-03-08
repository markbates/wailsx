package eventx

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/markbates/wailsx/eventx/msgx"
	"github.com/markbates/wailsx/wailsrun"
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

	const name = "test:event"

	msg := msgx.Message{
		Event: name,
		Data:  1,
		Text:  "my arg",
		Time:  wailstest.NowTime(),
	}

	event := Event{
		Name:      name,
		Data:      []msgx.Messenger{msg},
		EmittedAt: wailstest.NowTime(),
	}

	em = &Manager{
		data: EventsData{
			Emitted: map[string][]Event{
				"test:event": {
					event,
				},
			},
			Caught: map[string][]Event{
				"test:event": {
					event,
				},
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

	// fmt.Println(act)

	// f, err := os.Create("testdata/state.json")
	// r.NoError(err)
	// enc := json.NewEncoder(f)
	// r.NoError(enc.Encode(data))
	// r.NoError(f.Close())

	b, err = os.ReadFile("testdata/state.json")
	r.NoError(err)

	exp := string(b)

	exp = strings.TrimSpace(exp)

	r.Equal(exp, act)
}

func Test_Manager_DisableStateData(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	ctx := context.Background()

	const name = "test:event"

	em := NopManager()
	em.DisableStateData = true

	for i := 0; i < 5; i++ {
		err := em.EventsEmit(ctx, name, i)
		r.NoError(err)
	}

	sd, err := em.StateData(ctx)
	r.NoError(err)
	r.Nil(sd.Data)
}

func Test_Nil_Manager(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	var em *Manager

	ctx := context.Background()

	err := em.EventsEmit(ctx, "test:event")
	r.Error(err)

	exp := wailsrun.ErrNotAvailable("EventsEmit")
	r.Equal(exp, err)

	_, err = em.EventsOn(ctx, "test:event", func(args ...any) error {
		return nil
	})

	r.Error(err)

	exp = wailsrun.ErrNotAvailable("EventsOn")
	r.Equal(exp, err)

	_, err = em.EventsOnMultiple(ctx, "test:event", func(args ...any) error {
		return nil
	}, 1)

	r.Error(err)

	exp = wailsrun.ErrNotAvailable("EventsOnMultiple")
	r.Equal(exp, err)

	err = em.EventsOff(ctx, "test:event")

	r.Error(err)

	exp = wailsrun.ErrNotAvailable("EventsOff")
	r.Equal(exp, err)

	err = em.EventsOffAll(ctx)

	r.Error(err)

	exp = wailsrun.ErrNotAvailable("EventsOffAll")
	r.Equal(exp, err)

	_, err = em.EventsOnce(ctx, "test:event", func(args ...any) error {
		return nil
	})

	r.Error(err)

	exp = wailsrun.ErrNotAvailable("EventsOnce")
	r.Equal(exp, err)
}
