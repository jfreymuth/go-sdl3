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
// #cgo noescape SDL_AddGamepadMapping
// #cgo nocallback SDL_AddGamepadMapping
// #cgo noescape SDL_AddGamepadMappingsFromIO
// #cgo noescape SDL_AddGamepadMappingsFromFile
// #cgo nocallback SDL_AddGamepadMappingsFromFile
// #cgo noescape SDL_ReloadGamepadMappings
// #cgo nocallback SDL_ReloadGamepadMappings
// #cgo noescape SDL_GetGamepadMappings
// #cgo nocallback SDL_GetGamepadMappings
// #cgo noescape SDL_GetGamepadMappingForGUID
// #cgo nocallback SDL_GetGamepadMappingForGUID
// #cgo noescape SDL_GetGamepadMapping
// #cgo nocallback SDL_GetGamepadMapping
// #cgo noescape SDL_SetGamepadMapping
// #cgo nocallback SDL_SetGamepadMapping
// #cgo noescape SDL_HasGamepad
// #cgo nocallback SDL_HasGamepad
// #cgo noescape SDL_GetGamepads
// #cgo nocallback SDL_GetGamepads
// #cgo noescape SDL_IsGamepad
// #cgo nocallback SDL_IsGamepad
// #cgo noescape SDL_GetGamepadNameForID
// #cgo nocallback SDL_GetGamepadNameForID
// #cgo noescape SDL_GetGamepadPathForID
// #cgo nocallback SDL_GetGamepadPathForID
// #cgo noescape SDL_GetGamepadPlayerIndexForID
// #cgo nocallback SDL_GetGamepadPlayerIndexForID
// #cgo noescape SDL_GetGamepadGUIDForID
// #cgo nocallback SDL_GetGamepadGUIDForID
// #cgo noescape SDL_GetGamepadVendorForID
// #cgo nocallback SDL_GetGamepadVendorForID
// #cgo noescape SDL_GetGamepadProductForID
// #cgo nocallback SDL_GetGamepadProductForID
// #cgo noescape SDL_GetGamepadProductVersionForID
// #cgo nocallback SDL_GetGamepadProductVersionForID
// #cgo noescape SDL_GetGamepadTypeForID
// #cgo nocallback SDL_GetGamepadTypeForID
// #cgo noescape SDL_GetRealGamepadTypeForID
// #cgo nocallback SDL_GetRealGamepadTypeForID
// #cgo noescape SDL_GetGamepadMappingForID
// #cgo nocallback SDL_GetGamepadMappingForID
// #cgo noescape SDL_OpenGamepad
// #cgo nocallback SDL_OpenGamepad
// #cgo noescape SDL_GetGamepadFromID
// #cgo nocallback SDL_GetGamepadFromID
// #cgo noescape SDL_GetGamepadFromPlayerIndex
// #cgo nocallback SDL_GetGamepadFromPlayerIndex
// #cgo noescape SDL_GetGamepadProperties
// #cgo nocallback SDL_GetGamepadProperties
// #cgo noescape SDL_GetGamepadID
// #cgo nocallback SDL_GetGamepadID
// #cgo noescape SDL_GetGamepadName
// #cgo nocallback SDL_GetGamepadName
// #cgo noescape SDL_GetGamepadPath
// #cgo nocallback SDL_GetGamepadPath
// #cgo noescape SDL_GetGamepadType
// #cgo nocallback SDL_GetGamepadType
// #cgo noescape SDL_GetRealGamepadType
// #cgo nocallback SDL_GetRealGamepadType
// #cgo noescape SDL_GetGamepadPlayerIndex
// #cgo nocallback SDL_GetGamepadPlayerIndex
// #cgo noescape SDL_SetGamepadPlayerIndex
// #cgo nocallback SDL_SetGamepadPlayerIndex
// #cgo noescape SDL_GetGamepadVendor
// #cgo nocallback SDL_GetGamepadVendor
// #cgo noescape SDL_GetGamepadProduct
// #cgo nocallback SDL_GetGamepadProduct
// #cgo noescape SDL_GetGamepadProductVersion
// #cgo nocallback SDL_GetGamepadProductVersion
// #cgo noescape SDL_GetGamepadFirmwareVersion
// #cgo nocallback SDL_GetGamepadFirmwareVersion
// #cgo noescape SDL_GetGamepadSerial
// #cgo nocallback SDL_GetGamepadSerial
// #cgo noescape SDL_GetGamepadSteamHandle
// #cgo nocallback SDL_GetGamepadSteamHandle
// #cgo noescape SDL_GetGamepadConnectionState
// #cgo nocallback SDL_GetGamepadConnectionState
// #cgo noescape SDL_GetGamepadPowerInfo
// #cgo nocallback SDL_GetGamepadPowerInfo
// #cgo noescape SDL_GamepadConnected
// #cgo nocallback SDL_GamepadConnected
// #cgo noescape SDL_GetGamepadJoystick
// #cgo nocallback SDL_GetGamepadJoystick
// #cgo noescape SDL_SetGamepadEventsEnabled
// #cgo nocallback SDL_SetGamepadEventsEnabled
// #cgo noescape SDL_GamepadEventsEnabled
// #cgo nocallback SDL_GamepadEventsEnabled
// #cgo noescape SDL_UpdateGamepads
// #cgo nocallback SDL_UpdateGamepads
// #cgo noescape SDL_GetGamepadTypeFromString
// #cgo nocallback SDL_GetGamepadTypeFromString
// #cgo noescape SDL_GetGamepadStringForType
// #cgo nocallback SDL_GetGamepadStringForType
// #cgo noescape SDL_GetGamepadAxisFromString
// #cgo nocallback SDL_GetGamepadAxisFromString
// #cgo noescape SDL_GetGamepadStringForAxis
// #cgo nocallback SDL_GetGamepadStringForAxis
// #cgo noescape SDL_GamepadHasAxis
// #cgo nocallback SDL_GamepadHasAxis
// #cgo noescape SDL_GetGamepadAxis
// #cgo nocallback SDL_GetGamepadAxis
// #cgo noescape SDL_GetGamepadButtonFromString
// #cgo nocallback SDL_GetGamepadButtonFromString
// #cgo noescape SDL_GetGamepadStringForButton
// #cgo nocallback SDL_GetGamepadStringForButton
// #cgo noescape SDL_GamepadHasButton
// #cgo nocallback SDL_GamepadHasButton
// #cgo noescape SDL_GetGamepadButton
// #cgo nocallback SDL_GetGamepadButton
// #cgo noescape SDL_GetGamepadButtonLabelForType
// #cgo nocallback SDL_GetGamepadButtonLabelForType
// #cgo noescape SDL_GetGamepadButtonLabel
// #cgo nocallback SDL_GetGamepadButtonLabel
// #cgo noescape SDL_GetNumGamepadTouchpads
// #cgo nocallback SDL_GetNumGamepadTouchpads
// #cgo noescape SDL_GetNumGamepadTouchpadFingers
// #cgo nocallback SDL_GetNumGamepadTouchpadFingers
// #cgo noescape SDL_GetGamepadTouchpadFinger
// #cgo nocallback SDL_GetGamepadTouchpadFinger
// #cgo noescape SDL_GamepadHasSensor
// #cgo nocallback SDL_GamepadHasSensor
// #cgo noescape SDL_SetGamepadSensorEnabled
// #cgo noescape SDL_GamepadSensorEnabled
// #cgo nocallback SDL_GamepadSensorEnabled
// #cgo noescape SDL_GetGamepadSensorDataRate
// #cgo nocallback SDL_GetGamepadSensorDataRate
// #cgo noescape SDL_GetGamepadSensorData
// #cgo nocallback SDL_GetGamepadSensorData
// #cgo noescape SDL_RumbleGamepad
// #cgo noescape SDL_RumbleGamepadTriggers
// #cgo noescape SDL_SetGamepadLED
// #cgo noescape SDL_SendGamepadEffect
// #cgo noescape SDL_CloseGamepad
// #cgo noescape SDL_GetGamepadAppleSFSymbolsNameForButton
// #cgo nocallback SDL_GetGamepadAppleSFSymbolsNameForButton
// #cgo noescape SDL_GetGamepadAppleSFSymbolsNameForAxis
// #cgo nocallback SDL_GetGamepadAppleSFSymbolsNameForAxis
// #include <SDL3/SDL.h>
import "C"
import (
	"unsafe"
)

