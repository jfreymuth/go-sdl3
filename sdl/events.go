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
// #cgo noescape SDL_PumpEvents
// #cgo noescape SDL_PeepEvents
// #cgo noescape SDL_HasEvent
// #cgo nocallback SDL_HasEvent
// #cgo noescape SDL_HasEvents
// #cgo nocallback SDL_HasEvents
// #cgo noescape SDL_FlushEvent
// #cgo nocallback SDL_FlushEvent
// #cgo noescape SDL_FlushEvents
// #cgo nocallback SDL_FlushEvents
// #cgo noescape SDL_PollEvent
// #cgo noescape SDL_WaitEvent
// #cgo noescape SDL_WaitEventTimeout
// #cgo noescape SDL_PushEvent
// #cgo noescape wrap_SDL_SetEventFilter
// #cgo nocallback wrap_SDL_SetEventFilter
// #cgo noescape wrap_SDL_AddEventWatch
// #cgo nocallback wrap_SDL_AddEventWatch
// #cgo noescape wrap_SDL_RemoveEventWatch
// #cgo nocallback wrap_SDL_RemoveEventWatch
// #cgo noescape wrap_SDL_FilterEvents
// #cgo noescape SDL_SetEventEnabled
// #cgo nocallback SDL_SetEventEnabled
// #cgo noescape SDL_EventEnabled
// #cgo nocallback SDL_EventEnabled
// #cgo noescape SDL_RegisterEvents
// #cgo nocallback SDL_RegisterEvents
// #cgo noescape SDL_GetWindowFromEvent
// #cgo nocallback SDL_GetWindowFromEvent
// #include <SDL3/SDL.h>
// extern void wrap_SDL_SetEventFilter(uintptr_t userdata);
// extern void wrap_SDL_FilterEvents(uintptr_t userdata);
// extern bool wrap_SDL_AddEventWatch(uintptr_t userdata);
// extern void wrap_SDL_RemoveEventWatch(uintptr_t userdata);
import "C"
import (
	"runtime/cgo"
	"sync"
	"unsafe"
)

// # CategoryEvents
//
// Event queue management.
//
// It's extremely common – often required – that an app deal with SDL's event
// queue. Almost all useful information about interactions with the real world
// flow through here: the user interacting with the computer and app, hardware
// coming and going, the system changing in some way, etc.
//
// An app generally takes a moment, perhaps at the start of a new frame, to
// examine any events that have occured since the last time and process or
// ignore them. This is generally done by calling [PollEvent] in a loop
// until it returns false.
//
// There is other forms of control, too: [PeepEvents] has more
// functionality at the cost of more complexity, and [WaitEvent] can block
// the process until something interesting happens, which might be beneficial
// for certain types of programs on low-power hardware. One may also call
// [AddEventWatch] to set a callback when new events arrive.
//
// The app is free to generate their own events, too: [PushEvent] allows the
// app to put events onto the queue for later retrieval; [RegisterEvents]
// can guarantee that these events have a type that isn't in use by other
// parts of the system.

// The types of events that can be delivered.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_EventType
type EventType uint32

const (
	EventFirst                      EventType = 0     // Unused (do not remove)
	EventQuit                       EventType = 0x100 // User-requested quit
	EventLocaleChanged              EventType = 0x101 // The user's locale preferences have changed.
	EventSystemThemeChanged         EventType = 0x102 // The system theme changed
	EventDisplayOrientation         EventType = 0x151 // Display orientation has changed to data1
	EventDisplayAdded               EventType = 0x152 // Display has been added to the system
	EventDisplayRemoved             EventType = 0x153 // Display has been removed from the system
	EventDisplayMoved               EventType = 0x154 // Display has changed position
	EventDisplayDesktopModeChanged  EventType = 0x155 // Display has changed desktop mode
	EventDisplayCurrentModeChanged  EventType = 0x156 // Display has changed current mode
	EventDisplayContentScaleChanged EventType = 0x157 // Display has changed content scale
	EventDisplayFirst               EventType = EventDisplayOrientation
	EventDisplayLast                EventType = EventDisplayContentScaleChanged
	EventWindowShown                EventType = 0x202 // Window has been shown
	EventWindowHidden               EventType = 0x203 // Window has been hidden
	EventWindowExposed              EventType = 0x204 // Window has been exposed and should be redrawn, and can be redrawn directly from event watchers for this event
	EventWindowMoved                EventType = 0x205 // Window has been moved to data1, data2
	EventWindowResized              EventType = 0x206 // Window has been resized to data1xdata2
	EventWindowPixelSizeChanged     EventType = 0x207 // The pixel size of the window has changed to data1xdata2
	EventWindowMetalViewResized     EventType = 0x208 // The pixel size of a Metal view associated with the window has changed
	EventWindowMinimized            EventType = 0x209 // Window has been minimized
	EventWindowMaximized            EventType = 0x20A // Window has been maximized
	EventWindowRestored             EventType = 0x20B // Window has been restored to normal size and position
	EventWindowMouseEnter           EventType = 0x20C // Window has gained mouse focus
	EventWindowMouseLeave           EventType = 0x20D // Window has lost mouse focus
	EventWindowFocusGained          EventType = 0x20E // Window has gained keyboard focus
	EventWindowFocusLost            EventType = 0x20F // Window has lost keyboard focus
	EventWindowCloseRequested       EventType = 0x210 // The window manager requests that the window be closed
	EventWindowHitTest              EventType = 0x211 // Window had a hit test that wasn't [HittestNormal]
	EventWindowICCprofChanged       EventType = 0x212 // The ICC profile of the window's display has changed
	EventWindowDisplayChanged       EventType = 0x213 // Window has been moved to display data1
	EventWindowDisplayScaleChanged  EventType = 0x214 // Window display scale has been changed
	EventWindowSafeAreaChanged      EventType = 0x215 // The window safe area has been changed
	EventWindowOccluded             EventType = 0x216 // The window has been occluded
	EventWindowEnterFullscreen      EventType = 0x217 // The window has entered fullscreen mode
	EventWindowLeaveFullscreen      EventType = 0x218 // The window has left fullscreen mode
	EventWindowHDRStateChanged      EventType = 0x219 // Window HDR properties have changed
	EventWindowFirst                EventType = EventWindowShown
	EventWindowLast                 EventType = EventWindowHDRStateChanged
	EventKeyDown                    EventType = 0x300 // Key pressed
	EventKeyUp                      EventType = 0x301 // Key released
	EventTextEditing                EventType = 0x302 // Keyboard text editing (composition)
	EventTextInput                  EventType = 0x303 // Keyboard text input
	EventKeymapChanged              EventType = 0x304 // Keymap changed due to a system event such as an input language or keyboard layout change.
	EventKeyboardAdded              EventType = 0x305 // A new keyboard has been inserted into the system
	EventKeyboardRemoved            EventType = 0x306 // A keyboard has been removed
	EventTextEditingCandidates      EventType = 0x307 // Keyboard text editing candidates
	EventMouseMotion                EventType = 0x400 // Mouse moved
	EventMouseButtonDown            EventType = 0x401 // Mouse button pressed
	EventMouseButtonUp              EventType = 0x402 // Mouse button released
	EventMouseWheel                 EventType = 0x403 // Mouse wheel motion
	EventMouseAdded                 EventType = 0x404 // A new mouse has been inserted into the system
	EventMouseRemoved               EventType = 0x405 // A mouse has been removed
	EventJoystickAxisMotion         EventType = 0x600 // Joystick axis motion
	EventJoystickBallMotion         EventType = 0x601 // Joystick trackball motion
	EventJoystickHatMotion          EventType = 0x602 // Joystick hat position change
	EventJoystickButtonDown         EventType = 0x603 // Joystick button pressed
	EventJoystickButtonUp           EventType = 0x604 // Joystick button released
	EventJoystickAdded              EventType = 0x605 // A new joystick has been inserted into the system
	EventJoystickRemoved            EventType = 0x606 // An opened joystick has been removed
	EventJoystickBatteryUpdated     EventType = 0x607 // Joystick battery level change
	EventJoystickUpdateComplete     EventType = 0x608 // Joystick update is complete
	EventGamepadAxisMotion          EventType = 0x650 // Gamepad axis motion
	EventGamepadButtonDown          EventType = 0x651 // Gamepad button pressed
	EventGamepadButtonUp            EventType = 0x652 // Gamepad button released
	EventGamepadAdded               EventType = 0x653 // A new gamepad has been inserted into the system
	EventGamepadRemoved             EventType = 0x654 // A gamepad has been removed
	EventGamepadRemapped            EventType = 0x655 // The gamepad mapping was updated
	EventGamepadTouchpadDown        EventType = 0x656 // Gamepad touchpad was touched
	EventGamepadTouchpadMotion      EventType = 0x657 // Gamepad touchpad finger was moved
	EventGamepadTouchpadUp          EventType = 0x658 // Gamepad touchpad finger was lifted
	EventGamepadSensorUpdate        EventType = 0x659 // Gamepad sensor was updated
	EventGamepadUpdateComplete      EventType = 0x65A // Gamepad update is complete
	EventGamepadSteamHandleUpdated  EventType = 0x65B // Gamepad Steam handle has changed
	EventFingerDown                 EventType = 0x700
	EventFingerUp                   EventType = 0x701
	EventFingerMotion               EventType = 0x702
	EventFingerCanceled             EventType = 0x703
	EventClipboardUpdate            EventType = 0x900  // The clipboard or primary selection changed
	EventDropFile                   EventType = 0x1000 // The system requests a file open
	EventDropText                   EventType = 0x1001 // text/plain drag-and-drop event
	EventDropBegin                  EventType = 0x1002 // A new set of drops is beginning (empty filename)
	EventDropComplete               EventType = 0x1003 // Current set of drops is now complete (empty filename)
	EventDropPosition               EventType = 0x1004 // Position while moving over the window
	EventAudioDeviceAdded           EventType = 0x1100 // A new audio device is available
	EventAudioDeviceRemoved         EventType = 0x1101 // An audio device has been removed.
	EventAudioDeviceFormatChanged   EventType = 0x1102 // An audio device's format has been changed by the system.
	EventSensorUpdate               EventType = 0x1200 // A sensor was updated
	EventPenProximityIn             EventType = 0x1300 // Pressure-sensitive pen has become available
	EventPenProximityOut            EventType = 0x1301 // Pressure-sensitive pen has become unavailable
	EventPenDown                    EventType = 0x1302 // Pressure-sensitive pen touched drawing surface
	EventPenUp                      EventType = 0x1303 // Pressure-sensitive pen stopped touching drawing surface
	EventPenButtonDown              EventType = 0x1304 // Pressure-sensitive pen button pressed
	EventPenButtonUp                EventType = 0x1305 // Pressure-sensitive pen button released
	EventPenMotion                  EventType = 0x1306 // Pressure-sensitive pen is moving on the tablet
	EventPenAxis                    EventType = 0x1307 // Pressure-sensitive pen angle/pressure/etc changed
	EventCameraDeviceAdded          EventType = 0x1400 // A new camera device is available
	EventCameraDeviceRemoved        EventType = 0x1401 // A camera device has been removed.
	EventCameraDeviceApproved       EventType = 0x1402 // A camera device has been approved for use by the user.
	EventCameraDeviceDenied         EventType = 0x1403 // A camera device has been denied for use by the user.
	EventRenderTargetsReset         EventType = 0x2000 // The render targets have been reset and their contents need to be updated
	EventRenderDeviceReset          EventType = 0x2001 // The device has been reset and all textures need to be recreated
	EventRenderDeviceLost           EventType = 0x2002 // The device has been lost and can't be recovered.
	EventPrivate0                   EventType = 0x4000
	EventPrivate1                   EventType = 0x4001
	EventPrivate2                   EventType = 0x4002
	EventPrivate3                   EventType = 0x4003
	EventPollSentinel               EventType = 0x7F00 // Signals the end of an event poll cycle
	EventUser                       EventType = 0x8000
	EventLast                       EventType = 0xFFFF
	EventEnumPadding                EventType = 0x7FFFFFFF
)

