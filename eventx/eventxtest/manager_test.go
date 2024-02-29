package eventxtest

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Manager_On(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	em, err := NewManager()
	r.NoError(err)
	r.NotNil(em)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const event = "event:test"

	var many int
	ec, err := em.On(ctx, event, func(data ...any) error {
		if len(data) != 1 {
			return fmt.Errorf("expected 1 arg, got %d", len(data))
		}

		i, ok := data[0].(int)

		if !ok {
			return fmt.Errorf("expected int, got %T", data[0])
		}
		many += i

		return nil
	})
	r.NoError(err)

	err = em.Emit(ctx, event, 5)
	r.NoError(err)
	r.Equal(5, many)

	err = em.Emit(ctx, event, 7)
	r.NoError(err)
	r.Equal(12, many)

	r.NoError(ec())

	err = em.Emit(ctx, event, 42)
	r.NoError(err)
	r.Equal(12, many)

	r.NoError(err)
}

func Test_Manager_OnMultiple(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	em, err := NewManager()
	r.NoError(err)
	r.NotNil(em)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const event = "event:test"

	ec, err := em.OnMultiple(ctx, event, func(data ...any) error {
		return nil
	}, 5)
	r.NoError(err)
	r.NotNil(ec)

	for i := 0; i < 10; i++ {
		err = em.Emit(ctx, event)
		r.NoError(err)
	}

	data, ok := em.Callbacks[event]
	r.True(ok)
	r.Equal(5, data.Called)

	_, ok = em.Callbacks[event]
	r.True(ok)
	r.NoError(ec())

	data, ok = em.Callbacks[event]
	r.True(ok)
	r.True(data.Off)
}

func Test_Manager_Off(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	em, err := NewManager()
	r.NoError(err)
	r.NotNil(em)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const event = "event:test"

	ec, err := em.On(ctx, event, func(data ...any) error {
		return nil
	})
	r.NoError(err)
	r.NotNil(ec)

	err = em.Emit(ctx, event)
	r.NoError(err)

	data, ok := em.Callbacks[event]
	r.True(ok)
	r.Equal(1, data.Called)
	r.False(data.Off)

	err = em.Off(ctx, event)
	r.NoError(err)

	data, ok = em.Callbacks[event]
	r.True(ok)
	r.True(data.Off)
}

func Test_Manager_OffAll(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	em, err := NewManager()
	r.NoError(err)
	r.NotNil(em)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	events := []string{"event:test1", "event:test2", "event:test3"}

	for _, event := range events {
		_, err := em.On(ctx, event, func(data ...any) error {
			return nil
		})
		r.NoError(err)
	}

	r.Len(em.Callbacks, len(events))

	err = em.OffAll(ctx)
	r.NoError(err)

	r.Len(em.Callbacks, 3)

	for _, event := range events {
		data, ok := em.Callbacks[event]
		r.True(ok)
		r.True(data.Off)
	}

}

func Test_Manager_Once(t *testing.T) {
	t.Parallel()
	r := require.New(t)

	em, err := NewManager()
	r.NoError(err)
	r.NotNil(em)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	const event = "event:test"

	ec, err := em.Once(ctx, event, func(data ...any) error {
		return nil
	})
	r.NoError(err)
	r.NotNil(ec)

	data, ok := em.Callbacks[event]
	r.True(ok)
	r.Equal(1, data.MaxCalls)
	r.Equal(0, data.Called)

	for i := 0; i < 10; i++ {
		err = em.Emit(ctx, event)
		r.NoError(err)
	}

	data, ok = em.Callbacks[event]
	r.True(ok)
	r.Equal(1, data.Called)
	r.False(data.Off)

	r.NoError(ec())

	data, ok = em.Callbacks[event]
	r.True(ok)
	r.True(data.Off)

}
