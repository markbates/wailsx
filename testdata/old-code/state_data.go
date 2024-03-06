package wailsx

import (
	"fmt"
)

var _ StateDataProvider = StateData{}

type StateData struct {
	Name string `json:"name,omitempty"` // name of the data
	Data any    `json:"data,omitempty"` // data for the state
}

func (sd StateData) PluginName() string {
	return fmt.Sprintf("%T", sd)
}

func (sd StateData) StateData() (StateData, error) {
	return sd, nil
}
