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
// #cgo noescape SDL_HasKeyboard
// #cgo nocallback SDL_HasKeyboard
// #cgo noescape SDL_GetKeyboards
// #cgo nocallback SDL_GetKeyboards
// #cgo noescape SDL_GetKeyboardNameForID
// #cgo nocallback SDL_GetKeyboardNameForID
// #cgo noescape SDL_GetKeyboardFocus
// #cgo nocallback SDL_GetKeyboardFocus
// #cgo noescape SDL_GetKeyboardState
// #cgo nocallback SDL_GetKeyboardState
// #cgo noescape SDL_ResetKeyboard
// #cgo nocallback SDL_ResetKeyboard
// #cgo noescape SDL_GetModState
// #cgo nocallback SDL_GetModState
// #cgo noescape SDL_SetModState
// #cgo nocallback SDL_SetModState
// #cgo noescape SDL_GetKeyFromScancode
// #cgo nocallback SDL_GetKeyFromScancode
// #cgo noescape SDL_GetScancodeFromKey
// #cgo nocallback SDL_GetScancodeFromKey
// #cgo noescape SDL_SetScancodeName
// #cgo nocallback SDL_SetScancodeName
// #cgo noescape SDL_GetScancodeName
// #cgo nocallback SDL_GetScancodeName
// #cgo noescape SDL_GetScancodeFromName
// #cgo nocallback SDL_GetScancodeFromName
// #cgo noescape SDL_GetKeyName
// #cgo nocallback SDL_GetKeyName
// #cgo noescape SDL_GetKeyFromName
// #cgo nocallback SDL_GetKeyFromName
// #cgo noescape SDL_StartTextInput
// #cgo nocallback SDL_StartTextInput
// #cgo noescape SDL_StartTextInputWithProperties
// #cgo nocallback SDL_StartTextInputWithProperties
// #cgo noescape SDL_TextInputActive
// #cgo nocallback SDL_TextInputActive
// #cgo noescape SDL_StopTextInput
// #cgo nocallback SDL_StopTextInput
// #cgo noescape SDL_ClearComposition
// #cgo nocallback SDL_ClearComposition
// #cgo noescape SDL_SetTextInputArea
// #cgo nocallback SDL_SetTextInputArea
// #cgo noescape SDL_GetTextInputArea
// #cgo nocallback SDL_GetTextInputArea
// #cgo noescape SDL_HasScreenKeyboardSupport
// #cgo nocallback SDL_HasScreenKeyboardSupport
// #cgo noescape SDL_ScreenKeyboardShown
// #cgo nocallback SDL_ScreenKeyboardShown
// #include <SDL3/SDL.h>
import "C"
import "unsafe"

// # CategoryKeyboard
//
// SDL keyboard management.
//
// Please refer to the Best Keyboard Practices document for details on how
// best to accept keyboard input in various types of programs:
//
// https://wiki.libsdl.org/SDL3/BestKeyboardPractices

// This is a unique ID for a keyboard for the time it is connected to the
// system, and is never reused for the lifetime of the application.
//
// If the keyboard is disconnected and reconnected, it will get a new ID.
//
// The value 0 is an invalid ID.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_KeyboardID
type KeyboardID uint32

// Return whether a keyboard is currently connected.
//
// Returns true if a keyboard is connected, false otherwise.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasKeyboard
func HasKeyboard() bool {
	return (bool)(C.SDL_HasKeyboard())
}

