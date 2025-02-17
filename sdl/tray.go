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
// #cgo noescape SDL_CreateTray
// #cgo nocallback SDL_CreateTray
// #cgo noescape SDL_SetTrayIcon
// #cgo nocallback SDL_SetTrayIcon
// #cgo noescape SDL_SetTrayTooltip
// #cgo nocallback SDL_SetTrayTooltip
// #cgo noescape SDL_CreateTrayMenu
// #cgo nocallback SDL_CreateTrayMenu
// #cgo noescape SDL_CreateTraySubmenu
// #cgo nocallback SDL_CreateTraySubmenu
// #cgo noescape SDL_GetTrayMenu
// #cgo nocallback SDL_GetTrayMenu
// #cgo noescape SDL_GetTraySubmenu
// #cgo nocallback SDL_GetTraySubmenu
// #cgo noescape SDL_GetTrayEntries
// #cgo nocallback SDL_GetTrayEntries
// #cgo noescape SDL_RemoveTrayEntry
// #cgo nocallback SDL_RemoveTrayEntry
// #cgo noescape SDL_InsertTrayEntryAt
// #cgo nocallback SDL_InsertTrayEntryAt
// #cgo noescape SDL_SetTrayEntryLabel
// #cgo nocallback SDL_SetTrayEntryLabel
// #cgo noescape SDL_GetTrayEntryLabel
// #cgo nocallback SDL_GetTrayEntryLabel
// #cgo noescape SDL_SetTrayEntryChecked
// #cgo nocallback SDL_SetTrayEntryChecked
// #cgo noescape SDL_GetTrayEntryChecked
// #cgo nocallback SDL_GetTrayEntryChecked
// #cgo noescape SDL_SetTrayEntryEnabled
// #cgo nocallback SDL_SetTrayEntryEnabled
// #cgo noescape SDL_GetTrayEntryEnabled
// #cgo nocallback SDL_GetTrayEntryEnabled
// #cgo noescape wrap_SDL_SetTrayEntryCallback
// #cgo nocallback wrap_SDL_SetTrayEntryCallback
// #cgo noescape SDL_ClickTrayEntry
// #cgo noescape SDL_DestroyTray
// #cgo noescape SDL_GetTrayEntryParent
// #cgo nocallback SDL_GetTrayEntryParent
// #cgo noescape SDL_GetTrayMenuParentEntry
// #cgo nocallback SDL_GetTrayMenuParentEntry
// #cgo noescape SDL_GetTrayMenuParentTray
// #cgo nocallback SDL_GetTrayMenuParentTray
// #cgo noescape SDL_UpdateTrays
// #include <SDL3/SDL.h>
// extern void wrap_SDL_SetTrayEntryCallback(SDL_TrayEntry *entry, uintptr_t userdata);
import "C"
import (
	"runtime/cgo"
	"unsafe"
)

// # CategoryTray
//
// SDL offers a way to add items to the "system tray" (more correctly called
// the "notification area" on Windows). On platforms that offer this concept,
// an SDL app can add a tray icon, submenus, checkboxes, and clickable
// entries, and register a callback that is fired when the user clicks on
// these pieces.

// An opaque handle representing a toplevel system tray object.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Tray
type Tray C.struct_SDL_Tray

// An opaque handle representing a menu/submenu on a system tray object.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_TrayMenu
type TrayMenu C.struct_SDL_TrayMenu

// An opaque handle representing an entry on a system tray object.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_TrayEntry
type TrayEntry C.struct_SDL_TrayEntry

// Flags that control the creation of system tray entries.
//
// Some of these flags are required; exactly one of them must be specified at
// the time a tray entry is created. Other flags are optional; zero or more of
// those can be OR'ed together with the required flag.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_TrayEntryFlags
type TrayEntryFlags uint32

const (
	TrayentryButton   TrayEntryFlags = 0x00000001 //Make the entry a simple button. Required.
	TrayentryCheckbox TrayEntryFlags = 0x00000002 //Make the entry a checkbox. Required.
	TrayentrySubmenu  TrayEntryFlags = 0x00000004 //Prepare the entry to have a submenu. Required
	TrayentryDisabled TrayEntryFlags = 0x80000000 //Make the entry disabled. Optional.
	TrayentryChecked  TrayEntryFlags = 0x40000000 //Make the entry checked. This is valid only for checkboxes. Optional.
)

