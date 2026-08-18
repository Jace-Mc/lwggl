# LWGGL
- Welcome to lwggl.
- lwggl stands for light weight Go game library.
- A easy - to - use game library for Go.
- This library has bindings for GLFW, glad, and OpenGL.

# dependencies: (all are included)
- all dependencies are included.
- GLFW-3.6.1 ~ For OpenGL, Window and Inputs
- GLAD ~ For loading OpenGL extensions.
- OpenGL ~ Included with your operating system.

# needed dependencies:
- Go ~ must be installed for using the library. 
- you can install Go from: https://go.dev
- A C compiler: clang, gcc.

# installation
- In your command line:
```sh
# for GLFW, GLAD, and OpenGL run: 
go get github.com/Jace-Mc/lwggl
```

# examples
- More examples will be found in the examples directory.
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

    window := glfw.CreateWindow(200, 200, "My Window")
    glfw.MakeContextCurrent(window)
    glad.LoadGL()

    for !glfw.WindowShouldClose(window) {
        glfw.PollEvents()
        glfw.SwapBuffers(window)
    }
}
```

