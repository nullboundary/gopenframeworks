package gopenframeworks

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.1/glfw"
	"log"
	"runtime"
)

var (
	gofApp    BaseApp
	appWindow *Window
)

//BaseApp Interface
type BaseApp interface {
	Controllers
	Handlers
}

type Controllers interface {
	Setup()
	Update()
	Draw()
}

type Handlers interface {
	KeyPressed(key int)
	KeyReleased(key int)
	MouseMoved(x int, y int)
	MouseDragged(x int, y int, button int)
	MousePressed(x int, y int, button int)
	MouseReleased(x int, y int, button int)
	WindowResized(w int, h int)
}

func SetupOpenGL(w int, h int, window string) {

	appWindow = &Window{}
	appWindow.Width = w
	appWindow.Height = h
	appWindow.WType = window
	appWindow.Title = "gopenframeworks"

}

func RunApp(app BaseApp) {

	// lock glfw/gl calls to a single thread
	runtime.LockOSThread()

	appWindow.Create()
	defer glfw.Terminate()
	defer appWindow.GLFWindow.Destroy()

	initGL()
	err := newShaderProgram(vertexPixel, fragColor)
	if err != nil {
		log.Fatalf("shader error: %s", err)
	}
	defer gl.DeleteProgram(shadeProg)

	app.Setup()

	for !appWindow.GLFWindow.ShouldClose() {

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Update animation
		app.Update()

		//draw
		app.Draw()

		appWindow.GLFWindow.SwapBuffers()
		glfw.PollEvents()

		if appWindow.GLFWindow.GetKey(glfw.KeyEscape) == glfw.Press {
			appWindow.GLFWindow.SetShouldClose(true)
		}
	}

}

//Background set the background color of the context
func Background(r float32, g float32, b float32) {
	//	red := gl.GLclampf(r) //floating-point value, clamped to the range [0,1]
	//	green := gl.GLclampf(g)
	//	blue := gl.GLclampf(b)

	gl.ClearColor(r, g, b, 1.0)
}

func initGL() {

	err := gl.Init()
	if err != nil {
		log.Println(err)
	}

}
