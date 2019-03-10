package opengl

import (
	"fmt"
	"github.com/byliuyang/ui/graphics"
	"github.com/go-gl/gl/v4.1-core/gl"
)

type OpenGl struct {
	x          int
	y          int
	programRef Ref
	vao        Ref
}

func (o *OpenGl) GetX() int {
	return o.x
}

func (o *OpenGl) GetY() int {
	return o.y
}

func (OpenGl) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

func (o *OpenGl) MoveTo(x int, y int) {
	o.x = x
	o.y = y
}

func (o *OpenGl) FillRect(x int, y int, width int, height int) {
	gl.BindBuffer(gl.ARRAY_BUFFER, o.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
}

func NewOpenGl() graphics.Graphics {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version", version)

	programRef := newProgram()
	gl.UseProgram(programRef)

	points := []float32{
		0, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
	}

	vao := bufferData(points)

	return &OpenGl{
		programRef: programRef,
		vao:        vao,
	}
}

func bufferData(points []float32) Ref {
	var vbo Ref
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao Ref
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)
	return vao
}