package gopenframeworks

import (
	"github.com/go-gl/glfw/v3.1/glfw"
	"log"
)

type Window struct {
	Width     int
	Height    int
	WType     string //fullscreen or window
	Title     string //title for the window
	GLFWindow *glfw.Window
}

//Create initializes the main window
func (w *Window) Create() {

	//start glfw
	err := glfw.Init()
	if err != nil {
		panic("Failed to initialize GLFW")
	}

	//create window
	glfw.WindowHint(glfw.Samples, 4)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	glfw.WindowHint(glfw.OpenGLDebugContext, 1)

	log.Printf("%d %d %s", w.Height, w.Width, w.Title)

	w.GLFWindow, err = glfw.CreateWindow(w.Width, w.Height, w.Title, nil, nil)
	if err != nil {
		panic(err)
	}

	// Set callback functions
	w.GLFWindow.SetFramebufferSizeCallback(resizeCallback)
	w.GLFWindow.SetKeyCallback(keyCallBack)
	w.GLFWindow.SetMouseButtonCallback(mouseButtonCallBack)

	w.GLFWindow.MakeContextCurrent()
	glfw.SwapInterval(1)

	sizeW, sizeH := w.GLFWindow.GetFramebufferSize()
	resizeCallback(w.GLFWindow, sizeW, sizeH)

}

// resizeCallback sets a new window size
func resizeCallback(window *glfw.Window, width, height int) {

	log.Println(float64(width))

	if width < 1 {
		width = 1
	}

	if height < 1 {
		height = 1
	}

	//gofApp.WindowResized(width, height)
}

//exit upon ESC
func keyCallBack(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Press {
		return
	}

	if k == glfw.KeyEscape {
		window.SetShouldClose(true)
	}

}

func mouseButtonCallBack(window *glfw.Window, button glfw.MouseButton, action glfw.Action, mods glfw.ModifierKey) {

}
