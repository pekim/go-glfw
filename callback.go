// This is a generated file. DO NOT EDIT.

package glfw

import (
	"unsafe"

	ffi "github.com/go-webgpu/goffi/ffi"
)

/*
Generic function pointer used for returning client API function pointers
without forcing a cast from a regular pointer.

C documentation : [GLFWglproc]

# since

Added in version 3.0.

[GLFWglproc]: https://www.glfw.org/docs/latest/group__context.html#ga3d47c2d2fbe0be9c505d0e04e91a133c
*/
type GlProc uintptr

func GlProcNew(callback func()) GlProc {
	return GlProc(ffi.NewCallback(func() {
		callback()
	}))
}

/*
Generic function pointer used for returning Vulkan API function pointers
without forcing a cast from a regular pointer.

C documentation : [GLFWvkproc]

# since

Added in version 3.2.

[GLFWvkproc]: https://www.glfw.org/docs/latest/group__vulkan.html#ga70c01918dc9d233a4fbe0681a43018af
*/
type VkProc uintptr

func VkProcNew(callback func()) VkProc {
	return VkProc(ffi.NewCallback(func() {
		callback()
	}))
}

/*
This is the function pointer type for memory allocation callbacks.  A memory
allocation callback function has the following signature:

	void* function_name(size_t size, void* user)

This function must return either a memory block at least `size` bytes long,
or `NULL` if allocation failed.  Note that not all parts of GLFW handle allocation
failures gracefully yet.

This function must support being called during [Init] but before the library is
flagged as initialized, as well as during [Terminate] after the library is no
longer flagged as initialized.

Any memory allocated via this function will be deallocated via the same allocator
during library termination or earlier.

Any memory allocated via this function must be suitably aligned for any object type.
If you are using C99 or earlier, this alignment is platform-dependent but will be the
same as what `malloc` provides.  If you are using C11 or later, this is the value of
`alignof(max_align_t)`.

The size will always be greater than zero.  Allocations of size zero are filtered out
before reaching the custom allocator.

If this function returns `NULL`, GLFW will emit [OUT_OF_MEMORY].

This function must not call any GLFW function.

C documentation : [GLFWallocatefun]

# params
  - size - The minimum size, in bytes, of the memory block.
  - user - The user-defined pointer from the allocator.

# return

The address of the newly allocated memory block, or `NULL` if an
error occurred.

# pointer lifetime

The returned memory block must be valid at least until it
is deallocated.

# reentrancy

This function should not call any GLFW function.

# thread safety

This function must support being called from any thread that calls GLFW
functions.

# since

Added in version 3.4.

[GLFWallocatefun]: https://www.glfw.org/docs/latest/group__init.html#ga4306a564e9f60f4de8cc8f31731a3120
*/
type AllocateCallback uintptr

func AllocateCallbackNew(callback func(size ULong, user unsafe.Pointer) unsafe.Pointer) AllocateCallback {
	return AllocateCallback(ffi.NewCallback(func(c_size uint64, c_user unsafe.Pointer) {
		callback(ULong(c_size), c_user)
	}))
}

/*
This is the function pointer type for memory reallocation callbacks.
A memory reallocation callback function has the following signature:

	void* function_name(void* block, size_t size, void* user)

This function must return a memory block at least `size` bytes long, or
`NULL` if allocation failed.  Note that not all parts of GLFW handle allocation
failures gracefully yet.

This function must support being called during [Init] but before the library is
flagged as initialized, as well as during [Terminate] after the library is no
longer flagged as initialized.

Any memory allocated via this function will be deallocated via the same allocator
during library termination or earlier.

Any memory allocated via this function must be suitably aligned for any object type.
If you are using C99 or earlier, this alignment is platform-dependent but will be the
same as what `realloc` provides.  If you are using C11 or later, this is the value of
`alignof(max_align_t)`.

The block address will never be `NULL` and the size will always be greater than zero.
Reallocations of a block to size zero are converted into deallocations before reaching
the custom allocator.  Reallocations of `NULL` to a non-zero size are converted into
regular allocations before reaching the custom allocator.

If this function returns `NULL`, GLFW will emit [OUT_OF_MEMORY].

This function must not call any GLFW function.

C documentation : [GLFWreallocatefun]

# params
  - block - The address of the memory block to reallocate.
  - size - The new minimum size, in bytes, of the memory block.
  - user - The user-defined pointer from the allocator.

# return

The address of the newly allocated or resized memory block, or
`NULL` if an error occurred.

# pointer lifetime

The returned memory block must be valid at least until it
is deallocated.

# reentrancy

This function should not call any GLFW function.

# thread safety

This function must support being called from any thread that calls GLFW
functions.

# since

Added in version 3.4.

[GLFWreallocatefun]: https://www.glfw.org/docs/latest/group__init.html#ga3e88a829615d8efe8bec1746f7309c63
*/
type ReallocateCallback uintptr

