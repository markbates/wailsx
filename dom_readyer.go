package wailsx

import "context"

type DomReadyer interface {
	DomReady(ctx context.Context) error
}
