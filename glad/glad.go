// Welcome to lwggl.
//
// This is the glad/glad.go file, this holds all the glad bindings.
//
// Link to the GLAD github repository: [NIL]
//
// DEPENDENCIES: (all are included)
//
// [GLAD] ~ Has OpenGL Loaders, needed for OpenGL.
//
// LICENSE: zlib/libpng
// The license is shown in `LICENSE`
package glad

/*
// The GLAD header include.
#include "../external/glad.h"
*/
import "C"

// @fn glad.LoadGL
// 
// Loads all OpenGL, and GLAD, procs, functions, enums, etc.
// This allows the user to use all the OpenGL functions without any problems doing so.
//
// < No parameters >
func LoadGL() {
	C.gladLoadGL()
}