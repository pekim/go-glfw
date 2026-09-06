// This is a generated file. DO NOT EDIT.

package glfw

import (
	"unsafe"

	ffi "github.com/go-webgpu/goffi/ffi"
)

/*
This function initializes the GLFW library.  Before most GLFW functions can
be used, GLFW must be initialized, and before an application terminates GLFW
should be terminated in order to free any resources allocated during or
after initialization.

If this function fails, it calls [Terminate] before returning.  If it
succeeds, you should call [Terminate] before the application exits.

Additional calls to this function after successful initialization but before
termination will return `GLFW_TRUE` immediately.

The [PLATFORM] init hint controls which platforms are considered during
initialization.  This also depends on which platforms the library was compiled to
support.

C documentation : [glfwInit]

# return

`GLFW_TRUE` if successful, or `GLFW_FALSE` if an
[error] occurred.

# errors

Possible errors include [PLATFORM_UNAVAILABLE] and [PLATFORM_ERROR].

# macos

This function will change the current directory of the
application to the `Contents/Resources` subdirectory of the application's
bundle, if present.  This can be disabled with the [COCOA_CHDIR_RESOURCES] init hint.

# macos

This function will create the main menu and dock icon for the
application.  If GLFW finds a `MainMenu.nib` it is loaded and assumed to
contain a menu bar.  Otherwise a minimal menu bar is created manually with
common commands like Hide, Quit and About.  The About entry opens a minimal
about dialog with information from the application's bundle.  The menu bar
and dock icon can be disabled entirely with the [COCOA_MENUBAR] init
hint.

__Wayland, X11:__ If the library was compiled with support for both
Wayland and X11, and the [PLATFORM] init hint is set to
`GLFW_ANY_PLATFORM`, the `XDG_SESSION_TYPE` environment variable affects
which platform is picked.  If the environment variable is not set, or is set
to something other than `wayland` or `x11`, the regular detection mechanism
will be used instead.

# x11

This function will set the `LC_CTYPE` category of the
application locale according to the current environment if that category is
still "C".  This is because the "C" locale breaks Unicode text input.

# thread safety

This function must only be called from the main thread.

# since

Added in version 1.0.

[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwInit]: https://www.glfw.org/docs/latest/group__init.html#ga317aac130a235ab08c6db0834907d85e
*/
func Init() Int {
	var result Int
	_, err := ffi.CallFunction(
		cif_glfwInit,
		func_glfwInit,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function destroys all remaining windows and cursors, restores any
modified gamma ramps and frees any other allocated resources.  Once this
function is called, you must again call [Init] successfully before
you will be able to use most GLFW functions.

If GLFW has been successfully initialized, this function should be called
before the application exits.  If initialization fails, there is no need to
call this function, as it is called by [Init] before it returns
failure.

This function has no effect if GLFW is not initialized.

C documentation : [glfwTerminate]

# errors

Possible errors include [PLATFORM_ERROR].

This function may be called before [Init].

The contexts of any remaining windows must not be current on any
other thread when this function is called.

# reentrancy

This function must not be called from a callback.

# thread safety

This function must only be called from the main thread.

# since

Added in version 1.0.

[glfwTerminate]: https://www.glfw.org/docs/latest/group__init.html#gaaae48c0a18607ea4a4ba951d939f0901
*/
func Terminate() {
	_, err := ffi.CallFunction(
		cif_glfwTerminate,
		func_glfwTerminate,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function sets hints for the next initialization of GLFW.

The values you set hints to are never reset by GLFW, but they only take
effect during initialization.  Once GLFW has been initialized, any values
you set will be ignored until the library is terminated and initialized
again.

Some hints are platform specific.  These may be set on any platform but they
will only affect their specific platform.  Other platforms will ignore them.
Setting these hints requires no platform specific headers or functions.

C documentation : [glfwInitHint]

# params
  - hint - The [init hint] to set.
  - value - The new value of the init hint.

# errors

Possible errors include [INVALID_ENUM] and [INVALID_VALUE].

This function may be called before [Init].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.3.

[init hint]: https://www.glfw.org/docs/latest/intro_guide.html#init_hints
[glfwInitHint]: https://www.glfw.org/docs/latest/group__init.html#ga110fd1d3f0412822b4f1908c026f724a
*/
func InitHint(hint Int, value Int) {
	_, err := ffi.CallFunction(
		cif_glfwInitHint,
		func_glfwInitHint,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&hint),
			unsafe.Pointer(&value),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
To use the default allocator, call this function with a `NULL` argument.

If you specify an allocator struct, every member must be a valid function
pointer.  If any member is `NULL`, this function will emit [INVALID_VALUE] and the init allocator will be unchanged.

The functions in the allocator must fulfil a number of requirements.  See the
documentation for [AllocateCallback], [ReallocateCallback] and [DeallocateCallback] for details.

C documentation : [glfwInitAllocator]

# params
  - allocator - The allocator to use at the next initialization, or
    `NULL` to use the default one.

# errors

Possible errors include [INVALID_VALUE].

# pointer lifetime

The specified allocator is copied before this function
returns.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.4.

[glfwInitAllocator]: https://www.glfw.org/docs/latest/group__init.html#ga9dde93e9891fa7dd17e4194c9f3ae7c6
*/
func (allocator *Allocator) Init() {
	_, err := ffi.CallFunction(
		cif_glfwInitAllocator,
		func_glfwInitAllocator,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&allocator),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function retrieves the major, minor and revision numbers of the GLFW
library.  It is intended for when you are using GLFW as a shared library and
want to ensure that you are using the minimum required version.

Any or all of the version arguments may be `NULL`.

C documentation : [glfwGetVersion]

# params
  - major (out) - Where to store the major version number, or `NULL`.
  - minor (out) - Where to store the minor version number, or `NULL`.
  - rev (out) - Where to store the revision number, or `NULL`.

# errors

None.

This function may be called before [Init].

# thread safety

This function may be called from any thread.

# since

Added in version 1.0.

[glfwGetVersion]: https://www.glfw.org/docs/latest/group__init.html#ga9f8ffaacf3c269cc48eafbf8b9b71197
*/
func GetVersion() (Int, Int, Int) {
	var major Int
	var minor Int
	var rev Int
	_, err := ffi.CallFunction(
		cif_glfwGetVersion,
		func_glfwGetVersion,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(new(&major)),
			unsafe.Pointer(new(&minor)),
			unsafe.Pointer(new(&rev)),
		},
	)
	if err != nil {
		panic(err)
	}
	return major, minor, rev
}

/*
This function returns the compile-time generated
[version string] of the GLFW library binary.  It describes
the version, platforms, compiler and any platform or operating system specific
compile-time options.  It should not be confused with the OpenGL or OpenGL ES version
string, queried with `glGetString`.

__Do not use the version string__ to parse the GLFW library version.  The
[GetVersion] function provides the version of the running library
binary in numerical format.

__Do not use the version string__ to parse what platforms are supported.  The [PlatformSupported] function lets you query platform support.

C documentation : [glfwGetVersionString]

# return

The ASCII encoded GLFW version string.

# errors

None.

This function may be called before [Init].

# pointer lifetime

The returned string is static and compile-time generated.

# thread safety

This function may be called from any thread.

# since

Added in version 3.0.

[version string]: https://www.glfw.org/docs/latest/intro_guide.html#intro_version_string
[glfwGetVersionString]: https://www.glfw.org/docs/latest/group__init.html#ga026abd003c8e6501981ab1662062f1c0
*/
func GetVersionString() string {
	var result *byte
	_, err := ffi.CallFunction(
		cif_glfwGetVersionString,
		func_glfwGetVersionString,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{},
	)
	if err != nil {
		panic(err)
	}
	return goString(result)
}

// UNSUPPORTED glfwGetError : param description is "const char **"

/*
This function sets the error callback, which is called with an error code
and a human-readable description each time a GLFW error occurs.

The error code is set before the callback is called.  Calling [GetError] from the error callback will return the same value as the error
code argument.

The error callback is called on the thread where the error occurred.  If you
are using GLFW from multiple threads, your error callback needs to be
written accordingly.

Because the description string may have been generated specifically for that
error, it is not guaranteed to be valid after the callback has returned.  If
you wish to use it after the callback returns, you need to make a copy.

Once set, the error callback remains set even after the library has been
terminated.

C documentation : [glfwSetErrorCallback]

# params
  - callback - The new callback, or `NULL` to remove the currently set
    callback.

# return

The previously set callback, or `NULL` if no callback was set.

	void callback_name(int error_code, const char* description)

For more information about the callback parameters, see the
[callback pointer type].

# errors

None.

This function may be called before [Init].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[callback pointer type]: https://www.glfw.org/docs/latest/group__init.html#ga8184701785c096b3862a75cda1bf44a3
[glfwSetErrorCallback]: https://www.glfw.org/docs/latest/group__init.html#gaff45816610d53f0b83656092a4034f40
*/
func SetErrorCallback(callback ErrorCallback) ErrorCallback {
	var result ErrorCallback
	_, err := ffi.CallFunction(
		cif_glfwSetErrorCallback,
		func_glfwSetErrorCallback,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&callback),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function returns the platform that was selected during initialization.  The
returned value will be one of `GLFW_PLATFORM_WIN32`, `GLFW_PLATFORM_COCOA`,
`GLFW_PLATFORM_WAYLAND`, `GLFW_PLATFORM_X11` or `GLFW_PLATFORM_NULL`.

C documentation : [glfwGetPlatform]

# return

The currently selected platform, or zero if an error occurred.

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function may be called from any thread.

# since

Added in version 3.4.

[glfwGetPlatform]: https://www.glfw.org/docs/latest/group__init.html#ga6d6a983d38bd4e8fd786d7a9061d399e
*/
func GetPlatform() Int {
	var result Int
	_, err := ffi.CallFunction(
		cif_glfwGetPlatform,
		func_glfwGetPlatform,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function returns whether the library was compiled with support for the specified
platform.  The platform must be one of `GLFW_PLATFORM_WIN32`, `GLFW_PLATFORM_COCOA`,
`GLFW_PLATFORM_WAYLAND`, `GLFW_PLATFORM_X11` or `GLFW_PLATFORM_NULL`.

C documentation : [glfwPlatformSupported]

# params
  - platform - The platform to query.

# return

`GLFW_TRUE` if the platform is supported, or `GLFW_FALSE` otherwise.

# errors

Possible errors include [INVALID_ENUM].

This function may be called before [Init].

# thread safety

This function may be called from any thread.

# since

Added in version 3.4.

[glfwPlatformSupported]: https://www.glfw.org/docs/latest/group__init.html#ga8785d2b6b36632368d803e78079d38ed
*/
func PlatformSupported(platform Int) Int {
	var result Int
	_, err := ffi.CallFunction(
		cif_glfwPlatformSupported,
		func_glfwPlatformSupported,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&platform),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

// UNSUPPORTED glfwGetMonitors : result type is "GLFWmonitor **"

/*
This function returns the primary monitor.  This is usually the monitor
where elements like the task bar or global menu bar are located.

C documentation : [glfwGetPrimaryMonitor]

# return

The primary monitor, or `NULL` if no monitors were found or if an
[error] occurred.

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function must only be called from the main thread.

The primary monitor is always first in the array returned by [GetMonitors].

# since

Added in version 3.0.

[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwGetPrimaryMonitor]: https://www.glfw.org/docs/latest/group__monitor.html#gac3adb24947eb709e1874028272e5dfc5
*/
func GetPrimaryMonitor() *Monitor {
	var result *Monitor
	_, err := ffi.CallFunction(
		cif_glfwGetPrimaryMonitor,
		func_glfwGetPrimaryMonitor,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function returns the position, in screen coordinates, of the upper-left
corner of the specified monitor.

Any or all of the position arguments may be `NULL`.  If an error occurs, all
non-`NULL` position arguments will be set to zero.

C documentation : [glfwGetMonitorPos]

# params
  - monitor - The monitor to query.
  - xpos (out) - Where to store the monitor x-coordinate, or `NULL`.
  - ypos (out) - Where to store the monitor y-coordinate, or `NULL`.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[glfwGetMonitorPos]: https://www.glfw.org/docs/latest/group__monitor.html#ga102f54e7acc9149edbcf0997152df8c9
*/
func (monitor *Monitor) GetPos() (Int, Int) {
	var xpos Int
	var ypos Int
	_, err := ffi.CallFunction(
		cif_glfwGetMonitorPos,
		func_glfwGetMonitorPos,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&monitor),
			unsafe.Pointer(new(&xpos)),
			unsafe.Pointer(new(&ypos)),
		},
	)
	if err != nil {
		panic(err)
	}
	return xpos, ypos
}

/*
This function returns the position, in screen coordinates, of the upper-left
corner of the work area of the specified monitor along with the work area
size in screen coordinates. The work area is defined as the area of the
monitor not occluded by the window system task bar where present. If no
task bar exists then the work area is the monitor resolution in screen
coordinates.

Any or all of the position and size arguments may be `NULL`.  If an error
occurs, all non-`NULL` position and size arguments will be set to zero.

C documentation : [glfwGetMonitorWorkarea]

# params
  - monitor - The monitor to query.
  - xpos (out) - Where to store the monitor x-coordinate, or `NULL`.
  - ypos (out) - Where to store the monitor y-coordinate, or `NULL`.
  - width (out) - Where to store the monitor width, or `NULL`.
  - height (out) - Where to store the monitor height, or `NULL`.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.3.

[glfwGetMonitorWorkarea]: https://www.glfw.org/docs/latest/group__monitor.html#ga7387a3bdb64bfe8ebf2b9e54f5b6c9d0
*/
func (monitor *Monitor) GetWorkarea() (Int, Int, Int, Int) {
	var xpos Int
	var ypos Int
	var width Int
	var height Int
	_, err := ffi.CallFunction(
		cif_glfwGetMonitorWorkarea,
		func_glfwGetMonitorWorkarea,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&monitor),
			unsafe.Pointer(new(&xpos)),
			unsafe.Pointer(new(&ypos)),
			unsafe.Pointer(new(&width)),
			unsafe.Pointer(new(&height)),
		},
	)
	if err != nil {
		panic(err)
	}
	return xpos, ypos, width, height
}

/*
This function returns the size, in millimetres, of the display area of the
specified monitor.

Some platforms do not provide accurate monitor size information, either
because the monitor [EDID][] data is incorrect or because the driver does
not report it accurately.

Any or all of the size arguments may be `NULL`.  If an error occurs, all
non-`NULL` size arguments will be set to zero.

C documentation : [glfwGetMonitorPhysicalSize]

# params
  - monitor - The monitor to query.
  - widthMM (out) - Where to store the width, in millimetres, of the
    monitor's display area, or `NULL`.
  - heightMM (out) - Where to store the height, in millimetres, of the
    monitor's display area, or `NULL`.

# errors

Possible errors include [NOT_INITIALIZED].

# win32

On Windows 8 and earlier the physical size is calculated from
the current resolution and system DPI instead of querying the monitor EDID data.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[EDID]: https://en.wikipedia.org/wiki/Extended_display_identification_data
[glfwGetMonitorPhysicalSize]: https://www.glfw.org/docs/latest/group__monitor.html#ga7d8bffc6c55539286a6bd20d32a8d7ea
*/
func (monitor *Monitor) GetPhysicalSize() (Int, Int) {
	var widthMM Int
	var heightMM Int
	_, err := ffi.CallFunction(
		cif_glfwGetMonitorPhysicalSize,
		func_glfwGetMonitorPhysicalSize,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&monitor),
			unsafe.Pointer(new(&widthMM)),
			unsafe.Pointer(new(&heightMM)),
		},
	)
	if err != nil {
		panic(err)
	}
	return widthMM, heightMM
}

/*
This function retrieves the content scale for the specified monitor.  The
content scale is the ratio between the current DPI and the platform's
default DPI.  This is especially important for text and any UI elements.  If
the pixel dimensions of your UI scaled by this look appropriate on your
machine then it should appear at a reasonable size on other machines
regardless of their DPI and scaling settings.  This relies on the system DPI
and scaling settings being somewhat correct.

The content scale may depend on both the monitor resolution and pixel
density and on user settings.  It may be very different from the raw DPI
calculated from the physical size and current resolution.

C documentation : [glfwGetMonitorContentScale]

# params
  - monitor - The monitor to query.
  - xscale (out) - Where to store the x-axis content scale, or `NULL`.
  - yscale (out) - Where to store the y-axis content scale, or `NULL`.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# wayland

Fractional scaling information is not yet available for
monitors, so this function only returns integer content scales.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.3.

[glfwGetMonitorContentScale]: https://www.glfw.org/docs/latest/group__monitor.html#gad3152e84465fa620b601265ebfcdb21b
*/
func (monitor *Monitor) GetContentScale() (Float, Float) {
	var xscale Float
	var yscale Float
	_, err := ffi.CallFunction(
		cif_glfwGetMonitorContentScale,
		func_glfwGetMonitorContentScale,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&monitor),
			unsafe.Pointer(new(&xscale)),
			unsafe.Pointer(new(&yscale)),
		},
	)
	if err != nil {
		panic(err)
	}
	return xscale, yscale
}

/*
This function returns a human-readable name, encoded as UTF-8, of the
specified monitor.  The name typically reflects the make and model of the
monitor and is not guaranteed to be unique among the connected monitors.

C documentation : [glfwGetMonitorName]

# params
  - monitor - The monitor to query.

# return

The UTF-8 encoded name of the monitor, or `NULL` if an
[error] occurred.

# errors

Possible errors include [NOT_INITIALIZED].

# pointer lifetime

The returned string is allocated and freed by GLFW.  You
should not free it yourself.  It is valid until the specified monitor is
disconnected or the library is terminated.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwGetMonitorName]: https://www.glfw.org/docs/latest/group__monitor.html#ga7af83e13489d90379588fb331b9e4b68
*/
func (monitor *Monitor) GetName() string {
	var result *byte
	_, err := ffi.CallFunction(
		cif_glfwGetMonitorName,
		func_glfwGetMonitorName,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&monitor),
		},
	)
	if err != nil {
		panic(err)
	}
	return goString(result)
}

/*
This function sets the user-defined pointer of the specified monitor.  The
current value is retained until the monitor is disconnected.  The initial
value is `NULL`.

This function may be called from the monitor callback, even for a monitor
that is being disconnected.

C documentation : [glfwSetMonitorUserPointer]

# params
  - monitor - The monitor whose pointer to set.
  - pointer - The new value.

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function may be called from any thread.  Access is not
synchronized.

# since

Added in version 3.3.

[glfwSetMonitorUserPointer]: https://www.glfw.org/docs/latest/group__monitor.html#ga702750e24313a686d3637297b6e85fda
*/
func (monitor *Monitor) SetUserPointer(pointer unsafe.Pointer) {
	_, err := ffi.CallFunction(
		cif_glfwSetMonitorUserPointer,
		func_glfwSetMonitorUserPointer,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&monitor),
			unsafe.Pointer(&pointer),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function returns the current value of the user-defined pointer of the
specified monitor.  The initial value is `NULL`.

This function may be called from the monitor callback, even for a monitor
that is being disconnected.

C documentation : [glfwGetMonitorUserPointer]

# params
  - monitor - The monitor whose pointer to return.

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function may be called from any thread.  Access is not
synchronized.

# since

Added in version 3.3.

[glfwGetMonitorUserPointer]: https://www.glfw.org/docs/latest/group__monitor.html#ga1adbfbfb8cd58b23cfee82e574fbbdc5
*/
func (monitor *Monitor) GetUserPointer() unsafe.Pointer {
	var result unsafe.Pointer
	_, err := ffi.CallFunction(
		cif_glfwGetMonitorUserPointer,
		func_glfwGetMonitorUserPointer,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&monitor),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the monitor configuration callback, or removes the
currently set callback.  This is called when a monitor is connected to or
disconnected from the system.

C documentation : [glfwSetMonitorCallback]

# params
  - callback - The new callback, or `NULL` to remove the currently set
    callback.

# return

The previously set callback, or `NULL` if no callback was set or the
library had not been [initialized].

	void function_name(GLFWmonitor* monitor, int event)

For more information about the callback parameters, see the
[function pointer type].

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[initialized]: https://www.glfw.org/docs/latest/intro_guide.html#intro_init
[function pointer type]: https://www.glfw.org/docs/latest/group__monitor.html#gaabe16caca8dea952504dfdebdf4cd249
[glfwSetMonitorCallback]: https://www.glfw.org/docs/latest/group__monitor.html#gab39df645587c8518192aa746c2fb06c3
*/
func SetMonitorCallback(callback MonitorCallback) MonitorCallback {
	var result MonitorCallback
	_, err := ffi.CallFunction(
		cif_glfwSetMonitorCallback,
		func_glfwSetMonitorCallback,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&callback),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

// UNSUPPORTED glfwGetVideoModes : result type is "const GLFWvidmode *"

/*
This function returns the current video mode of the specified monitor.  If
you have created a full screen window for that monitor, the return value
will depend on whether that window is iconified.

C documentation : [glfwGetVideoMode]

# params
  - monitor - The monitor to query.

# return

The current mode of the monitor, or `NULL` if an
[error] occurred.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# pointer lifetime

The returned array is allocated and freed by GLFW.  You
should not free it yourself.  It is valid until the specified monitor is
disconnected or the library is terminated.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.  Replaces `glfwGetDesktopMode`.

[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwGetVideoMode]: https://www.glfw.org/docs/latest/group__monitor.html#gaba376fa7e76634b4788bddc505d6c9d5
*/
func (monitor *Monitor) GetVideoMode() *Vidmode {
	var result *Vidmode
	_, err := ffi.CallFunction(
		cif_glfwGetVideoMode,
		func_glfwGetVideoMode,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&monitor),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function generates an appropriately sized gamma ramp from the specified
exponent and then calls [SetGammaRamp] with it.  The value must be
a finite number greater than zero.

The software controlled gamma ramp is applied _in addition_ to the hardware
gamma correction, which today is usually an approximation of sRGB gamma.
This means that setting a perfectly linear ramp, or gamma 1.0, will produce
the default (usually sRGB-like) behavior.

For gamma correct rendering with OpenGL or OpenGL ES, see the [SRGB_CAPABLE] hint.

C documentation : [glfwSetGamma]

# params
  - monitor - The monitor whose gamma ramp to set.
  - gamma - The desired exponent.

# errors

Possible errors include [NOT_INITIALIZED], [INVALID_VALUE],
[PLATFORM_ERROR] and [FEATURE_UNAVAILABLE] (see remarks).

# wayland

Gamma handling is a privileged protocol, this function
will thus never be implemented and emits [FEATURE_UNAVAILABLE].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[glfwSetGamma]: https://www.glfw.org/docs/latest/group__monitor.html#ga6ac582625c990220785ddd34efa3169a
*/
func (monitor *Monitor) SetGamma(gamma Float) {
	_, err := ffi.CallFunction(
		cif_glfwSetGamma,
		func_glfwSetGamma,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&monitor),
			unsafe.Pointer(&gamma),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function returns the current gamma ramp of the specified monitor.

C documentation : [glfwGetGammaRamp]

# params
  - monitor - The monitor to query.

# return

The current gamma ramp, or `NULL` if an
[error] occurred.

# errors

Possible errors include [NOT_INITIALIZED], [PLATFORM_ERROR]
and [FEATURE_UNAVAILABLE] (see remarks).

# wayland

Gamma handling is a privileged protocol, this function
will thus never be implemented and emits [FEATURE_UNAVAILABLE] while
returning `NULL`.

# pointer lifetime

The returned structure and its arrays are allocated and
freed by GLFW.  You should not free them yourself.  They are valid until the
specified monitor is disconnected, this function is called again for that
monitor or the library is terminated.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwGetGammaRamp]: https://www.glfw.org/docs/latest/group__monitor.html#ga76ba90debcf0062b5c4b73052b24f96f
*/
func (monitor *Monitor) GetGammaRamp() *Gammaramp {
	var result *Gammaramp
	_, err := ffi.CallFunction(
		cif_glfwGetGammaRamp,
		func_glfwGetGammaRamp,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&monitor),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the current gamma ramp for the specified monitor.  The
original gamma ramp for that monitor is saved by GLFW the first time this
function is called and is restored by [Terminate].

The software controlled gamma ramp is applied _in addition_ to the hardware
gamma correction, which today is usually an approximation of sRGB gamma.
This means that setting a perfectly linear ramp, or gamma 1.0, will produce
the default (usually sRGB-like) behavior.

For gamma correct rendering with OpenGL or OpenGL ES, see the [SRGB_CAPABLE] hint.

C documentation : [glfwSetGammaRamp]

# params
  - monitor - The monitor whose gamma ramp to set.
  - ramp - The gamma ramp to use.

# errors

Possible errors include [NOT_INITIALIZED], [PLATFORM_ERROR]
and [FEATURE_UNAVAILABLE] (see remarks).

The size of the specified gamma ramp should match the size of the
current ramp for that monitor.

# win32

The gamma ramp size must be 256.

# wayland

Gamma handling is a privileged protocol, this function
will thus never be implemented and emits [FEATURE_UNAVAILABLE].

# pointer lifetime

The specified gamma ramp is copied before this function
returns.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[glfwSetGammaRamp]: https://www.glfw.org/docs/latest/group__monitor.html#ga583f0ffd0d29613d8cd172b996bbf0dd
*/
func (monitor *Monitor) SetGammaRamp(ramp *Gammaramp) {
	_, err := ffi.CallFunction(
		cif_glfwSetGammaRamp,
		func_glfwSetGammaRamp,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&monitor),
			unsafe.Pointer(&ramp),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function resets all window hints to their
[default values].

C documentation : [glfwDefaultWindowHints]

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[default values]: https://www.glfw.org/docs/latest/window_guide.html#window_hints_values
[glfwDefaultWindowHints]: https://www.glfw.org/docs/latest/group__window.html#gaa77c4898dfb83344a6b4f76aa16b9a4a
*/
func DefaultWindowHints() {
	_, err := ffi.CallFunction(
		cif_glfwDefaultWindowHints,
		func_glfwDefaultWindowHints,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function sets hints for the next call to [CreateWindow].  The
hints, once set, retain their values until changed by a call to this
function or [DefaultWindowHints], or until the library is terminated.

Only integer value hints can be set with this function.  String value hints
are set with [WindowHintString].

This function does not check whether the specified hint values are valid.
If you set hints to invalid values this will instead be reported by the next
call to [CreateWindow].

Some hints are platform specific.  These may be set on any platform but they
will only affect their specific platform.  Other platforms will ignore them.
Setting these hints requires no platform specific headers or functions.

C documentation : [glfwWindowHint]

# params
  - hint - The [window hint] to set.
  - value - The new value of the window hint.

# errors

Possible errors include [NOT_INITIALIZED] and [INVALID_ENUM].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.  Replaces `glfwOpenWindowHint`.

[window hint]: https://www.glfw.org/docs/latest/window_guide.html#window_hints
[glfwWindowHint]: https://www.glfw.org/docs/latest/group__window.html#ga7d9c8c62384b1e2821c4dc48952d2033
*/
func WindowHint(hint Int, value Int) {
	_, err := ffi.CallFunction(
		cif_glfwWindowHint,
		func_glfwWindowHint,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&hint),
			unsafe.Pointer(&value),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function sets hints for the next call to [CreateWindow].  The
hints, once set, retain their values until changed by a call to this
function or [DefaultWindowHints], or until the library is terminated.

Only string type hints can be set with this function.  Integer value hints
are set with [WindowHint].

This function does not check whether the specified hint values are valid.
If you set hints to invalid values this will instead be reported by the next
call to [CreateWindow].

Some hints are platform specific.  These may be set on any platform but they
will only affect their specific platform.  Other platforms will ignore them.
Setting these hints requires no platform specific headers or functions.

C documentation : [glfwWindowHintString]

# params
  - hint - The [window hint] to set.
  - value - The new value of the window hint.

# errors

Possible errors include [NOT_INITIALIZED] and [INVALID_ENUM].

# pointer lifetime

The specified string is copied before this function
returns.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.3.

[window hint]: https://www.glfw.org/docs/latest/window_guide.html#window_hints
[glfwWindowHintString]: https://www.glfw.org/docs/latest/group__window.html#ga8cb2782861c9d997bcf2dea97f363e5f
*/
func WindowHintString(hint Int, value string) {
	c_value := cString(value)
	_, err := ffi.CallFunction(
		cif_glfwWindowHintString,
		func_glfwWindowHintString,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&hint),
			unsafe.Pointer(&c_value),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function creates a window and its associated OpenGL or OpenGL ES
context.  Most of the options controlling how the window and its context
should be created are specified with [window hints].

Successful creation does not change which context is current.  Before you
can use the newly created context, you need to
[make it current].  For information about the `share`
parameter, see [Context object sharing].

The created window, framebuffer and context may differ from what you
requested, as not all parameters and hints are
[hard constraints].  This includes the size of the
window, especially for full screen windows.  To query the actual attributes
of the created window, framebuffer and context, see [GetAttrib], [GetSize] and [GetFramebufferSize].

To create a full screen window, you need to specify the monitor the window
will cover.  If no monitor is specified, the window will be windowed mode.
Unless you have a way for the user to choose a specific monitor, it is
recommended that you pick the primary monitor.  For more information on how
to query connected monitors, see [Retrieving monitors].

For full screen windows, the specified size becomes the resolution of the
window's _desired video mode_.  As long as a full screen window is not
iconified, the supported video mode most closely matching the desired video
mode is set for the specified monitor.  For more information about full
screen windows, including the creation of so called _windowed full screen_
or _borderless full screen_ windows, see ["Windowed full screen" windows].

Once you have created the window, you can switch it between windowed and
full screen mode with [SetMonitor].  This will not affect its
OpenGL or OpenGL ES context.

By default, newly created windows use the placement recommended by the
window system.  To create the window at a specific position, set the [POSITION_X] and [POSITION_Y] window hints before creation.  To
restore the default behavior, set either or both hints back to
`GLFW_ANY_POSITION`.

As long as at least one full screen window is not iconified, the screensaver
is prohibited from starting.

Window systems put limits on window sizes.  Very large or very small window
dimensions may be overridden by the window system on creation.  Check the
actual [size] after creation.

The [swap interval] is not set during window creation and
the initial value may vary depending on driver settings and defaults.

C documentation : [glfwCreateWindow]

# params
  - width - The desired width, in screen coordinates, of the window.
    This must be greater than zero.
  - height - The desired height, in screen coordinates, of the window.
    This must be greater than zero.
  - title - The initial, UTF-8 encoded window title.
  - monitor - The monitor to use for full screen mode, or `NULL` for
    windowed mode.
  - share - The window whose context to share resources with, or `NULL`
    to not share resources.

# return

The handle of the created window, or `NULL` if an
[error] occurred.

# errors

Possible errors include [NOT_INITIALIZED], [INVALID_ENUM], [INVALID_VALUE], [API_UNAVAILABLE], [VERSION_UNAVAILABLE], [FORMAT_UNAVAILABLE], [NO_WINDOW_CONTEXT] and [PLATFORM_ERROR].

# win32

Window creation will fail if the Microsoft GDI software
OpenGL implementation is the only one available.

# win32

If the executable has an icon resource named `GLFW_ICON,` it
will be set as the initial icon for the window.  If no such icon is present,
the `IDI_APPLICATION` icon will be used instead.  To set a different icon,
see [SetIcon].

# win32

The context to share resources with must not be current on
any other thread.

# macos

The OS only supports core profile contexts for OpenGL
versions 3.2 and later.  Before creating an OpenGL context of version 3.2 or
later you must set the [GLFW_OPENGL_PROFILE]
hint accordingly.  OpenGL 3.0 and 3.1 contexts are not supported at all
on macOS.

# macos

The GLFW window has no icon, as it is not a document
window, but the dock icon will be the same as the application bundle's icon.
For more information on bundles, see the
[Bundle Programming Guide][bundle-guide] in the Mac Developer Library.

# macos

On OS X 10.10 and later the window frame will not be rendered
at full resolution on Retina displays unless the
[GLFW_SCALE_FRAMEBUFFER]
hint is `GLFW_TRUE` and the `NSHighResolutionCapable` key is enabled in the
application bundle's `Info.plist`.  For more information, see
[High Resolution Guidelines for OS X][hidpi-guide] in the Mac Developer
Library.  The GLFW test and example programs use a custom `Info.plist`
template for this, which can be found as `CMake/Info.plist.in` in the source
tree.

# macos

When activating frame autosaving with
[GLFW_COCOA_FRAME_NAME], the specified
window size and position may be overridden by previously saved values.

# wayland

GLFW uses [libdecor][] where available to create its window
decorations.  This in turn uses server-side XDG decorations where available
and provides high quality client-side decorations on compositors like GNOME.
If both XDG decorations and libdecor are unavailable, GLFW falls back to
a very simple set of window decorations that only support moving, resizing
and the window manager's right-click menu.

# x11

Some window managers will not respect the placement of
initially hidden windows.

# x11

Due to the asynchronous nature of X11, it may take a moment for
a window to reach its requested state.  This means you may not be able to
query the final size, position or other attributes directly after window
creation.

# x11

The class part of the `WM_CLASS` window property will by
default be set to the window title passed to this function.  The instance
part will use the contents of the `RESOURCE_NAME` environment variable, if
present and not empty, or fall back to the window title.  Set the
[GLFW_X11_CLASS_NAME] and
[GLFW_X11_INSTANCE_NAME] window hints to
override this.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.  Replaces `glfwOpenWindow`.

[bundle-guide]: https://developer.apple.com/library/mac/documentation/CoreFoundation/Conceptual/CFBundles/
[hidpi-guide]: https://developer.apple.com/library/mac/documentation/GraphicsAnimation/Conceptual/HighResolutionOSX/Explained/Explained.html
[libdecor]: https://gitlab.freedesktop.org/libdecor/libdecor
[window hints]: https://www.glfw.org/docs/latest/window_guide.html#window_hints
[make it current]: https://www.glfw.org/docs/latest/context_guide.html#context_current
[hard constraints]: https://www.glfw.org/docs/latest/window_guide.html#window_hints_hard
[size]: https://www.glfw.org/docs/latest/window_guide.html#window_size
[swap interval]: https://www.glfw.org/docs/latest/window_guide.html#buffer_swap
[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[GLFW_OPENGL_PROFILE]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_OPENGL_PROFILE_hint
[GLFW_SCALE_FRAMEBUFFER]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_SCALE_FRAMEBUFFER_hint
[GLFW_COCOA_FRAME_NAME]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_COCOA_FRAME_NAME_hint
[GLFW_X11_CLASS_NAME]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_X11_CLASS_NAME_hint
[GLFW_X11_INSTANCE_NAME]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_X11_INSTANCE_NAME_hint
[Context object sharing]: https://www.glfw.org/docs/latest/context_guide.html#context_sharing
[Retrieving monitors]: https://www.glfw.org/docs/latest/monitor_guide.html#monitor_monitors
["Windowed full screen" windows]: https://www.glfw.org/docs/latest/window_guide.html#window_windowed_full_screen
[glfwCreateWindow]: https://www.glfw.org/docs/latest/group__window.html#ga3555a418df92ad53f917597fe2f64aeb
*/
func CreateWindow(width Int, height Int, title string, monitor *Monitor, share *Window) *Window {
	var result *Window
	c_title := cString(title)
	_, err := ffi.CallFunction(
		cif_glfwCreateWindow,
		func_glfwCreateWindow,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&width),
			unsafe.Pointer(&height),
			unsafe.Pointer(&c_title),
			unsafe.Pointer(&monitor),
			unsafe.Pointer(&share),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function destroys the specified window and its context.  On calling
this function, no further callbacks will be called for that window.

If the context of the specified window is current on the main thread, it is
detached before being destroyed.

C documentation : [glfwDestroyWindow]

# params
  - window - The window to destroy.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

The context of the specified window must not be current on any other
thread when this function is called.

# reentrancy

This function must not be called from a callback.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.  Replaces `glfwCloseWindow`.

[glfwDestroyWindow]: https://www.glfw.org/docs/latest/group__window.html#gacdf43e51376051d2c091662e9fe3d7b2
*/
func (window *Window) Destroy() {
	_, err := ffi.CallFunction(
		cif_glfwDestroyWindow,
		func_glfwDestroyWindow,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function returns the value of the close flag of the specified window.

C documentation : [glfwWindowShouldClose]

# params
  - window - The window to query.

# return

The value of the close flag.

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function may be called from any thread.  Access is not
synchronized.

# since

Added in version 3.0.

[glfwWindowShouldClose]: https://www.glfw.org/docs/latest/group__window.html#ga24e02fbfefbb81fc45320989f8140ab5
*/
func (window *Window) ShouldClose() Int {
	var result Int
	_, err := ffi.CallFunction(
		cif_glfwWindowShouldClose,
		func_glfwWindowShouldClose,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the value of the close flag of the specified window.
This can be used to override the user's attempt to close the window, or
to signal that it should be closed.

C documentation : [glfwSetWindowShouldClose]

# params
  - window - The window whose flag to change.
  - value - The new value.

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function may be called from any thread.  Access is not
synchronized.

# since

Added in version 3.0.

[glfwSetWindowShouldClose]: https://www.glfw.org/docs/latest/group__window.html#ga49c449dde2a6f87d996f4daaa09d6708
*/
func (window *Window) SetShouldClose(value Int) {
	_, err := ffi.CallFunction(
		cif_glfwSetWindowShouldClose,
		func_glfwSetWindowShouldClose,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&value),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function returns the window title, encoded as UTF-8, of the specified
window.  This is the title set previously by [CreateWindow]
or [SetTitle].

C documentation : [glfwGetWindowTitle]

# params
  - window - The window to query.

# return

The UTF-8 encoded window title, or `NULL` if an
[error] occurred.

# errors

Possible errors include [NOT_INITIALIZED].

The returned title is currently a copy of the title last set by [CreateWindow] or [SetTitle].  It does not include any
additional text which may be appended by the platform or another program.

# pointer lifetime

The returned string is allocated and freed by GLFW.  You
should not free it yourself.  It is valid until the next call to [GetTitle] or [SetTitle], or until the library is
terminated.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.4.

[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwGetWindowTitle]: https://www.glfw.org/docs/latest/group__window.html#gac6151765c54b789c4fe66c6bc6215953
*/
func (window *Window) GetTitle() string {
	var result *byte
	_, err := ffi.CallFunction(
		cif_glfwGetWindowTitle,
		func_glfwGetWindowTitle,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
		},
	)
	if err != nil {
		panic(err)
	}
	return goString(result)
}

/*
This function sets the window title, encoded as UTF-8, of the specified
window.

C documentation : [glfwSetWindowTitle]

# params
  - window - The window whose title to change.
  - title - The UTF-8 encoded window title.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# macos

The window title will not be updated until the next time you
process events.

# thread safety

This function must only be called from the main thread.

# since

Added in version 1.0.

[glfwSetWindowTitle]: https://www.glfw.org/docs/latest/group__window.html#ga5d877f09e968cef7a360b513306f17ff
*/
func (window *Window) SetTitle(title string) {
	c_title := cString(title)
	_, err := ffi.CallFunction(
		cif_glfwSetWindowTitle,
		func_glfwSetWindowTitle,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&c_title),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function sets the icon of the specified window.  If passed an array of
candidate images, those of or closest to the sizes desired by the system are
selected.  If no images are specified, the window reverts to its default
icon.

The pixels are 32-bit, little-endian, non-premultiplied RGBA, i.e. eight
bits per channel with the red channel first.  They are arranged canonically
as packed sequential rows, starting from the top-left corner.

The desired image sizes varies depending on platform and system settings.
The selected images will be rescaled as needed.  Good sizes include 16x16,
32x32 and 48x48.

C documentation : [glfwSetWindowIcon]

# params
  - window - The window whose icon to set.
  - count - The number of images in the specified array, or zero to
    revert to the default window icon.
  - images - The images to create the icon from.  This is ignored if
    count is zero.

# errors

Possible errors include [NOT_INITIALIZED], [INVALID_VALUE], [PLATFORM_ERROR] and [FEATURE_UNAVAILABLE] (see remarks).

# pointer lifetime

The specified image data is copied before this function
returns.

# macos

Regular windows do not have icons on macOS.  This function
will emit [FEATURE_UNAVAILABLE].  The dock icon will be the same as
the application bundle's icon.  For more information on bundles, see the
[Bundle Programming Guide][bundle-guide] in the Mac Developer Library.

# wayland

There is no existing protocol to change an icon, the
window will thus inherit the one defined in the application's desktop file.
This function will emit [FEATURE_UNAVAILABLE].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.2.

[bundle-guide]: https://developer.apple.com/library/mac/documentation/CoreFoundation/Conceptual/CFBundles/
[glfwSetWindowIcon]: https://www.glfw.org/docs/latest/group__window.html#gadd7ccd39fe7a7d1f0904666ae5932dc5
*/
func (window *Window) SetIcon(count Int, images *Image) {
	_, err := ffi.CallFunction(
		cif_glfwSetWindowIcon,
		func_glfwSetWindowIcon,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&count),
			unsafe.Pointer(&images),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function retrieves the position, in screen coordinates, of the
upper-left corner of the content area of the specified window.

Any or all of the position arguments may be `NULL`.  If an error occurs, all
non-`NULL` position arguments will be set to zero.

C documentation : [glfwGetWindowPos]

# params
  - window - The window to query.
  - xpos (out) - Where to store the x-coordinate of the upper-left corner of
    the content area, or `NULL`.
  - ypos (out) - Where to store the y-coordinate of the upper-left corner of
    the content area, or `NULL`.

# errors

Possible errors include [NOT_INITIALIZED], [PLATFORM_ERROR] and [FEATURE_UNAVAILABLE] (see remarks).

# wayland

There is no way for an application to retrieve the global
position of its windows.  This function will emit [FEATURE_UNAVAILABLE].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[glfwGetWindowPos]: https://www.glfw.org/docs/latest/group__window.html#ga73cb526c000876fd8ddf571570fdb634
*/
func (window *Window) GetPos() (Int, Int) {
	var xpos Int
	var ypos Int
	_, err := ffi.CallFunction(
		cif_glfwGetWindowPos,
		func_glfwGetWindowPos,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(new(&xpos)),
			unsafe.Pointer(new(&ypos)),
		},
	)
	if err != nil {
		panic(err)
	}
	return xpos, ypos
}

/*
This function sets the position, in screen coordinates, of the upper-left
corner of the content area of the specified windowed mode window.  If the
window is a full screen window, this function does nothing.

__Do not use this function__ to move an already visible window unless you
have very good reasons for doing so, as it will confuse and annoy the user.

The window manager may put limits on what positions are allowed.  GLFW
cannot and should not override these limits.

C documentation : [glfwSetWindowPos]

# params
  - window - The window to query.
  - xpos - The x-coordinate of the upper-left corner of the content area.
  - ypos - The y-coordinate of the upper-left corner of the content area.

# errors

Possible errors include [NOT_INITIALIZED], [PLATFORM_ERROR] and [FEATURE_UNAVAILABLE] (see remarks).

# wayland

There is no way for an application to set the global
position of its windows.  This function will emit [FEATURE_UNAVAILABLE].

# thread safety

This function must only be called from the main thread.

# since

Added in version 1.0.

[glfwSetWindowPos]: https://www.glfw.org/docs/latest/group__window.html#ga1abb6d690e8c88e0c8cd1751356dbca8
*/
func (window *Window) SetPos(xpos Int, ypos Int) {
	_, err := ffi.CallFunction(
		cif_glfwSetWindowPos,
		func_glfwSetWindowPos,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&xpos),
			unsafe.Pointer(&ypos),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function retrieves the size, in screen coordinates, of the content area
of the specified window.  If you wish to retrieve the size of the
framebuffer of the window in pixels, see [GetFramebufferSize].

Any or all of the size arguments may be `NULL`.  If an error occurs, all
non-`NULL` size arguments will be set to zero.

C documentation : [glfwGetWindowSize]

# params
  - window - The window whose size to retrieve.
  - width (out) - Where to store the width, in screen coordinates, of the
    content area, or `NULL`.
  - height (out) - Where to store the height, in screen coordinates, of the
    content area, or `NULL`.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# thread safety

This function must only be called from the main thread.

# since

Added in version 1.0.

[glfwGetWindowSize]: https://www.glfw.org/docs/latest/group__window.html#gaeea7cbc03373a41fb51cfbf9f2a5d4c6
*/
func (window *Window) GetSize() (Int, Int) {
	var width Int
	var height Int
	_, err := ffi.CallFunction(
		cif_glfwGetWindowSize,
		func_glfwGetWindowSize,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(new(&width)),
			unsafe.Pointer(new(&height)),
		},
	)
	if err != nil {
		panic(err)
	}
	return width, height
}

/*
This function sets the size limits of the content area of the specified
window.  If the window is full screen, the size limits only take effect
once it is made windowed.  If the window is not resizable, this function
does nothing.

The size limits are applied immediately to a windowed mode window and may
cause it to be resized.

The maximum dimensions must be greater than or equal to the minimum
dimensions and all must be greater than or equal to zero.

C documentation : [glfwSetWindowSizeLimits]

# params
  - window - The window to set limits for.
  - minwidth - The minimum width, in screen coordinates, of the content
    area, or `GLFW_DONT_CARE`.
  - minheight - The minimum height, in screen coordinates, of the
    content area, or `GLFW_DONT_CARE`.
  - maxwidth - The maximum width, in screen coordinates, of the content
    area, or `GLFW_DONT_CARE`.
  - maxheight - The maximum height, in screen coordinates, of the
    content area, or `GLFW_DONT_CARE`.

# errors

Possible errors include [NOT_INITIALIZED], [INVALID_VALUE] and [PLATFORM_ERROR].

If you set size limits and an aspect ratio that conflict, the
results are undefined.

# wayland

The size limits will not be applied until the window is
actually resized, either by the user or by the compositor.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.2.

[glfwSetWindowSizeLimits]: https://www.glfw.org/docs/latest/group__window.html#gac314fa6cec7d2d307be9963e2709cc90
*/
func (window *Window) SetSizeLimits(minwidth Int, minheight Int, maxwidth Int, maxheight Int) {
	_, err := ffi.CallFunction(
		cif_glfwSetWindowSizeLimits,
		func_glfwSetWindowSizeLimits,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&minwidth),
			unsafe.Pointer(&minheight),
			unsafe.Pointer(&maxwidth),
			unsafe.Pointer(&maxheight),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function sets the required aspect ratio of the content area of the
specified window.  If the window is full screen, the aspect ratio only takes
effect once it is made windowed.  If the window is not resizable, this
function does nothing.

The aspect ratio is specified as a numerator and a denominator and both
values must be greater than zero.  For example, the common 16:9 aspect ratio
is specified as 16 and 9, respectively.

If the numerator and denominator is set to `GLFW_DONT_CARE` then the aspect
ratio limit is disabled.

The aspect ratio is applied immediately to a windowed mode window and may
cause it to be resized.

C documentation : [glfwSetWindowAspectRatio]

# params
  - window - The window to set limits for.
  - numer - The numerator of the desired aspect ratio, or
    `GLFW_DONT_CARE`.
  - denom - The denominator of the desired aspect ratio, or
    `GLFW_DONT_CARE`.

# errors

Possible errors include [NOT_INITIALIZED], [INVALID_VALUE] and [PLATFORM_ERROR].

If you set size limits and an aspect ratio that conflict, the
results are undefined.

# wayland

The aspect ratio will not be applied until the window is
actually resized, either by the user or by the compositor.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.2.

[glfwSetWindowAspectRatio]: https://www.glfw.org/docs/latest/group__window.html#ga72ac8cb1ee2e312a878b55153d81b937
*/
func (window *Window) SetAspectRatio(numer Int, denom Int) {
	_, err := ffi.CallFunction(
		cif_glfwSetWindowAspectRatio,
		func_glfwSetWindowAspectRatio,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&numer),
			unsafe.Pointer(&denom),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function sets the size, in screen coordinates, of the content area of
the specified window.

For full screen windows, this function updates the resolution of its desired
video mode and switches to the video mode closest to it, without affecting
the window's context.  As the context is unaffected, the bit depths of the
framebuffer remain unchanged.

If you wish to update the refresh rate of the desired video mode in addition
to its resolution, see [SetMonitor].

The window manager may put limits on what sizes are allowed.  GLFW cannot
and should not override these limits.

C documentation : [glfwSetWindowSize]

# params
  - window - The window to resize.
  - width - The desired width, in screen coordinates, of the window
    content area.
  - height - The desired height, in screen coordinates, of the window
    content area.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# thread safety

This function must only be called from the main thread.

# since

Added in version 1.0.

[glfwSetWindowSize]: https://www.glfw.org/docs/latest/group__window.html#ga371911f12c74c504dd8d47d832d095cb
*/
func (window *Window) SetSize(width Int, height Int) {
	_, err := ffi.CallFunction(
		cif_glfwSetWindowSize,
		func_glfwSetWindowSize,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&width),
			unsafe.Pointer(&height),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function retrieves the size, in pixels, of the framebuffer of the
specified window.  If you wish to retrieve the size of the window in screen
coordinates, see [GetSize].

Any or all of the size arguments may be `NULL`.  If an error occurs, all
non-`NULL` size arguments will be set to zero.

C documentation : [glfwGetFramebufferSize]

# params
  - window - The window whose framebuffer to query.
  - width (out) - Where to store the width, in pixels, of the framebuffer,
    or `NULL`.
  - height (out) - Where to store the height, in pixels, of the framebuffer,
    or `NULL`.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[glfwGetFramebufferSize]: https://www.glfw.org/docs/latest/group__window.html#ga0e2637a4161afb283f5300c7f94785c9
*/
func (window *Window) GetFramebufferSize() (Int, Int) {
	var width Int
	var height Int
	_, err := ffi.CallFunction(
		cif_glfwGetFramebufferSize,
		func_glfwGetFramebufferSize,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(new(&width)),
			unsafe.Pointer(new(&height)),
		},
	)
	if err != nil {
		panic(err)
	}
	return width, height
}

/*
This function retrieves the size, in screen coordinates, of each edge of the
frame of the specified window.  This size includes the title bar, if the
window has one.  The size of the frame may vary depending on the
[window-related hints]([Window related hints]) used to create it.

Because this function retrieves the size of each window frame edge and not
the offset along a particular coordinate axis, the retrieved values will
always be zero or positive.

Any or all of the size arguments may be `NULL`.  If an error occurs, all
non-`NULL` size arguments will be set to zero.

C documentation : [glfwGetWindowFrameSize]

# params
  - window - The window whose frame size to query.
  - left (out) - Where to store the size, in screen coordinates, of the left
    edge of the window frame, or `NULL`.
  - top (out) - Where to store the size, in screen coordinates, of the top
    edge of the window frame, or `NULL`.
  - right (out) - Where to store the size, in screen coordinates, of the
    right edge of the window frame, or `NULL`.
  - bottom (out) - Where to store the size, in screen coordinates, of the
    bottom edge of the window frame, or `NULL`.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.1.

[Window related hints]: https://www.glfw.org/docs/latest/window_guide.html#window_hints_wnd
[glfwGetWindowFrameSize]: https://www.glfw.org/docs/latest/group__window.html#ga1a9fd382058c53101b21cf211898f1f1
*/
func (window *Window) GetFrameSize() (Int, Int, Int, Int) {
	var left Int
	var top Int
	var right Int
	var bottom Int
	_, err := ffi.CallFunction(
		cif_glfwGetWindowFrameSize,
		func_glfwGetWindowFrameSize,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(new(&left)),
			unsafe.Pointer(new(&top)),
			unsafe.Pointer(new(&right)),
			unsafe.Pointer(new(&bottom)),
		},
	)
	if err != nil {
		panic(err)
	}
	return left, top, right, bottom
}

/*
This function retrieves the content scale for the specified window.  The
content scale is the ratio between the current DPI and the platform's
default DPI.  This is especially important for text and any UI elements.  If
the pixel dimensions of your UI scaled by this look appropriate on your
machine then it should appear at a reasonable size on other machines
regardless of their DPI and scaling settings.  This relies on the system DPI
and scaling settings being somewhat correct.

On platforms where each monitors can have its own content scale, the window
content scale will depend on which monitor the system considers the window
to be on.

C documentation : [glfwGetWindowContentScale]

# params
  - window - The window to query.
  - xscale (out) - Where to store the x-axis content scale, or `NULL`.
  - yscale (out) - Where to store the y-axis content scale, or `NULL`.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.3.

[glfwGetWindowContentScale]: https://www.glfw.org/docs/latest/group__window.html#gaf5d31de9c19c4f994facea64d2b3106c
*/
func (window *Window) GetContentScale() (Float, Float) {
	var xscale Float
	var yscale Float
	_, err := ffi.CallFunction(
		cif_glfwGetWindowContentScale,
		func_glfwGetWindowContentScale,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(new(&xscale)),
			unsafe.Pointer(new(&yscale)),
		},
	)
	if err != nil {
		panic(err)
	}
	return xscale, yscale
}

/*
This function returns the opacity of the window, including any decorations.

The opacity (or alpha) value is a positive finite number between zero and
one, where zero is fully transparent and one is fully opaque.  If the system
does not support whole window transparency, this function always returns one.

The initial opacity value for newly created windows is one.

C documentation : [glfwGetWindowOpacity]

# params
  - window - The window to query.

# return

The opacity value of the specified window.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.3.

[glfwGetWindowOpacity]: https://www.glfw.org/docs/latest/group__window.html#gad09f0bd7a6307c4533b7061828480a84
*/
func (window *Window) GetOpacity() Float {
	var result Float
	_, err := ffi.CallFunction(
		cif_glfwGetWindowOpacity,
		func_glfwGetWindowOpacity,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the opacity of the window, including any decorations.

The opacity (or alpha) value is a positive finite number between zero and
one, where zero is fully transparent and one is fully opaque.

The initial opacity value for newly created windows is one.

A window created with framebuffer transparency may not use whole window
transparency.  The results of doing this are undefined.

C documentation : [glfwSetWindowOpacity]

# params
  - window - The window to set the opacity for.
  - opacity - The desired opacity of the specified window.

# errors

Possible errors include [NOT_INITIALIZED], [PLATFORM_ERROR] and [FEATURE_UNAVAILABLE] (see remarks).

# wayland

There is no way to set an opacity factor for a window.
This function will emit [FEATURE_UNAVAILABLE].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.3.

[glfwSetWindowOpacity]: https://www.glfw.org/docs/latest/group__window.html#gac31caeb3d1088831b13d2c8a156802e9
*/
func (window *Window) SetOpacity(opacity Float) {
	_, err := ffi.CallFunction(
		cif_glfwSetWindowOpacity,
		func_glfwSetWindowOpacity,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&opacity),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function iconifies (minimizes) the specified window if it was
previously restored.  If the window is already iconified, this function does
nothing.

If the specified window is a full screen window, GLFW restores the original
video mode of the monitor.  The window's desired video mode is set again
when the window is restored.

C documentation : [glfwIconifyWindow]

# params
  - window - The window to iconify.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# wayland

Once a window is iconified, [Restore] won’t
be able to restore it.  This is a design decision of the xdg-shell
protocol.

# thread safety

This function must only be called from the main thread.

# since

Added in version 2.1.

[glfwIconifyWindow]: https://www.glfw.org/docs/latest/group__window.html#ga1bb559c0ebaad63c5c05ad2a066779c4
*/
func (window *Window) Iconify() {
	_, err := ffi.CallFunction(
		cif_glfwIconifyWindow,
		func_glfwIconifyWindow,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function restores the specified window if it was previously iconified
(minimized) or maximized.  If the window is already restored, this function
does nothing.

If the specified window is an iconified full screen window, its desired
video mode is set again for its monitor when the window is restored.

C documentation : [glfwRestoreWindow]

# params
  - window - The window to restore.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# thread safety

This function must only be called from the main thread.

# since

Added in version 2.1.

[glfwRestoreWindow]: https://www.glfw.org/docs/latest/group__window.html#ga52527a5904b47d802b6b4bb519cdebc7
*/
func (window *Window) Restore() {
	_, err := ffi.CallFunction(
		cif_glfwRestoreWindow,
		func_glfwRestoreWindow,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function maximizes the specified window if it was previously not
maximized.  If the window is already maximized, this function does nothing.

If the specified window is a full screen window, this function does nothing.

C documentation : [glfwMaximizeWindow]

# params
  - window - The window to maximize.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# Thread Safety

This function may only be called from the main thread.

# since

Added in GLFW 3.2.

[glfwMaximizeWindow]: https://www.glfw.org/docs/latest/group__window.html#ga3f541387449d911274324ae7f17ec56b
*/
func (window *Window) Maximize() {
	_, err := ffi.CallFunction(
		cif_glfwMaximizeWindow,
		func_glfwMaximizeWindow,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function makes the specified window visible if it was previously
hidden.  If the window is already visible or is in full screen mode, this
function does nothing.

By default, windowed mode windows are focused when shown
Set the [GLFW_FOCUS_ON_SHOW] window hint
to change this behavior for all newly created windows, or change the
behavior for an existing window with [SetAttrib].

C documentation : [glfwShowWindow]

# params
  - window - The window to make visible.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# wayland

Because Wayland wants every frame of the desktop to be
complete, this function does not immediately make the window visible.
Instead it will become visible the next time the window framebuffer is
updated after this call.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[GLFW_FOCUS_ON_SHOW]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_FOCUS_ON_SHOW_hint
[glfwShowWindow]: https://www.glfw.org/docs/latest/group__window.html#ga61be47917b72536a148300f46494fc66
*/
func (window *Window) Show() {
	_, err := ffi.CallFunction(
		cif_glfwShowWindow,
		func_glfwShowWindow,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function hides the specified window if it was previously visible.  If
the window is already hidden or is in full screen mode, this function does
nothing.

C documentation : [glfwHideWindow]

# params
  - window - The window to hide.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[glfwHideWindow]: https://www.glfw.org/docs/latest/group__window.html#ga49401f82a1ba5f15db5590728314d47c
*/
func (window *Window) Hide() {
	_, err := ffi.CallFunction(
		cif_glfwHideWindow,
		func_glfwHideWindow,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function brings the specified window to front and sets input focus.
The window should already be visible and not iconified.

By default, both windowed and full screen mode windows are focused when
initially created.  Set the [GLFW_FOCUSED] to
disable this behavior.

Also by default, windowed mode windows are focused when shown
with [Show]. Set the
[GLFW_FOCUS_ON_SHOW] to disable this behavior.

__Do not use this function__ to steal focus from other applications unless
you are certain that is what the user wants.  Focus stealing can be
extremely disruptive.

For a less disruptive way of getting the user's attention, see
[attention requests].

C documentation : [glfwFocusWindow]

# params
  - window - The window to give input focus.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# wayland

The compositor will likely ignore focus requests unless
another window created by the same application already has input focus.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.2.

[GLFW_FOCUSED]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_FOCUSED_hint
[GLFW_FOCUS_ON_SHOW]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_FOCUS_ON_SHOW_hint
[attention requests]: https://www.glfw.org/docs/latest/window_guide.html#window_attention
[glfwFocusWindow]: https://www.glfw.org/docs/latest/group__window.html#ga873780357abd3f3a081d71a40aae45a1
*/
func (window *Window) Focus() {
	_, err := ffi.CallFunction(
		cif_glfwFocusWindow,
		func_glfwFocusWindow,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function requests user attention to the specified window.  On
platforms where this is not supported, attention is requested to the
application as a whole.

Once the user has given attention, usually by focusing the window or
application, the system will end the request automatically.

C documentation : [glfwRequestWindowAttention]

# params
  - window - The window to request attention to.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# macos

Attention is requested to the application as a whole, not the
specific window.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.3.

[glfwRequestWindowAttention]: https://www.glfw.org/docs/latest/group__window.html#ga2f8d59323fc4692c1d54ba08c863a703
*/
func (window *Window) RequestAttention() {
	_, err := ffi.CallFunction(
		cif_glfwRequestWindowAttention,
		func_glfwRequestWindowAttention,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function returns the handle of the monitor that the specified window is
in full screen on.

C documentation : [glfwGetWindowMonitor]

# params
  - window - The window to query.

# return

The monitor, or `NULL` if the window is in windowed mode or an
[error] occurred.

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwGetWindowMonitor]: https://www.glfw.org/docs/latest/group__window.html#ga4d766499ac02c60f02221a9dfab87299
*/
func (window *Window) GetMonitor() *Monitor {
	var result *Monitor
	_, err := ffi.CallFunction(
		cif_glfwGetWindowMonitor,
		func_glfwGetWindowMonitor,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the monitor that the window uses for full screen mode or,
if the monitor is `NULL`, makes it windowed mode.

When setting a monitor, this function updates the width, height and refresh
rate of the desired video mode and switches to the video mode closest to it.
The window position is ignored when setting a monitor.

When the monitor is `NULL`, the position, width and height are used to
place the window content area.  The refresh rate is ignored when no monitor
is specified.

If you only wish to update the resolution of a full screen window or the
size of a windowed mode window, see [SetSize].

When a window transitions from full screen to windowed mode, this function
restores any previous window settings such as whether it is decorated,
floating, resizable, has size or aspect ratio limits, etc.

C documentation : [glfwSetWindowMonitor]

# params
  - window - The window whose monitor, size or video mode to set.
  - monitor - The desired monitor, or `NULL` to set windowed mode.
  - xpos - The desired x-coordinate of the upper-left corner of the
    content area.
  - ypos - The desired y-coordinate of the upper-left corner of the
    content area.
  - width - The desired with, in screen coordinates, of the content
    area or video mode.
  - height - The desired height, in screen coordinates, of the content
    area or video mode.
  - refreshRate - The desired refresh rate, in Hz, of the video mode,
    or `GLFW_DONT_CARE`.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

The OpenGL or OpenGL ES context will not be destroyed or otherwise
affected by any resizing or mode switching, although you may need to update
your viewport if the framebuffer size has changed.

# wayland

The desired window position is ignored, as there is no way
for an application to set this property.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.2.

[glfwSetWindowMonitor]: https://www.glfw.org/docs/latest/group__window.html#ga81c76c418af80a1cce7055bccb0ae0a7
*/
func (window *Window) SetMonitor(monitor *Monitor, xpos Int, ypos Int, width Int, height Int, refreshRate Int) {
	_, err := ffi.CallFunction(
		cif_glfwSetWindowMonitor,
		func_glfwSetWindowMonitor,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&monitor),
			unsafe.Pointer(&xpos),
			unsafe.Pointer(&ypos),
			unsafe.Pointer(&width),
			unsafe.Pointer(&height),
			unsafe.Pointer(&refreshRate),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function returns the value of an attribute of the specified window or
its OpenGL or OpenGL ES context.

C documentation : [glfwGetWindowAttrib]

# params
  - window - The window to query.
  - attrib - The [window attribute] whose value to
    return.

# return

The value of the attribute, or zero if an
[error] occurred.

# errors

Possible errors include [NOT_INITIALIZED], [INVALID_ENUM] and [PLATFORM_ERROR].

Framebuffer related hints are not window attributes.  See [Framebuffer related attributes] for more information.

Zero is a valid value for many window and context related
attributes so you cannot use a return value of zero as an indication of
errors.  However, this function should not fail as long as it is passed
valid arguments and the library has been [initialized].

# wayland

The Wayland protocol provides no way to check whether a
window is iconfied, so [ICONIFIED] always returns `GLFW_FALSE`.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.  Replaces `glfwGetWindowParam` and
`glfwGetGLVersion`.

[window attribute]: https://www.glfw.org/docs/latest/window_guide.html#window_attribs
[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[initialized]: https://www.glfw.org/docs/latest/intro_guide.html#intro_init
[Framebuffer related attributes]: https://www.glfw.org/docs/latest/window_guide.html#window_attribs_fb
[glfwGetWindowAttrib]: https://www.glfw.org/docs/latest/group__window.html#gacccb29947ea4b16860ebef42c2cb9337
*/
func (window *Window) GetAttrib(attrib Int) Int {
	var result Int
	_, err := ffi.CallFunction(
		cif_glfwGetWindowAttrib,
		func_glfwGetWindowAttrib,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&attrib),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the value of an attribute of the specified window.

The supported attributes are [GLFW_DECORATED],
[GLFW_RESIZABLE],
[GLFW_FLOATING],
[GLFW_AUTO_ICONIFY] and
[GLFW_FOCUS_ON_SHOW].
[GLFW_MOUSE_PASSTHROUGH]

Some of these attributes are ignored for full screen windows.  The new
value will take effect if the window is later made windowed.

Some of these attributes are ignored for windowed mode windows.  The new
value will take effect if the window is later made full screen.

C documentation : [glfwSetWindowAttrib]

# params
  - window - The window to set the attribute for.
  - attrib - A supported window attribute.
  - value - `GLFW_TRUE` or `GLFW_FALSE`.

# errors

Possible errors include [NOT_INITIALIZED], [INVALID_ENUM], [INVALID_VALUE], [PLATFORM_ERROR] and [FEATURE_UNAVAILABLE] (see remarks).

Calling [GetAttrib] will always return the latest
value, even if that value is ignored by the current mode of the window.

# wayland

The [GLFW_FLOATING] window attribute is
not supported.  Setting this will emit [FEATURE_UNAVAILABLE].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.3.

[GLFW_DECORATED]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_DECORATED_attrib
[GLFW_RESIZABLE]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_RESIZABLE_attrib
[GLFW_FLOATING]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_FLOATING_attrib
[GLFW_AUTO_ICONIFY]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_AUTO_ICONIFY_attrib
[GLFW_FOCUS_ON_SHOW]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_FOCUS_ON_SHOW_attrib
[GLFW_MOUSE_PASSTHROUGH]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_MOUSE_PASSTHROUGH_attrib
[glfwSetWindowAttrib]: https://www.glfw.org/docs/latest/group__window.html#gace2afda29b4116ec012e410a6819033e

[GLFW_FLOATING]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_FLOATING_attrib
*/
func (window *Window) SetAttrib(attrib Int, value Int) {
	_, err := ffi.CallFunction(
		cif_glfwSetWindowAttrib,
		func_glfwSetWindowAttrib,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&attrib),
			unsafe.Pointer(&value),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function sets the user-defined pointer of the specified window.  The
current value is retained until the window is destroyed.  The initial value
is `NULL`.

C documentation : [glfwSetWindowUserPointer]

# params
  - window - The window whose pointer to set.
  - pointer - The new value.

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function may be called from any thread.  Access is not
synchronized.

# since

Added in version 3.0.

[glfwSetWindowUserPointer]: https://www.glfw.org/docs/latest/group__window.html#ga3d2fc6026e690ab31a13f78bc9fd3651
*/
func (window *Window) SetUserPointer(pointer unsafe.Pointer) {
	_, err := ffi.CallFunction(
		cif_glfwSetWindowUserPointer,
		func_glfwSetWindowUserPointer,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&pointer),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function returns the current value of the user-defined pointer of the
specified window.  The initial value is `NULL`.

C documentation : [glfwGetWindowUserPointer]

# params
  - window - The window whose pointer to return.

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function may be called from any thread.  Access is not
synchronized.

# since

Added in version 3.0.

[glfwGetWindowUserPointer]: https://www.glfw.org/docs/latest/group__window.html#gae77a4add0d2023ca21ff1443ced01653
*/
func (window *Window) GetUserPointer() unsafe.Pointer {
	var result unsafe.Pointer
	_, err := ffi.CallFunction(
		cif_glfwGetWindowUserPointer,
		func_glfwGetWindowUserPointer,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the position callback of the specified window, which is
called when the window is moved.  The callback is provided with the
position, in screen coordinates, of the upper-left corner of the content
area of the window.

C documentation : [glfwSetWindowPosCallback]

# params
  - window - The window whose callback to set.
  - callback - The new callback, or `NULL` to remove the currently set
    callback.

# return

The previously set callback, or `NULL` if no callback was set or the
library had not been [initialized].

	void function_name(GLFWwindow* window, int xpos, int ypos)

For more information about the callback parameters, see the
[function pointer type].

# errors

Possible errors include [NOT_INITIALIZED].

# wayland

This callback will never be called, as there is no way for
an application to know its global position.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[initialized]: https://www.glfw.org/docs/latest/intro_guide.html#intro_init
[function pointer type]: https://www.glfw.org/docs/latest/group__window.html#gabe287973a21a8f927cde4db06b8dcbe9
[glfwSetWindowPosCallback]: https://www.glfw.org/docs/latest/group__window.html#ga08bdfbba88934f9c4f92fd757979ac74
*/
func (window *Window) SetPosCallback(callback WindowposCallback) WindowposCallback {
	var result WindowposCallback
	_, err := ffi.CallFunction(
		cif_glfwSetWindowPosCallback,
		func_glfwSetWindowPosCallback,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&callback),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the size callback of the specified window, which is
called when the window is resized.  The callback is provided with the size,
in screen coordinates, of the content area of the window.

C documentation : [glfwSetWindowSizeCallback]

# params
  - window - The window whose callback to set.
  - callback - The new callback, or `NULL` to remove the currently set
    callback.

# return

The previously set callback, or `NULL` if no callback was set or the
library had not been [initialized].

	void function_name(GLFWwindow* window, int width, int height)

For more information about the callback parameters, see the
[function pointer type].

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function must only be called from the main thread.

# since

Added in version 1.0.

[initialized]: https://www.glfw.org/docs/latest/intro_guide.html#intro_init
[function pointer type]: https://www.glfw.org/docs/latest/group__window.html#gaec0282944bb810f6f3163ec02da90350
[glfwSetWindowSizeCallback]: https://www.glfw.org/docs/latest/group__window.html#gad91b8b047a0c4c6033c38853864c34f8
*/
func (window *Window) SetSizeCallback(callback WindowsizeCallback) WindowsizeCallback {
	var result WindowsizeCallback
	_, err := ffi.CallFunction(
		cif_glfwSetWindowSizeCallback,
		func_glfwSetWindowSizeCallback,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&callback),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the close callback of the specified window, which is
called when the user attempts to close the window, for example by clicking
the close widget in the title bar.

The close flag is set before this callback is called, but you can modify it
at any time with [SetShouldClose].

The close callback is not triggered by [Destroy].

C documentation : [glfwSetWindowCloseCallback]

# params
  - window - The window whose callback to set.
  - callback - The new callback, or `NULL` to remove the currently set
    callback.

# return

The previously set callback, or `NULL` if no callback was set or the
library had not been [initialized].

	void function_name(GLFWwindow* window)

For more information about the callback parameters, see the
[function pointer type].

# errors

Possible errors include [NOT_INITIALIZED].

# macos

Selecting Quit from the application menu will trigger the
close callback for all windows.

# thread safety

This function must only be called from the main thread.

# since

Added in version 2.5.

[initialized]: https://www.glfw.org/docs/latest/intro_guide.html#intro_init
[function pointer type]: https://www.glfw.org/docs/latest/group__window.html#gabf859b936d80961b7d39013a9694cc3e
[glfwSetWindowCloseCallback]: https://www.glfw.org/docs/latest/group__window.html#gada646d775a7776a95ac000cfc1885331
*/
func (window *Window) SetCloseCallback(callback WindowcloseCallback) WindowcloseCallback {
	var result WindowcloseCallback
	_, err := ffi.CallFunction(
		cif_glfwSetWindowCloseCallback,
		func_glfwSetWindowCloseCallback,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&callback),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the refresh callback of the specified window, which is
called when the content area of the window needs to be redrawn, for example
if the window has been exposed after having been covered by another window.

On compositing window systems such as Aero, Compiz, Aqua or Wayland, where
the window contents are saved off-screen, this callback may be called only
very infrequently or never at all.

C documentation : [glfwSetWindowRefreshCallback]

# params
  - window - The window whose callback to set.
  - callback - The new callback, or `NULL` to remove the currently set
    callback.

# return

The previously set callback, or `NULL` if no callback was set or the
library had not been [initialized].

	void function_name(GLFWwindow* window);

For more information about the callback parameters, see the
[function pointer type].

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function must only be called from the main thread.

# since

Added in version 2.5.

[initialized]: https://www.glfw.org/docs/latest/intro_guide.html#intro_init
[function pointer type]: https://www.glfw.org/docs/latest/group__window.html#ga431663a1427d2eb3a273bc398b6737b5
[glfwSetWindowRefreshCallback]: https://www.glfw.org/docs/latest/group__window.html#ga1c5c7eb889c33c7f4d10dd35b327654e
*/
func (window *Window) SetRefreshCallback(callback WindowrefreshCallback) WindowrefreshCallback {
	var result WindowrefreshCallback
	_, err := ffi.CallFunction(
		cif_glfwSetWindowRefreshCallback,
		func_glfwSetWindowRefreshCallback,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&callback),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the focus callback of the specified window, which is
called when the window gains or loses input focus.

After the focus callback is called for a window that lost input focus,
synthetic key and mouse button release events will be generated for all such
that had been pressed.  For more information, see [SetKeyCallback]
and [SetMouseButtonCallback].

C documentation : [glfwSetWindowFocusCallback]

# params
  - window - The window whose callback to set.
  - callback - The new callback, or `NULL` to remove the currently set
    callback.

# return

The previously set callback, or `NULL` if no callback was set or the
library had not been [initialized].

	void function_name(GLFWwindow* window, int focused)

For more information about the callback parameters, see the
[function pointer type].

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[initialized]: https://www.glfw.org/docs/latest/intro_guide.html#intro_init
[function pointer type]: https://www.glfw.org/docs/latest/group__window.html#gabc58c47e9d93f6eb1862d615c3680f46
[glfwSetWindowFocusCallback]: https://www.glfw.org/docs/latest/group__window.html#gac2d83c4a10f071baf841f6730528e66c
*/
func (window *Window) SetFocusCallback(callback WindowfocusCallback) WindowfocusCallback {
	var result WindowfocusCallback
	_, err := ffi.CallFunction(
		cif_glfwSetWindowFocusCallback,
		func_glfwSetWindowFocusCallback,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&callback),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the iconification callback of the specified window, which
is called when the window is iconified or restored.

C documentation : [glfwSetWindowIconifyCallback]

# params
  - window - The window whose callback to set.
  - callback - The new callback, or `NULL` to remove the currently set
    callback.

# return

The previously set callback, or `NULL` if no callback was set or the
library had not been [initialized].

	void function_name(GLFWwindow* window, int iconified)

For more information about the callback parameters, see the
[function pointer type].

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[initialized]: https://www.glfw.org/docs/latest/intro_guide.html#intro_init
[function pointer type]: https://www.glfw.org/docs/latest/group__window.html#ga35c658cccba236f26e7adee0e25f6a4f
[glfwSetWindowIconifyCallback]: https://www.glfw.org/docs/latest/group__window.html#gac793e9efd255567b5fb8b445052cfd3e
*/
func (window *Window) SetIconifyCallback(callback WindowiconifyCallback) WindowiconifyCallback {
	var result WindowiconifyCallback
	_, err := ffi.CallFunction(
		cif_glfwSetWindowIconifyCallback,
		func_glfwSetWindowIconifyCallback,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&callback),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the maximization callback of the specified window, which
is called when the window is maximized or restored.

C documentation : [glfwSetWindowMaximizeCallback]

# params
  - window - The window whose callback to set.
  - callback - The new callback, or `NULL` to remove the currently set
    callback.

# return

The previously set callback, or `NULL` if no callback was set or the
library had not been [initialized].

	void function_name(GLFWwindow* window, int maximized)

For more information about the callback parameters, see the
[function pointer type].

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.3.

[initialized]: https://www.glfw.org/docs/latest/intro_guide.html#intro_init
[function pointer type]: https://www.glfw.org/docs/latest/group__window.html#ga3017196fdaec33ac3e095765176c2a90
[glfwSetWindowMaximizeCallback]: https://www.glfw.org/docs/latest/group__window.html#gacbe64c339fbd94885e62145563b6dc93
*/
func (window *Window) SetMaximizeCallback(callback WindowmaximizeCallback) WindowmaximizeCallback {
	var result WindowmaximizeCallback
	_, err := ffi.CallFunction(
		cif_glfwSetWindowMaximizeCallback,
		func_glfwSetWindowMaximizeCallback,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&callback),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the framebuffer resize callback of the specified window,
which is called when the framebuffer of the specified window is resized.

C documentation : [glfwSetFramebufferSizeCallback]

# params
  - window - The window whose callback to set.
  - callback - The new callback, or `NULL` to remove the currently set
    callback.

# return

The previously set callback, or `NULL` if no callback was set or the
library had not been [initialized].

	void function_name(GLFWwindow* window, int width, int height)

For more information about the callback parameters, see the
[function pointer type].

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[initialized]: https://www.glfw.org/docs/latest/intro_guide.html#intro_init
[function pointer type]: https://www.glfw.org/docs/latest/group__window.html#gae18026e294dde685ed2e5f759533144d
[glfwSetFramebufferSizeCallback]: https://www.glfw.org/docs/latest/group__window.html#gab3fb7c3366577daef18c0023e2a8591f
*/
func (window *Window) SetFramebufferSizeCallback(callback FramebuffersizeCallback) FramebuffersizeCallback {
	var result FramebuffersizeCallback
	_, err := ffi.CallFunction(
		cif_glfwSetFramebufferSizeCallback,
		func_glfwSetFramebufferSizeCallback,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&callback),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the window content scale callback of the specified window,
which is called when the content scale of the specified window changes.

C documentation : [glfwSetWindowContentScaleCallback]

# params
  - window - The window whose callback to set.
  - callback - The new callback, or `NULL` to remove the currently set
    callback.

# return

The previously set callback, or `NULL` if no callback was set or the
library had not been [initialized].

	void function_name(GLFWwindow* window, float xscale, float yscale)

For more information about the callback parameters, see the
[function pointer type].

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.3.

[initialized]: https://www.glfw.org/docs/latest/intro_guide.html#intro_init
[function pointer type]: https://www.glfw.org/docs/latest/group__window.html#ga77f288a2d04bb3c77c7d9615d08cf70e
[glfwSetWindowContentScaleCallback]: https://www.glfw.org/docs/latest/group__window.html#gaf2832ebb5aa6c252a2d261de002c92d6
*/
func (window *Window) SetContentScaleCallback(callback WindowcontentscaleCallback) WindowcontentscaleCallback {
	var result WindowcontentscaleCallback
	_, err := ffi.CallFunction(
		cif_glfwSetWindowContentScaleCallback,
		func_glfwSetWindowContentScaleCallback,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&callback),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function processes only those events that are already in the event
queue and then returns immediately.  Processing events will cause the window
and input callbacks associated with those events to be called.

On some platforms, a window move, resize or menu operation will cause event
processing to block.  This is due to how event processing is designed on
those platforms.  You can use the
[window refresh callback] to redraw the contents of
your window when necessary during such operations.

Do not assume that callbacks you set will _only_ be called in response to
event processing functions like this one.  While it is necessary to poll for
events, window systems that require GLFW to register callbacks of its own
can pass events to GLFW in response to many window system function calls.
GLFW will pass those events on to the application callbacks before
returning.

Event processing is not required for joystick input to work.

C documentation : [glfwPollEvents]

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# reentrancy

This function must not be called from a callback.

# thread safety

This function must only be called from the main thread.

# since

Added in version 1.0.

[window refresh callback]: https://www.glfw.org/docs/latest/window_guide.html#window_refresh
[glfwPollEvents]: https://www.glfw.org/docs/latest/group__window.html#ga37bd57223967b4211d60ca1a0bf3c832
*/
func PollEvents() {
	_, err := ffi.CallFunction(
		cif_glfwPollEvents,
		func_glfwPollEvents,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function puts the calling thread to sleep until at least one event is
available in the event queue.  Once one or more events are available,
it behaves exactly like [PollEvents], i.e. the events in the queue
are processed and the function then returns immediately.  Processing events
will cause the window and input callbacks associated with those events to be
called.

Since not all events are associated with callbacks, this function may return
without a callback having been called even if you are monitoring all
callbacks.

On some platforms, a window move, resize or menu operation will cause event
processing to block.  This is due to how event processing is designed on
those platforms.  You can use the
[window refresh callback] to redraw the contents of
your window when necessary during such operations.

Do not assume that callbacks you set will _only_ be called in response to
event processing functions like this one.  While it is necessary to poll for
events, window systems that require GLFW to register callbacks of its own
can pass events to GLFW in response to many window system function calls.
GLFW will pass those events on to the application callbacks before
returning.

Event processing is not required for joystick input to work.

C documentation : [glfwWaitEvents]

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# reentrancy

This function must not be called from a callback.

# thread safety

This function must only be called from the main thread.

# since

Added in version 2.5.

[window refresh callback]: https://www.glfw.org/docs/latest/window_guide.html#window_refresh
[glfwWaitEvents]: https://www.glfw.org/docs/latest/group__window.html#ga554e37d781f0a997656c26b2c56c835e
*/
func WaitEvents() {
	_, err := ffi.CallFunction(
		cif_glfwWaitEvents,
		func_glfwWaitEvents,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function puts the calling thread to sleep until at least one event is
available in the event queue, or until the specified timeout is reached.  If
one or more events are available, it behaves exactly like [PollEvents], i.e. the events in the queue are processed and the function
then returns immediately.  Processing events will cause the window and input
callbacks associated with those events to be called.

The timeout value must be a positive finite number.

Since not all events are associated with callbacks, this function may return
without a callback having been called even if you are monitoring all
callbacks.

On some platforms, a window move, resize or menu operation will cause event
processing to block.  This is due to how event processing is designed on
those platforms.  You can use the
[window refresh callback] to redraw the contents of
your window when necessary during such operations.

Do not assume that callbacks you set will _only_ be called in response to
event processing functions like this one.  While it is necessary to poll for
events, window systems that require GLFW to register callbacks of its own
can pass events to GLFW in response to many window system function calls.
GLFW will pass those events on to the application callbacks before
returning.

Event processing is not required for joystick input to work.

C documentation : [glfwWaitEventsTimeout]

# params
  - timeout - The maximum amount of time, in seconds, to wait.

# errors

Possible errors include [NOT_INITIALIZED], [INVALID_VALUE] and [PLATFORM_ERROR].

# reentrancy

This function must not be called from a callback.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.2.

[window refresh callback]: https://www.glfw.org/docs/latest/window_guide.html#window_refresh
[glfwWaitEventsTimeout]: https://www.glfw.org/docs/latest/group__window.html#ga605a178db92f1a7f1a925563ef3ea2cf
*/
func WaitEventsTimeout(timeout Double) {
	_, err := ffi.CallFunction(
		cif_glfwWaitEventsTimeout,
		func_glfwWaitEventsTimeout,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&timeout),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function posts an empty event from the current thread to the event
queue, causing [WaitEvents] or [WaitEventsTimeout] to return.

C documentation : [glfwPostEmptyEvent]

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# thread safety

This function may be called from any thread.

# since

Added in version 3.1.

[glfwPostEmptyEvent]: https://www.glfw.org/docs/latest/group__window.html#gab5997a25187e9fd5c6f2ecbbc8dfd7e9
*/
func PostEmptyEvent() {
	_, err := ffi.CallFunction(
		cif_glfwPostEmptyEvent,
		func_glfwPostEmptyEvent,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function returns the value of an input option for the specified window.
The mode must be one of [CURSOR], [STICKY_KEYS],
[STICKY_MOUSE_BUTTONS], [LOCK_KEY_MODS] or
[RAW_MOUSE_MOTION].

C documentation : [glfwGetInputMode]

# params
  - window - The window to query.
  - mode - One of `GLFW_CURSOR`, `GLFW_STICKY_KEYS`,
    `GLFW_STICKY_MOUSE_BUTTONS`, `GLFW_LOCK_KEY_MODS` or
    `GLFW_RAW_MOUSE_MOTION`.

# errors

Possible errors include [NOT_INITIALIZED] and [INVALID_ENUM].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[glfwGetInputMode]: https://www.glfw.org/docs/latest/group__input.html#gaf5b859dbe19bdf434e42695ea45cc5f4
*/
func (window *Window) GetInputMode(mode Int) Int {
	var result Int
	_, err := ffi.CallFunction(
		cif_glfwGetInputMode,
		func_glfwGetInputMode,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&mode),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets an input mode option for the specified window.  The mode
must be one of [CURSOR], [STICKY_KEYS],
[STICKY_MOUSE_BUTTONS], [LOCK_KEY_MODS] or
[RAW_MOUSE_MOTION].

If the mode is `GLFW_CURSOR`, the value must be one of the following cursor
modes:
- `GLFW_CURSOR_NORMAL` makes the cursor visible and behaving normally.
- `GLFW_CURSOR_HIDDEN` makes the cursor invisible when it is over the
content area of the window but does not restrict the cursor from leaving.
- `GLFW_CURSOR_DISABLED` hides and grabs the cursor, providing virtual
and unlimited cursor movement.  This is useful for implementing for
example 3D camera controls.
- `GLFW_CURSOR_CAPTURED` makes the cursor visible and confines it to the
content area of the window.

If the mode is `GLFW_STICKY_KEYS`, the value must be either `GLFW_TRUE` to
enable sticky keys, or `GLFW_FALSE` to disable it.  If sticky keys are
enabled, a key press will ensure that [GetKey] returns `GLFW_PRESS`
the next time it is called even if the key had been released before the
call.  This is useful when you are only interested in whether keys have been
pressed but not when or in which order.

If the mode is `GLFW_STICKY_MOUSE_BUTTONS`, the value must be either
`GLFW_TRUE` to enable sticky mouse buttons, or `GLFW_FALSE` to disable it.
If sticky mouse buttons are enabled, a mouse button press will ensure that
[GetMouseButton] returns `GLFW_PRESS` the next time it is called even
if the mouse button had been released before the call.  This is useful when
you are only interested in whether mouse buttons have been pressed but not
when or in which order.

If the mode is `GLFW_LOCK_KEY_MODS`, the value must be either `GLFW_TRUE` to
enable lock key modifier bits, or `GLFW_FALSE` to disable them.  If enabled,
callbacks that receive modifier bits will also have the [MOD_CAPS_LOCK] bit set when the event was generated with Caps Lock on,
and the [MOD_NUM_LOCK] bit when Num Lock was on.

If the mode is `GLFW_RAW_MOUSE_MOTION`, the value must be either `GLFW_TRUE`
to enable raw (unscaled and unaccelerated) mouse motion when the cursor is
disabled, or `GLFW_FALSE` to disable it.  If raw motion is not supported,
attempting to set this will emit [FEATURE_UNAVAILABLE].  Call [RawMouseMotionSupported] to check for support.

C documentation : [glfwSetInputMode]

# params
  - window - The window whose input mode to set.
  - mode - One of `GLFW_CURSOR`, `GLFW_STICKY_KEYS`,
    `GLFW_STICKY_MOUSE_BUTTONS`, `GLFW_LOCK_KEY_MODS` or
    `GLFW_RAW_MOUSE_MOTION`.
  - value - The new value of the specified input mode.

# errors

Possible errors include [NOT_INITIALIZED], [INVALID_ENUM], [PLATFORM_ERROR] and [FEATURE_UNAVAILABLE] (see above).

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.  Replaces `glfwEnable` and `glfwDisable`.

[glfwSetInputMode]: https://www.glfw.org/docs/latest/group__input.html#gaa92336e173da9c8834558b54ee80563b
*/
func (window *Window) SetInputMode(mode Int, value Int) {
	_, err := ffi.CallFunction(
		cif_glfwSetInputMode,
		func_glfwSetInputMode,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&mode),
			unsafe.Pointer(&value),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function returns whether raw mouse motion is supported on the current
system.  This status does not change after GLFW has been initialized so you
only need to check this once.  If you attempt to enable raw motion on
a system that does not support it, [PLATFORM_ERROR] will be emitted.

Raw mouse motion is closer to the actual motion of the mouse across
a surface.  It is not affected by the scaling and acceleration applied to
the motion of the desktop cursor.  That processing is suitable for a cursor
while raw motion is better for controlling for example a 3D camera.  Because
of this, raw mouse motion is only provided when the cursor is disabled.

C documentation : [glfwRawMouseMotionSupported]

# return

`GLFW_TRUE` if raw mouse motion is supported on the current machine,
or `GLFW_FALSE` otherwise.

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.3.

[glfwRawMouseMotionSupported]: https://www.glfw.org/docs/latest/group__input.html#gae4ee0dbd0d256183e1ea4026d897e1c2
*/
func RawMouseMotionSupported() Int {
	var result Int
	_, err := ffi.CallFunction(
		cif_glfwRawMouseMotionSupported,
		func_glfwRawMouseMotionSupported,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function returns the name of the specified printable key, encoded as
UTF-8.  This is typically the character that key would produce without any
modifier keys, intended for displaying key bindings to the user.  For dead
keys, it is typically the diacritic it would add to a character.

__Do not use this function__ for [text input].  You will
break text input for many languages even if it happens to work for yours.

If the key is `GLFW_KEY_UNKNOWN`, the scancode is used to identify the key,
otherwise the scancode is ignored.  If you specify a non-printable key, or
`GLFW_KEY_UNKNOWN` and a scancode that maps to a non-printable key, this
function returns `NULL` but does not emit an error.

This behavior allows you to always pass in the arguments in the
[key callback] without modification.

The printable keys are:
- `GLFW_KEY_APOSTROPHE`
- `GLFW_KEY_COMMA`
- `GLFW_KEY_MINUS`
- `GLFW_KEY_PERIOD`
- `GLFW_KEY_SLASH`
- `GLFW_KEY_SEMICOLON`
- `GLFW_KEY_EQUAL`
- `GLFW_KEY_LEFT_BRACKET`
- `GLFW_KEY_RIGHT_BRACKET`
- `GLFW_KEY_BACKSLASH`
- `GLFW_KEY_WORLD_1`
- `GLFW_KEY_WORLD_2`
- `GLFW_KEY_0` to `GLFW_KEY_9`
- `GLFW_KEY_A` to `GLFW_KEY_Z`
- `GLFW_KEY_KP_0` to `GLFW_KEY_KP_9`
- `GLFW_KEY_KP_DECIMAL`
- `GLFW_KEY_KP_DIVIDE`
- `GLFW_KEY_KP_MULTIPLY`
- `GLFW_KEY_KP_SUBTRACT`
- `GLFW_KEY_KP_ADD`
- `GLFW_KEY_KP_EQUAL`

Names for printable keys depend on keyboard layout, while names for
non-printable keys are the same across layouts but depend on the application
language and should be localized along with other user interface text.

C documentation : [glfwGetKeyName]

# params
  - key - The key to query, or `GLFW_KEY_UNKNOWN`.
  - scancode - The scancode of the key to query.

# return

The UTF-8 encoded, layout-specific name of the key, or `NULL`.

# errors

Possible errors include [NOT_INITIALIZED], [INVALID_VALUE], [INVALID_ENUM] and [PLATFORM_ERROR].

The contents of the returned string may change when a keyboard
layout change event is received.

# pointer lifetime

The returned string is allocated and freed by GLFW.  You
should not free it yourself.  It is valid until the library is terminated.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.2.

[text input]: https://www.glfw.org/docs/latest/input_guide.html#input_char
[key callback]: https://www.glfw.org/docs/latest/input_guide.html#input_key
[glfwGetKeyName]: https://www.glfw.org/docs/latest/group__input.html#gaeaed62e69c3bd62b7ff8f7b19913ce4f
*/
func GetKeyName(key Int, scancode Int) string {
	var result *byte
	_, err := ffi.CallFunction(
		cif_glfwGetKeyName,
		func_glfwGetKeyName,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&key),
			unsafe.Pointer(&scancode),
		},
	)
	if err != nil {
		panic(err)
	}
	return goString(result)
}

/*
This function returns the platform-specific scancode of the specified key.

If the specified [key token] corresponds to a physical key not
supported on the current platform then this method will return `-1`.
Calling this function with anything other than a key token will return `-1`
and generate a [INVALID_ENUM] error.

C documentation : [glfwGetKeyScancode]

# params
  - key - Any [key token].

# return

The platform-specific scancode for the key, or `-1` if the key is
not supported on the current platform or an [error]
occurred.

# errors

Possible errors include [NOT_INITIALIZED] and [INVALID_ENUM].

# thread safety

This function may be called from any thread.

# since

Added in version 3.3.

[key token]: https://www.glfw.org/docs/latest/group__keys.html
[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwGetKeyScancode]: https://www.glfw.org/docs/latest/group__input.html#ga67ddd1b7dcbbaff03e4a76c0ea67103a

[key token]: https://www.glfw.org/docs/latest/group__keys.html
*/
func GetKeyScancode(key Int) Int {
	var result Int
	_, err := ffi.CallFunction(
		cif_glfwGetKeyScancode,
		func_glfwGetKeyScancode,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&key),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
window.

This function returns the last state reported for the specified key to the
specified window.  The returned state is one of `GLFW_PRESS` or
`GLFW_RELEASE`.  The action `GLFW_REPEAT` is only reported to the key callback.

If the [STICKY_KEYS] input mode is enabled, this function returns
`GLFW_PRESS` the first time you call it for a key that was pressed, even if
that key has already been released.

The key functions deal with physical keys, with [key tokens]
named after their use on the standard US keyboard layout.  If you want to
input text, use the Unicode character callback instead.

The [modifier key bit masks] are not key tokens and cannot be
used with this function.

__Do not use this function__ to implement [text input].

C documentation : [glfwGetKey]

# params
  - window - The desired window.
  - key - The desired [keyboard key].  `GLFW_KEY_UNKNOWN` is
    not a valid key for this function.

# return

One of `GLFW_PRESS` or `GLFW_RELEASE`.

# errors

Possible errors include [NOT_INITIALIZED] and [INVALID_ENUM].

# thread safety

This function must only be called from the main thread.

# since

Added in version 1.0.

[key tokens]: https://www.glfw.org/docs/latest/group__keys.html
[modifier key bit masks]: https://www.glfw.org/docs/latest/group__mods.html
[text input]: https://www.glfw.org/docs/latest/input_guide.html#input_char
[keyboard key]: https://www.glfw.org/docs/latest/group__keys.html
[glfwGetKey]: https://www.glfw.org/docs/latest/group__input.html#gadd341da06bc8d418b4dc3a3518af9ad2
*/
func (window *Window) GetKey(key Int) Int {
	var result Int
	_, err := ffi.CallFunction(
		cif_glfwGetKey,
		func_glfwGetKey,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&key),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
window.

This function returns the last state reported for the specified mouse button
to the specified window.  The returned state is one of `GLFW_PRESS` or
`GLFW_RELEASE`.

If the [STICKY_MOUSE_BUTTONS] input mode is enabled, this function
returns `GLFW_PRESS` the first time you call it for a mouse button that was
pressed, even if that mouse button has already been released.

C documentation : [glfwGetMouseButton]

# params
  - window - The desired window.
  - button - The desired [mouse button].

# return

One of `GLFW_PRESS` or `GLFW_RELEASE`.

# errors

Possible errors include [NOT_INITIALIZED] and [INVALID_ENUM].

# thread safety

This function must only be called from the main thread.

# since

Added in version 1.0.

[mouse button]: https://www.glfw.org/docs/latest/group__buttons.html
[glfwGetMouseButton]: https://www.glfw.org/docs/latest/group__input.html#gac1473feacb5996c01a7a5a33b5066704
*/
func (window *Window) GetMouseButton(button Int) Int {
	var result Int
	_, err := ffi.CallFunction(
		cif_glfwGetMouseButton,
		func_glfwGetMouseButton,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&button),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
the window.

This function returns the position of the cursor, in screen coordinates,
relative to the upper-left corner of the content area of the specified
window.

If the cursor is disabled (with `GLFW_CURSOR_DISABLED`) then the cursor
position is unbounded and limited only by the minimum and maximum values of
a `double`.

The coordinate can be converted to their integer equivalents with the
`floor` function.  Casting directly to an integer type works for positive
coordinates, but fails for negative ones.

Any or all of the position arguments may be `NULL`.  If an error occurs, all
non-`NULL` position arguments will be set to zero.

C documentation : [glfwGetCursorPos]

# params
  - window - The desired window.
  - xpos (out) - Where to store the cursor x-coordinate, relative to the
    left edge of the content area, or `NULL`.
  - ypos (out) - Where to store the cursor y-coordinate, relative to the to
    top edge of the content area, or `NULL`.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.  Replaces `glfwGetMousePos`.

[glfwGetCursorPos]: https://www.glfw.org/docs/latest/group__input.html#ga01d37b6c40133676b9cea60ca1d7c0cc
*/
func (window *Window) GetCursorPos() (Double, Double) {
	var xpos Double
	var ypos Double
	_, err := ffi.CallFunction(
		cif_glfwGetCursorPos,
		func_glfwGetCursorPos,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(new(&xpos)),
			unsafe.Pointer(new(&ypos)),
		},
	)
	if err != nil {
		panic(err)
	}
	return xpos, ypos
}

/*
window.

This function sets the position, in screen coordinates, of the cursor
relative to the upper-left corner of the content area of the specified
window.  The window must have input focus.  If the window does not have
input focus when this function is called, it fails silently.

__Do not use this function__ to implement things like camera controls.  GLFW
already provides the `GLFW_CURSOR_DISABLED` cursor mode that hides the
cursor, transparently re-centers it and provides unconstrained cursor
motion.  See [SetInputMode] for more information.

If the cursor mode is `GLFW_CURSOR_DISABLED` then the cursor position is
unconstrained and limited only by the minimum and maximum values of
a `double`.

C documentation : [glfwSetCursorPos]

# params
  - window - The desired window.
  - xpos - The desired x-coordinate, relative to the left edge of the
    content area.
  - ypos - The desired y-coordinate, relative to the top edge of the
    content area.

# errors

Possible errors include [NOT_INITIALIZED], [PLATFORM_ERROR] and [FEATURE_UNAVAILABLE] (see remarks).

# wayland

This function will only work when the cursor mode is
`GLFW_CURSOR_DISABLED`, otherwise it will emit [FEATURE_UNAVAILABLE].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.  Replaces `glfwSetMousePos`.

[glfwSetCursorPos]: https://www.glfw.org/docs/latest/group__input.html#ga04b03af936d906ca123c8f4ee08b39e7
*/
func (window *Window) SetCursorPos(xpos Double, ypos Double) {
	_, err := ffi.CallFunction(
		cif_glfwSetCursorPos,
		func_glfwSetCursorPos,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&xpos),
			unsafe.Pointer(&ypos),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
Creates a new custom cursor image that can be set for a window with [SetCursor].  The cursor can be destroyed with [Destroy].
Any remaining cursors are destroyed by [Terminate].

The pixels are 32-bit, little-endian, non-premultiplied RGBA, i.e. eight
bits per channel with the red channel first.  They are arranged canonically
as packed sequential rows, starting from the top-left corner.

The cursor hotspot is specified in pixels, relative to the upper-left corner
of the cursor image.  Like all other coordinate systems in GLFW, the X-axis
points to the right and the Y-axis points down.

C documentation : [glfwCreateCursor]

# params
  - image - The desired cursor image.
  - xhot - The desired x-coordinate, in pixels, of the cursor hotspot.
  - yhot - The desired y-coordinate, in pixels, of the cursor hotspot.

# return

The handle of the created cursor, or `NULL` if an
[error] occurred.

# errors

Possible errors include [NOT_INITIALIZED], [INVALID_VALUE] and [PLATFORM_ERROR].

# pointer lifetime

The specified image data is copied before this function
returns.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.1.

[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwCreateCursor]: https://www.glfw.org/docs/latest/group__input.html#ga556f604f73af156c0db0e97c081373c3
*/
func (image *Image) CreateCursor(xhot Int, yhot Int) *Cursor {
	var result *Cursor
	_, err := ffi.CallFunction(
		cif_glfwCreateCursor,
		func_glfwCreateCursor,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&image),
			unsafe.Pointer(&xhot),
			unsafe.Pointer(&yhot),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
Returns a cursor with a standard shape, that can be set for a window with
[SetCursor].  The images for these cursors come from the system
cursor theme and their exact appearance will vary between platforms.

Most of these shapes are guaranteed to exist on every supported platform but
a few may not be present.  See the table below for details.

Cursor shape                   | Windows | macOS | X11    | Wayland
------------------------------ | ------- | ----- | ------ | -------
[ARROW_CURSOR]         | Yes     | Yes   | Yes    | Yes
[IBEAM_CURSOR]         | Yes     | Yes   | Yes    | Yes
[CROSSHAIR_CURSOR]     | Yes     | Yes   | Yes    | Yes
[POINTING_HAND_CURSOR] | Yes     | Yes   | Yes    | Yes
[RESIZE_EW_CURSOR]     | Yes     | Yes   | Yes    | Yes
[RESIZE_NS_CURSOR]     | Yes     | Yes   | Yes    | Yes
[RESIZE_NWSE_CURSOR]   | Yes     | Yes<sup>1</sup> | Maybe<sup>2</sup> | Maybe<sup>2</sup>
[RESIZE_NESW_CURSOR]   | Yes     | Yes<sup>1</sup> | Maybe<sup>2</sup> | Maybe<sup>2</sup>
[RESIZE_ALL_CURSOR]    | Yes     | Yes   | Yes    | Yes
[NOT_ALLOWED_CURSOR]   | Yes     | Yes   | Maybe<sup>2</sup> | Maybe<sup>2</sup>

1) This uses a private system API and may fail in the future.

2) This uses a newer standard that not all cursor themes support.

If the requested shape is not available, this function emits a [CURSOR_UNAVAILABLE] error and returns `NULL`.

C documentation : [glfwCreateStandardCursor]

# params
  - shape - One of the [standard shapes].

# return

A new cursor ready to use or `NULL` if an
[error] occurred.

# errors

Possible errors include [NOT_INITIALIZED], [INVALID_ENUM], [CURSOR_UNAVAILABLE] and [PLATFORM_ERROR].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.1.

[standard shapes]: https://www.glfw.org/docs/latest/group__shapes.html
[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwCreateStandardCursor]: https://www.glfw.org/docs/latest/group__input.html#gaf2fb2eb2c9dd842d1cef8a34e3c6403e
*/
func CreateStandardCursor(shape Int) *Cursor {
	var result *Cursor
	_, err := ffi.CallFunction(
		cif_glfwCreateStandardCursor,
		func_glfwCreateStandardCursor,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&shape),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function destroys a cursor previously created with [CreateCursor].  Any remaining cursors will be destroyed by [Terminate].

If the specified cursor is current for any window, that window will be
reverted to the default cursor.  This does not affect the cursor mode.

C documentation : [glfwDestroyCursor]

# params
  - cursor - The cursor object to destroy.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# reentrancy

This function must not be called from a callback.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.1.

[glfwDestroyCursor]: https://www.glfw.org/docs/latest/group__input.html#ga81b952cd1764274d0db7fb3c5a79ba6a
*/
func (cursor *Cursor) Destroy() {
	_, err := ffi.CallFunction(
		cif_glfwDestroyCursor,
		func_glfwDestroyCursor,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&cursor),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function sets the cursor image to be used when the cursor is over the
content area of the specified window.  The set cursor will only be visible
when the [cursor mode] of the window is
`GLFW_CURSOR_NORMAL`.

On some platforms, the set cursor may not be visible unless the window also
has input focus.

C documentation : [glfwSetCursor]

# params
  - window - The window to set the cursor for.
  - cursor - The cursor to set, or `NULL` to switch back to the default
    arrow cursor.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.1.

[cursor mode]: https://www.glfw.org/docs/latest/input_guide.html#cursor_mode
[glfwSetCursor]: https://www.glfw.org/docs/latest/group__input.html#gad3b4f38c8d5dae036bc8fa959e18343e
*/
func (window *Window) SetCursor(cursor *Cursor) {
	_, err := ffi.CallFunction(
		cif_glfwSetCursor,
		func_glfwSetCursor,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&cursor),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function sets the key callback of the specified window, which is called
when a key is pressed, repeated or released.

The key functions deal with physical keys, with layout independent
[key tokens] named after their values in the standard US keyboard
layout.  If you want to input text, use the
[character callback] instead.

When a window loses input focus, it will generate synthetic key release
events for all pressed keys with associated key tokens.  You can tell these
events from user-generated events by the fact that the synthetic ones are
generated after the focus loss event has been processed, i.e. after the
[window focus callback] has been called.

The scancode of a key is specific to that platform or sometimes even to that
machine.  Scancodes are intended to allow users to bind keys that don't have
a GLFW key token.  Such keys have `key` set to `GLFW_KEY_UNKNOWN`, their
state is not saved and so it cannot be queried with [GetKey].

Sometimes GLFW needs to generate synthetic key events, in which case the
scancode may be zero.

C documentation : [glfwSetKeyCallback]

# params
  - window - The window whose callback to set.
  - callback - The new key callback, or `NULL` to remove the currently
    set callback.

# return

The previously set callback, or `NULL` if no callback was set or the
library had not been [initialized].

	void function_name(GLFWwindow* window, int key, int scancode, int action, int mods)

For more information about the callback parameters, see the
[function pointer type].

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function must only be called from the main thread.

# since

Added in version 1.0.

[key tokens]: https://www.glfw.org/docs/latest/group__keys.html
[character callback]: https://www.glfw.org/docs/latest/group__input.html#gab25c4a220fd8f5717718dbc487828996
[window focus callback]: https://www.glfw.org/docs/latest/group__window.html#gac2d83c4a10f071baf841f6730528e66c
[initialized]: https://www.glfw.org/docs/latest/intro_guide.html#intro_init
[function pointer type]: https://www.glfw.org/docs/latest/group__input.html#ga5bd751b27b90f865d2ea613533f0453c
[glfwSetKeyCallback]: https://www.glfw.org/docs/latest/group__input.html#ga1caf18159767e761185e49a3be019f8d
*/
func (window *Window) SetKeyCallback(callback KeyCallback) KeyCallback {
	var result KeyCallback
	_, err := ffi.CallFunction(
		cif_glfwSetKeyCallback,
		func_glfwSetKeyCallback,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&callback),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the character callback of the specified window, which is
called when a Unicode character is input.

The character callback is intended for Unicode text input.  As it deals with
characters, it is keyboard layout dependent, whereas the
[key callback] is not.  Characters do not map 1:1
to physical keys, as a key may produce zero, one or more characters.  If you
want to know whether a specific physical key was pressed or released, see
the key callback instead.

The character callback behaves as system text input normally does and will
not be called if modifier keys are held down that would prevent normal text
input on that platform, for example a Super (Command) key on macOS or Alt key
on Windows.

C documentation : [glfwSetCharCallback]

# params
  - window - The window whose callback to set.
  - callback - The new callback, or `NULL` to remove the currently set
    callback.

# return

The previously set callback, or `NULL` if no callback was set or the
library had not been [initialized].

	void function_name(GLFWwindow* window, unsigned int codepoint)

For more information about the callback parameters, see the
[function pointer type].

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function must only be called from the main thread.

# since

Added in version 2.4.

[key callback]: https://www.glfw.org/docs/latest/group__input.html#ga1caf18159767e761185e49a3be019f8d
[initialized]: https://www.glfw.org/docs/latest/intro_guide.html#intro_init
[function pointer type]: https://www.glfw.org/docs/latest/group__input.html#ga1ab90a55cf3f58639b893c0f4118cb6e
[glfwSetCharCallback]: https://www.glfw.org/docs/latest/group__input.html#gab25c4a220fd8f5717718dbc487828996
*/
func (window *Window) SetCharCallback(callback CharCallback) CharCallback {
	var result CharCallback
	_, err := ffi.CallFunction(
		cif_glfwSetCharCallback,
		func_glfwSetCharCallback,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&callback),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the character with modifiers callback of the specified
window, which is called when a Unicode character is input regardless of what
modifier keys are used.

The character with modifiers callback is intended for implementing custom
Unicode character input.  For regular Unicode text input, see the
[character callback].  Like the character
callback, the character with modifiers callback deals with characters and is
keyboard layout dependent.  Characters do not map 1:1 to physical keys, as
a key may produce zero, one or more characters.  If you want to know whether
a specific physical key was pressed or released, see the
[key callback] instead.

C documentation : [glfwSetCharModsCallback]

# params
  - window - The window whose callback to set.
  - callback - The new callback, or `NULL` to remove the currently set
    callback.

# return

The previously set callback, or `NULL` if no callback was set or an
[error] occurred.

	void function_name(GLFWwindow* window, unsigned int codepoint, int mods)

For more information about the callback parameters, see the
[function pointer type].

# deprecated

Scheduled for removal in version 4.0.

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.1.

[character callback]: https://www.glfw.org/docs/latest/group__input.html#gab25c4a220fd8f5717718dbc487828996
[key callback]: https://www.glfw.org/docs/latest/group__input.html#ga1caf18159767e761185e49a3be019f8d
[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[function pointer type]: https://www.glfw.org/docs/latest/group__input.html#gac3cf64f90b6219c05ac7b7822d5a4b8f
[glfwSetCharModsCallback]: https://www.glfw.org/docs/latest/group__input.html#ga0b7f4ad13c2b17435ff13b6dcfb4e43c
*/
func (window *Window) SetCharModsCallback(callback CharmodsCallback) CharmodsCallback {
	var result CharmodsCallback
	_, err := ffi.CallFunction(
		cif_glfwSetCharModsCallback,
		func_glfwSetCharModsCallback,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&callback),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the mouse button callback of the specified window, which
is called when a mouse button is pressed or released.

When a window loses input focus, it will generate synthetic mouse button
release events for all pressed mouse buttons.  You can tell these events
from user-generated events by the fact that the synthetic ones are generated
after the focus loss event has been processed, i.e. after the
[window focus callback] has been called.

C documentation : [glfwSetMouseButtonCallback]

# params
  - window - The window whose callback to set.
  - callback - The new callback, or `NULL` to remove the currently set
    callback.

# return

The previously set callback, or `NULL` if no callback was set or the
library had not been [initialized].

	void function_name(GLFWwindow* window, int button, int action, int mods)

For more information about the callback parameters, see the
[function pointer type].

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function must only be called from the main thread.

# since

Added in version 1.0.

[window focus callback]: https://www.glfw.org/docs/latest/group__window.html#gac2d83c4a10f071baf841f6730528e66c
[initialized]: https://www.glfw.org/docs/latest/intro_guide.html#intro_init
[function pointer type]: https://www.glfw.org/docs/latest/group__input.html#ga0184dcb59f6d85d735503dcaae809727
[glfwSetMouseButtonCallback]: https://www.glfw.org/docs/latest/group__input.html#ga6ab84420974d812bee700e45284a723c
*/
func (window *Window) SetMouseButtonCallback(callback MousebuttonCallback) MousebuttonCallback {
	var result MousebuttonCallback
	_, err := ffi.CallFunction(
		cif_glfwSetMouseButtonCallback,
		func_glfwSetMouseButtonCallback,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&callback),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the cursor position callback of the specified window,
which is called when the cursor is moved.  The callback is provided with the
position, in screen coordinates, relative to the upper-left corner of the
content area of the window.

C documentation : [glfwSetCursorPosCallback]

# params
  - window - The window whose callback to set.
  - callback - The new callback, or `NULL` to remove the currently set
    callback.

# return

The previously set callback, or `NULL` if no callback was set or the
library had not been [initialized].

	void function_name(GLFWwindow* window, double xpos, double ypos);

For more information about the callback parameters, see the
[function pointer type].

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.  Replaces `glfwSetMousePosCallback`.

[initialized]: https://www.glfw.org/docs/latest/intro_guide.html#intro_init
[function pointer type]: https://www.glfw.org/docs/latest/group__input.html#gad6fae41b3ac2e4209aaa87b596c57f68
[glfwSetCursorPosCallback]: https://www.glfw.org/docs/latest/group__input.html#gac1f879ab7435d54d4d79bb469fe225d7
*/
func (window *Window) SetCursorPosCallback(callback CursorposCallback) CursorposCallback {
	var result CursorposCallback
	_, err := ffi.CallFunction(
		cif_glfwSetCursorPosCallback,
		func_glfwSetCursorPosCallback,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&callback),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the cursor boundary crossing callback of the specified
window, which is called when the cursor enters or leaves the content area of
the window.

C documentation : [glfwSetCursorEnterCallback]

# params
  - window - The window whose callback to set.
  - callback - The new callback, or `NULL` to remove the currently set
    callback.

# return

The previously set callback, or `NULL` if no callback was set or the
library had not been [initialized].

	void function_name(GLFWwindow* window, int entered)

For more information about the callback parameters, see the
[function pointer type].

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[initialized]: https://www.glfw.org/docs/latest/intro_guide.html#intro_init
[function pointer type]: https://www.glfw.org/docs/latest/group__input.html#gaa93dc4818ac9ab32532909d53a337cbe
[glfwSetCursorEnterCallback]: https://www.glfw.org/docs/latest/group__input.html#gad27f8ad0142c038a281466c0966817d8
*/
func (window *Window) SetCursorEnterCallback(callback CursorenterCallback) CursorenterCallback {
	var result CursorenterCallback
	_, err := ffi.CallFunction(
		cif_glfwSetCursorEnterCallback,
		func_glfwSetCursorEnterCallback,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&callback),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the scroll callback of the specified window, which is
called when a scrolling device is used, such as a mouse wheel or scrolling
area of a touchpad.

The scroll callback receives all scrolling input, like that from a mouse
wheel or a touchpad scrolling area.

C documentation : [glfwSetScrollCallback]

# params
  - window - The window whose callback to set.
  - callback - The new scroll callback, or `NULL` to remove the
    currently set callback.

# return

The previously set callback, or `NULL` if no callback was set or the
library had not been [initialized].

	void function_name(GLFWwindow* window, double xoffset, double yoffset)

For more information about the callback parameters, see the
[function pointer type].

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.  Replaces `glfwSetMouseWheelCallback`.

[initialized]: https://www.glfw.org/docs/latest/intro_guide.html#intro_init
[function pointer type]: https://www.glfw.org/docs/latest/group__input.html#gaf656112c33de3efdb227fa58f0134cf5
[glfwSetScrollCallback]: https://www.glfw.org/docs/latest/group__input.html#ga571e45a030ae4061f746ed56cb76aede
*/
func (window *Window) SetScrollCallback(callback ScrollCallback) ScrollCallback {
	var result ScrollCallback
	_, err := ffi.CallFunction(
		cif_glfwSetScrollCallback,
		func_glfwSetScrollCallback,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&callback),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the path drop callback of the specified window, which is
called when one or more dragged paths are dropped on the window.

Because the path array and its strings may have been generated specifically
for that event, they are not guaranteed to be valid after the callback has
returned.  If you wish to use them after the callback returns, you need to
make a deep copy.

C documentation : [glfwSetDropCallback]

# params
  - window - The window whose callback to set.
  - callback - The new file drop callback, or `NULL` to remove the
    currently set callback.

# return

The previously set callback, or `NULL` if no callback was set or the
library had not been [initialized].

	void function_name(GLFWwindow* window, int path_count, const char* paths[])

For more information about the callback parameters, see the
[function pointer type].

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.1.

[initialized]: https://www.glfw.org/docs/latest/intro_guide.html#intro_init
[function pointer type]: https://www.glfw.org/docs/latest/group__input.html#gaaba73c3274062c18723b7f05862d94b2
[glfwSetDropCallback]: https://www.glfw.org/docs/latest/group__input.html#gab773f0ee0a07cff77a210cea40bc1f6b
*/
func (window *Window) SetDropCallback(callback DropCallback) DropCallback {
	var result DropCallback
	_, err := ffi.CallFunction(
		cif_glfwSetDropCallback,
		func_glfwSetDropCallback,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&callback),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function returns whether the specified joystick is present.

There is no need to call this function before other functions that accept
a joystick ID, as they all check for presence before performing any other
work.

C documentation : [glfwJoystickPresent]

# params
  - jid - The [joystick] to query.

# return

`GLFW_TRUE` if the joystick is present, or `GLFW_FALSE` otherwise.

# errors

Possible errors include [NOT_INITIALIZED], [INVALID_ENUM] and [PLATFORM_ERROR].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.  Replaces `glfwGetJoystickParam`.

[joystick]: https://www.glfw.org/docs/latest/group__joysticks.html
[glfwJoystickPresent]: https://www.glfw.org/docs/latest/group__input.html#gaed0966cee139d815317f9ffcba64c9f1
*/
func JoystickPresent(jid Int) Int {
	var result Int
	_, err := ffi.CallFunction(
		cif_glfwJoystickPresent,
		func_glfwJoystickPresent,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&jid),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function returns the values of all axes of the specified joystick.
Each element in the array is a value between -1.0 and 1.0.

If the specified joystick is not present this function will return `NULL`
but will not generate an error.  This can be used instead of first calling
[JoystickPresent].

C documentation : [glfwGetJoystickAxes]

# params
  - jid - The [joystick] to query.
  - count (out) - Where to store the number of axis values in the returned
    array.  This is set to zero if the joystick is not present or an error
    occurred.

# return

An array of axis values, or `NULL` if the joystick is not present or
an [error] occurred.

# errors

Possible errors include [NOT_INITIALIZED], [INVALID_ENUM] and [PLATFORM_ERROR].

# pointer lifetime

The returned array is allocated and freed by GLFW.  You
should not free it yourself.  It is valid until the specified joystick is
disconnected or the library is terminated.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.  Replaces `glfwGetJoystickPos`.

[joystick]: https://www.glfw.org/docs/latest/group__joysticks.html
[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwGetJoystickAxes]: https://www.glfw.org/docs/latest/group__input.html#gaeb1c0191d3140a233a682987c61eb408
*/
func GetJoystickAxes(jid Int) []Float {
	var result *Float
	var count_ Int
	count := &count_
	_, err := ffi.CallFunction(
		cif_glfwGetJoystickAxes,
		func_glfwGetJoystickAxes,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&jid),
			unsafe.Pointer(new(&count)),
		},
	)
	if err != nil {
		panic(err)
	}
	slice := make([]Float, count_)
	copy(slice, unsafe.Slice(result, count_))
	return slice
}

/*
This function returns the state of all buttons of the specified joystick.
Each element in the array is either `GLFW_PRESS` or `GLFW_RELEASE`.

For backward compatibility with earlier versions that did not have [GetJoystickHats], the button array also includes all hats, each
represented as four buttons.  The hats are in the same order as returned by
__glfwGetJoystickHats__ and are in the order _up_, _right_, _down_ and
_left_.  To disable these extra buttons, set the [JOYSTICK_HAT_BUTTONS] init hint before initialization.

If the specified joystick is not present this function will return `NULL`
but will not generate an error.  This can be used instead of first calling
[JoystickPresent].

C documentation : [glfwGetJoystickButtons]

# params
  - jid - The [joystick] to query.
  - count (out) - Where to store the number of button states in the returned
    array.  This is set to zero if the joystick is not present or an error
    occurred.

# return

An array of button states, or `NULL` if the joystick is not present
or an [error] occurred.

# errors

Possible errors include [NOT_INITIALIZED], [INVALID_ENUM] and [PLATFORM_ERROR].

# pointer lifetime

The returned array is allocated and freed by GLFW.  You
should not free it yourself.  It is valid until the specified joystick is
disconnected or the library is terminated.

# thread safety

This function must only be called from the main thread.

# since

Added in version 2.2.

[joystick]: https://www.glfw.org/docs/latest/group__joysticks.html
[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwGetJoystickButtons]: https://www.glfw.org/docs/latest/group__input.html#ga5ffe34739d3dc97efe432ed2d81d9938
*/
func GetJoystickButtons(jid Int) []UChar {
	var result *UChar
	var count_ Int
	count := &count_
	_, err := ffi.CallFunction(
		cif_glfwGetJoystickButtons,
		func_glfwGetJoystickButtons,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&jid),
			unsafe.Pointer(new(&count)),
		},
	)
	if err != nil {
		panic(err)
	}
	slice := make([]UChar, count_)
	copy(slice, unsafe.Slice(result, count_))
	return slice
}

/*
This function returns the state of all hats of the specified joystick.
Each element in the array is one of the following values:

Name                  | Value
----                  | -----
`GLFW_HAT_CENTERED`   | 0
`GLFW_HAT_UP`         | 1
`GLFW_HAT_RIGHT`      | 2
`GLFW_HAT_DOWN`       | 4
`GLFW_HAT_LEFT`       | 8
`GLFW_HAT_RIGHT_UP`   | `GLFW_HAT_RIGHT` \| `GLFW_HAT_UP`
`GLFW_HAT_RIGHT_DOWN` | `GLFW_HAT_RIGHT` \| `GLFW_HAT_DOWN`
`GLFW_HAT_LEFT_UP`    | `GLFW_HAT_LEFT` \| `GLFW_HAT_UP`
`GLFW_HAT_LEFT_DOWN`  | `GLFW_HAT_LEFT` \| `GLFW_HAT_DOWN`

The diagonal directions are bitwise combinations of the primary (up, right,
down and left) directions and you can test for these individually by ANDing
it with the corresponding direction.

	if (hats[2] & GLFW_HAT_RIGHT)
	{
	// State of hat 2 could be right-up, right or right-down
	}

If the specified joystick is not present this function will return `NULL`
but will not generate an error.  This can be used instead of first calling
[JoystickPresent].

C documentation : [glfwGetJoystickHats]

# params
  - jid - The [joystick] to query.
  - count (out) - Where to store the number of hat states in the returned
    array.  This is set to zero if the joystick is not present or an error
    occurred.

# return

An array of hat states, or `NULL` if the joystick is not present
or an [error] occurred.

# errors

Possible errors include [NOT_INITIALIZED], [INVALID_ENUM] and [PLATFORM_ERROR].

# pointer lifetime

The returned array is allocated and freed by GLFW.  You
should not free it yourself.  It is valid until the specified joystick is
disconnected, this function is called again for that joystick or the library
is terminated.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.3.

[joystick]: https://www.glfw.org/docs/latest/group__joysticks.html
[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwGetJoystickHats]: https://www.glfw.org/docs/latest/group__input.html#ga06e660841b3e79c54da4f54a932c5a2c
*/
func GetJoystickHats(jid Int) []UChar {
	var result *UChar
	var count_ Int
	count := &count_
	_, err := ffi.CallFunction(
		cif_glfwGetJoystickHats,
		func_glfwGetJoystickHats,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&jid),
			unsafe.Pointer(new(&count)),
		},
	)
	if err != nil {
		panic(err)
	}
	slice := make([]UChar, count_)
	copy(slice, unsafe.Slice(result, count_))
	return slice
}

/*
This function returns the name, encoded as UTF-8, of the specified joystick.
The returned string is allocated and freed by GLFW.  You should not free it
yourself.

If the specified joystick is not present this function will return `NULL`
but will not generate an error.  This can be used instead of first calling
[JoystickPresent].

C documentation : [glfwGetJoystickName]

# params
  - jid - The [joystick] to query.

# return

The UTF-8 encoded name of the joystick, or `NULL` if the joystick
is not present or an [error] occurred.

# errors

Possible errors include [NOT_INITIALIZED], [INVALID_ENUM] and [PLATFORM_ERROR].

# pointer lifetime

The returned string is allocated and freed by GLFW.  You
should not free it yourself.  It is valid until the specified joystick is
disconnected or the library is terminated.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[joystick]: https://www.glfw.org/docs/latest/group__joysticks.html
[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwGetJoystickName]: https://www.glfw.org/docs/latest/group__input.html#gac6a8e769e18e0bcfa9097793fc2c3978
*/
func GetJoystickName(jid Int) string {
	var result *byte
	_, err := ffi.CallFunction(
		cif_glfwGetJoystickName,
		func_glfwGetJoystickName,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&jid),
		},
	)
	if err != nil {
		panic(err)
	}
	return goString(result)
}

/*
This function returns the SDL compatible GUID, as a UTF-8 encoded
hexadecimal string, of the specified joystick.  The returned string is
allocated and freed by GLFW.  You should not free it yourself.

The GUID is what connects a joystick to a gamepad mapping.  A connected
joystick will always have a GUID even if there is no gamepad mapping
assigned to it.

If the specified joystick is not present this function will return `NULL`
but will not generate an error.  This can be used instead of first calling
[JoystickPresent].

The GUID uses the format introduced in SDL 2.0.5.  This GUID tries to
uniquely identify the make and model of a joystick but does not identify
a specific unit, e.g. all wired Xbox 360 controllers will have the same
GUID on that platform.  The GUID for a unit may vary between platforms
depending on what hardware information the platform specific APIs provide.

C documentation : [glfwGetJoystickGUID]

# params
  - jid - The [joystick] to query.

# return

The UTF-8 encoded GUID of the joystick, or `NULL` if the joystick
is not present or an [error] occurred.

# errors

Possible errors include [NOT_INITIALIZED], [INVALID_ENUM] and [PLATFORM_ERROR].

# pointer lifetime

The returned string is allocated and freed by GLFW.  You
should not free it yourself.  It is valid until the specified joystick is
disconnected or the library is terminated.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.3.

[joystick]: https://www.glfw.org/docs/latest/group__joysticks.html
[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwGetJoystickGUID]: https://www.glfw.org/docs/latest/group__input.html#ga6659411aec3c7fcef27780e2cb2d9600
*/
func GetJoystickGUID(jid Int) string {
	var result *byte
	_, err := ffi.CallFunction(
		cif_glfwGetJoystickGUID,
		func_glfwGetJoystickGUID,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&jid),
		},
	)
	if err != nil {
		panic(err)
	}
	return goString(result)
}

/*
This function sets the user-defined pointer of the specified joystick.  The
current value is retained until the joystick is disconnected.  The initial
value is `NULL`.

This function may be called from the joystick callback, even for a joystick
that is being disconnected.

C documentation : [glfwSetJoystickUserPointer]

# params
  - jid - The joystick whose pointer to set.
  - pointer - The new value.

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function may be called from any thread.  Access is not
synchronized.

# since

Added in version 3.3.

[glfwSetJoystickUserPointer]: https://www.glfw.org/docs/latest/group__input.html#ga6b2f72d64d636b48a727b437cbb7489e
*/
func SetJoystickUserPointer(jid Int, pointer unsafe.Pointer) {
	_, err := ffi.CallFunction(
		cif_glfwSetJoystickUserPointer,
		func_glfwSetJoystickUserPointer,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&jid),
			unsafe.Pointer(&pointer),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function returns the current value of the user-defined pointer of the
specified joystick.  The initial value is `NULL`.

This function may be called from the joystick callback, even for a joystick
that is being disconnected.

C documentation : [glfwGetJoystickUserPointer]

# params
  - jid - The joystick whose pointer to return.

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function may be called from any thread.  Access is not
synchronized.

# since

Added in version 3.3.

[glfwGetJoystickUserPointer]: https://www.glfw.org/docs/latest/group__input.html#ga18cefd7265d1fa04f3fd38a6746db5f3
*/
func GetJoystickUserPointer(jid Int) unsafe.Pointer {
	var result unsafe.Pointer
	_, err := ffi.CallFunction(
		cif_glfwGetJoystickUserPointer,
		func_glfwGetJoystickUserPointer,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&jid),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function returns whether the specified joystick is both present and has
a gamepad mapping.

If the specified joystick is present but does not have a gamepad mapping
this function will return `GLFW_FALSE` but will not generate an error.  Call
[JoystickPresent] to check if a joystick is present regardless of
whether it has a mapping.

C documentation : [glfwJoystickIsGamepad]

# params
  - jid - The [joystick] to query.

# return

`GLFW_TRUE` if a joystick is both present and has a gamepad mapping,
or `GLFW_FALSE` otherwise.

# errors

Possible errors include [NOT_INITIALIZED] and [INVALID_ENUM].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.3.

[joystick]: https://www.glfw.org/docs/latest/group__joysticks.html
[glfwJoystickIsGamepad]: https://www.glfw.org/docs/latest/group__input.html#gad0f676860f329d80f7e47e9f06a96f00
*/
func JoystickIsGamepad(jid Int) Int {
	var result Int
	_, err := ffi.CallFunction(
		cif_glfwJoystickIsGamepad,
		func_glfwJoystickIsGamepad,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&jid),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the joystick configuration callback, or removes the
currently set callback.  This is called when a joystick is connected to or
disconnected from the system.

For joystick connection and disconnection events to be delivered on all
platforms, you need to call one of the [event processing]
functions.  Joystick disconnection may also be detected and the callback
called by joystick functions.  The function will then return whatever it
returns if the joystick is not present.

C documentation : [glfwSetJoystickCallback]

# params
  - callback - The new callback, or `NULL` to remove the currently set
    callback.

# return

The previously set callback, or `NULL` if no callback was set or the
library had not been [initialized].

	void function_name(int jid, int event)

For more information about the callback parameters, see the
[function pointer type].

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.2.

[event processing]: https://www.glfw.org/docs/latest/input_guide.html#events
[initialized]: https://www.glfw.org/docs/latest/intro_guide.html#intro_init
[function pointer type]: https://www.glfw.org/docs/latest/group__input.html#gaa21ad5986ae9a26077a40142efb56243
[glfwSetJoystickCallback]: https://www.glfw.org/docs/latest/group__input.html#ga2f60a0e5b7bd8d1b7344dc0a7cb32b4c
*/
func SetJoystickCallback(callback JoystickCallback) JoystickCallback {
	var result JoystickCallback
	_, err := ffi.CallFunction(
		cif_glfwSetJoystickCallback,
		func_glfwSetJoystickCallback,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&callback),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function parses the specified ASCII encoded string and updates the
internal list with any gamepad mappings it finds.  This string may
contain either a single gamepad mapping or many mappings separated by
newlines.  The parser supports the full format of the `gamecontrollerdb.txt`
source file including empty lines and comments.

See [Gamepad mappings] for a description of the format.

If there is already a gamepad mapping for a given GUID in the internal list,
it will be replaced by the one passed to this function.  If the library is
terminated and re-initialized the internal list will revert to the built-in
default.

C documentation : [glfwUpdateGamepadMappings]

# params
  - string - The string containing the gamepad mappings.

# return

`GLFW_TRUE` if successful, or `GLFW_FALSE` if an
[error] occurred.

# errors

Possible errors include [NOT_INITIALIZED] and [INVALID_VALUE].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.3.

[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[Gamepad mappings]: https://www.glfw.org/docs/latest/input_guide.html#gamepad_mapping
[glfwUpdateGamepadMappings]: https://www.glfw.org/docs/latest/group__input.html#gaed5104612f2fa8e66aa6e846652ad00f
*/
func UpdateGamepadMappings(string string) Int {
	var result Int
	c_string := cString(string)
	_, err := ffi.CallFunction(
		cif_glfwUpdateGamepadMappings,
		func_glfwUpdateGamepadMappings,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&c_string),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function returns the human-readable name of the gamepad from the
gamepad mapping assigned to the specified joystick.

If the specified joystick is not present or does not have a gamepad mapping
this function will return `NULL` but will not generate an error.  Call
[JoystickPresent] to check whether it is present regardless of
whether it has a mapping.

C documentation : [glfwGetGamepadName]

# params
  - jid - The [joystick] to query.

# return

The UTF-8 encoded name of the gamepad, or `NULL` if the
joystick is not present, does not have a mapping or an
[error] occurred.

# errors

Possible errors include [NOT_INITIALIZED] and [INVALID_ENUM].

# pointer lifetime

The returned string is allocated and freed by GLFW.  You
should not free it yourself.  It is valid until the specified joystick is
disconnected, the gamepad mappings are updated or the library is terminated.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.3.

[joystick]: https://www.glfw.org/docs/latest/group__joysticks.html
[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwGetGamepadName]: https://www.glfw.org/docs/latest/group__input.html#ga8aea73a1a25cc6c0486a617019f56728
*/
func GetGamepadName(jid Int) string {
	var result *byte
	_, err := ffi.CallFunction(
		cif_glfwGetGamepadName,
		func_glfwGetGamepadName,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&jid),
		},
	)
	if err != nil {
		panic(err)
	}
	return goString(result)
}

/*
This function retrieves the state of the specified joystick remapped to
an Xbox-like gamepad.

If the specified joystick is not present or does not have a gamepad mapping
this function will return `GLFW_FALSE` but will not generate an error.  Call
[JoystickPresent] to check whether it is present regardless of
whether it has a mapping.

The Guide button may not be available for input as it is often hooked by the
system or the Steam client.

Not all devices have all the buttons or axes provided by [Gamepadstate].  Unavailable buttons and axes will always report
`GLFW_RELEASE` and 0.0 respectively.

C documentation : [glfwGetGamepadState]

# params
  - jid - The [joystick] to query.
  - state (out) - The gamepad input state of the joystick.

# return

`GLFW_TRUE` if successful, or `GLFW_FALSE` if no joystick is
connected, it has no gamepad mapping or an [error]
occurred.

# errors

Possible errors include [NOT_INITIALIZED] and [INVALID_ENUM].

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.3.

[joystick]: https://www.glfw.org/docs/latest/group__joysticks.html
[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwGetGamepadState]: https://www.glfw.org/docs/latest/group__input.html#gadccddea8bce6113fa459de379ddaf051
*/
func GetGamepadState(jid Int) (Gamepadstate, Int) {
	var result Int
	var state Gamepadstate
	_, err := ffi.CallFunction(
		cif_glfwGetGamepadState,
		func_glfwGetGamepadState,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&jid),
			unsafe.Pointer(new(&state)),
		},
	)
	if err != nil {
		panic(err)
	}
	return state, result
}

/*
This function sets the system clipboard to the specified, UTF-8 encoded
string.

C documentation : [glfwSetClipboardString]

# params
  - window - Deprecated.  Any valid window or `NULL`.
  - string - A UTF-8 encoded string.

# errors

Possible errors include [NOT_INITIALIZED] and [PLATFORM_ERROR].

# win32

The clipboard on Windows has a single global lock for reading and
writing.  GLFW tries to acquire it a few times, which is almost always enough.  If it
cannot acquire the lock then this function emits [PLATFORM_ERROR] and returns.
It is safe to try this multiple times.

# pointer lifetime

The specified string is copied before this function
returns.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[glfwSetClipboardString]: https://www.glfw.org/docs/latest/group__input.html#gaba1f022c5eb07dfac421df34cdcd31dd
*/
func (window *Window) SetClipboardString(string string) {
	c_string := cString(string)
	_, err := ffi.CallFunction(
		cif_glfwSetClipboardString,
		func_glfwSetClipboardString,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
			unsafe.Pointer(&c_string),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function returns the contents of the system clipboard, if it contains
or is convertible to a UTF-8 encoded string.  If the clipboard is empty or
if its contents cannot be converted, `NULL` is returned and a [FORMAT_UNAVAILABLE] error is generated.

C documentation : [glfwGetClipboardString]

# params
  - window - Deprecated.  Any valid window or `NULL`.

# return

The contents of the clipboard as a UTF-8 encoded string, or `NULL`
if an [error] occurred.

# errors

Possible errors include [NOT_INITIALIZED], [FORMAT_UNAVAILABLE] and [PLATFORM_ERROR].

# win32

The clipboard on Windows has a single global lock for reading and
writing.  GLFW tries to acquire it a few times, which is almost always enough.  If it
cannot acquire the lock then this function emits [PLATFORM_ERROR] and returns.
It is safe to try this multiple times.

# pointer lifetime

The returned string is allocated and freed by GLFW.  You
should not free it yourself.  It is valid until the next call to [GetClipboardString] or [SetClipboardString], or until the library
is terminated.

# thread safety

This function must only be called from the main thread.

# since

Added in version 3.0.

[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwGetClipboardString]: https://www.glfw.org/docs/latest/group__input.html#ga71a5b20808ea92193d65c21b82580355
*/
func (window *Window) GetClipboardString() string {
	var result *byte
	_, err := ffi.CallFunction(
		cif_glfwGetClipboardString,
		func_glfwGetClipboardString,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
		},
	)
	if err != nil {
		panic(err)
	}
	return goString(result)
}

/*
This function returns the current GLFW time, in seconds.  Unless the time
has been set using [SetTime] it measures time elapsed since GLFW was
initialized.

This function and [SetTime] are helper functions on top of [GetTimerFrequency] and [GetTimerValue].

The resolution of the timer is system dependent, but is usually on the order
of a few micro- or nanoseconds.  It uses the highest-resolution monotonic
time source on each operating system.

C documentation : [glfwGetTime]

# return

The current time, in seconds, or zero if an
[error] occurred.

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function may be called from any thread.  Reading and
writing of the internal base time is not atomic, so it needs to be
externally synchronized with calls to [SetTime].

# since

Added in version 1.0.

[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwGetTime]: https://www.glfw.org/docs/latest/group__input.html#gaa6cf4e7a77158a3b8fd00328b1720a4a
*/
func GetTime() Double {
	var result Double
	_, err := ffi.CallFunction(
		cif_glfwGetTime,
		func_glfwGetTime,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function sets the current GLFW time, in seconds.  The value must be
a positive finite number less than or equal to 18446744073.0, which is
approximately 584.5 years.

This function and [GetTime] are helper functions on top of [GetTimerFrequency] and [GetTimerValue].

C documentation : [glfwSetTime]

# params
  - time - The new value, in seconds.

# errors

Possible errors include [NOT_INITIALIZED] and [INVALID_VALUE].

The upper limit of GLFW time is calculated as
floor((2<sup>64</sup> - 1) / 10<sup>9</sup>) and is due to implementations
storing nanoseconds in 64 bits.  The limit may be increased in the future.

# thread safety

This function may be called from any thread.  Reading and
writing of the internal base time is not atomic, so it needs to be
externally synchronized with calls to [GetTime].

# since

Added in version 2.2.

[glfwSetTime]: https://www.glfw.org/docs/latest/group__input.html#gaf59589ef6e8b8c8b5ad184b25afd4dc0
*/
func SetTime(time Double) {
	_, err := ffi.CallFunction(
		cif_glfwSetTime,
		func_glfwSetTime,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&time),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function returns the current value of the raw timer, measured in
1&nbsp;/&nbsp;frequency seconds.  To get the frequency, call [GetTimerFrequency].

C documentation : [glfwGetTimerValue]

# return

The value of the timer, or zero if an
[error] occurred.

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function may be called from any thread.

# since

Added in version 3.2.

[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwGetTimerValue]: https://www.glfw.org/docs/latest/group__input.html#ga09b2bd37d328e0b9456c7ec575cc26aa
*/
func GetTimerValue() ULong {
	var result ULong
	_, err := ffi.CallFunction(
		cif_glfwGetTimerValue,
		func_glfwGetTimerValue,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function returns the frequency, in Hz, of the raw timer.

C documentation : [glfwGetTimerFrequency]

# return

The frequency of the timer, in Hz, or zero if an
[error] occurred.

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function may be called from any thread.

# since

Added in version 3.2.

[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwGetTimerFrequency]: https://www.glfw.org/docs/latest/group__input.html#ga3289ee876572f6e91f06df3a24824443
*/
func GetTimerFrequency() ULong {
	var result ULong
	_, err := ffi.CallFunction(
		cif_glfwGetTimerFrequency,
		func_glfwGetTimerFrequency,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
thread.

This function makes the OpenGL or OpenGL ES context of the specified window
current on the calling thread.  It can also detach the current context from
the calling thread without making a new one current by passing in `NULL`.

A context must only be made current on a single thread at a time and each
thread can have only a single current context at a time.  Making a context
current detaches any previously current context on the calling thread.

When moving a context between threads, you must detach it (make it
non-current) on the old thread before making it current on the new one.

By default, making a context non-current implicitly forces a pipeline flush.
On machines that support `GL_KHR_context_flush_control`, you can control
whether a context performs this flush by setting the
[GLFW_CONTEXT_RELEASE_BEHAVIOR]
hint.

The specified window must have an OpenGL or OpenGL ES context.  Specifying
a window without a context will generate a [NO_WINDOW_CONTEXT]
error.

C documentation : [glfwMakeContextCurrent]

# params
  - window - The window whose context to make current, or `NULL` to
    detach the current context.

If the previously current context was created via a different
context creation API than the one passed to this function, GLFW will still
detach the previous one from its API before making the new one current.

# errors

Possible errors include [NOT_INITIALIZED], [NO_WINDOW_CONTEXT] and [PLATFORM_ERROR].

# thread safety

This function may be called from any thread.

# since

Added in version 3.0.

[GLFW_CONTEXT_RELEASE_BEHAVIOR]: https://www.glfw.org/docs/latest/window_guide.html#GLFW_CONTEXT_RELEASE_BEHAVIOR_hint
[glfwMakeContextCurrent]: https://www.glfw.org/docs/latest/group__context.html#ga1c04dc242268f827290fe40aa1c91157
*/
func (window *Window) MakeContextCurrent() {
	_, err := ffi.CallFunction(
		cif_glfwMakeContextCurrent,
		func_glfwMakeContextCurrent,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function returns the window whose OpenGL or OpenGL ES context is
current on the calling thread.

C documentation : [glfwGetCurrentContext]

# return

The window whose context is current, or `NULL` if no window's
context is current.

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function may be called from any thread.

# since

Added in version 3.0.

[glfwGetCurrentContext]: https://www.glfw.org/docs/latest/group__context.html#gad94e80185397a6cf5fe2ab30567af71c
*/
func GetCurrentContext() *Window {
	var result *Window
	_, err := ffi.CallFunction(
		cif_glfwGetCurrentContext,
		func_glfwGetCurrentContext,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function swaps the front and back buffers of the specified window when
rendering with OpenGL or OpenGL ES.  If the swap interval is greater than
zero, the GPU driver waits the specified number of screen updates before
swapping the buffers.

The specified window must have an OpenGL or OpenGL ES context.  Specifying
a window without a context will generate a [NO_WINDOW_CONTEXT]
error.

This function does not apply to Vulkan.  If you are rendering with Vulkan,
see `vkQueuePresentKHR` instead.

C documentation : [glfwSwapBuffers]

# params
  - window - The window whose buffers to swap.

# errors

Possible errors include [NOT_INITIALIZED], [NO_WINDOW_CONTEXT] and [PLATFORM_ERROR].

__EGL:__ The context of the specified window must be current on the
calling thread.

# thread safety

This function may be called from any thread.

# since

Added in version 1.0.

[glfwSwapBuffers]: https://www.glfw.org/docs/latest/group__window.html#ga15a5a1ee5b3c2ca6b15ca209a12efd14
*/
func (window *Window) SwapBuffers() {
	_, err := ffi.CallFunction(
		cif_glfwSwapBuffers,
		func_glfwSwapBuffers,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&window),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function sets the swap interval for the current OpenGL or OpenGL ES
context, i.e. the number of screen updates to wait from the time [SwapBuffers] was called before swapping the buffers and returning.  This
is sometimes called _vertical synchronization_, _vertical retrace
synchronization_ or just _vsync_.

A context that supports either of the `WGL_EXT_swap_control_tear` and
`GLX_EXT_swap_control_tear` extensions also accepts _negative_ swap
intervals, which allows the driver to swap immediately even if a frame
arrives a little bit late.  You can check for these extensions with [ExtensionSupported].

A context must be current on the calling thread.  Calling this function
without a current context will cause a [NO_CURRENT_CONTEXT] error.

This function does not apply to Vulkan.  If you are rendering with Vulkan,
see the present mode of your swapchain instead.

C documentation : [glfwSwapInterval]

# params
  - interval - The minimum number of screen updates to wait for
    until the buffers are swapped by [SwapBuffers].

# errors

Possible errors include [NOT_INITIALIZED], [NO_CURRENT_CONTEXT] and [PLATFORM_ERROR].

This function is not called during context creation, leaving the
swap interval set to whatever is the default for that API.  This is done
because some swap interval extensions used by GLFW do not allow the swap
interval to be reset to zero once it has been set to a non-zero value.

Some GPU drivers do not honor the requested swap interval, either
because of a user setting that overrides the application's request or due to
bugs in the driver.

# thread safety

This function may be called from any thread.

# since

Added in version 1.0.

[glfwSwapInterval]: https://www.glfw.org/docs/latest/group__context.html#ga6d4e0cdf151b5e579bd67f13202994ed
*/
func SwapInterval(interval Int) {
	_, err := ffi.CallFunction(
		cif_glfwSwapInterval,
		func_glfwSwapInterval,
		unsafe.Pointer(nil),
		[]unsafe.Pointer{
			unsafe.Pointer(&interval),
		},
	)
	if err != nil {
		panic(err)
	}
}

/*
This function returns whether the specified
[API extension] is supported by the current OpenGL or
OpenGL ES context.  It searches both for client API extension and context
creation API extensions.

A context must be current on the calling thread.  Calling this function
without a current context will cause a [NO_CURRENT_CONTEXT] error.

As this functions retrieves and searches one or more extension strings each
call, it is recommended that you cache its results if it is going to be used
frequently.  The extension strings will not change during the lifetime of
a context, so there is no danger in doing this.

This function does not apply to Vulkan.  If you are using Vulkan, see [GetRequiredInstanceExtensions], `vkEnumerateInstanceExtensionProperties`
and `vkEnumerateDeviceExtensionProperties` instead.

C documentation : [glfwExtensionSupported]

# params
  - extension - The ASCII encoded name of the extension.

# return

`GLFW_TRUE` if the extension is available, or `GLFW_FALSE`
otherwise.

# errors

Possible errors include [NOT_INITIALIZED], [NO_CURRENT_CONTEXT], [INVALID_VALUE] and [PLATFORM_ERROR].

# thread safety

This function may be called from any thread.

# since

Added in version 1.0.

[API extension]: https://www.glfw.org/docs/latest/context_guide.html#context_glext
[glfwExtensionSupported]: https://www.glfw.org/docs/latest/group__context.html#ga87425065c011cef1ebd6aac75e059dfa
*/
func ExtensionSupported(extension string) Int {
	var result Int
	c_extension := cString(extension)
	_, err := ffi.CallFunction(
		cif_glfwExtensionSupported,
		func_glfwExtensionSupported,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&c_extension),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
context.

This function returns the address of the specified OpenGL or OpenGL ES
[core or extension function], if it is supported
by the current context.

A context must be current on the calling thread.  Calling this function
without a current context will cause a [NO_CURRENT_CONTEXT] error.

This function does not apply to Vulkan.  If you are rendering with Vulkan,
see [glfwGetInstanceProcAddress], `vkGetInstanceProcAddr` and
`vkGetDeviceProcAddr` instead.

C documentation : [glfwGetProcAddress]

# params
  - procname - The ASCII encoded name of the function.

# return

The address of the function, or `NULL` if an
[error] occurred.

# errors

Possible errors include [NOT_INITIALIZED], [NO_CURRENT_CONTEXT] and [PLATFORM_ERROR].

The address of a given function is not guaranteed to be the same
between contexts.

This function may return a non-`NULL` address despite the
associated version or extension not being available.  Always check the
context version or extension string first.

# pointer lifetime

The returned function pointer is valid until the context
is destroyed or the library is terminated.

# thread safety

This function may be called from any thread.

# since

Added in version 1.0.

[core or extension function]: https://www.glfw.org/docs/latest/context_guide.html#context_glext
[error]: https://www.glfw.org/docs/latest/intro_guide.html#error_handling
[glfwGetInstanceProcAddress]: https://www.glfw.org/docs/latest/group__vulkan.html#gadf228fac94c5fd8f12423ec9af9ff1e9
[glfwGetProcAddress]: https://www.glfw.org/docs/latest/group__context.html#ga35f1837e6f666781842483937612f163
*/
func GetProcAddress(procname string) GlProc {
	var result GlProc
	c_procname := cString(procname)
	_, err := ffi.CallFunction(
		cif_glfwGetProcAddress,
		func_glfwGetProcAddress,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{
			unsafe.Pointer(&c_procname),
		},
	)
	if err != nil {
		panic(err)
	}
	return result
}

/*
This function returns whether the Vulkan loader and any minimally functional
ICD have been found.

The availability of a Vulkan loader and even an ICD does not by itself guarantee that
surface creation or even instance creation is possible.  Call [GetRequiredInstanceExtensions] to check whether the extensions necessary for Vulkan
surface creation are available and [glfwGetPhysicalDevicePresentationSupport] to
check whether a queue family of a physical device supports image presentation.

C documentation : [glfwVulkanSupported]

# return

`GLFW_TRUE` if Vulkan is minimally available, or `GLFW_FALSE`
otherwise.

# errors

Possible errors include [NOT_INITIALIZED].

# thread safety

This function may be called from any thread.

# since

Added in version 3.2.

[glfwGetPhysicalDevicePresentationSupport]: https://www.glfw.org/docs/latest/group__vulkan.html#gaff3823355cdd7e2f3f9f4d9ea9518d92
[glfwVulkanSupported]: https://www.glfw.org/docs/latest/group__vulkan.html#ga2e7f30931e02464b5bc8d0d4b6f9fe2b
*/
func VulkanSupported() Int {
	var result Int
	_, err := ffi.CallFunction(
		cif_glfwVulkanSupported,
		func_glfwVulkanSupported,
		unsafe.Pointer(&result),
		[]unsafe.Pointer{},
	)
	if err != nil {
		panic(err)
	}
	return result
}

// UNSUPPORTED glfwGetRequiredInstanceExtensions : result type is "const char **"
