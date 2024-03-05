package eventx_test

// func Test_EventManager_On(t *testing.T) {
// 	t.Parallel()
// 	r := require.New(t)

// 	em, ec := newEventManager()
// 	_ = ec

// 	const evt = "event:test"

// 	em.EventsOnFn = func(ctx context.Context, name string, callback wailsrun.CallbackFn) (wailsrun.CancelFn, error) {
// 		if name != evt {
// 			return nil, wailstest.ErrTest
// 		}

// 		if err := callback(); err != nil {
// 			return nil, err
// 		}

// 		return func() error {
// 			return nil
// 		}, nil
// 	}

// 	tcs := []struct {
// 		name string
// 		cb   wailsrun.CallbackFn
// 		err  bool
// 	}{
// 		{
// 			name: "no error",
// 			cb:   func(data ...any) error { return nil },
// 		},
// 	}

// 	for _, tc := range tcs {
// 		tc := tc
// 		t.Run(tc.name, func(t *testing.T) {
// 			t.Parallel()

// 			_, err := em.EventsOn(context.Background(), evt, tc.cb)

// 			if !tc.err {
// 				r.NoError(err)
// 				return
// 			}

// 			r.Error(err)
// 			r.True(errors.Is(err, wailstest.ErrTest))
// 		})
// 	}

// }