// # CategoryGamepad
//
// SDL provides a low-level joystick API, which just treats joysticks as an
// arbitrary pile of buttons, axes, and hat switches. If you're planning to
// write your own control configuration screen, this can give you a lot of
// flexibility, but that's a lot of work, and most things that we consider
// "joysticks" now are actually console-style gamepads. So SDL provides the
// gamepad API on top of the lower-level joystick functionality.
//
// The difference betweena joystick and a gamepad is that a gamepad tells you
// _where_ a button or axis is on the device. You don't speak to gamepads in
// terms of arbitrary numbers like "button 3" or "axis 2" but in standard
// locations: the d-pad, the shoulder buttons, triggers, A/B/X/Y (or
// X/O/Square/Triangle, if you will).
//
// One turns a joystick into a gamepad by providing a magic configuration
// string, which tells SDL the details of a specific device: when you see this
// specific hardware, if button 2 gets pressed, this is actually D-Pad Up,
// etc.
//
// SDL has many popular controllers configured out of the box, and users can
// add their own controller details through an environment variable if it's
// otherwise unknown to SDL.
//
// In order to use these functions, [Init]() must have been called with the
// [InitGamepad] flag. This causes SDL to scan the system for gamepads, and
// load appropriate drivers.
//
// If you would like to receive gamepad updates while the application is in
// the background, you should set the following hint before calling
// [Init]: [HintJoystickAllowBackgroundEvents]
//
// Gamepads support various optional features such as rumble, color LEDs,
// touchpad, gyro, etc. The support for these features varies depending on the
// controller and OS support available. You can check for LED and rumble
// capabilities at runtime by calling [Gamepad.Properties] and checking
// the various capability properties. You can check for touchpad by calling
// [Gamepad.NumTouchpads] and check for gyro and accelerometer by
// calling [Gamepad.HasSensor].
//
// By default SDL will try to use the most capable driver available, but you
// can tune which OS drivers to use with the various joystick hints in
// SDL_hints.h.
//
// Your application should always support gamepad hotplugging. On some
// platforms like Xbox, Steam Deck, etc., this is a requirement for
// certification. On other platforms, like macOS and Windows when using
// Windows.Gaming.Input, controllers may not be available at startup and will
// come in at some point after you've started processing events.

// The structure used to identify an SDL gamepad
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Gamepad
type Gamepad C.struct_SDL_Gamepad

// Standard gamepad types.
//
// This type does not necessarily map to first-party controllers from
// Microsoft/Sony/Nintendo; in many cases, third-party controllers can report
// as these, either because they were designed for a specific console, or they
// simply most closely match that console's controllers (does it have A/B/X/Y
// buttons or X/O/Square/Triangle? Does it have a touchpad? etc).
//
// https://wiki.libsdl.org/SDL3/SDL_GamepadType
type GamepadType uint32

const (
	GamepadTypeUnknown GamepadType = iota
	GamepadTypeStandard
	GamepadTypeXBOX360
	GamepadTypeXboxone
	GamepadTypePS3
	GamepadTypePS4
	GamepadTypePS5
	GamepadTypeNintendoSwitchPro
	GamepadTypeNintendoSwitchJoyconLeft
	GamepadTypeNintendoSwitchJoyconRight
	GamepadTypeNintendoSwitchJoyconPair
	GamepadTypeCount = iota
)

// The list of buttons available on a gamepad
//
// For controllers that use a diamond pattern for the face buttons, the
// south/east/west/north buttons below correspond to the locations in the
// diamond pattern. For Xbox controllers, this would be A/B/X/Y, for Nintendo
// Switch controllers, this would be B/A/Y/X, for PlayStation controllers this
// would be Cross/Circle/Square/Triangle.
//
// For controllers that don't use a diamond pattern for the face buttons, the
// south/east/west/north buttons indicate the buttons labeled A, B, C, D, or
// 1, 2, 3, 4, or for controllers that aren't labeled, they are the primary,
// secondary, etc. buttons.
//
// The activate action is often the south button and the cancel action is
// often the east button, but in some regions this is reversed, so your game
// should allow remapping actions based on user preferences.
//
// You can query the labels for the face buttons using
// [Gamepad.ButtonLabel]
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GamepadButton
type GamepadButton byte

const (
	GamepadButtonSouth GamepadButton = iota // Bottom face button (e.g. Xbox A button)
	GamepadButtonEast                       // Right face button (e.g. Xbox B button)
	GamepadButtonWest                       // Left face button (e.g. Xbox X button)
	GamepadButtonNorth                      // Top face button (e.g. Xbox Y button)
	GamepadButtonBack
	GamepadButtonGuide
	GamepadButtonStart
	GamepadButtonLeftStick
	GamepadButtonRightStick
	GamepadButtonLeftShoulder
	GamepadButtonRightShoulder
	GamepadButtonDpadUp
	GamepadButtonDpadDown
	GamepadButtonDpadLeft
	GamepadButtonDpadRight
	GamepadButtonMisc1        // Additional button (e.g. Xbox Series X share button, PS5 microphone button, Nintendo Switch Pro capture button, Amazon Luna microphone button, Google Stadia capture button)
	GamepadButtonRightPaddle1 // Upper or primary paddle, under your right hand (e.g. Xbox Elite paddle P1)
	GamepadButtonLeftPaddle1  // Upper or primary paddle, under your left hand (e.g. Xbox Elite paddle P3)
	GamepadButtonRightPaddle2 // Lower or secondary paddle, under your right hand (e.g. Xbox Elite paddle P2)
	GamepadButtonLeftPaddle2  // Lower or secondary paddle, under your left hand (e.g. Xbox Elite paddle P4)
	GamepadButtonTouchpad     // PS4/PS5 touchpad button
	GamepadButtonMisc2        // Additional button
	GamepadButtonMisc3        // Additional button
	GamepadButtonMisc4        // Additional button
	GamepadButtonMisc5        // Additional button
	GamepadButtonMisc6        // Additional button
	GamepadButtonCount        = iota

	GamepadButtonInvalid = 255
)

