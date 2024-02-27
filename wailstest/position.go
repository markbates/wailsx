package wailstest

import "context"

func PositionGet(a int, b int) func(ctx context.Context) (int, int) {
	return func(ctx context.Context) (int, int) {
		return a, b
	}
}