func ReallocateCallbackNew(callback func(block unsafe.Pointer, size ULong, user unsafe.Pointer) unsafe.Pointer) ReallocateCallback {
	return ReallocateCallback(ffi.NewCallback(func(c_block unsafe.Pointer, c_size uint64, c_user unsafe.Pointer) {
		callback(c_block, ULong(c_size), c_user)
	}))
}

/*
This is the function pointer type for memory deallocation callbacks.
A memory deallocation callback function has the following signature:

	void function_name(void* block, void* user)

This function may deallocate the specified memory block.  This memory block
will have been allocated with the same allocator.

This function must support being called during [Init] but before the library is
flagged as initialized, as well as during [Terminate] after the library is no
longer flagged as initialized.

The block address will never be `NULL`.  Deallocations of `NULL` are filtered out
before reaching the custom allocator.

If this function returns `NULL`, GLFW will emit [OUT_OF_MEMORY].

This function must not call any GLFW function.

C documentation : [GLFWdeallocatefun]

# params
  - block - The address of the memory block to deallocate.
  - user - The user-defined pointer from the allocator.

# pointer lifetime

The specified memory block will not be accessed by GLFW
after this function is called.

# reentrancy

This function should not call any GLFW function.

# thread safety

This function must support being called from any thread that calls GLFW
functions.

# since

Added in version 3.4.

[GLFWdeallocatefun]: https://www.glfw.org/docs/latest/group__init.html#ga7181615eda94c4b07bd72bdcee39fa28
*/
type DeallocateCallback uintptr

func DeallocateCallbackNew(callback func(block unsafe.Pointer, user unsafe.Pointer)) DeallocateCallback {
	return DeallocateCallback(ffi.NewCallback(func(c_block unsafe.Pointer, c_user unsafe.Pointer) {
		callback(c_block, c_user)
	}))
}

/*
This is the function pointer type for error callbacks.  An error callback
function has the following signature:

	void callback_name(int error_code, const char* description)

C documentation : [GLFWerrorfun]

# params
  - error_code - An [error code].  Future releases may add
    more error codes.
  - description - A UTF-8 encoded string describing the error.

# pointer lifetime

The error description string is valid until the callback
function returns.

# since

Added in version 3.0.

[error code]: https://www.glfw.org/docs/latest/group__errors.html
[GLFWerrorfun]: https://www.glfw.org/docs/latest/group__init.html#ga8184701785c096b3862a75cda1bf44a3
*/
type ErrorCallback uintptr

func ErrorCallbackNew(callback func(error_code Int, description string)) ErrorCallback {
	return ErrorCallback(ffi.NewCallback(func(c_error_code int32, c_description *byte) {
		callback(Int(c_error_code), goString(c_description))
	}))
}

/*
This is the function pointer type for window position callbacks.  A window
position callback function has the following signature:

	void callback_name(GLFWwindow* window, int xpos, int ypos)

C documentation : [GLFWwindowposfun]

# params
  - window - The window that was moved.
  - xpos - The new x-coordinate, in screen coordinates, of the
    upper-left corner of the content area of the window.
  - ypos - The new y-coordinate, in screen coordinates, of the
    upper-left corner of the content area of the window.

# since

Added in version 3.0.

[GLFWwindowposfun]: https://www.glfw.org/docs/latest/group__window.html#gabe287973a21a8f927cde4db06b8dcbe9
*/
type WindowposCallback uintptr

