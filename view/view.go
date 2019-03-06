package view

import "github.com/byliuyang/ui/graphics"

type View interface {
	Render(g graphics.Graphics)
}
