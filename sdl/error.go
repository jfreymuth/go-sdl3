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
// #cgo noescape SDL_OutOfMemory
// #cgo nocallback SDL_OutOfMemory
// #cgo noescape SDL_GetError
// #cgo nocallback SDL_GetError
// #cgo noescape SDL_ClearError
// #cgo nocallback SDL_ClearError
// #include <SDL3/SDL.h>
// void wrap_SDL_SetError(_GoString_ msg) {
//     SDL_SetError("%.*s", _GoStringLen(msg), _GoStringPtr(msg));
// }
import "C"

// # CategoryError
//
// Simple error message routines for SDL.
//
// Most apps will interface with these APIs in exactly one function: when
// almost any SDL function call reports failure, you can get a human-readable
// string of the problem from [GetError].
//
// These strings are maintained per-thread, and apps are welcome to set their
// own errors, which is popular when building libraries on top of SDL for
// other apps to consume. These strings are set by calling [SetError].
//
// A common usage pattern is to have a function that returns true for success
// and false for failure, and do this when something fails:
//
//	if (something_went_wrong) {
//	return SDL_SetError("The thing broke in this specific way: %d", errcode);
//	}
//
// It's also common to just return false in this case if the failing thing
// is known to call [SetError], so errors simply propagate through.

// Set an error indicating that memory allocation failed.
//
// This function does not do any memory allocation.
//
// Returns false.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_OutOfMemory
func OutOfMemory() bool {
	return (bool)(C.SDL_OutOfMemory())
}

// Retrieve a message about the last error that occurred on the current
// thread.
//
// Functions in this package that return an error will call [GetError] on
// failure, and also call [ClearError], meaning this function will usually only
// return an error if SDL sets the error string after a successful operation.
//
// It is possible for multiple errors to occur before calling [GetError].
// Only the last error is returned.
//
// The message is only applicable when an SDL function has signaled an error.
// You must check the return values of SDL function calls to determine when to
// appropriately call [GetError]. You should *not* use the results of
// [GetError] to decide if an error has occurred! Sometimes SDL will set
// an error string even when reporting success.
//
// SDL will *not* clear the error string for successful API calls. You *must*
// check return values for failure cases before you can assume the error
// string applies.
//
// Error strings are set per-thread, so an error set in a different thread
// will not interfere with the current thread's operation.
//
// The returned value is a thread-local string which will remain valid until
// the current thread's error string is changed. The caller should make a copy
// if the value is needed after the next SDL API call.
//
// Returns a message with information about the specific error that occurred,
// or an empty string if there hasn't been an error message set since
// the last call to [ClearError].
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetError
func GetError() error {
	e := C.SDL_GetError()
	if e == nil {
		return nil
	}
	err := sdlError(C.GoString(e))
	return err
}

func getError() error {
	e := C.SDL_GetError()
	if e == nil {
		return sdlError("Unknown error")
	}
	err := sdlError(C.GoString(e))
	ClearError()
	return err
}

type sdlError string

func (err sdlError) Error() string {
	if err == "" {
		return "sdl: Unknown error"
	}
	return "sdl: " + string(err)
}

// Clear any previous error message for this thread.
//
// Returns true.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ClearError
func ClearError() bool {
	return (bool)(C.SDL_ClearError())
}

// Set the SDL error message for the current thread.
//
// Calling this function will replace any previous error message that was set.
//
// This function always returns false, since SDL frequently uses false to
// signify a failing result, leading to this idiom:
//
//	if (error_code) {
//	    return SDL_SetError("This operation has failed: %d", error_code);
//	}
//
// message: the error message
//
// Returns false.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
func SetError(message string) bool {
	C.wrap_SDL_SetError(message)
	return false
}
