/*
Portions of this file are copied from the SDL3 headers. The SDL3 license is
reproduced below:

This software is provided 'as-is', without any express or implied
warranty.  In no event will the authors be held liable for any damages
arising from the use of this software.

Permission is granted to anyone to use this software for any purpose,
including commercial applications, and to alter it and redistribute it
freely, subject to the following restrictions:

1. The origin of this software must not be misrepresented; you must not
   claim that you wrote the original software. If you use this software
   in a product, an acknowledgment in the product documentation would be
   appreciated but is not required.
2. Altered source versions must be plainly marked as such, and must not be
   misrepresented as being the original software.
3. This notice may not be removed or altered from any source distribution.
*/

package sdl

// #cgo pkg-config: sdl3
// #cgo noescape SDL_GetNumVideoDrivers
// #cgo nocallback SDL_GetNumVideoDrivers
// #cgo noescape SDL_GetVideoDriver
// #cgo nocallback SDL_GetVideoDriver
// #cgo noescape SDL_GetCurrentVideoDriver
// #cgo nocallback SDL_GetCurrentVideoDriver
// #cgo noescape SDL_GetSystemTheme
// #cgo nocallback SDL_GetSystemTheme
// #cgo noescape SDL_GetDisplays
// #cgo nocallback SDL_GetDisplays
// #cgo noescape SDL_GetPrimaryDisplay
// #cgo nocallback SDL_GetPrimaryDisplay
// #cgo noescape SDL_GetDisplayProperties
// #cgo nocallback SDL_GetDisplayProperties
// #cgo noescape SDL_GetDisplayName
// #cgo nocallback SDL_GetDisplayName
// #cgo noescape SDL_GetDisplayBounds
// #cgo nocallback SDL_GetDisplayBounds
// #cgo noescape SDL_GetDisplayUsableBounds
// #cgo nocallback SDL_GetDisplayUsableBounds
// #cgo noescape SDL_GetNaturalDisplayOrientation
// #cgo nocallback SDL_GetNaturalDisplayOrientation
// #cgo noescape SDL_GetCurrentDisplayOrientation
// #cgo nocallback SDL_GetCurrentDisplayOrientation
// #cgo noescape SDL_GetDisplayContentScale
// #cgo nocallback SDL_GetDisplayContentScale
// #cgo noescape SDL_GetFullscreenDisplayModes
// #cgo nocallback SDL_GetFullscreenDisplayModes
// #cgo noescape SDL_GetClosestFullscreenDisplayMode
// #cgo nocallback SDL_GetClosestFullscreenDisplayMode
// #cgo noescape SDL_GetDesktopDisplayMode
// #cgo nocallback SDL_GetDesktopDisplayMode
// #cgo noescape SDL_GetCurrentDisplayMode
// #cgo nocallback SDL_GetCurrentDisplayMode
// #cgo noescape SDL_GetDisplayForPoint
// #cgo nocallback SDL_GetDisplayForPoint
// #cgo noescape SDL_GetDisplayForRect
// #cgo nocallback SDL_GetDisplayForRect
// #cgo noescape SDL_GetDisplayForWindow
// #cgo nocallback SDL_GetDisplayForWindow
// #cgo noescape SDL_GetWindowPixelDensity
// #cgo nocallback SDL_GetWindowPixelDensity
// #cgo noescape SDL_GetWindowDisplayScale
// #cgo nocallback SDL_GetWindowDisplayScale
// #cgo noescape SDL_SetWindowFullscreenMode
// #cgo nocallback SDL_SetWindowFullscreenMode
// #cgo noescape SDL_GetWindowFullscreenMode
// #cgo nocallback SDL_GetWindowFullscreenMode
// #cgo noescape SDL_GetWindowICCProfile
// #cgo nocallback SDL_GetWindowICCProfile
// #cgo noescape SDL_GetWindowPixelFormat
// #cgo nocallback SDL_GetWindowPixelFormat
// #cgo noescape SDL_GetWindows
// #cgo nocallback SDL_GetWindows
// #cgo noescape SDL_CreateWindow
// #cgo nocallback SDL_CreateWindow
// #cgo noescape SDL_CreatePopupWindow
// #cgo nocallback SDL_CreatePopupWindow
// #cgo noescape SDL_CreateWindowWithProperties
// #cgo nocallback SDL_CreateWindowWithProperties
// #cgo noescape SDL_GetWindowID
// #cgo nocallback SDL_GetWindowID
// #cgo noescape SDL_GetWindowFromID
// #cgo nocallback SDL_GetWindowFromID
// #cgo noescape SDL_GetWindowParent
// #cgo nocallback SDL_GetWindowParent
// #cgo noescape SDL_GetWindowProperties
// #cgo nocallback SDL_GetWindowProperties
// #cgo noescape SDL_GetWindowFlags
// #cgo nocallback SDL_GetWindowFlags
// #cgo noescape SDL_SetWindowTitle
// #cgo nocallback SDL_SetWindowTitle
// #cgo noescape SDL_GetWindowTitle
// #cgo nocallback SDL_GetWindowTitle
// #cgo noescape SDL_SetWindowIcon
// #cgo nocallback SDL_SetWindowIcon
// #cgo noescape SDL_SetWindowPosition
// #cgo nocallback SDL_SetWindowPosition
// #cgo noescape SDL_GetWindowPosition
// #cgo nocallback SDL_GetWindowPosition
// #cgo noescape SDL_SetWindowSize
// #cgo nocallback SDL_SetWindowSize
// #cgo noescape SDL_GetWindowSize
// #cgo nocallback SDL_GetWindowSize
// #cgo noescape SDL_GetWindowSafeArea
// #cgo nocallback SDL_GetWindowSafeArea
// #cgo noescape SDL_SetWindowAspectRatio
// #cgo nocallback SDL_SetWindowAspectRatio
// #cgo noescape SDL_GetWindowAspectRatio
// #cgo nocallback SDL_GetWindowAspectRatio
// #cgo noescape SDL_GetWindowBordersSize
// #cgo nocallback SDL_GetWindowBordersSize
// #cgo noescape SDL_GetWindowSizeInPixels
// #cgo nocallback SDL_GetWindowSizeInPixels
// #cgo noescape SDL_SetWindowMinimumSize
// #cgo nocallback SDL_SetWindowMinimumSize
// #cgo noescape SDL_GetWindowMinimumSize
// #cgo nocallback SDL_GetWindowMinimumSize
// #cgo noescape SDL_SetWindowMaximumSize
// #cgo nocallback SDL_SetWindowMaximumSize
// #cgo noescape SDL_GetWindowMaximumSize
// #cgo nocallback SDL_GetWindowMaximumSize
// #cgo noescape SDL_SetWindowBordered
// #cgo nocallback SDL_SetWindowBordered
// #cgo noescape SDL_SetWindowResizable
// #cgo nocallback SDL_SetWindowResizable
// #cgo noescape SDL_SetWindowAlwaysOnTop
// #cgo nocallback SDL_SetWindowAlwaysOnTop
// #cgo noescape SDL_ShowWindow
// #cgo nocallback SDL_ShowWindow
// #cgo noescape SDL_HideWindow
// #cgo nocallback SDL_HideWindow
// #cgo noescape SDL_RaiseWindow
// #cgo nocallback SDL_RaiseWindow
// #cgo noescape SDL_MaximizeWindow
// #cgo nocallback SDL_MaximizeWindow
// #cgo noescape SDL_MinimizeWindow
// #cgo nocallback SDL_MinimizeWindow
// #cgo noescape SDL_RestoreWindow
// #cgo nocallback SDL_RestoreWindow
// #cgo noescape SDL_SetWindowFullscreen
// #cgo nocallback SDL_SetWindowFullscreen
// #cgo noescape SDL_SyncWindow
// #cgo nocallback SDL_SyncWindow
// #cgo noescape SDL_WindowHasSurface
// #cgo nocallback SDL_WindowHasSurface
// #cgo noescape SDL_GetWindowSurface
// #cgo noescape SDL_SetWindowSurfaceVSync
// #cgo nocallback SDL_SetWindowSurfaceVSync
// #cgo noescape SDL_GetWindowSurfaceVSync
// #cgo nocallback SDL_GetWindowSurfaceVSync
// #cgo noescape SDL_UpdateWindowSurface
// #cgo nocallback SDL_UpdateWindowSurface
// #cgo noescape SDL_UpdateWindowSurfaceRects
// #cgo nocallback SDL_UpdateWindowSurfaceRects
// #cgo noescape SDL_DestroyWindowSurface
// #cgo noescape SDL_SetWindowKeyboardGrab
// #cgo nocallback SDL_SetWindowKeyboardGrab
// #cgo noescape SDL_SetWindowMouseGrab
// #cgo nocallback SDL_SetWindowMouseGrab
// #cgo noescape SDL_GetWindowKeyboardGrab
// #cgo nocallback SDL_GetWindowKeyboardGrab
// #cgo noescape SDL_GetWindowMouseGrab
// #cgo nocallback SDL_GetWindowMouseGrab
// #cgo noescape SDL_GetGrabbedWindow
// #cgo nocallback SDL_GetGrabbedWindow
// #cgo noescape SDL_SetWindowMouseRect
// #cgo nocallback SDL_SetWindowMouseRect
// #cgo noescape SDL_GetWindowMouseRect
// #cgo nocallback SDL_GetWindowMouseRect
// #cgo noescape SDL_SetWindowOpacity
// #cgo nocallback SDL_SetWindowOpacity
// #cgo noescape SDL_GetWindowOpacity
// #cgo nocallback SDL_GetWindowOpacity
// #cgo noescape SDL_SetWindowParent
// #cgo nocallback SDL_SetWindowParent
// #cgo noescape SDL_SetWindowModal
// #cgo nocallback SDL_SetWindowModal
// #cgo noescape SDL_SetWindowFocusable
// #cgo nocallback SDL_SetWindowFocusable
// #cgo noescape SDL_ShowWindowSystemMenu
// #cgo nocallback SDL_ShowWindowSystemMenu
// #cgo noescape SDL_SetWindowHitTest
// #cgo nocallback SDL_SetWindowHitTest
// #cgo noescape wrap_SDL_SetWindowHitTest
// #cgo nocallback wrap_SDL_SetWindowHitTest
// #cgo noescape SDL_SetWindowShape
// #cgo nocallback SDL_SetWindowShape
// #cgo noescape SDL_FlashWindow
// #cgo nocallback SDL_FlashWindow
// #cgo noescape SDL_DestroyWindow
// #cgo noescape SDL_ScreenSaverEnabled
// #cgo nocallback SDL_ScreenSaverEnabled
// #cgo noescape SDL_EnableScreenSaver
// #cgo nocallback SDL_EnableScreenSaver
// #cgo noescape SDL_DisableScreenSaver
// #cgo nocallback SDL_DisableScreenSaver
// #cgo noescape SDL_GL_LoadLibrary
// #cgo nocallback SDL_GL_LoadLibrary
// #cgo noescape SDL_GL_GetProcAddress
// #cgo nocallback SDL_GL_GetProcAddress
// #cgo noescape SDL_EGL_GetProcAddress
// #cgo nocallback SDL_EGL_GetProcAddress
// #cgo noescape SDL_GL_UnloadLibrary
// #cgo nocallback SDL_GL_UnloadLibrary
// #cgo noescape SDL_GL_ExtensionSupported
// #cgo nocallback SDL_GL_ExtensionSupported
// #cgo noescape SDL_GL_ResetAttributes
// #cgo nocallback SDL_GL_ResetAttributes
// #cgo noescape SDL_GL_SetAttribute
// #cgo nocallback SDL_GL_SetAttribute
// #cgo noescape SDL_GL_GetAttribute
// #cgo nocallback SDL_GL_GetAttribute
// #cgo noescape SDL_GL_CreateContext
// #cgo nocallback SDL_GL_CreateContext
// #cgo noescape SDL_GL_MakeCurrent
// #cgo nocallback SDL_GL_MakeCurrent
// #cgo noescape SDL_GL_GetCurrentWindow
// #cgo nocallback SDL_GL_GetCurrentWindow
// #cgo noescape SDL_GL_GetCurrentContext
// #cgo nocallback SDL_GL_GetCurrentContext
// #cgo noescape SDL_EGL_GetCurrentDisplay
// #cgo noescape SDL_EGL_GetCurrentConfig
// #cgo noescape SDL_EGL_GetWindowSurface
// #cgo noescape SDL_GL_SetSwapInterval
// #cgo nocallback SDL_GL_SetSwapInterval
// #cgo noescape SDL_GL_GetSwapInterval
// #cgo nocallback SDL_GL_GetSwapInterval
// #cgo noescape SDL_GL_SwapWindow
// #cgo nocallback SDL_GL_SwapWindow
// #cgo noescape SDL_GL_DestroyContext
// #cgo nocallback SDL_GL_DestroyContext
// #include <SDL3/SDL.h>
// extern bool wrap_SDL_SetWindowHitTest(SDL_Window *window, uintptr_t userdata);
// extern void wrap_SDL_EGL_SetAttributeCallbacks(uintptr_t userdata);
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

// # CategoryVideo
//
// SDL's video subsystem is largely interested in abstracting window
// management from the underlying operating system. You can create windows,
// manage them in various ways, set them fullscreen, and get events when
// interesting things happen with them, such as the mouse or keyboard
// interacting with a window.
//
// The video subsystem is also interested in abstracting away some
// platform-specific differences in OpenGL: context creation, swapping
// buffers, etc. This may be crucial to your app, but also you are not
// required to use OpenGL at all. In fact, SDL can provide rendering to those
// windows as well, either with an easy-to-use
// [2D API](https://wiki.libsdl.org/SDL3/CategoryRender)
// or with a more-powerful
// [GPU API](https://wiki.libsdl.org/SDL3/CategoryGPU)
// . Of course, it can simply get out of your way and give you the window
// handles you need to use Vulkan, Direct3D, Metal, or whatever else you like
// directly, too.
//
// The video subsystem covers a lot of functionality, out of necessity, so it
// is worth perusing the list of functions just to see what's available, but
// most apps can get by with simply creating a window and listening for
// events, so start with [CreateWindow] and [PollEvent].

// This is a unique ID for a display for the time it is connected to the
// system, and is never reused for the lifetime of the application.
//
// If the display is disconnected and reconnected, it will get a new ID.
//
// The value 0 is an invalid ID.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DisplayID
type DisplayID uint32

// This is a unique ID for a window.
//
// The value 0 is an invalid ID.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WindowID
type WindowID uint32

