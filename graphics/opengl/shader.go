package opengl

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"strings"
)

type ShaderType int

const (
	VertexShader ShaderType = iota
	FragmentShader
)

type ShaderBuilder struct {
	source     string
	shaderType uint32
}

func newShaderBuilder() *ShaderBuilder {
	return &ShaderBuilder{}
}

func (b *ShaderBuilder) Source(source string) *ShaderBuilder {
	b.source = source
	return b
}

func (b *ShaderBuilder) ShaderType(shaderType ShaderType) *ShaderBuilder {
	switch shaderType {
	case VertexShader:
		b.shaderType = gl.VERTEX_SHADER
	case FragmentShader:
		b.shaderType = gl.FRAGMENT_SHADER
	}
	return b
}

func (b *ShaderBuilder) Build() (Ref, error) {
	shader := gl.CreateShader(b.shaderType)
	sourcePtr, free := gl.Strs(b.source + "\x00")
	gl.ShaderSource(shader, 1, sourcePtr, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))
		return 0, fmt.Errorf("failed to compiler %v: %v", b.source, log)
	}

	return shader, nil
}

func newVertexShader() (Ref, error) {
	return newShaderBuilder().
		ShaderType(VertexShader).
		Source(`
#version 410

in vec3 vertex;

void main() {
    gl_Position = vec4(vertex, 1.0);
}
`).Build()
}

func newFragmentShader() (Ref, error) {
	return newShaderBuilder().
		ShaderType(FragmentShader).
		Source(`
#version 410
out vec4 outputColor;

void main() {
	outputColor = vec4(1, 1, 1, 1.0);
}
`).Build()
}