// The set of gamepad button labels
//
// This isn't a complete set, just the face buttons to make it easy to show
// button prompts.
//
// For a complete set, you should look at the button and gamepad type and have
// a set of symbols that work well with your art style.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GamepadButtonLabel
type GamepadButtonLabel uint32

const (
	GamepadButtonLabelUnknown GamepadButtonLabel = iota
	GamepadButtonLabelA
	GamepadButtonLabelB
	GamepadButtonLabelX
	GamepadButtonLabelY
	GamepadButtonLabelCross
	GamepadButtonLabelCircle
	GamepadButtonLabelSquare
	GamepadButtonLabelTriangle
)

// The list of axes available on a gamepad
//
// Thumbstick axis values range from [JoystickAxisMin] to
// [JoystickAxisMax], and are centered within ~8000 of zero, though
// advanced UI will allow users to set or autodetect the dead zone, which
// varies between gamepads.
//
// Trigger axis values range from 0 (released) to [JoystickAxisMax] (fully
// pressed) when reported by [Gamepad.Axis]. Note that this is not the
// same range that will be reported by the lower-level [Joystick.Axis].
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GamepadAxis
type GamepadAxis byte

const (
	GamepadAxisLeftX GamepadAxis = iota
	GamepadAxisLeftY
	GamepadAxisRightX
	GamepadAxisRightY
	GamepadAxisLeftTrigger
	GamepadAxisRightTrigger
	GamepadAxisCount = iota

	GamepadAxisInvalid GamepadAxis = 255
)

// Types of gamepad control bindings.
//
// A gamepad is a collection of bindings that map arbitrary joystick buttons,
// axes and hat switches to specific positions on a generic console-style
// gamepad. This enum is used as part of SDL_GamepadBinding to specify those
// mappings.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GamepadBindingType
//TODOtype GamepadBindingType uint32
//TODO
//TODOconst (
//TODO	GamepadBindtypeNone GamepadBindingType = 0
//TODO	GamepadBindtypeButton
//TODO	GamepadBindtypeAxis
//TODO	GamepadBindtypeHat
//TODO)

// A mapping between one joystick input to a gamepad control.
//
// A gamepad has a collection of several bindings, to say, for example, when
// joystick button number 5 is pressed, that should be treated like the
// gamepad's "start" button.
//
// SDL has these bindings built-in for many popular controllers, and can add
// more with a simple text string. Those strings are parsed into a collection
// of these structs to make it easier to operate on the data.
//
// This struct is available since SDL 3.2.0.
//TODOtype GamepadBinding struct {
//TODO	InputType  GamepadBindingType //
//TODO	Button     int                //
//TODO	Axis       int                //
//TODO	AxisMin    int                //
//TODO	AxisMax    int                //
//TODO	Hat        int                //
//TODO	HatMask    int                //
//TODO	OutputType GamepadBindingType //
//TODO	Button     GamepadButton      //
//TODO	Axis       GamepadAxis        //
//TODO	AxisMin    int                //
//TODO	AxisMax    int                //
//TODO}

// Add support for gamepads that SDL is unaware of or change the binding of an
// existing gamepad.
//
// The mapping string has the format "GUID,name,mapping", where GUID is the
// string value from [GUID.String], name is the human readable string for
// the device and mappings are gamepad mappings to joystick ones. Under
// Windows there is a reserved GUID of "xinput" that covers all XInput
// devices. The mapping format for joystick is:
//
//   - `bX`: a joystick button, index X
//   - `hX.Y`: hat X with value Y
//   - `aX`: axis X of the joystick
//
// Buttons can be used as a gamepad axes and vice versa.
//
// If a device with this GUID is already plugged in, SDL will generate an
// [EventGamepadAdded] event.
//
// This string shows an example of a valid mapping for a gamepad:
//
//	"341a3608000000000000504944564944,Afterglow PS3 Controller,a:b1,b:b2,y:b3,x:b0,start:b9,guide:b12,back:b8,dpup:h0.1,dpleft:h0.8,dpdown:h0.4,dpright:h0.2,leftshoulder:b4,rightshoulder:b5,leftstick:b10,rightstick:b11,leftx:a0,lefty:a1,rightx:a2,righty:a3,lefttrigger:b6,righttrigger:b7"
//
// mapping: the mapping string.
//
// Returns (true, nil) if a new mapping is added, (false, nil) if an existing
// mapping is updated, or a non-nil error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AddGamepadMapping
func AddGamepadMapping(mapping string) (new bool, err error) {
	r := C.SDL_AddGamepadMapping(tmpstring(mapping))
	if r == -1 {
		return false, getError()
	}
	return r == 1, nil
}

// Load a set of gamepad mappings from an [IOStream].
//
// You can call this function several times, if needed, to load different
// database files.
//
// If a new mapping is loaded for an already known gamepad GUID, the later
// version will overwrite the one currently loaded.
//
// Any new mappings for already plugged in controllers will generate
// [EventGamepadAdded] events.
//
// Mappings not belonging to the current platform or with no platform field
// specified will be ignored (i.e. mappings for Linux will be ignored in
// Windows, etc).
//
// This function will load the text database entirely in memory before
// processing it, so take this into consideration if you are in a memory
// constrained environment.
//
// src: the data stream for the mappings to be added.
//
// closeio: if true, calls [CloseIO] on src before returning, even
// in the case of an error.
//
// Returns the number of mappings added or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
func AddGamepadMappingsFromIO(src *IOStream, closeio bool) (int, error) {
	n := (int)(C.SDL_AddGamepadMappingsFromIO((*C.SDL_IOStream)(src), (C.bool)(closeio)))
	if n == -1 {
		return 0, getError()
	}
	return n, nil
}

// Load a set of gamepad mappings from a file.
//
// You can call this function several times, if needed, to load different
// database files.
//
// If a new mapping is loaded for an already known gamepad GUID, the later
// version will overwrite the one currently loaded.
//
// Any new mappings for already plugged in controllers will generate
// [EventGamepadAdded] events.
//
// Mappings not belonging to the current platform or with no platform field
// specified will be ignored (i.e. mappings for Linux will be ignored in
// Windows, etc).
//
// file: the mappings file to load.
//
// Returns the number of mappings added or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AddGamepadMappingsFromFile
func AddGamepadMappingsFromFile(file string) (int, error) {
	n := (int)(C.SDL_AddGamepadMappingsFromFile(tmpstring(file)))
	if n == -1 {
		return 0, getError()
	}
	return n, nil
}

// Reinitialize the SDL mapping database to its initial state.
//
// This will generate gamepad events as needed if device mappings change.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReloadGamepadMappings
func ReloadGamepadMappings() bool {
	return (bool)(C.SDL_ReloadGamepadMappings())
}

// Get the current gamepad mappings.
//
// Returns an array of the mapping strings or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappings
func GetGamepadMappings() ([]string, error) {
	var count C.int
	mappings := C.SDL_GetGamepadMappings(&count)
	if mappings == nil {
		return nil, getError()
	}
	result := make([]string, count)
	for i, m := range unsafe.Slice(mappings, count) {
		result[i] = C.GoString(m)
	}
	C.SDL_free(unsafe.Pointer(mappings))
	return result, nil
}

