package wailsx

import "context"

type Shutdowner interface {
	Shutdown(ctx context.Context) error
}
