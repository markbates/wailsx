//go:build !dev || !desktop || !production || !wails

package wailsrun

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_API(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	errs := ErrNotAvailable.Error()
	r.Equal(string(ErrNotAvailable), errs)

	type errFn func() error

	ctx := context.Background()
	tcs := []errFn{
		func() error {
			return BrowserOpenURL(ctx, "http://example.com")
		},
		func() error {
			_, err := ClipboardGetText(ctx)
			return err
		},
		func() error {
			return ClipboardSetText(ctx, "test")
		},
		func() error {
			return EventsEmit(ctx, "test")
		},
		func() error {
			return EventsOff(ctx, "test")
		},
		func() error {
			return EventsOffAll(ctx)
		},
		func() error {
			_, err := EventsOn(ctx, "test", nil)
			return err
		},
		func() error {
			_, err := EventsOnMultiple(ctx, "test", nil, 1)
			return err
		},
		func() error {
			_, err := EventsOnce(ctx, "test", nil)
			return err
		},
		func() error {
			return Hide(ctx)
		},
		func() error {
			return LogDebug(ctx, "test")
		},
		func() error {
			return LogDebugf(ctx, "test")
		},
		func() error {
			return LogError(ctx, "test")
		},
		func() error {
			return LogErrorf(ctx, "test")
		},
		func() error {
			return LogFatal(ctx, "test")
		},
		func() error {
			return LogFatalf(ctx, "test")
		},
		func() error {
			return LogInfo(ctx, "test")
		},
		func() error {
			return LogInfof(ctx, "test")
		},
		func() error {
			return LogPrint(ctx, "test")
		},
		func() error {
			return LogPrintf(ctx, "test")
		},
		func() error {
			return LogSetLogLevel(ctx, 0)
		},
		func() error {
			return LogTrace(ctx, "test")
		},
		func() error {
			return LogTracef(ctx, "test")
		},
		func() error {
			return LogWarning(ctx, "test")
		},
		func() error {
			return LogWarningf(ctx, "test")
		},
		func() error {
			return MenuSetApplicationMenu(ctx, nil)
		},
		func() error {
			return MenuUpdateApplicationMenu(ctx)
		},
		func() error {
			_, err := MessageDialog(ctx, MessageDialogOptions{})
			return err
		},
		func() error {
			_, err := OpenDirectoryDialog(ctx, OpenDialogOptions{})
			return err
		},
		func() error {
			_, err := OpenFileDialog(ctx, OpenDialogOptions{})
			return err
		},

		func() error {
			_, err := OpenMultipleFilesDialog(ctx, OpenDialogOptions{})
			return err
		},
		func() error {
			return Quit(ctx)
		},
		func() error {
			_, err := SaveFileDialog(ctx, SaveDialogOptions{})
			return err
		},
		func() error {
			return Show(ctx)
		},
		func() error {
			return WindowCenter(ctx)
		},
		func() error {
			return WindowExecJS(ctx, "test")
		},
		func() error {
			return WindowFullscreen(ctx)
		},
		func() error {
			_, _, err := WindowGetPosition(ctx)
			return err
		},
		func() error {
			_, _, err := WindowGetSize(ctx)
			return err
		},
		func() error {
			return WindowHide(ctx)
		},
		func() error {
			_, err := WindowIsFullscreen(ctx)
			return err
		},
		func() error {
			_, err := WindowIsMaximised(ctx)
			return err
		},
		func() error {
			_, err := WindowIsMinimised(ctx)
			return err
		},
		func() error {
			_, err := WindowIsNormal(ctx)
			return err
		},
		func() error {
			return WindowMaximise(ctx)
		},
		func() error {
			return WindowMinimise(ctx)
		},
		func() error {
			return WindowPrint(ctx)
		},
		func() error {
			return WindowReload(ctx)
		},
		func() error {
			return WindowReloadApp(ctx)
		},
		func() error {
			return WindowSetAlwaysOnTop(ctx, false)
		},
		func() error {
			return WindowSetBackgroundColour(ctx, 0, 0, 0, 0)
		},
		func() error {
			return WindowSetDarkTheme(ctx)
		},
		func() error {
			return WindowSetLightTheme(ctx)
		},
		func() error {
			return WindowSetMaxSize(ctx, 0, 0)
		},
		func() error {
			return WindowSetMinSize(ctx, 0, 0)
		},
		func() error {
			return WindowSetPosition(ctx, 0, 0)
		},
		func() error {
			return WindowSetSize(ctx, 0, 0)
		},
		func() error {
			return WindowSetSystemDefaultTheme(ctx)
		},
		func() error {
			return WindowSetTitle(ctx, "test")
		},
		func() error {
			return WindowShow(ctx)
		},
		func() error {
			return WindowToggleMaximise(ctx)
		},
		func() error {
			return WindowUnfullscreen(ctx)
		},
		func() error {
			return WindowUnmaximise(ctx)
		},
		func() error {
			return WindowUnminimise(ctx)
		},
		func() error {
			_, err := ScreenGetAll(ctx)
			return err
		},
	}

	for _, tc := range tcs {
		err := tc()
		r.Error(err)
		r.True(errors.Is(err, ErrNotAvailable))
	}

}