// Get a list of currently connected keyboards.
//
// Note that this will include any device or virtual driver that includes
// keyboard functionality, including some mice, KVM switches, motherboard
// power buttons, etc. You should wait for input from a device before you
// consider it actively in use.
//
// count: a pointer filled in with the number of keyboards returned, may
// be NULL.
//
// Returns a 0 terminated array of keyboards instance IDs or NULL on failure;
// call SDL_GetError() for more information. This should be freed
// with SDL_free() when it is no longer needed.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetKeyboards
func GetKeyboards() ([]KeyboardID, error) {
	var count C.int
	kb := C.SDL_GetKeyboards(&count)
	if kb == nil {
		return nil, getError()
	}
	result := make([]KeyboardID, count)
	for i, k := range unsafe.Slice(kb, count) {
		result[i] = KeyboardID(k)
	}
	C.SDL_free(unsafe.Pointer(kb))
	return result, nil
}

// Get the name of a keyboard.
//
// This function returns "" if the keyboard doesn't have a name.
//
// id: the keyboard instance ID.
//
// Returns the name of the selected keyboard or NULL on failure; call
// SDL_GetError() for more information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetKeyboardNameForID
func (id KeyboardID) Name() (string, error) {
	name := C.SDL_GetKeyboardNameForID((C.SDL_KeyboardID)(id))
	if name == nil {
		return "", getError()
	}
	return C.GoString(name), nil
}

// Query the window which currently has keyboard focus.
//
// Returns the window with keyboard focus.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetKeyboardFocus
func GetKeyboardFocus() *Window {
	return (*Window)(C.SDL_GetKeyboardFocus())
}

// Get a snapshot of the current state of the keyboard.
//
// The pointer returned is a pointer to an internal SDL array. It will be
// valid for the whole lifetime of the application and should not be freed by
// the caller.
//
// A array element with a value of true means that the key is pressed and a
// value of false means that it is not. Indexes into this array are obtained
// by using SDL_Scancode values.
//
// Use SDL_PumpEvents() to update the state array.
//
// This function gives you the current state after all events have been
// processed, so if a key or button has been pressed and released before you
// process events, then the pressed state will never show up in the
// SDL_GetKeyboardState() calls.
//
// Note: This function doesn't take into account whether shift has been
// pressed or not.
//
// numkeys: if non-NULL, receives the length of the returned array.
//
// Returns a pointer to an array of key states.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetKeyboardState
func GetKeyboardState() []bool {
	var num C.int
	state := (*bool)(C.SDL_GetKeyboardState(&num))
	return unsafe.Slice(state, num)
}

// Clear the state of the keyboard.
//
// This function will generate key up events for all pressed keys.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ResetKeyboard
func ResetKeyboard() {
	C.SDL_ResetKeyboard()
}

// Get the current key modifier state for the keyboard.
//
// Returns an OR'd combination of the modifier keys for the keyboard. See
// SDL_Keymod for details.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetModState
func GetModState() Keymod {
	return (Keymod)(C.SDL_GetModState())
}

// Set the current key modifier state for the keyboard.
//
// The inverse of SDL_GetModState(), SDL_SetModState() allows you to impose
// modifier key states on your application. Simply pass your desired modifier
// states into `modstate`. This value may be a bitwise, OR'd combination of
// SDL_Keymod values.
//
// This does not change the keyboard state, only the key modifier flags that
// SDL reports.
//
// modstate: the desired SDL_Keymod for the keyboard.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetModState
func SetModState(modstate Keymod) {
	C.SDL_SetModState((C.SDL_Keymod)(modstate))
}

// Get the key code corresponding to the given scancode according to the
// current keyboard layout.
//
// If you want to get the keycode as it would be delivered in key events,
// including options specified in SDL_HINT_KEYCODE_OPTIONS, then you should
// pass `key_event` as true. Otherwise this function simply translates the
// scancode based on the given modifier state.
//
// scancode: the desired SDL_Scancode to query.
//
// modstate: the modifier state to use when translating the scancode to
// a keycode.
//
// key_event: true if the keycode will be used in key events.
//
// Returns the SDL_Keycode that corresponds to the given SDL_Scancode.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetKeyFromScancode
func GetKeyFromScancode(scancode Scancode, modstate Keymod, keyEvent bool) Keycode {
	return (Keycode)(C.SDL_GetKeyFromScancode((C.SDL_Scancode)(scancode), (C.SDL_Keymod)(modstate), (C.bool)(keyEvent)))
}

