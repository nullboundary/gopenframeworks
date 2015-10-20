Gopenframeworks
======================
Gopenframeworks (temporary project name) is a proof of concept exploring the possibility of using go to create a graphics framework inspired by openframeworks. All contributions to the framework are welcome, but this framework is NOT ready for use yet.

Installation
======================

* Make sure that Go is installed on your computer.
* Follow the instructions to install GLFW.
https://github.com/go-gl/glfw3
* Once you have installed Go and GLFW:
```
go get github.com/nullboundary/gopenframeworks
```



Example
======================

```go
package main

import (
	"fmt"
	gof "github.com/nullboundary/gopenframeworks"
)

var y float32 = 100
var width float32 = 1024
var height float32 = 768

//--------------------------------------------------------------
func (app testApp) Setup() {
	fmt.Println("Hello World")
	gof.Background(0, 0, 0.2)

}

//--------------------------------------------------------------
func (app testApp) Update() {

	y = y - 1
	if y < 0 {
		y = height
	}
}

//--------------------------------------------------------------
func (app testApp) Draw() {

	redTriangle := gof.Triangle(0, 0, 0, width, 0, 0, width/2, height, 0)
	redTriangle.Fill(1, 0, 0, 1.0) //r,g,b,a
	redTriangle.Draw()

	cutOutRect := gof.Rectangle((width/2)-150.0, 0, 300, 300) //x,y,w,h
	cutOutRect.Fill(0, 0, 0.2, 0.4)                 //r,g,b,a
	cutOutRect.Draw()
	
	movingLine := gof.NewLine(0, y, 0, width, y, 0)
	movingLine.Fill(1, 0, 0, 1.0)
	movingLine.Draw()
}
```
![Image](examples/example.png?raw=true)
