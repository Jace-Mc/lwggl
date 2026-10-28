// Welcome to lwggl.
//
// This is the glfw/glfw.go file, this contains all the GLFW bindings.
// No librarys are external that includes:
// GLAD and GLFW, and any other external librarys.
//
// Link to the GLFW website:
// https://www.glfw.org
//
// DEPENDENCIES: (all are included)
//
// [GLFW] ~ For OpenGL Window, context, and input.
//
// LICENSE: zlib/libpng
// The license is shown in `LICENSE`
package glfw

/*
#cgo CFLAGS: -I ../external/glfw/include
#include "GLFW/glfw3.h"
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// Typedef for int (GLFWbool)
type GLFWbool int

// Typedef for float64(double) (GLFWtime)
type GLFWtime float64

// enum of boolean types
const (
	TRUE  GLFWbool = C.GLFW_TRUE
	FALSE GLFWbool = C.GLFW_FALSE
)

// Typedef for int (GLFWkey)
type GLFWkey int

// enum of glfw keys.
const (
	KEY_A      GLFWkey = C.GLFW_KEY_A
	KEY_B      GLFWkey = C.GLFW_KEY_B
	KEY_C      GLFWkey = C.GLFW_KEY_C
	KEY_D      GLFWkey = C.GLFW_KEY_D
	KEY_E      GLFWkey = C.GLFW_KEY_E
	KEY_F      GLFWkey = C.GLFW_KEY_F
	KEY_G      GLFWkey = C.GLFW_KEY_G
	KEY_H      GLFWkey = C.GLFW_KEY_H
	KEY_I      GLFWkey = C.GLFW_KEY_I
	KEY_J      GLFWkey = C.GLFW_KEY_J
	KEY_K      GLFWkey = C.GLFW_KEY_K
	KEY_L      GLFWkey = C.GLFW_KEY_L
	KEY_M      GLFWkey = C.GLFW_KEY_M
	KEY_N      GLFWkey = C.GLFW_KEY_N
	KEY_O      GLFWkey = C.GLFW_KEY_O
	KEY_P      GLFWkey = C.GLFW_KEY_P
	KEY_Q      GLFWkey = C.GLFW_KEY_Q
	KEY_R      GLFWkey = C.GLFW_KEY_R
	KEY_S      GLFWkey = C.GLFW_KEY_S
	KEY_T      GLFWkey = C.GLFW_KEY_T
	KEY_U      GLFWkey = C.GLFW_KEY_U
	KEY_V      GLFWkey = C.GLFW_KEY_V
	KEY_W      GLFWkey = C.GLFW_KEY_W
	KEY_X      GLFWkey = C.GLFW_KEY_X
	KEY_Y      GLFWkey = C.GLFW_KEY_Y
	KEY_Z      GLFWkey = C.GLFW_KEY_Z
	KEY_ESCAPE GLFWkey = C.GLFW_KEY_ESCAPE
)

// enum of glfw key events.
const (
	KEY_PRESS   GLFWbool = C.GLFW_PRESS
	KEY_RELEASE GLFWbool = C.GLFW_RELEASE
)

// The struct GLFWwindow.
// This is a handle struct so no attributes.
type GLFWwindow struct{}

// The struct _GLFWvideoMode.
// This is a core struct so there are (5) attributes.
type _GLFWvideoMode struct {
	width  int
	height int
	x      int
	y      int
}

type _GLFWhints struct {
	resizable   GLFWbool
	minimizable GLFWbool
}

// The struct _GLFWplatform.
// This is a reciever struct so it has no attributes.
// This struct needs to be here for _glfw.platform.<functions || variables>
type _GLFWplatform struct{}

// This function converts C.integers to booleans, C.int function && variables can be converted.
func (p _GLFWplatform) boolifyfunction(function C.int) bool {
	return function == C.int(TRUE)
}

// This function destroys the GLFWwindow via an abstract function.
func (p _GLFWplatform) destroyWindow(window _GLFWwindow) {
	C.glfwDestroyWindow(window.handle)
}

// Terminates GLFW itself, by directly calling the function C.glfwTerminate()
func (p _GLFWplatform) terminate() {
	C.glfwTerminate()
}

// Creates the GLFW window object, but does not initialize GLFW.
// You must call `glfw.Init()`
func (p _GLFWplatform) create(window _GLFWwindow) bool {
	// Converts the Go string to a char* C string,
	// after the titleconv is used it frees the string.
	// defer is called on `C.free` for this exact same reason.
	titleconv := C.CString(window.title)
	defer C.free(unsafe.Pointer(titleconv))

	// Creates the C `GLFWwindow*` object,
	// Since this is a binding, and not a full on GLFW implmentation.
	// As you notice parameters: `monitor` and `share` are nil,
	// because no functions for these parameters have been implemented in this binding yet.
	// Also `window.videoMode.width` and `window.videoMode.height` have also been converted to C.int.
	window.handle = C.glfwCreateWindow(C.int(window.videoMode.width),
		C.int(window.videoMode.height),
		titleconv,
		nil,
		nil)

	// Automatic Error checking logic for required for `glfw.CreateWindow`,
	// If the C.glfwCreateWindow returns nil or `NULL` (because of errors) this function will return false.
	// Else the `return true` logic will take priority.
	if window.handle == nil {
		return false
	}

	return true
}

// Swaps the OpenGL Buffers.
func (p _GLFWplatform) swapBuffers(window _GLFWwindow) {
	// Does what is says,
	// It swaps the OpenGL buffers after the OpenGL context provided the window object has been made current,
	// if not this function is literally useless.
	// But it is completely up to the user to call the needed functions.
	C.glfwSwapBuffers(window.handle)
}

// Polls the input events by directly calling `C.glfwPollEvents`
func (p _GLFWplatform) pollEvents() {
	// Directly poll GLFW input events.
	// This function is literally no simpler than that all it is,
	// It's not a complex explanation at all.
	C.glfwPollEvents()
}

// Makes the window context current.
func (p _GLFWplatform) makeContextCurrent(window _GLFWwindow) {
	// Makes the OpenGL context current on the `GLFWwindow*` object provided,
	// when this function is not called the user is not able to swap the opengl buffers,
	// since the context gives way for all the other configurations.
	C.glfwMakeContextCurrent(window.handle)
}

// Gets the current key found that was input but needs `glfwPollEvents` to work.
func (p _GLFWplatform) GetKey(window _GLFWwindow, key GLFWkey) GLFWbool {
	_keyconv := GLFWbool(C.glfwGetKey(window.handle, C.int(key)))
	return _keyconv
}

// Gets the current time since the window started drawing.
func (p _GLFWplatform) GetTime() GLFWtime {
	_timeconv := GLFWtime(C.glfwGetTime())
	return _timeconv
}

// This function initializes GLFW so that most functions, from C can be used here.
func (p _GLFWplatform) initialize() bool {
	// Converts C's glfwInit (int) to a Go usable boolean.
	_initconv := _glfw.platform.boolifyfunction(C.glfwInit())
	return _initconv
}

// This function checks if the GLFW window should close or not.
func (p _GLFWplatform) windowShouldClose(window _GLFWwindow) bool {
	// Converts C's glfwWindowShould (int) to a Go usable boolean.
	_shouldcloseconv := _glfw.platform.boolifyfunction(C.glfwWindowShouldClose(window.handle))
	return _shouldcloseconv
}

func (p _GLFWplatform) setWindowShouldClose(window _GLFWwindow, value int) {
	C.glfwSetWindowShouldClose(window.handle, C.int(value))
}

func (l _GLFWlibrary) invalidError(format string, va ...any) {
	fmt.Printf(format, va...)
}

type _GLFWwindow struct {
	videoMode _GLFWvideoMode
	title     string
	handle    *C.GLFWwindow
}

type _GLFWlibrary struct {
	platform _GLFWplatform
	hints    _GLFWhints
}

var _glfw _GLFWlibrary = _GLFWlibrary{}
var _window _GLFWwindow = _GLFWwindow{}

// @fn glfw.Init
//
// Initializes GLFW.
// This function initializes GLFW so that all function in the library are usable,
// If any errors are detected in the library,
// this function will return `glfw.FALSE` or `0` or anything that results to `0`,
// if not this function will return `glfw.TRUE` or `1` or anything that results to `1`.
//
// < No parameters >
func Init() bool {
	// So if `_glfw.platform.initialize()` returns true || false (normally because of errors during initialization or no errors),
	// it does one of these:
	// If true then the library keeps on running.
	// If false then that error checking logic is completely up for the user to decide with the logic provided.
	return _glfw.platform.initialize()
}

// @fn glfw.CreateWindow
//
// This Creates a GLFWwindow,
// This function unlike the C version does not have monitor, share arguments.
// Why? I don't want to add them yet.
//
// Initializes a GLFWwindow object, that during initialization creates the window object for your current platform.
// Whilst the internal functions of this module does error checking for this function.
// Normally if any errors are detected with the logic provided,
// it is up for you to decide how to configure the error checking logic.
//
// @params:
//   - width: the width of the window being created.
//   - height: the height of the window being created.
//   - title: the title of the window being created.
func CreateWindow(width int, height int, title string) *GLFWwindow {
	handle := new(GLFWwindow)

	_window.videoMode.width = width
	_window.videoMode.height = height
	_window.title = title

	// So if the _glfw.platform.create function returns false (normally because of errors), it returns nil.
	// If true then it just calls the function and returns the `return handle` logic after it.
	// It is completely up to the user to do error checking with the logic provided.
	if !_glfw.platform.create(_window) {
		_glfw.invalidError("Could not create GLFWwindow\n")
		return nil
	}

	return handle
}

// @fn glfw.SwapBuffers
//
// This Swaps Buffers.
// Swaps the OpenGL buffers of the window after the window is created,
// And has the GLX, CGL, or WGL Context has been made current.
// Once OpenGL buffers have been swapped the User may be able to use OpenGL functions for graphics programming on the window.
//
// @params:
//   - window: the window object that will have it's buffers swapped.
func SwapBuffers(window *GLFWwindow) {
	// Calls the `_glfw.platform.swapBuffers` function,
	// which then calls C.glfwSwapBuffers.
	_glfw.platform.swapBuffers(_window)
}

// @fn glfw.PollEvents
//
// This Polls Events.
// Polls Input, Window events that is when the events arise this function Allows these events to do their job.
// This is also used to display the window, `SwapBuffers` only Swaps the OpenGL,
// When Using vulkan you would only use this function.
//
// < No parameters >
func PollEvents() {
	// Polls any events found while using glfw,
	// these could range from input to window events.
	// It is like always said completely up to the user to call this function.
	_glfw.platform.pollEvents()
}

// @fn glfw.GetKey
//
// Gets the Current Key.
// This function gets the current key specified events,
// When the Key specified is Pressed (clicked) It returns `glfw.PRESS` or `1` or anything that result to `1`,
// When the Key specified is Released (declicked) It returns `glfw.RELEASE` or `0` or anything that results to `0`.
//
// @params:
//   - window: the window object that will use the Key.
//   - key: the key specified to have it's events found.
func GetKey(window *GLFWwindow, key GLFWkey) GLFWbool {
	// Checks to see if the key provided provides a specific event wanted by the user,
	// If per say:
	// I was to call `glfw.GetKey(glfw.KEY_A)`.
	// This will either return `glfw.KEY_PRESS` or `glfw.KEY_RELEASE`,
	// if the key is released it will return `0` (glfw.KEY_RELEASE),
	// else the key is pressed it will return `1` (glfw.KEY_PRESS).
	return _glfw.platform.GetKey(_window, key)
}

// @fn glfw.GetTime
//
// Gets the Current Time.
// This function gets the current time elapsed since the program was first ran,
// This will always reset everytime the program is exited and ran again.
// Used for frames per second, dt `deltaTime` calculations, etc.
//
// < No parameters >
func GetTime() GLFWtime {
	return _glfw.platform.GetTime()
}

// @fn glfw.MakeContextCurrent
//
// Makes the context current.
// This function makes the OpenGL context current,
// For window object created.
// This allows the user to swap the buffers and go on with their graphics programming journey with simplicity.
//
// @params:
//   - window: the window object that will have it's context made current.
func MakeContextCurrent(window *GLFWwindow) {
	// Makes the OpenGL context current,
	// So that the user may be able to swap buffers, use openGL function, etc.
	// If this function is not called most graphics functions are just paperweights.
	_glfw.platform.makeContextCurrent(_window)
}

// @fn glfw.SetWindowPos
//
// Sets the window position.
// This function sets the position of the window,
// Uses the coordinates `x` and `y`.
//
// @params:
//   - window: the window object that needs it's position set.
//   - x: the x position of the window object created.
//   - y: the y position of the window object created.
func SetWindowPos(window *GLFWwindow, x int, y int) {
	_window.videoMode.x = x
	_window.videoMode.y = y
}

// @fn glfw.WindowShouldClose
//
// Checks if the window should close or not.
// This is the window's mainLoop all drawing code must be placed here,
// When the input events are polled if the window sends a window closed or open message,
// The window will do it's corrosponding action, or if you have your own logic in this loop it will do so.
// All Drawing code must be placed inside this loop.
// If the window has closed, this function will return `glfw.TRUE` or `1` or anything that results to `1`,
// If the window is still open, this function will return `glfw.FALSE` or `0` or anything that results to `0`.
//
// @params:
//   - window: the window object created that should close.
func WindowShouldClose(window *GLFWwindow) bool {
	// Checks if the window should close.
	// It's no more confusing than that.
	return _glfw.platform.windowShouldClose(_window)
}

// @fn glfw.SetWindowShouldClose
//
// Sets if the window should close or not.
// Unlike WindowShouldClose, it sets if the window should close,
// it does not check if the window should close.
// Normally used to exit the windowShouldClose loop.
// `value` can be set to `glfw.TRUE` or `glfw.FALSE`.
//
// @params:
//   - window: the window object that will be set to close or not.
//   - value: the state of the window object that will be set to close or not.
func SetWindowShouldClose(window *GLFWwindow, value GLFWbool) {
	// Sets if the window should close.
	// No more confusing.
	_glfw.platform.setWindowShouldClose(_window, int(value))
}

// @fn glfw.DestroyWindow
//
// Destroys the window.
// This function destroys the window object that means that all allocated will be deallocated,
// even though the `glfw.Terminate` free everything, this function destroys specific stuff allocated.
// This function unlike the `glfw.Terminate` function should be called after not before the mainLoop.
//
// @params:
//   - window: the window object created that will be destroyed.
func DestroyWindow(window *GLFWwindow) {
	// Destroys the GLFW window.
	_glfw.platform.destroyWindow(_window)
}

// @fn glfw.Terminate
//
// Terminates the GLFWlibrary.
// This function terminates the entire library without this function nothing will work,
// this function also deallocates everything the GLFW library uses internally.
// This however does not destroy or deallocate objects or variables in your program that deallocation is entirely up to you.
//
// < No parameters >
func Terminate() {
	// Terminates GLFW.
	_glfw.platform.terminate()
}

//==================================================
// lwggl creator marks:
// END OF FILE!
//
// ** NOTHING SHALL BE WRITTEN BEYOND THIS MARKER **
//==================================================
