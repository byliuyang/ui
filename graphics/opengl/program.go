package opengl

import "github.com/go-gl/gl/v4.1-core/gl"

func newProgram() Ref {
	vertexShaderRef, err := newVertexShader()
	if err != nil {
		panic("fail to compile vertex shader")
	}

	fragmentShaderRef, err := newFragmentShader()
	if err != nil {
		panic("fail to compile fragment shader")
	}

	programRef := gl.CreateProgram()
	gl.AttachShader(programRef, vertexShaderRef)
	gl.AttachShader(programRef, fragmentShaderRef)
	gl.LinkProgram(programRef)
	return programRef
}
