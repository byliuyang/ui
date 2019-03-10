package view

import "github.com/byliuyang/ui/graphics"

type Container struct {
	Base

	PaddingLeft   int
	PaddingTop    int
	PaddingRight  int
	PaddingBottom int

	BackgroundColor Color
}

func NewContainer() View {
	return &Container{}
}

func (v *Container) Render(g graphics.Graphics) {
	x := g.GetX() + v.MarginLeft
	y := g.GetY() + v.MarginTop

	width := v.PaddingLeft + v.Width + v.PaddingRight
	height := v.PaddingTop + v.Height + v.PaddingBottom

	g.FillRect(x, y, width, height)

	newX := x + v.PaddingLeft + v.Width + v.PaddingRight + v.MarginRight
	newY := y + v.PaddingTop + v.Height + v.PaddingBottom + v.MarginBottom

	g.MoveTo(newX, newY)
}
