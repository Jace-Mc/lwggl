package opengl

/*
#include "../external/glad.c"

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

func Vertex2f(x float32, y float32) {
	C.Vertex2f(C.GLfloat(x), C.GLfloat(y))
}

func Vertex3f(x float32, y float32, z float32) {
	C.Vertex3f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func Color3f(x float32, y float32, z float32) {
	C.Color3f(C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}

func Color4f(w float32, x float32, y float32, z float32) {
	C.Color4f(C.GLfloat(w), C.GLfloat(x), C.GLfloat(y), C.GLfloat(z))
}