package glfw

/*
#ifdef __linux__
	#cgo CFLAGS: -D_GLFW_X11 -D_GNU_SOURCE
	#cgo LDFLAGS: -lX11 -lGL -lEGL -lm
#endif

#include "../external/glfw/src/init.c"
#include "../external/glfw/src/window.c"
#include "../external/glfw/src/platform.c"
#include "../external/glfw/src/input.c"
#include "../external/glfw/src/context.c"
#include "../external/glfw/src/monitor.c"
#include "../external/glfw/src/vulkan.c"
#include "../external/glfw/src/null_init.c"
#include "../external/glfw/src/null_joystick.c"
#include "../external/glfw/src/null_monitor.c"
	
#ifdef __linux__
	#ifdef _GLFW_X11	
		#include "../external/glfw/src/x11_init.c"
		#include "../external/glfw/src/x11_window.c"
		#include "../external/glfw/src/x11_monitor.c"
		#include "../external/glfw/src/xkb_unicode.c"
		#include "../external/glfw/src/glx_context.c"
	#endif

	#ifdef _GLFW_WAYLAND 
		#include "../external/glfw/src/wl_init.c"
		#include "../external/glfw/src/wl_window.c"
		#include "../external/glfw/src/wl_monitor.c"
	#endif

	#include "../external/glfw/src/posix_thread.c"
	#include "../external/glfw/src/posix_module.c"
	#include "../external/glfw/src/posix_time.c"
	#include "../external/glfw/src/posix_poll.c"

	#include "../external/glfw/src/egl_context.c"
	#include "../external/glfw/src/osmesa_context.c"

	#include "../external/glfw/src/linux_joystick.c"	
#endif
*/
import "C"