func WindowposCallbackNew(callback func(window *Window, xpos Int, ypos Int)) WindowposCallback {
	return WindowposCallback(ffi.NewCallback(func(c_window *Window, c_xpos int32, c_ypos int32) {
		callback(c_window, Int(c_xpos), Int(c_ypos))
	}))
}

/*
This is the function pointer type for window size callbacks.  A window size
callback function has the following signature:

	void callback_name(GLFWwindow* window, int width, int height)

C documentation : [GLFWwindowsizefun]

# params
  - window - The window that was resized.
  - width - The new width, in screen coordinates, of the window.
  - height - The new height, in screen coordinates, of the window.

# since

Added in version 1.0.

[GLFWwindowsizefun]: https://www.glfw.org/docs/latest/group__window.html#gaec0282944bb810f6f3163ec02da90350
*/
type WindowsizeCallback uintptr

func WindowsizeCallbackNew(callback func(window *Window, width Int, height Int)) WindowsizeCallback {
	return WindowsizeCallback(ffi.NewCallback(func(c_window *Window, c_width int32, c_height int32) {
		callback(c_window, Int(c_width), Int(c_height))
	}))
}

/*
This is the function pointer type for window close callbacks.  A window
close callback function has the following signature:

	void function_name(GLFWwindow* window)

C documentation : [GLFWwindowclosefun]

# params
  - window - The window that the user attempted to close.

# since

Added in version 2.5.

[GLFWwindowclosefun]: https://www.glfw.org/docs/latest/group__window.html#gabf859b936d80961b7d39013a9694cc3e
*/
type WindowcloseCallback uintptr

func WindowcloseCallbackNew(callback func(window *Window)) WindowcloseCallback {
	return WindowcloseCallback(ffi.NewCallback(func(c_window *Window) {
		callback(c_window)
	}))
}

/*
This is the function pointer type for window content refresh callbacks.
A window content refresh callback function has the following signature:

	void function_name(GLFWwindow* window);

C documentation : [GLFWwindowrefreshfun]

# params
  - window - The window whose content needs to be refreshed.

# since

Added in version 2.5.

[GLFWwindowrefreshfun]: https://www.glfw.org/docs/latest/group__window.html#ga431663a1427d2eb3a273bc398b6737b5
*/
type WindowrefreshCallback uintptr

func WindowrefreshCallbackNew(callback func(window *Window)) WindowrefreshCallback {
	return WindowrefreshCallback(ffi.NewCallback(func(c_window *Window) {
		callback(c_window)
	}))
}

/*
This is the function pointer type for window focus callbacks.  A window
focus callback function has the following signature:

	void function_name(GLFWwindow* window, int focused)

C documentation : [GLFWwindowfocusfun]

# params
  - window - The window that gained or lost input focus.
  - focused - `GLFW_TRUE` if the window was given input focus, or
    `GLFW_FALSE` if it lost it.

# since

Added in version 3.0.

[GLFWwindowfocusfun]: https://www.glfw.org/docs/latest/group__window.html#gabc58c47e9d93f6eb1862d615c3680f46
*/
type WindowfocusCallback uintptr

func WindowfocusCallbackNew(callback func(window *Window, focused Int)) WindowfocusCallback {
	return WindowfocusCallback(ffi.NewCallback(func(c_window *Window, c_focused int32) {
		callback(c_window, Int(c_focused))
	}))
}

/*
This is the function pointer type for window iconify callbacks.  A window
iconify callback function has the following signature:

	void function_name(GLFWwindow* window, int iconified)

C documentation : [GLFWwindowiconifyfun]

# params
  - window - The window that was iconified or restored.
  - iconified - `GLFW_TRUE` if the window was iconified, or
    `GLFW_FALSE` if it was restored.

# since

Added in version 3.0.

[GLFWwindowiconifyfun]: https://www.glfw.org/docs/latest/group__window.html#ga35c658cccba236f26e7adee0e25f6a4f
*/
type WindowiconifyCallback uintptr

func WindowiconifyCallbackNew(callback func(window *Window, iconified Int)) WindowiconifyCallback {
	return WindowiconifyCallback(ffi.NewCallback(func(c_window *Window, c_iconified int32) {
		callback(c_window, Int(c_iconified))
	}))
}

