package view

import (
	"github.com/byliuyang/ui/graphics"
)

type View struct {
	Width int
	Height int

	MarginLeft int
	MarginTop int
	MarginRight int
	MarginBottom int
}


func (v *View) Render(g graphics.Graphics) {
	newX := g.GetX() + v.MarginLeft + v.Width + v.MarginRight
	newY := g.GetY() + v.MarginTop + v.Height + v.MarginBottom

	g.MoveTo(newX, newY)
}