// Get the gamepad mapping string for a given [GUID].
//
// guid: a structure containing the GUID for which a mapping is desired.
//
// Returns a mapping string an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappingForGUID
func GetGamepadMappingForGUID(guid GUID) (string, error) {
	mapping := C.SDL_GetGamepadMappingForGUID(C.SDL_GUID{guid})
	if mapping == nil {
		return "", getError()
	}
	result := C.GoString(mapping)
	C.SDL_free(unsafe.Pointer(mapping))
	return result, nil
}

// Get the current mapping of a gamepad.
//
// Details about mappings are discussed with [AddGamepadMapping].
//
// gamepad: the gamepad you want to get the current mapping for.
//
// Returns a string that has the gamepad's mapping or an error if no mapping is
// available.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadMapping
func (gamepad *Gamepad) Mapping() (string, error) {
	mapping := C.SDL_GetGamepadMapping((*C.SDL_Gamepad)(gamepad))
	if mapping == nil {
		return "", getError()
	}
	result := C.GoString(mapping)
	C.SDL_free(unsafe.Pointer(mapping))
	return result, nil
}

// Set the current mapping of a joystick or gamepad.
//
// Details about mappings are discussed with [AddGamepadMapping].
//
// id: the joystick instance ID.
//
// mapping: the mapping to use for this device, or NULL to clear the
// mapping.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetGamepadMapping
func (id JoystickID) SetGamepadMapping(mapping string) error {
	if !C.SDL_SetGamepadMapping((C.SDL_JoystickID)(id), tmpstring(mapping)) {
		return getError()
	}
	return nil
}

// Return whether a gamepad is currently connected.
//
// Returns true if a gamepad is connected, false otherwise.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasGamepad
func HasGamepad() bool {
	return (bool)(C.SDL_HasGamepad())
}

// Get a list of currently connected gamepads.
//
// Returns a slice of joystick instance IDs or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepads
func GetGamepads() ([]JoystickID, error) {
	var count C.int
	g := C.SDL_GetGamepads(&count)
	if g == nil {
		return nil, getError()
	}
	result := make([]JoystickID, count)
	for i, g := range unsafe.Slice(g, count) {
		result[i] = JoystickID(g)
	}
	C.SDL_free(unsafe.Pointer(g))
	return result, nil
}

// Check if the given joystick is supported by the gamepad interface.
//
// id: the joystick instance ID.
//
// Returns true if the given joystick is supported by the gamepad interface,
// false if it isn't or it's an invalid index.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_IsGamepad
func (id JoystickID) IsGamepad() bool {
	return (bool)(C.SDL_IsGamepad((C.SDL_JoystickID)(id)))
}

// Get the implementation dependent name of a gamepad.
//
// This can be called before any gamepads are opened.
//
// id: the joystick instance ID.
//
// Returns the name of the selected gamepad or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadNameForID
func (id JoystickID) GamepadName() (string, error) {
	name := C.SDL_GetGamepadNameForID((C.SDL_JoystickID)(id))
	if name == nil {
		return "", getError()
	}
	return C.GoString(name), nil
}

// Get the implementation dependent path of a gamepad.
//
// This can be called before any gamepads are opened.
//
// id: the joystick instance ID.
//
// Returns the path of the selected gamepad or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadPathForID
func (id JoystickID) GamepadPath() (string, error) {
	name := C.SDL_GetGamepadPathForID((C.SDL_JoystickID)(id))
	if name == nil {
		return "", getError()
	}
	return C.GoString(name), nil
}

// Get the player index of a gamepad.
//
// This can be called before any gamepads are opened.
//
// id: the joystick instance ID.
//
// Returns the player index of a gamepad, or -1 if it's not available.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadPlayerIndexForID
func (id JoystickID) GamepadPlayerIndex() int {
	return (int)(C.SDL_GetGamepadPlayerIndexForID((C.SDL_JoystickID)(id)))
}

// Get the implementation-dependent GUID of a gamepad.
//
// This can be called before any gamepads are opened.
//
// id: the joystick instance ID.
//
// Returns the GUID of the selected gamepad. If called on an invalid index,
// this function returns a zero GUID.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadGUIDForID
func (id JoystickID) GamepadGUID() GUID {
	return (GUID)(C.SDL_GetGamepadGUIDForID((C.SDL_JoystickID)(id)).data)
}

// Get the USB vendor ID of a gamepad, if available.
//
// This can be called before any gamepads are opened. If the vendor ID isn't
// available this function returns 0.
//
// id: the joystick instance ID.
//
// Returns the USB vendor ID of the selected gamepad. If called on an invalid
// index, this function returns zero.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadVendorForID
func (id JoystickID) GamepadVendor() uint16 {
	return (uint16)(C.SDL_GetGamepadVendorForID((C.SDL_JoystickID)(id)))
}

// Get the USB product ID of a gamepad, if available.
//
// This can be called before any gamepads are opened. If the product ID isn't
// available this function returns 0.
//
// id: the joystick instance ID.
//
// Returns the USB product ID of the selected gamepad. If called on an
// invalid index, this function returns zero.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductForID
func (id JoystickID) GamepadProduct() uint16 {
	return (uint16)(C.SDL_GetGamepadProductForID((C.SDL_JoystickID)(id)))
}

// Get the product version of a gamepad, if available.
//
// This can be called before any gamepads are opened. If the product version
// isn't available this function returns 0.
//
// id: the joystick instance ID.
//
// Returns the product version of the selected gamepad. If called on an
// invalid index, this function returns zero.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductVersionForID
func (id JoystickID) GamepadProductVersion() uint16 {
	return (uint16)(C.SDL_GetGamepadProductVersionForID((C.SDL_JoystickID)(id)))
}

// Get the type of a gamepad.
//
// This can be called before any gamepads are opened.
//
// id: the joystick instance ID.
//
// Returns the gamepad type.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadTypeForID
func (id JoystickID) GamepadType() GamepadType {
	return (GamepadType)(C.SDL_GetGamepadTypeForID((C.SDL_JoystickID)(id)))
}

// Get the type of a gamepad, ignoring any mapping override.
//
// This can be called before any gamepads are opened.
//
// id: the joystick instance ID.
//
// Returns the gamepad type.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRealGamepadTypeForID
func (id JoystickID) RealGamepadType() GamepadType {
	return (GamepadType)(C.SDL_GetRealGamepadTypeForID((C.SDL_JoystickID)(id)))
}

// Get the mapping of a gamepad.
//
// This can be called before any gamepads are opened.
//
// id: the joystick instance ID.
//
// Returns the mapping string. Returns an empty string if no mapping is
// available.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadMappingForID
func (id JoystickID) GamepadMapping() string {
	m := C.SDL_GetGamepadMappingForID((C.SDL_JoystickID)(id))
	result := C.GoString(m)
	C.SDL_free(unsafe.Pointer(m))
	return result
}

