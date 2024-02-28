package wailstest

type es string

func (e es) Error() string {
	return string(e)
}

// ERR is a test error returned
// by the test error functions
const ERR = es("test error")
