package windowx

import (
	"context"
	"fmt"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

var _ MaximiseManagerDataProvider = &Maximiser{}

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
	WindowFullscreenFn     func(ctx context.Context) error
	WindowIsFullscreenFn   func(ctx context.Context) (bool, error)
	WindowIsMaximisedFn    func(ctx context.Context) (bool, error)
	WindowIsMinimisedFn    func(ctx context.Context) (bool, error)
	WindowIsNormalFn       func(ctx context.Context) (bool, error)
	WindowMaximiseFn       func(ctx context.Context) error
	WindowMinimiseFn       func(ctx context.Context) error
	WindowToggleMaximiseFn func(ctx context.Context) error
	WindowUnfullscreenFn   func(ctx context.Context) error
	WindowUnmaximiseFn     func(ctx context.Context) error
	WindowUnminimiseFn     func(ctx context.Context) error

	data MaximiserData
}

func (mm *Maximiser) WindowFullscreen(ctx context.Context) error {
	if mm == nil {
		return fmt.Errorf("maximiser manager is nil")
	}

	return safe.Run(func() error {
		if mm.WindowFullscreenFn == nil {
			return wailsrun.WindowFullscreen(ctx)
		}

		err := mm.WindowFullscreenFn(ctx)
		if err != nil {
			return err
		}

		return mm.data.SetFullscreen()
	})
}

func (mm *Maximiser) WindowIsFullscreen(ctx context.Context) (bool, error) {
	if mm == nil {
		return false, fmt.Errorf("maximiser manager is nil")
	}

	var b bool

	err := safe.Run(func() error {
		if mm.WindowIsFullscreenFn == nil {
			mm.WindowIsFullscreenFn = wailsrun.WindowIsFullscreen
		}

		var err error
		b, err = mm.WindowIsFullscreenFn(ctx)
		return err
	})

	return b, err
}

func (mm *Maximiser) WindowIsMaximised(ctx context.Context) (bool, error) {
	if mm == nil {
		return false, fmt.Errorf("maximiser manager is nil")
	}

	var b bool

	err := safe.Run(func() error {
		if mm.WindowIsMaximisedFn == nil {
			mm.WindowIsMaximisedFn = wailsrun.WindowIsMaximised
		}

		var err error
		b, err = mm.WindowIsMaximisedFn(ctx)
		return err
	})

	return b, err
}

func (mm *Maximiser) WindowIsMinimised(ctx context.Context) (bool, error) {
	if mm == nil {
		return false, fmt.Errorf("maximiser manager is nil")
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
		return fmt.Errorf("maximiser manager is nil")
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
		return fmt.Errorf("maximiser manager is nil")
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
		return fmt.Errorf("maximiser manager is nil")
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
		return fmt.Errorf("maximiser manager is nil")
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
		return fmt.Errorf("maximiser manager is nil")
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
		return false, fmt.Errorf("maximiser manager is nil")
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
		return fmt.Errorf("maximiser manager is nil")
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
