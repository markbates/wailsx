package dialogx

import "context"

type DialogManager interface {
	MessageDialog(ctx context.Context, opts MessageDialogOptions) (string, error)
	OpenDirectoryDialog(ctx context.Context, opts OpenDialogOptions) (string, error)
	OpenFileDialog(ctx context.Context, opts OpenDialogOptions) (string, error)
	OpenMultipleFilesDialog(ctx context.Context, opts OpenDialogOptions) ([]string, error)
	SaveFileDialog(ctx context.Context, opts SaveDialogOptions) (string, error)
}
