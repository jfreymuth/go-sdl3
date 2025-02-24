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
// #include <SDL3/SDL.h>
import "C"
import (
	"runtime"
	"unsafe"
)

// # CategoryMessagebox
//
// SDL offers a simple message box API, which is useful for simple alerts,
// such as informing the user when something fatal happens at startup without
// the need to build a UI for it (or informing the user _before_ your UI is
// ready).
//
// These message boxes are native system dialogs where possible.
//
// There is both a customizable function ([ShowMessageBox]) that offers
// lots of options for what to display and reports on what choice the user
// made, and also a much-simplified version ([ShowSimpleMessageBox]),
// merely takes a text message and title, and waits until the user presses a
// single "OK" UI button. Often, this is all that is necessary.

// Message box flags.
//
// If supported will display warning icon, etc.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MessageBoxFlags
type MessageBoxFlags uint32

const (
	MessageboxError              MessageBoxFlags = 0x00000010 //error dialog
	MessageboxWarning            MessageBoxFlags = 0x00000020 //warning dialog
	MessageboxInformation        MessageBoxFlags = 0x00000040 //informational dialog
	MessageboxButtonsLeftToRight MessageBoxFlags = 0x00000080 //buttons placed left to right
	MessageboxButtonsRightToLeft MessageBoxFlags = 0x00000100 //buttons placed right to left
)

// [MessageBoxButtonData] flags.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MessageBoxButtonFlags
type MessageBoxButtonFlags uint32

const (
	MessageboxButtonReturnkeyDefault MessageBoxButtonFlags = 0x00000001 //Marks the default button when return is hit
	MessageboxButtonEscapekeyDefault MessageBoxButtonFlags = 0x00000002 //Marks the default button when escape is hit
)

// Individual button data.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MessageBoxButtonData
type MessageBoxButtonData struct {
	Flags    MessageBoxButtonFlags
	ButtonID int    // User defined button id (value returned via [ShowMessageBox])
	Text     string // The UTF-8 button text
}

// RGB value used in a message box color scheme
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MessageBoxColor
type MessageBoxColor struct {
	R, G, B byte
}

// An enumeration of indices inside the colors array of
// [MessageBoxColorScheme].
//
// https://wiki.libsdl.org/SDL3/SDL_MessageBoxColorType
type MessageBoxColorType uint32

const (
	MessageboxColorBackground MessageBoxColorType = iota
	MessageboxColorText
	MessageboxColorButtonBorder
	MessageboxColorButtonBackground
	MessageboxColorButtonSelected
	MessageboxColorCount = iota // Size of the colors array of [MessageBoxColorScheme].
)

// A set of colors to use for message box dialogs
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MessageBoxColorScheme
type MessageBoxColorScheme struct {
	Colors [MessageboxColorCount]MessageBoxColor
}

// MessageBox structure containing title, text, window, etc.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MessageBoxData
type MessageBoxData struct {
	Flags       MessageBoxFlags
	Window      *Window                // Parent window, can be nil
	Title       string                 // UTF-8 title
	Message     string                 // UTF-8 message text
	Buttons     []MessageBoxButtonData //
	ColorScheme MessageBoxColorScheme  // [MessageBoxColorScheme], can be NULL to use system settings
}

// Create a modal message box.
//
// If your needs aren't complex, it might be easier to use
// [ShowSimpleMessageBox].
//
// This function should be called on the thread that created the parent
// window, or on the main thread if the messagebox has no parent. It will
// block execution of that thread until the user clicks a button or closes the
// messagebox.
//
// This function may be called at any time, even before [Init]. This makes
// it useful for reporting errors like a failure to create a renderer or
// OpenGL context.
//
// On X11, SDL rolls its own dialog box with X11 primitives instead of a
// formal toolkit like GTK+ or Qt.
//
// Note that if [Init] would fail because there isn't any available video
// target, this function is likely to fail for the same reasons. If this is a
// concern, check the return value from this function and fall back to writing
// to stderr if you can.
//
// messageboxdata: the [MessageBoxData] structure with title, text and
// other options.
//
// Returns the selected ButtonID on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MessageBoxButtonFlags
func ShowMessageBox(messageboxdata *MessageBoxData) (int, error) {
	var pinner runtime.Pinner
	buttons := make([]C.SDL_MessageBoxButtonData, len(messageboxdata.Buttons))
	for i, b := range messageboxdata.Buttons {
		buttons[i].flags = C.SDL_MessageBoxButtonFlags(b.Flags)
		buttons[i].buttonID = C.int(b.ButtonID)
		buttons[i].text = tmpstring(b.Text)
		pinner.Pin(buttons[i].text)
	}
	var colors [MessageboxColorCount]C.SDL_MessageBoxColor
	for i, c := range messageboxdata.ColorScheme.Colors {
		colors[i].r = C.Uint8(c.R)
		colors[i].g = C.Uint8(c.G)
		colors[i].b = C.Uint8(c.B)
	}
	data := &C.SDL_MessageBoxData{
		C.SDL_MessageBoxFlags(messageboxdata.Flags),
		(*C.SDL_Window)(messageboxdata.Window),
		tmpstring(messageboxdata.Title),
		tmpstring(messageboxdata.Message),
		C.int(len(buttons)),
		unsafe.SliceData(buttons),
		&C.SDL_MessageBoxColorScheme{colors},
	}
	pinner.Pin(data.title)
	pinner.Pin(data.message)
	pinner.Pin(data.buttons)
	pinner.Pin(data.colorScheme)
	var buttonID C.int
	ok := C.SDL_ShowMessageBox(data, &buttonID)
	pinner.Unpin()
	if !ok {
		return 0, getError()
	}
	return int(buttonID), nil
}

// Display a simple modal message box.
//
// If your needs aren't complex, this function is preferred over
// [ShowMessageBox].
//
// flags may be any of the following:
//
//   - [MessageboxError]: error dialog
//   - [MessageboxWarning]: warning dialog
//   - [MessageboxInformation]: informational dialog
//
// This function should be called on the thread that created the parent
// window, or on the main thread if the messagebox has no parent. It will
// block execution of that thread until the user clicks a button or closes the
// messagebox.
//
// This function may be called at any time, even before [Init]. This makes
// it useful for reporting errors like a failure to create a renderer or
// OpenGL context.
//
// On X11, SDL rolls its own dialog box with X11 primitives instead of a
// formal toolkit like GTK+ or Qt.
//
// Note that if [Init] would fail because there isn't any available video
// target, this function is likely to fail for the same reasons. If this is a
// concern, check the return value from this function and fall back to writing
// to stderr if you can.
//
// flags: a [MessageBoxFlags] value.
//
// title: UTF-8 title text.
//
// message: UTF-8 message text.
//
// window: the parent window, or nil for no parent.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ShowSimpleMessageBox
func ShowSimpleMessageBox(flags MessageBoxFlags, title string, message string, window *Window) error {
	ok := C.SDL_ShowSimpleMessageBox((C.SDL_MessageBoxFlags)(flags), tmpstring(title), tmpstring(message), (*C.SDL_Window)(window))
	if !ok {
		return getError()
	}
	return nil
}