// Open a gamepad for use.
//
// id: the joystick instance ID.
//
// Returns a gamepad identifier or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_OpenGamepad
func OpenGamepad(id JoystickID) (*Gamepad, error) {
	g := (*Gamepad)(C.SDL_OpenGamepad((C.SDL_JoystickID)(id)))
	if g == nil {
		return nil, getError()
	}
	return g, nil
}

// Get the [Gamepad] associated with a joystick instance ID, if it has been
// opened.
//
// id: the joystick instance ID of the gamepad.
//
// Returns a [Gamepad] on success or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadFromID
func (id JoystickID) Gamepad() (*Gamepad, error) {
	g := (*Gamepad)(C.SDL_GetGamepadFromID((C.SDL_JoystickID)(id)))
	if g == nil {
		return nil, getError()
	}
	return g, nil
}

// Get the [Gamepad] associated with a player index.
//
// playerIndex: the player index, which different from the instance ID.
//
// Returns the [Gamepad] associated with a player index.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadFromPlayerIndex
func GetGamepadFromPlayerIndex(playerIndex int32) (*Gamepad, error) {
	g := (*Gamepad)(C.SDL_GetGamepadFromPlayerIndex((C.int)(playerIndex)))
	if g == nil {
		return nil, getError()
	}
	return g, nil
}

// Get the properties associated with an opened gamepad.
//
// These properties are shared with the underlying joystick object.
//
// The following read-only properties are provided by SDL:
//
//   - [PropGamepadCapMonoLEDBoolean]: true if this gamepad has an LED
//     that has adjustable brightness
//   - [PropGamepadCapRGBLEDBoolean]: true if this gamepad has an LED
//     that has adjustable color
//   - [PropGamepadCapPlayerLEDBoolean]: true if this gamepad has a
//     player LED
//   - [PropGamepadCapRumbleBoolean]: true if this gamepad has
//     left/right rumble
//   - [PropGamepadCapTriggerRumbleBoolean]: true if this gamepad has
//     simple trigger rumble
//
// gamepad: a gamepad identifier previously returned by [OpenGamepad].
//
// Returns a valid property ID on success or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadProperties
func (gamepad *Gamepad) Properties() (PropertiesID, error) {
	props := (PropertiesID)(C.SDL_GetGamepadProperties((*C.SDL_Gamepad)(gamepad)))
	if props == 0 {
		return 0, getError()
	}
	return props, nil
}

const PropGamepadCapMonoLEDBoolean = PropJoystickCapMonoLEDBoolean
const PropGamepadCapRGBLEDBoolean = PropJoystickCapRGBLEDBoolean
const PropGamepadCapPlayerLEDBoolean = PropJoystickCapPlayerLEDBoolean
const PropGamepadCapRumbleBoolean = PropJoystickCapRumbleBoolean
const PropGamepadCapTriggerRumbleBoolean = PropJoystickCapTriggerRumbleBoolean

// Get the instance ID of an opened gamepad.
//
// gamepad: a gamepad identifier previously returned by [OpenGamepad].
//
// Returns the instance ID of the specified gamepad on success or an error on
// failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadID
func (gamepad *Gamepad) ID() (JoystickID, error) {
	id := (JoystickID)(C.SDL_GetGamepadID((*C.SDL_Gamepad)(gamepad)))
	if id == 0 {
		return 0, getError()
	}
	return id, nil
}

// Get the implementation-dependent name for an opened gamepad.
//
// gamepad: a gamepad identifier previously returned by [OpenGamepad].
//
// Returns the implementation dependent name for the gamepad, or an empty string
// if there is no name or the identifier passed is invalid.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadName
func (gamepad *Gamepad) Name() string {
	return C.GoString(C.SDL_GetGamepadName((*C.SDL_Gamepad)(gamepad)))
}

// Get the implementation-dependent path for an opened gamepad.
//
// gamepad: a gamepad identifier previously returned by [OpenGamepad].
//
// Returns the implementation dependent path for the gamepad, or NULL if
// there is no path or the identifier passed is invalid.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadPath
func (gamepad *Gamepad) Path() string {
	return C.GoString(C.SDL_GetGamepadPath((*C.SDL_Gamepad)(gamepad)))
}

// Get the type of an opened gamepad.
//
// gamepad: the gamepad object to query.
//
// Returns the gamepad type, or [GamepadTypeUnknown] if it's not
// available.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadType
func (gamepad *Gamepad) Type() GamepadType {
	return (GamepadType)(C.SDL_GetGamepadType((*C.SDL_Gamepad)(gamepad)))
}

// Get the type of an opened gamepad, ignoring any mapping override.
//
// gamepad: the gamepad object to query.
//
// Returns the gamepad type, or [GamepadTypeUnknown] if it's not
// available.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRealGamepadType
func (gamepad *Gamepad) RealType() GamepadType {
	return (GamepadType)(C.SDL_GetRealGamepadType((*C.SDL_Gamepad)(gamepad)))
}

// Get the player index of an opened gamepad.
//
// For XInput gamepads this returns the XInput user index.
//
// gamepad: the gamepad object to query.
//
// Returns the player index for gamepad, or -1 if it's not available.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadPlayerIndex
func (gamepad *Gamepad) PlayerIndex() int {
	return (int)(C.SDL_GetGamepadPlayerIndex((*C.SDL_Gamepad)(gamepad)))
}

// Set the player index of an opened gamepad.
//
// gamepad: the gamepad object to adjust.
//
// playerIndex: player index to assign to this gamepad, or -1 to clear
// the player index and turn off player LEDs.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetGamepadPlayerIndex
func (gamepad *Gamepad) SetPlayerIndex(playerIndex int) error {
	if !C.SDL_SetGamepadPlayerIndex((*C.SDL_Gamepad)(gamepad), (C.int)(playerIndex)) {
		return getError()
	}
	return nil
}

// Get the USB vendor ID of an opened gamepad, if available.
//
// If the vendor ID isn't available this function returns 0.
//
// gamepad: the gamepad object to query.
//
// Returns the USB vendor ID, or zero if unavailable.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadVendor
func (gamepad *Gamepad) Vendor() uint16 {
	return (uint16)(C.SDL_GetGamepadVendor((*C.SDL_Gamepad)(gamepad)))
}

// Get the USB product ID of an opened gamepad, if available.
//
// If the product ID isn't available this function returns 0.
//
// gamepad: the gamepad object to query.
//
// Returns the USB product ID, or zero if unavailable.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadProduct
func (gamepad *Gamepad) Product() uint16 {
	return (uint16)(C.SDL_GetGamepadProduct((*C.SDL_Gamepad)(gamepad)))
}

// Get the product version of an opened gamepad, if available.
//
// If the product version isn't available this function returns 0.
//
// gamepad: the gamepad object to query.
//
// Returns the USB product version, or zero if unavailable.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadProductVersion
func (gamepad *Gamepad) ProductVersion() uint16 {
	return (uint16)(C.SDL_GetGamepadProductVersion((*C.SDL_Gamepad)(gamepad)))
}

