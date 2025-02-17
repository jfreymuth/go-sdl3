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
// #cgo noescape SDL_GetVersion
// #cgo nocallback SDL_GetVersion
// #cgo noescape SDL_GetRevision
// #cgo nocallback SDL_GetRevision
// #include <SDL3/SDL.h>
import "C"

// # CategoryVersion
//
// Functionality to query the current SDL version, both as headers the app was
// compiled against, and a library the app is linked to.

// The current major version of SDL headers.
//
// If this were SDL version 3.2.1, this value would be 3.
//
// This macro is available since SDL 3.2.0.
const MajorVersion = 3

// The current minor version of the SDL headers.
//
// If this were SDL version 3.2.1, this value would be 2.
//
// This macro is available since SDL 3.2.0.
const MinorVersion = 2

// The current micro (or patchlevel) version of the SDL headers.
//
// If this were SDL version 3.2.1, this value would be 1.
//
// This macro is available since SDL 3.2.0.
const MicroVersion = 0

// This macro turns the version numbers into a numeric value.
//
// (1,2,3) becomes 1002003.
//
// major: the major version number.
//
// minor: the minorversion number.
//
// patch: the patch version number.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_VERSIONNUM
func Versionnum(major, minor, patch int) int {
	return major*1000000 + minor*1000 + patch
}

// This macro extracts the major version from a version number
//
// 1002003 becomes 1.
//
// version: the version number.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_VERSIONNUM_MAJOR
func VersionnumMajor(version int) int {
	return version / 1000000
}

// This macro extracts the minor version from a version number
//
// 1002003 becomes 2.
//
// version: the version number.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_VERSIONNUM_MINOR
func VersionnumMinor(version int) int {
	return version / 1000 % 1000
}

// This macro extracts the micro version from a version number
//
// 1002003 becomes 3.
//
// version: the version number.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_VERSIONNUM_MICRO
func VersionnumMicro(version int) int {
	return version % 1000
}

// This is the version number macro for the current SDL version.
//
// This macro is available since SDL 3.2.0.
const Version = 3002000

// This macro will evaluate to true if compiled with SDL at least X.Y.Z.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_VERSION_ATLEAST
func VersionAtLeast(major, minor, micro int) bool {
	return Version >= Versionnum(major, minor, micro)
}

// Get the version of SDL that is linked against your program.
//
// If you are linking to SDL dynamically, then it is possible that the current
// version will be different than the version you compiled against. This
// function returns the current version, while SDL_VERSION is the version you
// compiled with.
//
// This function may be called safely at any time, even before SDL_Init().
//
// Returns the version of the linked library.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetVersion
func GetVersion() int {
	return (int)(C.SDL_GetVersion())
}

// Get the code revision of SDL that is linked against your program.
//
// This value is the revision of the code you are linked with and may be
// different from the code you are compiling with, which is found in the
// constant SDL_REVISION.
//
// The revision is arbitrary string (a hash value) uniquely identifying the
// exact revision of the SDL library in use, and is only useful in comparing
// against other revisions. It is NOT an incrementing number.
//
// If SDL wasn't built from a git repository with the appropriate tools, this
// will return an empty string.
//
// You shouldn't use this function for anything but logging it for debugging
// purposes. The string is not intended to be reliable in any way.
//
// Returns an arbitrary string, uniquely identifying the exact revision of
// the SDL library in use.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRevision
func GetRevision() string {
	return C.GoString(C.SDL_GetRevision())
}