/*
This is the function pointer type for window maximize callbacks.  A window
maximize callback function has the following signature:

	void function_name(GLFWwindow* window, int maximized)

C documentation : [GLFWwindowmaximizefun]

# params
  - window - The window that was maximized or restored.
  - maximized - `GLFW_TRUE` if the window was maximized, or
    `GLFW_FALSE` if it was restored.

# since

Added in version 3.3.

[GLFWwindowmaximizefun]: https://www.glfw.org/docs/latest/group__window.html#ga3017196fdaec33ac3e095765176c2a90
*/
type WindowmaximizeCallback uintptr

func WindowmaximizeCallbackNew(callback func(window *Window, maximized Int)) WindowmaximizeCallback {
	return WindowmaximizeCallback(ffi.NewCallback(func(c_window *Window, c_maximized int32) {
		callback(c_window, Int(c_maximized))
	}))
}

/*
This is the function pointer type for framebuffer size callbacks.
A framebuffer size callback function has the following signature:

	void function_name(GLFWwindow* window, int width, int height)

C documentation : [GLFWframebuffersizefun]

# params
  - window - The window whose framebuffer was resized.
  - width - The new width, in pixels, of the framebuffer.
  - height - The new height, in pixels, of the framebuffer.

# since

Added in version 3.0.

[GLFWframebuffersizefun]: https://www.glfw.org/docs/latest/group__window.html#gae18026e294dde685ed2e5f759533144d
*/
type FramebuffersizeCallback uintptr

func FramebuffersizeCallbackNew(callback func(window *Window, width Int, height Int)) FramebuffersizeCallback {
	return FramebuffersizeCallback(ffi.NewCallback(func(c_window *Window, c_width int32, c_height int32) {
		callback(c_window, Int(c_width), Int(c_height))
	}))
}

/*
This is the function pointer type for window content scale callbacks.
A window content scale callback function has the following signature:

	void function_name(GLFWwindow* window, float xscale, float yscale)

C documentation : [GLFWwindowcontentscalefun]

# params
  - window - The window whose content scale changed.
  - xscale - The new x-axis content scale of the window.
  - yscale - The new y-axis content scale of the window.

# since

Added in version 3.3.

[GLFWwindowcontentscalefun]: https://www.glfw.org/docs/latest/group__window.html#ga77f288a2d04bb3c77c7d9615d08cf70e
*/
type WindowcontentscaleCallback uintptr

func WindowcontentscaleCallbackNew(callback func(window *Window, xscale Float, yscale Float)) WindowcontentscaleCallback {
	return WindowcontentscaleCallback(ffi.NewCallback(func(c_window *Window, c_xscale float32, c_yscale float32) {
		callback(c_window, Float(c_xscale), Float(c_yscale))
	}))
}

/*
This is the function pointer type for mouse button callback functions.
A mouse button callback function has the following signature:

	void function_name(GLFWwindow* window, int button, int action, int mods)

C documentation : [GLFWmousebuttonfun]

# params
  - window - The window that received the event.
  - button - The [mouse button] that was pressed or
    released.
  - action - One of `GLFW_PRESS` or `GLFW_RELEASE`.  Future releases
    may add more actions.
  - mods - Bit field describing which [modifier keys] were
    held down.

# since

Added in version 1.0.

[mouse button]: https://www.glfw.org/docs/latest/group__buttons.html
[modifier keys]: https://www.glfw.org/docs/latest/group__mods.html
[GLFWmousebuttonfun]: https://www.glfw.org/docs/latest/group__input.html#ga0184dcb59f6d85d735503dcaae809727
*/
type MousebuttonCallback uintptr

func MousebuttonCallbackNew(callback func(window *Window, button Int, action Int, mods Int)) MousebuttonCallback {
	return MousebuttonCallback(ffi.NewCallback(func(c_window *Window, c_button int32, c_action int32, c_mods int32) {
		callback(c_window, Int(c_button), Int(c_action), Int(c_mods))
	}))
}

