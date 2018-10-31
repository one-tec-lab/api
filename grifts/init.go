package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/one-tec-lab/api/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
