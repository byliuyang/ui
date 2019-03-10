package view

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockGraphics struct {
	x int
	y int
}

func (g *MockGraphics) Clear() {
	g.x = 0
	g.y = 0
}

func (g *MockGraphics) FillRect(x int, y int, width int, height int) {
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
	testCases := []struct {
		width  int
		height int

		marginLeft   int
		marginTop    int
		marginRight  int
		marginBottom int

		paddingLeft   int
		paddingTop    int
		paddingRight  int
		paddingBottom int

		backgroundColor Color

		testName string

		expectedX int
		expectedY int
	}{
		{
			width:  10,
			height: 20,

			marginLeft:   0,
			marginTop:    0,
			marginRight:  0,
			marginBottom: 0,

			paddingLeft:   0,
			paddingTop:    0,
			paddingRight:  0,
			paddingBottom: 0,

			backgroundColor: Color{},

			testName: "width_height",

			expectedX: 10,
			expectedY: 20,
		},
		{
			width:  0,
			height: 0,

			marginLeft:   1,
			marginTop:    2,
			marginRight:  3,
			marginBottom: 4,

			paddingLeft:   0,
			paddingTop:    0,
			paddingRight:  0,
			paddingBottom: 0,

			backgroundColor: Color{},

			testName: "margin",

			expectedX: 4,
			expectedY: 6,
		},
		{
			width:  0,
			height: 0,

			marginLeft:   0,
			marginTop:    0,
			marginRight:  0,
			marginBottom: 0,

			paddingLeft:   1,
			paddingTop:    2,
			paddingRight:  3,
			paddingBottom: 4,

			backgroundColor: Color{},

			testName: "padding",

			expectedX: 4,
			expectedY: 6,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.testName, func(t1 *testing.T) {
			container := Container{
				Base: Base{
					Width:  testCase.width,
					Height: testCase.height,

					MarginLeft:   testCase.marginLeft,
					MarginTop:    testCase.marginTop,
					MarginRight:  testCase.marginRight,
					MarginBottom: testCase.marginBottom,
				},

				PaddingLeft:   testCase.paddingLeft,
				PaddingTop:    testCase.paddingTop,
				PaddingRight:  testCase.paddingRight,
				PaddingBottom: testCase.paddingBottom,

				BackgroundColor: testCase.backgroundColor,
			}

			g := &MockGraphics{
				x: 0,
				y: 0,
			}

			assert.Equal(t1, 0, g.x)
			assert.Equal(t1, 0, g.y)

			container.Render(g)

			assert.Equal(t1, testCase.expectedX, g.x)
			assert.Equal(t1, testCase.expectedY, g.y)
		})
	}
}
