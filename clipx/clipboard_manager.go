package clipx

import (
	"context"
)

type ClipboardManager interface {
	ClipboardGetText(ctx context.Context) (string, error)
	ClipboardSetText(ctx context.Context, text string) error
}

type ClipboardManagerDataProvider interface {
	ClipboardManager
	StateData(ctx context.Context) (string, error)
}

type RestorableClipboardManager interface {
	ClipboardManager
	RestoreClipboard(ctx context.Context, data string) error
}
