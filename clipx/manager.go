package clipx

import (
	"context"
	"sync"

	"github.com/markbates/safe"
	"github.com/markbates/wailsx/statedata"
	"github.com/markbates/wailsx/wailsrun"
)

var _ ClipboardManagerDataProvider = &Manager{}

func NopManager() *Manager {
	m := &Manager{}
	m.ClipboardGetTextFn = func(ctx context.Context) (string, error) {
		m.mu.RLock()
		defer m.mu.RUnlock()
		return m.Content, nil
	}

	m.ClipboardSetTextFn = func(ctx context.Context, text string) error {
		m.mu.Lock()
		m.Content = text
		m.mu.Unlock()
		return nil
	}

	return m
}

type Manager struct {
	Content string `json:"content,omitempty"`

	ClipboardGetTextFn func(ctx context.Context) (string, error)    `json:"-"`
	ClipboardSetTextFn func(ctx context.Context, text string) error `json:"-"`

	mu sync.RWMutex
}

func (m *Manager) ClipboardGetText(ctx context.Context) (s string, err error) {
	if m == nil {
		return wailsrun.ClipboardGetText(ctx)
	}

	err = safe.Run(func() error {
		fn := m.ClipboardGetTextFn
		if fn == nil {
			fn = wailsrun.ClipboardGetText
		}

		s, err = fn(ctx)
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return s, nil
}

func (m *Manager) ClipboardSetText(ctx context.Context, text string) error {
	if m == nil {
		return wailsrun.ClipboardSetText(ctx, text)
	}

	return safe.Run(func() error {
		fn := m.ClipboardSetTextFn
		if fn == nil {
			fn = wailsrun.ClipboardSetText
		}

		if err := fn(ctx, text); err != nil {
			return err
		}

		m.mu.Lock()
		m.Content = text
		m.mu.Unlock()
		return nil
	})
}

func (m *Manager) StateData(ctx context.Context) (statedata.Data[string], error) {
	sd := statedata.Data[string]{
		Name: ClipboardManagerStateDataProviderName,
	}

	if m == nil {
		return sd, nil
	}

	m.mu.RLock()
	defer m.mu.RUnlock()

	sd.Data = m.Content
	return sd, nil
}
