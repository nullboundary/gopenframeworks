package main

import (
	gof "github.com/nullboundary/gopenframeworks"
)

type testApp struct {
	appName    string
	width      int
	height     int
	windowType string
}

func main() {

	app := new(testApp)
	app.appName = "particle"
	app.width = 1024
	app.height = 768
	app.windowType = "WINDOW" //or FULLSCREEN

	gof.SetupOpenGL(app.width, app.height, app.windowType)

	gof.RunApp(app)

}
