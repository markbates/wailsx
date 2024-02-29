package wailstest

type es string

func (e es) Error() string {
	return string(e)
}

// ErrTest is a test error returned
// by the test error functions
const ErrTest = es("test error")
