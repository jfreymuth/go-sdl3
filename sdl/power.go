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
// #cgo noescape SDL_GetPowerInfo
// #cgo nocallback SDL_GetPowerInfo
// #include <SDL3/SDL.h>
import "C"

// # CategoryPower
//
// SDL power management routines.
//
// There is a single function in this category: [GetPowerInfo].
//
// This function is useful for games on the go. This allows an app to know if
// it's running on a draining battery, which can be useful if the app wants to
// reduce processing, or perhaps framerate, to extend the duration of the
// battery's charge. Perhaps the app just wants to show a battery meter when
// fullscreen, or alert the user when the power is getting extremely low, so
// they can save their game.

// The basic state for the system's power supply.
//
// These are results returned by [GetPowerInfo].
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PowerState
type PowerState uint32

const PowerstateError PowerState = 0xFFFFFFFF

const (
	PowerstateUnknown   PowerState = iota // cannot determine power status
	PowerstateOnBattery                   // Not plugged in, running on the battery
	PowerstateNoBattery                   // Plugged in, no battery available
	PowerstateCharging                    // Plugged in, charging battery
	PowerstateCharged                     // Plugged in, battery charged
)

// Get the current power supply details.
//
// You should never take a battery status as absolute truth. Batteries
// (especially failing batteries) are delicate hardware, and the values
// reported here are best estimates based on what that hardware reports. It's
// not uncommon for older batteries to lose stored power much faster than it
// reports, or completely drain when reporting it has 20 percent left, etc.
//
// Battery status can change at any time; if you are concerned with power
// state, you should call this function frequently, and perhaps ignore changes
// until they seem to be stable for a few seconds.
//
// It's possible a platform can only report battery percentage or time left
// but not both.
//
// seconds: the seconds of battery life left. This will be -1 if we
// can't determine a value or there is no battery.
//
// percent: the percentage of battery life left, between 0 and 100. This will be
// -1 we can't determine a value or there is no battery.
//
// Returns the current battery state or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetPowerInfo
func GetPowerInfo() (state PowerState, seconds, percent int, err error) {
	var cseconds, cpercent C.int
	state = (PowerState)(C.SDL_GetPowerInfo(&cseconds, &cpercent))
	if state == PowerstateError {
		return PowerstateError, -1, -1, getError()
	}
	return state, int(cseconds), int(cpercent), nil
}
