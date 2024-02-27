package wailstest

import "context"

func PositionGet(a int, b int) func(ctx context.Context) (int, int, error) {
	return func(ctx context.Context) (int, int, error) {
		return a, b, nil
	}
}