// Get the scancode corresponding to the given key code according to the
// current keyboard layout.
//
// Note that there may be multiple scancode+modifier states that can generate
// this keycode, this will just return the first one found.
//
// key: the desired SDL_Keycode to query.
//
// modstate: a pointer to the modifier state that would be used when the
// scancode generates this key, may be NULL.
//
// Returns the SDL_Scancode that corresponds to the given SDL_Keycode.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetScancodeFromKey
func GetScancodeFromKey(key Keycode) (Scancode, Keymod) {
	var mod C.SDL_Keymod
	return (Scancode)(C.SDL_GetScancodeFromKey((C.SDL_Keycode)(key), &mod)), Keymod(mod)
}

// Set a human-readable name for a scancode.
//
// scancode: the desired SDL_Scancode.
//
// name: the name to use for the scancode, encoded as UTF-8. The string
// is not copied, so the pointer given to this function must stay
// valid while SDL is being used.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetScancodeName
func SetScancodeName(scancode Scancode, name string) error {
	if !C.SDL_SetScancodeName((C.SDL_Scancode)(scancode), tmpstring(name)) {
		return getError()
	}
	return nil
}

// Get a human-readable name for a scancode.
//
// **Warning**: The returned name is by design not stable across platforms,
// e.g. the name for `SDL_SCANCODE_LGUI` is "Left GUI" under Linux but "Left
// Windows" under Microsoft Windows, and some scancodes like
// `SDL_SCANCODE_NONUSBACKSLASH` don't have any name at all. There are even
// scancodes that share names, e.g. `SDL_SCANCODE_RETURN` and
// `SDL_SCANCODE_RETURN2` (both called "Return"). This function is therefore
// unsuitable for creating a stable cross-platform two-way mapping between
// strings and scancodes.
//
// scancode: the desired SDL_Scancode to query.
//
// Returns a pointer to the name for the scancode. If the scancode doesn't
// have a name this function returns an empty string ("").
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetScancodeName
func GetScancodeName(scancode Scancode) string {
	return C.GoString(C.SDL_GetScancodeName((C.SDL_Scancode)(scancode)))
}

// Get a scancode from a human-readable name.
//
// name: the human-readable scancode name.
//
// Returns the SDL_Scancode, or `SDL_SCANCODE_UNKNOWN` if the name wasn't
// recognized; call SDL_GetError() for more information.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetScancodeFromName
func GetScancodeFromName(name string) (Scancode, error) {
	scancode := (Scancode)(C.SDL_GetScancodeFromName(tmpstring(name)))
	if scancode == ScancodeUnknown {
		return ScancodeUnknown, getError()
	}
	return scancode, nil
}

// Get a human-readable name for a key.
//
// If the key doesn't have a name, this function returns an empty string ("").
//
// Letters will be presented in their uppercase form, if applicable.
//
// key: the desired SDL_Keycode to query.
//
// Returns a UTF-8 encoded string of the key name.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetKeyName
func GetKeyName(key Keycode) string {
	return C.GoString(C.SDL_GetKeyName((C.SDL_Keycode)(key)))
}

// Get a key code from a human-readable name.
//
// name: the human-readable key name.
//
// Returns key code, or `SDLK_UNKNOWN` if the name wasn't recognized; call
// SDL_GetError() for more information.
//
// This function is not thread safe.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetKeyFromName
func GetKeyFromName(name string) (Keycode, error) {
	keycode := (Keycode)(C.SDL_GetKeyFromName(tmpstring(name)))
	if keycode == KeyUnknown {
		return KeyUnknown, getError()
	}
	return keycode, nil
}

