package windowx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

var _ Toggler = Toggle{}

func NewNOOPToggle() Toggle {
	return Toggle{
		HideFn:       func(ctx context.Context) error { return nil },
		ShowFn:       func(ctx context.Context) error { return nil },
		WindowHideFn: func(ctx context.Context) error { return nil },
	}
}

type Toggle struct {
	HideFn       func(ctx context.Context) error
	ShowFn       func(ctx context.Context) error
	WindowHideFn func(ctx context.Context) error
}

func (t Toggle) Hide(ctx context.Context) error {
	return safe.Run(func() error {
		if t.HideFn == nil {
			return wailsrun.Hide(ctx)
		}
		return t.HideFn(ctx)
	})
}

func (t Toggle) Show(ctx context.Context) error {
	return safe.Run(func() error {
		if t.ShowFn == nil {
			return wailsrun.Show(ctx)
		}
		return t.ShowFn(ctx)
	})
}

func (t Toggle) WindowHide(ctx context.Context) error {
	return safe.Run(func() error {
		if t.WindowHideFn == nil {
			return wailsrun.WindowHide(ctx)
		}
		return t.WindowHideFn(ctx)
	})
}