// Get the firmware version of an opened gamepad, if available.
//
// If the firmware version isn't available this function returns 0.
//
// gamepad: the gamepad object to query.
//
// Returns the gamepad firmware version, or zero if unavailable.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadFirmwareVersion
func (gamepad *Gamepad) FirmwareVersion() uint16 {
	return (uint16)(C.SDL_GetGamepadFirmwareVersion((*C.SDL_Gamepad)(gamepad)))
}

// Get the serial number of an opened gamepad, if available.
//
// Returns the serial number of the gamepad, or an empty string if it is not
// available.
//
// gamepad: the gamepad object to query.
//
// Returns the serial number, or an empty string if unavailable.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadSerial
func GetGamepadSerial(gamepad *Gamepad) string {
	return C.GoString(C.SDL_GetGamepadSerial((*C.SDL_Gamepad)(gamepad)))
}

// Get the Steam Input handle of an opened gamepad, if available.
//
// Returns an InputHandle_t for the gamepad that can be used with Steam Input
// API: https://partner.steamgames.com/doc/api/ISteamInput
//
// gamepad: the gamepad object to query.
//
// Returns the gamepad handle, or 0 if unavailable.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadSteamHandle
func (gamepad *Gamepad) SteamHandle() uint64 {
	return (uint64)(C.SDL_GetGamepadSteamHandle((*C.SDL_Gamepad)(gamepad)))
}

// Get the connection state of a gamepad.
//
// gamepad: the gamepad object to query.
//
// Returns the connection state on success or [JoystickConnectionInvalid] and
// an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadConnectionState
func (gamepad *Gamepad) ConnectionState() (JoystickConnectionState, error) {
	state := (JoystickConnectionState)(C.SDL_GetGamepadConnectionState((*C.SDL_Gamepad)(gamepad)))
	if state == JoystickConnectionInvalid {
		return JoystickConnectionInvalid, getError()
	}
	return state, nil
}

// Get the battery state of a gamepad.
//
// You should never take a battery status as absolute truth. Batteries
// (especially failing batteries) are delicate hardware, and the values
// reported here are best estimates based on what that hardware reports. It's
// not uncommon for older batteries to lose stored power much faster than it
// reports, or completely drain when reporting it has 20 percent left, etc.
//
// gamepad: the gamepad object to query.
//
// percent: the percentage of battery life left, between 0 and 100. This will be
// -1 we can't determine a value or there is no battery.
//
// Returns the current battery state.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadPowerInfo
func (gamepad *Gamepad) PowerInfo() (state PowerState, percent int) {
	var cpercent C.int
	state = (PowerState)(C.SDL_GetGamepadPowerInfo((*C.SDL_Gamepad)(gamepad), &cpercent))
	return state, int(cpercent)
}

// Check if a gamepad has been opened and is currently connected.
//
// gamepad: a gamepad identifier previously returned by [OpenGamepad].
//
// Returns true if the gamepad has been opened and is currently connected, or
// false if not.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GamepadConnected
func (gamepad *Gamepad) Connected() bool {
	return (bool)(C.SDL_GamepadConnected((*C.SDL_Gamepad)(gamepad)))
}

// Get the underlying joystick from a gamepad.
//
// This function will give you a [Joystick] object, which allows you to use
// the [Joystick] functions with a [Gamepad] object. This would be useful
// for getting a joystick's position at any given time, even if it hasn't
// moved (moving it would produce an event, which would have the axis' value).
//
// The pointer returned is owned by the [Gamepad]. You should not call
// [Joystick.Close] on it, for example, since doing so will likely cause
// SDL to crash.
//
// gamepad: the gamepad object that you want to get a joystick from.
//
// Returns a [Joystick] object or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadJoystick
func (gamepad *Gamepad) Joystick() (*Joystick, error) {
	j := (*Joystick)(C.SDL_GetGamepadJoystick((*C.SDL_Gamepad)(gamepad)))
	if j == nil {
		return nil, getError()
	}
	return j, nil
}

// Set the state of gamepad event processing.
//
// If gamepad events are disabled, you must call [UpdateGamepads] yourself
// and check the state of the gamepad when you want gamepad information.
//
// enabled: whether to process gamepad events or not.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetGamepadEventsEnabled
func SetGamepadEventsEnabled(enabled bool) {
	C.SDL_SetGamepadEventsEnabled((C.bool)(enabled))
}

// Query the state of gamepad event processing.
//
// If gamepad events are disabled, you must call [UpdateGamepads] yourself
// and check the state of the gamepad when you want gamepad information.
//
// Returns true if gamepad events are being processed, false otherwise.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GamepadEventsEnabled
func GamepadEventsEnabled() bool {
	return (bool)(C.SDL_GamepadEventsEnabled())
}

// Get the SDL joystick layer bindings for a gamepad.
//
// gamepad: a gamepad.
//
// count: a pointer filled in with the number of bindings returned.
//
// Returns a NULL terminated array of pointers to bindings or NULL on
// failure; call SDL_GetError() for more information. This is a
// single allocation that should be freed with SDL_free() when it is
// no longer needed.
//
// This function is available since SDL 3.2.0.
//TODO func GetGamepadBindings(gamepad *Gamepad, count *int32) **GamepadBinding {
//TODO 	return (**GamepadBinding)(C.SDL_GetGamepadBindings((*C.SDL_Gamepad)(gamepad), (*C.int)(count)))
//TODO }

// Manually pump gamepad updates if not using the loop.
//
// This function is called automatically by the event loop if events are
// enabled. Under such circumstances, it will not be necessary to call this
// function.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_UpdateGamepads
func UpdateGamepads() {
	C.SDL_UpdateGamepads()
}

// Convert a string into [GamepadType] enum.
//
// This function is called internally to translate [Gamepad] mapping strings
// for the underlying joystick device into the consistent [Gamepad] mapping.
// You do not normally need to call this function unless you are parsing
// [Gamepad] mappings in your own code.
//
// str: string representing a [GamepadType] type.
//
// Returns the [GamepadType] enum corresponding to the input string, or
// [GamepadTypeUnknown] if no match was found.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadTypeFromString
func GetGamepadTypeFromString(str string) GamepadType {
	return (GamepadType)(C.SDL_GetGamepadTypeFromString(tmpstring(str)))
}

// Convert from a [GamepadType] enum to a string.
//
// type: an enum value for a given [GamepadType].
//
// Returns a string for the given type, or an empty string if an invalid type is
// specified. The string returned is of the format used by [Gamepad] mapping
// strings.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForType
func (typ GamepadType) String() string {
	return C.GoString(C.SDL_GetGamepadStringForType((C.SDL_GamepadType)(typ)))
}

// Convert a string into [GamepadAxis] enum.
//
// This function is called internally to translate [Gamepad] mapping strings
// for the underlying joystick device into the consistent [Gamepad] mapping.
// You do not normally need to call this function unless you are parsing
// [Gamepad] mappings in your own code.
//
// Note specially that "righttrigger" and "lefttrigger" map to
// [GamepadAxisRightTrigger] and [GamepadAxisLeftTrigger],
// respectively.
//
// str: string representing a [Gamepad] axis.
//
// Returns the [GamepadAxis] enum corresponding to the input string, or
// [GamepadAxisInvalid] if no match was found.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadAxisFromString
func GetGamepadAxisFromString(str string) GamepadAxis {
	return (GamepadAxis)(C.SDL_GetGamepadAxisFromString(tmpstring(str)))
}