// Display state change event data ([Event.Display])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DisplayEvent
type DisplayEvent struct {
	Type      EventType // EventDisplay*
	Timestamp uint64    // In nanoseconds, populated using [GetTicksNS]
	DisplayID DisplayID // The associated display
	Data1     int       // event dependent data
	Data2     int       // event dependent data
}

// Window state change event data ([Event.Window])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WindowEvent
type WindowEvent struct {
	Type      EventType // EventWindow*
	Timestamp uint64    // In nanoseconds, populated using [GetTicksNS]
	WindowID  WindowID  // The associated window
	Data1     int       // event dependent data
	Data2     int       // event dependent data
}

// Keyboard device event structure ([Event.KeyboardDevice])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_KeyboardDeviceEvent
type KeyboardDeviceEvent struct {
	Type      EventType  // [EventKeyboardAdded] or [EventKeyboardRemoved]
	Timestamp uint64     // In nanoseconds, populated using [GetTicksNS]
	Which     KeyboardID // The keyboard instance id
}

// Keyboard button event structure ([Event.Keyboard])
//
// The key is the base [Keycode] generated by pressing the scancode
// using the current keyboard layout, applying any options specified in
// [HintKeycodeOptions]. You can get the [Keycode] corresponding to the
// event scancode and modifiers directly from the keyboard layout, bypassing
// [HintKeycodeOptions], by calling [GetKeyFromScancode].
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_KeyboardEvent
type KeyboardEvent struct {
	Type      EventType  // [EventKeyDown] or [EventKeyUp]
	Timestamp uint64     // In nanoseconds, populated using [GetTicksNS]
	WindowID  WindowID   // The window with keyboard focus, if any
	Which     KeyboardID // The keyboard instance id, or 0 if unknown or virtual
	Scancode  Scancode   // SDL physical key code
	Key       Keycode    // SDL virtual key code
	Mod       Keymod     // current key modifiers
	Raw       uint16     // The platform dependent scancode for this event
	Down      bool       // true if the key is pressed
	Repeat    bool       // true if this is a key repeat
}

// Keyboard text editing event structure ([Event.TextEditing])
//
// The start cursor is the position, in UTF-8 characters, where new typing
// will be inserted into the editing text. The length is the number of UTF-8
// characters that will be replaced by new typing.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_TextEditingEvent
type TextEditingEvent struct {
	Type      EventType // [EventTextEditing]
	Timestamp uint64    // In nanoseconds, populated using [GetTicksNS]
	WindowID  WindowID  // The window with keyboard focus, if any
	Text      string    // The editing text
	Start     int       // The start cursor of selected editing text, or -1 if not set
	Length    int       // The length of selected editing text, or -1 if not set
}

// Keyboard IME candidates event structure ([Event.TextEditingCandidates])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_TextEditingCandidatesEvent
type TextEditingCandidatesEvent struct {
	Type              EventType // [EventTextEditingCandidates]
	Timestamp         uint64    // In nanoseconds, populated using [GetTicksNS]
	WindowID          WindowID  // The window with keyboard focus, if any
	Candidates        []string  // The list of candidates, or nil if there are no candidates available
	SelectedCandidate int32     // The index of the selected candidate, or -1 if no candidate is selected
	Horizontal        bool      // true if the list is horizontal, false if it's vertical
}

// Keyboard text input event structure ([Event.TextInput])
//
// This event will never be delivered unless text input is enabled by calling
// [StartTextInput]. Text input is disabled by default!
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_TextInputEvent
type TextInputEvent struct {
	Type      EventType // [EventTextInput]
	Timestamp uint64    // In nanoseconds, populated using [GetTicksNS]
	WindowID  WindowID  // The window with keyboard focus, if any
	Text      string    // The input text, UTF-8 encoded
}

// Mouse device event structure ([Event.MouseDevice])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MouseDeviceEvent
type MouseDeviceEvent struct {
	Type      EventType // [EventMouseAdded] or [EventMouseRemoved]
	Timestamp uint64    // In nanoseconds, populated using [GetTicksNS]
	Which     MouseID   // The mouse instance id
}

// Mouse motion event structure ([Event.MouseMotion])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MouseMotionEvent
type MouseMotionEvent struct {
	Type      EventType        // [EventMouseMotion]
	Timestamp uint64           // In nanoseconds, populated using [GetTicksNS]
	WindowID  WindowID         // The window with mouse focus, if any
	Which     MouseID          // The mouse instance id in relative mode, [TouchMouseID] for touch events, or 0
	State     MouseButtonFlags // The current button state
	X         float32          // X coordinate, relative to window
	Y         float32          // Y coordinate, relative to window
	XRel      float32          // The relative motion in the X direction
	YRel      float32          // The relative motion in the Y direction
}

