package wailsrun

import (
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/logger"
	wailsrun "github.com/wailsapp/wails/v2/pkg/runtime"
)

type ErrNotAvailable string

func (e ErrNotAvailable) Error() string {
	return fmt.Sprintf("wails api calls are not available in this environment: %q", string(e))
}

// const (
// 	ErrNotAvailable = es("wails api calls are not available in this environment")
// )

// type es string

// func (e es) Error() string {
// 	return string(e)
// }

const (
	InfoDialog     = wailsrun.InfoDialog
	WarningDialog  = wailsrun.WarningDialog
	ErrorDialog    = wailsrun.ErrorDialog
	QuestionDialog = wailsrun.QuestionDialog

	DEBUG   LogLevel = logger.DEBUG
	ERROR   LogLevel = logger.ERROR
	INFO    LogLevel = logger.INFO
	TRACE   LogLevel = logger.TRACE
	WARNING LogLevel = logger.WARNING

	// extra log levels
	FATAL LogLevel = LogLevel(42)
	PRINT LogLevel = LogLevel(43)
)

type DialogType = wailsrun.DialogType
type EnvironmentInfo = wailsrun.EnvironmentInfo
type FileFilter = wailsrun.FileFilter
type MessageDialogOptions = wailsrun.MessageDialogOptions
type OpenDialogOptions = wailsrun.OpenDialogOptions
type SaveDialogOptions = wailsrun.SaveDialogOptions
type LogLevel = logger.LogLevel
