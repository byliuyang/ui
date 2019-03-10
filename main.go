package main

import (
	"github.com/byliuyang/ui/graphics/opengl"
	"github.com/byliuyang/ui/view"
	"github.com/byliuyang/ui/window"
	"runtime"
)

func main() {
	runtime.LockOSThread()

	win := window.NewGlfw()
	openGl := opengl.NewOpenGl()
	win.Show(openGl, view.NewRoot())
}
