# LWGGL
- Welcome to lwggl.
- lwggl stands for light weight Go game library.
- A easy - to - use game library for Go.
- This library has bindings for GLFW, glad, and OpenGL.

# dependencies: (all are included)
- all dependencies are included.
- GLFW-3.6.1 ~ For OpenGL, Window and Inputs
- GLAD ~ For loading OpenGL extensions.

# needed dependencies (to be installed by the user):
- Go ~ must be installed for using the library. 
- you can install Go from: https://go.dev
- A C compiler: clang, gcc.
- CGO must be enabled.
- OpenGL must be installed.

# installation
- In your command line:
```sh
# for lwggl's run: 
go get github.com/Jace-Mc/lwggl
```

# examples
- More examples will be found in the tests directory.
```go
package main
import (
    "github.com/Jace-Mc/lwggl/glfw"
    "github.com/Jace-Mc/lwggl/glad"
    "github.com/Jace-Mc/lwggl/opengl"
)

func main() {
    glfw.Init()
    defer glfw.Terminate()

    window := glfw.CreateWindow(200, 200, "My First lwggl Window")
    glfw.MakeContextCurrent(window)
    glad.LoadGL()

    for !glfw.WindowShouldClose(window) {
        glfw.PollEvents()
        glfw.SwapBuffers(window)
    }
}
```

# changelog:
## v1.0.5:
* Added function `GetTime` (gets the current time), can be used for fps targets, getting time, etc.
* Fixed function/s `GetKey`
* Added function/s `_glfw.platform.destroyWindow` and DestroyWindow these both are an abstraction of GLFW's glfwDestroyWindow function.
* Added the section `lwggl functions` for showing all the functions in lwggl
* Properly configured `external/lwglfw.go` and `external/lwextra.go`, so `cgo` does not have any errors
* Renamed `external/lwbuild.go` to `external/lwextra.go`
## v1.0.4
* Added function/s: `_glfw.platform.boolifyfunction` (can boolify functions, and variables)                    
## v1.0.3                
* Removed `utils/` and `utils.go`
* Fixed how comments were describing functions, and enums (const ()), and structs.
## v1.0.2
* Removed the _GLFWwindow variable `shouldClose`, now using `C.WindowShouldClose` 
* Bug Fixes (if there are bugs).
## v1.0.1
* Both OpenGL, and GLAD are up and running.

# lwggl functions:
```go
// In lwggl v1.0.5:
func Init() bool 
func CreateWindow(width int, height int, title string) GLFWwindow 
func SwapBuffers(window GLFWwindow) 
func PollEvents()
func GetKey(window GLFWwindow, key GLFWkey) GLFWbool 
func GetTime() GLFWtime 
func MakeContextCurrent(window GLFWwindow) 
func WindowHint(Type int, value int)
func SetWindowPos(window GLFWwindow, x int, y int) 
func WindowShouldClose(window GLFWwindow) bool 
func SetWindowShouldClose(window GLFWwindow, value GLFWbool) 
func DestroyWindow(window GLFWwindow) 
func Terminate() 
```
------------------------------------
* End of changelog