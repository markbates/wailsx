package msgx

import "time"

func NewMessages(event string, now time.Time, args ...any) Messengers {
	var msgs = make([]Messenger, 0, len(args))
	for _, a := range args {
		msgs = append(msgs, NewMessage(event, now, a))
	}
	return msgs
}

func NewMessage(event string, now time.Time, arg any) Messenger {
	if now.IsZero() {
		now = time.Now()
	}

	msg := Message{
		Event: event,
		Time:  now,
	}

	switch t := arg.(type) {
	case Messenger:
		return t
	case error:
		msg.Text = t.Error()
		msg.Data = t
		return ErrorMessage{
			Err:     t,
			Message: msg,
		}
	case string:
		msg.Text = t
		msg.Data = t
	default:
		msg.Data = t
	}

	return msg
}
