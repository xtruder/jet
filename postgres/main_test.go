//go:generate go test ./. -testparrot.record -testparrot.splitfiles
package postgres

import (
	"testing"

	testparrot "github.com/xtruder/go-testparrot"
)

func TestMain(m *testing.M) {
	testparrot.Run(m)
}
