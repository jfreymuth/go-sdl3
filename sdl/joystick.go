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
// #cgo noescape SDL_HasJoystick
// #cgo nocallback SDL_HasJoystick
// #cgo noescape SDL_GetJoysticks
// #cgo nocallback SDL_GetJoysticks
// #cgo noescape SDL_GetJoystickNameForID
// #cgo nocallback SDL_GetJoystickNameForID
// #cgo noescape SDL_GetJoystickPathForID
// #cgo nocallback SDL_GetJoystickPathForID
// #cgo noescape SDL_GetJoystickPlayerIndexForID
// #cgo nocallback SDL_GetJoystickPlayerIndexForID
// #cgo noescape SDL_GetJoystickGUIDForID
// #cgo nocallback SDL_GetJoystickGUIDForID
// #cgo noescape SDL_GetJoystickVendorForID
// #cgo nocallback SDL_GetJoystickVendorForID
// #cgo noescape SDL_GetJoystickProductForID
// #cgo nocallback SDL_GetJoystickProductForID
// #cgo noescape SDL_GetJoystickProductVersionForID
// #cgo nocallback SDL_GetJoystickProductVersionForID
// #cgo noescape SDL_GetJoystickTypeForID
// #cgo nocallback SDL_GetJoystickTypeForID
// #cgo noescape SDL_OpenJoystick
// #cgo nocallback SDL_OpenJoystick
// #cgo noescape SDL_GetJoystickFromID
// #cgo nocallback SDL_GetJoystickFromID
// #cgo noescape SDL_GetJoystickFromPlayerIndex
// #cgo nocallback SDL_GetJoystickFromPlayerIndex
// #cgo noescape wrap_SDL_AttachVirtualJoystick
// #cgo noescape SDL_DetachVirtualJoystick
// #cgo noescape SDL_IsJoystickVirtual
// #cgo nocallback SDL_IsJoystickVirtual
// #cgo noescape SDL_SetJoystickVirtualAxis
// #cgo nocallback SDL_SetJoystickVirtualAxis
// #cgo noescape SDL_SetJoystickVirtualBall
// #cgo nocallback SDL_SetJoystickVirtualBall
// #cgo noescape SDL_SetJoystickVirtualButton
// #cgo nocallback SDL_SetJoystickVirtualButton
// #cgo noescape SDL_SetJoystickVirtualHat
// #cgo nocallback SDL_SetJoystickVirtualHat
// #cgo noescape SDL_SetJoystickVirtualTouchpad
// #cgo nocallback SDL_SetJoystickVirtualTouchpad
// #cgo noescape SDL_SendJoystickVirtualSensorData
// #cgo nocallback SDL_SendJoystickVirtualSensorData
// #cgo noescape SDL_GetJoystickProperties
// #cgo nocallback SDL_GetJoystickProperties
// #cgo noescape SDL_GetJoystickName
// #cgo nocallback SDL_GetJoystickName
// #cgo noescape SDL_GetJoystickPath
// #cgo nocallback SDL_GetJoystickPath
// #cgo noescape SDL_GetJoystickPlayerIndex
// #cgo nocallback SDL_GetJoystickPlayerIndex
// #cgo noescape SDL_SetJoystickPlayerIndex
// #cgo nocallback SDL_SetJoystickPlayerIndex
// #cgo noescape SDL_GetJoystickGUID
// #cgo nocallback SDL_GetJoystickGUID
// #cgo noescape SDL_GetJoystickVendor
// #cgo nocallback SDL_GetJoystickVendor
// #cgo noescape SDL_GetJoystickProduct
// #cgo nocallback SDL_GetJoystickProduct
// #cgo noescape SDL_GetJoystickProductVersion
// #cgo nocallback SDL_GetJoystickProductVersion
// #cgo noescape SDL_GetJoystickFirmwareVersion
// #cgo nocallback SDL_GetJoystickFirmwareVersion
// #cgo noescape SDL_GetJoystickSerial
// #cgo nocallback SDL_GetJoystickSerial
// #cgo noescape SDL_GetJoystickType
// #cgo nocallback SDL_GetJoystickType
// #cgo noescape SDL_GetJoystickGUIDInfo
// #cgo nocallback SDL_GetJoystickGUIDInfo
// #cgo noescape SDL_JoystickConnected
// #cgo nocallback SDL_JoystickConnected
// #cgo noescape SDL_GetJoystickID
// #cgo nocallback SDL_GetJoystickID
// #cgo noescape SDL_GetNumJoystickAxes
// #cgo nocallback SDL_GetNumJoystickAxes
// #cgo noescape SDL_GetNumJoystickBalls
// #cgo nocallback SDL_GetNumJoystickBalls
// #cgo noescape SDL_GetNumJoystickHats
// #cgo nocallback SDL_GetNumJoystickHats
// #cgo noescape SDL_GetNumJoystickButtons
// #cgo nocallback SDL_GetNumJoystickButtons
// #cgo noescape SDL_SetJoystickEventsEnabled
// #cgo nocallback SDL_SetJoystickEventsEnabled
// #cgo noescape SDL_JoystickEventsEnabled
// #cgo nocallback SDL_JoystickEventsEnabled
// #cgo noescape SDL_UpdateJoysticks
// #cgo nocallback SDL_UpdateJoysticks
// #cgo noescape SDL_GetJoystickAxis
// #cgo nocallback SDL_GetJoystickAxis
// #cgo noescape SDL_GetJoystickAxisInitialState
// #cgo nocallback SDL_GetJoystickAxisInitialState
// #cgo noescape SDL_GetJoystickBall
// #cgo nocallback SDL_GetJoystickBall
// #cgo noescape SDL_GetJoystickHat
// #cgo nocallback SDL_GetJoystickHat
// #cgo noescape SDL_GetJoystickButton
// #cgo nocallback SDL_GetJoystickButton
// #cgo noescape SDL_RumbleJoystick
// #cgo noescape SDL_RumbleJoystickTriggers
// #cgo noescape SDL_SetJoystickLED
// #cgo noescape SDL_SendJoystickEffect
// #cgo noescape SDL_CloseJoystick
// #cgo noescape SDL_GetJoystickConnectionState
// #cgo nocallback SDL_GetJoystickConnectionState
// #cgo noescape SDL_GetJoystickPowerInfo
// #cgo nocallback SDL_GetJoystickPowerInfo
// #include <SDL3/SDL.h>
// extern SDL_JoystickID wrap_SDL_AttachVirtualJoystick(SDL_VirtualJoystickDesc *desc, uintptr_t userdata);
import "C"
import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

