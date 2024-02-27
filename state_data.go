package wailsx

import (
	"fmt"

	"github.com/markbates/plugins"
)

type StateDataPlugin interface {
	plugins.Plugin
	StateData() (StateData, error)
}

var _ StateDataPlugin = StateData{}

type StateData struct {
	Name string // name of the data
	Data any    // data for the state
}

func (sd StateData) PluginName() string {
	return fmt.Sprintf("%T", sd)
}

func (sd StateData) StateData() (StateData, error) {
	return sd, nil
}
