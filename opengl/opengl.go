// Welcome to lwggl.
//
// This is the opengl/opengl.go file, this holds all OpenGL bindings.
// Will be adding more functions to this module.
// Since most glad functions are defines.
//
// DEPENDENCIES: (all are included.)
//
// [OpenGL] ~ Holds all Graphics functions that have something to use with OpenGL.
//
// LICENSE: zlib/libpng
// The license is shown in `LICENSE`
package opengl

/*
// The GLAD header include.
#include "../external/glad.h"

// Since cgo does not except macros as functions,
// they have to be converted to void functions.

// Converts 'GLvertex2f' to go usable.
void Vertex2f(GLfloat x, GLfloat y)
{
	glVertex2f(x, y);
}

// Converts 'GLvertex3f' to go usable.
void Vertex3f(GLfloat x, GLfloat y, GLfloat z)
{
	glVertex3f(x, y, z);
}

// Converts 'GLcolor3f' to go usable.
void Color3f(GLfloat x, GLfloat y, GLfloat z)
{
	glColor3f(x, y, z);
}

// Converts 'GLcolor4f' to go usable.
void Color4f(GLfloat w, GLfloat x, GLfloat y, GLfloat z)
{
	glColor4f(w, x, y, z);
}

// Converts 'GLclear' to go usable.
void Clear(GLenum mode)
{
	glClear(mode);
}

// Converts 'GLclearcolor' to go usable
void ClearColor(float r, float g, float b, float a)
{
	glClearColor(r, g, b, a);
}

*/
import "C"

type clampf float32
type float float32

type GLenum uint

const (
	ColorBufferBit GLenum = C.GL_COLOR_BUFFER_BIT
)

func Vertex2f(x float, y float) {
	C.Vertex2f(C.GLfloat(x), C.GLfloat(y))
}

func Vertex3f(x float, y float, z float) {
	C.Vertex3f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func Color3f(x clampf, y clampf, z clampf) {
	C.Color3f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func Color4f(w clampf, x clampf, y clampf, z clampf) {
	C.Color4f(C.GLfloat(w), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func Clear(mode GLenum) {
	C.Clear(C.GLenum(mode))
}

func ClearColor(r float, g float, b float, a float) {
	C.ClearColor(C.float(r), C.float(g), C.float(b), C.float(a))
}