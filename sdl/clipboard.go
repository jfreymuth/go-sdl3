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
// #cgo noescape SDL_SetClipboardText
// #cgo noescape SDL_GetClipboardText
// #cgo noescape SDL_HasClipboardText
// #cgo nocallback SDL_HasClipboardText
// #cgo noescape SDL_SetPrimarySelectionText
// #cgo nocallback SDL_SetPrimarySelectionText
// #cgo noescape SDL_GetPrimarySelectionText
// #cgo nocallback SDL_GetPrimarySelectionText
// #cgo noescape SDL_HasPrimarySelectionText
// #cgo nocallback SDL_HasPrimarySelectionText
// #cgo noescape wrap_SDL_SetClipboardData
// #cgo noescape SDL_ClearClipboardData
// #cgo noescape SDL_GetClipboardData
// #cgo noescape SDL_HasClipboardData
// #cgo nocallback SDL_HasClipboardData
// #cgo noescape SDL_GetClipboardMimeTypes
// #cgo nocallback SDL_GetClipboardMimeTypes
// #include <SDL3/SDL.h>
// extern bool wrap_SDL_SetClipboardData(uintptr_t userdata, const char **mime_types, size_t num_mime_types);
import "C"
import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

// # CategoryClipboard
//
// SDL provides access to the system clipboard, both for reading information
// from other processes and publishing information of its own.
//
// This is not just text! SDL apps can access and publish data by mimetype.
//
// ## Basic use (text)
//
// Obtaining and publishing simple text to the system clipboard is as easy as
// calling SDL_GetClipboardText() and SDL_SetClipboardText(), respectively.
// These deal with C strings in UTF-8 encoding. Data transmission and encoding
// conversion is completely managed by SDL.
//
// ## Clipboard callbacks (data other than text)
//
// Things get more complicated when the clipboard contains something other
// than text. Not only can the system clipboard contain data of any type, in
// some cases it can contain the same data in different formats! For example,
// an image painting app might let the user copy a graphic to the clipboard,
// and offers it in .BMP, .JPG, or .PNG format for other apps to consume.
//
// Obtaining clipboard data ("pasting") like this is a matter of calling
// SDL_GetClipboardData() and telling it the mimetype of the data you want.
// But how does one know if that format is available? SDL_HasClipboardData()
// can report if a specific mimetype is offered, and
// SDL_GetClipboardMimeTypes() can provide the entire list of mimetypes
// available, so the app can decide what to do with the data and what formats
// it can support.
//
// Setting the clipboard ("copying") to arbitrary data is done with
// SDL_SetClipboardData. The app does not provide the data in this call, but
// rather the mimetypes it is willing to provide and a callback function.
// During the callback, the app will generate the data. This allows massive
// data sets to be provided to the clipboard, without any data being copied
// before it is explicitly requested. More specifically, it allows an app to
// offer data in multiple formats without providing a copy of all of them
// upfront. If the app has an image that it could provide in PNG or JPG
// format, it doesn't have to encode it to either of those unless and until
// something tries to paste it.
//
// ## Primary Selection
//
// The X11 and Wayland video targets have a concept of the "primary selection"
// in addition to the usual clipboard. This is generally highlighted (but not
// explicitly copied) text from various apps. SDL offers APIs for this through
// SDL_GetPrimarySelectionText() and SDL_SetPrimarySelectionText(). SDL offers
// these APIs on platforms without this concept, too, but only so far that it
// will keep a copy of a string that the app sets for later retrieval; the
// operating system will not ever attempt to change the string externally if
// it doesn't support a primary selection.

// Put UTF-8 text into the clipboard.
//
// text: the text to store in the clipboard.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetClipboardText
func SetClipboardText(text string) error {
	if !C.SDL_SetClipboardText(tmpstring(text)) {
		return getError()
	}
	return nil
}

// Get UTF-8 text from the clipboard.
//
// This functions returns an empty string if there was not enough memory left
// for a copy of the clipboard's content.
//
// Returns the clipboard text on success or an empty string on failure; call
// SDL_GetError() for more information. This should be freed with
// SDL_free() when it is no longer needed.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetClipboardText
func GetClipboardText() (string, error) {
	ctext := C.SDL_GetClipboardText()
	text := C.GoString(ctext)
	C.SDL_free(unsafe.Pointer(ctext))
	if text == "" {
		return "", getError()
	}
	return text, nil
}

// Query whether the clipboard exists and contains a non-empty text string.
//
// Returns true if the clipboard has text, or false if it does not.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasClipboardText
func HasClipboardText() bool {
	return (bool)(C.SDL_HasClipboardText())
}

// Put UTF-8 text into the primary selection.
//
// text: the text to store in the primary selection.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetPrimarySelectionText
func SetPrimarySelectionText(text string) error {
	if !C.SDL_SetPrimarySelectionText(tmpstring(text)) {
		return getError()
	}
	return nil
}

// Get UTF-8 text from the primary selection.
//
// This functions returns an empty string if there was not enough memory left
// for a copy of the primary selection's content.
//
// Returns the primary selection text on success or an empty string on
// failure; call SDL_GetError() for more information. This should be
// freed with SDL_free() when it is no longer needed.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetPrimarySelectionText
func GetPrimarySelectionText() (string, error) {
	ctext := C.SDL_GetPrimarySelectionText()
	text := C.GoString(ctext)
	C.SDL_free(unsafe.Pointer(ctext))
	if text == "" {
		return "", getError()
	}
	return text, nil
}

