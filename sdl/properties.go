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
// #cgo noescape SDL_GetGlobalProperties
// #cgo nocallback SDL_GetGlobalProperties
// #cgo noescape SDL_CreateProperties
// #cgo nocallback SDL_CreateProperties
// #cgo noescape SDL_CopyProperties
// #cgo noescape SDL_LockProperties
// #cgo nocallback SDL_LockProperties
// #cgo noescape SDL_UnlockProperties
// #cgo nocallback SDL_UnlockProperties
// #cgo noescape wrap_SDL_SetPointerPropertyWithCleanup
// #cgo noescape SDL_SetPointerProperty
// #cgo noescape SDL_SetStringProperty
// #cgo noescape SDL_SetNumberProperty
// #cgo noescape SDL_SetFloatProperty
// #cgo noescape SDL_SetBooleanProperty
// #cgo noescape SDL_HasProperty
// #cgo nocallback SDL_HasProperty
// #cgo noescape SDL_GetPropertyType
// #cgo nocallback SDL_GetPropertyType
// #cgo noescape SDL_GetPointerProperty
// #cgo nocallback SDL_GetPointerProperty
// #cgo noescape SDL_GetStringProperty
// #cgo nocallback SDL_GetStringProperty
// #cgo noescape SDL_GetNumberProperty
// #cgo nocallback SDL_GetNumberProperty
// #cgo noescape SDL_GetFloatProperty
// #cgo nocallback SDL_GetFloatProperty
// #cgo noescape SDL_GetBooleanProperty
// #cgo nocallback SDL_GetBooleanProperty
// #cgo noescape SDL_ClearProperty
// #cgo noescape wrap_SDL_EnumerateProperties
// #cgo noescape SDL_DestroyProperties
// #include <SDL3/SDL.h>
// extern bool wrap_SDL_SetPointerPropertyWithCleanup(SDL_PropertiesID props, const char *name, uintptr_t value, uintptr_t userdata);
// extern bool wrap_SDL_EnumerateProperties(SDL_PropertiesID props, uintptr_t userdata);
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

// # CategoryProperties
//
// A property is a variable that can be created and retrieved by name at
// runtime.
//
// All properties are part of a property group ([PropertiesID]). A property
// group can be created with the [CreateProperties] function and destroyed
// with the [PropertiesID.Destroy] function.
//
// Properties can be added to and retrieved from a property group through the
// following functions:
//
//   - [PropertiesID.SetPointer] and [PropertiesID.Pointer] operate on `void*`
//     pointer types.
//   - [PropertiesID.SetString] and [PropertiesID.String] operate on string types.
//   - [PropertiesID.SetNumber] and [PropertiesID.Number] operate on signed 64-bit
//     integer types.
//   - [PropertiesID.SetFloat] and [PropertiesID.Float] operate on floating point
//     types.
//   - [PropertiesID.SetBoolean] and [PropertiesID.Boolean] operate on boolean
//     types.
//
// Properties can be removed from a group by using [PropertiesID.Clear].

// SDL properties ID
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PropertiesID
type PropertiesID uint32

// SDL property type
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PropertyType
type PropertyType uint32

const (
	PropertyTypeInvalid PropertyType = iota
	PropertyTypePointer
	PropertyTypeString
	PropertyTypeNumber
	PropertyTypeFloat
	PropertyTypeBoolean
)

// Get the global SDL properties.
//
// Returns a valid property ID or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetGlobalProperties
func GetGlobalProperties() (PropertiesID, error) {
	props := (PropertiesID)(C.SDL_GetGlobalProperties())
	if props == 0 {
		return 0, getError()
	}
	return props, nil
}

// Create a group of properties.
//
// All properties are automatically destroyed when [Quit] is called.
//
// Returns an ID for a new group of properties or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateProperties
func CreateProperties() (PropertiesID, error) {
	props := (PropertiesID)(C.SDL_CreateProperties())
	if props == 0 {
		return 0, getError()
	}
	return props, nil
}

