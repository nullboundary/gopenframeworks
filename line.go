package gopenframeworks

import (
	"github.com/go-gl/gl"
)

type Line struct {
	vertices []float32
}

func NewLine(x1 float32, y1 float32, z1 float32, x2 float32, y2 float32, z2 float32) *Line {

	l := &Line{}
	l.vertices = []float32{x1, y1, z1, x2, y2, z2}
	return l
}

func (l *Line) Draw() {

	toVertexBuffer(l.vertices)

	attribLoc := shadeProg.GetAttribLocation("position")
	attribLoc.AttribPointer(3, gl.FLOAT, false, 0, nil) //what format our vertex array data in the buffer object is stored in
	attribLoc.EnableArray()
	defer attribLoc.DisableArray()

	gl.DrawArrays(gl.LINES, 0, len(l.vertices))
}

func (l *Line) Fill(red float32, green float32, blue float32, alpha float32) {
	colorUniform := shadeProg.GetUniformLocation("color")
	colorUniform.Uniform4f(red, green, blue, alpha)
}

func (l *Line) Stroke(red float32, green float32, blue float32, alpha float32) {

}