// Query whether the primary selection exists and contains a non-empty text
// string.
//
// Returns true if the primary selection has text, or false if it does not.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasPrimarySelectionText
func HasPrimarySelectionText() bool {
	return (bool)(C.SDL_HasPrimarySelectionText())
}

// Callback function that will be called when data for the specified mime-type
// is requested by the OS.
//
// The callback function is called with NULL as the mime_type when the
// clipboard is cleared or new data is set. The clipboard is automatically
// cleared in SDL_Quit().
//
// userdata: a pointer to provided user data.
//
// mime_type: the requested mime-type.
//
// size: a pointer filled in with the length of the returned data.
//
// Returns a pointer to the data for the provided mime-type. Returning NULL
// or setting length to 0 will cause no data to be sent to the
// "receiver". It is up to the receiver to handle this. Essentially
// returning no data is more or less undefined behavior and may cause
// breakage in receiving applications. The returned data will not be
// freed so it needs to be retained and dealt with internally.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ClipboardDataCallback
type ClipboardDataCallback func(mimeType string) string

var clipboardData runtime.Pinner

//export cb_ClipboardDataCallback
func cb_ClipboardDataCallback(userdata uintptr, mimeType *C.char, size *C.size_t) unsafe.Pointer {
	h := cgo.Handle(userdata)
	data := h.Value().(ClipboardDataCallback)(C.GoString(mimeType))
	*size = C.size_t(len(data))
	ptr := unsafe.Pointer(unsafe.StringData(data))
	clipboardData.Pin(ptr)
	return ptr
}

//export cb_ClipboardCleanupCallback
func cb_ClipboardCleanupCallback(userdata uintptr) {
	h := cgo.Handle(userdata)
	h.Delete()
	clipboardData.Unpin()
}

// Offer clipboard data to the OS.
//
// Tell the operating system that the application is offering clipboard data
// for each of the provided mime-types. Once another application requests the
// data the callback function will be called, allowing it to generate and
// respond with the data for the requested mime-type.
//
// The size of text data does not include any terminator, and the text does
// not need to be null terminated (e.g. you can directly copy a portion of a
// document).
//
// callback: a function pointer to the function that provides the
// clipboard data.
//
// cleanup: a function pointer to the function that cleans up the
// clipboard data.
//
// userdata: an opaque pointer that will be forwarded to the callbacks.
//
// mime_types: a list of mime-types that are being offered.
//
// num_mime_types: the number of mime-types in the mime_types list.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetClipboardData
func SetClipboardData(callback ClipboardDataCallback, mimeTypes []string) error {
	h := cgo.NewHandle(callback)
	cmime := make([]*C.char, len(mimeTypes))
	var pinner runtime.Pinner
	for i, m := range mimeTypes {
		cmime[i] = tmpstring(m)
		pinner.Pin(cmime[i])
	}
	ok := C.wrap_SDL_SetClipboardData(C.uintptr_t(h), unsafe.SliceData(cmime), C.size_t(len(mimeTypes)))
	pinner.Unpin()
	if !ok {
		h.Delete()
		return getError()
	}
	return nil
}

// Clear the clipboard data.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ClearClipboardData
func ClearClipboardData() error {
	if !C.SDL_ClearClipboardData() {
		return getError()
	}
	return nil
}

// Get the data from clipboard for a given mime type.
//
// The size of text data does not include the terminator, but the text is
// guaranteed to be null terminated.
//
// mime_type: the mime type to read from the clipboard.
//
// size: a pointer filled in with the length of the returned data.
//
// Returns the retrieved data buffer or NULL on failure; call SDL_GetError()
// for more information. This should be freed with SDL_free() when it
// is no longer needed.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetClipboardData
func GetClipboardData(mimeType string) (string, error) {
	var size C.size_t
	cdata := C.SDL_GetClipboardData(tmpstring(mimeType), &size)
	if cdata == nil {
		return "", getError()
	}
	data := C.GoString((*C.char)(cdata))
	C.SDL_free(cdata)
	return data, nil
}

// Query whether there is data in the clipboard for the provided mime type.
//
// mime_type: the mime type to check for data for.
//
// Returns true if there exists data in clipboard for the provided mime type,
// false if it does not.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasClipboardData
func HasClipboardData(mimeType string) bool {
	return (bool)(C.SDL_HasClipboardData(tmpstring(mimeType)))
}

// Retrieve the list of mime types available in the clipboard.
//
// num_mime_types: a pointer filled with the number of mime types, may
// be NULL.
//
// Returns a null terminated array of strings with mime types, or NULL on
// failure; call SDL_GetError() for more information. This should be
// freed with SDL_free() when it is no longer needed.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetClipboardMimeTypes
func GetClipboardMimeTypes() ([]string, error) {
	var num C.size_t
	result := C.SDL_GetClipboardMimeTypes(&num)
	if result == nil {
		return nil, getError()
	}
	mimeTypes := make([]string, num)
	for i, r := range unsafe.Slice(result, num) {
		mimeTypes[i] = C.GoString(r)
	}
	C.SDL_free(unsafe.Pointer(result))
	return mimeTypes, nil
}
