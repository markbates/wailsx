package wailstest

import "fmt"

type ErrorLogger struct {
	Messages []string
}

func (el *ErrorLogger) LogError(message string) {
	if el == nil {
		return
	}

	el.Messages = append(el.Messages, message)
}

func (el *ErrorLogger) LogErrorf(format string, args ...any) {
	if el == nil {
		return
	}

	el.Messages = append(el.Messages, fmt.Sprintf(format, args...))
}
