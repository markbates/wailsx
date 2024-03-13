package msgx

type ErrorMessenger interface {
	Messenger
	MsgError() error
}
