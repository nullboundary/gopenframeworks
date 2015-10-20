package gopenframeworks

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
)

type Point4 mgl32.Vec4
type Vec2 mgl32.Vec2
type Vec3 mgl32.Vec3
type Color4 mgl32.Vec4

type Shape2D interface {
	Fill(red float32, green float32, blue float32, alpha float32)
	Stroke(red float32, green float32, blue float32, alpha float32)
	Draw()
}

func toVertexBuffer(vertexPos []float32) uint32 {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)                                                 //bind to array vbo target
	gl.BufferData(gl.ARRAY_BUFFER, len(vertexPos)*4, gl.Ptr(vertexPos), gl.STREAM_DRAW) //add vertex data to vbo

	return vbo
}

func Mesh(pos ...float32) {

	toVertexBuffer(pos)

	attribLoc := uint32(gl.GetAttribLocation(shadeProg, gl.Str("position\x00")))
	gl.EnableVertexAttribArray(attribLoc)
	gl.VertexAttribPointer(attribLoc, 3, gl.FLOAT, false, 0, gl.PtrOffset(0))
	defer gl.DisableVertexAttribArray(attribLoc)

	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(pos)))
	//gl.DrawArrays(gl.LINES, 0, len(pos))

}
