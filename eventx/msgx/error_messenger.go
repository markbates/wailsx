package msgx

import "time"

type ErrorMessenger interface {
	MsgError() error
	MsgEvent() string
	MsgText() string
	MsgTime() time.Time
	MsgData() any
}