// Mouse button event structure ([Event.MouseButton])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MouseButtonEvent
type MouseButtonEvent struct {
	Type      EventType // [EventMouseButtonDown] or [EventMouseButtonUp]
	Timestamp uint64    // In nanoseconds, populated using [GetTicksNS]
	WindowID  WindowID  // The window with mouse focus, if any
	Which     MouseID   // The mouse instance id in relative mode, [TouchMouseID] for touch events, or 0
	Button    byte      // The mouse button index
	Down      bool      // true if the button is pressed
	Clicks    byte      // 1 for single-click, 2 for double-click, etc.
	X         float32   // X coordinate, relative to window
	Y         float32   // Y coordinate, relative to window
}

// Mouse wheel event structure ([Event.MouseWheel])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MouseWheelEvent
type MouseWheelEvent struct {
	Type      EventType           // [EventMouseWheel]
	Timestamp uint64              // In nanoseconds, populated using [GetTicksNS]
	WindowID  WindowID            // The window with mouse focus, if any
	Which     MouseID             // The mouse instance id in relative mode or 0
	X         float32             // The amount scrolled horizontally, positive to the right and negative to the left
	Y         float32             // The amount scrolled vertically, positive away from the user and negative toward the user
	Direction MouseWheelDirection // Set to a [MouseWheelDirection] value. When [MousewheelFlipped] the values in X and Y will be opposite. Multiply by -1 to change them back
	MouseX    float32             // X coordinate, relative to window
	MouseY    float32             // Y coordinate, relative to window
}

// Joystick axis motion event structure ([Event.JoyAxis])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_JoyAxisEvent
type JoyAxisEvent struct {
	Type      EventType  // [EventJoystickAxisMotion]
	Timestamp uint64     // In nanoseconds, populated using [GetTicksNS]
	Which     JoystickID // The joystick instance id
	Axis      byte       // The joystick axis index
	Value     int16      // The axis value (range: -32768 to 32767)
}

// Joystick trackball motion event structure ([Event.JoyBall])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_JoyBallEvent
type JoyBallEvent struct {
	Type      EventType  // [EventJoystickBallMotion]
	Timestamp uint64     // In nanoseconds, populated using [GetTicksNS]
	Which     JoystickID // The joystick instance id
	Ball      byte       // The joystick trackball index
	XRel      int16      // The relative motion in the X direction
	YRel      int16      // The relative motion in the Y direction
}

// Joystick hat position change event structure ([Event.JoyHat])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_JoyHatEvent
type JoyHatEvent struct {
	Type      EventType        // [EventJoystickHatMotion]
	Timestamp uint64           // In nanoseconds, populated using [GetTicksNS]
	Which     JoystickID       // The joystick instance id
	Hat       byte             // The joystick hat index
	Value     JoystickHatState // The hat position value. Note that zero means the POV is centered.
}

// Joystick button event structure ([Event.JoyButton])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_JoyButtonEvent
type JoyButtonEvent struct {
	Type      EventType  // [EventJoystickButtonDown] or [EventJoystickButtonUp]
	Timestamp uint64     // In nanoseconds, populated using [GetTicksNS]
	Which     JoystickID // The joystick instance id
	Button    byte       // The joystick button index
	Down      bool       // true if the button is pressed
}

// Joystick device event structure ([Event.JoyDevice])
//
// SDL will send [EventJoystickAdded] events for devices that are already plugged in
// during [Init].
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_JoyDeviceEvent
type JoyDeviceEvent struct {
	Type      EventType  // [EventJoystickAdded] or [EventJoystickRemoved] or [EventJoystickUpdateComplete]
	Timestamp uint64     // In nanoseconds, populated using [GetTicksNS]
	Which     JoystickID // The joystick instance id
}

// Joystick battery level change event structure ([Event.JoyBattery])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_JoyBatteryEvent
type JoyBatteryEvent struct {
	Type      EventType  // [EventJoystickBatteryUpdated]
	Timestamp uint64     // In nanoseconds, populated using [GetTicksNS]
	Which     JoystickID // The joystick instance id
	State     PowerState // The joystick battery state
	Percent   int        // The joystick battery percent charge remaining
}

// Gamepad axis motion event structure ([Event.GamepadAxis])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GamepadAxisEvent
type GamepadAxisEvent struct {
	Type      EventType   // [EventGamepadAxisMotion]
	Timestamp uint64      // In nanoseconds, populated using [GetTicksNS]
	Which     JoystickID  // The joystick instance id
	Axis      GamepadAxis // The gamepad axis
	Value     int16       // The axis value (range: -32768 to 32767)
}

// Gamepad button event structure ([Event.GamepadButton])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GamepadButtonEvent
type GamepadButtonEvent struct {
	Type      EventType     // [EventGamepadButtonDown] or [EventGamepadButtonUp]
	Timestamp uint64        // In nanoseconds, populated using [GetTicksNS]
	Which     JoystickID    // The joystick instance id
	Button    GamepadButton // The gamepad button
	Down      bool          // true if the button is pressed
}

// Gamepad device event structure ([Event.GamepadDevice])
//
// Joysticks that are supported gamepads receive both a [JoyDeviceEvent]
// and a [GamepadDeviceEvent].
//
// SDL will send [EventGamepadAdded] events for joysticks that are already
// plugged in during [Init] and are recognized as gamepads. It will also send
// events for joysticks that get gamepad mappings at runtime.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GamepadDeviceEvent
type GamepadDeviceEvent struct {
	Type      EventType  // [EventGamepadAdded], [EventGamepadRemoved], or [EventGamepadRemapped], [EventGamepadUpdateComplete] or [EventGamepadSteamHandleUpdated]
	Timestamp uint64     // In nanoseconds, populated using [GetTicksNS]
	Which     JoystickID // The joystick instance id
}

// Gamepad touchpad event structure ([Event.GamepadTouchpad])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GamepadTouchpadEvent
type GamepadTouchpadEvent struct {
	Type      EventType  // [EventGamepadTouchpadDown] or [EventGamepadTouchpadMotion] or [EventGamepadTouchpadUp]
	Timestamp uint64     // In nanoseconds, populated using [GetTicksNS]
	Which     JoystickID // The joystick instance id
	Touchpad  int        // The index of the touchpad
	Finger    int        // The index of the finger on the touchpad
	X         float32    // Normalized in the range 0...1 with 0 being on the left
	Y         float32    // Normalized in the range 0...1 with 0 being at the top
	Pressure  float32    // Normalized in the range 0...1
}

// Gamepad sensor event structure ([Event.GamepadSensor])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GamepadSensorEvent
type GamepadSensorEvent struct {
	Type            EventType  // [EventGamepadSensorUpdate]
	Timestamp       uint64     // In nanoseconds, populated using [GetTicksNS]
	Which           JoystickID // The joystick instance id
	Sensor          int        // The type of the sensor, one of the values of [SensorType]
	Data            [3]float32 // Up to 3 values from the sensor
	SensorTimestamp uint64     // The timestamp of the sensor reading in nanoseconds, not necessarily synchronized with the system clock
}

// Audio device event structure ([Event.AudioDevice])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AudioDeviceEvent
type AudioDeviceEvent struct {
	Type      EventType     // [EventAudioDeviceAdded], or [EventAudioDeviceRemoved], or [EventAudioDeviceFormatChanged]
	Timestamp uint64        // In nanoseconds, populated using [GetTicksNS]
	Which     AudioDeviceID // [AudioDeviceID] for the device being added or removed or changing
	Recording bool          // false if a playback device, true if a recording device.
}

// Camera device event structure ([Event.CameraDevice])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CameraDeviceEvent
type CameraDeviceEvent struct {
	Type      EventType // [EventCameraDeviceAdded], [EventCameraDeviceRemoved], [EventCameraDeviceApproved], [EventCameraDeviceDenied]
	Timestamp uint64    // In nanoseconds, populated using [GetTicksNS]
	Which     CameraID  // [CameraID] for the device being added or removed or changing
}

// Renderer event structure ([Event.Render])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderEvent
type RenderEvent struct {
	Type      EventType // [EventRenderTargetsReset], [EventRenderDeviceReset], [EventRenderDeviceLost]
	Timestamp uint64    // In nanoseconds, populated using [GetTicksNS]
	WindowID  WindowID  // The window containing the renderer in question.
}

