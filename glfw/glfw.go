package glfw

/*
#include "../external/glfw/include/GLFW/glfw3.h"
#include <stdlib.h>

GLFWwindow* window;

void InitGLFW(void) 
{
	if (!glfwInit()) exit(1); 
}

void CreateWindow(int width, int height, const char* title)
{
	window int = glfwCreateWindow(width, height, title, NULL, NULL);
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

*/
import "C"
import "unsafe"

type GLFWbool = int

const (
	WINDOW_RESIZABLE int = 0x0023
	WINDOW_MINIMIZABLE int = 0x0026
)

const (
	TRUE = 1
	FALSE = 0
)

const (
	KEY_A int = C.GLFW_KEY_A
	KEY_B int = C.GLFW_KEY_B
	KEY_C int = C.GLFW_KEY_C
	KEY_D int = C.GLFW_KEY_D
	KEY_E int = C.GLFW_KEY_E
	KEY_F int = C.GLFW_KEY_F
	KEY_G int = C.GLFW_KEY_G
	KEY_H int = C.GLFW_KEY_H
	KEY_I int = C.GLFW_KEY_I
	KEY_J int = C.GLFW_KEY_J
	KEY_K int = C.GLFW_KEY_K
	KEY_L int = C.GLFW_KEY_L
	KEY_M int = C.GLFW_KEY_M
	KEY_N int = C.GLFW_KEY_N
	KEY_O int = C.GLFW_KEY_O
	KEY_P int = C.GLFW_KEY_P
	KEY_Q int = C.GLFW_KEY_Q
	KEY_R int = C.GLFW_KEY_R
	KEY_S int = C.GLFW_KEY_S
	KEY_T int = C.GLFW_KEY_T
	KEY_U int = C.GLFW_KEY_U
	KEY_V int = C.GLFW_KEY_V
	KEY_W int = C.GLFW_KEY_W
	KEY_X int = C.GLFW_KEY_X
	KEY_Y int = C.GLFW_KEY_Y
	KEY_Z int = C.GLFW_KEY_Z
	KEY_RELEASE int = C.GLFW_KEY_RELEASE
	KEY_PRESS int = C.GLFW_KEY_PRESS
)

type GLFWwindow struct {}

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
	return C.glfwGetKey(C.int(key))
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
			_glfw.hints.resizable = value

		case WINDOW_MINIMIZABLE:
			_glfw.hints.minimizable = value
	}
}

/**
 * Sets Window Position. 
 */
func SetWindowPos(x int, y int) {
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