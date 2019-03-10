package window

import (
	"github.com/byliuyang/ui/graphics"
	"github.com/byliuyang/ui/view"
)

type DrawFunc func()

type Window interface {
	Show(g graphics.Graphics, view view.View)
}
