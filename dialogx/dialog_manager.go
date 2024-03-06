package dialogx

import (
	"context"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/wailsrun"
)

var _ Dialoger = DialogManager{}

type DialogManager struct {
	MessageDialogFn           func(ctx context.Context, opts MessageDialogOptions) (string, error)
	OpenDirectoryDialogFn     func(ctx context.Context, opts OpenDialogOptions) (string, error)
	OpenFileDialogFn          func(ctx context.Context, opts OpenDialogOptions) (string, error)
	OpenMultipleFilesDialogFn func(ctx context.Context, opts OpenDialogOptions) ([]string, error)
	SaveFileDialogFn          func(ctx context.Context, opts SaveDialogOptions) (string, error)
}

func (dm DialogManager) MessageDialog(ctx context.Context, opts MessageDialogOptions) (fp string, err error) {
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

func (dm DialogManager) OpenDirectoryDialog(ctx context.Context, opts OpenDialogOptions) (fp string, err error) {
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

func (dm DialogManager) OpenFileDialog(ctx context.Context, opts OpenDialogOptions) (fp string, err error) {
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

func (dm DialogManager) OpenMultipleFilesDialog(ctx context.Context, opts OpenDialogOptions) (fp []string, err error) {
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

func (dm DialogManager) SaveFileDialog(ctx context.Context, opts SaveDialogOptions) (fp string, err error) {
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