// A callback that is invoked when a tray entry is selected.
//
// userdata: an optional pointer to pass extra data to the callback when
// it will be invoked.
//
// entry: the tray entry that was selected.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_TrayCallback
type TrayCallback func(entry *TrayEntry)

//export cb_TrayCallback
func cb_TrayCallback(userdata uintptr, entry *C.SDL_TrayEntry) {
	cgo.Handle(userdata).Value().(TrayCallback)((*TrayEntry)(entry))
}

// Create an icon to be placed in the operating system's tray, or equivalent.
//
// Many platforms advise not using a system tray unless persistence is a
// necessary feature. Avoid needlessly creating a tray icon, as the user may
// feel like it clutters their interface.
//
// Using tray icons require the video subsystem.
//
// icon: a surface to be used as icon. May be NULL.
//
// tooltip: a tooltip to be displayed when the mouse hovers the icon in
// UTF-8 encoding. Not supported on all platforms. May be NULL.
//
// Returns The newly created system tray icon.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateTray
func CreateTray(icon *Surface, tooltip string) (*Tray, error) {
	t := (*Tray)(C.SDL_CreateTray(icon.internal, tmpstring(tooltip)))
	if t == nil {
		return nil, getError()
	}
	return t, nil
}

// Updates the system tray icon's icon.
//
// tray: the tray icon to be updated.
//
// icon: the new icon. May be NULL.
//
// This function should be called on the thread that created the
// tray.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetTrayIcon
func (tray *Tray) SetIcon(icon *Surface) {
	C.SDL_SetTrayIcon((*C.SDL_Tray)(tray), icon.internal)
}

// Updates the system tray icon's tooltip.
//
// tray: the tray icon to be updated.
//
// tooltip: the new tooltip in UTF-8 encoding. May be NULL.
//
// This function should be called on the thread that created the
// tray.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetTrayTooltip
func (tray *Tray) SetTooltip(tooltip string) {
	C.SDL_SetTrayTooltip((*C.SDL_Tray)(tray), tmpstring(tooltip))
}

// Create a menu for a system tray.
//
// This should be called at most once per tray icon.
//
// This function does the same thing as SDL_CreateTraySubmenu(), except that
// it takes a SDL_Tray instead of a SDL_TrayEntry.
//
// A menu does not need to be destroyed; it will be destroyed with the tray.
//
// tray: the tray to bind the menu to.
//
// Returns the newly created menu.
//
// This function should be called on the thread that created the
// tray.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateTrayMenu
func (tray *Tray) CreateMenu() *TrayMenu {
	return (*TrayMenu)(C.SDL_CreateTrayMenu((*C.SDL_Tray)(tray)))
}

// Create a submenu for a system tray entry.
//
// This should be called at most once per tray entry.
//
// This function does the same thing as SDL_CreateTrayMenu, except that it
// takes a SDL_TrayEntry instead of a SDL_Tray.
//
// A menu does not need to be destroyed; it will be destroyed with the tray.
//
// entry: the tray entry to bind the menu to.
//
// Returns the newly created menu.
//
// This function should be called on the thread that created the
// tray.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateTraySubmenu
func (entry *TrayEntry) CreateSubmenu() *TrayMenu {
	return (*TrayMenu)(C.SDL_CreateTraySubmenu((*C.SDL_TrayEntry)(entry)))
}

// Gets a previously created tray menu.
//
// You should have called SDL_CreateTrayMenu() on the tray object. This
// function allows you to fetch it again later.
//
// This function does the same thing as SDL_GetTraySubmenu(), except that it
// takes a SDL_Tray instead of a SDL_TrayEntry.
//
// A menu does not need to be destroyed; it will be destroyed with the tray.
//
// tray: the tray entry to bind the menu to.
//
// Returns the newly created menu.
//
// This function should be called on the thread that created the
// tray.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTrayMenu
func (tray *Tray) Menu() *TrayMenu {
	return (*TrayMenu)(C.SDL_GetTrayMenu((*C.SDL_Tray)(tray)))
}

