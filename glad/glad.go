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

/*
 * This functions loads OpenGL, so that all the opengl functions are usable.
 */
func LoadGL() {
	C.gladLoadGL()
}