// Touch finger event structure ([Event.TouchFinger])
//
// Coordinates in this event are normalized. X and Y are normalized to a
// range between 0.0f and 1.0f, relative to the window, so (0,0) is the top
// left and (1,1) is the bottom right. Delta coordinates DX and DY are
// normalized in the ranges of -1.0f (traversed all the way from the bottom or
// right to all the way up or left) to 1.0f (traversed all the way from the
// top or left to all the way down or right).
//
// Note that while the coordinates are _normalized_, they are not _clamped_,
// which means in some circumstances you can get a value outside of this
// range. For example, a renderer using logical presentation might give a
// negative value when the touch is in the letterboxing. Some platforms might
// report a touch outside of the window, which will also be outside of the
// range.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_TouchFingerEvent
type TouchFingerEvent struct {
	Type      EventType // [EventFingerDown], [EventFingerUp], [EventFingerMotion], or [EventFingerCanceled]
	Timestamp uint64    // In nanoseconds, populated using [GetTicksNS]
	TouchID   TouchID   // The touch device id
	FingerID  FingerID  //
	X         float32   // Normalized in the range 0...1
	Y         float32   // Normalized in the range 0...1
	DX        float32   // Normalized in the range -1...1
	DY        float32   // Normalized in the range -1...1
	Pressure  float32   // Normalized in the range 0...1
	WindowID  WindowID  // The window underneath the finger, if any
}

// Pressure-sensitive pen proximity event structure ([Event.PenMotion])
//
// When a pen becomes visible to the system (it is close enough to a tablet,
// etc), SDL will send an [EventPenProximityIn] event with the new pen's
// ID. This ID is valid until the pen leaves proximity again (has been removed
// from the tablet's area, the tablet has been unplugged, etc). If the same
// pen reenters proximity again, it will be given a new ID.
//
// Note that "proximity" means "close enough for the tablet to know the tool
// is there." The pen touching and lifting off from the tablet while not
// leaving the area are handled by [EventPenDown] and [EventPenUp].
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PenProximityEvent
type PenProximityEvent struct {
	Type      EventType // [EventPenProximityIn] or [EventPenProximityOut]
	Timestamp uint64    // In nanoseconds, populated using [GetTicksNS]
	WindowID  WindowID  // The window with pen focus, if any
	Which     PenID     // The pen instance id
}

// Pressure-sensitive pen motion event structure ([Event.PenMotion])
//
// Depending on the hardware, you may get motion events when the pen is not
// touching a tablet, for tracking a pen even when it isn't drawing. You
// should listen for [EventPenDown] and [EventPenUp] events, or check
// `pen_state & [PenInputDown]` to decide if a pen is "drawing" when
// dealing with pen motion.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PenMotionEvent
type PenMotionEvent struct {
	Type      EventType     // [EventPenMotion]
	Timestamp uint64        // In nanoseconds, populated using [GetTicksNS]
	WindowID  WindowID      // The window with pen focus, if any
	Which     PenID         // The pen instance id
	PenState  PenInputFlags // Complete pen input state at time of event
	X         float32       // X coordinate, relative to window
	Y         float32       // Y coordinate, relative to window
}

// Pressure-sensitive pen touched event structure ([Event.PenTouch])
//
// These events come when a pen touches a surface (a tablet, etc), or lifts
// off from one.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PenTouchEvent
type PenTouchEvent struct {
	Type      EventType     // [EventPenDown] or [EventPenUp]
	Timestamp uint64        // In nanoseconds, populated using [GetTicksNS]
	WindowID  WindowID      // The window with pen focus, if any
	Which     PenID         // The pen instance id
	PenState  PenInputFlags // Complete pen input state at time of event
	X         float32       // X coordinate, relative to window
	Y         float32       // Y coordinate, relative to window
	Eraser    bool          // true if eraser end is used (not all pens support this).
	Down      bool          // true if the pen is touching or false if the pen is lifted off
}

// Pressure-sensitive pen button event structure ([Event.PenButton])
//
// This is for buttons on the pen itself that the user might click. The pen
// itself pressing down to draw triggers a [EventPenDown] event instead.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PenButtonEvent
type PenButtonEvent struct {
	Type      EventType     // [EventPenButtonDown] or [EventPenButtonUp]
	Timestamp uint64        // In nanoseconds, populated using [GetTicksNS]
	WindowID  WindowID      // The window with mouse focus, if any
	Which     PenID         // The pen instance id
	PenState  PenInputFlags // Complete pen input state at time of event
	X         float32       // X coordinate, relative to window
	Y         float32       // Y coordinate, relative to window
	Button    byte          // The pen button index (first button is 1).
	Down      bool          // true if the button is pressed
}

// Pressure-sensitive pen pressure / angle event structure ([Event.PenAxis])
//
// You might get some of these events even if the pen isn't touching the
// tablet.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PenAxisEvent
type PenAxisEvent struct {
	Type      EventType     // [EventPenAxis]
	Timestamp uint64        // In nanoseconds, populated using [GetTicksNS]
	WindowID  WindowID      // The window with pen focus, if any
	Which     PenID         // The pen instance id
	PenState  PenInputFlags // Complete pen input state at time of event
	X         float32       // X coordinate, relative to window
	Y         float32       // Y coordinate, relative to window
	Axis      PenAxis       // Axis that has changed
	Value     float32       // New value of axis
}

// An event used to drop text or request a file open by the system
// ([Event.Drop])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DropEvent
type DropEvent struct {
	Type      EventType // [EventDropBegin] or [EventDropFile] or [EventDropText] or [EventDropComplete] or [EventDropPosition]
	Timestamp uint64    // In nanoseconds, populated using [GetTicksNS]
	WindowID  WindowID  // The window that was dropped on, if any
	X         float32   // X coordinate, relative to window (not on begin)
	Y         float32   // Y coordinate, relative to window (not on begin)
	Source    string    // The source app that sent this drop event, or empty if that isn't available
	Data      string    // The text for [EventDropText] and the file name for [EventDropFile], empty for other events
}

// An event triggered when the clipboard contents have changed
// ([Event.Clipboard])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ClipboardEvent
type ClipboardEvent struct {
	Type      EventType // [EventClipboardUpdate]
	Timestamp uint64    // In nanoseconds, populated using [GetTicksNS]
	Owner     bool      // are we owning the clipboard (internal update)
	MimeTypes []string  // current mime types
}

// Sensor event structure ([Event.Sensor])
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SensorEvent
type SensorEvent struct {
	Type            EventType  // [EventSensorUpdate]
	Timestamp       uint64     // In nanoseconds, populated using [GetTicksNS]
	Which           SensorID   // The instance ID of the sensor
	Data            [6]float32 // Up to 6 values from the sensor - additional values can be queried using [Sensor.Read]
	SensorTimestamp uint64     // The timestamp of the sensor reading in nanoseconds, not necessarily synchronized with the system clock
}

// The "quit requested" event
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_QuitEvent
type QuitEvent struct {
	Type      EventType // [EventQuit]
	Timestamp uint64    // In nanoseconds, populated using [GetTicksNS]
}

// A user-defined event type ([Event.User])
//
// This event is unique; it is never created by SDL, but only by the
// application. The event can be pushed onto the event queue using
// [PushEvent]. The contents of the structure members are completely up to
// the programmer; the only requirement is that Type is a value obtained
// from [RegisterEvents].
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_UserEvent
type UserEvent struct {
	Type      EventType      // [EventUser] through [EventLast]-1
	Timestamp uint64         // In nanoseconds, populated using [GetTicksNS]
	WindowID  WindowID       // The associated window if any
	Code      int            // User defined event code
	Data1     unsafe.Pointer // User defined data pointer
	Data2     unsafe.Pointer // User defined data pointer
}

// The structure for all events in SDL.
//
// The [Event] structure is the core of all event handling in SDL. [Event]
// is a union of all event structures used in SDL.
//
// # TODO not about pointers, convert
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Event
type Event struct {
	internal C.SDL_Event
}

func (e *Event) Type() EventType {
	return EventType((*C.SDL_CommonEvent)(unsafe.Pointer(e))._type)
}

func (e *Event) Timestamp() uint64 {
	return uint64((*C.SDL_CommonEvent)(unsafe.Pointer(e)).timestamp)
}

// Display returns the event converted to a [DisplayEvent],
// or nil if the event type does not match
func (e *Event) Display() *DisplayEvent {
	typ := e.Type()
	if typ < EventDisplayFirst || typ > EventDisplayLast {
		return nil
	}
	ev := (*C.SDL_DisplayEvent)(unsafe.Pointer(e))
	return &DisplayEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		DisplayID(ev.displayID),
		int(ev.data1),
		int(ev.data2),
	}
}

