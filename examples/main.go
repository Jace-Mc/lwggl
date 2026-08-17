package examples

import (
	GLFW "github.com/Jace-Mc/lwggl/glfw"
	"github.com/Jace-Mc/lwggl/glad"
)

func main() {
	GLFW.Init()
	defer GLFW.Terminate()

	window := GLFW.CreateWindow(400, 400, "My Window")
	GLFW.MakeContextCurrent(window)
	glad.LoadGL()

	for !GLFW.WindowShouldClose(window) {
		if GLFW.GetKey(GLFW.KEY_E) == GLFW.KEY_PRESS {
			GLFW.SetWindowPos(10, 20)
		}
		GLFW.PollEvents()
		GLFW.SwapBuffers(window)
	}
}