// # CategoryJoystick
//
// SDL joystick support.
//
// This is the lower-level joystick handling. If you want the simpler option,
// where what each button does is well-defined, you should use the gamepad API
// instead.
//
// The term "id" is the current instantiation of a joystick device in
// the system, if the joystick is removed and then re-inserted then it will
// get a new id, id's are monotonically increasing
// identifiers of a joystick plugged in.
//
// The term "player_index" is the number assigned to a player on a specific
// controller. For XInput controllers this returns the XInput user index. Many
// joysticks will not be able to supply this information.
//
// SDL_GUID is used as a stable 128-bit identifier for a joystick device that
// does not change over time. It identifies class of the device (a X360 wired
// controller for example). This identifier is platform dependent.
//
// In order to use these functions, SDL_Init() must have been called with the
// SDL_INIT_JOYSTICK flag. This causes SDL to scan the system for joysticks,
// and load appropriate drivers.
//
// If you would like to receive joystick updates while the application is in
// the background, you should set the following hint before calling
// SDL_Init(): SDL_HINT_JOYSTICK_ALLOW_BACKGROUND_EVENTS

// The joystick structure used to identify an SDL joystick.
//
// This is opaque data.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Joystick
type Joystick C.struct_SDL_Joystick

// This is a unique ID for a joystick for the time it is connected to the
// system, and is never reused for the lifetime of the application.
//
// If the joystick is disconnected and reconnected, it will get a new ID.
//
// The value 0 is an invalid ID.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_JoystickID
type JoystickID uint32

// An enum of some common joystick types.
//
// In some cases, SDL can identify a low-level joystick as being a certain
// type of device, and will report it through SDL_GetJoystickType (or
// SDL_GetJoystickTypeForID).
//
// This is by no means a complete list of everything that can be plugged into
// a computer.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_JoystickType
type JoystickType uint32

const (
	JoystickTypeUnknown JoystickType = iota
	JoystickTypeGamepad
	JoystickTypeWheel
	JoystickTypeArcadeStick
	JoystickTypeFlightStick
	JoystickTypeDancePad
	JoystickTypeGuitar
	JoystickTypeDrumKit
	JoystickTypeArcadePad
	JoystickTypeThrottle
	JoystickTypeCount = iota
)

// Possible connection states for a joystick device.
//
// This is used by SDL_GetJoystickConnectionState to report how a device is
// connected to the system.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_JoystickConnectionState
type JoystickConnectionState uint32

const JoystickConnectionInvalid JoystickConnectionState = 0xFFFFFFFF

const (
	JoystickConnectionUnknown JoystickConnectionState = iota
	JoystickConnectionWired
	JoystickConnectionWireless
)

// The largest value an SDL_Joystick's axis can report.
//
// This macro is available since SDL 3.2.0.
const JoystickAxisMax = 32767

// The smallest value an SDL_Joystick's axis can report.
//
// This is a negative number!
//
// This macro is available since SDL 3.2.0.
const JoystickAxisMin = -32768

// Locking for atomic access to the joystick API.
//
// The SDL joystick functions are thread-safe, however you can lock the
// joysticks while processing to guarantee that the joystick list won't change
// and joystick and gamepad events will not be delivered.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_LockJoysticks
func LockJoysticks() {
	C.SDL_LockJoysticks()
}

// Unlocking for atomic access to the joystick API.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_UnlockJoysticks
func UnlockJoysticks() {
	C.SDL_UnlockJoysticks()
}

// Return whether a joystick is currently connected.
//
// Returns true if a joystick is connected, false otherwise.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasJoystick
func HasJoystick() bool {
	return (bool)(C.SDL_HasJoystick())
}

// Get a list of currently connected joysticks.
//
// count: a pointer filled in with the number of joysticks returned, may
// be NULL.
//
// Returns a 0 terminated array of joystick instance IDs or NULL on failure;
// call SDL_GetError() for more information. This should be freed
// with SDL_free() when it is no longer needed.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoysticks
func GetJoysticks() ([]JoystickID, error) {
	var count C.int
	j := C.SDL_GetJoysticks(&count)
	if j == nil {
		return nil, getError()
	}
	result := make([]JoystickID, count)
	for i, j := range unsafe.Slice(j, count) {
		result[i] = JoystickID(j)
	}
	C.SDL_free(unsafe.Pointer(j))
	return result, nil
}

// Get the implementation dependent name of a joystick.
//
// This can be called before any joysticks are opened.
//
// id: the joystick instance ID.
//
// Returns the name of the selected joystick. If no name can be found, this
// function returns NULL; call SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickNameForID
func (id JoystickID) Name() (string, error) {
	name := C.SDL_GetJoystickNameForID((C.SDL_JoystickID)(id))
	if name == nil {
		return "", getError()
	}
	return C.GoString(name), nil
}

// Get the implementation dependent path of a joystick.
//
// This can be called before any joysticks are opened.
//
// id: the joystick instance ID.
//
// Returns the path of the selected joystick. If no path can be found, this
// function returns NULL; call SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickPathForID
func (id JoystickID) Path() (string, error) {
	name := C.SDL_GetJoystickPathForID((C.SDL_JoystickID)(id))
	if name == nil {
		return "", getError()
	}
	return C.GoString(name), nil
}

// Get the player index of a joystick.
//
// This can be called before any joysticks are opened.
//
// id: the joystick instance ID.
//
// Returns the player index of a joystick, or -1 if it's not available.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickPlayerIndexForID
func (id JoystickID) PlayerIndex() int32 {
	return (int32)(C.SDL_GetJoystickPlayerIndexForID((C.SDL_JoystickID)(id)))
}

// Get the implementation-dependent GUID of a joystick.
//
// This can be called before any joysticks are opened.
//
// id: the joystick instance ID.
//
// Returns the GUID of the selected joystick. If called with an invalid
// id, this function returns a zero GUID.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUIDForID
func (id JoystickID) GUID() GUID {
	return (GUID)(C.SDL_GetJoystickGUIDForID((C.SDL_JoystickID)(id)).data)
}