// Window returns the event converted to a [WindowEvent],
// or nil if the event type does not match
func (e *Event) Window() *WindowEvent {
	typ := e.Type()
	if typ < EventWindowFirst || typ > EventWindowLast {
		return nil
	}
	ev := (*C.SDL_WindowEvent)(unsafe.Pointer(e))
	return &WindowEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		WindowID(ev.windowID),
		int(ev.data1),
		int(ev.data2),
	}
}

// KeyboardDevice returns the event converted to a [KeyboardDeviceEv],
// or nil if the event type does not match
func (e *Event) KeyboardDevice() *KeyboardDeviceEvent {
	typ := e.Type()
	if typ < EventKeyboardAdded || typ > EventKeyboardRemoved {
		return nil
	}
	ev := (*C.SDL_KeyboardDeviceEvent)(unsafe.Pointer(e))
	return &KeyboardDeviceEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		KeyboardID(ev.which),
	}
}

// Keyboard returns the event converted to a [KeyboardEvent],
// or nil if the event type does not match
func (e *Event) Keyboard() *KeyboardEvent {
	typ := e.Type()
	if typ < EventKeyDown || typ > EventKeyUp {
		return nil
	}
	ev := (*C.SDL_KeyboardEvent)(unsafe.Pointer(e))
	return &KeyboardEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		WindowID(ev.windowID),
		KeyboardID(ev.which),
		Scancode(ev.scancode),
		Keycode(ev.key),
		Keymod(ev.mod),
		uint16(ev.raw),
		bool(ev.down),
		bool(ev.repeat),
	}
}

// TextEditing returns the event converted to a [TextEditingEvent],
// or nil if the event type does not match
func (e *Event) TextEditing() *TextEditingEvent {
	typ := e.Type()
	if typ != EventTextEditing {
		return nil
	}
	ev := (*C.SDL_TextEditingEvent)(unsafe.Pointer(e))
	return &TextEditingEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		WindowID(ev.windowID),
		C.GoString(ev.text),
		int(ev.start),
		int(ev.length),
	}
}

// TextEditingCandidates returns the event converted to a [TextEditingCandi],
// or nil if the event type does not match
func (e *Event) TextEditingCandidates() *TextEditingCandidatesEvent {
	typ := e.Type()
	if typ != EventTextEditingCandidates {
		return nil
	}
	ev := (*C.SDL_TextEditingCandidatesEvent)(unsafe.Pointer(e))
	candidates := make([]string, ev.num_candidates)
	for i, c := range unsafe.Slice(ev.candidates, ev.num_candidates) {
		candidates[i] = C.GoString(c)
	}
	return &TextEditingCandidatesEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		WindowID(ev.windowID),
		candidates,
		int32(ev.selected_candidate),
		bool(ev.horizontal),
	}
}

// TextInput returns the event converted to a [TextInputEvent],
// or nil if the event type does not match
func (e *Event) TextInput() *TextInputEvent {
	typ := e.Type()
	if typ != EventTextInput {
		return nil
	}
	ev := (*C.SDL_TextInputEvent)(unsafe.Pointer(e))
	return &TextInputEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		WindowID(ev.windowID),
		C.GoString(ev.text),
	}
}

// MouseDevice returns the event converted to a [MouseDeviceEvent],
// or nil if the event type does not match
func (e *Event) MouseDevice() *MouseDeviceEvent {
	typ := e.Type()
	if typ < EventMouseAdded || typ > EventMouseRemoved {
		return nil
	}
	ev := (*C.SDL_MouseDeviceEvent)(unsafe.Pointer(e))
	return &MouseDeviceEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		MouseID(ev.which),
	}
}

// MouseMotion returns the event converted to a [MouseMotionEvent],
// or nil if the event type does not match
func (e *Event) MouseMotion() *MouseMotionEvent {
	typ := e.Type()
	if typ != EventMouseMotion {
		return nil
	}
	ev := (*C.SDL_MouseMotionEvent)(unsafe.Pointer(e))
	return &MouseMotionEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		WindowID(ev.windowID),
		MouseID(ev.which),
		MouseButtonFlags(ev.state),
		float32(ev.x),
		float32(ev.y),
		float32(ev.xrel),
		float32(ev.yrel),
	}
}

// MouseButton returns the event converted to a [MouseButtonEvent],
// or nil if the event type does not match
func (e *Event) MouseButton() *MouseButtonEvent {
	typ := e.Type()
	if typ < EventMouseButtonDown || typ > EventMouseButtonUp {
		return nil
	}
	ev := (*C.SDL_MouseButtonEvent)(unsafe.Pointer(e))
	return &MouseButtonEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		WindowID(ev.windowID),
		MouseID(ev.which),
		byte(ev.button),
		bool(ev.down),
		byte(ev.clicks),
		float32(ev.x),
		float32(ev.y),
	}
}

// MouseWheel returns the event converted to a [MouseWheelEvent],
// or nil if the event type does not match
func (e *Event) MouseWheel() *MouseWheelEvent {
	typ := e.Type()
	if typ != EventMouseWheel {
		return nil
	}
	ev := (*C.SDL_MouseWheelEvent)(unsafe.Pointer(e))
	return &MouseWheelEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		WindowID(ev.windowID),
		MouseID(ev.which),
		float32(ev.x),
		float32(ev.y),
		MouseWheelDirection(ev.direction),
		float32(ev.mouse_x),
		float32(ev.mouse_y),
	}
}

// JoyAxis returns the event converted to a [JoyAxisEvent],
// or nil if the event type does not match
func (e *Event) JoyAxis() *JoyAxisEvent {
	typ := e.Type()
	if typ != EventJoystickAxisMotion {
		return nil
	}
	ev := (*C.SDL_JoyAxisEvent)(unsafe.Pointer(e))
	return &JoyAxisEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		JoystickID(ev.which),
		byte(ev.axis),
		int16(ev.value),
	}
}

// JoyBall returns the event converted to a [JoyBallEvent],
// or nil if the event type does not match
func (e *Event) JoyBall() *JoyBallEvent {
	typ := e.Type()
	if typ != EventJoystickBallMotion {
		return nil
	}
	ev := (*C.SDL_JoyBallEvent)(unsafe.Pointer(e))
	return &JoyBallEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		JoystickID(ev.which),
		byte(ev.ball),
		int16(ev.xrel),
		int16(ev.yrel),
	}
}

// JoyHat returns the event converted to a [JoyHatEvent],
// or nil if the event type does not match
func (e *Event) JoyHat() *JoyHatEvent {
	typ := e.Type()
	if typ != EventJoystickHatMotion {
		return nil
	}
	ev := (*C.SDL_JoyHatEvent)(unsafe.Pointer(e))
	return &JoyHatEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		JoystickID(ev.which),
		byte(ev.hat),
		JoystickHatState(ev.value),
	}
}

// JoyButton returns the event converted to a [JoyButtonEvent],
// or nil if the event type does not match
func (e *Event) JoyButton() *JoyButtonEvent {
	typ := e.Type()
	if typ < EventJoystickButtonDown || typ > EventJoystickButtonUp {
		return nil
	}
	ev := (*C.SDL_JoyButtonEvent)(unsafe.Pointer(e))
	return &JoyButtonEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		JoystickID(ev.which),
		byte(ev.button),
		bool(ev.down),
	}
}

// JoyDevice returns the event converted to a [JoyDeviceEvent],
// or nil if the event type does not match
func (e *Event) JoyDevice() *JoyDeviceEvent {
	typ := e.Type()
	if typ < EventJoystickAxisMotion || typ > EventJoystickUpdateComplete {
		return nil
	}
	ev := (*C.SDL_JoyDeviceEvent)(unsafe.Pointer(e))
	return &JoyDeviceEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		JoystickID(ev.which),
	}
}

// JoyBattery returns the event converted to a [JoyBatteryEvent],
// or nil if the event type does not match
func (e *Event) JoyBattery() *JoyBatteryEvent {
	typ := e.Type()
	if typ != EventJoystickBatteryUpdated {
		return nil
	}
	ev := (*C.SDL_JoyBatteryEvent)(unsafe.Pointer(e))
	return &JoyBatteryEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		JoystickID(ev.which),
		PowerState(ev.state),
		int(ev.percent),
	}
}

