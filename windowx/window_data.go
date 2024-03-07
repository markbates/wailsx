package windowx

type WindowData struct {
	*MaximiserData `json:"maximiser,omitempty"`
	*PositionData  `json:"positioner,omitempty"`
	*ThemeData     `json:"themer,omitempty"`
}
