package examples
import "github.com/Jace-Mc/lwggl/glfw"

func main() {
	glfw.Init()
	defer glfw.Terminate()

	window := glfw.CreateWindow(400, 400, "My lwggl window")
	glfw.MakeContextCurrent(window)

	for !glfw.WindowShouldClose(window) {
		glfw.PollEvents()
		glfw.SwapBuffers(window)
	}
}