// The pointer to the global wl_display object used by the Wayland video
// backend.
//
// Can be set before the video subsystem is initialized to import an external
// wl_display object from an application or toolkit for use in SDL, or read
// after initialization to export the wl_display used by the Wayland video
// backend. Setting this property after the video subsystem has been
// initialized has no effect, and reading it when the video subsystem is
// uninitialized will either return the user provided value, if one was set
// prior to initialization, or NULL. See docs/README-wayland.md for more
// information.
const PropGlobalVideoWaylandWLDisplayPointer = "SDL.video.wayland.wl_display"

// System theme.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SystemTheme
type SystemTheme uint32

const (
	SystemThemeUnknown SystemTheme = iota // Unknown system theme
	SystemThemeLight                      // Light colored system theme
	SystemThemeDark                       // Dark colored system theme
)

// The structure that defines a display mode.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DisplayMode
type DisplayMode struct {
	DisplayID              DisplayID   // the display this mode is associated with
	Format                 PixelFormat // pixel format
	W                      int         // width
	H                      int         // height
	PixelDensity           float32     // scale converting size to pixels (e.g. a 1920x1080 mode with 2.0 scale would have 3840x2160 pixels)
	RefreshRate            float32     // refresh rate (or 0.0f for unspecified)
	RefreshRateNumerator   int         // precise refresh rate numerator (or 0 for unspecified)
	RefreshRateDenominator int         // precise refresh rate denominator
	internal               *C.SDL_DisplayModeData
}

// Display orientation values; the way a display is rotated.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DisplayOrientation
type DisplayOrientation uint32

const (
	OrientationUnknown          DisplayOrientation = iota // The display orientation can't be determined
	OrientationLandscape                                  // The display is in landscape mode, with the right side up, relative to portrait mode
	OrientationLandscapeFlipped                           // The display is in landscape mode, with the left side up, relative to portrait mode
	OrientationPortrait                                   // The display is in portrait mode
	OrientationPortraitFlipped                            // The display is in portrait mode, upside down
)

// The struct used as an opaque handle to a window.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Window
type Window C.struct_SDL_Window

// The flags on a window.
//
// These cover a lot of true/false, or on/off, window state. Some of it is
// immutable after being set through [CreateWindow], some of it can be
// changed on existing windows by the app, and some of it might be altered by
// the user or system outside of the app's control.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WindowFlags
type WindowFlags uint64

const (
	WindowFullscreen        WindowFlags = 0x0000000000000001 //window is in fullscreen mode
	WindowOpenGL            WindowFlags = 0x0000000000000002 //window usable with OpenGL context
	WindowOccluded          WindowFlags = 0x0000000000000004 //window is occluded
	WindowHidden            WindowFlags = 0x0000000000000008 //window is neither mapped onto the desktop nor shown in the taskbar/dock/window list; SDL_ShowWindow() is required for it to become visible
	WindowBorderless        WindowFlags = 0x0000000000000010 //no window decoration
	WindowResizable         WindowFlags = 0x0000000000000020 //window can be resized
	WindowMinimized         WindowFlags = 0x0000000000000040 //window is minimized
	WindowMaximized         WindowFlags = 0x0000000000000080 //window is maximized
	WindowMouseGrabbed      WindowFlags = 0x0000000000000100 //window has grabbed mouse input
	WindowInputFocus        WindowFlags = 0x0000000000000200 //window has input focus
	WindowMouseFocus        WindowFlags = 0x0000000000000400 //window has mouse focus
	WindowExternal          WindowFlags = 0x0000000000000800 //window not created by SDL
	WindowModal             WindowFlags = 0x0000000000001000 //window is modal
	WindowHighPixelDensity  WindowFlags = 0x0000000000002000 //window uses high pixel density back buffer if possible
	WindowMouseCapture      WindowFlags = 0x0000000000004000 //window has mouse captured (unrelated to MOUSE_GRABBED)
	WindowMouseRelativeMode WindowFlags = 0x0000000000008000 //window has relative mode enabled
	WindowAlwaysOnTop       WindowFlags = 0x0000000000010000 //window should always be above others
	WindowUtility           WindowFlags = 0x0000000000020000 //window should be treated as a utility window, not showing in the task bar and window list
	WindowTooltip           WindowFlags = 0x0000000000040000 //window should be treated as a tooltip and does not get mouse or keyboard focus, requires a parent window
	WindowPopupMenu         WindowFlags = 0x0000000000080000 //window should be treated as a popup menu, requires a parent window
	WindowKeyboardGrabbed   WindowFlags = 0x0000000000100000 //window has grabbed keyboard input
	WindowVulkan            WindowFlags = 0x0000000010000000 //window usable for Vulkan surface
	WindowMetal             WindowFlags = 0x0000000020000000 //window usable for Metal view
	WindowTransparent       WindowFlags = 0x0000000040000000 //window with transparent buffer
	WindowNotFocusable      WindowFlags = 0x0000000080000000 //window should not be focusable
)

// A magic value used with [WindowposUndefined].
//
// Generally this macro isn't used directly, but rather through
// [WindowposUndefined] or [WindowposUndefinedDisplay].
//
// This macro is available since SDL 3.2.0.
const WindowposUndefinedMask = 0x1FFF0000

// Used to indicate that you don't care what the window position is.
//
// If you _really_ don't care, [WindowposUndefined] is the same, but always
// uses the primary display instead of specifying one.
//
// x: the [DisplayID] of the display to use.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WINDOWPOS_UNDEFINED_DISPLAY
func WindowposUndefinedDisplay(x DisplayID) int {
	return WindowposUndefinedMask | int(x)
}

// Used to indicate that you don't care what the window position/display is.
//
// This always uses the primary display.
//
// This macro is available since SDL 3.2.0.
const WindowposUndefined = WindowposUndefinedMask

// A macro to test if the window position is marked as "undefined."
//
// x: the window position value.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WINDOWPOS_ISUNDEFINED
func WindowposIsUndefined(x int) bool {
	return x&0xFFFF0000 == WindowposUndefinedMask
}

// A magic value used with [WindowposCentered].
//
// Generally this macro isn't used directly, but rather through
// [WindowposCentered] or [WindowposCenteredDisplay].
//
// This macro is available since SDL 3.2.0.
const WindowposCenteredMask = 0x2FFF0000

// Used to indicate that the window position should be centered.
//
// [WindowposCentered] is the same, but always uses the primary display
// instead of specifying one.
//
// x: the [DisplayID] of the display to use.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WINDOWPOS_CENTERED_DISPLAY
func WindowposCenteredDisplay(x DisplayID) int {
	return WindowposCenteredMask | int(x)
}

// Used to indicate that the window position should be centered.
//
// This always uses the primary display.
//
// This macro is available since SDL 3.2.0.
const WindowposCentered = WindowposCenteredMask

// A macro to test if the window position is marked as "centered."
//
// x: the window position value.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WINDOWPOS_ISCENTERED
func WindowposIsCentered(x int) bool {
	return x&0xFFFF0000 == WindowposCenteredMask
}

// Window flash operation.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_FlashOperation
type FlashOperation uint32

const (
	FlashCancel       FlashOperation = iota // Cancel any window flash state
	FlashBriefly                            // Flash the window briefly to get attention
	FlashUntilFocused                       // Flash the window until it gets focus
)

// An opaque handle to an OpenGL context.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GLContext
type GLContext *C.struct_SDL_GLContextState

// Opaque type for an EGL display.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_EGLDisplay
type EGLDisplay unsafe.Pointer

// Opaque type for an EGL config.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_EGLConfig
type EGLConfig unsafe.Pointer

// Opaque type for an EGL surface.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_EGLSurface
type EGLSurface unsafe.Pointer

// An EGL attribute, used when creating an EGL context.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_EGLAttrib
type EGLAttrib C.SDL_EGLAttrib

// An EGL integer attribute, used when creating an EGL surface.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_EGLint
type EGLint C.SDL_EGLint

// EGL platform attribute initialization callback.
//
// This is called when SDL is attempting to create an EGL context, to let the
// app add extra attributes to its eglGetPlatformDisplay() call.
//
// The callback should return a slice EGL attributes. If this function returns
// nil, the [CreateWindow] process will fail gracefully.
//
// The arrays returned by each callback will be appended to the existing
// attribute arrays defined by SDL.
//
// Returns a slice of attributes.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_EGLAttribArrayCallback
type EGLAttribArrayCallback func() []EGLAttrib

// EGL surface/context attribute initialization callback types.
//
// This is called when SDL is attempting to create an EGL surface, to let the
// app add extra attributes to its eglCreateWindowSurface() or
// eglCreateContext calls.
//
// For convenience, the [EGLDisplay] and [EGLConfig] to use are provided to the
// callback.
//
// The callback should return a slice EGL attributes. If this function returns
// nil, the [CreateWindow] process will fail gracefully.
//
// The arrays returned by each callback will be appended to the existing
// attribute arrays defined by SDL.
//
// userdata: an app-controlled pointer that is passed to the callback.
//
// display: the EGL display to be used.
//
// config: the EGL config to be used.
//
// Returns a slice of attributes.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_EGLIntArrayCallback
type EGLIntArrayCallback func(display EGLDisplay, config EGLConfig) []EGLint

// An enumeration of OpenGL configuration attributes.
//
// While you can set most OpenGL attributes normally, the attributes listed
// above must be known before SDL creates the window that will be used with
// the OpenGL context. These attributes are set and read with
// [GL_SetAttribute] and [GL_GetAttribute].
//
// In some cases, these attributes are minimum requests; the GL does not
// promise to give you exactly what you asked for. It's possible to ask for a
// 16-bit depth buffer and get a 24-bit one instead, for example, or to ask
// for no stencil buffer and still have one available. Context creation should
// fail if the GL can't provide your requested attributes at a minimum, but
// you should check to see exactly what you got.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GLAttr
type GLAttr uint32

const (
	GLRedSize                  GLAttr = iota // the minimum number of bits for the red channel of the color buffer; defaults to 3.
	GLGreenSize                              // the minimum number of bits for the green channel of the color buffer; defaults to 3.
	GLBlueSize                               // the minimum number of bits for the blue channel of the color buffer; defaults to 2.
	GLAlphaSize                              // the minimum number of bits for the alpha channel of the color buffer; defaults to 0.
	GLBufferSize                             // the minimum number of bits for frame buffer size; defaults to 0.
	GLDoublebuffer                           // whether the output is single or double buffered; defaults to double buffering on.
	GLDepthSize                              // the minimum number of bits in the depth buffer; defaults to 16.
	GLStencilSize                            // the minimum number of bits in the stencil buffer; defaults to 0.
	GLAccumRedSize                           // the minimum number of bits for the red channel of the accumulation buffer; defaults to 0.
	GLAccumGreenSize                         // the minimum number of bits for the green channel of the accumulation buffer; defaults to 0.
	GLAccumBlueSize                          // the minimum number of bits for the blue channel of the accumulation buffer; defaults to 0.
	GLAccumAlphaSize                         // the minimum number of bits for the alpha channel of the accumulation buffer; defaults to 0.
	GLStereo                                 // whether the output is stereo 3D; defaults to off.
	GLMultisamplebuffers                     // the number of buffers used for multisample anti-aliasing; defaults to 0.
	GLMultisamplesamples                     // the number of samples used around the current pixel used for multisample anti-aliasing.
	GLAcceleratedVisual                      // set to 1 to require hardware acceleration, set to 0 to force software rendering; defaults to allow either.
	GLRetainedBacking                        // not used (deprecated).
	GLContextMajorVersion                    // OpenGL context major version.
	GLContextMinorVersion                    // OpenGL context minor version.
	GLContextFlags                           // some combination of 0 or more of elements of the [GLContextFlag] enumeration; defaults to 0.
	GLContextProfileMask                     // type of GL context (Core, Compatibility, ES). See [GLProfile]; default value depends on platform.
	GLShareWithCurrentContext                // OpenGL context sharing; defaults to 0.
	GLFramebufferSrgbCapable                 // requests sRGB capable visual; defaults to 0.
	GLContextReleaseBehavior                 // sets context the release behavior. See [GLContextReleaseFlag]; defaults to FLUSH.
	GLContextResetNotification               // set context reset notification. See [GLContextResetNotificationType]; defaults to NO_NOTIFICATION.
	GLContextNoError
	GLFloatbuffers
	GLEglPlatform
)

// Possible values to be set for the [GLContextProfileMask] attribute.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GLProfile
type GLProfile uint32

const (
	GLContextProfileCore          GLProfile = 0x0001 //OpenGL Core Profile context
	GLContextProfileCompatibility GLProfile = 0x0002 //OpenGL Compatibility Profile context
	GLContextProfileEs            GLProfile = 0x0004 //GLX_CONTEXT_ES2_PROFILE_BIT_EXT
)

// Possible flags to be set for the [GLContextFlags] attribute.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GLContextFlag
type GLContextFlag uint32

const (
	GLContextDebugFlag             GLContextFlag = 0x0001
	GLContextForwardCompatibleFlag GLContextFlag = 0x0002
	GLContextRobustAccessFlag      GLContextFlag = 0x0004
	GLContextResetIsolationFlag    GLContextFlag = 0x0008
)

// Possible values to be set for the [GLContextReleaseBehavior]
// attribute.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GLContextReleaseFlag
type GLContextReleaseFlag uint32

const (
	GLContextReleaseBehaviorNone  GLContextReleaseFlag = 0x0000
	GLContextReleaseBehaviorFlush GLContextReleaseFlag = 0x0001
)

// Possible values to be set [GLContextResetNotification] attribute.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GLContextResetNotification
type GLContextResetNotificationType uint32

const (
	GLContextResetNoNotification GLContextResetNotificationType = 0x0000
	GLContextResetLoseContext    GLContextResetNotificationType = 0x0001
)

// Get the number of video drivers compiled into SDL.
//
// Returns the number of built in video drivers.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetNumVideoDrivers
func GetNumVideoDrivers() int {
	return (int)(C.SDL_GetNumVideoDrivers())
}

// Get the name of a built in video driver.
//
// The video drivers are presented in the order in which they are normally
// checked during initialization.
//
// The names of drivers are all simple, low-ASCII identifiers, like "cocoa",
// "x11" or "windows". These never have Unicode characters, and are not meant
// to be proper names.
//
// index: the index of a video driver.
//
// Returns the name of the video driver with the given **index**.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetVideoDriver
func GetVideoDriver(index int) string {
	return C.GoString(C.SDL_GetVideoDriver((C.int)(index)))
}

// Get the name of the currently initialized video driver.
//
// The names of drivers are all simple, low-ASCII identifiers, like "cocoa",
// "x11" or "windows". These never have Unicode characters, and are not meant
// to be proper names.
//
// Returns the name of the current video driver or "" if no driver has been
// initialized.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetCurrentVideoDriver
func GetCurrentVideoDriver() string {
	return C.GoString(C.SDL_GetCurrentVideoDriver())
}