// Copy a group of properties.
//
// Copy all the properties from one group of properties to another, with the
// exception of properties requiring cleanup (set using
// [PropertiesID.SetPointerWithCleanup]), which will not be copied. Any
// property that already exists on dst will be overwritten.
//
// props: the properties to copy.
//
// dst: the destination properties.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CopyProperties
func (props PropertiesID) Copy(dst PropertiesID) error {
	if !C.SDL_CopyProperties((C.SDL_PropertiesID)(props), (C.SDL_PropertiesID)(dst)) {
		return getError()
	}
	return nil
}

// Lock a group of properties.
//
// Obtain a multi-threaded lock for these properties. Other threads will wait
// while trying to lock these properties until they are unlocked. Properties
// must be unlocked before they are destroyed.
//
// The lock is automatically taken when setting individual properties, this
// function is only needed when you want to set several properties atomically
// or want to guarantee that properties being queried aren't freed in another
// thread.
//
// props: the properties to lock.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_LockProperties
func (props PropertiesID) Lock() error {
	if !C.SDL_LockProperties((C.SDL_PropertiesID)(props)) {
		return getError()
	}
	return nil
}

// Unlock a group of properties.
//
// props: the properties to unlock.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_UnlockProperties
func (props PropertiesID) Unlock() {
	C.SDL_UnlockProperties((C.SDL_PropertiesID)(props))
}

// A callback used to free resources when a property is deleted.
//
// This should release any resources associated with value that are no
// longer needed.
//
// This callback is set per-property. Different properties in the same group
// can have different cleanup callbacks.
//
// This callback will be called _during_ [PropertiesID.SetPointerWithCleanup] if
// the function fails for any reason.
//
// value: the pointer assigned to the property to clean up.
//
// This callback may fire without any locks held; if this is a
// concern, the app should provide its own locking.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CleanupPropertyCallback
type CleanupPropertyCallback func(value uintptr)

//export cb_CleanupPropertyCallback
func cb_CleanupPropertyCallback(userdata uintptr, value uintptr) {
	h := cgo.Handle(userdata)
	h.Value().(CleanupPropertyCallback)(value)
	h.Delete()
}

// Set a pointer property in a group of properties with a cleanup function
// that is called when the property is deleted.
//
// The cleanup function is also called if setting the property fails for any
// reason.
//
// For simply setting basic data types, like numbers, bools, or strings, use
// [PropertiesID.SetNumber], [PropertiesID.SetBoolean], or
// [PropertiesID.SetString] instead, as those functions will handle cleanup on
// your behalf. This function is only for more complex, custom data.
//
// props: the properties to modify.
//
// name: the name of the property to modify.
//
// value: the new value of the property, or 0 to delete the property.
//
// cleanup: the function to call when this property is deleted.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetPointerPropertyWithCleanup
func (props PropertiesID) SetPointerWithCleanup(name string, value uintptr, cleanup CleanupPropertyCallback) error {
	h := cgo.NewHandle(cleanup)
	if !C.wrap_SDL_SetPointerPropertyWithCleanup((C.SDL_PropertiesID)(props), tmpstring(name), C.uintptr_t(value), C.uintptr_t(h)) {
		return getError()
	}
	return nil
}

func propHandleCleanup(value uintptr) {
	cgo.Handle(value).Delete()
}

// Set a pointer property in a group of properties.
//
// props: the properties to modify.
//
// name: the name of the property to modify.
//
// value: the new value of the property, or nil to delete the property.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetPointerProperty
func (props PropertiesID) SetPointer(name string, value unsafe.Pointer) error {
	if !C.SDL_SetPointerProperty((C.SDL_PropertiesID)(props), tmpstring(name), value) {
		return getError()
	}
	return nil
}

