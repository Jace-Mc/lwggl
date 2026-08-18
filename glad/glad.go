package glad

// #include "../external/glad.c"
import "C"

/**
 * This functions loads OpenGL, so that all the opengl functions are usable.
 */
func LoadGL() {
	C.gladLoadGL()
}