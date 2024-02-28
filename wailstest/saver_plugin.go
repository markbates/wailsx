package wailstest

import (
	"context"
	"fmt"
)

type SaverPlugin struct {
	Saved bool
}

func (s *SaverPlugin) PluginName() string {
	return fmt.Sprintf("%T", s)
}

func (s *SaverPlugin) Save(ctx context.Context) error {
	s.Saved = true
	return nil
}
