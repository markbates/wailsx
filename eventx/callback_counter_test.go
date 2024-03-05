package eventx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CallbackCounter_Catch_All(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	cc := &CallbackCounter{}

	for i := 0; i < 10; i++ {
		b, err := cc.Catch()
		r.NoError(err)
		r.True(b)
	}
	r.Equal(10, cc.Called)
}

func Test_CallbackCounter_Catch_Off(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	cc := &CallbackCounter{
		Off: true,
	}

	b, err := cc.Catch()
	r.NoError(err)
	r.False(b)
	r.Equal(0, cc.Called)
}

func Test_CallbackCounter_Catch_MaxCalls(t *testing.T) {
	t.Parallel()

	r := require.New(t)

	cc := &CallbackCounter{
		MaxCalls: 5,
	}

	for i := 0; i < 10; i++ {
		b, err := cc.Catch()
		r.NoError(err)
		if i < 5 {
			r.True(b)
		} else {
			r.False(b)
		}
	}
	r.Equal(5, cc.Called)
}
