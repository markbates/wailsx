package wailsx

import (
	"github.com/markbates/wailsx/eventx"
	"github.com/markbates/wailsx/logx"
	"github.com/markbates/wailsx/windowx"
)

type API interface {
	eventx.EventManager
	logx.Logger
	windowx.WindowManager
}
