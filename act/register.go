package act

import (
	"github.com/df-mc/we/brush"
)

// init registers all brush.Action implementations in the act package.
func init() {
	brush.RegisterAction("Fill", func() brush.Action { return Fill{} })
	brush.RegisterAction("Replace", func() brush.Action { return Replace{} })
}
