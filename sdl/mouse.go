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
// #cgo noescape SDL_HasMouse
// #cgo nocallback SDL_HasMouse
// #cgo noescape SDL_GetMice
// #cgo nocallback SDL_GetMice
// #cgo noescape SDL_GetMouseNameForID
// #cgo nocallback SDL_GetMouseNameForID
// #cgo noescape SDL_GetMouseFocus
// #cgo nocallback SDL_GetMouseFocus
// #cgo noescape SDL_GetMouseState
// #cgo nocallback SDL_GetMouseState
// #cgo noescape SDL_GetGlobalMouseState
// #cgo nocallback SDL_GetGlobalMouseState
// #cgo noescape SDL_GetRelativeMouseState
// #cgo nocallback SDL_GetRelativeMouseState
// #cgo noescape SDL_WarpMouseInWindow
// #cgo nocallback SDL_WarpMouseInWindow
// #cgo noescape SDL_WarpMouseGlobal
// #cgo nocallback SDL_WarpMouseGlobal
// #cgo noescape SDL_SetWindowRelativeMouseMode
// #cgo nocallback SDL_SetWindowRelativeMouseMode
// #cgo noescape SDL_GetWindowRelativeMouseMode
// #cgo nocallback SDL_GetWindowRelativeMouseMode
// #cgo noescape SDL_CaptureMouse
// #cgo nocallback SDL_CaptureMouse
// #cgo noescape SDL_CreateCursor
// #cgo nocallback SDL_CreateCursor
// #cgo noescape SDL_CreateColorCursor
// #cgo nocallback SDL_CreateColorCursor
// #cgo noescape SDL_CreateSystemCursor
// #cgo nocallback SDL_CreateSystemCursor
// #cgo noescape SDL_SetCursor
// #cgo nocallback SDL_SetCursor
// #cgo noescape SDL_GetCursor
// #cgo nocallback SDL_GetCursor
// #cgo noescape SDL_GetDefaultCursor
// #cgo nocallback SDL_GetDefaultCursor
// #cgo noescape SDL_DestroyCursor
// #cgo noescape SDL_ShowCursor
// #cgo nocallback SDL_ShowCursor
// #cgo noescape SDL_HideCursor
// #cgo nocallback SDL_HideCursor
// #cgo noescape SDL_CursorVisible
// #cgo nocallback SDL_CursorVisible
// #include <SDL3/SDL.h>
import "C"
import "unsafe"

// # CategoryMouse
//
// Any GUI application has to deal with the mouse, and SDL provides functions
// to manage mouse input and the displayed cursor.
//
// Most interactions with the mouse will come through the event subsystem.
// Moving a mouse generates an SDL_EVENT_MOUSE_MOTION event, pushing a button
// generates SDL_EVENT_MOUSE_BUTTON_DOWN, etc, but one can also query the
// current state of the mouse at any time with SDL_GetMouseState().
//
// For certain games, it's useful to disassociate the mouse cursor from mouse
// input. An FPS, for example, would not want the player's motion to stop as
// the mouse hits the edge of the window. For these scenarios, use
// SDL_SetWindowRelativeMouseMode(), which hides the cursor, grabs mouse input
// to the window, and reads mouse input no matter how far it moves.
//
// Games that want the system to track the mouse but want to draw their own
// cursor can use SDL_HideCursor() and SDL_ShowCursor(). It might be more
// efficient to let the system manage the cursor, if possible, using
// SDL_SetCursor() with a custom image made through SDL_CreateColorCursor(),
// or perhaps just a specific system cursor from SDL_CreateSystemCursor().
//
// SDL can, on many platforms, differentiate between multiple connected mice,
// allowing for interesting input scenarios and multiplayer games. They can be
// enumerated with SDL_GetMice(), and SDL will send SDL_EVENT_MOUSE_ADDED and
// SDL_EVENT_MOUSE_REMOVED events as they are connected and unplugged.
//
// Since many apps only care about basic mouse input, SDL offers a virtual
// mouse device for touch and pen input, which often can make a desktop
// application work on a touchscreen phone without any code changes. Apps that
// care about touch/pen separately from mouse input should filter out events
// with a `which` field of SDL_TOUCH_MOUSEID/SDL_PEN_MOUSEID.

// This is a unique ID for a mouse for the time it is connected to the system,
// and is never reused for the lifetime of the application.
//
// If the mouse is disconnected and reconnected, it will get a new ID.
//
// The value 0 is an invalid ID.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MouseID
type MouseID uint32

