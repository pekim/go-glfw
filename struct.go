// This is a generated file. DO NOT EDIT.

package glfw

import (
	"structs"
	"unsafe"
)

/*
Opaque monitor object.

See [Monitor objects]

C documentation : [GLFWmonitor]

# since

Added in version 3.0.

[Monitor objects]: https://www.glfw.org/docs/latest/monitor_guide.html#monitor_object
[GLFWmonitor]: https://www.glfw.org/docs/latest/group__monitor.html#ga8d9efd1cde9426692c73fe40437d0ae3
*/
type Monitor struct {
	_ structs.HostLayout
}

/*
Opaque window object.

See [Window objects]

C documentation : [GLFWwindow]

# since

Added in version 3.0.

[Window objects]: https://www.glfw.org/docs/latest/window_guide.html#window_object
[GLFWwindow]: https://www.glfw.org/docs/latest/group__window.html#ga3c96d80d363e67d13a41b5d1821f3242
*/
type Window struct {
	_ structs.HostLayout
}

/*
Opaque cursor object.

See [Cursor objects]

C documentation : [GLFWcursor]

# since

Added in version 3.1.

[Cursor objects]: https://www.glfw.org/docs/latest/input_guide.html#cursor_object
[GLFWcursor]: https://www.glfw.org/docs/latest/group__input.html#ga89261ae18c75e863aaf2656ecdd238f4
*/
type Cursor struct {
	_ structs.HostLayout
}

/*
This describes a single video mode.

C documentation : [GLFWvidmode]

# since

Added in version 1.0.

[GLFWvidmode]: https://www.glfw.org/docs/latest/group__monitor.html#ga902c2816ac9b34b757282daab59b2565
*/
type Vidmode struct {
	_ structs.HostLayout
	/*
	   The width, in screen coordinates, of the video mode.
	*/
	Width Int
	/*
	   The height, in screen coordinates, of the video mode.
	*/
	Height Int
	/*
	   The bit depth of the red channel of the video mode.
	*/
	RedBits Int
	/*
	   The bit depth of the green channel of the video mode.
	*/
	GreenBits Int
	/*
	   The bit depth of the blue channel of the video mode.
	*/
	BlueBits Int
	/*
	   The refresh rate, in Hz, of the video mode.
	*/
	RefreshRate Int
}

/*
This describes the gamma ramp for a monitor.

C documentation : [GLFWgammaramp]

# since

Added in version 3.0.

[GLFWgammaramp]: https://www.glfw.org/docs/latest/group__monitor.html#ga939cf093cb0af0498b7b54dc2e181404
*/
type Gammaramp struct {
	_ structs.HostLayout
	/*
	   An array of value describing the response of the red channel.
	*/
	red *UShort
	/*
	   An array of value describing the response of the green channel.
	*/
	green *UShort
	/*
	   An array of value describing the response of the blue channel.
	*/
	blue *UShort
	/*
	   The number of elements in each array.
	*/
	size UInt
}

/*
This describes a single 2D image.  See the documentation for each related
function what the expected pixel format is.

C documentation : [GLFWimage]

# since

Added in version 2.1.

[GLFWimage]: https://www.glfw.org/docs/latest/group__window.html#ga7cc0a09de172fa7250872046f8c4d2ca
*/
type Image struct {
	_ structs.HostLayout
	/*
	   The width, in pixels, of this image.
	*/
	width Int
	/*
	   The height, in pixels, of this image.
	*/
	height Int
	/*
	   The pixel data of this image, arranged left-to-right, top-to-bottom.
	*/
	pixels *UChar
}

/*
This describes the input state of a gamepad.

C documentation : [GLFWgamepadstate]

# since

Added in version 3.3.

[GLFWgamepadstate]: https://www.glfw.org/docs/latest/group__input.html#ga61acfb1f28f751438dd221225c5e725d
*/
type Gamepadstate struct {
	_ structs.HostLayout
	/*
	   The states of each [gamepad button], `GLFW_PRESS`
	   or `GLFW_RELEASE`.

	   [gamepad button]: https://www.glfw.org/docs/latest/group__gamepad__buttons.html
	*/
	Buttons [15]UChar
	/*
	   The states of each [gamepad axis], in the range -1.0
	   to 1.0 inclusive.

	   [gamepad axis]: https://www.glfw.org/docs/latest/group__gamepad__axes.html
	*/
	Axes [6]Float
}

/*
This describes a custom heap memory allocator for GLFW.  To set an allocator, pass it
to [Init] before initializing the library.

C documentation : [GLFWallocator]

# since

Added in version 3.4.

[GLFWallocator]: https://www.glfw.org/docs/latest/group__init.html#ga145c57d7f2aeda0b704a5a4ba1d6104b
*/
type Allocator struct {
	_ structs.HostLayout
	/*
	   The memory allocation function.  See [AllocateCallback] for details about
	   allocation function.
	*/
	Allocate AllocateCallback
	/*
	   The memory reallocation function.  See [ReallocateCallback] for details about
	   reallocation function.
	*/
	Reallocate ReallocateCallback
	/*
	   The memory deallocation function.  See [DeallocateCallback] for details about
	   deallocation function.
	*/
	Deallocate DeallocateCallback
	/*
	   The user pointer for this custom allocator.  This value will be passed to the
	   allocator functions.
	*/
	User unsafe.Pointer
}
