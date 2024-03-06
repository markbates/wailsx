package eventx

import (
	"time"

	"github.com/markbates/wailsx/eventx/msgx"
)

type Event struct {
	Name      string          `json:"name,omitempty"`
	Data      msgx.Messengers `json:"data,omitempty"`
	EmittedAt time.Time       `json:"emitted_at,omitempty"`
}

func NewEvent(name string, now time.Time, args ...any) (Event, error) {
	if now.IsZero() {
		now = time.Now()
	}

	return Event{
		Name:      name,
		Data:      msgx.NewMessages(name, time.Now(), args...),
		EmittedAt: now,
	}, nil
}
