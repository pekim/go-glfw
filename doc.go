/*
Package glfw implements Go bindings for glfw, without use of cgo.

# Initialisation

[Initialise] must be called before calling any other function.
If an api function is called before [Initialise] has been called,
a panic will result.

# glfw version

Only functions supported by the available version of glfw can be called.
Calling a function that is only supported in a later version will result in a panic.

The [GetVersion] function can be used to find the available library version.
*/
package glfw
