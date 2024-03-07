package wailsx

import (
	"github.com/markbates/wailsx/eventx"
	"github.com/markbates/wailsx/windowx"
)

type APIData struct {
	Events *eventx.EventsData  `json:"events,omitempty"`
	Window *windowx.WindowData `json:"window,omitempty"`
}
