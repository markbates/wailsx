package windowx

import "context"

type ReloadManager interface {
	WindowReload(ctx context.Context) error
	WindowReloadApp(ctx context.Context) error
}