// Set a string property in a group of properties.
//
// This function makes a copy of the string; the caller does not have to
// preserve the data after this call completes.
//
// props: the properties to modify.
//
// name: the name of the property to modify.
//
// value: the new value of the property, or nil to delete the property.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetStringProperty
func (props PropertiesID) SetString(name, value string) error {
	if !C.SDL_SetStringProperty((C.SDL_PropertiesID)(props), tmpstring(name), tmpstring(value)) {
		return getError()
	}
	return nil
}

// Set an integer property in a group of properties.
//
// props: the properties to modify.
//
// name: the name of the property to modify.
//
// value: the new value of the property.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetNumberProperty
func (props PropertiesID) SetNumber(name string, value int64) error {
	if !C.SDL_SetNumberProperty((C.SDL_PropertiesID)(props), tmpstring(name), (C.Sint64)(value)) {
		return getError()
	}
	return nil
}

// Set a floating point property in a group of properties.
//
// props: the properties to modify.
//
// name: the name of the property to modify.
//
// value: the new value of the property.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetFloatProperty
func (props PropertiesID) SetFloat(name string, value float32) error {
	if !C.SDL_SetFloatProperty((C.SDL_PropertiesID)(props), tmpstring(name), (C.float)(value)) {
		return getError()
	}
	return nil
}

// Set a boolean property in a group of properties.
//
// props: the properties to modify.
//
// name: the name of the property to modify.
//
// value: the new value of the property.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetBooleanProperty
func (props PropertiesID) SetBoolean(name string, value bool) error {
	if !C.SDL_SetBooleanProperty((C.SDL_PropertiesID)(props), tmpstring(name), (C.bool)(value)) {
		return getError()
	}
	return nil
}

// Return whether a property exists in a group of properties.
//
// props: the properties to query.
//
// name: the name of the property to query.
//
// Returns true if the property exists, or false if it doesn't.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_HasProperty
func (props PropertiesID) Has(name string) bool {
	return (bool)(C.SDL_HasProperty((C.SDL_PropertiesID)(props), tmpstring(name)))
}

// Get the type of a property in a group of properties.
//
// props: the properties to query.
//
// name: the name of the property to query.
//
// Returns the type of the property, or [PropertyTypeInvalid] if it is
// not set.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetPropertyType
func (props PropertiesID) Type(name string) PropertyType {
	return (PropertyType)(C.SDL_GetPropertyType((C.SDL_PropertiesID)(props), tmpstring(name)))
}

// Get a pointer property from a group of properties.
//
// By convention, the names of properties that SDL exposes on objects will
// start with "SDL.", and properties that SDL uses internally will start with
// "SDL.internal.". These should be considered read-only and should not be
// modified by applications.
//
// props: the properties to query.
//
// name: the name of the property to query.
//
// defaultValue: the default value of the property.
//
// Returns the value of the property, or defaultValue if it is not set or
// not a pointer property.
//
// It is safe to call this function from any thread, although
// the data returned is not protected and could potentially be
// freed if you call [PropertiesID.SetPointer] or
// [PropertiesID.Clear] on these properties from another thread.
// If you need to avoid this, use [PropertiesID.Lock] and
// [PropertiesID.Unlock].
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetPointerProperty
func (props PropertiesID) Pointer(name string, defaultValue unsafe.Pointer) unsafe.Pointer {
	return C.SDL_GetPointerProperty((C.SDL_PropertiesID)(props), tmpstring(name), defaultValue)
}

// Get a string property from a group of properties.
//
// props: the properties to query.
//
// name: the name of the property to query.
//
// defaultValue: the default value of the property.
//
// Returns the value of the property, or defaultValue if it is not set or
// not a string property.
//
// It is safe to call this function from any thread, although
// the data returned is not protected and could potentially be
// freed if you call [PropertiesID.SetString] or
// [PropertiesID.Clear] on these properties from another thread.
// If you need to avoid this, use [PropertiesID.Lock] and
// [PropertiesID.Unlock].
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetStringProperty
func (props PropertiesID) String(name string, defaultValue string) string {
	return C.GoString(C.SDL_GetStringProperty((C.SDL_PropertiesID)(props), tmpstring(name), tmpstring(defaultValue)))
}

