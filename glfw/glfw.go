/*
 Welcome to lwggl.

 This is the glfw/glfw.go file, this converts the C GLFW to the GO GLFW.
 No librarys are external that includes:
 GLAD and GLFW, and any other external librarys.
 
 Link to the GLFW website:
 https://www.glfw.org 

 Link to the GLAD github repository:

 DEPENDENCIES: (all are included)
 [GLAD] ~ For OpenGL loading.
 [GLFW] ~ For OpenGL Window, context, and input.
 
 LICENSE: zlib/libpng 
 The license is shown in `LICENSE`
 */
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

// enum of window hints
const (
	WINDOW_RESIZABLE int = 0x0023
	WINDOW_MINIMIZABLE int = 0x0026
) 

// Typedef for int (GLFWbool)
type GLFWbool int

// Typedef for float64(double) (GLFWtime)
type GLFWtime float64

// enum of boolean types
const (
	TRUE GLFWbool = C.GLFW_TRUE
	FALSE GLFWbool = C.GLFW_FALSE
) 

// Typedef for int (GLFWkey)
type GLFWkey int

// enum of glfw keys.
const (
	KeyA GLFWkey = C.GLFW_KEY_A
	KeyB GLFWkey = C.GLFW_KEY_B
	KeyC GLFWkey = C.GLFW_KEY_C	
	KeyD GLFWkey = C.GLFW_KEY_D	
	KeyE GLFWkey = C.GLFW_KEY_E	
	KeyF GLFWkey = C.GLFW_KEY_F	
	KeyG GLFWkey = C.GLFW_KEY_G
	KeyH GLFWkey = C.GLFW_KEY_H	
	KeyI GLFWkey = C.GLFW_KEY_I	
	KeyJ GLFWkey = C.GLFW_KEY_J	
	KeyK GLFWkey = C.GLFW_KEY_K
	KeyL GLFWkey = C.GLFW_KEY_L
	KeyM GLFWkey = C.GLFW_KEY_M
	KeyN GLFWkey = C.GLFW_KEY_N
	KeyO GLFWkey = C.GLFW_KEY_O
	KeyP GLFWkey = C.GLFW_KEY_P
	KeyQ GLFWkey = C.GLFW_KEY_Q
	KeyR GLFWkey = C.GLFW_KEY_R
	KeyS GLFWkey = C.GLFW_KEY_S
	KeyT GLFWkey = C.GLFW_KEY_T
	KeyU GLFWkey = C.GLFW_KEY_U
	KeyV GLFWkey = C.GLFW_KEY_V
	KeyW GLFWkey = C.GLFW_KEY_W
	KeyX GLFWkey = C.GLFW_KEY_X
	KeyY GLFWkey = C.GLFW_KEY_Y
	KeyZ GLFWkey = C.GLFW_KEY_Z
	KeyEscape GLFWkey = C.GLFW_KEY_ESCAPE
)

// enum of Key events.
const (
	KeyPress GLFWbool = C.GLFW_PRESS
	KeyRelease GLFWbool = C.GLFW_RELEASE
) 

// The struct GLFWwindow.
// This is a handle struct so no attributes.
type GLFWwindow struct {}

// The struct _GLFWvideoMode.
// This is a core struct so there are (5) attributes.
type _GLFWvideoMode struct {
	width int
	height int
	title string
	x int
	y int
}

type _GLFWhints struct {
	resizable GLFWbool
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
	titleconv := C.CString(window.videoMode.title)
	defer C.free(unsafe.Pointer(titleconv))

	C.CreateWindow(C.int(_window.videoMode.width), C.int(_window.videoMode.height), titleconv)
}

// Swaps the OpenGL Buffers.
func (p _GLFWplatform) swapBuffers(window _GLFWwindow) {
	C.SwapBuffers()
}

// Polls the input events by directly calling `C.glfwPollEvents`
func (p _GLFWplatform) pollEvents() {
	C.glfwPollEvents()
}

// Makes the window context current.
func (p _GLFWplatform) makeContextCurrent(window _GLFWwindow) {
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
	_initconv  := _glfw.platform.boolifyfunction(C.glfwInit())
	return _initconv  
}

type _GLFWwindow struct {
	videoMode _GLFWvideoMode
}

type _GLFWlibrary struct {
	platform _GLFWplatform
	hints _GLFWhints
}

var _glfw _GLFWlibrary = _GLFWlibrary{}
var _window _GLFWwindow = _GLFWwindow{}

/**
 * initializes GLFW.
 */ 
func Init() bool {
	return _glfw.platform.initialize()
}

/*
 * Creates a GLFW window.
 */
func CreateWindow(width int, height int, title string) GLFWwindow {
	_window.videoMode.width = width
	_window.videoMode.height = height
	_window.videoMode.title = title

	_glfw.platform.create(_window)

	return GLFWwindow{}
}

/*
 * Swaps OpenGL Buffers.
 */
func SwapBuffers(window GLFWwindow) {
	_glfw.platform.swapBuffers(_window)
}

/*
 * Poll Window Events.
 */
func PollEvents() {
	_glfw.platform.pollEvents()
}

/*
 * Gets the Key found via the event.
 */
func GetKey(window GLFWwindow, key GLFWkey) GLFWbool {
	return _glfw.platform.GetKey(key)
}

/*
 * Gets the time since the window started drawing.
 */
func GetTime() GLFWtime {
	return _glfw.platform.GetTime()
}

/*
 * Makes OpenGL Context Current.
 */
func MakeContextCurrent(window GLFWwindow) {
	_glfw.platform.makeContextCurrent(_window)
}

/*
 * Hints stuff to the window. 
 */
func WindowHint(Type int, value int) {
	switch (Type) {
		case WINDOW_RESIZABLE:
			_glfw.hints.resizable = GLFWbool(value)

		case WINDOW_MINIMIZABLE:
			_glfw.hints.minimizable = GLFWbool(value) 
	}
}

/*
 * Sets Window Position. 
 */
func SetWindowPos(window GLFWwindow, x int, y int) {
	_window.videoMode.x = x
	_window.videoMode.y = y
}

/* 
 * Checks if windowshouldclose or not.
 */
func WindowShouldClose(window GLFWwindow) bool {
	_booleanconv := _glfw.platform.boolifyfunction(C.WindowShouldClose())
	return _booleanconv
}

/*
 * sets if window should close or not.
 */
func SetWindowShouldClose(window GLFWwindow, value GLFWbool) {
	C.SetWindowShouldClose(C.int(value))
}

/*
 * destroys the GLFW window.
 */
func DestroyWindow(window GLFWwindow) {
	_glfw.platform.destroyWindow()
}

/*******************************
 * terminates GLFW.
 * USAGES: defer glfw.terminate()
 */
func Terminate() {
	_glfw.platform.terminate()
}

//==================================================
// lwggl creator marks:
// END OF FILE!
// 
// ** NOTHING SHALL BE WRITTEN BEYOND THIS MARKER **
//==================================================