// Get the current system theme.
//
// Returns the current system theme, light, dark, or unknown.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSystemTheme
func GetSystemTheme() SystemTheme {
	return (SystemTheme)(C.SDL_GetSystemTheme())
}

// Get a list of currently connected displays.
//
// Returns a slice of display instance IDs or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetDisplays
func GetDisplays() ([]DisplayID, error) {
	var count C.int
	displays := C.SDL_GetDisplays(&count)
	if displays == nil {
		return nil, getError()
	}
	result := make([]DisplayID, *displays)
	for i, d := range unsafe.Slice(displays, count) {
		result[i] = DisplayID(d)
	}
	C.SDL_free(unsafe.Pointer(displays))
	return result, nil
}

// Return the primary display.
//
// Returns the instance ID of the primary display or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetPrimaryDisplay
func GetPrimaryDisplay() (DisplayID, error) {
	display := C.SDL_GetPrimaryDisplay()
	if display == 0 {
		return 0, getError()
	}
	return DisplayID(display), nil
}

// Get the properties associated with a display.
//
// The following read-only properties are provided by SDL:
//
//   - [PropDisplayHdrEnabledBoolean]: true if the display has HDR
//     headroom above the SDR white point. This is for informational and
//     diagnostic purposes only, as not all platforms provide this information
//     at the display level.
//
// On KMS/DRM:
//
//   - [PropDisplayKmsdrmPanelOrientationNumber]: the "panel
//     orientation" property for the display in degrees of clockwise rotation.
//     Note that this is provided only as a hint, and the application is
//     responsible for any coordinate transformations needed to conform to the
//     requested display orientation.
//
// displayID: the instance ID of the display to query.
//
// Returns a valid property ID or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetDisplayProperties
func (displayID DisplayID) Properties() (PropertiesID, error) {
	props := C.SDL_GetDisplayProperties((C.SDL_DisplayID)(displayID))
	if props == 0 {
		return 0, getError()
	}
	return PropertiesID(props), nil
}

const PropDisplayHdrEnabledBoolean = "SDL.display.HDR_enabled"
const PropDisplayKmsdrmPanelOrientationNumber = "SDL.display.KMSDRM.panel_orientation"

// Get the name of a display.
//
// displayID: the instance ID of the display to query.
//
// Returns the name of a display or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetDisplayName
func (displayID DisplayID) Name() (string, error) {
	name := C.SDL_GetDisplayName((C.SDL_DisplayID)(displayID))
	if name == nil {
		return "", getError()
	}
	return C.GoString(name), nil
}

// Get the desktop area represented by a display.
//
// The primary display is often located at (0,0), but may be placed at a
// different location depending on monitor layout.
//
// displayID: the instance ID of the display to query.
//
// Returns the display bounds or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetDisplayBounds
func (displayID DisplayID) Bounds() (Rect, error) {
	var rect C.SDL_Rect
	if !C.SDL_GetDisplayBounds((C.SDL_DisplayID)(displayID), &rect) {
		return Rect{}, getError()
	}
	return Rect{int(rect.x), int(rect.y), int(rect.w), int(rect.h)}, nil
}

// Get the usable desktop area represented by a display, in screen
// coordinates.
//
// This is the same area as [DisplayID.Bounds] reports, but with portions
// reserved by the system removed. For example, on Apple's macOS, this
// subtracts the area occupied by the menu bar and dock.
//
// Setting a window to be fullscreen generally bypasses these unusable areas,
// so these are good guidelines for the maximum space available to a
// non-fullscreen window.
//
// displayID: the instance ID of the display to query.
//
// Returns the display bounds or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetDisplayUsableBounds
func (displayID DisplayID) UsableBounds() (Rect, error) {
	var rect C.SDL_Rect
	if !C.SDL_GetDisplayUsableBounds((C.SDL_DisplayID)(displayID), &rect) {
		return Rect{}, getError()
	}
	return Rect{int(rect.x), int(rect.y), int(rect.w), int(rect.h)}, nil
}

// Get the orientation of a display when it is unrotated.
//
// displayID: the instance ID of the display to query.
//
// Returns the [DisplayOrientation] enum value of the display, or
// [OrientationUnknown] if it isn't available.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetNaturalDisplayOrientation
func (displayID DisplayID) NaturalOrientation() DisplayOrientation {
	return (DisplayOrientation)(C.SDL_GetNaturalDisplayOrientation((C.SDL_DisplayID)(displayID)))
}

// Get the orientation of a display.
//
// displayID: the instance ID of the display to query.
//
// Returns the [DisplayOrientation] enum value of the display, or
// [OrientationUnknown] if it isn't available.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetCurrentDisplayOrientation
func (displayID DisplayID) CurrentOrientation() DisplayOrientation {
	return (DisplayOrientation)(C.SDL_GetCurrentDisplayOrientation((C.SDL_DisplayID)(displayID)))
}

// Get the content scale of a display.
//
// The content scale is the expected scale for content based on the DPI
// settings of the display. For example, a 4K display might have a 2.0 (200%)
// display scale, which means that the user expects UI elements to be twice as
// big on this display, to aid in readability.
//
// After window creation, [Window.DisplayScale] should be used to query
// the content scale factor for individual windows instead of querying the
// display for a window and calling this function, as the per-window content
// scale factor may differ from the base value of the display it is on,
// particularly on high-DPI and/or multi-monitor desktop configurations.
//
// displayID: the instance ID of the display to query.
//
// Returns the content scale of the display or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetDisplayContentScale
func (displayID DisplayID) ContentScale() (float32, error) {
	scale := C.SDL_GetDisplayContentScale((C.SDL_DisplayID)(displayID))
	if scale == 0 {
		return 0, getError()
	}
	return float32(scale), nil
}

// Get a list of fullscreen display modes available on a display.
//
// The display modes are sorted in this priority:
//
//   - w -> largest to smallest
//   - h -> largest to smallest
//   - bits per pixel -> more colors to fewer colors
//   - packed pixel layout -> largest to smallest
//   - refresh rate -> highest to lowest
//   - pixel density -> lowest to highest
//
// displayID: the instance ID of the display to query.
//
// Returns a slice of display mode pointers or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetFullscreenDisplayModes
func (displayID DisplayID) FullscreenDisplayModes() ([]*DisplayMode, error) {
	var count C.int
	modes := C.SDL_GetFullscreenDisplayModes((C.SDL_DisplayID)(displayID), &count)
	if modes == nil {
		return nil, getError()
	}
	result := make([]*DisplayMode, count)
	for i, m := range unsafe.Slice(modes, count) {
		result[i] = &DisplayMode{
			DisplayID(m.displayID),
			PixelFormat(m.format),
			int(m.w),
			int(m.h),
			float32(m.pixel_density),
			float32(m.refresh_rate),
			int(m.refresh_rate_numerator),
			int(m.refresh_rate_denominator),
			m.internal,
		}
	}
	C.SDL_free(unsafe.Pointer(modes))
	return result, nil
}

// Get the closest match to the requested display mode.
//
// The available display modes are scanned and the
// closest mode matching the requested mode is returned. The mode format and
// refresh rate default to the desktop mode if they are set to 0. The modes
// are scanned with size being first priority, format being second priority,
// and finally checking the refresh rate. If all the available modes are too
// small, then false is returned.
//
// displayID: the instance ID of the display to query.
//
// w: the width in pixels of the desired display mode.
//
// h: the height in pixels of the desired display mode.
//
// refreshRate: the refresh rate of the desired display mode, or 0.0f
// for the desktop refresh rate.
//
// includeHighDensityModes: boolean to include high density modes in
// the search.
//
// Returns the closest display mode equal to or larger than the desired
// mode or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetClosestFullscreenDisplayMode
func (displayID DisplayID) ClosestFullscreenDisplayMode(w int, h int, refreshRate float32, includeHighDensityModes bool) (*DisplayMode, error) {
	var closest C.SDL_DisplayMode
	ok := C.SDL_GetClosestFullscreenDisplayMode((C.SDL_DisplayID)(displayID), (C.int)(w), (C.int)(h), (C.float)(refreshRate), (C.bool)(includeHighDensityModes), &closest)
	if !ok {
		return nil, getError()
	}
	return &DisplayMode{
		DisplayID(closest.displayID),
		PixelFormat(closest.format),
		int(closest.w),
		int(closest.h),
		float32(closest.pixel_density),
		float32(closest.refresh_rate),
		int(closest.refresh_rate_numerator),
		int(closest.refresh_rate_denominator),
		closest.internal,
	}, nil
}

// Get information about the desktop's display mode.
//
// There's a difference between this function and [DisplayID.CurrentDisplayMode]
// when SDL runs fullscreen and has changed the resolution. In that case this
// function will return the previous native display mode, and not the current
// display mode.
//
// displayID: the instance ID of the display to query.
//
// Returns a pointer to the desktop display mode or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetDesktopDisplayMode
func (displayID DisplayID) DesktopDisplayMode() (*DisplayMode, error) {
	mode := C.SDL_GetDesktopDisplayMode((C.SDL_DisplayID)(displayID))
	if mode == nil {
		return nil, getError()
	}
	return &DisplayMode{
		DisplayID(mode.displayID),
		PixelFormat(mode.format),
		int(mode.w),
		int(mode.h),
		float32(mode.pixel_density),
		float32(mode.refresh_rate),
		int(mode.refresh_rate_numerator),
		int(mode.refresh_rate_denominator),
		mode.internal,
	}, nil
}

// Get information about the current display mode.
//
// There's a difference between this function and [DisplayID.DesktopDisplayMode]
// when SDL runs fullscreen and has changed the resolution. In that case this
// function will return the current display mode, and not the previous native
// display mode.
//
// displayID: the instance ID of the display to query.
//
// Returns a pointer to the desktop display mode or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetCurrentDisplayMode
func (displayID DisplayID) CurrentDisplayMode() (*DisplayMode, error) {
	mode := C.SDL_GetCurrentDisplayMode((C.SDL_DisplayID)(displayID))
	if mode == nil {
		return nil, getError()
	}
	return &DisplayMode{
		DisplayID(mode.displayID),
		PixelFormat(mode.format),
		int(mode.w),
		int(mode.h),
		float32(mode.pixel_density),
		float32(mode.refresh_rate),
		int(mode.refresh_rate_numerator),
		int(mode.refresh_rate_denominator),
		mode.internal,
	}, nil
}

// Get the display containing a point.
//
// point: the point to query.
//
// Returns the instance ID of the display containing the point or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetDisplayForPoint
func GetDisplayForPoint(point Point) (DisplayID, error) {
	display := C.SDL_GetDisplayForPoint(&C.SDL_Point{C.int(point.X), C.int(point.Y)})
	if display == 0 {
		return 0, getError()
	}
	return DisplayID(display), nil
}

// Get the display primarily containing a rect.
//
// rect: the rect to query.
//
// Returns the instance ID of the display entirely containing the rect or
// closest to the center of the rect on success an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetDisplayForRect
func GetDisplayForRect(rect Rect) (DisplayID, error) {
	display := C.SDL_GetDisplayForRect(&C.SDL_Rect{C.int(rect.X), C.int(rect.Y), C.int(rect.W), C.int(rect.H)})
	if display == 0 {
		return 0, getError()
	}
	return DisplayID(display), nil
}

// Get the display associated with a window.
//
// window: the window to query.
//
// Returns the instance ID of the display containing the center of the window
// on success or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetDisplayForWindow
func GetDisplayForWindow(window *Window) (DisplayID, error) {
	display := C.SDL_GetDisplayForWindow((*C.SDL_Window)(window))
	if display == 0 {
		return 0, getError()
	}
	return DisplayID(display), nil
}

// Get the pixel density of a window.
//
// This is a ratio of pixel size to window size. For example, if the window is
// 1920x1080 and it has a high density back buffer of 3840x2160 pixels, it
// would have a pixel density of 2.0.
//
// window: the window to query.
//
// Returns the pixel density or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowPixelDensity
func (window *Window) PixelDensity() (float32, error) {
	pd := C.SDL_GetWindowPixelDensity((*C.SDL_Window)(window))
	if pd == 0 {
		return 0, getError()
	}
	return float32(pd), nil
}

// Get the content display scale relative to a window's pixel size.
//
// This is a combination of the window pixel density and the display content
// scale, and is the expected scale for displaying content in this window. For
// example, if a 3840x2160 window had a display scale of 2.0, the user expects
// the content to take twice as many pixels and be the same physical size as
// if it were being displayed in a 1920x1080 window with a display scale of
// 1.0.
//
// Conceptually this value corresponds to the scale display setting, and is
// updated when that setting is changed, or the window moves to a display with
// a different scale setting.
//
// window: the window to query.
//
// Returns the display scale, or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowDisplayScale
func (window *Window) DisplayScale() (float32, error) {
	scale := C.SDL_GetWindowDisplayScale((*C.SDL_Window)(window))
	if scale == 0 {
		return 0, getError()
	}
	return float32(scale), nil
}

// Set the display mode to use when a window is visible and fullscreen.
//
// This only affects the display mode used when the window is fullscreen. To
// change the window size when the window is not fullscreen, use
// [Window.SetSize]().
//
// If the window is currently in the fullscreen state, this request is
// asynchronous on some windowing systems and the new mode dimensions may not
// be applied immediately upon the return of this function. If an immediate
// change is required, call [Window.Sync] to block until the changes have
// taken effect.
//
// When the new mode takes effect, an [EventWindowResized] and/or an
// [EventWindowPixelSizeChanged] event will be emitted with the new mode
// dimensions.
//
// window: the window to affect.
//
// mode: a pointer to the display mode to use, which can be nil for
// borderless fullscreen desktop mode, or one of the fullscreen
// modes returned by [DisplayID.FullscreenDisplayModes] to set an
// exclusive fullscreen mode.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SetFullscreenMode
func (window *Window) SetFullscreenMode(mode *DisplayMode) error {
	var cmode *C.SDL_DisplayMode
	if mode != nil {
		cmode = &C.SDL_DisplayMode{
			C.SDL_DisplayID(mode.DisplayID),
			C.SDL_PixelFormat(mode.Format),
			C.int(mode.W),
			C.int(mode.H),
			C.float(mode.PixelDensity),
			C.float(mode.RefreshRate),
			C.int(mode.RefreshRateNumerator),
			C.int(mode.RefreshRateDenominator),
			mode.internal,
		}
	}
	if !C.SDL_SetWindowFullscreenMode((*C.SDL_Window)(window), cmode) {
		return getError()
	}
	return nil
}

