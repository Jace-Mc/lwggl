package external

// #cgo CFLAGS: -D_GLFW_X11 -D_GNU_SOURCE
// #cgo LDFLAGS: -lX11 -lGL
// #include "lwglfw.c"
import "C"
