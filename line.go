package gopenframeworks

import (
	"github.com/go-gl/gl/v4.1-core/gl"
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
	attribLoc := uint32(gl.GetAttribLocation(shadeProg, gl.Str("position\x00")))
	gl.EnableVertexAttribArray(attribLoc)

	//what format our vertex array data in the buffer object is stored in
	gl.VertexAttribPointer(attribLoc, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))

	defer gl.DisableVertexAttribArray(attribLoc)
	numVertices := int32(len(l.vertices))
	gl.DrawArrays(gl.LINES, 0, numVertices)
}

func (l *Line) Fill(red float32, green float32, blue float32, alpha float32) {
	colorUniform := gl.GetUniformLocation(shadeProg, gl.Str("color\x00"))
	gl.Uniform4f(colorUniform, red, green, blue, alpha)
}

func (l *Line) Stroke(red float32, green float32, blue float32, alpha float32) {

}