// Get the USB vendor ID of a joystick, if available.
//
// This can be called before any joysticks are opened. If the vendor ID isn't
// available this function returns 0.
//
// id: the joystick instance ID.
//
// Returns the USB vendor ID of the selected joystick. If called with an
// invalid id, this function returns 0.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickVendorForID
func (id JoystickID) Vendor() uint16 {
	return (uint16)(C.SDL_GetJoystickVendorForID((C.SDL_JoystickID)(id)))
}

// Get the USB product ID of a joystick, if available.
//
// This can be called before any joysticks are opened. If the product ID isn't
// available this function returns 0.
//
// id: the joystick instance ID.
//
// Returns the USB product ID of the selected joystick. If called with an
// invalid id, this function returns 0.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductForID
func (id JoystickID) Product() uint16 {
	return (uint16)(C.SDL_GetJoystickProductForID((C.SDL_JoystickID)(id)))
}

// Get the product version of a joystick, if available.
//
// This can be called before any joysticks are opened. If the product version
// isn't available this function returns 0.
//
// id: the joystick instance ID.
//
// Returns the product version of the selected joystick. If called with an
// invalid id, this function returns 0.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductVersionForID
func (id JoystickID) ProductVersion() uint16 {
	return (uint16)(C.SDL_GetJoystickProductVersionForID((C.SDL_JoystickID)(id)))
}

// Get the type of a joystick, if available.
//
// This can be called before any joysticks are opened.
//
// id: the joystick instance ID.
//
// Returns the SDL_JoystickType of the selected joystick. If called with an
// invalid id, this function returns
// `SDL_JOYSTICK_TYPE_UNKNOWN`.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickTypeForID
func (id JoystickID) Type() JoystickType {
	return (JoystickType)(C.SDL_GetJoystickTypeForID((C.SDL_JoystickID)(id)))
}

// Open a joystick for use.
//
// The joystick subsystem must be initialized before a joystick can be opened
// for use.
//
// id: the joystick instance ID.
//
// Returns a joystick identifier or NULL on failure; call SDL_GetError() for
// more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_OpenJoystick
func OpenJoystick(id JoystickID) (*Joystick, error) {
	j := (*Joystick)(C.SDL_OpenJoystick((C.SDL_JoystickID)(id)))
	if j == nil {
		return nil, getError()
	}
	return j, nil
}

// Get the SDL_Joystick associated with an instance ID, if it has been opened.
//
// id: the instance ID to get the SDL_Joystick for.
//
// Returns an SDL_Joystick on success or NULL on failure or if it hasn't been
// opened yet; call SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickFromID
func GetJoystickFromID(id JoystickID) (*Joystick, error) {
	j := (*Joystick)(C.SDL_GetJoystickFromID((C.SDL_JoystickID)(id)))
	if j == nil {
		return nil, getError()
	}
	return j, nil
}

// Get the SDL_Joystick associated with a player index.
//
// player_index: the player index to get the SDL_Joystick for.
//
// Returns an SDL_Joystick on success or NULL on failure; call SDL_GetError()
// for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickFromPlayerIndex
func GetJoystickFromPlayerIndex(player_index int) (*Joystick, error) {
	j := (*Joystick)(C.SDL_GetJoystickFromPlayerIndex((C.int)(player_index)))
	if j == nil {
		return nil, getError()
	}
	return j, nil
}

// The structure that describes a virtual joystick touchpad.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_VirtualJoystickTouchpadDesc
type VirtualJoystickTouchpadDesc struct {
	NFingers int // the number of simultaneous fingers on this touchpad
}

// The structure that describes a virtual joystick sensor.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_VirtualJoystickSensorDesc
type VirtualJoystickSensorDesc struct {
	Type SensorType // the type of this sensor
	Rate float32    // the update frequency of this sensor, may be 0.0f
}

// The structure that describes a virtual joystick.
//
// This structure should be initialized using SDL_INIT_INTERFACE(). All
// elements of this structure are optional.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_VirtualJoystickDesc
type VirtualJoystickDesc struct {
	Type       JoystickType                  // `SDL_JoystickType`
	VendorId   uint16                        // the USB vendor ID of this joystick
	ProductId  uint16                        // the USB product ID of this joystick
	NAxes      int                           // the number of axes on this joystick
	NButtons   int                           // the number of buttons on this joystick
	NBalls     int                           // the number of balls on this joystick
	NHats      int                           // the number of hats on this joystick
	ButtonMask GamepadButton                 // A mask of which buttons are valid for this controller e.g. (1 << SDL_GAMEPAD_BUTTON_SOUTH)
	AxisMask   GamepadAxis                   // A mask of which axes are valid for this controller e.g. (1 << SDL_GAMEPAD_AXIS_LEFTX)
	Name       string                        // the name of the joystick
	Touchpads  []VirtualJoystickTouchpadDesc // A pointer to an array of touchpad descriptions, required if `ntouchpads` is > 0
	Sensors    []VirtualJoystickSensorDesc   // A pointer to an array of sensor descriptions, required if `nsensors` is > 0

	Update            func()                                                    // Called when the joystick state should be updated
	SetPlayerIndex    func(playerIndex int)                                     // Called when the player index is set
	Rumble            func(lowFrequencyRumble, highFrequencyRumble uint16) bool // Implements SDL_RumbleJoystick()
	RumbleTriggers    func(leftRumble, rightRumble uint16) bool                 // Implements SDL_RumbleJoystickTriggers()
	SetLED            func(red, green, blue byte) bool                          // Implements SDL_SetJoystickLED()
	SendEffect        func(data []byte) bool                                    // Implements SDL_SendJoystickEffect()
	SetSensorsEnabled func(enabled bool) bool                                   // Implements SDL_SetGamepadSensorEnabled()
	Cleanup           func()                                                    // Cleans up the userdata when the joystick is detached
}

//export cb_VirtualJoystickUpdate
func cb_VirtualJoystickUpdate(userdata uintptr) {
	h := cgo.Handle(userdata)
	h.Value().(*VirtualJoystickDesc).Update()
}

