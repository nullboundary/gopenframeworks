package gopenframeworks

import ()

type Rect struct {
	vertices []float32
}

func Rectangle(x float32, y float32, width float32, height float32) *Rect {

	var x1 = x
	var x2 = x + width
	var y1 = y
	var y2 = y + height

	/*x1, y1,
	  x2, y1,
	  x1, y2,
	  x1, y2,
	  x2, y1,
	  x2, y2*/

	rect := &Rect{}
	rect.vertices = []float32{x1, y1, 0, x2, y1, 0, x1, y2, 0, x1, y2, 0, x2, y1, 0, x2, y2, 0}

	return rect

}

func (r *Rect) Draw() {
	Mesh(r.vertices...)
}

func (r *Rect) Fill(red float32, green float32, blue float32, alpha float32) {
	colorUniform := shadeProg.GetUniformLocation("color")
	colorUniform.Uniform4f(red, green, blue, alpha)
}

func (r *Rect) Stroke(red float32, green float32, blue float32, alpha float32) {

}