// Convert from a [GamepadAxis] enum to a string.
//
// axis: an enum value for a given [GamepadAxis].
//
// Returns a string for the given axis, or an empty string if an invalid axis is
// specified. The string returned is of the format used by [Gamepad] mapping
// strings.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForAxis
func (axis GamepadAxis) String() string {
	return C.GoString(C.SDL_GetGamepadStringForAxis((C.SDL_GamepadAxis)(int8(axis))))
}

// Query whether a gamepad has a given axis.
//
// This merely reports whether the gamepad's mapping defined this axis, as
// that is all the information SDL has about the physical device.
//
// gamepad: a gamepad.
//
// axis: an axis enum value (a [GamepadAxis] value).
//
// Returns true if the gamepad has this axis, false otherwise.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GamepadHasAxis
func (gamepad *Gamepad) HasAxis(axis GamepadAxis) bool {
	return (bool)(C.SDL_GamepadHasAxis((*C.SDL_Gamepad)(gamepad), (C.SDL_GamepadAxis)(int8(axis))))
}

// Get the current state of an axis control on a gamepad.
//
// The axis indices start at index 0.
//
// For thumbsticks, the state is a value ranging from -32768 (up/left) to
// 32767 (down/right).
//
// Triggers range from 0 when released to 32767 when fully pressed, and never
// return a negative value. Note that this differs from the value reported by
// the lower-level [Joystick.Axis], which normally uses the full range.
//
// gamepad: a gamepad.
//
// axis: an axis index (one of the [GamepadAxis] values).
//
// Returns axis state (including 0) on success or 0 (also) on failure; call
// [GetError] for more information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadAxis
func (gamepad *Gamepad) Axis(axis GamepadAxis) int16 {
	return (int16)(C.SDL_GetGamepadAxis((*C.SDL_Gamepad)(gamepad), (C.SDL_GamepadAxis)(int8(axis))))
}

// Convert a string into a [GamepadButton] enum.
//
// This function is called internally to translate [Gamepad] mapping strings
// for the underlying joystick device into the consistent [Gamepad] mapping.
// You do not normally need to call this function unless you are parsing
// [Gamepad] mappings in your own code.
//
// str: string representing a [Gamepad] axis.
//
// Returns the [GamepadButton] enum corresponding to the input string, or
// [GamepadButtonInvalid] if no match was found.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonFromString
func GetGamepadButtonFromString(str string) GamepadButton {
	return (GamepadButton)(C.SDL_GetGamepadButtonFromString(tmpstring(str)))
}

// Convert from a [GamepadButton] enum to a string.
//
// button: an enum value for a given [GamepadButton].
//
// Returns a string for the given button, or NULL if an invalid button is
// specified. The string returned is of the format used by [Gamepad] mapping
// strings.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadStringForButton
func (button GamepadButton) String() string {
	return C.GoString(C.SDL_GetGamepadStringForButton((C.SDL_GamepadButton)(int8(button))))
}

// Query whether a gamepad has a given button.
//
// This merely reports whether the gamepad's mapping defined this button, as
// that is all the information SDL has about the physical device.
//
// gamepad: a gamepad.
//
// button: a button enum value (a [GamepadButton] value).
//
// Returns true if the gamepad has this button, false otherwise.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GamepadHasButton
func (gamepad *Gamepad) HasButton(button GamepadButton) bool {
	return (bool)(C.SDL_GamepadHasButton((*C.SDL_Gamepad)(gamepad), (C.SDL_GamepadButton)(int8(button))))
}

// Get the current state of a button on a gamepad.
//
// gamepad: a gamepad.
//
// button: a button index (one of the [GamepadButton] values).
//
// Returns true if the button is pressed, false otherwise.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadButton
func (gamepad *Gamepad) Button(button GamepadButton) bool {
	return (bool)(C.SDL_GetGamepadButton((*C.SDL_Gamepad)(gamepad), (C.SDL_GamepadButton)(int8(button))))
}

// Get the label of a button on a gamepad.
//
// type: the type of gamepad to check.
//
// button: a button index (one of the [GamepadButton] values).
//
// Returns the [GamepadButtonLabel] enum corresponding to the button label.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonLabelForType
func GetGamepadButtonLabelForType(typ GamepadType, button GamepadButton) GamepadButtonLabel {
	return (GamepadButtonLabel)(C.SDL_GetGamepadButtonLabelForType((C.SDL_GamepadType)(typ), (C.SDL_GamepadButton)(int8(button))))
}

// Get the label of a button on a gamepad.
//
// gamepad: a gamepad.
//
// button: a button index (one of the [GamepadButton] values).
//
// Returns the [GamepadButtonLabel] enum corresponding to the button label.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadButtonLabel
func (gamepad *Gamepad) ButtonLabel(button GamepadButton) GamepadButtonLabel {
	return (GamepadButtonLabel)(C.SDL_GetGamepadButtonLabel((*C.SDL_Gamepad)(gamepad), (C.SDL_GamepadButton)(int8(button))))
}

// Get the number of touchpads on a gamepad.
//
// gamepad: a gamepad.
//
// Returns number of touchpads.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetNumGamepadTouchpads
func (gamepad *Gamepad) NumTouchpads() int {
	return (int)(C.SDL_GetNumGamepadTouchpads((*C.SDL_Gamepad)(gamepad)))
}

// Get the number of supported simultaneous fingers on a touchpad on a game
// gamepad.
//
// gamepad: a gamepad.
//
// touchpad: a touchpad.
//
// Returns number of supported simultaneous fingers.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetNumGamepadTouchpadFingers
func (gamepad *Gamepad) NumTouchpadFingers(touchpad int) int {
	return (int)(C.SDL_GetNumGamepadTouchpadFingers((*C.SDL_Gamepad)(gamepad), (C.int)(touchpad)))
}

// Get the current state of a finger on a touchpad on a gamepad.
//
// gamepad: a gamepad.
//
// touchpad: a touchpad.
//
// finger: a finger.
//
// down: true if the finger is down, false
// otherwise.
//
// x: the x position, normalized 0 to 1, with the
// origin in the upper left.
//
// y: the y position, normalized 0 to 1, with the
// origin in the upper left.
//
// pressure: the pressure value.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadTouchpadFinger
func (gamepad *Gamepad) TouchpadFinger(touchpad int, finger int) (down bool, x, y, pressure float32, err error) {
	if !C.SDL_GetGamepadTouchpadFinger((*C.SDL_Gamepad)(gamepad), (C.int)(touchpad), (C.int)(finger), (*C.bool)(&down), (*C.float)(&x), (*C.float)(&y), (*C.float)(&pressure)) {
		return false, 0, 0, 0, getError()
	}
	return
}

// Return whether a gamepad has a particular sensor.
//
// gamepad: the gamepad to query.
//
// type: the type of sensor to query.
//
// Returns true if the sensor exists, false otherwise.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GamepadHasSensor
func (gamepad *Gamepad) HasSensor(typ SensorType) bool {
	return (bool)(C.SDL_GamepadHasSensor((*C.SDL_Gamepad)(gamepad), (C.SDL_SensorType)(typ)))
}