//export cb_VirtualJoystickSetPlayerIndex
func cb_VirtualJoystickSetPlayerIndex(userdata uintptr, playerIndex C.int) {
	h := cgo.Handle(userdata)
	h.Value().(*VirtualJoystickDesc).SetPlayerIndex((int)(playerIndex))
}

//export cb_VirtualJoystickRumble
func cb_VirtualJoystickRumble(userdata uintptr, lowFrequencyRumble, highFrequencyRumble C.Uint16) {
	h := cgo.Handle(userdata)
	h.Value().(*VirtualJoystickDesc).Rumble(uint16(lowFrequencyRumble), uint16(highFrequencyRumble))
}

//export cb_VirtualJoystickRumbleTriggers
func cb_VirtualJoystickRumbleTriggers(userdata uintptr, leftRumble, rightRumble C.Uint16) {
	h := cgo.Handle(userdata)
	h.Value().(*VirtualJoystickDesc).RumbleTriggers(uint16(leftRumble), uint16(rightRumble))
}

//export cb_VirtualJoystickSetLED
func cb_VirtualJoystickSetLED(userdata uintptr, red, green, blue C.Uint8) {
	h := cgo.Handle(userdata)
	h.Value().(*VirtualJoystickDesc).SetLED(byte(red), byte(green), byte(blue))
}

//export cb_VirtualJoystickSendEffect
func cb_VirtualJoystickSendEffect(userdata uintptr, data unsafe.Pointer, size C.int) {
	h := cgo.Handle(userdata)
	h.Value().(*VirtualJoystickDesc).SendEffect(unsafe.Slice((*byte)(data), size))
}

//export cb_VirtualJoystickSetSensorsEnabled
func cb_VirtualJoystickSetSensorsEnabled(userdata uintptr, enabled C.bool) {
	h := cgo.Handle(userdata)
	h.Value().(*VirtualJoystickDesc).SetSensorsEnabled(bool(enabled))
}

//export cb_VirtualJoystickCleanup
func cb_VirtualJoystickCleanup(userdata uintptr) {
	h := cgo.Handle(userdata)
	h.Value().(*VirtualJoystickDesc).Cleanup()
	h.Delete()
}

// Attach a new virtual joystick.
//
// desc: joystick description, initialized using SDL_INIT_INTERFACE().
//
// Returns the joystick instance ID, or 0 on failure; call SDL_GetError() for
// more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AttachVirtualJoystick
func AttachVirtualJoystick(desc *VirtualJoystickDesc) (JoystickID, error) {
	h := cgo.NewHandle(desc)
	touchpads := make([]C.SDL_VirtualJoystickTouchpadDesc, len(desc.Touchpads))
	for i, t := range desc.Touchpads {
		touchpads[i] = C.SDL_VirtualJoystickTouchpadDesc{nfingers: C.Uint16(t.NFingers)}
	}
	sensors := make([]C.SDL_VirtualJoystickSensorDesc, len(desc.Sensors))
	for i, s := range desc.Sensors {
		sensors[i] = C.SDL_VirtualJoystickSensorDesc{_type: C.SDL_SensorType(s.Type), rate: C.float(s.Rate)}
	}
	cdesc := &C.SDL_VirtualJoystickDesc{
		_type:       C.Uint16(desc.Type),
		vendor_id:   C.Uint16(desc.VendorId),
		product_id:  C.Uint16(desc.ProductId),
		naxes:       C.Uint16(desc.NAxes),
		nbuttons:    C.Uint16(desc.NButtons),
		nballs:      C.Uint16(desc.NBalls),
		nhats:       C.Uint16(desc.NHats),
		ntouchpads:  C.Uint16(len(desc.Touchpads)),
		nsensors:    C.Uint16(len(desc.Sensors)),
		button_mask: C.Uint32(desc.ButtonMask),
		axis_mask:   C.Uint32(desc.AxisMask),
		name:        tmpstring(desc.Name),
		touchpads:   unsafe.SliceData(touchpads),
		sensors:     unsafe.SliceData(sensors),
	}
	var pinner runtime.Pinner
	pinner.Pin(cdesc.touchpads)
	pinner.Pin(cdesc.sensors)
	id := (JoystickID)(C.wrap_SDL_AttachVirtualJoystick(cdesc, C.uintptr_t(h)))
	pinner.Unpin()
	if id == 0 {
		h.Delete()
		return id, getError()
	}
	return id, nil
}

// Detach a virtual joystick.
//
// id: the joystick instance ID, previously returned from
// SDL_AttachVirtualJoystick().
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DetachVirtualJoystick
func DetachVirtualJoystick(id JoystickID) bool {
	return (bool)(C.SDL_DetachVirtualJoystick((C.SDL_JoystickID)(id)))
}

// Query whether or not a joystick is virtual.
//
// id: the joystick instance ID.
//
// Returns true if the joystick is virtual, false otherwise.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_IsJoystickVirtual
func (id JoystickID) Virtual() bool {
	return (bool)(C.SDL_IsJoystickVirtual((C.SDL_JoystickID)(id)))
}

// Set the state of an axis on an opened virtual joystick.
//
// Please note that values set here will not be applied until the next call to
// SDL_UpdateJoysticks, which can either be called directly, or can be called
// indirectly through various other SDL APIs, including, but not limited to
// the following: SDL_PollEvent, SDL_PumpEvents, SDL_WaitEventTimeout,
// SDL_WaitEvent.
//
// Note that when sending trigger axes, you should scale the value to the full
// range of Sint16. For example, a trigger at rest would have the value of
// `SDL_JOYSTICK_AXIS_MIN`.
//
// joystick: the virtual joystick on which to set state.
//
// axis: the index of the axis on the virtual joystick to update.
//
// value: the new value for the specified axis.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualAxis
func (joystick *Joystick) SetVirtualAxis(axis int, value int16) error {
	if !C.SDL_SetJoystickVirtualAxis((*C.SDL_Joystick)(joystick), (C.int)(axis), (C.Sint16)(value)) {
		return getError()
	}
	return nil
}

