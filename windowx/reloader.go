package windowx

import "context"

type Reloader interface {
	WindowReload(ctx context.Context) error
	WindowReloadApp(ctx context.Context) error
}
