package external

/*

#ifdef __linux__
#cgo CFLAGS: -D_GLFW_X11 -D_GNU_SOURCE
#cgo LDFLAGS: -lX11 -lGL.
#endif

#include "glfw/src/init.c"
#include "glfw/src/window.c"
#include "glfw/src/platform.c"
#include "glfw/src/input.c"
#include "glfw/src/context.c"
#include "glfw/src/monitor.c"
*/
import "C"