// Generate ball motion on an opened virtual joystick.
//
// Please note that values set here will not be applied until the next call to
// SDL_UpdateJoysticks, which can either be called directly, or can be called
// indirectly through various other SDL APIs, including, but not limited to
// the following: SDL_PollEvent, SDL_PumpEvents, SDL_WaitEventTimeout,
// SDL_WaitEvent.
//
// joystick: the virtual joystick on which to set state.
//
// ball: the index of the ball on the virtual joystick to update.
//
// xrel: the relative motion on the X axis.
//
// yrel: the relative motion on the Y axis.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualBall
func (joystick *Joystick) SetVirtualBall(ball int, xrel int16, yrel int16) error {
	if !C.SDL_SetJoystickVirtualBall((*C.SDL_Joystick)(joystick), (C.int)(ball), (C.Sint16)(xrel), (C.Sint16)(yrel)) {
		return getError()
	}
	return nil
}

// Set the state of a button on an opened virtual joystick.
//
// Please note that values set here will not be applied until the next call to
// SDL_UpdateJoysticks, which can either be called directly, or can be called
// indirectly through various other SDL APIs, including, but not limited to
// the following: SDL_PollEvent, SDL_PumpEvents, SDL_WaitEventTimeout,
// SDL_WaitEvent.
//
// joystick: the virtual joystick on which to set state.
//
// button: the index of the button on the virtual joystick to update.
//
// down: true if the button is pressed, false otherwise.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualButton
func (joystick *Joystick) SetVirtualButton(button int, down bool) error {
	if !C.SDL_SetJoystickVirtualButton((*C.SDL_Joystick)(joystick), (C.int)(button), (C.bool)(down)) {
		return getError()
	}
	return nil
}

// Set the state of a hat on an opened virtual joystick.
//
// Please note that values set here will not be applied until the next call to
// SDL_UpdateJoysticks, which can either be called directly, or can be called
// indirectly through various other SDL APIs, including, but not limited to
// the following: SDL_PollEvent, SDL_PumpEvents, SDL_WaitEventTimeout,
// SDL_WaitEvent.
//
// joystick: the virtual joystick on which to set state.
//
// hat: the index of the hat on the virtual joystick to update.
//
// value: the new value for the specified hat.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualHat
func (joystick *Joystick) SetVirtualHat(hat int, value byte) error {
	if !C.SDL_SetJoystickVirtualHat((*C.SDL_Joystick)(joystick), (C.int)(hat), (C.Uint8)(value)) {
		return getError()
	}
	return nil
}

// Set touchpad finger state on an opened virtual joystick.
//
// Please note that values set here will not be applied until the next call to
// SDL_UpdateJoysticks, which can either be called directly, or can be called
// indirectly through various other SDL APIs, including, but not limited to
// the following: SDL_PollEvent, SDL_PumpEvents, SDL_WaitEventTimeout,
// SDL_WaitEvent.
//
// joystick: the virtual joystick on which to set state.
//
// touchpad: the index of the touchpad on the virtual joystick to
// update.
//
// finger: the index of the finger on the touchpad to set.
//
// down: true if the finger is pressed, false if the finger is released.
//
// x: the x coordinate of the finger on the touchpad, normalized 0 to 1,
// with the origin in the upper left.
//
// y: the y coordinate of the finger on the touchpad, normalized 0 to 1,
// with the origin in the upper left.
//
// pressure: the pressure of the finger.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetJoystickVirtualTouchpad
func (joystick *Joystick) SetVirtualTouchpad(touchpad int, finger int, down bool, x float32, y float32, pressure float32) error {
	if !C.SDL_SetJoystickVirtualTouchpad((*C.SDL_Joystick)(joystick), (C.int)(touchpad), (C.int)(finger), (C.bool)(down), (C.float)(x), (C.float)(y), (C.float)(pressure)) {
		return getError()
	}
	return nil
}

// Send a sensor update for an opened virtual joystick.
//
// Please note that values set here will not be applied until the next call to
// SDL_UpdateJoysticks, which can either be called directly, or can be called
// indirectly through various other SDL APIs, including, but not limited to
// the following: SDL_PollEvent, SDL_PumpEvents, SDL_WaitEventTimeout,
// SDL_WaitEvent.
//
// joystick: the virtual joystick on which to set state.
//
// type: the type of the sensor on the virtual joystick to update.
//
// sensor_timestamp: a 64-bit timestamp in nanoseconds associated with
// the sensor reading.
//
// data: the data associated with the sensor reading.
//
// num_values: the number of values pointed to by `data`.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SendJoystickVirtualSensorData
func (joystick *Joystick) SendVirtualSensorData(typ SensorType, sensor_timestamp uint64, data []float32) error {
	if !C.SDL_SendJoystickVirtualSensorData((*C.SDL_Joystick)(joystick), (C.SDL_SensorType)(typ), (C.Uint64)(sensor_timestamp), (*C.float)(unsafe.SliceData(data)), (C.int)(len(data))) {
		return getError()
	}
	return nil
}

// Get the properties associated with a joystick.
//
// The following read-only properties are provided by SDL:
//
// - `SDL_PROP_JOYSTICK_CAP_MONO_LED_BOOLEAN`: true if this joystick has an
// LED that has adjustable brightness
// - `SDL_PROP_JOYSTICK_CAP_RGB_LED_BOOLEAN`: true if this joystick has an LED
// that has adjustable color
// - `SDL_PROP_JOYSTICK_CAP_PLAYER_LED_BOOLEAN`: true if this joystick has a
// player LED
// - `SDL_PROP_JOYSTICK_CAP_RUMBLE_BOOLEAN`: true if this joystick has
// left/right rumble
// - `SDL_PROP_JOYSTICK_CAP_TRIGGER_RUMBLE_BOOLEAN`: true if this joystick has
// simple trigger rumble
//
// joystick: the SDL_Joystick obtained from SDL_OpenJoystick().
//
// Returns a valid property ID on success or 0 on failure; call
// SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickProperties
func (joystick *Joystick) Properties() (PropertiesID, error) {
	props := (PropertiesID)(C.SDL_GetJoystickProperties((*C.SDL_Joystick)(joystick)))
	if props == 0 {
		return 0, getError()
	}
	return props, nil
}

const PropJoystickCapMonoLedBoolean = "SDL.joystick.cap.mono_led"
const PropJoystickCapRgbLedBoolean = "SDL.joystick.cap.rgb_led"
const PropJoystickCapPlayerLedBoolean = "SDL.joystick.cap.player_led"
const PropJoystickCapRumbleBoolean = "SDL.joystick.cap.rumble"
const PropJoystickCapTriggerRumbleBoolean = "SDL.joystick.cap.trigger_rumble"

