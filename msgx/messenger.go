package msgx

import "time"

type Messenger interface {
	MsgEvent() string
	MsgText() string
	MsgTime() time.Time
	MsgData() any
}