// Query the display mode to use when a window is visible at fullscreen.
//
// window: the window to query.
//
// Returns a pointer to the exclusive fullscreen mode to use or nil for
// borderless fullscreen desktop mode.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowFullscreenMode
func (window *Window) FullscreenMode() *DisplayMode {
	mode := C.SDL_GetWindowFullscreenMode((*C.SDL_Window)(window))
	if mode == nil {
		return nil
	}
	return &DisplayMode{
		DisplayID(mode.displayID),
		PixelFormat(mode.format),
		int(mode.w),
		int(mode.h),
		float32(mode.pixel_density),
		float32(mode.refresh_rate),
		int(mode.refresh_rate_numerator),
		int(mode.refresh_rate_denominator),
		mode.internal,
	}
}

// Get the raw ICC profile data for the screen the window is currently on.
//
// window: the window to query.
//
// Returns the raw ICC profile data on success or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowICCProfile
func (window *Window) ICCProfile() ([]byte, error) {
	var size C.size_t
	data := C.SDL_GetWindowICCProfile((*C.SDL_Window)(window), &size)
	if data == nil {
		return nil, getError()
	}
	result := C.GoBytes(data, C.int(size))
	C.SDL_free(data)
	return result, nil
}

// Get the pixel format associated with the window.
//
// window: the window to query.
//
// Returns the pixel format of the window on success or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowPixelFormat
func (window *Window) PixelFormat() (PixelFormat, error) {
	pf := (PixelFormat)(C.SDL_GetWindowPixelFormat((*C.SDL_Window)(window)))
	if pf == PixelformatUnknown {
		return PixelformatUnknown, getError()
	}
	return pf, nil
}

// Get a list of valid windows.
//
// Returns a alice of [Window] pointers or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindows
func GetWindows() ([]*Window, error) {
	var count C.int
	windows := C.SDL_GetWindows(&count)
	if windows == nil {
		return nil, getError()
	}
	result := make([]*Window, count)
	for i, w := range unsafe.Slice(windows, count) {
		result[i] = (*Window)(w)
	}
	C.SDL_free(unsafe.Pointer(windows))
	return result, nil
}

// Create a window with the specified dimensions and flags.
//
// flags may be any of the following OR'd together:
//
//   - [WindowFullscreen]: fullscreen window at desktop resolution
//   - [WindowOpenGL]: window usable with an OpenGL context
//   - [WindowOccluded]: window partially or completely obscured by another
//     window
//   - [WindowHidden]: window is not visible
//   - [WindowBorderless]: no window decoration
//   - [WindowResizable]: window can be resized
//   - [WindowMinimized]: window is minimized
//   - [WindowMaximized]: window is maximized
//   - [WindowMouseGrabbed]: window has grabbed mouse focus
//   - [WindowInputFocus]: window has input focus
//   - [WindowMouseFocus]: window has mouse focus
//   - [WindowExternal]: window not created by SDL
//   - [WindowModal]: window is modal
//   - [WindowHighPixelDensity]: window uses high pixel density back
//     buffer if possible
//   - [WindowMouseCapture]: window has mouse captured (unrelated to
//     MOUSE_GRABBED)
//   - [WindowAlwaysOnTop]: window should always be above others
//   - [WindowUtility]: window should be treated as a utility window, not
//     showing in the task bar and window list
//   - [WindowTooltip]: window should be treated as a tooltip and does not
//     get mouse or keyboard focus, requires a parent window
//   - [WindowPopupMenu]: window should be treated as a popup menu,
//     requires a parent window
//   - [WindowKeyboardGrabbed]: window has grabbed keyboard input
//   - [WindowVulkan]: window usable with a Vulkan instance
//   - [WindowMetal]: window usable with a Metal instance
//   - [WindowTransparent]: window with transparent buffer
//   - [WindowNotFocusable]: window should not be focusable
//
// The [Window] is implicitly shown if [WindowHidden] is not set.
//
// On Apple's macOS, you **must** set the NSHighResolutionCapable Info.plist
// property to YES, otherwise you will not receive a High-DPI OpenGL canvas.
//
// The window pixel size may differ from its window coordinate size if the
// window is on a high pixel density display. Use [Window.Size] to query
// the client area's size in window coordinates, and
// [Window.SizeInPixels] or [Renderer.OutputSize] to query the
// drawable size in pixels. Note that the drawable size can vary after the
// window is created and should be queried again if you get an
// [EventWindowPixelSizeChanged] event.
//
// If the window is created with any of the [WindowOpenGL] or
// [WindowVulkan] flags, then the corresponding LoadLibrary function
// ([GL_LoadLibrary] or [Vulkan_LoadLibrary]) is called and the
// corresponding UnloadLibrary function is called by [Window.Destroy].
//
// If [WindowVulkan] is specified and there isn't a working Vulkan driver,
// [CreateWindow] will fail, because [Vulkan_LoadLibrary] will fail.
//
// If [WindowMetal] is specified on an OS that does not support Metal,
// [CreateWindow] will fail.
//
// If you intend to use this window with a [Renderer], you should use
// [CreateWindowAndRenderer] instead of this function, to avoid window
// flicker.
//
// On non-Apple devices, SDL requires you to either not link to the Vulkan
// loader or link to a dynamic library version. This limitation may be removed
// in a future version of SDL.
//
// title: the title of the window, in UTF-8 encoding.
//
// w: the width of the window.
//
// h: the height of the window.
//
// flags: 0, or one or more [WindowFlags] OR'd together.
//
// Returns the window that was created or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateWindow
func CreateWindow(title string, w int, h int, flags WindowFlags) (*Window, error) {
	window := (*Window)(C.SDL_CreateWindow(tmpstring(title), (C.int)(w), (C.int)(h), (C.SDL_WindowFlags)(flags)))
	if window == nil {
		return nil, getError()
	}
	return window, nil
}

// Create a child popup window of the specified parent window.
//
// The flags parameter **must** contain at least one of the following:
//
//   - [WindowTooltip]: The popup window is a tooltip and will not pass any
//     input events.
//   - [WindowPopupMenu]: The popup window is a popup menu. The topmost
//     popup menu will implicitly gain the keyboard focus.
//
// The following flags are not relevant to popup window creation and will be
// ignored:
//
//   - [WindowMinimized]
//   - [WindowMaximized]
//   - [WindowFullscreen]
//   - [WindowBorderless]
//
// The following flags are incompatible with popup window creation and will
// cause it to fail:
//
//   - [WindowUtility]
//   - [WindowModal]
//
// The parent parameter **must** be non-null and a valid window. The parent of
// a popup window can be either a regular, toplevel window, or another popup
// window.
//
// Popup windows cannot be minimized, maximized, made fullscreen, raised,
// flash, be made a modal window, be the parent of a toplevel window, or grab
// the mouse and/or keyboard. Attempts to do so will fail.
//
// Popup windows implicitly do not have a border/decorations and do not appear
// on the taskbar/dock or in lists of windows such as alt-tab menus.
//
// If a parent window is hidden or destroyed, any child popup windows will be
// recursively hidden or destroyed as well. Child popup windows not explicitly
// hidden will be restored when the parent is shown.
//
// parent: the parent of the window, must not be nil.
//
// offset_x: the x position of the popup window relative to the origin
// of the parent.
//
// offset_y: the y position of the popup window relative to the origin
// of the parent window.
//
// w: the width of the window.
//
// h: the height of the window.
//
// flags: [WindowTooltip] or [WindowPopupMenu], and zero or more
// additional [WindowFlags] OR'd together.
//
// Returns the window that was created or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreatePopupWindow
func CreatePopupWindow(parent *Window, offset_x int, offset_y int, w int, h int, flags WindowFlags) (*Window, error) {
	window := (*Window)(C.SDL_CreatePopupWindow((*C.SDL_Window)(parent), (C.int)(offset_x), (C.int)(offset_y), (C.int)(w), (C.int)(h), (C.SDL_WindowFlags)(flags)))
	if window == nil {
		return nil, getError()
	}
	return window, nil
}

// Create a window with the specified properties.
//
// These are the supported properties:
//
//   - [PropWindowCreateAlwaysOnTopBoolean]: true if the window should
//     be always on top
//   - [PropWindowCreateBorderlessBoolean]: true if the window has no
//     window decoration
//   - [PropWindowCreateExternalGraphicsContextBoolean]: true if the
//     window will be used with an externally managed graphics context.
//   - [PropWindowCreateFocusableBoolean]: true if the window should
//     accept keyboard input (defaults true)
//   - [PropWindowCreateFullscreenBoolean]: true if the window should
//     start in fullscreen mode at desktop resolution
//   - [PropWindowCreateHeightNumber]: the height of the window
//   - [PropWindowCreateHiddenBoolean]: true if the window should start
//     hidden
//   - [PropWindowCreateHighPixelDensityBoolean]: true if the window
//     uses a high pixel density buffer if possible
//   - [PropWindowCreateMaximizedBoolean]: true if the window should
//     start maximized
//   - [PropWindowCreateMenuBoolean]: true if the window is a popup menu
//   - [PropWindowCreateMetalBoolean]: true if the window will be used
//     with Metal rendering
//   - [PropWindowCreateMinimizedBoolean]: true if the window should
//     start minimized
//   - [PropWindowCreateModalBoolean]: true if the window is modal to
//     its parent
//   - [PropWindowCreateMouseGrabbedBoolean]: true if the window starts
//     with grabbed mouse focus
//   - [PropWindowCreateOpenGLBoolean]: true if the window will be used
//     with OpenGL rendering
//   - [PropWindowCreateParentPointer]: a [Window] that will be the
//     parent of this window, required for windows with the "tooltip", "menu",
//     and "modal" properties
//   - [PropWindowCreateResizableBoolean]: true if the window should be
//     resizable
//   - [PropWindowCreateTitleString]: the title of the window, in UTF-8
//     encoding
//   - [PropWindowCreateTransparentBoolean]: true if the window show
//     transparent in the areas with alpha of 0
//   - [PropWindowCreateTooltipBoolean]: true if the window is a tooltip
//   - [PropWindowCreateUtilityBoolean]: true if the window is a utility
//     window, not showing in the task bar and window list
//   - [PropWindowCreateVulkanBoolean]: true if the window will be used
//     with Vulkan rendering
//   - [PropWindowCreateWidthNumber]: the width of the window
//   - [PropWindowCreateXNumber]: the x position of the window, or
//     [WindowposCentered], defaults to [WindowposUndefined]. This is
//     relative to the parent for windows with the "tooltip" or "menu" property
//     set.
//   - [PropWindowCreateYNumber]: the y position of the window, or
//     [WindowposCentered], defaults to [WindowposUndefined]. This is
//     relative to the parent for windows with the "tooltip" or "menu" property
//     set.
//
// These are additional supported properties on macOS:
//
//   - [PropWindowCreateCocoaWindowPointer]: the
//     (__unsafe_unretained) NSWindow associated with the window, if you want
//     to wrap an existing window.
//   - [PropWindowCreateCocoaViewPointer]: the (__unsafe_unretained)
//     NSView associated with the window, defaults to [window contentView]
//
// These are additional supported properties on Wayland:
//
//   - [PropWindowCreateWaylandSurfaceRoleCustomBoolean] - true if
//     the application wants to use the Wayland surface for a custom role and
//     does not want it attached to an XDG toplevel window. See
//     [README/wayland](README/wayland) for more information on using custom
//     surfaces.
//   - [PropWindowCreateWaylandCreateEglWindowBoolean] - true if the
//     application wants an associated wl_egl_window object to be created and
//     attached to the window, even if the window does not have the OpenGL
//     property or [WindowOpenGL] flag set.
//   - [PropWindowCreateWaylandWlSurfacePointer] - the wl_surface
//     associated with the window, if you want to wrap an existing window. See
//     [README/wayland](README/wayland) for more information.
//
// These are additional supported properties on Windows:
//
//   - [PropWindowCreateWIN32HwndPointer]: the HWND associated with the
//     window, if you want to wrap an existing window.
//   - [PropWindowCreateWIN32PixelFormatHwndPointer]: optional,
//     another window to share pixel format with, useful for OpenGL windows
//
// These are additional supported properties with X11:
//
//   - [PropWindowCreateX11WindowNumber]: the X11 Window associated
//     with the window, if you want to wrap an existing window.
//
// The window is implicitly shown if the "hidden" property is not set.
//
// Windows with the "tooltip" and "menu" properties are popup windows and have
// the behaviors and guidelines outlined in [CreatePopupWindow].
//
// If this window is being created to be used with an [Renderer], you should
// not add a graphics API specific property
// ([PropWindowCreateOpenGLBoolean], etc), as SDL will handle that
// internally when it chooses a renderer. However, SDL might need to recreate
// your window at that point, which may cause the window to appear briefly,
// and then flicker as it is recreated. The correct approach to this is to
// create the window with the [PropWindowCreateHiddenBoolean] property
// set to true, then create the renderer, then show the window with
// [Window.Show].
//
// props: the properties to use.
//
// Returns the window that was created or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateWindowWithProperties
func CreateWindowWithProperties(props PropertiesID) (*Window, error) {
	window := (*Window)(C.SDL_CreateWindowWithProperties((C.SDL_PropertiesID)(props)))
	if window == nil {
		return nil, getError()
	}
	return window, nil
}

const PropWindowCreateAlwaysOnTopBoolean = "SDL.window.create.always_on_top"
const PropWindowCreateBorderlessBoolean = "SDL.window.create.borderless"
const PropWindowCreateFocusableBoolean = "SDL.window.create.focusable"
const PropWindowCreateExternalGraphicsContextBoolean = "SDL.window.create.external_graphics_context"
const PropWindowCreateFlagsNumber = "SDL.window.create.flags"
const PropWindowCreateFullscreenBoolean = "SDL.window.create.fullscreen"
const PropWindowCreateHeightNumber = "SDL.window.create.height"
const PropWindowCreateHiddenBoolean = "SDL.window.create.hidden"
const PropWindowCreateHighPixelDensityBoolean = "SDL.window.create.high_pixel_density"
const PropWindowCreateMaximizedBoolean = "SDL.window.create.maximized"
const PropWindowCreateMenuBoolean = "SDL.window.create.menu"
const PropWindowCreateMetalBoolean = "SDL.window.create.metal"
const PropWindowCreateMinimizedBoolean = "SDL.window.create.minimized"
const PropWindowCreateModalBoolean = "SDL.window.create.modal"
const PropWindowCreateMouseGrabbedBoolean = "SDL.window.create.mouse_grabbed"
const PropWindowCreateOpenGLBoolean = "SDL.window.create.opengl"
const PropWindowCreateParentPointer = "SDL.window.create.parent"
const PropWindowCreateResizableBoolean = "SDL.window.create.resizable"
const PropWindowCreateTitleString = "SDL.window.create.title"
const PropWindowCreateTransparentBoolean = "SDL.window.create.transparent"
const PropWindowCreateTooltipBoolean = "SDL.window.create.tooltip"
const PropWindowCreateUtilityBoolean = "SDL.window.create.utility"
const PropWindowCreateVulkanBoolean = "SDL.window.create.vulkan"
const PropWindowCreateWidthNumber = "SDL.window.create.width"
const PropWindowCreateXNumber = "SDL.window.create.x"
const PropWindowCreateYNumber = "SDL.window.create.y"
const PropWindowCreateCocoaWindowPointer = "SDL.window.create.cocoa.window"
const PropWindowCreateCocoaViewPointer = "SDL.window.create.cocoa.view"
const PropWindowCreateWaylandSurfaceRoleCustomBoolean = "SDL.window.create.wayland.surface_role_custom"
const PropWindowCreateWaylandCreateEglWindowBoolean = "SDL.window.create.wayland.create_egl_window"
const PropWindowCreateWaylandWlSurfacePointer = "SDL.window.create.wayland.wl_surface"
const PropWindowCreateWIN32HwndPointer = "SDL.window.create.win32.hwnd"
const PropWindowCreateWIN32PixelFormatHwndPointer = "SDL.window.create.win32.pixel_format_hwnd"
const PropWindowCreateX11WindowNumber = "SDL.window.create.x11.window"

