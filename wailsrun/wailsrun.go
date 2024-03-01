package wailsrun

import (
	wailsrun "github.com/wailsapp/wails/v2/pkg/runtime"
)

const (
	ErrNotAvailable = es("wails api calls are not available in this environment")
)

type es string

func (e es) Error() string {
	return string(e)
}

const (
	InfoDialog     = wailsrun.InfoDialog
	WarningDialog  = wailsrun.WarningDialog
	ErrorDialog    = wailsrun.ErrorDialog
	QuestionDialog = wailsrun.QuestionDialog
)

type DialogType = wailsrun.DialogType
type EnvironmentInfo = wailsrun.EnvironmentInfo
type FileFilter = wailsrun.FileFilter
type MessageDialogOptions = wailsrun.MessageDialogOptions
type OpenDialogOptions = wailsrun.OpenDialogOptions
type SaveDialogOptions = wailsrun.SaveDialogOptions
type Screen = wailsrun.Screen