// Set whether data reporting for a gamepad sensor is enabled.
//
// gamepad: the gamepad to update.
//
// type: the type of sensor to enable/disable.
//
// enabled: whether data reporting should be enabled.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetGamepadSensorEnabled
func (gamepad *Gamepad) SetSensorEnabled(typ SensorType, enabled bool) error {
	if !C.SDL_SetGamepadSensorEnabled((*C.SDL_Gamepad)(gamepad), (C.SDL_SensorType)(typ), (C.bool)(enabled)) {
		return getError()
	}
	return nil
}

// Query whether sensor data reporting is enabled for a gamepad.
//
// gamepad: the gamepad to query.
//
// type: the type of sensor to query.
//
// Returns true if the sensor is enabled, false otherwise.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GamepadSensorEnabled
func (gamepad *Gamepad) SensorEnabled(typ SensorType) bool {
	return (bool)(C.SDL_GamepadSensorEnabled((*C.SDL_Gamepad)(gamepad), (C.SDL_SensorType)(typ)))
}

// Get the data rate (number of events per second) of a gamepad sensor.
//
// gamepad: the gamepad to query.
//
// type: the type of sensor to query.
//
// Returns the data rate, or 0.0f if the data rate is not available.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadSensorDataRate
func (gamepad *Gamepad) SensorDataRate(typ SensorType) float32 {
	return (float32)(C.SDL_GetGamepadSensorDataRate((*C.SDL_Gamepad)(gamepad), (C.SDL_SensorType)(typ)))
}

// Get the current state of a gamepad sensor.
//
// The number of values and interpretation of the data is sensor dependent.
// See [SensorType] for the details for each type of sensor.
//
// gamepad: the gamepad to query.
//
// type: the type of sensor to query.
//
// data: a pointer filled with the current sensor state.
//
// num_values: the number of values to write to data.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadSensorData
func (gamepad *Gamepad) SensorData(typ SensorType, data []float32) error {
	if !C.SDL_GetGamepadSensorData((*C.SDL_Gamepad)(gamepad), (C.SDL_SensorType)(typ), (*C.float)(unsafe.SliceData(data)), (C.int)(len(data))) {
		return getError()
	}
	return nil
}

// Start a rumble effect on a gamepad.
//
// Each call to this function cancels any previous rumble effect, and calling
// it with 0 intensity stops any rumbling.
//
// This function requires you to process SDL events or call
// [UpdateJoysticks] to update rumble state.
//
// gamepad: the gamepad to vibrate.
//
// lowFrequencyRumble: the intensity of the low frequency (left)
// rumble motor, from 0 to 0xFFFF.
//
// highFrequencyRumble: the intensity of the high frequency (right)
// rumble motor, from 0 to 0xFFFF.
//
// durationMS: the duration of the rumble effect, in milliseconds.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RumbleGamepad
func (gamepad *Gamepad) Rumble(lowFrequencyRumble uint16, highFrequencyRumble uint16, durationMS uint32) error {
	if !C.SDL_RumbleGamepad((*C.SDL_Gamepad)(gamepad), (C.Uint16)(lowFrequencyRumble), (C.Uint16)(highFrequencyRumble), (C.Uint32)(durationMS)) {
		return getError()
	}
	return nil
}

// Start a rumble effect in the gamepad's triggers.
//
// Each call to this function cancels any previous trigger rumble effect, and
// calling it with 0 intensity stops any rumbling.
//
// Note that this is rumbling of the _triggers_ and not the gamepad as a
// whole. This is currently only supported on Xbox One gamepads. If you want
// the (more common) whole-gamepad rumble, use [RumbleGamepad] instead.
//
// This function requires you to process SDL events or call
// [UpdateJoysticks] to update rumble state.
//
// gamepad: the gamepad to vibrate.
//
// leftRumble: the intensity of the left trigger rumble motor, from 0
// to 0xFFFF.
//
// rightRumble: the intensity of the right trigger rumble motor, from 0
// to 0xFFFF.
//
// durationMS: the duration of the rumble effect, in milliseconds.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RumbleGamepadTriggers
func (gamepad *Gamepad) RumbleTriggers(leftRumble uint16, rightRumble uint16, durationMS uint32) error {
	if !C.SDL_RumbleGamepadTriggers((*C.SDL_Gamepad)(gamepad), (C.Uint16)(leftRumble), (C.Uint16)(rightRumble), (C.Uint32)(durationMS)) {
		return getError()
	}
	return nil
}

// Update a gamepad's LED color.
//
// An example of a joystick LED is the light on the back of a PlayStation 4's
// DualShock 4 controller.
//
// For gamepads with a single color LED, the maximum of the RGB values will be
// used as the LED brightness.
//
// gamepad: the gamepad to update.
//
// red: the intensity of the red LED.
//
// green: the intensity of the green LED.
//
// blue: the intensity of the blue LED.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetGamepadLED
func (gamepad *Gamepad) SetLED(red byte, green byte, blue byte) error {
	if !C.SDL_SetGamepadLED((*C.SDL_Gamepad)(gamepad), (C.Uint8)(red), (C.Uint8)(green), (C.Uint8)(blue)) {
		return getError()
	}
	return nil
}

// Send a gamepad specific effect packet.
//
// gamepad: the gamepad to affect.
//
// data: the data to send to the gamepad.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SendGamepadEffect
func (gamepad *Gamepad) SendEffect(data []byte) error {
	if !C.SDL_SendGamepadEffect((*C.SDL_Gamepad)(gamepad), unsafe.Pointer(unsafe.SliceData(data)), (C.int)(len(data))) {
		return getError()
	}
	return nil
}

// Close a gamepad previously opened with [OpenGamepad].
//
// gamepad: a gamepad identifier previously returned by [OpenGamepad].
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CloseGamepad
func (gamepad *Gamepad) Close() {
	C.SDL_CloseGamepad((*C.SDL_Gamepad)(gamepad))
}

// Return the sfSymbolsName for a given button on a gamepad on Apple
// platforms.
//
// gamepad: the gamepad to query.
//
// button: a button on the gamepad.
//
// Returns the sfSymbolsName or an empty string if the name can't be found.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadAppleSFSymbolsNameForButton
func (gamepad *Gamepad) AppleSFSymbolsNameForButton(button GamepadButton) string {
	return C.GoString(C.SDL_GetGamepadAppleSFSymbolsNameForButton((*C.SDL_Gamepad)(gamepad), (C.SDL_GamepadButton)(int8(button))))
}

// Return the sfSymbolsName for a given axis on a gamepad on Apple platforms.
//
// gamepad: the gamepad to query.
//
// axis: an axis on the gamepad.
//
// Returns the sfSymbolsName or an empty string if the name can't be found.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGamepadAppleSFSymbolsNameForAxis
func (gamepad *Gamepad) AppleSFSymbolsNameForAxis(axis GamepadAxis) string {
	return C.GoString(C.SDL_GetGamepadAppleSFSymbolsNameForAxis((*C.SDL_Gamepad)(gamepad), (C.SDL_GamepadAxis)(axis)))
}
