package wailsx

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func newEmitter() (Emitter, *EmitCatcher) {
	ec := &EmitCatcher{}
	return Emitter{
		EmitFn:          ec.Emit,
		DisableWildcard: true,
	}, ec
}

func assertJSON(t testing.TB, fp string, msg Messenger) {
	t.Helper()

	r := require.New(t)

	b, err := json.MarshalIndent(msg, "", "  ")
	r.NoError(err)

	act := string(b)
	act = strings.TrimSpace(act)

	fmt.Println(act)

	fp = filepath.Join("testdata", fp+".json")

	// os.MkdirAll(filepath.Dir(fp), 0755)
	// f, err := os.Create(fp)
	// r.NoError(err)
	// f.Write([]byte(act))
	// r.NoError(f.Close())

	b, err = os.ReadFile(fp)
	r.NoError(err)

	exp := string(b)
	exp = strings.TrimSpace(exp)

	r.Equal(exp, act)
}

func nowTime() time.Time {
	return time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
}

func oldTime() time.Time {
	return time.Date(1976, 1, 1, 0, 0, 0, 0, time.UTC)
}
