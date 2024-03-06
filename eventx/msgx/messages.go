package msgx

type Messengers []Messenger

func (msgs Messengers) Any() []any {
	args := make([]any, len(msgs))
	for i, m := range msgs {
		args[i] = m
	}
	return args
}
