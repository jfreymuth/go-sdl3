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
// #cgo noescape SDL_GUIDToString
// #cgo nocallback SDL_GUIDToString
// #cgo noescape SDL_StringToGUID
// #cgo nocallback SDL_StringToGUID
// #include <SDL3/SDL.h>
import "C"

// # CategoryGUID
//
// A GUID is a 128-bit value that represents something that is uniquely
// identifiable by this value: "globally unique."
//
// SDL provides functions to convert a GUID to/from a string.

// A GUID is a 128-bit identifier for an input device that identifies
// that device across runs of SDL programs on the same platform.
//
// If the device is detached and then re-attached to a different port, or if
// the base system is rebooted, the device should still report the same GUID.
//
// GUIDs are as precise as possible but are not guaranteed to distinguish
// physically distinct but equivalent devices. For example, two game
// controllers from the same vendor with the same product ID and revision may
// have the same GUID.
//
// GUIDs may be platform-dependent (i.e., the same device may report different
// GUIDs on different operating systems).
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GUID
type GUID [16]C.Uint8

// Get an ASCII string representation for a given GUID.
//
// guid: the GUID you wish to convert to string.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GUIDToString
func (g GUID) String() string {
	var buf [33]C.char
	C.SDL_GUIDToString(C.SDL_GUID{g}, &buf[0], C.int(len(buf)))
	return C.GoString(&buf[0])
}

// Convert a GUID string into a [GUID] structure.
//
// Performs no error checking. If this function is given a string containing
// an invalid GUID, the function will silently succeed, but the GUID generated
// will not be useful.
//
// s: string containing an ASCII representation of a GUID.
//
// Returns a GUID structure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_StringToGUID
func StringToGUID(s string) GUID {
	return (GUID)(C.SDL_StringToGUID(tmpstring(s)).data)
}
