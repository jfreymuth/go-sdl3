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
// #cgo noescape SDL_GetTicks
// #cgo nocallback SDL_GetTicks
// #cgo noescape SDL_GetTicksNS
// #cgo nocallback SDL_GetTicksNS
// #cgo noescape SDL_GetPerformanceCounter
// #cgo nocallback SDL_GetPerformanceCounter
// #cgo noescape SDL_GetPerformanceFrequency
// #cgo nocallback SDL_GetPerformanceFrequency
// #include <SDL3/SDL.h>
import "C"

// # CategoryTimer
//
// SDL provides time management functionality. It is useful for dealing with
// (usually) small durations of time.
//
// This is not to be confused with _calendar time_ management, which is
// provided by [CategoryTime](CategoryTime).
//
// This category covers measuring time elapsed ([GetTicks],
// [GetPerformanceCounter]), putting a thread to sleep for a certain
// amount of time (SDL_Delay(), SDL_DelayNS(), SDL_DelayPrecise()), and firing
// a callback function after a certain amount of time has elasped
// (SDL_AddTimer(), etc).
//
// There are also useful macros to convert between time units, like
// SDL_SECONDS_TO_NS() and such.

// Get the number of milliseconds since SDL library initialization.
//
// Returns an unsigned 64-bit value representing the number of milliseconds
// since the SDL library initialized.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTicks
func GetTicks() uint64 {
	return (uint64)(C.SDL_GetTicks())
}

// Get the number of nanoseconds since SDL library initialization.
//
// Returns an unsigned 64-bit value representing the number of nanoseconds
// since the SDL library initialized.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTicksNS
func GetTicksNS() uint64 {
	return (uint64)(C.SDL_GetTicksNS())
}

// Get the current value of the high resolution counter.
//
// This function is typically used for profiling.
//
// The counter values are only meaningful relative to each other. Differences
// between values can be converted to times by using
// [GetPerformanceFrequency].
//
// Returns the current counter value.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetPerformanceCounter
func GetPerformanceCounter() uint64 {
	return (uint64)(C.SDL_GetPerformanceCounter())
}

// Get the count per second of the high resolution counter.
//
// Returns a platform-specific count per second.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetPerformanceFrequency
func GetPerformanceFrequency() uint64 {
	return (uint64)(C.SDL_GetPerformanceFrequency())
}