// GamepadAxis returns the event converted to a [GamepadAxisEvent],
// or nil if the event type does not match
func (e *Event) GamepadAxis() *GamepadAxisEvent {
	typ := e.Type()
	if typ != EventGamepadAxisMotion {
		return nil
	}
	ev := (*C.SDL_GamepadAxisEvent)(unsafe.Pointer(e))
	return &GamepadAxisEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		JoystickID(ev.which),
		GamepadAxis(ev.axis),
		int16(ev.value),
	}
}

// GamepadButton returns the event converted to a [GamepadButtonEve],
// or nil if the event type does not match
func (e *Event) GamepadButton() *GamepadButtonEvent {
	typ := e.Type()
	if typ < EventGamepadButtonDown || typ > EventGamepadButtonUp {
		return nil
	}
	ev := (*C.SDL_GamepadButtonEvent)(unsafe.Pointer(e))
	return &GamepadButtonEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		JoystickID(ev.which),
		GamepadButton(ev.button),
		bool(ev.down),
	}
}

// GamepadDevice returns the event converted to a [GamepadDeviceEve],
// or nil if the event type does not match
func (e *Event) GamepadDevice() *GamepadDeviceEvent {
	typ := e.Type()
	if typ < EventGamepadAxisMotion || typ > EventGamepadSteamHandleUpdated {
		return nil
	}
	ev := (*C.SDL_GamepadDeviceEvent)(unsafe.Pointer(e))
	return &GamepadDeviceEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		JoystickID(ev.which),
	}
}

// GamepadTouchpad returns the event converted to a [GamepadTouchpadE],
// or nil if the event type does not match
func (e *Event) GamepadTouchpad() *GamepadTouchpadEvent {
	typ := e.Type()
	if typ < EventGamepadTouchpadDown || typ > EventGamepadTouchpadUp {
		return nil
	}
	ev := (*C.SDL_GamepadTouchpadEvent)(unsafe.Pointer(e))
	return &GamepadTouchpadEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		JoystickID(ev.which),
		int(ev.touchpad),
		int(ev.finger),
		float32(ev.x),
		float32(ev.y),
		float32(ev.pressure),
	}
}

// GamepadSensor returns the event converted to a [GamepadSensorEve],
// or nil if the event type does not match
func (e *Event) GamepadSensor() *GamepadSensorEvent {
	typ := e.Type()
	if typ != EventGamepadSensorUpdate {
		return nil
	}
	ev := (*C.SDL_GamepadSensorEvent)(unsafe.Pointer(e))
	return &GamepadSensorEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		JoystickID(ev.which),
		int(ev.sensor),
		[3]float32{float32(ev.data[0]), float32(ev.data[1]), float32(ev.data[2])},
		uint64(ev.sensor_timestamp),
	}
}

// AudioDevice returns the event converted to a [AudioDeviceEvent],
// or nil if the event type does not match
func (e *Event) AudioDevice() *AudioDeviceEvent {
	typ := e.Type()
	if typ < EventAudioDeviceAdded || typ > EventAudioDeviceFormatChanged {
		return nil
	}
	ev := (*C.SDL_AudioDeviceEvent)(unsafe.Pointer(e))
	return &AudioDeviceEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		AudioDeviceID(ev.which),
		bool(ev.recording),
	}
}

// CameraDevice returns the event converted to a [CameraDeviceEven],
// or nil if the event type does not match
func (e *Event) CameraDevice() *CameraDeviceEvent {
	typ := e.Type()
	if typ < EventCameraDeviceAdded || typ > EventCameraDeviceDenied {
		return nil
	}
	ev := (*C.SDL_CameraDeviceEvent)(unsafe.Pointer(e))
	return &CameraDeviceEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		CameraID(ev.which),
	}
}

// Render returns the event converted to a [RenderEvent],
// or nil if the event type does not match
func (e *Event) Render() *RenderEvent {
	typ := e.Type()
	if typ < EventRenderTargetsReset || typ > EventRenderDeviceLost {
		return nil
	}
	ev := (*C.SDL_RenderEvent)(unsafe.Pointer(e))
	return &RenderEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		WindowID(ev.windowID),
	}
}

// TouchFinger returns the event converted to a [TouchFingerEvent],
// or nil if the event type does not match
func (e *Event) TouchFinger() *TouchFingerEvent {
	typ := e.Type()
	if typ < EventFingerDown || typ > EventFingerCanceled {
		return nil
	}
	ev := (*C.SDL_TouchFingerEvent)(unsafe.Pointer(e))
	return &TouchFingerEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		TouchID(ev.touchID),
		FingerID(ev.fingerID),
		float32(ev.x),
		float32(ev.y),
		float32(ev.dx),
		float32(ev.dy),
		float32(ev.pressure),
		WindowID(ev.windowID),
	}
}

// PenProximity returns the event converted to a [PenProximityEven],
// or nil if the event type does not match
func (e *Event) PenProximity() *PenProximityEvent {
	typ := e.Type()
	if typ < EventPenProximityIn || typ > EventPenProximityOut {
		return nil
	}
	ev := (*C.SDL_PenProximityEvent)(unsafe.Pointer(e))
	return &PenProximityEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		WindowID(ev.windowID),
		PenID(ev.which),
	}
}

// PenMotion returns the event converted to a [PenMotionEvent],
// or nil if the event type does not match
func (e *Event) PenMotion() *PenMotionEvent {
	typ := e.Type()
	if typ != EventPenMotion {
		return nil
	}
	ev := (*C.SDL_PenMotionEvent)(unsafe.Pointer(e))
	return &PenMotionEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		WindowID(ev.windowID),
		PenID(ev.which),
		PenInputFlags(ev.pen_state),
		float32(ev.x),
		float32(ev.y),
	}
}

// PenTouch returns the event converted to a [PenTouchEvent],
// or nil if the event type does not match
func (e *Event) PenTouch() *PenTouchEvent {
	typ := e.Type()
	if typ < EventPenDown || typ > EventPenUp {
		return nil
	}
	ev := (*C.SDL_PenTouchEvent)(unsafe.Pointer(e))
	return &PenTouchEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		WindowID(ev.windowID),
		PenID(ev.which),
		PenInputFlags(ev.pen_state),
		float32(ev.x),
		float32(ev.y),
		bool(ev.eraser),
		bool(ev.down),
	}
}

// PenButton returns the event converted to a [PenButtonEvent],
// or nil if the event type does not match
func (e *Event) PenButton() *PenButtonEvent {
	typ := e.Type()
	if typ < EventPenButtonDown || typ > EventPenButtonUp {
		return nil
	}
	ev := (*C.SDL_PenButtonEvent)(unsafe.Pointer(e))
	return &PenButtonEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		WindowID(ev.windowID),
		PenID(ev.which),
		PenInputFlags(ev.pen_state),
		float32(ev.x),
		float32(ev.y),
		byte(ev.button),
		bool(ev.down),
	}
}

// PenAxis returns the event converted to a [PenAxisEvent],
// or nil if the event type does not match
func (e *Event) PenAxis() *PenAxisEvent {
	typ := e.Type()
	if typ != EventPenAxis {
		return nil
	}
	ev := (*C.SDL_PenAxisEvent)(unsafe.Pointer(e))
	return &PenAxisEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		WindowID(ev.windowID),
		PenID(ev.which),
		PenInputFlags(ev.pen_state),
		float32(ev.x),
		float32(ev.y),
		PenAxis(ev.axis),
		float32(ev.value),
	}
}

// Drop returns the event converted to a [DropEvent],
// or nil if the event type does not match
func (e *Event) Drop() *DropEvent {
	typ := e.Type()
	if typ < EventDropFile || typ > EventDropPosition {
		return nil
	}
	ev := (*C.SDL_DropEvent)(unsafe.Pointer(e))
	return &DropEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		WindowID(ev.windowID),
		float32(ev.x),
		float32(ev.y),
		C.GoString(ev.source),
		C.GoString(ev.data),
	}
}

// Clipboard returns the event converted to a [ClipboardEvent],
// or nil if the event type does not match
func (e *Event) Clipboard() *ClipboardEvent {
	typ := e.Type()
	if typ != EventClipboardUpdate {
		return nil
	}
	ev := (*C.SDL_ClipboardEvent)(unsafe.Pointer(e))
	mimeTypes := make([]string, ev.num_mime_types)
	for i, c := range unsafe.Slice(ev.mime_types, ev.num_mime_types) {
		mimeTypes[i] = C.GoString(c)
	}
	return &ClipboardEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		bool(ev.owner),
		mimeTypes,
	}
}

