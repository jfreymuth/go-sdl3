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
// #cgo noescape SDL_Init
// #cgo nocallback SDL_Init
// #cgo noescape SDL_InitSubSystem
// #cgo nocallback SDL_InitSubSystem
// #cgo noescape SDL_QuitSubSystem
// #cgo noescape SDL_WasInit
// #cgo nocallback SDL_WasInit
// #cgo noescape SDL_Quit
// #cgo noescape SDL_IsMainThread
// #cgo nocallback SDL_IsMainThread
// #cgo noescape wrap_SDL_RunOnMainThread
// #cgo noescape SDL_SetAppMetadata
// #cgo nocallback SDL_SetAppMetadata
// #cgo noescape SDL_SetAppMetadataProperty
// #cgo nocallback SDL_SetAppMetadataProperty
// #cgo noescape SDL_GetAppMetadataProperty
// #cgo nocallback SDL_GetAppMetadataProperty
// #include <SDL3/SDL.h>
// extern bool wrap_SDL_RunOnMainThread(uintptr_t userdata, bool wait);
import "C"
import (
	"runtime/cgo"
)

// # CategoryInit
//
// All SDL programs need to initialize the library before starting to work
// with it.
//
// Almost everything can simply call SDL_Init() near startup, with a handful
// of flags to specify subsystems to touch. These are here to make sure SDL
// does not even attempt to touch low-level pieces of the operating system
// that you don't intend to use. For example, you might be using SDL for video
// and input but chose an external library for audio, and in this case you
// would just need to leave off the `SDL_INIT_AUDIO` flag to make sure that
// external library has complete control.
//
// Most apps, when terminating, should call SDL_Quit(). This will clean up
// (nearly) everything that SDL might have allocated, and crucially, it'll
// make sure that the display's resolution is back to what the user expects if
// you had previously changed it for your game.
//
// SDL3 apps are strongly encouraged to call SDL_SetAppMetadata() at startup
// to fill in details about the program. This is completely optional, but it
// helps in small ways (we can provide an About dialog box for the macOS menu,
// we can name the app in the system's audio mixer, etc). Those that want to
// provide a _lot_ of information should look at the more-detailed
// SDL_SetAppMetadataProperty().

// Initialization flags for SDL_Init and/or SDL_InitSubSystem
//
// These are the flags which may be passed to SDL_Init(). You should specify
// the subsystems which you will be using in your application.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_InitFlags
type InitFlags uint32

const (
	InitAudio    InitFlags = 0x00000010 //`SDL_INIT_AUDIO` implies `SDL_INIT_EVENTS`
	InitVideo    InitFlags = 0x00000020 //`SDL_INIT_VIDEO` implies `SDL_INIT_EVENTS`, should be initialized on the main thread
	InitJoystick InitFlags = 0x00000200 //`SDL_INIT_JOYSTICK` implies `SDL_INIT_EVENTS`, should be initialized on the same thread as SDL_INIT_VIDEO on Windows if you don't set SDL_HINT_JOYSTICK_THREAD
	InitHaptic   InitFlags = 0x00001000
	InitGamepad  InitFlags = 0x00002000 //`SDL_INIT_GAMEPAD` implies `SDL_INIT_JOYSTICK`
	InitEvents   InitFlags = 0x00004000
	InitSensor   InitFlags = 0x00008000 //`SDL_INIT_SENSOR` implies `SDL_INIT_EVENTS`
	InitCamera   InitFlags = 0x00010000 //`SDL_INIT_CAMERA` implies `SDL_INIT_EVENTS`
)

// Initialize the SDL library.
//
// SDL_Init() simply forwards to calling SDL_InitSubSystem(). Therefore, the
// two may be used interchangeably. Though for readability of your code
// SDL_InitSubSystem() might be preferred.
//
// The file I/O (for example: SDL_IOFromFile) and threading (SDL_CreateThread)
// subsystems are initialized by default. Message boxes
// (SDL_ShowSimpleMessageBox) also attempt to work without initializing the
// video subsystem, in hopes of being useful in showing an error dialog when
// SDL_Init fails. You must specifically initialize other subsystems if you
// use them in your application.
//
// Logging (such as SDL_Log) works without initialization, too.
//
// `flags` may be any of the following OR'd together:
//
// - `SDL_INIT_AUDIO`: audio subsystem; automatically initializes the events
// subsystem
// - `SDL_INIT_VIDEO`: video subsystem; automatically initializes the events
// subsystem, should be initialized on the main thread.
// - `SDL_INIT_JOYSTICK`: joystick subsystem; automatically initializes the
// events subsystem
// - `SDL_INIT_HAPTIC`: haptic (force feedback) subsystem
// - `SDL_INIT_GAMEPAD`: gamepad subsystem; automatically initializes the
// joystick subsystem
// - `SDL_INIT_EVENTS`: events subsystem
// - `SDL_INIT_SENSOR`: sensor subsystem; automatically initializes the events
// subsystem
// - `SDL_INIT_CAMERA`: camera subsystem; automatically initializes the events
// subsystem
//
// Subsystem initialization is ref-counted, you must call SDL_QuitSubSystem()
// for each SDL_InitSubSystem() to correctly shutdown a subsystem manually (or
// call SDL_Quit() to force shutdown). If a subsystem is already loaded then
// this call will increase the ref-count and return.
//
// Consider reporting some basic metadata about your application before
// calling SDL_Init, using either SDL_SetAppMetadata() or
// SDL_SetAppMetadataProperty().
//
// flags: subsystem initialization flags.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Init
func Init(flags InitFlags) error {
	if !C.SDL_Init((C.SDL_InitFlags)(flags)) {
		return getError()
	}
	return nil
}

