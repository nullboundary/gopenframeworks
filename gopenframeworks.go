package gopenframeworks

import (
	"github.com/go-gl/gl"
	glfw "github.com/go-gl/glfw3"
	"github.com/go-gl/glh"
	"github.com/go-gl/glu"
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
	compileShaders(vertexPixel, fragColor)
	defer shadeProg.Delete()

	app.Setup()

	for !appWindow.GLFWindow.ShouldClose() {

		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Update animation
		app.Update()

		//draw
		app.Draw()

		appWindow.GLFWindow.SwapBuffers()
		glfw.PollEvents()
		//checkGLErrors()

		if appWindow.GLFWindow.GetKey(glfw.KeyEscape) == glfw.Press {
			appWindow.GLFWindow.SetShouldClose(true)
		}
	}

}

//Background set the background color of the context
func Background(r float32, g float32, b float32) {

	red := gl.GLclampf(r)
	green := gl.GLclampf(g)
	blue := gl.GLclampf(b)

	gl.ClearColor(red, green, blue, 1.0)
}

func initGL() {

	errno := gl.Init()
	if errno != 0 {
		err := glh.CheckGLError()
		log.Println(err)
	}
	checkGLErrors()

}

func checkGLErrors() {

	errno := gl.GetError()
	for errno != gl.NO_ERROR {

		str, err := glu.ErrorString(errno)
		if err != nil {
			log.Printf("Unknown GL error: %d", errno)
		}

		log.Printf("Opengl Error:%s", str)
		errno = gl.GetError()
	}

}