/*
This is the function pointer type for cursor position callbacks.  A cursor
position callback function has the following signature:

	void function_name(GLFWwindow* window, double xpos, double ypos);

C documentation : [GLFWcursorposfun]

# params
  - window - The window that received the event.
  - xpos - The new cursor x-coordinate, relative to the left edge of
    the content area.
  - ypos - The new cursor y-coordinate, relative to the top edge of the
    content area.

# since

Added in version 3.0.  Replaces `GLFWmouseposfun`.

[GLFWcursorposfun]: https://www.glfw.org/docs/latest/group__input.html#gad6fae41b3ac2e4209aaa87b596c57f68
*/
type CursorposCallback uintptr

func CursorposCallbackNew(callback func(window *Window, xpos Double, ypos Double)) CursorposCallback {
	return CursorposCallback(ffi.NewCallback(func(c_window *Window, c_xpos float64, c_ypos float64) {
		callback(c_window, Double(c_xpos), Double(c_ypos))
	}))
}

/*
This is the function pointer type for cursor enter/leave callbacks.
A cursor enter/leave callback function has the following signature:

	void function_name(GLFWwindow* window, int entered)

C documentation : [GLFWcursorenterfun]

# params
  - window - The window that received the event.
  - entered - `GLFW_TRUE` if the cursor entered the window's content
    area, or `GLFW_FALSE` if it left it.

# since

Added in version 3.0.

[GLFWcursorenterfun]: https://www.glfw.org/docs/latest/group__input.html#gaa93dc4818ac9ab32532909d53a337cbe
*/
type CursorenterCallback uintptr

func CursorenterCallbackNew(callback func(window *Window, entered Int)) CursorenterCallback {
	return CursorenterCallback(ffi.NewCallback(func(c_window *Window, c_entered int32) {
		callback(c_window, Int(c_entered))
	}))
}

/*
This is the function pointer type for scroll callbacks.  A scroll callback
function has the following signature:

	void function_name(GLFWwindow* window, double xoffset, double yoffset)

C documentation : [GLFWscrollfun]

# params
  - window - The window that received the event.
  - xoffset - The scroll offset along the x-axis.
  - yoffset - The scroll offset along the y-axis.

# since

Added in version 3.0.  Replaces `GLFWmousewheelfun`.

[GLFWscrollfun]: https://www.glfw.org/docs/latest/group__input.html#gaf656112c33de3efdb227fa58f0134cf5
*/
type ScrollCallback uintptr

func ScrollCallbackNew(callback func(window *Window, xoffset Double, yoffset Double)) ScrollCallback {
	return ScrollCallback(ffi.NewCallback(func(c_window *Window, c_xoffset float64, c_yoffset float64) {
		callback(c_window, Double(c_xoffset), Double(c_yoffset))
	}))
}

/*
This is the function pointer type for keyboard key callbacks.  A keyboard
key callback function has the following signature:

	void function_name(GLFWwindow* window, int key, int scancode, int action, int mods)

C documentation : [GLFWkeyfun]

# params
  - window - The window that received the event.
  - key - The [keyboard key] that was pressed or released.
  - scancode - The platform-specific scancode of the key.
  - action - `GLFW_PRESS`, `GLFW_RELEASE` or `GLFW_REPEAT`.  Future
    releases may add more actions.
  - mods - Bit field describing which [modifier keys] were
    held down.

# since

Added in version 1.0.

[keyboard key]: https://www.glfw.org/docs/latest/group__keys.html
[modifier keys]: https://www.glfw.org/docs/latest/group__mods.html
[GLFWkeyfun]: https://www.glfw.org/docs/latest/group__input.html#ga5bd751b27b90f865d2ea613533f0453c
*/
type KeyCallback uintptr

func KeyCallbackNew(callback func(window *Window, key Int, scancode Int, action Int, mods Int)) KeyCallback {
	return KeyCallback(ffi.NewCallback(func(c_window *Window, c_key int32, c_scancode int32, c_action int32, c_mods int32) {
		callback(c_window, Int(c_key), Int(c_scancode), Int(c_action), Int(c_mods))
	}))
}