// Start accepting Unicode text input events in a window.
//
// This function will enable text input (SDL_EVENT_TEXT_INPUT and
// SDL_EVENT_TEXT_EDITING events) in the specified window. Please use this
// function paired with SDL_StopTextInput().
//
// Text input events are not received by default.
//
// On some platforms using this function shows the screen keyboard and/or
// activates an IME, which can prevent some key press events from being passed
// through.
//
// window: the window to enable text input.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_StartTextInput
func (window *Window) StartTextInput() error {
	if !C.SDL_StartTextInput((*C.SDL_Window)(window)) {
		return getError()
	}
	return nil
}

// Text input type.
//
// These are the valid values for SDL_PROP_TEXTINPUT_TYPE_NUMBER. Not every
// value is valid on every platform, but where a value isn't supported, a
// reasonable fallback will be used.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_TextInputType
type TextInputType uint32

const (
	TextinputTypeText                  TextInputType = iota // The input is text
	TextinputTypeTextName                                   // The input is a person's name
	TextinputTypeTextEmail                                  // The input is an e-mail address
	TextinputTypeTextUsername                               // The input is a username
	TextinputTypeTextPasswordHidden                         // The input is a secure password that is hidden
	TextinputTypeTextPasswordVisible                        // The input is a secure password that is visible
	TextinputTypeNumber                                     // The input is a number
	TextinputTypeNumberPasswordHidden                       // The input is a secure PIN that is hidden
	TextinputTypeNumberPasswordVisible                      // The input is a secure PIN that is visible
)

// Auto capitalization type.
//
// These are the valid values for SDL_PROP_TEXTINPUT_CAPITALIZATION_NUMBER.
// Not every value is valid on every platform, but where a value isn't
// supported, a reasonable fallback will be used.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Capitalization
type Capitalization uint32

const (
	CapitalizeNone      Capitalization = iota // No auto-capitalization will be done
	CapitalizeSentences                       // The first letter of sentences will be capitalized
	CapitalizeWords                           // The first letter of words will be capitalized
	CapitalizeLetters                         // All letters will be capitalized
)

// Start accepting Unicode text input events in a window, with properties
// describing the input.
//
// This function will enable text input (SDL_EVENT_TEXT_INPUT and
// SDL_EVENT_TEXT_EDITING events) in the specified window. Please use this
// function paired with SDL_StopTextInput().
//
// Text input events are not received by default.
//
// On some platforms using this function shows the screen keyboard and/or
// activates an IME, which can prevent some key press events from being passed
// through.
//
// These are the supported properties:
//
// - `SDL_PROP_TEXTINPUT_TYPE_NUMBER` - an SDL_TextInputType value that
// describes text being input, defaults to SDL_TEXTINPUT_TYPE_TEXT.
// - `SDL_PROP_TEXTINPUT_CAPITALIZATION_NUMBER` - an SDL_Capitalization value
// that describes how text should be capitalized, defaults to
// SDL_CAPITALIZE_SENTENCES for normal text entry, SDL_CAPITALIZE_WORDS for
// SDL_TEXTINPUT_TYPE_TEXT_NAME, and SDL_CAPITALIZE_NONE for e-mail
// addresses, usernames, and passwords.
// - `SDL_PROP_TEXTINPUT_AUTOCORRECT_BOOLEAN` - true to enable auto completion
// and auto correction, defaults to true.
// - `SDL_PROP_TEXTINPUT_MULTILINE_BOOLEAN` - true if multiple lines of text
// are allowed. This defaults to true if SDL_HINT_RETURN_KEY_HIDES_IME is
// "0" or is not set, and defaults to false if SDL_HINT_RETURN_KEY_HIDES_IME
// is "1".
//
// On Android you can directly specify the input type:
//
// - `SDL_PROP_TEXTINPUT_ANDROID_INPUTTYPE_NUMBER` - the text input type to
// use, overriding other properties. This is documented at
// https://developer.android.com/reference/android/text/InputType
//
// window: the window to enable text input.
//
// props: the properties to use.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_StartTextInputWithProperties
func (window *Window) StartTextInputWithProperties(props PropertiesID) error {
	if !C.SDL_StartTextInputWithProperties((*C.SDL_Window)(window), (C.SDL_PropertiesID)(props)) {
		return getError()
	}
	return nil
}

