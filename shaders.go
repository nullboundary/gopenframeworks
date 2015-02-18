package gopenframeworks

import (
	"github.com/go-gl/gl"
	"log"
)

var shadeProg gl.Program

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

func compileShaders(vert string, frag string) {

	vao := gl.GenVertexArray()
	vao.Bind()

	//setup default shaders
	vertex_shader := gl.CreateShader(gl.VERTEX_SHADER)
	vertex_shader.Source(vert)
	vertex_shader.Compile()
	log.Println(vertex_shader.GetInfoLog())
	defer vertex_shader.Delete()

	fragment_shader := gl.CreateShader(gl.FRAGMENT_SHADER)
	fragment_shader.Source(frag)
	fragment_shader.Compile()
	log.Println(fragment_shader.GetInfoLog())
	defer fragment_shader.Delete()

	shadeProg = gl.CreateProgram()
	shadeProg.AttachShader(vertex_shader)
	shadeProg.AttachShader(fragment_shader)

	shadeProg.BindFragDataLocation(0, "outColor")
	shadeProg.Link()
	shadeProg.Use()

	resUniform := shadeProg.GetUniformLocation("resolution")
	resUniform.Uniform2f(float32(appWindow.Width), float32(appWindow.Height))

}
