package wailsrun

type CallbackFn func(data ...any) error

type CancelFn func() error
