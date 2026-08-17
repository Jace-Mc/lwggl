package examples

import (
	"github.com/Jace-Mc/lwggl/glfw"
	"github.com/Jace-Mc/lwggl/glad"
)

func main() {
	glfw.Init()
	defer glfw.Terminate()

	window := glfw.CreateWindow(400, 400, "My GLFW window")
	glfw.MakeContextCurrent(window)
	glad.LoadGL()

	for !glfw.WindowShouldClose(window) {
		glfw.PollEvents()
		glfw.SwapBuffers(window)
	}
}
