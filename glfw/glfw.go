package glfw

/*
#cgo CFLAGS: -I ../external/glfw/include
#include <GLFW/glfw3.h>
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

*/
import "C"
import "unsafe"

// enum
const (
	WINDOW_RESIZABLE int = 0x0023
	WINDOW_MINIMIZABLE int = 0x0026
) // WindowHints

type GLFWbool int

// enum 
const (
	TRUE GLFWbool = 1
	FALSE GLFWbool = 0
) // Boolean Types

type GLFWkey int

// enum
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
	KeyRelease GLFWkey = C.GLFW_RELEASE
	KeyPress GLFWkey = C.GLFW_PRESS
) // GLFWKeys

//=======================
// The struct GLFWwindow.
//=======================
type GLFWwindow struct {}

//===========================
// The struct _GLFWvideoMode.
//===========================
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

type _GLFWplatform struct {}

func (p _GLFWplatform) terminate() {
	C.glfwTerminate()
}

func (p _GLFWplatform) create(window _GLFWwindow) {
	titleconv := C.CString(window.videoMode.title)
	defer C.free(unsafe.Pointer(titleconv))

	C.CreateWindow(C.int(_window.videoMode.width), C.int(_window.videoMode.height), titleconv)
	C.SetWindowPosition(C.int(_window.videoMode.x), C.int(_window.videoMode.y))
}

func (p _GLFWplatform) swapBuffers(window _GLFWwindow) {
	C.SwapBuffers()
}

func (p _GLFWplatform) pollEvents() {
	C.glfwPollEvents()
}

func (p _GLFWplatform) makeContextCurrent(window _GLFWwindow) {
	C.MakeContextCurrent()
}

func (p _GLFWplatform) initialize() {
	C.InitGLFW()
}

type _GLFWwindow struct {
	videoMode _GLFWvideoMode
	shouldClose bool
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
func Init() {
	_glfw.platform.initialize()
}

/**
 * Creates a GLFW window.
 */
func CreateWindow(width int, height int, title string) GLFWwindow {
	_window.videoMode.width = width
	_window.videoMode.height = height
	_window.videoMode.title = title

	_glfw.platform.create(_window)

	return GLFWwindow{}
}

/**
 * Swaps OpenGL Buffers.
 */
func SwapBuffers(window GLFWwindow) {
	_glfw.platform.swapBuffers(_window)
}

/**
 * Poll Window Events.
 */
func PollEvents() {
	_glfw.platform.pollEvents()
}

func GetKey(key int) GLFWbool {
	_getkey := GLFWbool(C.GetGLFWKey(C.int(key)))
	return _getkey
}

/**
 * Makes OpenGL Context Current.
 */
func MakeContextCurrent(window GLFWwindow) {
	_glfw.platform.makeContextCurrent(_window)
}

/**
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

/**
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
	return _window.shouldClose
}

/*
 * sets if window should close or not.
 */
func SetWindowShouldClose(window GLFWwindow, value int) {
	C.SetWindowShouldClose(C.int(value))
}

/*
 * terminates GLFW.
 * USAGES: defer glfw.terminate()
 */
func Terminate() {
	_glfw.platform.terminate()
}