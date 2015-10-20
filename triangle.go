package gopenframeworks

import (
	"github.com/go-gl/gl/v4.1-core/gl"
)

type Tri struct {
	vertices []float32
}

func Triangle(x1 float32, y1 float32, z1 float32, x2 float32, y2 float32, z2 float32, x3 float32, y3 float32, z3 float32) *Tri {

	t := &Tri{}
	t.vertices = []float32{x1, y1, z1, x2, y2, z2, x3, y3, z3}

	return t
}

func (t *Tri) Draw() {
	Mesh(t.vertices...)
}

func (r *Tri) Fill(red float32, green float32, blue float32, alpha float32) {
	colorUniform := gl.GetUniformLocation(shadeProg, gl.Str("color"+"\x00"))
	gl.Uniform4f(colorUniform, red, green, blue, alpha)
}

func (r *Tri) Stroke(red float32, green float32, blue float32, alpha float32) {

}
