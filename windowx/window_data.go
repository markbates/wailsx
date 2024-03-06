package windowx

type WindowData struct {
	*MaximiserData `json:"maximiser_data,omitempty"`
	*PositionData  `json:"positioner_data,omitempty"`
	*ThemeData     `json:"themer_data,omitempty"`
}
