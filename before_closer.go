package wailsx

import "context"

type BeforeCloser interface {
	BeforeClose(ctx context.Context) error
}
