package view

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockGraphics struct {
	x int
	y int
}

func (g *MockGraphics) GetX() int {
	return g.x
}

func (g *MockGraphics) GetY() int {
	return g.y
}

func (g *MockGraphics) MoveTo(x int, y int) {
	g.x = x
	g.y = y
}

func TestView_Render(t *testing.T) {
	view := View{
		Width:  10,
		Height: 20,

		MarginLeft:   1,
		MarginTop:    2,
		MarginRight:  3,
		MarginBottom: 4,
	}

	g := &MockGraphics{
		x: 0,
		y: 0,
	}

	view.Render(g)
	assert.Equal(t, 14, g.x)
	assert.Equal(t, 26, g.y)
}