// Compatibility function to initialize the SDL library.
//
// This function and SDL_Init() are interchangeable.
//
// flags: any of the flags used by SDL_Init(); see SDL_Init for details.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_InitSubSystem
func InitSubSystem(flags InitFlags) error {
	if !C.SDL_InitSubSystem((C.SDL_InitFlags)(flags)) {
		return getError()
	}
	return nil
}

// Shut down specific SDL subsystems.
//
// You still need to call SDL_Quit() even if you close all open subsystems
// with SDL_QuitSubSystem().
//
// flags: any of the flags used by SDL_Init(); see SDL_Init for details.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_QuitSubSystem
func QuitSubSystem(flags InitFlags) {
	C.SDL_QuitSubSystem((C.SDL_InitFlags)(flags))
}

// Get a mask of the specified subsystems which are currently initialized.
//
// flags: any of the flags used by SDL_Init(); see SDL_Init for details.
//
// Returns a mask of all initialized subsystems if `flags` is 0, otherwise it
// returns the initialization status of the specified subsystems.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WasInit
func WasInit(flags InitFlags) InitFlags {
	return (InitFlags)(C.SDL_WasInit((C.SDL_InitFlags)(flags)))
}

// Clean up all initialized subsystems.
//
// You should call this function even if you have already shutdown each
// initialized subsystem with SDL_QuitSubSystem(). It is safe to call this
// function even in the case of errors in initialization.
//
// You can use this function with atexit() to ensure that it is run when your
// application is shutdown, but it is not wise to do this from a library or
// other dynamically loaded code.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Quit
func Quit() {
	C.SDL_Quit()
}

// Return whether this is the main thread.
//
// On Apple platforms, the main thread is the thread that runs your program's
// main() entry point. On other platforms, the main thread is the one that
// calls SDL_Init(SDL_INIT_VIDEO), which should usually be the one that runs
// your program's main() entry point. If you are using the main callbacks,
// SDL_AppInit(), SDL_AppIterate(), and SDL_AppQuit() are all called on the
// main thread.
//
// Returns true if this thread is the main thread, or false otherwise.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_IsMainThread
func IsMainThread() bool {
	return (bool)(C.SDL_IsMainThread())
}

// Callback run on the main thread.
//
// userdata: an app-controlled pointer that is passed to the callback.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MainThreadCallback
type MainThreadCallback func()

//export cb_MainThreadCallback
func cb_MainThreadCallback(userdata uintptr) {
	h := (cgo.Handle)(userdata)
	h.Value().(MainThreadCallback)()
	h.Delete()
}

// Call a function on the main thread during event processing.
//
// If this is called on the main thread, the callback is executed immediately.
// If this is called on another thread, this callback is queued for execution
// on the main thread during event processing.
//
// Be careful of deadlocks when using this functionality. You should not have
// the main thread wait for the current thread while this function is being
// called with `wait_complete` true.
//
// callback: the callback to call on the main thread.
//
// userdata: a pointer that is passed to `callback`.
//
// wait_complete: true to wait for the callback to complete, false to
// return immediately.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RunOnMainThread
func RunOnMainThread(callback MainThreadCallback, wait_complete bool) error {
	h := cgo.NewHandle(callback)
	ok := C.wrap_SDL_RunOnMainThread(C.uintptr_t(h), (C.bool)(wait_complete))
	if !ok {
		h.Delete()
		return getError()
	}
	return nil
}

