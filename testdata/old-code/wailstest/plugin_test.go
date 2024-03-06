package wailstest_test

import (
	"testing"

	"github.com/markbates/wailsx"
	. "github.com/markbates/wailsx/wailstest"
)

func Test_Plugin_Implementations(t *testing.T) {
	t.Parallel()

	var _ wailsx.Saver = &SaverPlugin{}
	var _ wailsx.Shutdowner = &ShutdownerPlugin{}
	var _ wailsx.Startuper = &StartuperPlugin{}
}
