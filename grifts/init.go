package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/pavanas/brokestest/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