// Get the implementation dependent name of a joystick.
//
// joystick: the SDL_Joystick obtained from SDL_OpenJoystick().
//
// Returns the name of the selected joystick. If no name can be found, this
// function returns NULL; call SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickName
func (joystick *Joystick) Name() (string, error) {
	name := C.SDL_GetJoystickName((*C.SDL_Joystick)(joystick))
	if name == nil {
		return "", getError()
	}
	return C.GoString(name), nil
}

// Get the implementation dependent path of a joystick.
//
// joystick: the SDL_Joystick obtained from SDL_OpenJoystick().
//
// Returns the path of the selected joystick. If no path can be found, this
// function returns NULL; call SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickPath
func (joystick *Joystick) Path() (string, error) {
	path := C.SDL_GetJoystickPath((*C.SDL_Joystick)(joystick))
	if path == nil {
		return "", getError()
	}
	return C.GoString(path), nil
}

// Get the player index of an opened joystick.
//
// For XInput controllers this returns the XInput user index. Many joysticks
// will not be able to supply this information.
//
// joystick: the SDL_Joystick obtained from SDL_OpenJoystick().
//
// Returns the player index, or -1 if it's not available.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickPlayerIndex
func (joystick *Joystick) PlayerIndex() int {
	return (int)(C.SDL_GetJoystickPlayerIndex((*C.SDL_Joystick)(joystick)))
}

// Set the player index of an opened joystick.
//
// joystick: the SDL_Joystick obtained from SDL_OpenJoystick().
//
// player_index: player index to assign to this joystick, or -1 to clear
// the player index and turn off player LEDs.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetJoystickPlayerIndex
func (joystick *Joystick) SetPlayerIndex(playerIndex int) error {
	if !C.SDL_SetJoystickPlayerIndex((*C.SDL_Joystick)(joystick), (C.int)(playerIndex)) {
		return getError()
	}
	return nil
}

// Get the implementation-dependent GUID for the joystick.
//
// This function requires an open joystick.
//
// joystick: the SDL_Joystick obtained from SDL_OpenJoystick().
//
// Returns the GUID of the given joystick. If called on an invalid index,
// this function returns a zero GUID; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUID
func (joystick *Joystick) GUID() (GUID, error) {
	g := (GUID)(C.SDL_GetJoystickGUID((*C.SDL_Joystick)(joystick)).data)
	if g == (GUID{}) {
		return GUID{}, getError()
	}
	return g, nil
}

// Get the USB vendor ID of an opened joystick, if available.
//
// If the vendor ID isn't available this function returns 0.
//
// joystick: the SDL_Joystick obtained from SDL_OpenJoystick().
//
// Returns the USB vendor ID of the selected joystick, or 0 if unavailable.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickVendor
func (joystick *Joystick) Vendor() uint16 {
	return (uint16)(C.SDL_GetJoystickVendor((*C.SDL_Joystick)(joystick)))
}

// Get the USB product ID of an opened joystick, if available.
//
// If the product ID isn't available this function returns 0.
//
// joystick: the SDL_Joystick obtained from SDL_OpenJoystick().
//
// Returns the USB product ID of the selected joystick, or 0 if unavailable.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickProduct
func (joystick *Joystick) Product() uint16 {
	return (uint16)(C.SDL_GetJoystickProduct((*C.SDL_Joystick)(joystick)))
}

// Get the product version of an opened joystick, if available.
//
// If the product version isn't available this function returns 0.
//
// joystick: the SDL_Joystick obtained from SDL_OpenJoystick().
//
// Returns the product version of the selected joystick, or 0 if unavailable.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickProductVersion
func (joystick *Joystick) ProductVersion() uint16 {
	return (uint16)(C.SDL_GetJoystickProductVersion((*C.SDL_Joystick)(joystick)))
}

// Get the firmware version of an opened joystick, if available.
//
// If the firmware version isn't available this function returns 0.
//
// joystick: the SDL_Joystick obtained from SDL_OpenJoystick().
//
// Returns the firmware version of the selected joystick, or 0 if
// unavailable.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickFirmwareVersion
func (joystick *Joystick) FirmwareVersion() uint16 {
	return (uint16)(C.SDL_GetJoystickFirmwareVersion((*C.SDL_Joystick)(joystick)))
}

// Get the serial number of an opened joystick, if available.
//
// Returns the serial number of the joystick, or NULL if it is not available.
//
// joystick: the SDL_Joystick obtained from SDL_OpenJoystick().
//
// Returns the serial number of the selected joystick, or NULL if
// unavailable.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickSerial
func (joystick *Joystick) Serial() string {
	return C.GoString(C.SDL_GetJoystickSerial((*C.SDL_Joystick)(joystick)))
}

// Get the type of an opened joystick.
//
// joystick: the SDL_Joystick obtained from SDL_OpenJoystick().
//
// Returns the SDL_JoystickType of the selected joystick.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickType
func (joystick *Joystick) Type() JoystickType {
	return (JoystickType)(C.SDL_GetJoystickType((*C.SDL_Joystick)(joystick)))
}

// Get the device information encoded in a SDL_GUID structure.
//
// guid: the SDL_GUID you wish to get info about.
//
// vendor: a pointer filled in with the device VID, or 0 if not
// available.
//
// product: a pointer filled in with the device PID, or 0 if not
// available.
//
// version: a pointer filled in with the device version, or 0 if not
// available.
//
// crc16: a pointer filled in with a CRC used to distinguish different
// products with the same VID/PID, or 0 if not available.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickGUIDInfo
func GetJoystickGUIDInfo(guid GUID) (vendor, product, version, crc16 uint16) {
	C.SDL_GetJoystickGUIDInfo(C.SDL_GUID{guid}, (*C.Uint16)(&vendor), (*C.Uint16)(&product), (*C.Uint16)(&version), (*C.Uint16)(&crc16))
	return
}

// Get the status of a specified joystick.
//
// joystick: the joystick to query.
//
// Returns true if the joystick has been opened, false if it has not; call
// SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_JoystickConnected
func (joystick *Joystick) Connected() bool {
	return (bool)(C.SDL_JoystickConnected((*C.SDL_Joystick)(joystick)))
}

