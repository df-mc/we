package act

import (
	"github.com/df-mc/we/brush"
)

func init() {
	brush.RegisterAction("Fill", func() brush.Action { return Fill{} })
}