// The structure used to identify an SDL cursor.
//
// This is opaque data.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Cursor
type Cursor C.struct_SDL_Cursor

// Cursor types for SDL_CreateSystemCursor().
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SystemCursor
type SystemCursor uint32

const (
	SystemCursorDefault    SystemCursor = iota // Default cursor. Usually an arrow.
	SystemCursorText                           // Text selection. Usually an I-beam.
	SystemCursorWait                           // Wait. Usually an hourglass or watch or spinning ball.
	SystemCursorCrosshair                      // Crosshair.
	SystemCursorProgress                       // Program is busy but still interactive. Usually it's WAIT with an arrow.
	SystemCursorNWSEResize                     // Double arrow pointing northwest and southeast.
	SystemCursorNESWResize                     // Double arrow pointing northeast and southwest.
	SystemCursorEWResize                       // Double arrow pointing west and east.
	SystemCursorNSResize                       // Double arrow pointing north and south.
	SystemCursorMove                           // Four pointed arrow pointing north, south, east, and west.
	SystemCursorNotAllowed                     // Not permitted. Usually a slashed circle or crossbones.
	SystemCursorPointer                        // Pointer that indicates a link. Usually a pointing hand.
	SystemCursorNWResize                       // Window resize top-left. This may be a single arrow or a double arrow like NWSE_RESIZE.
	SystemCursorNResize                        // Window resize top. May be NS_RESIZE.
	SystemCursorNEResize                       // Window resize top-right. May be NESW_RESIZE.
	SystemCursorEResize                        // Window resize right. May be EW_RESIZE.
	SystemCursorSEResize                       // Window resize bottom-right. May be NWSE_RESIZE.
	SystemCursorSResize                        // Window resize bottom. May be NS_RESIZE.
	SystemCursorSWResize                       // Window resize bottom-left. May be NESW_RESIZE.
	SystemCursorWResize                        // Window resize left. May be EW_RESIZE.
	SystemCursorCount      = iota
)

// Scroll direction types for the Scroll event
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MouseWheelDirection
type MouseWheelDirection uint32

const (
	MousewheelNormal  MouseWheelDirection = iota // The scroll direction is normal
	MousewheelFlipped                            // The scroll direction is flipped / natural
)

// A bitmask of pressed mouse buttons, as reported by SDL_GetMouseState, etc.
//
// - Button 1: Left mouse button
// - Button 2: Middle mouse button
// - Button 3: Right mouse button
// - Button 4: Side mouse button 1
// - Button 5: Side mouse button 2
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MouseButtonFlags
type MouseButtonFlags uint32

const ButtonLeft = 1
const ButtonMiddle = 2
const ButtonRight = 3
const ButtonX1 = 4
const ButtonX2 = 5

const (
	ButtonLeftMask   MouseButtonFlags = 1 << (ButtonLeft - 1)
	ButtonMiddleMask MouseButtonFlags = 1 << (ButtonMiddle - 1)
	ButtonRightMask  MouseButtonFlags = 1 << (ButtonRight - 1)
	ButtonX1Mask     MouseButtonFlags = 1 << (ButtonX1 - 1)
	ButtonX2Mask     MouseButtonFlags = 1 << (ButtonX2 - 1)
)

// Return whether a mouse is currently connected.
//
// Returns true if a mouse is connected, false otherwise.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasMouse
func HasMouse() bool {
	return (bool)(C.SDL_HasMouse())
}

// Get a list of currently connected mice.
//
// Note that this will include any device or virtual driver that includes
// mouse functionality, including some game controllers, KVM switches, etc.
// You should wait for input from a device before you consider it actively in
// use.
//
// count: a pointer filled in with the number of mice returned, may be
// NULL.
//
// Returns a 0 terminated array of mouse instance IDs or NULL on failure;
// call SDL_GetError() for more information. This should be freed
// with SDL_free() when it is no longer needed.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetMice
func GetMice() ([]MouseID, error) {
	var count C.int
	m := C.SDL_GetMice(&count)
	if m == nil {
		return nil, getError()
	}
	mice := make([]MouseID, count)
	for i, m := range unsafe.Slice(m, count) {
		mice[i] = MouseID(m)
	}
	C.SDL_free(unsafe.Pointer(m))
	return mice, nil
}

// Get the name of a mouse.
//
// This function returns "" if the mouse doesn't have a name.
//
// id: the mouse instance ID.
//
// Returns the name of the selected mouse, or NULL on failure; call
// SDL_GetError() for more information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetMouseNameForID
func (id MouseID) Name() (string, error) {
	name := C.SDL_GetMouseNameForID((C.SDL_MouseID)(id))
	if name == nil {
		return "", getError()
	}
	return C.GoString(name), nil
}

