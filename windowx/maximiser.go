package windowx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

var _ MaximiseManagerDataProvider = &Maximiser{}
var _ RestorableMaximiseManager = &Maximiser{}

func NopMaximiser() *Maximiser {
	return &Maximiser{
		WindowFullscreenFn:     func(ctx context.Context) error { return nil },
		WindowIsFullscreenFn:   func(ctx context.Context) (bool, error) { return false, nil },
		WindowIsMaximisedFn:    func(ctx context.Context) (bool, error) { return false, nil },
		WindowIsMinimisedFn:    func(ctx context.Context) (bool, error) { return false, nil },
		WindowIsNormalFn:       func(ctx context.Context) (bool, error) { return false, nil },
		WindowMaximiseFn:       func(ctx context.Context) error { return nil },
		WindowMinimiseFn:       func(ctx context.Context) error { return nil },
		WindowToggleMaximiseFn: func(ctx context.Context) error { return nil },
		WindowUnfullscreenFn:   func(ctx context.Context) error { return nil },
		WindowUnmaximiseFn:     func(ctx context.Context) error { return nil },
		WindowUnminimiseFn:     func(ctx context.Context) error { return nil },
	}
}

type Maximiser struct {
	WindowFullscreenFn     func(ctx context.Context) error         `json:"-"`
	WindowIsFullscreenFn   func(ctx context.Context) (bool, error) `json:"-"`
	WindowIsMaximisedFn    func(ctx context.Context) (bool, error) `json:"-"`
	WindowIsMinimisedFn    func(ctx context.Context) (bool, error) `json:"-"`
	WindowIsNormalFn       func(ctx context.Context) (bool, error) `json:"-"`
	WindowMaximiseFn       func(ctx context.Context) error         `json:"-"`
	WindowMinimiseFn       func(ctx context.Context) error         `json:"-"`
	WindowToggleMaximiseFn func(ctx context.Context) error         `json:"-"`
	WindowUnfullscreenFn   func(ctx context.Context) error         `json:"-"`
	WindowUnmaximiseFn     func(ctx context.Context) error         `json:"-"`
	WindowUnminimiseFn     func(ctx context.Context) error         `json:"-"`

	data MaximiserData
}

func (mm *Maximiser) WindowFullscreen(ctx context.Context) error {
	if mm == nil {
		return wailsrun.WindowFullscreen(ctx)
	}

	return safe.Run(func() error {
		fn := mm.WindowFullscreenFn
		if fn == nil {
			fn = wailsrun.WindowFullscreen
		}

		if err := fn(ctx); err != nil {
			return err
		}

		return mm.data.SetFullscreen()
	})
}

func (mm *Maximiser) WindowIsFullscreen(ctx context.Context) (bool, error) {
	if mm == nil {
		return wailsrun.WindowIsFullscreen(ctx)
	}

	var b bool

	err := safe.Run(func() error {
		fn := mm.WindowIsFullscreenFn
		if fn == nil {
			fn = wailsrun.WindowIsFullscreen
		}

		var err error
		b, err = fn(ctx)
		return err
	})

	return b, err
}

func (mm *Maximiser) WindowIsMaximised(ctx context.Context) (bool, error) {
	if mm == nil {
		return wailsrun.WindowIsMaximised(ctx)
	}

	var b bool

	err := safe.Run(func() error {
		fn := mm.WindowIsMaximisedFn
		if fn == nil {
			fn = wailsrun.WindowIsMaximised
		}

		var err error
		b, err = fn(ctx)
		return err
	})

	if err != nil {
		return false, err
	}

	return b, nil
}

func (mm *Maximiser) WindowIsMinimised(ctx context.Context) (bool, error) {
	if mm == nil {
		return wailsrun.WindowIsMinimised(ctx)
	}

	var b bool

	err := safe.Run(func() error {
		if mm.WindowIsMinimisedFn == nil {
			mm.WindowIsMinimisedFn = wailsrun.WindowIsMinimised
		}

		var err error
		b, err = mm.WindowIsMinimisedFn(ctx)
		return err
	})

	return b, err
}

func (mm *Maximiser) WindowMaximise(ctx context.Context) error {
	if mm == nil {
		return wailsrun.WindowMaximise(ctx)
	}

	return safe.Run(func() error {
		if mm.WindowMaximiseFn == nil {
			return wailsrun.WindowMaximise(ctx)
		}

		err := mm.WindowMaximiseFn(ctx)
		if err != nil {
			return err
		}

		return mm.data.SetMaximised()
	})
}

func (mm *Maximiser) WindowMinimise(ctx context.Context) error {
	if mm == nil {
		return wailsrun.WindowMinimise(ctx)
	}

	return safe.Run(func() error {
		if mm.WindowMinimiseFn == nil {
			return wailsrun.WindowMinimise(ctx)
		}

		err := mm.WindowMinimiseFn(ctx)
		if err != nil {
			return err
		}

		return mm.data.SetMinimised()
	})
}

func (mm *Maximiser) WindowUnfullscreen(ctx context.Context) error {
	if mm == nil {
		return wailsrun.WindowUnfullscreen(ctx)
	}

	return safe.Run(func() error {
		if mm.WindowUnfullscreenFn == nil {
			return wailsrun.WindowUnfullscreen(ctx)
		}

		err := mm.WindowUnfullscreenFn(ctx)
		if err != nil {
			return err
		}

		return mm.data.SetUnfullscreen()
	})
}

func (mm *Maximiser) WindowUnmaximise(ctx context.Context) error {
	if mm == nil {
		return wailsrun.WindowUnmaximise(ctx)
	}

	return safe.Run(func() error {
		if mm.WindowUnmaximiseFn == nil {
			return wailsrun.WindowUnmaximise(ctx)
		}

		err := mm.WindowUnmaximiseFn(ctx)
		if err != nil {
			return err
		}

		return mm.data.SetUnmaximised()
	})
}

func (mm *Maximiser) WindowUnminimise(ctx context.Context) error {
	if mm == nil {
		return wailsrun.WindowUnminimise(ctx)
	}

	return safe.Run(func() error {
		if mm.WindowUnminimiseFn == nil {
			return wailsrun.WindowUnminimise(ctx)
		}

		err := mm.WindowUnminimiseFn(ctx)
		if err != nil {
			return err
		}

		return mm.data.SetUnminimised()
	})
}

func (mm *Maximiser) WindowIsNormal(ctx context.Context) (bool, error) {
	if mm == nil {
		return wailsrun.WindowIsNormal(ctx)
	}

	var b bool

	err := safe.Run(func() error {
		if mm.WindowIsNormalFn == nil {
			mm.WindowIsNormalFn = wailsrun.WindowIsNormal
		}

		var err error
		b, err = mm.WindowIsNormalFn(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	return b, err
}

func (mm *Maximiser) WindowToggleMaximise(ctx context.Context) error {
	if mm == nil {
		return wailsrun.WindowToggleMaximise(ctx)
	}

	return safe.Run(func() error {
		fn := mm.WindowToggleMaximiseFn

		if fn == nil {
			fn = wailsrun.WindowToggleMaximise
		}

		if err := fn(ctx); err != nil {
			return err
		}

		return mm.data.ToggleMaximised()
	})
}

func (mm *Maximiser) MaximiseRestore(ctx context.Context, data *MaximiserData) error {
	return nil
}
