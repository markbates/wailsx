package dialogx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

var _ DialogManager = Manager{}

func NopManager() Manager {
	return Manager{
		MessageDialogFn:           func(ctx context.Context, opts MessageDialogOptions) (string, error) { return "", nil },
		OpenDirectoryDialogFn:     func(ctx context.Context, opts OpenDialogOptions) (string, error) { return "", nil },
		OpenFileDialogFn:          func(ctx context.Context, opts OpenDialogOptions) (string, error) { return "", nil },
		OpenMultipleFilesDialogFn: func(ctx context.Context, opts OpenDialogOptions) ([]string, error) { return nil, nil },
		SaveFileDialogFn:          func(ctx context.Context, opts SaveDialogOptions) (string, error) { return "", nil },
	}
}

type Manager struct {
	MessageDialogFn           func(ctx context.Context, opts MessageDialogOptions) (string, error) `json:"-"`
	OpenDirectoryDialogFn     func(ctx context.Context, opts OpenDialogOptions) (string, error)    `json:"-"`
	OpenFileDialogFn          func(ctx context.Context, opts OpenDialogOptions) (string, error)    `json:"-"`
	OpenMultipleFilesDialogFn func(ctx context.Context, opts OpenDialogOptions) ([]string, error)  `json:"-"`
	SaveFileDialogFn          func(ctx context.Context, opts SaveDialogOptions) (string, error)    `json:"-"`
}

func (dm Manager) MessageDialog(ctx context.Context, opts MessageDialogOptions) (fp string, err error) {
	err = safe.Run(func() error {
		fn := dm.MessageDialogFn
		if fn == nil {
			fn = wailsrun.MessageDialog
		}

		fp, err = fn(ctx, opts)
		return err
	})

	if err != nil {
		return "", err
	}

	return fp, err
}

func (dm Manager) OpenDirectoryDialog(ctx context.Context, opts OpenDialogOptions) (fp string, err error) {
	err = safe.Run(func() error {
		fn := dm.OpenDirectoryDialogFn
		if fn == nil {
			fn = wailsrun.OpenDirectoryDialog
		}

		fp, err = fn(ctx, opts)
		return err
	})

	if err != nil {
		return "", err
	}

	return fp, err
}

func (dm Manager) OpenFileDialog(ctx context.Context, opts OpenDialogOptions) (fp string, err error) {
	err = safe.Run(func() error {
		fn := dm.OpenFileDialogFn
		if fn == nil {
			fn = wailsrun.OpenFileDialog
		}

		fp, err = fn(ctx, opts)
		return err
	})

	if err != nil {
		return "", err
	}

	return fp, err
}

func (dm Manager) OpenMultipleFilesDialog(ctx context.Context, opts OpenDialogOptions) (fp []string, err error) {
	err = safe.Run(func() error {
		fn := dm.OpenMultipleFilesDialogFn
		if fn == nil {
			fn = wailsrun.OpenMultipleFilesDialog
		}

		fp, err = fn(ctx, opts)
		return err
	})

	if err != nil {
		return nil, err
	}

	return fp, err
}

func (dm Manager) SaveFileDialog(ctx context.Context, opts SaveDialogOptions) (fp string, err error) {
	err = safe.Run(func() error {
		fn := dm.SaveFileDialogFn
		if fn == nil {
			fn = wailsrun.SaveFileDialog
		}

		fp, err = fn(ctx, opts)
		return err
	})

	if err != nil {
		return "", err
	}

	return fp, err
}