// Get the window which currently has mouse focus.
//
// Returns the window with mouse focus.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetMouseFocus
func GetMouseFocus() *Window {
	return (*Window)(C.SDL_GetMouseFocus())
}

// Query SDL's cache for the synchronous mouse button state and the
// window-relative SDL-cursor position.
//
// This function returns the cached synchronous state as SDL understands it
// from the last pump of the event queue.
//
// To query the platform for immediate asynchronous state, use
// SDL_GetGlobalMouseState.
//
// Passing non-NULL pointers to `x` or `y` will write the destination with
// respective x or y coordinates relative to the focused window.
//
// In Relative Mode, the SDL-cursor's position usually contradicts the
// platform-cursor's position as manually calculated from
// SDL_GetGlobalMouseState() and SDL_GetWindowPosition.
//
// x: a pointer to receive the SDL-cursor's x-position from the focused
// window's top left corner, can be NULL if unused.
//
// y: a pointer to receive the SDL-cursor's y-position from the focused
// window's top left corner, can be NULL if unused.
//
// Returns a 32-bit bitmask of the button state that can be bitwise-compared
// against the SDL_BUTTON_MASK(X) macro.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetMouseState
func GetMouseState() (x, y float32, buttons MouseButtonFlags) {
	buttons = (MouseButtonFlags)(C.SDL_GetMouseState((*C.float)(&x), (*C.float)(&y)))
	return
}

// Query the platform for the asynchronous mouse button state and the
// desktop-relative platform-cursor position.
//
// This function immediately queries the platform for the most recent
// asynchronous state, more costly than retrieving SDL's cached state in
// SDL_GetMouseState().
//
// Passing non-NULL pointers to `x` or `y` will write the destination with
// respective x or y coordinates relative to the desktop.
//
// In Relative Mode, the platform-cursor's position usually contradicts the
// SDL-cursor's position as manually calculated from SDL_GetMouseState() and
// SDL_GetWindowPosition.
//
// This function can be useful if you need to track the mouse outside of a
// specific window and SDL_CaptureMouse() doesn't fit your needs. For example,
// it could be useful if you need to track the mouse while dragging a window,
// where coordinates relative to a window might not be in sync at all times.
//
// x: a pointer to receive the platform-cursor's x-position from the
// desktop's top left corner, can be NULL if unused.
//
// y: a pointer to receive the platform-cursor's y-position from the
// desktop's top left corner, can be NULL if unused.
//
// Returns a 32-bit bitmask of the button state that can be bitwise-compared
// against the SDL_BUTTON_MASK(X) macro.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGlobalMouseState
func GetGlobalMouseState() (x, y float32, buttons MouseButtonFlags) {
	buttons = (MouseButtonFlags)(C.SDL_GetGlobalMouseState((*C.float)(&x), (*C.float)(&y)))
	return
}

// Query SDL's cache for the synchronous mouse button state and accumulated
// mouse delta since last call.
//
// This function returns the cached synchronous state as SDL understands it
// from the last pump of the event queue.
//
// To query the platform for immediate asynchronous state, use
// SDL_GetGlobalMouseState.
//
// Passing non-NULL pointers to `x` or `y` will write the destination with
// respective x or y deltas accumulated since the last call to this function
// (or since event initialization).
//
// This function is useful for reducing overhead by processing relative mouse
// inputs in one go per-frame instead of individually per-event, at the
// expense of losing the order between events within the frame (e.g. quickly
// pressing and releasing a button within the same frame).
//
// x: a pointer to receive the x mouse delta accumulated since last
// call, can be NULL if unused.
//
// y: a pointer to receive the y mouse delta accumulated since last
// call, can be NULL if unused.
//
// Returns a 32-bit bitmask of the button state that can be bitwise-compared
// against the SDL_BUTTON_MASK(X) macro.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRelativeMouseState
func GetRelativeMouseState() (x, y float32, buttons MouseButtonFlags) {
	buttons = (MouseButtonFlags)(C.SDL_GetRelativeMouseState((*C.float)(&x), (*C.float)(&y)))
	return
}