/*
This is the function pointer type for Unicode character callbacks.
A Unicode character callback function has the following signature:

	void function_name(GLFWwindow* window, unsigned int codepoint)

C documentation : [GLFWcharfun]

# params
  - window - The window that received the event.
  - codepoint - The Unicode code point of the character.

# since

Added in version 2.4.

[GLFWcharfun]: https://www.glfw.org/docs/latest/group__input.html#ga1ab90a55cf3f58639b893c0f4118cb6e
*/
type CharCallback uintptr

func CharCallbackNew(callback func(window *Window, codepoint UInt)) CharCallback {
	return CharCallback(ffi.NewCallback(func(c_window *Window, c_codepoint uint32) {
		callback(c_window, UInt(c_codepoint))
	}))
}

/*
callbacks.

This is the function pointer type for Unicode character with modifiers
callbacks.  It is called for each input character, regardless of what
modifier keys are held down.  A Unicode character with modifiers callback
function has the following signature:

	void function_name(GLFWwindow* window, unsigned int codepoint, int mods)

C documentation : [GLFWcharmodsfun]

# params
  - window - The window that received the event.
  - codepoint - The Unicode code point of the character.
  - mods - Bit field describing which [modifier keys] were
    held down.

# deprecated

Scheduled for removal in version 4.0.

# since

Added in version 3.1.

[modifier keys]: https://www.glfw.org/docs/latest/group__mods.html
[GLFWcharmodsfun]: https://www.glfw.org/docs/latest/group__input.html#gac3cf64f90b6219c05ac7b7822d5a4b8f
*/
type CharmodsCallback uintptr

func CharmodsCallbackNew(callback func(window *Window, codepoint UInt, mods Int)) CharmodsCallback {
	return CharmodsCallback(ffi.NewCallback(func(c_window *Window, c_codepoint uint32, c_mods int32) {
		callback(c_window, UInt(c_codepoint), Int(c_mods))
	}))
}

/*
This is the function pointer type for path drop callbacks.  A path drop
callback function has the following signature:

	void function_name(GLFWwindow* window, int path_count, const char* paths[])

C documentation : [GLFWdropfun]

# params
  - window - The window that received the event.
  - path_count - The number of dropped paths.
  - paths - The UTF-8 encoded file and/or directory path names.

# pointer lifetime

The path array and its strings are valid until the
callback function returns.

# since

Added in version 3.1.

[GLFWdropfun]: https://www.glfw.org/docs/latest/group__input.html#gaaba73c3274062c18723b7f05862d94b2
*/
type DropCallback uintptr

// UNSUPPORTED GLFWdropfun : param paths is "const char *[]"

/*
This is the function pointer type for monitor configuration callbacks.
A monitor callback function has the following signature:

	void function_name(GLFWmonitor* monitor, int event)

C documentation : [GLFWmonitorfun]

# params
  - monitor - The monitor that was connected or disconnected.
  - event - One of `GLFW_CONNECTED` or `GLFW_DISCONNECTED`.  Future
    releases may add more events.

# since

Added in version 3.0.

[GLFWmonitorfun]: https://www.glfw.org/docs/latest/group__monitor.html#gaabe16caca8dea952504dfdebdf4cd249
*/
type MonitorCallback uintptr

func MonitorCallbackNew(callback func(monitor *Monitor, event Int)) MonitorCallback {
	return MonitorCallback(ffi.NewCallback(func(c_monitor *Monitor, c_event int32) {
		callback(c_monitor, Int(c_event))
	}))
}

/*
This is the function pointer type for joystick configuration callbacks.
A joystick configuration callback function has the following signature:

	void function_name(int jid, int event)

C documentation : [GLFWjoystickfun]

# params
  - jid - The joystick that was connected or disconnected.
  - event - One of `GLFW_CONNECTED` or `GLFW_DISCONNECTED`.  Future
    releases may add more events.

# since

Added in version 3.2.

[GLFWjoystickfun]: https://www.glfw.org/docs/latest/group__input.html#gaa21ad5986ae9a26077a40142efb56243
*/
type JoystickCallback uintptr

func JoystickCallbackNew(callback func(jid Int, event Int)) JoystickCallback {
	return JoystickCallback(ffi.NewCallback(func(c_jid int32, c_event int32) {
		callback(Int(c_jid), Int(c_event))
	}))
}