// Get the instance ID of an opened joystick.
//
// joystick: an SDL_Joystick structure containing joystick information.
//
// Returns the instance ID of the specified joystick on success or 0 on
// failure; call SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickID
func (joystick *Joystick) ID() (JoystickID, error) {
	id := (JoystickID)(C.SDL_GetJoystickID((*C.SDL_Joystick)(joystick)))
	if id == 0 {
		return 0, getError()
	}
	return id, nil
}

// Get the number of general axis controls on a joystick.
//
// Often, the directional pad on a game controller will either look like 4
// separate buttons or a POV hat, and not axes, but all of this is up to the
// device and platform.
//
// joystick: an SDL_Joystick structure containing joystick information.
//
// Returns the number of axis controls/number of axes on success or -1 on
// failure; call SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickAxes
func (joystick *Joystick) NumAxes() (int, error) {
	num := (int)(C.SDL_GetNumJoystickAxes((*C.SDL_Joystick)(joystick)))
	if num == -1 {
		return -1, getError()
	}
	return num, nil
}

// Get the number of trackballs on a joystick.
//
// Joystick trackballs have only relative motion events associated with them
// and their state cannot be polled.
//
// Most joysticks do not have trackballs.
//
// joystick: an SDL_Joystick structure containing joystick information.
//
// Returns the number of trackballs on success or -1 on failure; call
// SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickBalls
func (joystick *Joystick) NumBalls() (int, error) {
	num := (int)(C.SDL_GetNumJoystickBalls((*C.SDL_Joystick)(joystick)))
	if num == -1 {
		return -1, getError()
	}
	return num, nil
}

// Get the number of POV hats on a joystick.
//
// joystick: an SDL_Joystick structure containing joystick information.
//
// Returns the number of POV hats on success or -1 on failure; call
// SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickHats
func (joystick *Joystick) NumHats() (int, error) {
	num := (int)(C.SDL_GetNumJoystickHats((*C.SDL_Joystick)(joystick)))
	if num == -1 {
		return -1, getError()
	}
	return num, nil
}

// Get the number of buttons on a joystick.
//
// joystick: an SDL_Joystick structure containing joystick information.
//
// Returns the number of buttons on success or -1 on failure; call
// SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetNumJoystickButtons
func (joystick *Joystick) NumButtons() (int, error) {
	num := (int)(C.SDL_GetNumJoystickButtons((*C.SDL_Joystick)(joystick)))
	if num == -1 {
		return -1, getError()
	}
	return num, nil
}

// Set the state of joystick event processing.
//
// If joystick events are disabled, you must call SDL_UpdateJoysticks()
// yourself and check the state of the joystick when you want joystick
// information.
//
// enabled: whether to process joystick events or not.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetJoystickEventsEnabled
func SetJoystickEventsEnabled(enabled bool) {
	C.SDL_SetJoystickEventsEnabled((C.bool)(enabled))
}

// Query the state of joystick event processing.
//
// If joystick events are disabled, you must call SDL_UpdateJoysticks()
// yourself and check the state of the joystick when you want joystick
// information.
//
// Returns true if joystick events are being processed, false otherwise.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_JoystickEventsEnabled
func JoystickEventsEnabled() bool {
	return (bool)(C.SDL_JoystickEventsEnabled())
}

// Update the current state of the open joysticks.
//
// This is called automatically by the event loop if any joystick events are
// enabled.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_UpdateJoysticks
func UpdateJoysticks() {
	C.SDL_UpdateJoysticks()
}

// Get the current state of an axis control on a joystick.
//
// SDL makes no promises about what part of the joystick any given axis refers
// to. Your game should have some sort of configuration UI to let users
// specify what each axis should be bound to. Alternately, SDL's higher-level
// Game Controller API makes a great effort to apply order to this lower-level
// interface, so you know that a specific axis is the "left thumb stick," etc.
//
// The value returned by SDL_GetJoystickAxis() is a signed integer (-32768 to
// 32767) representing the current position of the axis. It may be necessary
// to impose certain tolerances on these values to account for jitter.
//
// joystick: an SDL_Joystick structure containing joystick information.
//
// axis: the axis to query; the axis indices start at index 0.
//
// Returns a 16-bit signed integer representing the current position of the
// axis or 0 on failure; call SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickAxis
func (joystick *Joystick) Axis(axis int) int16 {
	return (int16)(C.SDL_GetJoystickAxis((*C.SDL_Joystick)(joystick), (C.int)(axis)))
}

// Get the initial state of an axis control on a joystick.
//
// The state is a value ranging from -32768 to 32767.
//
// The axis indices start at index 0.
//
// joystick: an SDL_Joystick structure containing joystick information.
//
// axis: the axis to query; the axis indices start at index 0.
//
// state: upon return, the initial value is supplied here.
//
// Returns true if this axis has any initial value, or false if not.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickAxisInitialState
func (joystick *Joystick) AxisInitialState(axis int) (state int16, hasInitialValue bool) {
	hasInitialValue = (bool)(C.SDL_GetJoystickAxisInitialState((*C.SDL_Joystick)(joystick), (C.int)(axis), (*C.Sint16)(&state)))
	return
}

// Get the ball axis change since the last poll.
//
// Trackballs can only return relative motion since the last call to
// SDL_GetJoystickBall(), these motion deltas are placed into `dx` and `dy`.
//
// Most joysticks do not have trackballs.
//
// joystick: the SDL_Joystick to query.
//
// ball: the ball index to query; ball indices start at index 0.
//
// dx: stores the difference in the x axis position since the last poll.
//
// dy: stores the difference in the y axis position since the last poll.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickBall
func (joystick *Joystick) Ball(ball int) (dx, dy int, err error) {
	var cdx, cdy C.int
	if !C.SDL_GetJoystickBall((*C.SDL_Joystick)(joystick), (C.int)(ball), &cdx, &cdy) {
		return 0, 0, getError()
	}
	return int(cdx), int(cdy), nil
}

// Get the current state of a POV hat on a joystick.
//
// The returned value will be one of the `SDL_HAT_*` values.
//
// joystick: an SDL_Joystick structure containing joystick information.
//
// hat: the hat index to get the state from; indices start at index 0.
//
// Returns the current hat position.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickHat
func (joystick *Joystick) Hat(hat int) byte {
	return (byte)(C.SDL_GetJoystickHat((*C.SDL_Joystick)(joystick), (C.int)(hat)))
}

