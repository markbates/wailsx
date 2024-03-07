package wailsrun

type Screen struct {
	IsCurrent bool `json:"isCurrent"`
	IsPrimary bool `json:"isPrimary"`

	// Size is the size of the screen in logical pixel space, used when setting sizes in Wails
	Size ScreenSize `json:"size"`

	// PhysicalSize is the physical size of the screen in pixels
	PhysicalSize ScreenSize `json:"physicalSize"`
}

type ScreenSize struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}
