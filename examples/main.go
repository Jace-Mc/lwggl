package examples

import (
	"github.com/Jace-Mc/lwggl/glad"
	"github.com/Jace-Mc/lwggl/glfw"
)

func main() {
	glfw.Init()
	defer glfw.Terminate()

	glfw.WindowHint(glfw.WINDOW_RESIZABLE, glfw.FALSE)
	window := glfw.CreateWindow(400, 400, "my window")
	glfw.MakeContextCurrent(window)
	glad.LoadGL()

	for !glfw.WindowShouldClose(window) {
		glfw.PollEvents()
		glfw.SwapBuffers(window)
	}
}