// This is a generated file. DO NOT EDIT.

package structtest

import glfw "github.com/pekim/glfw"

// #include "../generate/glfw3.h"
import "C"

var structs = []struct {
	name string
	c    any
	go_  any
}{
	{
		c:    C.GLFWmonitor{},
		go_:  glfw.Monitor{},
		name: "GLFWmonitor",
	}, {
		c:    C.GLFWwindow{},
		go_:  glfw.Window{},
		name: "GLFWwindow",
	}, {
		c:    C.GLFWcursor{},
		go_:  glfw.Cursor{},
		name: "GLFWcursor",
	}, {
		c:    C.GLFWvidmode{},
		go_:  glfw.Vidmode{},
		name: "GLFWvidmode",
	}, {
		c:    C.GLFWgammaramp{},
		go_:  glfw.Gammaramp{},
		name: "GLFWgammaramp",
	}, {
		c:    C.GLFWimage{},
		go_:  glfw.Image{},
		name: "GLFWimage",
	}, {
		c:    C.GLFWgamepadstate{},
		go_:  glfw.Gamepadstate{},
		name: "GLFWgamepadstate",
	}, {
		c:    C.GLFWallocator{},
		go_:  glfw.Allocator{},
		name: "GLFWallocator",
	},
}
