package windowx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

var _ Maximiser = MaximiserManager{}

type MaximiserManager struct {
	WindowFullscreenFn   func(ctx context.Context) error
	WindowIsFullscreenFn func(ctx context.Context) (bool, error)
	WindowIsMaximisedFn  func(ctx context.Context) (bool, error)
	WindowIsMinimisedFn  func(ctx context.Context) (bool, error)
	WindowIsNormalFn     func(ctx context.Context) (bool, error)
	WindowMaximiseFn     func(ctx context.Context) error
	WindowMinimiseFn     func(ctx context.Context) error
	WindowUnfullscreenFn func(ctx context.Context) error
	WindowUnmaximiseFn   func(ctx context.Context) error
	WindowUnminimiseFn   func(ctx context.Context) error
}

func (mm MaximiserManager) WindowFullscreen(ctx context.Context) error {
	return safe.Run(func() error {
		if mm.WindowFullscreenFn == nil {
			return wailsrun.WindowFullscreen(ctx)
		}

		return mm.WindowFullscreenFn(ctx)
	})
}

func (mm MaximiserManager) WindowIsFullscreen(ctx context.Context) (bool, error) {
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

func (mm MaximiserManager) WindowIsMaximised(ctx context.Context) (bool, error) {
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

func (mm MaximiserManager) WindowIsMinimised(ctx context.Context) (bool, error) {
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

func (mm MaximiserManager) WindowMaximise(ctx context.Context) error {
	return safe.Run(func() error {
		if mm.WindowMaximiseFn == nil {
			return wailsrun.WindowMaximise(ctx)
		}

		return mm.WindowMaximiseFn(ctx)
	})
}

func (mm MaximiserManager) WindowMinimise(ctx context.Context) error {
	return safe.Run(func() error {
		if mm.WindowMinimiseFn == nil {
			return wailsrun.WindowMinimise(ctx)
		}

		return mm.WindowMinimiseFn(ctx)
	})
}

func (mm MaximiserManager) WindowUnfullscreen(ctx context.Context) error {
	return safe.Run(func() error {
		if mm.WindowUnfullscreenFn == nil {
			return wailsrun.WindowUnfullscreen(ctx)
		}

		return mm.WindowUnfullscreenFn(ctx)
	})
}

func (mm MaximiserManager) WindowUnmaximise(ctx context.Context) error {
	return safe.Run(func() error {
		if mm.WindowUnmaximiseFn == nil {
			return wailsrun.WindowUnmaximise(ctx)
		}

		return mm.WindowUnmaximiseFn(ctx)
	})
}

func (mm MaximiserManager) WindowUnminimise(ctx context.Context) error {
	return safe.Run(func() error {
		if mm.WindowUnminimiseFn == nil {
			return wailsrun.WindowUnminimise(ctx)
		}

		return mm.WindowUnminimiseFn(ctx)
	})
}

func (mm MaximiserManager) WindowIsNormal(ctx context.Context) (bool, error) {
	var b bool

	err := safe.Run(func() error {
		if mm.WindowIsNormalFn == nil {
			mm.WindowIsNormalFn = wailsrun.WindowIsNormal
		}

		var err error
		b, err = mm.WindowIsNormalFn(ctx)
		return err
	})

	return b, err
}