// Gets a previously created tray entry submenu.
//
// You should have called SDL_CreateTraySubmenu() on the entry object. This
// function allows you to fetch it again later.
//
// This function does the same thing as SDL_GetTrayMenu(), except that it
// takes a SDL_TrayEntry instead of a SDL_Tray.
//
// A menu does not need to be destroyed; it will be destroyed with the tray.
//
// entry: the tray entry to bind the menu to.
//
// Returns the newly created menu.
//
// This function should be called on the thread that created the
// tray.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTraySubmenu
func (entry *TrayEntry) Submenu() *TrayMenu {
	return (*TrayMenu)(C.SDL_GetTraySubmenu((*C.SDL_TrayEntry)(entry)))
}

// Returns a list of entries in the menu, in order.
//
// menu: The menu to get entries from.
//
// count: An optional pointer to obtain the number of entries in the
// menu.
//
// Returns a NULL-terminated list of entries within the given menu. The
// pointer becomes invalid when any function that inserts or deletes
// entries in the menu is called.
//
// This function should be called on the thread that created the
// tray.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTrayEntries
func (menu *TrayMenu) Entries() []*TrayEntry {
	var count C.int
	e := (**TrayEntry)(unsafe.Pointer(C.SDL_GetTrayEntries((*C.SDL_TrayMenu)(menu), &count)))
	return unsafe.Slice(e, count)
}

// Removes a tray entry.
//
// entry: The entry to be deleted.
//
// This function should be called on the thread that created the
// tray.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RemoveTrayEntry
func (entry *TrayEntry) Remove() {
	C.SDL_RemoveTrayEntry((*C.SDL_TrayEntry)(entry))
}

// Insert a tray entry at a given position.
//
// If label is NULL, the entry will be a separator. Many functions won't work
// for an entry that is a separator.
//
// An entry does not need to be destroyed; it will be destroyed with the tray.
//
// menu: the menu to append the entry to.
//
// pos: the desired position for the new entry. Entries at or following
// this place will be moved. If pos is -1, the entry is appended.
//
// label: the text to be displayed on the entry, in UTF-8 encoding, or
// NULL for a separator.
//
// flags: a combination of flags, some of which are mandatory.
//
// Returns the newly created entry, or NULL if pos is out of bounds.
//
// This function should be called on the thread that created the
// tray.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_InsertTrayEntryAt
func (menu *TrayMenu) InsertEntryAt(pos int, label string, flags TrayEntryFlags) *TrayEntry {
	return (*TrayEntry)(C.SDL_InsertTrayEntryAt((*C.SDL_TrayMenu)(menu), (C.int)(pos), tmpstring(label), (C.SDL_TrayEntryFlags)(flags)))
}

// Sets the label of an entry.
//
// An entry cannot change between a separator and an ordinary entry; that is,
// it is not possible to set a non-NULL label on an entry that has a NULL
// label (separators), or to set a NULL label to an entry that has a non-NULL
// label. The function will silently fail if that happens.
//
// entry: the entry to be updated.
//
// label: the new label for the entry in UTF-8 encoding.
//
// This function should be called on the thread that created the
// tray.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetTrayEntryLabel
func (entry *TrayEntry) SetLabel(label string) {
	C.SDL_SetTrayEntryLabel((*C.SDL_TrayEntry)(entry), tmpstring(label))
}

// Gets the label of an entry.
//
// If the returned value is NULL, the entry is a separator.
//
// entry: the entry to be read.
//
// Returns the label of the entry in UTF-8 encoding.
//
// This function should be called on the thread that created the
// tray.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTrayEntryLabel
func (entry *TrayEntry) Label() string {
	return C.GoString(C.SDL_GetTrayEntryLabel((*C.SDL_TrayEntry)(entry)))
}

// Sets whether or not an entry is checked.
//
// The entry must have been created with the SDL_TRAYENTRY_CHECKBOX flag.
//
// entry: the entry to be updated.
//
// checked: true if the entry should be checked; false otherwise.
//
// This function should be called on the thread that created the
// tray.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetTrayEntryChecked
func (entry *TrayEntry) SetChecked(checked bool) {
	C.SDL_SetTrayEntryChecked((*C.SDL_TrayEntry)(entry), (C.bool)(checked))
}

