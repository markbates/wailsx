package windowx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

var _ Maximizer = MaximizerManager{}

type MaximizerManager struct {
	WindowFullscreenFn   func(ctx context.Context) error
	WindowIsFullscreenFn func(ctx context.Context) (bool, error)
	WindowIsMaximisedFn  func(ctx context.Context) (bool, error)
	WindowIsMinimisedFn  func(ctx context.Context) (bool, error)
	WindowMaximiseFn     func(ctx context.Context) error
	WindowMinimiseFn     func(ctx context.Context) error
	WindowUnfullscreenFn func(ctx context.Context) error
	WindowUnmaximiseFn   func(ctx context.Context) error
	WindowUnminimiseFn   func(ctx context.Context) error
}

func (mm MaximizerManager) WindowFullscreen(ctx context.Context) error {
	return safe.Run(func() error {
		if mm.WindowFullscreenFn == nil {
			return wailsrun.WindowFullscreen(ctx)
		}
		return mm.WindowFullscreenFn(ctx)
	})
}

func (mm MaximizerManager) WindowIsFullscreen(ctx context.Context) (bool, error) {
	var b bool

	err := safe.Run(func() error {
		if mm.WindowIsFullscreenFn == nil {
			var err error
			b, err = wailsrun.WindowIsFullscreen(ctx)
			return err
		}
		var err error
		b, err = mm.WindowIsFullscreenFn(ctx)
		return err
	})

	return b, err
}

func (mm MaximizerManager) WindowIsMaximised(ctx context.Context) (bool, error) {
	var b bool

	err := safe.Run(func() error {
		if mm.WindowIsMaximisedFn == nil {
			var err error
			b, err = wailsrun.WindowIsMaximised(ctx)
			return err
		}
		var err error
		b, err = mm.WindowIsMaximisedFn(ctx)
		return err
	})

	return b, err
}

func (mm MaximizerManager) WindowIsMinimised(ctx context.Context) (bool, error) {
	var b bool

	err := safe.Run(func() error {
		if mm.WindowIsMinimisedFn == nil {
			var err error
			b, err = wailsrun.WindowIsMinimised(ctx)
			return err
		}
		var err error
		b, err = mm.WindowIsMinimisedFn(ctx)
		return err
	})

	return b, err
}

func (mm MaximizerManager) WindowMaximise(ctx context.Context) error {
	return safe.Run(func() error {
		if mm.WindowMaximiseFn == nil {
			return wailsrun.WindowMaximise(ctx)
		}
		return mm.WindowMaximiseFn(ctx)
	})
}

func (mm MaximizerManager) WindowMinimise(ctx context.Context) error {
	return safe.Run(func() error {
		if mm.WindowMinimiseFn == nil {
			return wailsrun.WindowMinimise(ctx)
		}
		return mm.WindowMinimiseFn(ctx)
	})
}

func (mm MaximizerManager) WindowUnfullscreen(ctx context.Context) error {
	return safe.Run(func() error {
		if mm.WindowUnfullscreenFn == nil {
			return wailsrun.WindowUnfullscreen(ctx)
		}
		return mm.WindowUnfullscreenFn(ctx)
	})
}

func (mm MaximizerManager) WindowUnmaximise(ctx context.Context) error {
	return safe.Run(func() error {
		if mm.WindowUnmaximiseFn == nil {
			return wailsrun.WindowUnmaximise(ctx)
		}
		return mm.WindowUnmaximiseFn(ctx)
	})
}

func (mm MaximizerManager) WindowUnminimise(ctx context.Context) error {
	return safe.Run(func() error {
		if mm.WindowUnminimiseFn == nil {
			return wailsrun.WindowUnminimise(ctx)
		}
		return mm.WindowUnminimiseFn(ctx)
	})
}

func (mm MaximizerManager) WindowIsNormal(ctx context.Context) (bool, error) {
	var b bool

	err := safe.Run(func() error {
		if mm.WindowIsMinimisedFn == nil {
			var err error
			b, err = wailsrun.WindowIsNormal(ctx)
			return err
		}
		var err error
		b, err = mm.WindowIsMinimisedFn(ctx)
		return err
	})

	return b, err
}