// Get a number property from a group of properties.
//
// You can use [PropertiesID.Type] to query whether the property exists and
// is a number property.
//
// props: the properties to query.
//
// name: the name of the property to query.
//
// defaultValue: the default value of the property.
//
// Returns the value of the property, or defaultValue if it is not set or
// not a number property.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetNumberProperty
func (props PropertiesID) Number(name string, defaultValue int64) int64 {
	return (int64)(C.SDL_GetNumberProperty((C.SDL_PropertiesID)(props), tmpstring(name), (C.Sint64)(defaultValue)))
}

// Get a floating point property from a group of properties.
//
// You can use [PropertiesID.Type] to query whether the property exists and
// is a floating point property.
//
// props: the properties to query.
//
// name: the name of the property to query.
//
// defaultValue: the default value of the property.
//
// Returns the value of the property, or defaultValue if it is not set or
// not a float property.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetFloatProperty
func (props PropertiesID) Float(name string, defaultValue float32) float32 {
	return (float32)(C.SDL_GetFloatProperty((C.SDL_PropertiesID)(props), tmpstring(name), (C.float)(defaultValue)))
}

// Get a boolean property from a group of properties.
//
// You can use [PropertiesID.Type] to query whether the property exists and
// is a boolean property.
//
// props: the properties to query.
//
// name: the name of the property to query.
//
// defaultValue: the default value of the property.
//
// Returns the value of the property, or defaultValue if it is not set or
// not a boolean property.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetBooleanProperty
func (props PropertiesID) Boolean(name string, defaultValue bool) bool {
	return (bool)(C.SDL_GetBooleanProperty((C.SDL_PropertiesID)(props), tmpstring(name), (C.bool)(defaultValue)))
}

// Clear a property from a group of properties.
//
// props: the properties to modify.
//
// name: the name of the property to clear.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ClearProperty
func (props PropertiesID) Clear(name string) error {
	if !C.SDL_ClearProperty((C.SDL_PropertiesID)(props), tmpstring(name)) {
		return getError()
	}
	return nil
}

// A callback used to enumerate all the properties in a group of properties.
//
// This callback is called from [PropertiesID.Enumerate], and is called once
// per property in the set.
//
// props: the [PropertiesID] that is being enumerated.
//
// name: the next property name in the enumeration.
//
// [PropertiesID.Enumerate] holds a lock on props during this
// callback.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_EnumeratePropertiesCallback
type EnumeratePropertiesCallback func(props PropertiesID, name string)

//export cb_EnumeratePropertiesCallback
func cb_EnumeratePropertiesCallback(userdata uintptr, props C.SDL_PropertiesID, name *C.char) {
	h := cgo.Handle(userdata)
	h.Value().(EnumeratePropertiesCallback)(PropertiesID(props), C.GoString(name))
}

// Enumerate the properties contained in a group of properties.
//
// The callback function is called for each property in the group of
// properties. The properties are locked during enumeration.
//
// props: the properties to query.
//
// callback: the function to call for each property.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_EnumerateProperties
func (props PropertiesID) Enumerate(callback EnumeratePropertiesCallback) error {
	h := cgo.NewHandle(callback)
	ok := C.wrap_SDL_EnumerateProperties((C.SDL_PropertiesID)(props), C.uintptr_t(h))
	h.Delete()
	if !ok {
		return getError()
	}
	return nil
}

// Destroy a group of properties.
//
// All properties are deleted and their cleanup functions will be called, if
// any.
//
// props: the properties to destroy.
//
// This function should not be called while these properties are
// locked or other threads might be setting or getting values
// from these properties.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DestroyProperties
func (props PropertiesID) Destroy() {
	C.SDL_DestroyProperties((C.SDL_PropertiesID)(props))
}
