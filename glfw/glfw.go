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

GLFWwindow* window;

void InitGLFW(void)
{	
	if (!glfwInit()) exit(1);
}

void CreateWindow(int width, int height, const char* title)
{
	window = glfwCreateWindow(width, height, title, NULL, NULL);
}

int WindowShouldClose(void)
{
	return glfwWindowShouldClose(window);
}

void SetWindowPosition(int x, int y)
{
	glfwSetWindowPos(window, x, y);
}

void MakeContextCurrent(void)
{
	glfwMakeContextCurrent(window);
}

void SwapBuffers(void)
{
	glfwSwapBuffers(window);
}

void SetWindowShouldClose(int value)
{
	glfwSetWindowShouldClose(window, value);
}

int GetGLFWKey(int key)
{
	return glfwGetKey(window, key);
}

void DestroyGLFWWindow(void)
{
	glfwDestroyWindow(window);
}
*/
import "C"
import "unsafe"

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
	KeyA      GLFWkey = C.GLFW_KEY_A
	KeyB      GLFWkey = C.GLFW_KEY_B
	KeyC      GLFWkey = C.GLFW_KEY_C
	KeyD      GLFWkey = C.GLFW_KEY_D
	KeyE      GLFWkey = C.GLFW_KEY_E
	KeyF      GLFWkey = C.GLFW_KEY_F
	KeyG      GLFWkey = C.GLFW_KEY_G
	KeyH      GLFWkey = C.GLFW_KEY_H
	KeyI      GLFWkey = C.GLFW_KEY_I
	KeyJ      GLFWkey = C.GLFW_KEY_J
	KeyK      GLFWkey = C.GLFW_KEY_K
	KeyL      GLFWkey = C.GLFW_KEY_L
	KeyM      GLFWkey = C.GLFW_KEY_M
	KeyN      GLFWkey = C.GLFW_KEY_N
	KeyO      GLFWkey = C.GLFW_KEY_O
	KeyP      GLFWkey = C.GLFW_KEY_P
	KeyQ      GLFWkey = C.GLFW_KEY_Q
	KeyR      GLFWkey = C.GLFW_KEY_R
	KeyS      GLFWkey = C.GLFW_KEY_S
	KeyT      GLFWkey = C.GLFW_KEY_T
	KeyU      GLFWkey = C.GLFW_KEY_U
	KeyV      GLFWkey = C.GLFW_KEY_V
	KeyW      GLFWkey = C.GLFW_KEY_W
	KeyX      GLFWkey = C.GLFW_KEY_X
	KeyY      GLFWkey = C.GLFW_KEY_Y
	KeyZ      GLFWkey = C.GLFW_KEY_Z
	KeyEscape GLFWkey = C.GLFW_KEY_ESCAPE
)

// enum of glfw key events.
const (
	KeyPress   GLFWbool = C.GLFW_PRESS
	KeyRelease GLFWbool = C.GLFW_RELEASE
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
type _GLFWplatform struct {} 

// This function converts C.integers to booleans, C.int function && variables can be converted.
func (p _GLFWplatform) boolifyfunction(function C.int) bool {
	return function == C.int(TRUE)
}

// This function destroys the GLFWwindow via an abstract function.
func (p _GLFWplatform) destroyWindow() {
	C.DestroyGLFWWindow()
}

// Terminates GLFW itself, by directly calling the function C.glfwTerminate()
func (p _GLFWplatform) terminate() {
	C.glfwTerminate()
}

// Creates the GLFW window object, but does not initialize GLFW.
// You must call `glfw.Init()`
func (p _GLFWplatform) create(window _GLFWwindow) {
	titleconv := C.CString(window.title)
	defer C.free(unsafe.Pointer(titleconv))

	C.CreateWindow(C.int(_window.videoMode.width), C.int(_window.videoMode.height), titleconv)
}

// Swaps the OpenGL Buffers.
func (p _GLFWplatform) swapBuffers() {
	C.SwapBuffers()
}

// Polls the input events by directly calling `C.glfwPollEvents`
func (p _GLFWplatform) pollEvents() {
	C.glfwPollEvents()
}

// Makes the window context current.
func (p _GLFWplatform) makeContextCurrent() {
	C.MakeContextCurrent()
}

// Gets the current key found that was input but needs `glfwPollEvents` to work.
func (p _GLFWplatform) GetKey(key GLFWkey) GLFWbool {
	_keyconv := GLFWbool(C.GetGLFWKey(C.int(key)))
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
func (p _GLFWplatform) windowShouldClose() bool {
	// Converts C's glfwWindowShould (int) to a Go usable boolean.
	_shouldcloseconv := _glfw.platform.boolifyfunction(C.WindowShouldClose())
	return _shouldcloseconv
}

func (p _GLFWplatform) setWindowShouldClose(value int) {
	C.SetWindowShouldClose(C.int(value))
}

type _GLFWwindow struct {
	videoMode _GLFWvideoMode
	title string
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
	return _glfw.platform.initialize()
}

// @fn glfw.CreateWindow
//
// This Creates a GLFWwindow,
// This function unlike the C version does not have monitor, share arguments.
// Why? I don't want to add them yet.
//
// Creates a GLFWwindow using with function that need *GLFWwindow as a parameter.
// Normally if any errors are detected whilst this function is used, C will do its normal error checking.
//
// @params:
//     - width: the width of the window being created.
//     - height: the height of the window being created.
//     - title: the title of the window being created.
func CreateWindow(width int, height int, title string) *GLFWwindow {
	handle := new(GLFWwindow)

	_window.videoMode.width = width
	_window.videoMode.height = height
	_window.title = title

	_glfw.platform.create(_window)

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
//     - window: the window object that will have it's buffers swapped.
func SwapBuffers(window *GLFWwindow) {
	_glfw.platform.swapBuffers()
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
//     - window: the window object that will use the Key.
//     - key: the key specified to have it's events found.
func GetKey(window *GLFWwindow, key GLFWkey) GLFWbool {
	return _glfw.platform.GetKey(key)
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
//     - window: the window object that will have it's context made current.
func MakeContextCurrent(window *GLFWwindow) {
	_glfw.platform.makeContextCurrent()
}

// @fn glfw.SetWindowPos 
// 
// Sets the window position.
// This function sets the position of the window,
// Uses the coordinates `x` and `y`.
// 
// @params:
//     - window: the window object that needs it's position set.
//     - x: the x position of the window object created.
//     - y: the y position of the window object created.
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
//     - window: the window object created that should close.
func WindowShouldClose(window *GLFWwindow) bool {
	return _glfw.platform.windowShouldClose()
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
//     - window: the window object that will be set to close or not.
//     - value: the state of the window object that will be set to close or not.
func SetWindowShouldClose(window *GLFWwindow, value GLFWbool) {
	_glfw.platform.setWindowShouldClose(int(value))
}

// @fn glfw.DestroyWindow
//
// Destroys the window.
// This function destroys the window object that means that all allocated will be deallocated,
// even though the `glfw.Terminate` free everything, this function destroys specific stuff allocated.
// This function unlike the `glfw.Terminate` function should be called after not before the mainLoop.
//
// @params:
//     - window: the window object created that will be destroyed.
func DestroyWindow(window *GLFWwindow) {
	_glfw.platform.destroyWindow()
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
	_glfw.platform.terminate()
}

//==================================================
// lwggl creator marks:
// END OF FILE!
//
// ** NOTHING SHALL BE WRITTEN BEYOND THIS MARKER **
//==================================================