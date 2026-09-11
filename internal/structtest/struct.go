// This is a generated file. DO NOT EDIT.

package structtest

import goglfw "github.com/pekim/go-glfw"

// #include "../generate/glfw3.h"
import "C"

var structs = []struct {
	name string
	c    any
	go_  any
}{
	{
		c:    C.GLFWmonitor{},
		go_:  goglfw.Monitor{},
		name: "GLFWmonitor",
	}, {
		c:    C.GLFWwindow{},
		go_:  goglfw.Window{},
		name: "GLFWwindow",
	}, {
		c:    C.GLFWcursor{},
		go_:  goglfw.Cursor{},
		name: "GLFWcursor",
	}, {
		c:    C.GLFWvidmode{},
		go_:  goglfw.Vidmode{},
		name: "GLFWvidmode",
	}, {
		c:    C.GLFWgammaramp{},
		go_:  goglfw.Gammaramp{},
		name: "GLFWgammaramp",
	}, {
		c:    C.GLFWimage{},
		go_:  goglfw.Image{},
		name: "GLFWimage",
	}, {
		c:    C.GLFWgamepadstate{},
		go_:  goglfw.Gamepadstate{},
		name: "GLFWgamepadstate",
	}, {
		c:    C.GLFWallocator{},
		go_:  goglfw.Allocator{},
		name: "GLFWallocator",
	},
}
