package gopenframeworks

import (
	"fmt"
	"github.com/go-gl/gl/v4.1-core/gl"
	"strings"
)

var shadeProg uint32

const (
	//vertexPixel converts pixel resolution to opengl coordinate space 0 to 1
	vertexPixel = `#version 330

in vec2 position;
uniform vec2 resolution;

void main() {
   // convert the rectangle from pixels to 0.0 to 1.0
   vec2 zeroToOne = position / resolution;

   // convert from 0->1 to 0->2
   vec2 zeroToTwo = zeroToOne * 2.0;

   // convert from 0->2 to -1->+1 (clipspace)
   vec2 clipSpace = zeroToTwo - 1.0;

   gl_Position = vec4(clipSpace * vec2(1, -1), 0, 1);
}`

	fragColor = `#version 330
uniform vec4 color;
out vec4 outColor;

void main() {
   outColor = color;
}`
)

//compileShader
func compileShader(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csource := gl.Str(source)
	gl.ShaderSource(shader, 1, &csource, nil)
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

//newShaderProgram
func newShaderProgram(vert string, frag string) error {

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)

	vertexShader, err := compileShader(vert+"\x00", gl.VERTEX_SHADER)
	if err != nil {
		return err
	}
	defer gl.DeleteShader(vertexShader)

	fragmentShader, err := compileShader(frag+"\x00", gl.FRAGMENT_SHADER)
	if err != nil {
		return err
	}
	defer gl.DeleteShader(fragmentShader)

	shadeProg = gl.CreateProgram()

	gl.AttachShader(shadeProg, vertexShader)
	gl.AttachShader(shadeProg, fragmentShader)
	gl.BindFragDataLocation(shadeProg, 0, gl.Str("outColor\x00"))
	gl.LinkProgram(shadeProg)
	gl.UseProgram(shadeProg)

	var status int32
	gl.GetProgramiv(shadeProg, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shadeProg, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shadeProg, logLength, nil, gl.Str(log))

		return fmt.Errorf("failed to link program: %v", log)
	}

	resUniform := gl.GetUniformLocation(shadeProg, gl.Str("resolution\x00"))
	gl.Uniform2f(resUniform, float32(appWindow.Width), float32(appWindow.Height))

	return nil

}