// Get the numeric ID of a window.
//
// The numeric ID is what [WindowEvent] references, and is necessary to map
// these events to specific [Window] objects.
//
// window: the window to query.
//
// Returns the ID of the window on success or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowID
func (window *Window) ID() (WindowID, error) {
	id := (WindowID)(C.SDL_GetWindowID((*C.SDL_Window)(window)))
	if id == 0 {
		return 0, getError()
	}
	return id, nil
}

// Get a window from a stored ID.
//
// The numeric ID is what [WindowEvent] references, and is necessary to map
// these events to specific [Window] objects.
//
// id: the ID of the window.
//
// Returns the window associated with id or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowFromID
func GetWindowFromID(id WindowID) (*Window, error) {
	window := (*Window)(C.SDL_GetWindowFromID((C.SDL_WindowID)(id)))
	if window == nil {
		return nil, getError()
	}
	return window, nil
}

// Get parent of a window.
//
// window: the window to query.
//
// Returns the parent of the window on success or nil if the window has no
// parent.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowParent
func (window *Window) Parent() *Window {
	return (*Window)(C.SDL_GetWindowParent((*C.SDL_Window)(window)))
}

// Get the properties associated with a window.
//
// The following read-only properties are provided by SDL:
//
//   - [PropWindowShapePointer]: the surface associated with a shaped
//     window
//   - [PropWindowHDREnabledBoolean]: true if the window has HDR
//     headroom above the SDR white point. This property can change dynamically
//     when [EventWindowHDRStateChanged] is sent.
//   - [PropWindowSDRWhiteLevelFloat]: the value of SDR white in the
//     [ColorspaceSRGBLinear] colorspace. On Windows this corresponds to the
//     SDR white level in scRGB colorspace, and on Apple platforms this is
//     always 1.0 for EDR content. This property can change dynamically when
//     [EventWindowHDRStateChanged] is sent.
//   - [PropWindowHDRHeadroomFloat]: the additional high dynamic range
//     that can be displayed, in terms of the SDR white point. When HDR is not
//     enabled, this will be 1.0. This property can change dynamically when
//     [EventWindowHDRStateChanged] is sent.
//
// On Android:
//
//   - [PropWindowAndroidWindowPointer]: the ANativeWindow associated
//     with the window
//   - [PropWindowAndroidSurfacePointer]: the [EGLSurface] associated with
//     the window
//
// On iOS:
//
//   - [PropWindowUikitWindowPointer]: the (__unsafe_unretained)
//     UIWindow associated with the window
//   - [PropWindowUikitMetalViewTagNumber]: the NSInteger tag
//     associated with metal views on the window
//   - [PropWindowUikitOpenglFramebufferNumber]: the OpenGL view's
//     framebuffer object. It must be bound when rendering to the screen using
//     OpenGL.
//   - [PropWindowUikitOpenglRenderbufferNumber]: the OpenGL view's
//     renderbuffer object. It must be bound when [GL_SwapWindow] is called.
//   - [PropWindowUikitOpenglResolveFramebufferNumber]: the OpenGL
//     view's resolve framebuffer, when MSAA is used.
//
// On KMS/DRM:
//
//   - [PropWindowKMSDRMDeviceIndexNumber]: the device index associated
//     with the window (e.g. the X in /dev/dri/cardX)
//   - [PropWindowKMSDRMDRMFDNumber]: the DRM FD associated with the
//     window
//   - [PropWindowKMSDRMGBMDevicePointer]: the GBM device associated
//     with the window
//
// On macOS:
//
//   - [PropWindowCocoaWindowPointer]: the (__unsafe_unretained)
//     NSWindow associated with the window
//   - [PropWindowCocoaMetalViewTagNumber]: the NSInteger tag
//     assocated with metal views on the window
//
// On OpenVR:
//
//   - [PropWindowOpenVROverlayID]: the OpenVR Overlay Handle ID for the
//     associated overlay window.
//
// On Vivante:
//
//   - [PropWindowVivanteDisplayPointer]: the EGLNativeDisplayType
//     associated with the window
//   - [PropWindowVivanteWindowPointer]: the EGLNativeWindowType
//     associated with the window
//   - [PropWindowVivanteSurfacePointer]: the EGLSurface associated with
//     the window
//
// On Windows:
//
//   - [PropWindowWin32HWNDPointer]: the HWND associated with the window
//   - [PropWindowWin32HDCPointer]: the HDC associated with the window
//   - [PropWindowWin32InstancePointer]: the HINSTANCE associated with
//     the window
//
// On Wayland:
//
// Note: The xdg_* window objects do not internally persist across window
// show/hide calls. They will be null if the window is hidden and must be
// queried each time it is shown.
//
//   - [PropWindowWaylandDisplayPointer]: the wl_display associated with
//     the window
//   - [PropWindowWaylandSurfacePointer]: the wl_surface associated with
//     the window
//   - [PropWindowWaylandViewportPointer]: the wp_viewport associated
//     with the window
//   - [PropWindowWaylandEGLWindowPointer]: the wl_egl_window
//     associated with the window
//   - [PropWindowWaylandXDGSurfacePointer]: the xdg_surface associated
//     with the window
//   - [PropWindowWaylandXDGToplevelPointer]: the xdg_toplevel role
//     associated with the window
//   - [PropWindowWaylandXDGToplevelExportHandleString]: the export
//     handle associated with the window
//   - [PropWindowWaylandXDGPopupPointer]: the xdg_popup role
//     associated with the window
//   - [PropWindowWaylandXDGPositionerPointer]: the xdg_positioner
//     associated with the window, in popup mode
//
// On X11:
//
//   - [PropWindowX11DisplayPointer]: the X11 Display associated with
//     the window
//   - [PropWindowX11ScreenNumber]: the screen number associated with
//     the window
//   - [PropWindowX11WindowNumber]: the X11 Window associated with the
//     window
//
// window: the window to query.
//
// Returns a valid property ID on success or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowProperties
func (window *Window) Properties() (PropertiesID, error) {
	props := (PropertiesID)(C.SDL_GetWindowProperties((*C.SDL_Window)(window)))
	if props == 0 {
		return 0, getError()
	}
	return props, nil
}

const PropWindowShapePointer = "SDL.window.shape"
const PropWindowHDREnabledBoolean = "SDL.window.HDR_enabled"
const PropWindowSDRWhiteLevelFloat = "SDL.window.SDR_white_level"
const PropWindowHDRHeadroomFloat = "SDL.window.HDR_headroom"
const PropWindowAndroidWindowPointer = "SDL.window.android.window"
const PropWindowAndroidSurfacePointer = "SDL.window.android.surface"
const PropWindowUikitWindowPointer = "SDL.window.uikit.window"
const PropWindowUikitMetalViewTagNumber = "SDL.window.uikit.metal_view_tag"
const PropWindowUikitOpenglFramebufferNumber = "SDL.window.uikit.opengl.framebuffer"
const PropWindowUikitOpenglRenderbufferNumber = "SDL.window.uikit.opengl.renderbuffer"
const PropWindowUikitOpenglResolveFramebufferNumber = "SDL.window.uikit.opengl.resolve_framebuffer"
const PropWindowKMSDRMDeviceIndexNumber = "SDL.window.kmsdrm.dev_index"
const PropWindowKMSDRMDRMFDNumber = "SDL.window.kmsdrm.drm_fd"
const PropWindowKMSDRMGBMDevicePointer = "SDL.window.kmsdrm.gbm_dev"
const PropWindowCocoaWindowPointer = "SDL.window.cocoa.window"
const PropWindowCocoaMetalViewTagNumber = "SDL.window.cocoa.metal_view_tag"
const PropWindowOpenVROverlayID = "SDL.window.openvr.overlay_id"
const PropWindowVivanteDisplayPointer = "SDL.window.vivante.display"
const PropWindowVivanteWindowPointer = "SDL.window.vivante.window"
const PropWindowVivanteSurfacePointer = "SDL.window.vivante.surface"
const PropWindowWin32HWNDPointer = "SDL.window.win32.hwnd"
const PropWindowWin32HDCPointer = "SDL.window.win32.hdc"
const PropWindowWin32InstancePointer = "SDL.window.win32.instance"
const PropWindowWaylandDisplayPointer = "SDL.window.wayland.display"
const PropWindowWaylandSurfacePointer = "SDL.window.wayland.surface"
const PropWindowWaylandViewportPointer = "SDL.window.wayland.viewport"
const PropWindowWaylandEGLWindowPointer = "SDL.window.wayland.egl_window"
const PropWindowWaylandXDGSurfacePointer = "SDL.window.wayland.xdg_surface"
const PropWindowWaylandXDGToplevelPointer = "SDL.window.wayland.xdg_toplevel"
const PropWindowWaylandXDGToplevelExportHandleString = "SDL.window.wayland.xdg_toplevel_export_handle"
const PropWindowWaylandXDGPopupPointer = "SDL.window.wayland.xdg_popup"
const PropWindowWaylandXDGPositionerPointer = "SDL.window.wayland.xdg_positioner"
const PropWindowX11DisplayPointer = "SDL.window.x11.display"
const PropWindowX11ScreenNumber = "SDL.window.x11.screen"
const PropWindowX11WindowNumber = "SDL.window.x11.window"

// Get the window flags.
//
// window: the window to query.
//
// Returns a mask of the [WindowFlags] associated with window.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowFlags
func (window *Window) Flags() WindowFlags {
	return (WindowFlags)(C.SDL_GetWindowFlags((*C.SDL_Window)(window)))
}

// Set the title of a window.
//
// window: the window to change.
//
// title: the desired window title.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowTitle
func (window *Window) SetTitle(title string) error {
	if !C.SDL_SetWindowTitle((*C.SDL_Window)(window), tmpstring(title)) {
		return getError()
	}
	return nil
}

// Get the title of a window.
//
// window: the window to query.
//
// Returns the title of the window or "" if there is no title.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowTitle
func (window *Window) Title() string {
	return C.GoString(C.SDL_GetWindowTitle((*C.SDL_Window)(window)))
}

// Set the icon for a window.
//
// If this function is passed a surface with alternate representations, the
// surface will be interpreted as the content to be used for 100% display
// scale, and the alternate representations will be used for high DPI
// situations. For example, if the original surface is 32x32, then on a 2x
// macOS display or 200% display scale on Windows, a 64x64 version of the
// image will be used, if available. If a matching version of the image isn't
// available, the closest larger size image will be downscaled to the
// appropriate size and be used instead, if available. Otherwise, the closest
// smaller image will be upscaled and be used instead.
//
// window: the window to change.
//
// icon: a [Surface] structure containing the icon for the window.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowIcon
func (window *Window) SetIcon(icon *Surface) error {
	if !C.SDL_SetWindowIcon((*C.SDL_Window)(window), icon.internal) {
		return getError()
	}
	return nil
}

// Request that the window's position be set.
//
// If the window is in an exclusive fullscreen or maximized state, this
// request has no effect.
//
// This can be used to reposition fullscreen-desktop windows onto a different
// display, however, as exclusive fullscreen windows are locked to a specific
// display, they can only be repositioned programmatically via
// [Window.SetFullscreenMode].
//
// On some windowing systems this request is asynchronous and the new
// coordinates may not have have been applied immediately upon the return of
// this function. If an immediate change is required, call [Window.Sync] to
// block until the changes have taken effect.
//
// When the window position changes, an [EventWindowMoved] event will be
// emitted with the window's new coordinates. Note that the new coordinates
// may not match the exact coordinates requested, as some windowing systems
// can restrict the position of the window in certain scenarios (e.g.
// constraining the position so the window is always within desktop bounds).
// Additionally, as this is just a request, it can be denied by the windowing
// system.
//
// window: the window to reposition.
//
// x: the x coordinate of the window, or [WindowposCentered] or
// [WindowposUndefined].
//
// y: the y coordinate of the window, or [WindowposCentered] or
// [WindowposUndefined].
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowPosition
func (window *Window) SetPosition(x int, y int) error {
	if !C.SDL_SetWindowPosition((*C.SDL_Window)(window), (C.int)(x), (C.int)(y)) {
		return getError()
	}
	return nil
}

// Get the position of a window.
//
// This is the current position of the window as last reported by the
// windowing system.
//
// window: the window to query.
//
// Returns the window position on success or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowPosition
func (window *Window) Position() (x, y int, err error) {
	var cx, cy C.int
	if !C.SDL_GetWindowPosition((*C.SDL_Window)(window), &cx, &cy) {
		return 0, 0, getError()
	}
	return int(cx), int(cy), nil
}

// Request that the size of a window's client area be set.
//
// If the window is in a fullscreen or maximized state, this request has no
// effect.
//
// To change the exclusive fullscreen mode of a window, use
// [Window.SetFullscreenMode].
//
// On some windowing systems, this request is asynchronous and the new window
// size may not have have been applied immediately upon the return of this
// function. If an immediate change is required, call [Window.Sync] to
// block until the changes have taken effect.
//
// When the window size changes, an [EventWindowResized] event will be
// emitted with the new window dimensions. Note that the new dimensions may
// not match the exact size requested, as some windowing systems can restrict
// the window size in certain scenarios (e.g. constraining the size of the
// content area to remain within the usable desktop bounds). Additionally, as
// this is just a request, it can be denied by the windowing system.
//
// window: the window to change.
//
// w: the width of the window, must be > 0.
//
// h: the height of the window, must be > 0.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowSize
func (window *Window) SetSize(w int, h int) error {
	if !C.SDL_SetWindowSize((*C.SDL_Window)(window), (C.int)(w), (C.int)(h)) {
		return getError()
	}
	return nil
}

