package grifts

import (
	"github.com/Hagelund96/prosjekt/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