// Specify basic metadata about your app.
//
// You can optionally provide metadata about your app to SDL. This is not
// required, but strongly encouraged.
//
// There are several locations where SDL can make use of metadata (an "About"
// box in the macOS menu bar, the name of the app can be shown on some audio
// mixers, etc). Any piece of metadata can be left as NULL, if a specific
// detail doesn't make sense for the app.
//
// This function should be called as early as possible, before SDL_Init.
// Multiple calls to this function are allowed, but various state might not
// change once it has been set up with a previous call to this function.
//
// Passing a NULL removes any previous metadata.
//
// This is a simplified interface for the most important information. You can
// supply significantly more detailed metadata with
// SDL_SetAppMetadataProperty().
//
// appname: The name of the application ("My Game 2: Bad Guy's
// Revenge!").
//
// appversion: The version of the application ("1.0.0beta5" or a git
// hash, or whatever makes sense).
//
// appidentifier: A unique string in reverse-domain format that
// identifies this app ("com.example.mygame2").
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetAppMetadata
func SetAppMetadata(appname string, appversion string, appidentifier string) error {
	if !C.SDL_SetAppMetadata(tmpstring(appname), tmpstring(appversion), tmpstring(appidentifier)) {
		return getError()
	}
	return nil
}

// Specify metadata about your app through a set of properties.
//
// You can optionally provide metadata about your app to SDL. This is not
// required, but strongly encouraged.
//
// There are several locations where SDL can make use of metadata (an "About"
// box in the macOS menu bar, the name of the app can be shown on some audio
// mixers, etc). Any piece of metadata can be left out, if a specific detail
// doesn't make sense for the app.
//
// This function should be called as early as possible, before SDL_Init.
// Multiple calls to this function are allowed, but various state might not
// change once it has been set up with a previous call to this function.
//
// Once set, this metadata can be read using SDL_GetAppMetadataProperty().
//
// These are the supported properties:
//
// - `SDL_PROP_APP_METADATA_NAME_STRING`: The human-readable name of the
// application, like "My Game 2: Bad Guy's Revenge!". This will show up
// anywhere the OS shows the name of the application separately from window
// titles, such as volume control applets, etc. This defaults to "SDL
// Application".
// - `SDL_PROP_APP_METADATA_VERSION_STRING`: The version of the app that is
// running; there are no rules on format, so "1.0.3beta2" and "April 22nd,
// 2024" and a git hash are all valid options. This has no default.
// - `SDL_PROP_APP_METADATA_IDENTIFIER_STRING`: A unique string that
// identifies this app. This must be in reverse-domain format, like
// "com.example.mygame2". This string is used by desktop compositors to
// identify and group windows together, as well as match applications with
// associated desktop settings and icons. If you plan to package your
// application in a container such as Flatpak, the app ID should match the
// name of your Flatpak container as well. This has no default.
// - `SDL_PROP_APP_METADATA_CREATOR_STRING`: The human-readable name of the
// creator/developer/maker of this app, like "MojoWorkshop, LLC"
// - `SDL_PROP_APP_METADATA_COPYRIGHT_STRING`: The human-readable copyright
// notice, like "Copyright (c) 2024 MojoWorkshop, LLC" or whatnot. Keep this
// to one line, don't paste a copy of a whole software license in here. This
// has no default.
// - `SDL_PROP_APP_METADATA_URL_STRING`: A URL to the app on the web. Maybe a
// product page, or a storefront, or even a GitHub repository, for user's
// further information This has no default.
// - `SDL_PROP_APP_METADATA_TYPE_STRING`: The type of application this is.
// Currently this string can be "game" for a video game, "mediaplayer" for a
// media player, or generically "application" if nothing else applies.
// Future versions of SDL might add new types. This defaults to
// "application".
//
// name: the name of the metadata property to set.
//
// value: the value of the property, or NULL to remove that property.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetAppMetadataProperty
func SetAppMetadataProperty(name string, value string) error {
	if !C.SDL_SetAppMetadataProperty(tmpstring(name), tmpstring(value)) {
		return getError()
	}
	return nil
}

const PropAppMetadataNameString = "SDL.app.metadata.name"
const PropAppMetadataVersionString = "SDL.app.metadata.version"
const PropAppMetadataIdentifierString = "SDL.app.metadata.identifier"
const PropAppMetadataCreatorString = "SDL.app.metadata.creator"
const PropAppMetadataCopyrightString = "SDL.app.metadata.copyright"
const PropAppMetadataUrlString = "SDL.app.metadata.url"
const PropAppMetadataTypeString = "SDL.app.metadata.type"

// Get metadata about your app.
//
// This returns metadata previously set using SDL_SetAppMetadata() or
// SDL_SetAppMetadataProperty(). See SDL_SetAppMetadataProperty() for the list
// of available properties and their meanings.
//
// name: the name of the metadata property to get.
//
// Returns the current value of the metadata property, or the default if it
// is not set, NULL for properties with no default.
//
// It is safe to call this function from any thread, although
// the string returned is not protected and could potentially be
// freed if you call SDL_SetAppMetadataProperty() to set that
// property from another thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetAppMetadataProperty
func GetAppMetadataProperty(name string) string {
	return C.GoString(C.SDL_GetAppMetadataProperty(tmpstring(name)))
}
