package clipx

import (
	"context"

	"github.com/markbates/wailsx/statedata"
)

type ClipboardManager interface {
	ClipboardGetText(ctx context.Context) (string, error)
	ClipboardSetText(ctx context.Context, text string) error
}

type ClipboardManagerDataProvider interface {
	ClipboardManager
	StateData(ctx context.Context) (statedata.Data[string], error)
}
