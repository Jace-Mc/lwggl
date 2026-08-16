package glfw

type GLFWbool = int

const (
	WINDOW_RESIZABLE = 0x0030
	WINDOW_MINIMIZABLE = 0x0060
)

const (
	TRUE = 1
	FALSE = 0
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

func (p _GLFWplatform) terminate() {}
func (p _GLFWplatform) create(window _GLFWwindow) {}
func (p _GLFWplatform) swapBuffers(window _GLFWwindow) {}
func (p _GLFWplatform) pollEvents() {}
func (p _GLFWplatform) makeContextCurrent(window _GLFWwindow) {}

type _GLFWwindow struct {
	videoMode _GLFWvideoMode

	shouldClose bool
}

type _GLFWlibrary struct {
	platform _GLFWplatform
	hints _GLFWhints
	initialize bool
}

var _glfw _GLFWlibrary = _GLFWlibrary{}
var _window _GLFWwindow = _GLFWwindow{}

/**
 * initializes GLFW.
 */ 
func Init() bool {
	return _glfw.initialize
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
func SetWindowShouldClose(window GLFWwindow, value bool) {
	_window.shouldClose = value
}

/*
 * terminates GLFW.
 * USAGES: defer glfw.terminate()
 */
func Terminate() {
	_glfw.platform.terminate()
}