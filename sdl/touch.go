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
// #cgo noescape SDL_GetTouchDevices
// #cgo nocallback SDL_GetTouchDevices
// #cgo noescape SDL_GetTouchDeviceName
// #cgo nocallback SDL_GetTouchDeviceName
// #cgo noescape SDL_GetTouchDeviceType
// #cgo nocallback SDL_GetTouchDeviceType
// #cgo noescape SDL_GetTouchFingers
// #cgo nocallback SDL_GetTouchFingers
// #include <SDL3/SDL.h>
import "C"
import "unsafe"

// # CategoryTouch
//
// SDL offers touch input, on platforms that support it. It can manage
// multiple touch devices and track multiple fingers on those devices.
//
// Touches are mostly dealt with through the event system, in the
// [EventFingerDown], [EventFingerMotion], and [EventFingerUp]
// events, but there are also functions to query for hardware details, etc.
//
// The touch system, by default, will also send virtual mouse events; this can
// be useful for making a some desktop apps work on a phone without
// significant changes. For apps that care about mouse and touch input
// separately, they should ignore mouse events that have a Which field of
// [TouchMouseID].

// A unique ID for a touch device.
//
// This ID is valid for the time the device is connected to the system, and is
// never reused for the lifetime of the application.
//
// The value 0 is an invalid ID.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_TouchID
type TouchID uint64

// A unique ID for a single finger on a touch device.
//
// This ID is valid for the time the finger (stylus, etc) is touching and will
// be unique for all fingers currently in contact, so this ID tracks the
// lifetime of a single continuous touch. This value may represent an index, a
// pointer, or some other unique ID, depending on the platform.
//
// The value 0 is an invalid ID.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_FingerID
type FingerID uint64

// An enum that describes the type of a touch device.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_TouchDeviceType
type TouchDeviceType uint32

const (
	TouchDeviceDirect           TouchDeviceType = iota // touch screen with window-relative coordinates
	TouchDeviceIndirectAbsolute                        // trackpad with absolute device coordinates
	TouchDeviceIndirectRelative                        // trackpad with screen cursor-relative coordinates
)

// Data about a single finger in a multitouch event.
//
// Each touch event is a collection of fingers that are simultaneously in
// contact with the touch device (so a "touch" can be a "multitouch," in
// reality), and this struct reports details of the specific fingers.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Finger
type Finger struct {
	Id       FingerID // the finger ID
	X        float32  // the x-axis location of the touch event, normalized (0...1)
	Y        float32  // the y-axis location of the touch event, normalized (0...1)
	Pressure float32  // the quantity of pressure applied, normalized (0...1)
}

// The [MouseID] for mouse events simulated with touch input.
//
// This macro is available since SDL 3.2.0.
const TouchMouseID MouseID = 0xFFFFFFFF

// The [TouchID] for touch events simulated with mouse input.
//
// This macro is available since SDL 3.2.0.
const MouseTouchID TouchID = 0xFFFFFFFF

// Get a list of registered touch devices.
//
// On some platforms SDL first sees the touch device if it was actually used.
// Therefore the returned list might be empty, although devices are available.
// After using all devices at least once the number will be correct.
//
// Returns a slice of touch device IDs or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTouchDevices
func GetTouchDevices() ([]TouchID, error) {
	var count C.int
	ids := C.SDL_GetTouchDevices(&count)
	if ids == nil {
		return nil, getError()
	}
	result := make([]TouchID, count)
	for i, id := range unsafe.Slice(ids, count) {
		result[i] = TouchID(id)
	}
	C.SDL_free(unsafe.Pointer(ids))
	return result, nil
}

// Get the touch device name as reported from the driver.
//
// touchID: the touch device instance ID.
//
// Returns touch device name or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTouchDeviceName
func GetTouchDeviceName(touchID TouchID) (string, error) {
	name := C.SDL_GetTouchDeviceName((C.SDL_TouchID)(touchID))
	if name == nil {
		return "", getError()
	}
	return C.GoString(name), nil
}

// Get the type of the given touch device.
//
// touchID: the ID of a touch device.
//
// Returns touch device type.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTouchDeviceType
func GetTouchDeviceType(touchID TouchID) TouchDeviceType {
	return (TouchDeviceType)(C.SDL_GetTouchDeviceType((C.SDL_TouchID)(touchID)))
}

// Get a list of active fingers for a given touch device.
//
// touchID: the ID of a touch device.
//
// Returns a slice of [Finger] structs or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTouchFingers
func GetTouchFingers(touchID TouchID) ([]Finger, error) {
	var count C.int
	fingers := C.SDL_GetTouchFingers((C.SDL_TouchID)(touchID), &count)
	if fingers == nil {
		return nil, getError()
	}
	result := make([]Finger, count)
	for i, f := range unsafe.Slice(fingers, count) {
		result[i] = Finger{
			FingerID(f.id),
			float32(f.x),
			float32(f.y),
			float32(f.pressure),
		}
	}
	C.SDL_free(unsafe.Pointer(fingers))
	return result, nil
}
