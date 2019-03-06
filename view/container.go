package view

import "github.com/byliuyang/ui/graphics"

type Container struct {
	Base

	PaddingLeft   int
	PaddingTop    int
	PaddingRight  int
	PaddingBottom int
}

func NewContainer() View {
	return &Container{}
}

func (v *Container) Render(g graphics.Graphics) {
	newX := g.GetX() + v.MarginLeft + v.PaddingLeft + v.Width + v.PaddingRight + v.MarginRight
	newY := g.GetY() + v.MarginTop + v.PaddingTop + v.Height + v.PaddingBottom + v.MarginBottom

	g.MoveTo(newX, newY)
}
