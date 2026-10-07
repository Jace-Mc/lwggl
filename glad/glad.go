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