package windowx

import "context"

type Toggler interface {
	Hide(ctx context.Context) error
	Show(ctx context.Context) error
	WindowHide(ctx context.Context) error
}
