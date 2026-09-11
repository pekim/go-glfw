package main

import (
	"fmt"
	"runtime"

	"github.com/pekim/go-glfw"
)

func init() {
	// This is important to ensure that only a single thread makes calls to
	// openGL apis.
	runtime.LockOSThread()
}

func main() {
	err := glfw.Initialise()
	if err != nil {
		panic(err)
	}
	glfw.Init()

	glfw.WindowHint(glfw.CONTEXT_CREATION_API, glfw.EGL_CONTEXT_API)
	window := glfw.CreateWindow(800, 600, "example", nil, nil)

	window.SetSizeCallback(glfw.WindowsizeCallbackNew(func(_window *glfw.Window, width, height glfw.Int) {
		fmt.Println("window :", width, height)
	}))
	window.SetFramebufferSizeCallback(glfw.FramebuffersizeCallbackNew(func(_window *glfw.Window, width, height glfw.Int) {
		fmt.Println("framebuffer size :", width, height)
	}))
	window.SetFocusCallback(glfw.WindowfocusCallbackNew(func(_window *glfw.Window, focused glfw.Int) {
		fmt.Println("focused :", focused == glfw.TRUE)
		// window.SetShouldClose(TRUE)
	}))
	window.SetCloseCallback(glfw.WindowcloseCallbackNew(func(_window *glfw.Window) {
		fmt.Println("close window :", window.GetTitle())
	}))
	window.MakeContextCurrent()
	window.Show()

	for window.ShouldClose() != glfw.TRUE {
		window.SwapBuffers()
		glfw.WaitEvents()
	}
}
