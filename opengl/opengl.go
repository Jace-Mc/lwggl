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

*/
import "C"

type float float32

func Vertex2f(x float, y float) {
	C.Vertex2f(C.GLfloat(x), C.GLfloat(y))
}

func Vertex3f(x float, y float, z float) {
	C.Vertex3f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func Color3f(x float, y float, z float) {
	C.Color3f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func Color4f(w float, x float, y float, z float) {
	C.Color4f(C.GLfloat(w), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}