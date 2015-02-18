Gopenframeworks
======================
Gopenframeworks (temporary project name) is a proof of concept exploring the possiblity of using go to create a graphics framework inspired by openframeworks. All contributions to the framework are welcome, but this framework is ```NOT``` ready for use yet. 

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

	redTriangle := gof.Triangle(100, 120, 0, 120, 120, 0, 310, 340, 0)
	redTriangle.Fill(1, 0, 0, 1.0) //r,g,b,a
	redTriangle.Draw()

	greenRect := gof.Rectangle(300, 129, 339, 599) //x,y,w,h
	greenRect.Fill(0, 1, 0.5, 0.2)                 //r,g,b,a
	greenRect.Draw()
	
	movingLine := gof.NewLine(0, y, 0, width, y, 0)
	movingLine.Fill(1, 0, 0, 1.0)
	movingLine.Draw()
}
```