// Gets whether or not an entry is checked.
//
// The entry must have been created with the SDL_TRAYENTRY_CHECKBOX flag.
//
// entry: the entry to be read.
//
// Returns true if the entry is checked; false otherwise.
//
// This function should be called on the thread that created the
// tray.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTrayEntryChecked
func (entry *TrayEntry) Checked() bool {
	return (bool)(C.SDL_GetTrayEntryChecked((*C.SDL_TrayEntry)(entry)))
}

// Sets whether or not an entry is enabled.
//
// entry: the entry to be updated.
//
// enabled: true if the entry should be enabled; false otherwise.
//
// This function should be called on the thread that created the
// tray.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetTrayEntryEnabled
func (entry *TrayEntry) SetEnabled(enabled bool) {
	C.SDL_SetTrayEntryEnabled((*C.SDL_TrayEntry)(entry), (C.bool)(enabled))
}

// Gets whether or not an entry is enabled.
//
// entry: the entry to be read.
//
// Returns true if the entry is enabled; false otherwise.
//
// This function should be called on the thread that created the
// tray.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTrayEntryEnabled
func (entry *TrayEntry) Enabled() bool {
	return (bool)(C.SDL_GetTrayEntryEnabled((*C.SDL_TrayEntry)(entry)))
}

// Sets a callback to be invoked when the entry is selected.
//
// entry: the entry to be updated.
//
// callback: a callback to be invoked when the entry is selected.
//
// userdata: an optional pointer to pass extra data to the callback when
// it will be invoked.
//
// This function should be called on the thread that created the
// tray.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetTrayEntryCallback
func (entry *TrayEntry) SetCallback(callback TrayCallback) {
	C.wrap_SDL_SetTrayEntryCallback((*C.SDL_TrayEntry)(entry), C.uintptr_t(cgo.NewHandle(callback)))
}

// Simulate a click on a tray entry.
//
// entry: The entry to activate.
//
// This function should be called on the thread that created the
// tray.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ClickTrayEntry
func (entry *TrayEntry) Click() {
	C.SDL_ClickTrayEntry((*C.SDL_TrayEntry)(entry))
}

// Destroys a tray object.
//
// This also destroys all associated menus and entries.
//
// tray: the tray icon to be destroyed.
//
// This function should be called on the thread that created the
// tray.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DestroyTray
func (tray *Tray) Destroy() {
	C.SDL_DestroyTray((*C.SDL_Tray)(tray))
}

// Gets the menu containing a certain tray entry.
//
// entry: the entry for which to get the parent menu.
//
// Returns the parent menu.
//
// This function should be called on the thread that created the
// tray.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTrayEntryParent
func (entry *TrayEntry) Parent() *TrayMenu {
	return (*TrayMenu)(C.SDL_GetTrayEntryParent((*C.SDL_TrayEntry)(entry)))
}

// Gets the entry for which the menu is a submenu, if the current menu is a
// submenu.
//
// Either this function or SDL_GetTrayMenuParentTray() will return non-NULL
// for any given menu.
//
// menu: the menu for which to get the parent entry.
//
// Returns the parent entry, or NULL if this menu is not a submenu.
//
// This function should be called on the thread that created the
// tray.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTrayMenuParentEntry
func (menu *TrayMenu) ParentEntry() *TrayEntry {
	return (*TrayEntry)(C.SDL_GetTrayMenuParentEntry((*C.SDL_TrayMenu)(menu)))
}

// Gets the tray for which this menu is the first-level menu, if the current
// menu isn't a submenu.
//
// Either this function or SDL_GetTrayMenuParentEntry() will return non-NULL
// for any given menu.
//
// menu: the menu for which to get the parent enttrayry.
//
// Returns the parent tray, or NULL if this menu is a submenu.
//
// This function should be called on the thread that created the
// tray.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTrayMenuParentTray
func (menu *TrayMenu) ParentTray() *Tray {
	return (*Tray)(C.SDL_GetTrayMenuParentTray((*C.SDL_TrayMenu)(menu)))
}

// Update the trays.
//
// This is called automatically by the event loop and is only needed if you're
// using trays but aren't handling SDL events.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_UpdateTrays
func UpdateTrays() {
	C.SDL_UpdateTrays()
}
