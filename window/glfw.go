package window

import (
	"github.com/byliuyang/ui/graphics"
	"github.com/byliuyang/ui/view"
	"github.com/go-gl/glfw/v3.2/glfw"
)

type Glfw struct {
	draw   DrawFunc
	window *glfw.Window
}

func (fw *Glfw) Show(g graphics.Graphics, view view.View) {
	defer glfw.Terminate()

	for !fw.window.ShouldClose() {
		g.Clear()

		view.Render(g)

		glfw.PollEvents()
		fw.window.SwapBuffers()
	}
}

type GlfwWindowBuilder struct {
	isResizable int
	hasBorder   int
	width       int
	height      int
}

func NewGlfwWindowBuilder() *GlfwWindowBuilder {
	return &GlfwWindowBuilder{}
}

func (b *GlfwWindowBuilder) IsResizable(isResizable bool) *GlfwWindowBuilder {
	b.isResizable = glfw.False
	if isResizable {
		b.isResizable = glfw.True
	}
	return b
}

func (b *GlfwWindowBuilder) Width(width int) *GlfwWindowBuilder {
	b.width = width
	return b
}

func (b *GlfwWindowBuilder) Height(height int) *GlfwWindowBuilder {
	b.height = height
	return b
}

func (b *GlfwWindowBuilder) HasBorder(hasBorder bool) *GlfwWindowBuilder {
	b.hasBorder = glfw.False
	if hasBorder {
		b.hasBorder = glfw.True
	}
	return b
}

func (b *GlfwWindowBuilder) Build() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, b.isResizable)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.Decorated, b.hasBorder)

	window, err := glfw.CreateWindow(b.width, b.height, "", nil, nil)
	if err != nil {
		panic(err)
	}

	window.MakeContextCurrent()
	return window
}

func NewGlfw() Window {
	glfwWindow := NewGlfwWindowBuilder().
		IsResizable(true).
		HasBorder(true).
		Width(800).
		Height(600).
		Build()

	window := &Glfw{
		window: glfwWindow,
	}

	return window
}