// Sensor returns the event converted to a [SensorEvent],
// or nil if the event type does not match
func (e *Event) Sensor() *SensorEvent {
	typ := e.Type()
	if typ != EventSensorUpdate {
		return nil
	}
	ev := (*C.SDL_SensorEvent)(unsafe.Pointer(e))
	return &SensorEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		SensorID(ev.which),
		[6]float32{float32(ev.data[0]), float32(ev.data[1]), float32(ev.data[2]), float32(ev.data[3]), float32(ev.data[4]), float32(ev.data[5])},
		uint64(ev.sensor_timestamp),
	}
}

// Sensor returns the event converted to a [QuitEvent],
// or nil if the event type does not match
func (e *Event) Quit() *QuitEvent {
	typ := e.Type()
	if typ != EventQuit {
		return nil
	}
	ev := (*C.SDL_QuitEvent)(unsafe.Pointer(e))
	return &QuitEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
	}
}

// User returns the event converted to a [UserEvent],
// or nil if the event type does not match
func (e *Event) User() *UserEvent {
	typ := e.Type()
	if typ < EventUser || typ >= EventLast {
		return nil
	}
	ev := (*C.SDL_UserEvent)(unsafe.Pointer(e))
	return &UserEvent{
		EventType(ev._type),
		uint64(ev.timestamp),
		WindowID(ev.windowID),
		int(ev.code),
		ev.data1,
		ev.data2,
	}
}

func (e *Event) SetUser(ev *UserEvent) {
	*(*C.SDL_UserEvent)(unsafe.Pointer(e)) = C.SDL_UserEvent{
		_type:     C.Uint32(ev.Type),
		timestamp: C.Uint64(ev.Timestamp),
		windowID:  C.SDL_WindowID(ev.WindowID),
		code:      C.Sint32(ev.Code),
		data1:     ev.Data1,
		data2:     ev.Data2,
	}
}

// Pump the event loop, gathering events from the input devices.
//
// This function updates the event queue and internal input device state.
//
// PumpEvents gathers all the pending input information from devices and
// places it in the event queue. Without calls to PumpEvents no events
// would ever be placed on the queue. Often the need for calls to
// PumpEvents is hidden from the user since [PollEvent] and
// [WaitEvent] implicitly call PumpEvents. However, if you are not
// polling or waiting for events (e.g. you are filtering them), then you must
// call PumpEvents to force an event queue update.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PumpEvents
func PumpEvents() {
	C.SDL_PumpEvents()
}

// The type of action to request from [PeepEvents].
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_EventAction
type EventAction uint32

const (
	AddEvent  EventAction = iota // Add events to the back of the queue.
	PeekEvent                    // Check but don't remove events from the queue front.
	GetEvent                     // Retrieve/remove events from the front of the queue.
)

// Check the event queue for messages and optionally return them.
//
// action may be any of the following:
//
//   - [AddEvent]: up to len(events) events will be added to the back of the
//     event queue.
//   - [PeekEvent]: len(events) events at the front of the event queue,
//     within the specified minimum and maximum type, will be returned to the
//     caller and will _not_ be removed from the queue. If you pass nil for
//     events, the total number of matching events will be returned.
//   - [GetEvent]: up to len(events) events at the front of the event queue,
//     within the specified minimum and maximum type, will be returned to the
//     caller and will be removed from the queue.
//
// You may have to call [PumpEvents] before calling this function.
// Otherwise, the events may not be ready to be filtered when you call
// PeepEvents.
//
// events: destination buffer for the retrieved events, may be nil to
// leave the events in the queue and return the number of events
// that would have been stored.
//
// numevents: if action is [AddEvent], the number of events to add
// back to the event queue; if action is [PeekEvent] or
// [GetEvent], the maximum number of events to retrieve.
//
// action: action to take.
//
// minType: minimum value of the event type to be considered;
// [EventFirst] is a safe choice.
//
// maxType: maximum value of the event type to be considered;
// [EventLast] is a safe choice.
//
// Returns the number of events actually stored or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PeepEvents
func PeepEvents(events []Event, action EventAction, minType, maxType EventType) (int, error) {
	n := C.SDL_PeepEvents((*C.SDL_Event)(unsafe.Pointer(unsafe.SliceData(events))), (C.int)(len(events)), (C.SDL_EventAction)(action), (C.Uint32)(minType), (C.Uint32)(maxType))
	if n == -1 {
		return 0, getError()
	}
	return int(n), nil
}

// Check for the existence of a certain event type in the event queue.
//
// If you need to check for a range of event types, use [HasEvents]
// instead.
//
// typ: the type of event to be queried; see [EventType] for details.
//
// Returns true if events matching typ are present, or false if events
// matching typ are not present.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasEvent
func HasEvent(typ EventType) bool {
	return (bool)(C.SDL_HasEvent((C.Uint32)(typ)))
}

// Check for the existence of certain event types in the event queue.
//
// If you need to check for a single event type, use [HasEvent] instead.
//
// minType: the low end of event type to be queried, inclusive; see
// [EventType] for details.
//
// maxType: the high end of event type to be queried, inclusive; see
// [EventType] for details.
//
// Returns true if events with type >= minType and <= maxType are
// present, or false if not.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasEvents
func HasEvents(minType, maxType EventType) bool {
	return (bool)(C.SDL_HasEvents((C.Uint32)(minType), (C.Uint32)(maxType)))
}

// Clear events of a specific type from the event queue.
//
// This will unconditionally remove any events from the queue that match
// typ. If you need to remove a range of event types, use [FlushEvents]
// instead.
//
// It's also normal to just ignore events you don't care about in your event
// loop without calling this function.
//
// This function only affects currently queued events. If you want to make
// sure that all pending OS events are flushed, you can call [PumpEvents]
// on the main thread immediately before the flush call.
//
// If you have user events with custom data that needs to be freed, you should
// use [PeepEvents] to remove and clean up those events before calling
// this function.
//
// typ: the type of event to be cleared; see [EventType] for details.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_FlushEvent
func FlushEvent(typ EventType) {
	C.SDL_FlushEvent((C.Uint32)(typ))
}

// Clear events of a range of types from the event queue.
//
// This will unconditionally remove any events from the queue that are in the
// range of minType to maxType, inclusive. If you need to remove a single
// event type, use [FlushEvent] instead.
//
// It's also normal to just ignore events you don't care about in your event
// loop without calling this function.
//
// This function only affects currently queued events. If you want to make
// sure that all pending OS events are flushed, you can call [PumpEvents]
// on the main thread immediately before the flush call.
//
// minType: the low end of event type to be cleared, inclusive; see
// [EventType] for details.
//
// maxType: the high end of event type to be cleared, inclusive; see
// [EventType] for details.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_FlushEvents
func FlushEvents(minType, maxType EventType) {
	C.SDL_FlushEvents((C.Uint32)(minType), (C.Uint32)(maxType))
}

// Poll for currently pending events.
//
// If event is not nil, the next event is removed from the queue and stored
// in the [Event] structure pointed to by event. The 1 returned refers to
// this event, immediately stored in the SDL Event structure -- not an event
// to follow.
//
// If event is nil, it simply returns true if there is an event in the queue,
// but will not remove it from the queue.
//
// As this function may implicitly call [PumpEvents], you can only call
// this function in the thread that set the video mode.
//
// PollEvent is the favored way of receiving system events since it can
// be done from the main loop and does not suspend the main loop while waiting
// on an event to be posted.
//
// The common practice is to fully process the event queue once every frame,
// usually as a first step before updating the game's state:
//
//	while (game_is_still_running) {
//		SDL_Event event;
//		while (SDL_PollEvent(&event)) {  // poll until all events are handled!
//			// decide what to do with this event.
//		}
//		// update game state, draw the current frame
//	}
//
// event: the [Event] structure to be filled with the next event from
// the queue, or nil.
//
// Returns true if this got an event or false if there are none available.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PollEvent
func PollEvent(event *Event) bool {
	var cevent *C.SDL_Event
	if event != nil {
		cevent = &event.internal
	}
	return (bool)(C.SDL_PollEvent(cevent))
}