const PropTextinputTypeNumber = "SDL.textinput.type"
const PropTextinputCapitalizationNumber = "SDL.textinput.capitalization"
const PropTextinputAutocorrectBoolean = "SDL.textinput.autocorrect"
const PropTextinputMultilineBoolean = "SDL.textinput.multiline"
const PropTextinputAndroidInputtypeNumber = "SDL.textinput.android.inputtype"

// Check whether or not Unicode text input events are enabled for a window.
//
// window: the window to check.
//
// Returns true if text input events are enabled else false.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_TextInputActive
func (window *Window) TextInputActive() bool {
	return (bool)(C.SDL_TextInputActive((*C.SDL_Window)(window)))
}

// Stop receiving any text input events in a window.
//
// If SDL_StartTextInput() showed the screen keyboard, this function will hide
// it.
//
// window: the window to disable text input.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_StopTextInput
func (window *Window) StopTextInput() error {
	if !C.SDL_StopTextInput((*C.SDL_Window)(window)) {
		return getError()
	}
	return nil
}

// Dismiss the composition window/IME without disabling the subsystem.
//
// window: the window to affect.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ClearComposition
func (window *Window) ClearComposition() error {
	if !C.SDL_ClearComposition((*C.SDL_Window)(window)) {
		return getError()
	}
	return nil
}

// Set the area used to type Unicode text input.
//
// Native input methods may place a window with word suggestions near the
// cursor, without covering the text being entered.
//
// window: the window for which to set the text input area.
//
// rect: the SDL_Rect representing the text input area, in window
// coordinates, or NULL to clear it.
//
// cursor: the offset of the current cursor location relative to
// `rect->x`, in window coordinates.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetTextInputArea
func (window *Window) SetTextInputArea(rect *Rect, cursor int) error {
	var crect *C.SDL_Rect
	if rect != nil {
		crect = &C.SDL_Rect{C.int(rect.X), C.int(rect.Y), C.int(rect.W), C.int(rect.H)}
	}
	if !C.SDL_SetTextInputArea((*C.SDL_Window)(window), crect, (C.int)(cursor)) {
		return getError()
	}
	return nil
}

// Get the area used to type Unicode text input.
//
// This returns the values previously set by SDL_SetTextInputArea().
//
// window: the window for which to query the text input area.
//
// rect: a pointer to an SDL_Rect filled in with the text input area,
// may be NULL.
//
// cursor: a pointer to the offset of the current cursor location
// relative to `rect->x`, may be NULL.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTextInputArea
func (window *Window) TextInputArea() (rect Rect, cursor int, err error) {
	var crect C.SDL_Rect
	var ccursor C.int
	if !C.SDL_GetTextInputArea((*C.SDL_Window)(window), &crect, &ccursor) {
		return Rect{}, 0, getError()
	}
	return Rect{int(crect.x), int(crect.y), int(crect.w), int(crect.h)}, int(ccursor), nil
}

// Check whether the platform has screen keyboard support.
//
// Returns true if the platform has some screen keyboard support or false if
// not.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasScreenKeyboardSupport
func HasScreenKeyboardSupport() bool {
	return (bool)(C.SDL_HasScreenKeyboardSupport())
}

// Check whether the screen keyboard is shown for given window.
//
// window: the window for which screen keyboard should be queried.
//
// Returns true if screen keyboard is shown or false if not.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ScreenKeyboardShown
func (window *Window) ScreenKeyboardShown() bool {
	return (bool)(C.SDL_ScreenKeyboardShown((*C.SDL_Window)(window)))
}