// Move the mouse cursor to the given position within the window.
//
// This function generates a mouse motion event if relative mode is not
// enabled. If relative mode is enabled, you can force mouse events for the
// warp by setting the SDL_HINT_MOUSE_RELATIVE_WARP_MOTION hint.
//
// Note that this function will appear to succeed, but not actually move the
// mouse when used over Microsoft Remote Desktop.
//
// window: the window to move the mouse into, or NULL for the current
// mouse focus.
//
// x: the x coordinate within the window.
//
// y: the y coordinate within the window.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WarpMouseInWindow
func WarpMouseInWindow(window *Window, x float32, y float32) {
	C.SDL_WarpMouseInWindow((*C.SDL_Window)(window), (C.float)(x), (C.float)(y))
}

// Move the mouse to the given position in global screen space.
//
// This function generates a mouse motion event.
//
// A failure of this function usually means that it is unsupported by a
// platform.
//
// Note that this function will appear to succeed, but not actually move the
// mouse when used over Microsoft Remote Desktop.
//
// x: the x coordinate.
//
// y: the y coordinate.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WarpMouseGlobal
func WarpMouseGlobal(x float32, y float32) bool {
	return (bool)(C.SDL_WarpMouseGlobal((C.float)(x), (C.float)(y)))
}

// Set relative mouse mode for a window.
//
// While the window has focus and relative mouse mode is enabled, the cursor
// is hidden, the mouse position is constrained to the window, and SDL will
// report continuous relative mouse motion even if the mouse is at the edge of
// the window.
//
// If you'd like to keep the mouse position fixed while in relative mode you
// can use SDL_SetWindowMouseRect(). If you'd like the cursor to be at a
// specific location when relative mode ends, you should use
// SDL_WarpMouseInWindow() before disabling relative mode.
//
// This function will flush any pending mouse motion for this window.
//
// window: the window to change.
//
// enabled: true to enable relative mode, false to disable.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetWindowRelativeMouseMode
func (window *Window) SetRelativeMouseMode(enabled bool) error {
	if !C.SDL_SetWindowRelativeMouseMode((*C.SDL_Window)(window), (C.bool)(enabled)) {
		return getError()
	}
	return nil
}

// Query whether relative mouse mode is enabled for a window.
//
// window: the window to query.
//
// Returns true if relative mode is enabled for a window or false otherwise.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowRelativeMouseMode
func (window *Window) RelativeMouseMode() bool {
	return (bool)(C.SDL_GetWindowRelativeMouseMode((*C.SDL_Window)(window)))
}

// Capture the mouse and to track input outside an SDL window.
//
// Capturing enables your app to obtain mouse events globally, instead of just
// within your window. Not all video targets support this function. When
// capturing is enabled, the current window will get all mouse events, but
// unlike relative mode, no change is made to the cursor and it is not
// restrained to your window.
//
// This function may also deny mouse input to other windows--both those in
// your application and others on the system--so you should use this function
// sparingly, and in small bursts. For example, you might want to track the
// mouse while the user is dragging something, until the user releases a mouse
// button. It is not recommended that you capture the mouse for long periods
// of time, such as the entire time your app is running. For that, you should
// probably use SDL_SetWindowRelativeMouseMode() or SDL_SetWindowMouseGrab(),
// depending on your goals.
//
// While captured, mouse events still report coordinates relative to the
// current (foreground) window, but those coordinates may be outside the
// bounds of the window (including negative values). Capturing is only allowed
// for the foreground window. If the window loses focus while capturing, the
// capture will be disabled automatically.
//
// While capturing is enabled, the current window will have the
// `SDL_WINDOW_MOUSE_CAPTURE` flag set.
//
// Please note that SDL will attempt to "auto capture" the mouse while the
// user is pressing a button; this is to try and make mouse behavior more
// consistent between platforms, and deal with the common case of a user
// dragging the mouse outside of the window. This means that if you are
// calling SDL_CaptureMouse() only to deal with this situation, you do not
// have to (although it is safe to do so). If this causes problems for your
// app, you can disable auto capture by setting the
// `SDL_HINT_MOUSE_AUTO_CAPTURE` hint to zero.
//
// enabled: true to enable capturing, false to disable.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CaptureMouse
func CaptureMouse(enabled bool) error {
	if !C.SDL_CaptureMouse((C.bool)(enabled)) {
		return getError()
	}
	return nil
}

