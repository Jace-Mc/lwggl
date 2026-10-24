package main

import (
	"fmt"
	"os"

	"github.com/Jace-Mc/lwggl/glfw"
	"github.com/Jace-Mc/lwggl/opengl"
)

func main() {
	if !glfw.Init() {
		fmt.Println("Could not initialize GLFW")
		os.Exit(1)
	}

	defer glfw.Terminate()

	window := glfw.CreateWindow(400, 400, "My lwggl window")
	glfw.MakeContextCurrent(window)

	for !glfw.WindowShouldClose(window) {
		glfw.PollEvents()
		glfw.SwapBuffers(window)
	}

	glfw.DestroyWindow(window)
}