// Get the size of a window's client area.
//
// The window pixel size may differ from its window coordinate size if the
// window is on a high pixel density display. Use [Window.SizeInPixels]
// or [Renderer.OutputSize] to get the real client area size in pixels.
//
// window: the window to query the width and height from.
//
// Returns the window size or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowSize
func (window *Window) Size() (w, h int, err error) {
	var cw, ch C.int
	if !C.SDL_GetWindowSize((*C.SDL_Window)(window), &cw, &ch) {
		return 0, 0, getError()
	}
	return int(cw), int(ch), nil
}

// Get the safe area for this window.
//
// Some devices have portions of the screen which are partially obscured or
// not interactive, possibly due to on-screen controls, curved edges, camera
// notches, TV overscan, etc. This function provides the area of the window
// which is safe to have interactable content. You should continue rendering
// into the rest of the window, but it should not contain visually important
// or interactible content.
//
// window: the window to query.
//
// Returns the client area that is safe for interactive content or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowSafeArea
func (window *Window) SafeArea() (Rect, error) {
	var rect C.SDL_Rect
	if !C.SDL_GetWindowSafeArea((*C.SDL_Window)(window), &rect) {
		return Rect{}, getError()
	}
	return Rect{int(rect.x), int(rect.y), int(rect.w), int(rect.h)}, nil
}

// Request that the aspect ratio of a window's client area be set.
//
// The aspect ratio is the ratio of width divided by height, e.g. 2560x1600
// would be 1.6. Larger aspect ratios are wider and smaller aspect ratios are
// narrower.
//
// If, at the time of this request, the window in a fixed-size state, such as
// maximized or fullscreen, the request will be deferred until the window
// exits this state and becomes resizable again.
//
// On some windowing systems, this request is asynchronous and the new window
// aspect ratio may not have have been applied immediately upon the return of
// this function. If an immediate change is required, call [Window.Sync] to
// block until the changes have taken effect.
//
// When the window size changes, an [EventWindowResized] event will be
// emitted with the new window dimensions. Note that the new dimensions may
// not match the exact aspect ratio requested, as some windowing systems can
// restrict the window size in certain scenarios (e.g. constraining the size
// of the content area to remain within the usable desktop bounds).
// Additionally, as this is just a request, it can be denied by the windowing
// system.
//
// window: the window to change.
//
// minAspect: the minimum aspect ratio of the window, or 0.0f for no
// limit.
//
// maxAspect: the maximum aspect ratio of the window, or 0.0f for no
// limit.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowAspectRatio
func (window *Window) SetAspectRatio(minAspect float32, maxAspect float32) error {
	if !C.SDL_SetWindowAspectRatio((*C.SDL_Window)(window), (C.float)(minAspect), (C.float)(maxAspect)) {
		return getError()
	}
	return nil
}

// Get the size of a window's client area.
//
// window: the window to query the width and height from.
//
// Returns the window aspect ratio or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowAspectRatio
func (window *Window) AspectRatio() (minAspect, maxAspect float32, err error) {
	var min, max C.float
	if !C.SDL_GetWindowAspectRatio((*C.SDL_Window)(window), &min, &max) {
		return 0, 0, getError()
	}
	return float32(min), float32(max), getError()
}

// Get the size of a window's borders (decorations) around the client area.
//
// Note: This function may fail on systems where the window has not yet been
// decorated by the display server (for example, immediately after calling
// [CreateWindow]). It is recommended that you wait at least until the
// window has been presented and composited, so that the window system has a
// chance to decorate the window and provide the border dimensions to SDL.
//
// This function also returns false if getting the information is not
// supported.
//
// window: the window to query the size values of the border
// (decorations) from.
//
// Returns the sizes of the borders or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowBordersSize
func (window *Window) BordersSize() (top, left, bottom, right int, err error) {
	var ctop, cleft, cbottom, cright C.int
	if !C.SDL_GetWindowBordersSize((*C.SDL_Window)(window), &ctop, &cleft, &cbottom, &cright) {
		return 0, 0, 0, 0, getError()
	}
	return int(ctop), int(cleft), int(cbottom), int(cright), nil
}

// Get the size of a window's client area, in pixels.
//
// window: the window from which the drawable size should be queried.
//
// Returns the window size in pixels or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowSizeInPixels
func (window *Window) SizeInPixels() (w, h int, err error) {
	var cw, ch C.int
	if !C.SDL_GetWindowSizeInPixels((*C.SDL_Window)(window), &cw, &ch) {
		return 0, 0, getError()
	}
	return int(cw), int(ch), nil
}

// Set the minimum size of a window's client area.
//
// window: the window to change.
//
// minW: the minimum width of the window, or 0 for no limit.
//
// minH: the minimum height of the window, or 0 for no limit.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowMinimumSize
func (window *Window) SetMinimumSize(minW int, minH int) error {
	if !C.SDL_SetWindowMinimumSize((*C.SDL_Window)(window), (C.int)(minW), (C.int)(minH)) {
		return getError()
	}
	return nil
}

// Get the minimum size of a window's client area.
//
// window: the window to query.
//
// Returns the minimum size or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowMinimumSize
func (window *Window) GetMinimumSize() (w, h int, err error) {
	var cw, ch C.int
	if !C.SDL_GetWindowMinimumSize((*C.SDL_Window)(window), &cw, &ch) {
		return 0, 0, getError()
	}
	return int(cw), int(ch), nil
}

// Set the maximum size of a window's client area.
//
// window: the window to change.
//
// maxW: the maximum width of the window, or 0 for no limit.
//
// maxH: the maximum height of the window, or 0 for no limit.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowMaximumSize
func (window *Window) SetMaximumSize(maxW int, maxH int) error {
	if !C.SDL_SetWindowMaximumSize((*C.SDL_Window)(window), (C.int)(maxW), (C.int)(maxH)) {
		return getError()
	}
	return nil
}

// Get the maximum size of a window's client area.
//
// window: the window to query.
//
// Returns the maximum size of the window or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowMaximumSize
func (window *Window) GetMaximumSize() (w, h int, err error) {
	var cw, ch C.int
	if !C.SDL_GetWindowMaximumSize((*C.SDL_Window)(window), &cw, &ch) {
		return 0, 0, getError()
	}
	return int(cw), int(ch), nil
}

// Set the border state of a window.
//
// This will add or remove the window's [WindowBorderless] flag and add
// or remove the border from the actual window. This is a no-op if the
// window's border already matches the requested state.
//
// You can't change the border state of a fullscreen window.
//
// window: the window of which to change the border state.
//
// bordered: false to remove border, true to add border.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowBordered
func (window *Window) SetBordered(bordered bool) error {
	if !C.SDL_SetWindowBordered((*C.SDL_Window)(window), (C.bool)(bordered)) {
		return getError()
	}
	return nil
}

// Set the user-resizable state of a window.
//
// This will add or remove the window's [WindowResizable] flag and
// allow/disallow user resizing of the window. This is a no-op if the window's
// resizable state already matches the requested state.
//
// You can't change the resizable state of a fullscreen window.
//
// window: the window of which to change the resizable state.
//
// resizable: true to allow resizing, false to disallow.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowResizable
func (window *Window) SetResizable(resizable bool) error {
	if !C.SDL_SetWindowResizable((*C.SDL_Window)(window), (C.bool)(resizable)) {
		return getError()
	}
	return nil
}

// Set the window to always be above the others.
//
// This will add or remove the window's [WindowAlwaysOnTop] flag. This
// will bring the window to the front and keep the window above the rest.
//
// window: the window of which to change the always on top state.
//
// onTop: true to set the window always on top, false to disable.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowAlwaysOnTop
func (window *Window) SetAlwaysOnTop(onTop bool) error {
	if !C.SDL_SetWindowAlwaysOnTop((*C.SDL_Window)(window), (C.bool)(onTop)) {
		return getError()
	}
	return nil
}

// Show a window.
//
// window: the window to show.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ShowWindow
func (window *Window) Show() error {
	if !C.SDL_ShowWindow((*C.SDL_Window)(window)) {
		return getError()
	}
	return nil
}

// Hide a window.
//
// window: the window to hide.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HideWindow
func (window *Window) Hide() error {
	if !C.SDL_HideWindow((*C.SDL_Window)(window)) {
		return getError()
	}
	return nil
}

// Request that a window be raised above other windows and gain the input
// focus.
//
// The result of this request is subject to desktop window manager policy,
// particularly if raising the requested window would result in stealing focus
// from another application. If the window is successfully raised and gains
// input focus, an [EventWindowFocusGained] event will be emitted, and
// the window will have the [WindowInputFocus] flag set.
//
// window: the window to raise.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RaiseWindow
func (window *Window) Raise() error {
	if !C.SDL_RaiseWindow((*C.SDL_Window)(window)) {
		return getError()
	}
	return nil
}

// Request that the window be made as large as possible.
//
// Non-resizable windows can't be maximized. The window must have the
// [WindowResizable] flag set, or this will have no effect.
//
// On some windowing systems this request is asynchronous and the new window
// state may not have have been applied immediately upon the return of this
// function. If an immediate change is required, call [Window.Sync] to
// block until the changes have taken effect.
//
// When the window state changes, an [EventWindowMaximized] event will be
// emitted. Note that, as this is just a request, the windowing system can
// deny the state change.
//
// When maximizing a window, whether the constraints set via
// [Window.SetMaximumSize] are honored depends on the policy of the window
// manager. Win32 and macOS enforce the constraints when maximizing, while X11
// and Wayland window managers may vary.
//
// window: the window to maximize.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MaximizeWindow
func (window *Window) Maximize() error {
	if !C.SDL_MaximizeWindow((*C.SDL_Window)(window)) {
		return getError()
	}
	return nil
}

// Request that the window be minimized to an iconic representation.
//
// If the window is in a fullscreen state, this request has no direct effect.
// It may alter the state the window is returned to when leaving fullscreen.
//
// On some windowing systems this request is asynchronous and the new window
// state may not have been applied immediately upon the return of this
// function. If an immediate change is required, call [Window.Sync] to
// block until the changes have taken effect.
//
// When the window state changes, an [EventWindowMinimized] event will be
// emitted. Note that, as this is just a request, the windowing system can
// deny the state change.
//
// window: the window to minimize.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MinimizeWindow
func (window *Window) Minimize() error {
	if !C.SDL_MinimizeWindow((*C.SDL_Window)(window)) {
		return getError()
	}
	return nil
}

// Request that the size and position of a minimized or maximized window be
// restored.
//
// If the window is in a fullscreen state, this request has no direct effect.
// It may alter the state the window is returned to when leaving fullscreen.
//
// On some windowing systems this request is asynchronous and the new window
// state may not have have been applied immediately upon the return of this
// function. If an immediate change is required, call [Window.Sync] to
// block until the changes have taken effect.
//
// When the window state changes, an [EventWindowRestored] event will be
// emitted. Note that, as this is just a request, the windowing system can
// deny the state change.
//
// window: the window to restore.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RestoreWindow
func (window *Window) Restore() error {
	if !C.SDL_RestoreWindow((*C.SDL_Window)(window)) {
		return getError()
	}
	return nil
}

// Request that the window's fullscreen state be changed.
//
// By default a window in fullscreen state uses borderless fullscreen desktop
// mode, but a specific exclusive display mode can be set using
// [Window.SetFullscreenMode].
//
// On some windowing systems this request is asynchronous and the new
// fullscreen state may not have have been applied immediately upon the return
// of this function. If an immediate change is required, call [Window.Sync]
// to block until the changes have taken effect.
//
// When the window state changes, an [EventWindowEnterFullscreen] or
// [EventWindowLeaveFullscreen] event will be emitted. Note that, as this
// is just a request, it can be denied by the windowing system.
//
// window: the window to change.
//
// fullscreen: true for fullscreen mode, false for windowed mode.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowFullscreen
func (window *Window) SetFullscreen(fullscreen bool) error {
	if !C.SDL_SetWindowFullscreen((*C.SDL_Window)(window), (C.bool)(fullscreen)) {
		return getError()
	}
	return nil
}

// Block until any pending window state is finalized.
//
// On asynchronous windowing systems, this acts as a synchronization barrier
// for pending window state. It will attempt to wait until any pending window
// state has been applied and is guaranteed to return within finite time. Note
// that for how long it can potentially block depends on the underlying window
// system, as window state changes may involve somewhat lengthy animations
// that must complete before the window is in its final requested state.
//
// On windowing systems where changes are immediate, this does nothing.
//
// window: the window for which to wait for the pending state to be
// applied.
//
// Returns true on success or false if the operation timed out before the
// window was in the requested state.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SyncWindow
func (window *Window) Sync() bool {
	return (bool)(C.SDL_SyncWindow((*C.SDL_Window)(window)))
}

// Return whether the window has a surface associated with it.
//
// window: the window to query.
//
// Returns true if there is a surface associated with the window, or false
// otherwise.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WindowHasSurface
func (window *Window) HasSurface() bool {
	return (bool)(C.SDL_WindowHasSurface((*C.SDL_Window)(window)))
}

// Get the SDL surface associated with the window.
//
// A new surface will be created with the optimal format for the window, if
// necessary. This surface will be freed when the window is destroyed. Do not
// free this surface.
//
// This surface will be invalidated if the window is resized. After resizing a
// window this function must be called again to return a valid surface.
//
// You may not combine this with 3D or the rendering API on this window.
//
// This function is affected by [HintFramebufferAcceleration].
//
// window: the window to query.
//
// # Returns the surface associated with the window or an error
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowSurface
func (window *Window) Surface() (*Surface, error) {
	s := C.SDL_GetWindowSurface((*C.SDL_Window)(window))
	if s == nil {
		return nil, getError()
	}
	return &Surface{s, nil}, nil
}

// Toggle VSync for the window surface.
//
// When a window surface is created, vsync defaults to
// [WindowSurfaceVSyncDisabled].
//
// The vsync parameter can be 1 to synchronize present with every vertical
// refresh, 2 to synchronize present with every second vertical refresh, etc.,
// [WindowSurfaceVSyncAdaptive] for late swap tearing (adaptive vsync),
// or [WindowSurfaceVSyncDisabled] to disable. Not every value is
// supported by every driver, so you should check the return value to see
// whether the requested setting is supported.
//
// window: the window.
//
// vsync: the vertical refresh sync interval.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowSurfaceVSync
func (window *Window) SetSurfaceVSync(vsync int) error {
	if !C.SDL_SetWindowSurfaceVSync((*C.SDL_Window)(window), (C.int)(vsync)) {
		return getError()
	}
	return nil
}

const WindowSurfaceVSyncDisabled = 0
const WindowSurfaceVSyncAdaptive = -1