// Wait indefinitely for the next available event.
//
// If event is not nil, the next event is removed from the queue and stored
// in the [Event] structure pointed to by event.
//
// As this function may implicitly call [PumpEvents], you can only call
// this function in the thread that initialized the video subsystem.
//
// event: the [Event] structure to be filled in with the next event
// from the queue, or nil.
//
// Returns nil on success or an error if there was an error while waiting for
// events.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WaitEvent
func WaitEvent(event *Event) error {
	var cevent *C.SDL_Event
	if event != nil {
		cevent = &event.internal
	}
	if !C.SDL_WaitEvent(cevent) {
		return getError()
	}
	return nil
}

// Wait until the specified timeout (in milliseconds) for the next available
// event.
//
// If event is not nil, the next event is removed from the queue and stored
// in the [Event] structure pointed to by event.
//
// As this function may implicitly call [PumpEvents], you can only call
// this function in the thread that initialized the video subsystem.
//
// The timeout is not guaranteed, the actual wait time could be longer due to
// system scheduling.
//
// event: the [Event] structure to be filled in with the next event
// from the queue, or nil.
//
// timeoutMS: the maximum number of milliseconds to wait for the next
// available event.
//
// Returns true if this got an event or false if the timeout elapsed without
// any events available.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WaitEventTimeout
func WaitEventTimeout(event *Event, timeoutMS int32) bool {
	var cevent *C.SDL_Event
	if event != nil {
		cevent = &event.internal
	}
	return (bool)(C.SDL_WaitEventTimeout(cevent, (C.Sint32)(timeoutMS)))
}

// Add an event to the event queue.
//
// The event queue can actually be used as a two way communication channel.
// Not only can events be read from the queue, but the user can also push
// their own events onto it. event is a pointer to the event structure you
// wish to push onto the queue. The event is copied into the queue, and the
// caller may dispose of the memory pointed to after PushEvent returns.
//
// Note: Pushing device input events onto the queue doesn't modify the state
// of the device within SDL.
//
// Note: Events pushed onto the queue with PushEvent get passed through
// the event filter but events added with [PeepEvents] do not.
//
// For pushing application-specific events, please use [RegisterEvents] to
// get an event type that does not conflict with other code that also wants
// its own custom event types.
//
// event: the [Event] to be added to the queue.
//
// Returns nil on success, an error if the event was filtered or on failure.
// A common reason for error is the event queue being full.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PushEvent
func PushEvent(event *Event) error {
	if !C.SDL_PushEvent(&event.internal) {
		return getError()
	}
	return nil
}

// A function pointer used for callbacks that watch the event queue.
//
// event: the event that triggered the callback.
//
// Returns true to permit event to be added to the queue, and false to
// disallow it. When used with [AddEventWatch], the return value is
// ignored.
//
// SDL may call this callback at any time from any thread; the
// application is responsible for locking resources the callback
// touches that need to be protected.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_EventFilter
type EventFilter func(event *Event) bool

//export cb_EventFilter
func cb_EventFilter(userdata uintptr, event *C.SDL_Event) bool {
	return cgo.Handle(userdata).Value().(EventFilter)(&Event{*event})
}

// EventWatch is a struct containing an EventFilter.
//
// Since go functions can not be compared, the only way to remove an
// EventWatch is to pass a pointer to the same struct to RemoveEventWatch
type EventWatch struct {
	Filter EventFilter
}

// Set up a filter to process all events before they are added to the internal
// event queue.
//
// If you just want to see events without modifying them or preventing them
// from being queued, you should use [AddEventWatch] instead.
//
// If the filter function returns true when called, then the event will be
// added to the internal queue. If it returns false, then the event will be
// dropped from the queue, but the internal state will still be updated. This
// allows selective filtering of dynamically arriving events.
//
// **WARNING**: Be very careful of what you do in the event filter function,
// as it may run in a different thread!
//
// On platforms that support it, if the quit event is generated by an
// interrupt signal (e.g. pressing Ctrl-C), it will be delivered to the
// application at the next event poll.
//
// Note: Disabled events never make it to the event filter function; see
// [SetEventEnabled].
//
// Note: Events pushed onto the queue with [PushEvent] get passed through
// the event filter, but events pushed onto the queue with [PeepEvents] do
// not.
//
// filter: an [EventFilter] function to call when an event happens.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetEventFilter
func SetEventFilter(filter EventFilter) {
	eventFilterLock.Lock()
	if eventFilter != 0 {
		eventFilter.Delete()
	}
	eventFilter = cgo.NewHandle(filter)
	C.wrap_SDL_SetEventFilter(C.uintptr_t(eventFilter))
	eventFilterLock.Unlock()
}

var eventFilter cgo.Handle
var eventFilterLock sync.Mutex

// Query the current event filter.
//
// This function can be used to "chain" filters, by saving the existing filter
// before replacing it with a function that will call that saved filter.
//
// filter: the current callback function will be stored here.
//
// Returns true on success or false if there is no event filter set.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetEventFilter
func GetEventFilter() EventFilter {
	eventFilterLock.Lock()
	if eventFilter == 0 {
		eventFilterLock.Unlock()
		return nil
	}
	filter := eventFilter.Value().(EventFilter)
	eventFilterLock.Unlock()
	return filter
}

// Add a callback to be triggered when an event is added to the event queue.
//
// filter will be called when an event happens, and its return value is
// ignored.
//
// **WARNING**: Be very careful of what you do in the event filter function,
// as it may run in a different thread!
//
// If the quit event is generated by a signal (e.g. SIGINT), it will bypass
// the internal queue and be delivered to the watch callback immediately, and
// arrive at the next event poll.
//
// Note: the callback is called for events posted by the user through
// [PushEvent], but not for disabled events, nor for events by a filter
// callback set with [SetEventFilter], nor for events posted by the user
// through [PeepEvents].
//
// watch: a struct containing an [EventFilter] function to call when an event
// happens.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AddEventWatch
func AddEventWatch(watch *EventWatch) error {
	h := cgo.NewHandle(watch.Filter)
	eventWatchesLock.Lock()
	if !C.wrap_SDL_AddEventWatch(C.uintptr_t(h)) {
		eventWatchesLock.Unlock()
		h.Delete()
		return getError()
	}
	eventWatches[watch] = h
	eventWatchesLock.Unlock()
	return nil
}

var eventWatches = map[*EventWatch]cgo.Handle{}
var eventWatchesLock sync.Mutex

// Remove an event watch callback added with [AddEventWatch].
//
// This function takes the same input as [AddEventWatch] to identify and
// delete the corresponding callback.
//
// watch: the same pointer originally passed to [AddEventWatch].
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RemoveEventWatch
func RemoveEventWatch(watch *EventWatch) {
	eventWatchesLock.Lock()
	h, ok := eventWatches[watch]
	delete(eventWatches, watch)
	eventWatchesLock.Unlock()
	if !ok {
		return
	}
	C.wrap_SDL_RemoveEventWatch(C.uintptr_t(h))
	h.Delete()
}

// Run a specific filter function on the current event queue, removing any
// events for which the filter returns false.
//
// See [SetEventFilter] for more information. Unlike [SetEventFilter],
// this function does not change the filter permanently, it only uses the
// supplied filter until this function returns.
//
// filter: the [EventFilter] function to call when an event happens.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_FilterEvents
func FilterEvents(filter EventFilter) {
	h := cgo.NewHandle(filter)
	C.wrap_SDL_FilterEvents(C.uintptr_t(h))
	h.Delete()
}

// Set the state of processing events by type.
//
// type: the type of event; see [EventType] for details.
//
// enabled: whether to process the event or not.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetEventEnabled
func SetEventEnabled(typ EventType, enabled bool) {
	C.SDL_SetEventEnabled((C.Uint32)(typ), (C.bool)(enabled))
}

// Query the state of processing events by type.
//
// type: the type of event; see [EventType] for details.
//
// Returns true if the event is being processed, false otherwise.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_EventEnabled
func EventEnabled(typ EventType) bool {
	return (bool)(C.SDL_EventEnabled((C.Uint32)(typ)))
}

// Allocate a set of user-defined events, and return the beginning event
// number for that set of events.
//
// numevents: the number of events to be allocated.
//
// Returns the beginning event number, or 0 if numevents is invalid or if
// there are not enough user-defined events left.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RegisterEvents
func RegisterEvents(numevents int) EventType {
	return (EventType)(C.SDL_RegisterEvents((C.int)(numevents)))
}

// Get window associated with an event.
//
// event: an event containing a WindowID.
//
// Returns the associated window on success or nil if there is none.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetWindowFromEvent
func GetWindowFromEvent(event *Event) *Window {
	return (*Window)(C.SDL_GetWindowFromEvent(&event.internal))
}