const (
	HatCentered  = 0x00
	HatUp        = 0x01
	HatRight     = 0x02
	HatDown      = 0x04
	HatLeft      = 0x08
	HatRightUp   = HatRight | HatUp
	HatRightDown = HatRight | HatDown
	HatLeftUp    = HatLeft | HatUp
	HatLeftDown  = HatLeft | HatDown
)

// Get the current state of a button on a joystick.
//
// joystick: an SDL_Joystick structure containing joystick information.
//
// button: the button index to get the state from; indices start at
// index 0.
//
// Returns true if the button is pressed, false otherwise.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickButton
func (joystick *Joystick) Button(button int) bool {
	return (bool)(C.SDL_GetJoystickButton((*C.SDL_Joystick)(joystick), (C.int)(button)))
}

// Start a rumble effect.
//
// Each call to this function cancels any previous rumble effect, and calling
// it with 0 intensity stops any rumbling.
//
// This function requires you to process SDL events or call
// SDL_UpdateJoysticks() to update rumble state.
//
// joystick: the joystick to vibrate.
//
// low_frequency_rumble: the intensity of the low frequency (left)
// rumble motor, from 0 to 0xFFFF.
//
// high_frequency_rumble: the intensity of the high frequency (right)
// rumble motor, from 0 to 0xFFFF.
//
// duration_ms: the duration of the rumble effect, in milliseconds.
//
// Returns true, or false if rumble isn't supported on this joystick.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RumbleJoystick
func (joystick *Joystick) Rumble(lowFrequencyRumble, highFrequencyRumble, durationMS uint32) bool {
	return (bool)(C.SDL_RumbleJoystick((*C.SDL_Joystick)(joystick), (C.Uint16)(lowFrequencyRumble), (C.Uint16)(highFrequencyRumble), (C.Uint32)(durationMS)))
}

// Start a rumble effect in the joystick's triggers.
//
// Each call to this function cancels any previous trigger rumble effect, and
// calling it with 0 intensity stops any rumbling.
//
// Note that this is rumbling of the _triggers_ and not the game controller as
// a whole. This is currently only supported on Xbox One controllers. If you
// want the (more common) whole-controller rumble, use SDL_RumbleJoystick()
// instead.
//
// This function requires you to process SDL events or call
// SDL_UpdateJoysticks() to update rumble state.
//
// joystick: the joystick to vibrate.
//
// left_rumble: the intensity of the left trigger rumble motor, from 0
// to 0xFFFF.
//
// right_rumble: the intensity of the right trigger rumble motor, from 0
// to 0xFFFF.
//
// duration_ms: the duration of the rumble effect, in milliseconds.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RumbleJoystickTriggers
func (joystick *Joystick) RumbleTriggers(leftRumble, rightRumble, durationMS uint32) error {
	if !C.SDL_RumbleJoystickTriggers((*C.SDL_Joystick)(joystick), (C.Uint16)(leftRumble), (C.Uint16)(rightRumble), (C.Uint32)(durationMS)) {
		return getError()
	}
	return nil
}

// Update a joystick's LED color.
//
// An example of a joystick LED is the light on the back of a PlayStation 4's
// DualShock 4 controller.
//
// For joysticks with a single color LED, the maximum of the RGB values will
// be used as the LED brightness.
//
// joystick: the joystick to update.
//
// red: the intensity of the red LED.
//
// green: the intensity of the green LED.
//
// blue: the intensity of the blue LED.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetJoystickLED
func (joystick *Joystick) SetLED(red byte, green byte, blue byte) error {
	if !C.SDL_SetJoystickLED((*C.SDL_Joystick)(joystick), (C.Uint8)(red), (C.Uint8)(green), (C.Uint8)(blue)) {
		return getError()
	}
	return nil
}

// Send a joystick specific effect packet.
//
// joystick: the joystick to affect.
//
// data: the data to send to the joystick.
//
// size: the size of the data to send to the joystick.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SendJoystickEffect
func (joystick *Joystick) SendEffect(data []byte) error {
	if !C.SDL_SendJoystickEffect((*C.SDL_Joystick)(joystick), unsafe.Pointer(unsafe.SliceData(data)), (C.int)(len(data))) {
		return getError()
	}
	return nil
}

// Close a joystick previously opened with SDL_OpenJoystick().
//
// joystick: the joystick device to close.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CloseJoystick
func (joystick *Joystick) Close() {
	C.SDL_CloseJoystick((*C.SDL_Joystick)(joystick))
}

// Get the connection state of a joystick.
//
// joystick: the joystick to query.
//
// Returns the connection state on success or
// `SDL_JOYSTICK_CONNECTION_INVALID` on failure; call SDL_GetError()
// for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickConnectionState
func (joystick *Joystick) ConnectionState() (JoystickConnectionState, error) {
	state := (JoystickConnectionState)(C.SDL_GetJoystickConnectionState((*C.SDL_Joystick)(joystick)))
	if state == JoystickConnectionInvalid {
		return JoystickConnectionInvalid, getError()
	}
	return state, nil
}

// Get the battery state of a joystick.
//
// You should never take a battery status as absolute truth. Batteries
// (especially failing batteries) are delicate hardware, and the values
// reported here are best estimates based on what that hardware reports. It's
// not uncommon for older batteries to lose stored power much faster than it
// reports, or completely drain when reporting it has 20 percent left, etc.
//
// joystick: the joystick to query.
//
// percent: a pointer filled in with the percentage of battery life
// left, between 0 and 100, or NULL to ignore. This will be
// filled in with -1 we can't determine a value or there is no
// battery.
//
// Returns the current battery state or `SDL_POWERSTATE_ERROR` on failure;
// call SDL_GetError() for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetJoystickPowerInfo
func (joystick *Joystick) PowerInfo() (state PowerState, percent int, err error) {
	var cpercent C.int
	state = (PowerState)(C.SDL_GetJoystickPowerInfo((*C.SDL_Joystick)(joystick), &cpercent))
	if state == PowerstateError {
		return PowerstateError, -1, getError()
	}
	return state, int(cpercent), nil
}
