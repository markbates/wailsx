package eventxtest

// import (
// 	"context"
// 	"fmt"
// 	"testing"

// 	"github.com/stretchr/testify/require"
// )

// func Test_Manager_On(t *testing.T) {
// 	t.Parallel()
// 	r := require.New(t)

// 	em, err := NewManager()
// 	r.NoError(err)
// 	r.NotNil(em)

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	const event = "event:test"

// 	var many int
// 	ec, err := em.EventsOn(ctx, event, func(data ...any) error {
// 		if len(data) != 1 {
// 			return fmt.Errorf("expected 1 arg, got %d", len(data))
// 		}

// 		i, ok := data[0].(int)

// 		if !ok {
// 			return fmt.Errorf("expected int, got %T", data[0])
// 		}
// 		many += i

// 		return nil
// 	})
// 	r.NoError(err)

// 	err = em.EventsEmit(ctx, event, 5)
// 	r.NoError(err)
// 	r.Equal(5, many)

// 	err = em.EventsEmit(ctx, event, 7)
// 	r.NoError(err)
// 	r.Equal(12, many)

// 	r.NoError(ec())

// 	err = em.EventsEmit(ctx, event, 42)
// 	r.NoError(err)
// 	r.Equal(12, many)

// 	r.NoError(err)
// }

// func Test_Manager_OnMultiple(t *testing.T) {
// 	t.Parallel()
// 	r := require.New(t)

// 	em, err := NewManager()
// 	r.NoError(err)
// 	r.NotNil(em)

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	const event = "event:test"

// 	ec, err := em.EventsOnMultiple(ctx, event, func(data ...any) error {
// 		return nil
// 	}, 5)
// 	r.NoError(err)
// 	r.NotNil(ec)

// 	for i := 0; i < 10; i++ {
// 		err = em.EventsEmit(ctx, event)
// 		r.NoError(err)
// 	}

// 	data, ok := em.Callbacks[event]
// 	r.True(ok)
// 	r.Equal(5, data.Called)

// 	_, ok = em.Callbacks[event]
// 	r.True(ok)
// 	r.NoError(ec())

// 	data, ok = em.Callbacks[event]
// 	r.True(ok)
// 	r.True(data.Off)
// }

// func Test_Manager_Off(t *testing.T) {
// 	t.Parallel()
// 	r := require.New(t)

// 	em, err := NewManager()
// 	r.NoError(err)
// 	r.NotNil(em)

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	const event = "event:test"

// 	ec, err := em.EventsOn(ctx, event, func(data ...any) error {
// 		return nil
// 	})
// 	r.NoError(err)
// 	r.NotNil(ec)

// 	err = em.EventsEmit(ctx, event)
// 	r.NoError(err)

// 	data, ok := em.Callbacks[event]
// 	r.True(ok)
// 	r.Equal(1, data.Called)
// 	r.False(data.Off)

// 	err = em.EventsOff(ctx, event)
// 	r.NoError(err)

// 	data, ok = em.Callbacks[event]
// 	r.True(ok)
// 	r.True(data.Off)
// }

// func Test_Manager_OffAll(t *testing.T) {
// 	t.Parallel()
// 	r := require.New(t)

// 	em, err := NewManager()
// 	r.NoError(err)
// 	r.NotNil(em)

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	events := []string{"event:test1", "event:test2", "event:test3"}

// 	for _, event := range events {
// 		_, err := em.EventsOn(ctx, event, func(data ...any) error {
// 			return nil
// 		})
// 		r.NoError(err)
// 	}

// 	r.Len(em.Callbacks, len(events))

// 	err = em.EventsOffAll(ctx)
// 	r.NoError(err)

// 	r.Len(em.Callbacks, 3)

// 	for _, event := range events {
// 		data, ok := em.Callbacks[event]
// 		r.True(ok)
// 		r.True(data.Off)
// 	}

// }

// func Test_Manager_Once(t *testing.T) {
// 	t.Parallel()
// 	r := require.New(t)

// 	em, err := NewManager()
// 	r.NoError(err)
// 	r.NotNil(em)

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	const event = "event:test"

// 	ec, err := em.EventsOnce(ctx, event, func(data ...any) error {
// 		return nil
// 	})
// 	r.NoError(err)
// 	r.NotNil(ec)

// 	data, ok := em.Callbacks[event]
// 	r.True(ok)
// 	r.Equal(1, data.MaxCalls)
// 	r.Equal(0, data.Called)

// 	for i := 0; i < 10; i++ {
// 		err = em.EventsEmit(ctx, event)
// 		r.NoError(err)
// 	}

// 	data, ok = em.Callbacks[event]
// 	r.True(ok)
// 	r.Equal(1, data.Called)
// 	r.False(data.Off)

// 	r.NoError(ec())

// 	data, ok = em.Callbacks[event]
// 	r.True(ok)
// 	r.True(data.Off)

// }
