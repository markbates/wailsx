package wailsx

import "context"

type Startuper interface {
	Startup(ctx context.Context) error
}
