package wailstest

import "context"

// WindowGetPosition returns a function that returns the given values
// if a && b == -1, the returned function will error out
func WindowGetPosition(a int, b int) func(ctx context.Context) (int, int, error) {
	return func(ctx context.Context) (int, int, error) {
		if a == -1 && b == -1 {
			return 0, 0, ERR
		}

		return a, b, nil
	}
}

// WindowSetPosition returns a function that returns the given values
// if a && b == -1, the returned function will error out
func WindowGetSize(a int, b int) func(ctx context.Context) (int, int, error) {
	return func(ctx context.Context) (int, int, error) {
		if a == -1 && b == -1 {
			return 0, 0, ERR
		}

		return a, b, nil
	}
}
