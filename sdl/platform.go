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
// #cgo noescape SDL_GetPlatform
// #cgo nocallback SDL_GetPlatform
// #include <SDL3/SDL.h>
import "C"

// # CategoryPlatform
//
// SDL provides a means to identify the app's platform, both at compile time
// and runtime.

// Get the name of the platform.
//
// Here are the names returned for some (but not all) supported platforms:
//
//   - "Windows"
//   - "macOS"
//   - "Linux"
//   - "iOS"
//   - "Android"
//
// Returns the name of the platform. If the correct platform name is not
// available, returns a string beginning with the text "Unknown".
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetPlatform
func GetPlatform() string {
	return C.GoString(C.SDL_GetPlatform())
}