// Get VSync for the window surface.
//
// window: the window to query.
//
// vsync: an int filled with the current vertical refresh sync interval.
// See [Window.SetSurfaceVSync] for the meaning of the value.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowSurfaceVSync
func (window *Window) SurfaceVSync() (int, error) {
	var vsync C.int
	if !C.SDL_GetWindowSurfaceVSync((*C.SDL_Window)(window), &vsync) {
		return 0, getError()
	}
	return int(vsync), nil
}

// Copy the window surface to the screen.
//
// This is the function you use to reflect any changes to the surface on the
// screen.
//
// This function is equivalent to the SDL 1.2 API SDL_Flip().
//
// window: the window to update.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_UpdateWindowSurface
func (window *Window) UpdateSurface() error {
	if !C.SDL_UpdateWindowSurface((*C.SDL_Window)(window)) {
		return getError()
	}
	return nil
}

// Copy areas of the window surface to the screen.
//
// This is the function you use to reflect changes to portions of the surface
// on the screen.
//
// This function is equivalent to the SDL 1.2 API SDL_UpdateRects().
//
// Note that this function will update _at least_ the rectangles specified,
// but this is only intended as an optimization; in practice, this might
// update more of the screen (or all of the screen!), depending on what method
// SDL uses to send pixels to the system.
//
// window: the window to update.
//
// rects: a slice of [Rect] structures representing areas of the
// surface to copy, in pixels.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_UpdateWindowSurfaceRects
func (window *Window) UpdateSurfaceRects(rects []Rect) error {
	crects := make([]C.SDL_Rect, len(rects))
	for i, r := range rects {
		crects[i] = C.SDL_Rect{C.int(r.X), C.int(r.Y), C.int(r.W), C.int(r.H)}
	}
	if !C.SDL_UpdateWindowSurfaceRects((*C.SDL_Window)(window), unsafe.SliceData(crects), (C.int)(len(rects))) {
		return getError()
	}
	return nil
}

// Destroy the surface associated with the window.
//
// window: the window to update.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DestroyWindowSurface
func (window *Window) DestroySurface() error {
	if !C.SDL_DestroyWindowSurface((*C.SDL_Window)(window)) {
		return getError()
	}
	return nil
}

// Set a window's keyboard grab mode.
//
// Keyboard grab enables capture of system keyboard shortcuts like Alt+Tab or
// the Meta/Super key. Note that not all system keyboard shortcuts can be
// captured by applications (one example is Ctrl+Alt+Del on Windows).
//
// This is primarily intended for specialized applications such as VNC clients
// or VM frontends. Normal games should not use keyboard grab.
//
// When keyboard grab is enabled, SDL will continue to handle Alt+Tab when the
// window is full-screen to ensure the user is not trapped in your
// application. If you have a custom keyboard shortcut to exit fullscreen
// mode, you may suppress this behavior with
// [HintAllowAltTabWhileGrabbed].
//
// If the caller enables a grab while another window is currently grabbed, the
// other window loses its grab in favor of the caller's window.
//
// window: the window for which the keyboard grab mode should be set.
//
// grabbed: this is true to grab keyboard, and false to release.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowKeyboardGrab
func (window *Window) SetKeyboardGrab(grabbed bool) error {
	if !C.SDL_SetWindowKeyboardGrab((*C.SDL_Window)(window), (C.bool)(grabbed)) {
		return getError()
	}
	return nil
}

// Set a window's mouse grab mode.
//
// Mouse grab confines the mouse cursor to the window.
//
// window: the window for which the mouse grab mode should be set.
//
// grabbed: this is true to grab mouse, and false to release.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowMouseGrab
func (window *Window) SetMouseGrab(grabbed bool) error {
	if !C.SDL_SetWindowMouseGrab((*C.SDL_Window)(window), (C.bool)(grabbed)) {
		return getError()
	}
	return nil
}

// Get a window's keyboard grab mode.
//
// window: the window to query.
//
// Returns true if keyboard is grabbed, and false otherwise.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowKeyboardGrab
func (window *Window) GetKeyboardGrab() bool {
	return (bool)(C.SDL_GetWindowKeyboardGrab((*C.SDL_Window)(window)))
}

// Get a window's mouse grab mode.
//
// window: the window to query.
//
// Returns true if mouse is grabbed, and false otherwise.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowMouseGrab
func (window *Window) GetMouseGrab() bool {
	return (bool)(C.SDL_GetWindowMouseGrab((*C.SDL_Window)(window)))
}

// Get the window that currently has an input grab enabled.
//
// Returns the window if input is grabbed or nil otherwise.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGrabbedWindow
func GetGrabbedWindow() *Window {
	return (*Window)(C.SDL_GetGrabbedWindow())
}

// Confines the cursor to the specified area of a window.
//
// Note that this does NOT grab the cursor, it only defines the area a cursor
// is restricted to when the window has mouse focus.
//
// window: the window that will be associated with the barrier.
//
// rect: a rectangle area in window-relative coordinates. If nil the
// barrier for the specified window will be destroyed.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowMouseRect
func (window *Window) SetMouseRect(rect *Rect) error {
	var crect *C.SDL_Rect
	if rect != nil {
		crect = &C.SDL_Rect{C.int(rect.X), C.int(rect.Y), C.int(rect.W), C.int(rect.H)}
	}
	if !C.SDL_SetWindowMouseRect((*C.SDL_Window)(window), crect) {
		return getError()
	}
	return nil
}

// Get the mouse confinement rectangle of a window.
//
// window: the window to query.
//
// Returns a pointer to the mouse confinement rectangle of a window, or nil
// if there isn't one.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowMouseRect
func (window *Window) GetMouseRect() *Rect {
	rect := C.SDL_GetWindowMouseRect((*C.SDL_Window)(window))
	if rect == nil {
		return nil
	}
	return &Rect{int(rect.x), int(rect.y), int(rect.w), int(rect.h)}
}

// Set the opacity for a window.
//
// The parameter opacity will be clamped internally between 0.0f
// (transparent) and 1.0f (opaque).
//
// This function also returns false if setting the opacity isn't supported.
//
// window: the window which will be made transparent or opaque.
//
// opacity: the opacity value (0.0f - transparent, 1.0f - opaque).
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowOpacity
func (window *Window) SetOpacity(opacity float32) error {
	if !C.SDL_SetWindowOpacity((*C.SDL_Window)(window), (C.float)(opacity)) {
		return getError()
	}
	return nil
}

// Get the opacity of a window.
//
// If transparency isn't supported on this platform, opacity will be returned
// as 1.0f without error.
//
// window: the window to get the current opacity value from.
//
// Returns the opacity, (0.0f - transparent, 1.0f - opaque), or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowOpacity
func (window *Window) Opacity() (float32, error) {
	opacity := (float32)(C.SDL_GetWindowOpacity((*C.SDL_Window)(window)))
	if opacity == -1 {
		return -1, getError()
	}
	return opacity, nil
}

// Set the window as a child of a parent window.
//
// If the window is already the child of an existing window, it will be
// reparented to the new owner. Setting the parent window to nil unparents
// the window and removes child window status.
//
// If a parent window is hidden or destroyed, the operation will be
// recursively applied to child windows. Child windows hidden with the parent
// that did not have their hidden status explicitly set will be restored when
// the parent is shown.
//
// Attempting to set the parent of a window that is currently in the modal
// state will fail. Use [Window.SetModal] to cancel the modal status before
// attempting to change the parent.
//
// Popup windows cannot change parents and attempts to do so will fail.
//
// Setting a parent window that is currently the sibling or descendent of the
// child window results in undefined behavior.
//
// window: the window that should become the child of a parent.
//
// parent: the new parent window for the child window.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowParent
func (window *Window) SetParent(parent *Window) error {
	if !C.SDL_SetWindowParent((*C.SDL_Window)(window), (*C.SDL_Window)(parent)) {
		return getError()
	}
	return nil
}

// Toggle the state of the window as modal.
//
// To enable modal status on a window, the window must currently be the child
// window of a parent, or toggling modal status on will fail.
//
// window: the window on which to set the modal state.
//
// modal: true to toggle modal status on, false to toggle it off.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowModal
func (window *Window) SetModal(modal bool) error {
	if !C.SDL_SetWindowModal((*C.SDL_Window)(window), (C.bool)(modal)) {
		return getError()
	}
	return nil
}

// Set whether the window may have input focus.
//
// window: the window to set focusable state.
//
// focusable: true to allow input focus, false to not allow input focus.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowFocusable
func (window *Window) SetFocusable(focusable bool) error {
	if !C.SDL_SetWindowFocusable((*C.SDL_Window)(window), (C.bool)(focusable)) {
		return getError()
	}
	return nil
}

// Display the system-level window menu.
//
// This default window menu is provided by the system and on some platforms
// provides functionality for setting or changing privileged state on the
// window, such as moving it between workspaces or displays, or toggling the
// always-on-top property.
//
// On platforms or desktops where this is unsupported, this function does
// nothing.
//
// window: the window for which the menu will be displayed.
//
// x: the x coordinate of the menu, relative to the origin (top-left) of
// the client area.
//
// y: the y coordinate of the menu, relative to the origin (top-left) of
// the client area.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ShowWindowSystemMenu
func (window *Window) ShowSystemMenu(x int, y int) error {
	if !C.SDL_ShowWindowSystemMenu((*C.SDL_Window)(window), (C.int)(x), (C.int)(y)) {
		return getError()
	}
	return nil
}

// Possible return values from the [HitTest] callback.
//
// This function should only be called on the main thread.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HitTestResult
type HitTestResult uint32

const (
	HittestNormal            HitTestResult = iota // Region is normal. No special properties.
	HittestDraggable                              // Region can drag entire window.
	HittestResizeTopleft                          // Region is the resizable top-left corner border.
	HittestResizeTop                              // Region is the resizable top border.
	HittestResizeTopright                         // Region is the resizable top-right corner border.
	HittestResizeRight                            // Region is the resizable right border.
	HittestResizeBottomright                      // Region is the resizable bottom-right corner border.
	HittestResizeBottom                           // Region is the resizable bottom border.
	HittestResizeBottomleft                       // Region is the resizable bottom-left corner border.
	HittestResizeLeft                             // Region is the resizable left border.
)

// Callback used for hit-testing.
//
// win: the [Window] where hit-testing was set on.
//
// area: a [Point] which should be hit-tested.
//
// Returns a [HitTestResult] value.
//
// https://wiki.libsdl.org/SDL3/SDL_HitTest
type HitTest func(win *Window, area Point) HitTestResult

// export cb_HitTest
func cb_HitTest(win *C.SDL_Window, area *C.SDL_Point, data uintptr) C.SDL_HitTestResult {
	h := cgo.Handle(data)
	return (C.SDL_HitTestResult)(h.Value().(HitTest)((*Window)(win), Point{int(area.x), int(area.y)}))
}

// Provide a callback that decides if a window region has special properties.
//
// Normally windows are dragged and resized by decorations provided by the
// system window manager (a title bar, borders, etc), but for some apps, it
// makes sense to drag them from somewhere else inside the window itself; for
// example, one might have a borderless window that wants to be draggable from
// any part, or simulate its own title bar, etc.
//
// This function lets the app provide a callback that designates pieces of a
// given window as special. This callback is run during event processing if we
// need to tell the OS to treat a region of the window specially; the use of
// this callback is known as "hit testing."
//
// Mouse input may not be delivered to your application if it is within a
// special area; the OS will often apply that input to moving the window or
// resizing the window and not deliver it to the application.
//
// Specifying nil for a callback disables hit-testing. Hit-testing is
// disabled by default.
//
// Platforms that don't support this functionality will return false
// unconditionally, even if you're attempting to disable hit-testing.
//
// Your callback may fire at any time, and its firing does not indicate any
// specific behavior (for example, on Windows, this certainly might fire when
// the OS is deciding whether to drag your window, but it fires for lots of
// other reasons, too, some unrelated to anything you probably care about _and
// when the mouse isn't actually at the location it is testing_). Since this
// can fire at any time, you should try to keep your callback efficient,
// devoid of allocations, etc.
//
// window: the window to set hit-testing on.
//
// callback: the function to call when doing a hit-test.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowHitTest
func (window *Window) SetHitTest(callback HitTest) error {
	if callback == nil {
		if !C.SDL_SetWindowHitTest((*C.SDL_Window)(window), nil, nil) {
			return getError()
		}
		return nil
	}
	h := cgo.NewHandle(callback)
	if !C.wrap_SDL_SetWindowHitTest((*C.SDL_Window)(window), C.uintptr_t(h)) {
		h.Delete()
		return getError()
	}
	return nil
}

// Set the shape of a transparent window.
//
// This sets the alpha channel of a transparent window and any fully
// transparent areas are also transparent to mouse clicks. If you are using
// something besides the SDL render API, then you are responsible for drawing
// the alpha channel of the window to match the shape alpha channel to get
// consistent cross-platform results.
//
// The shape is copied inside this function, so you can free it afterwards. If
// your shape surface changes, you should call [Window.SetShape] again to
// update the window. This is an expensive operation, so should be done
// sparingly.
//
// The window must have been created with the [WindowTransparent] flag.
//
// window: the window.
//
// shape: the surface representing the shape of the window, or nil to
// remove any current shape.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowShape
func (window *Window) SetShape(shape *Surface) error {
	var s *C.SDL_Surface
	if shape != nil {
		s = shape.internal
	}
	if !C.SDL_SetWindowShape((*C.SDL_Window)(window), s) {
		return getError()
	}
	return nil
}

// Request a window to demand attention from the user.
//
// window: the window to be flashed.
//
// operation: the operation to perform.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_FlashWindow
func (window *Window) Flash(operation FlashOperation) error {
	if !C.SDL_FlashWindow((*C.SDL_Window)(window), (C.SDL_FlashOperation)(operation)) {
		return getError()
	}
	return nil
}

// Destroy a window.
//
// Any child windows owned by the window will be recursively destroyed as
// well.
//
// Note that on some platforms, the visible window may not actually be removed
// from the screen until the SDL event loop is pumped again, even though the
// [Window] is no longer valid after this call.
//
// window: the window to destroy.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DestroyWindow
func (window *Window) Destroy() {
	C.SDL_DestroyWindow((*C.SDL_Window)(window))
}

// Check whether the screensaver is currently enabled.
//
// The screensaver is disabled by default.
//
// The default can also be changed using [HintVideoAllowScreensaver].
//
// Returns true if the screensaver is enabled, false if it is disabled.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ScreenSaverEnabled
func ScreenSaverEnabled() bool {
	return (bool)(C.SDL_ScreenSaverEnabled())
}

// Allow the screen to be blanked by a screen saver.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_EnableScreenSaver
func EnableScreenSaver() error {
	if !C.SDL_EnableScreenSaver() {
		return getError()
	}
	return nil
}

