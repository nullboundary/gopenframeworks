package gopenframeworks

import (
	"github.com/go-gl/gl"
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

func toVertexBuffer(vertexPos []float32) gl.Buffer {

	vbo := gl.GenBuffer()                                                       //creates the vbo object
	vbo.Bind(gl.ARRAY_BUFFER)                                                   //bind to array vbo target
	gl.BufferData(gl.ARRAY_BUFFER, len(vertexPos)*4, vertexPos, gl.STREAM_DRAW) //add vertex data to vbo

	return vbo
}

func Mesh(pos ...float32) {

	toVertexBuffer(pos)

	attribLoc := shadeProg.GetAttribLocation("position")
	attribLoc.AttribPointer(3, gl.FLOAT, false, 0, nil) //what format our vertex array data in the buffer object is stored in
	attribLoc.EnableArray()
	defer attribLoc.DisableArray()

	gl.DrawArrays(gl.TRIANGLES, 0, len(pos))
	//gl.DrawArrays(gl.LINES, 0, len(pos))

}