// Create a cursor using the specified bitmap data and mask (in MSB format).
//
// `mask` has to be in MSB (Most Significant Bit) format.
//
// The cursor width (`w`) must be a multiple of 8 bits.
//
// The cursor is created in black and white according to the following:
//
// - data=0, mask=1: white
// - data=1, mask=1: black
// - data=0, mask=0: transparent
// - data=1, mask=0: inverted color if possible, black if not.
//
// Cursors created with this function must be freed with SDL_DestroyCursor().
//
// If you want to have a color cursor, or create your cursor from an
// SDL_Surface, you should use SDL_CreateColorCursor(). Alternately, you can
// hide the cursor and draw your own as part of your game's rendering, but it
// will be bound to the framerate.
//
// Also, SDL_CreateSystemCursor() is available, which provides several
// readily-available system cursors to pick from.
//
// data: the color value for each pixel of the cursor.
//
// mask: the mask value for each pixel of the cursor.
//
// w: the width of the cursor.
//
// h: the height of the cursor.
//
// hot_x: the x-axis offset from the left of the cursor image to the
// mouse x position, in the range of 0 to `w` - 1.
//
// hot_y: the y-axis offset from the top of the cursor image to the
// mouse y position, in the range of 0 to `h` - 1.
//
// Returns a new cursor with the specified parameters on success or NULL on
// failure; call SDL_GetError() for more information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateCursor
func CreateCursor(data []byte, mask []byte, w int, h int, hotX int, hotY int) (*Cursor, error) {
	c := (*Cursor)(C.SDL_CreateCursor((*C.Uint8)(unsafe.SliceData(data)), (*C.Uint8)(unsafe.SliceData(mask)), (C.int)(w), (C.int)(h), (C.int)(hotX), (C.int)(hotY)))
	if c == nil {
		return nil, getError()
	}
	return c, nil
}

// Create a color cursor.
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
// surface: an SDL_Surface structure representing the cursor image.
//
// hot_x: the x position of the cursor hot spot.
//
// hot_y: the y position of the cursor hot spot.
//
// Returns the new cursor on success or NULL on failure; call SDL_GetError()
// for more information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateColorCursor
func CreateColorCursor(surface *Surface, hotX int, hotY int) (*Cursor, error) {
	c := (*Cursor)(C.SDL_CreateColorCursor(surface.internal, (C.int)(hotX), (C.int)(hotY)))
	if c == nil {
		return nil, getError()
	}
	return c, nil
}

// Create a system cursor.
//
// id: an SDL_SystemCursor enum value.
//
// Returns a cursor on success or NULL on failure; call SDL_GetError() for
// more information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateSystemCursor
func CreateSystemCursor(id SystemCursor) (*Cursor, error) {
	c := (*Cursor)(C.SDL_CreateSystemCursor((C.SDL_SystemCursor)(id)))
	if c == nil {
		return nil, getError()
	}
	return c, nil
}

// Set the active cursor.
//
// This function sets the currently active cursor to the specified one. If the
// cursor is currently visible, the change will be immediately represented on
// the display. SDL_SetCursor(NULL) can be used to force cursor redraw, if
// this is desired for any reason.
//
// cursor: a cursor to make active.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetCursor
func SetCursor(cursor *Cursor) error {
	if !C.SDL_SetCursor((*C.SDL_Cursor)(cursor)) {
		return getError()
	}
	return nil
}

// Get the active cursor.
//
// This function returns a pointer to the current cursor which is owned by the
// library. It is not necessary to free the cursor with SDL_DestroyCursor().
//
// Returns the active cursor or NULL if there is no mouse.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetCursor
func GetCursor() *Cursor {
	return (*Cursor)(C.SDL_GetCursor())
}

// Get the default cursor.
//
// You do not have to call SDL_DestroyCursor() on the return value, but it is
// safe to do so.
//
// Returns the default cursor on success or NULL on failuree; call
// SDL_GetError() for more information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetDefaultCursor
func GetDefaultCursor() (*Cursor, error) {
	c := (*Cursor)(C.SDL_GetDefaultCursor())
	if c == nil {
		return nil, getError()
	}
	return c, nil
}

// Free a previously-created cursor.
//
// Use this function to free cursor resources created with SDL_CreateCursor(),
// SDL_CreateColorCursor() or SDL_CreateSystemCursor().
//
// cursor: the cursor to free.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DestroyCursor
func (cursor *Cursor) Destroy() {
	C.SDL_DestroyCursor((*C.SDL_Cursor)(cursor))
}

// Show the cursor.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ShowCursor
func ShowCursor() error {
	if !C.SDL_ShowCursor() {
		return getError()
	}
	return nil
}

// Hide the cursor.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HideCursor
func HideCursor() error {
	if !C.SDL_HideCursor() {
		return getError()
	}
	return nil
}

// Return whether the cursor is currently being shown.
//
// Returns `true` if the cursor is being shown, or `false` if the cursor is
// hidden.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CursorVisible
func CursorVisible() bool {
	return (bool)(C.SDL_CursorVisible())
}