// Prevent the screen from being blanked by a screen saver.
//
// If you disable the screensaver, it is automatically re-enabled when SDL
// quits.
//
// The screensaver is disabled by default, but this may by changed by
// [HintVideoAllowScreensaver].
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DisableScreenSaver
func DisableScreenSaver() error {
	if !C.SDL_DisableScreenSaver() {
		return getError()
	}
	return nil
}

// \name OpenGL support functions

// Dynamically load an OpenGL library.
//
// This should be done after initializing the video driver, but before
// creating any OpenGL windows. If no OpenGL library is loaded, the default
// library will be loaded upon creation of the first OpenGL window.
//
// If you do this, you need to retrieve all of the GL functions used in your
// program from the dynamic library using [GL_GetProcAddress]().
//
// path: the platform dependent OpenGL library name, or "" to open the
// default OpenGL library.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GL_LoadLibrary
func GL_LoadLibrary(path string) error {
	if !C.SDL_GL_LoadLibrary(tmpstring(path)) {
		return getError()
	}
	return nil
}

// Get an OpenGL function by name.
//
// If the GL library is loaded at runtime with [GL_LoadLibrary], then all
// GL functions must be retrieved this way. Usually this is used to retrieve
// function pointers to OpenGL extensions.
//
// There are some quirks to looking up OpenGL functions that require some
// extra care from the application. If you code carefully, you can handle
// these quirks without any platform-specific code, though:
//
//   - On Windows, function pointers are specific to the current GL context;
//     this means you need to have created a GL context and made it current
//     before calling [GL_GetProcAddress]. If you recreate your context or
//     create a second context, you should assume that any existing function
//     pointers aren't valid to use with it. This is (currently) a
//     Windows-specific limitation, and in practice lots of drivers don't suffer
//     this limitation, but it is still the way the wgl API is documented to
//     work and you should expect crashes if you don't respect it. Store a copy
//     of the function pointers that comes and goes with context lifespan.
//   - On X11, function pointers returned by this function are valid for any
//     context, and can even be looked up before a context is created at all.
//     This means that, for at least some common OpenGL implementations, if you
//     look up a function that doesn't exist, you'll get a non-nil result that
//     is _NOT_ safe to call. You must always make sure the function is actually
//     available for a given GL context before calling it, by checking for the
//     existence of the appropriate extension with [GL_ExtensionSupported],
//     or verifying that the version of OpenGL you're using offers the function
//     as core functionality.
//   - Some OpenGL drivers, on all platforms, *will* return nil if a function
//     isn't supported, but you can't count on this behavior. Check for
//     extensions you use, and if you get a nil anyway, act as if that
//     extension wasn't available. This is probably a bug in the driver, but you
//     can code defensively for this scenario anyhow.
//   - Just because you're on Linux/Unix, don't assume you'll be using X11.
//     Next-gen display servers are waiting to replace it, and may or may not
//     make the same promises about function pointers.
//   - OpenGL function pointers must be declared APIENTRY as in the example
//     code. This will ensure the proper calling convention is followed on
//     platforms where this matters (Win32) thereby avoiding stack corruption.
//
// proc: the name of an OpenGL function.
//
// Returns a pointer to the named OpenGL function. The returned pointer
// should be cast to the appropriate function signature.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GL_GetProcAddress
func GL_GetProcAddress(proc string) unsafe.Pointer {
	return (unsafe.Pointer)(C.SDL_GL_GetProcAddress(tmpstring(proc)))
}

// Get an EGL library function by name.
//
// If an EGL library is loaded, this function allows applications to get entry
// points for EGL functions. This is useful to provide to an EGL API and
// extension loader.
//
// proc: the name of an EGL function.
//
// Returns a pointer to the named EGL function. The returned pointer should
// be cast to the appropriate function signature.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_EGL_GetProcAddress
func EGL_GetProcAddress(proc string) unsafe.Pointer {
	return (unsafe.Pointer)(C.SDL_EGL_GetProcAddress(tmpstring(proc)))
}

// Unload the OpenGL library previously loaded by [GL_LoadLibrary].
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GL_UnloadLibrary
func GL_UnloadLibrary() {
	C.SDL_GL_UnloadLibrary()
}

// Check if an OpenGL extension is supported for the current context.
//
// This function operates on the current GL context; you must have created a
// context and it must be current before calling this function. Do not assume
// that all contexts you create will have the same set of extensions
// available, or that recreating an existing context will offer the same
// extensions again.
//
// While it's probably not a massive overhead, this function is not an O(1)
// operation. Check the extensions you care about after creating the GL
// context and save that information somewhere instead of calling the function
// every time you need to know.
//
// extension: the name of the extension to check.
//
// Returns true if the extension is supported, false otherwise.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GL_ExtensionSupported
func GL_ExtensionSupported(extension string) bool {
	return (bool)(C.SDL_GL_ExtensionSupported(tmpstring(extension)))
}

// Reset all previously set OpenGL context attributes to their default values.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GL_ResetAttributes
func GL_ResetAttributes() {
	C.SDL_GL_ResetAttributes()
	if currentEGLAttributeCallbacks != 0 {
		currentEGLAttributeCallbacks.Delete()
		currentEGLAttributeCallbacks = 0
	}
}

// Set an OpenGL window attribute before window creation.
//
// This function sets the OpenGL attribute attr to value. The requested
// attributes should be set before creating an OpenGL window. You should use
// [GL_GetAttribute] to check the values after creating the OpenGL
// context, since the values obtained can differ from the requested ones.
//
// attr: an [GLAttr] enum value specifying the OpenGL attribute to
// set.
//
// value: the desired value for the attribute.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GL_SetAttribute
func GL_SetAttribute(attr GLAttr, value int32) error {
	if !C.SDL_GL_SetAttribute((C.SDL_GLAttr)(attr), (C.int)(value)) {
		return getError()
	}
	return nil
}

// Get the actual value for an attribute from the current context.
//
// attr: an [GLAttr] enum value specifying the OpenGL attribute to
// get.
//
// Returns the current value of attr or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GL_GetAttribute
func GL_GetAttribute(attr GLAttr) (int, error) {
	var value C.int
	if !C.SDL_GL_GetAttribute((C.SDL_GLAttr)(attr), &value) {
		return 0, getError()
	}
	return int(value), nil
}

// Create an OpenGL context for an OpenGL window, and make it current.
//
// Windows users new to OpenGL should note that, for historical reasons, GL
// functions added after OpenGL version 1.1 are not available by default.
// Those functions must be loaded at run-time, either with an OpenGL
// extension-handling library or with [GL_GetProcAddress] and its related
// functions.
//
// [GLContext] is opaque to the application.
//
// window: the window to associate with the context.
//
// Returns the OpenGL context associated with window or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GL_CreateContext
func GL_CreateContext(window *Window) (GLContext, error) {
	ctx := (GLContext)(C.SDL_GL_CreateContext((*C.SDL_Window)(window)))
	if ctx == nil {
		return nil, getError()
	}
	return ctx, nil
}

// Set up an OpenGL context for rendering into an OpenGL window.
//
// The context must have been created with a compatible window.
//
// window: the window to associate with the context.
//
// context: the OpenGL context to associate with the window.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GL_MakeCurrent
func GL_MakeCurrent(window *Window, context GLContext) error {
	if !C.SDL_GL_MakeCurrent((*C.SDL_Window)(window), (C.SDL_GLContext)(context)) {
		return getError()
	}
	return nil
}

// Get the currently active OpenGL window.
//
// Returns the currently active OpenGL window on success or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GL_GetCurrentWindow
func GL_GetCurrentWindow() (*Window, error) {
	window := (*Window)(C.SDL_GL_GetCurrentWindow())
	if window == nil {
		return nil, getError()
	}
	return window, nil
}

// Get the currently active OpenGL context.
//
// Returns the currently active OpenGL context or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GL_GetCurrentContext
func GL_GetCurrentContext() (GLContext, error) {
	ctx := (GLContext)(C.SDL_GL_GetCurrentContext())
	if ctx == nil {
		return nil, getError()
	}
	return ctx, nil
}

// Get the currently active EGL display.
//
// Returns the currently active EGL display or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_EGL_GetCurrentDisplay
func EGL_GetCurrentDisplay() (EGLDisplay, error) {
	display := (EGLDisplay)(C.SDL_EGL_GetCurrentDisplay())
	if display == nil {
		return nil, getError()
	}
	return display, nil
}

// Get the currently active EGL config.
//
// Returns the currently active EGL config or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_EGL_GetCurrentConfig
func EGL_GetCurrentConfig() (EGLConfig, error) {
	config := (EGLConfig)(C.SDL_EGL_GetCurrentConfig())
	if config == nil {
		return nil, getError()
	}
	return config, nil
}

// Get the EGL surface associated with the window.
//
// window: the window to query.
//
// Returns the [EGLSurface] pointer associated with the window or an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_EGL_GetWindowSurface
func EGL_GetWindowSurface(window *Window) (EGLSurface, error) {
	surface := (EGLSurface)(C.SDL_EGL_GetWindowSurface((*C.SDL_Window)(window)))
	if surface == nil {
		return nil, getError()
	}
	return surface, nil
}

type eglAttributeCallbacks struct {
	PlatformAttribCallback EGLAttribArrayCallback
	SurfaceAttribCallback  EGLIntArrayCallback
	ContextAttribCallback  EGLIntArrayCallback
}

//export cb_EGLAttribArrayCallback
func cb_EGLAttribArrayCallback(userdata uintptr) *C.SDL_EGLAttrib {
	h := cgo.Handle(userdata)
	array := h.Value().(*eglAttributeCallbacks).PlatformAttribCallback()
	if len(array) == 0 {
		return nil
	}
	result := (*C.SDL_EGLAttrib)(C.SDL_malloc(C.size_t(len(array)) * C.size_t(unsafe.Sizeof(array[0]))))
	copy(unsafe.Slice((*EGLAttrib)(result), len(array)), array)
	return result
}

//export cb_EGLSurfaceArrayCallback
func cb_EGLSurfaceArrayCallback(userdata uintptr, display C.SDL_EGLDisplay, config C.SDL_EGLConfig) *C.SDL_EGLint {
	h := cgo.Handle(userdata)
	array := h.Value().(*eglAttributeCallbacks).SurfaceAttribCallback(EGLDisplay(display), EGLConfig(config))
	if len(array) == 0 {
		return nil
	}
	result := (*C.SDL_EGLint)(C.SDL_malloc(C.size_t(len(array)) * C.size_t(unsafe.Sizeof(array[0]))))
	copy(unsafe.Slice((*EGLint)(result), len(array)), array)
	return result
}

//export cb_EGLContextArrayCallback
func cb_EGLContextArrayCallback(userdata uintptr, display C.SDL_EGLDisplay, config C.SDL_EGLConfig) *C.SDL_EGLint {
	h := cgo.Handle(userdata)
	array := h.Value().(*eglAttributeCallbacks).ContextAttribCallback(EGLDisplay(display), EGLConfig(config))
	if len(array) == 0 {
		return nil
	}
	result := (*C.SDL_EGLint)(C.SDL_malloc(C.size_t(len(array)) * C.size_t(unsafe.Sizeof(array[0]))))
	copy(unsafe.Slice((*EGLint)(result), len(array)), array)
	return result
}

// Sets the callbacks for defining custom [EGLAttrib] arrays for EGL
// initialization.
//
// NOTE: These callback pointers will be reset after [GL_ResetAttributes].
//
// platformAttribCallback: callback for attributes to pass to
// eglGetPlatformDisplay.
//
// surfaceAttribCallback: callback for attributes to pass to
// eglCreateSurface.
//
// contextAttribCallback: callback for attributes to pass to
// eglCreateContext.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
func EGL_SetAttributeCallbacks(platformAttribCallback EGLAttribArrayCallback, surfaceAttribCallback EGLIntArrayCallback, contextAttribCallback EGLIntArrayCallback) {
	if currentEGLAttributeCallbacks != 0 {
		currentEGLAttributeCallbacks.Delete()
	}
	h := cgo.NewHandle(&eglAttributeCallbacks{
		platformAttribCallback,
		surfaceAttribCallback,
		contextAttribCallback,
	})
	C.wrap_SDL_EGL_SetAttributeCallbacks(C.uintptr_t(h))
	currentEGLAttributeCallbacks = h
}

var currentEGLAttributeCallbacks cgo.Handle

// Set the swap interval for the current OpenGL context.
//
// Some systems allow specifying -1 for the interval, to enable adaptive
// vsync. Adaptive vsync works the same as vsync, but if you've already missed
// the vertical retrace for a given frame, it swaps buffers immediately, which
// might be less jarring for the user during occasional framerate drops. If an
// application requests adaptive vsync and the system does not support it,
// this function will fail and return false. In such a case, you should
// probably retry the call with 1 for the interval.
//
// Adaptive vsync is implemented for some glX drivers with
// GLX_EXT_swap_control_tear, and for some Windows drivers with
// WGL_EXT_swap_control_tear.
//
// Read more on the Khronos wiki:
// https://www.khronos.org/opengl/wiki/Swap_Interval#Adaptive_Vsync
//
// interval: 0 for immediate updates, 1 for updates synchronized with
// the vertical retrace, -1 for adaptive vsync.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GL_SetSwapInterval
func GL_SetSwapInterval(interval int32) error {
	if !C.SDL_GL_SetSwapInterval((C.int)(interval)) {
		return getError()
	}
	return nil
}

// Get the swap interval for the current OpenGL context.
//
// If the system can't determine the swap interval, or there isn't a valid
// current context, this function will set *interval to 0 as a safe default.
//
// interval: output interval value. 0 if there is no vertical retrace
// synchronization, 1 if the buffer swap is synchronized with
// the vertical retrace, and -1 if late swaps happen
// immediately instead of waiting for the next retrace.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GL_GetSwapInterval
func GL_GetSwapInterval() (int, error) {
	var interval C.int
	if !C.SDL_GL_GetSwapInterval(&interval) {
		return 0, getError()
	}
	return int(interval), nil
}

// Update a window with OpenGL rendering.
//
// This is used with double-buffered OpenGL contexts, which are the default.
//
// On macOS, make sure you bind 0 to the draw framebuffer before swapping the
// window, otherwise nothing will happen. If you aren't using
// glBindFramebuffer(), this is the default and you won't have to do anything
// extra.
//
// window: the window to change.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GL_SwapWindow
func GL_SwapWindow(window *Window) error {
	if !C.SDL_GL_SwapWindow((*C.SDL_Window)(window)) {
		return getError()
	}
	return nil
}

// Delete an OpenGL context.
//
// context: the OpenGL context to be deleted.
//
// Returns nil on success or an error on failure.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GL_DestroyContext
func GL_DestroyContext(context GLContext) error {
	if !C.SDL_GL_DestroyContext((C.SDL_GLContext)(context)) {
		return getError()
	}
	return nil
}
