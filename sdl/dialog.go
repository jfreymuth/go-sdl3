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
// #cgo noescape wrap_SDL_ShowOpenFileDialog
// #cgo noescape wrap_SDL_ShowSaveFileDialog
// #cgo noescape wrap_SDL_ShowOpenFolderDialog
// #cgo noescape wrap_SDL_ShowFileDialogWithProperties
// #include <SDL3/SDL.h>
// extern void wrap_SDL_ShowOpenFileDialog(uintptr_t userdata, SDL_Window *window, const SDL_DialogFileFilter *filters, int nfilters, const char *default_location, bool allow_many);
// extern void wrap_SDL_ShowSaveFileDialog(uintptr_t userdata, SDL_Window *window, const SDL_DialogFileFilter *filters, int nfilters, const char *default_location);
// extern void wrap_SDL_ShowOpenFolderDialog(uintptr_t userdata, SDL_Window *window, const char *default_location, bool allow_many);
// extern void wrap_SDL_ShowFileDialogWithProperties(SDL_FileDialogType type, uintptr_t userdata, SDL_PropertiesID props);
import "C"
import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

// # CategoryDialog
//
// File dialog support.
//
// SDL offers file dialogs, to let users select files with native GUI
// interfaces. There are "open" dialogs, "save" dialogs, and folder selection
// dialogs. The app can control some details, such as filtering to specific
// files, or whether multiple files can be selected by the user.
//
// Note that launching a file dialog is a non-blocking operation; control
// returns to the app immediately, and a callback is called later (possibly in
// another thread) when the user makes a choice.

// An entry for filters for file dialogs.
//
// Name is a user-readable label for the filter (for example, "Office
// document").
//
// Pattern is a semicolon-separated list of file extensions (for example,
// "doc;docx"). File extensions may only contain alphanumeric characters,
// hyphens, underscores and periods. Alternatively, the whole string can be a
// single asterisk ("*"), which serves as an "All files" filter.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DialogFileFilter
type DialogFileFilter struct {
	Name    string
	Pattern string
}

// Callback used by file dialog functions.
//
// The specific usage is described in each function.
//
// The meaning of the arguments depends on which are nil:
//
//   - err != nil: an error occurred.
//   - err == nil && filelist == nil: the user either didn't choose any file or
//     canceled the dialog.
//   - err == nil && filelist != nil: the user chose one or more files. filelist
//     contains the selected paths.
//
// The filter argument is the index of the filter that was selected, or -1 if
// no filter was selected or if the platform or method doesn't support
// fetching the selected filter.
//
// In Android, the filelist are `content://` URIs. They should be opened
// using [IOFromFile] with appropriate modes. This applies both to open
// and save file dialog.
//
// filelist: the file(s) chosen by the user.
//
// filter: index of the selected filter.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DialogFileCallback
type DialogFileCallback func(filelist []string, filter int, err error)

//export cb_DialogFileCallback
func cb_DialogFileCallback(userdata uintptr, filelist **C.char, filter C.int) {
	h := cgo.Handle(userdata)
	var err error
	var files []string
	if filelist == nil {
		err = getError()
	} else {
		for *filelist != nil {
			files = append(files, C.GoString(*filelist))
			filelist = (**C.char)(unsafe.Add(unsafe.Pointer(filelist), unsafe.Sizeof(*filelist)))
		}
	}
	h.Value().(DialogFileCallback)(files, int(filter), err)
	h.Delete()
}

// Displays a dialog that lets the user select a file on their filesystem.
//
// This is an asynchronous function; it will return immediately, and the
// result will be passed to the callback.
//
// The callback will be invoked with a null-terminated list of files the user
// chose. The list will be empty if the user canceled the dialog.
//
// Note that the callback may be called from a different thread than the one
// the function was invoked on.
//
// Depending on the platform, the user may be allowed to input paths that
// don't yet exist.
//
// On Linux, dialogs may require XDG Portals, which requires DBus, which
// requires an event-handling loop. Apps that do not use SDL to handle events
// should add a call to [PumpEvents] in their main loop.
//
// callback: a function pointer to be invoked when the user selects a
// file and accepts, or cancels the dialog, or an error
// occurs.
//
// window: the window that the dialog should be modal for, may be nil.
// Not all platforms support this option.
//
// filters: a list of filters, may be nil. Not all platforms support
// this option, and platforms that do support it may allow the
// user to ignore the filters.
//
// defaultLocation: the default folder or file to start the dialog at,
// may be empty. Not all platforms support this option.
//
// allowMany: if true, the user will be allowed to select multiple
// entries. Not all platforms support this option.
//
// This function should be called only from the main thread. The
// callback may be invoked from the same thread or from a
// different one, depending on the OS's constraints.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ShowOpenFileDialog
func ShowOpenFileDialog(callback DialogFileCallback, window *Window, filters []DialogFileFilter, defaultLocation string, allowMany bool) {
	cfilters := make([]C.SDL_DialogFileFilter, len(filters))
	var pinner runtime.Pinner
	for i, f := range filters {
		cfilters[i].name = tmpstring(f.Name)
		cfilters[i].pattern = tmpstring(f.Pattern)
		pinner.Pin(cfilters[i].name)
		pinner.Pin(cfilters[i].pattern)
	}
	C.wrap_SDL_ShowOpenFileDialog(C.uintptr_t(cgo.NewHandle(callback)), (*C.SDL_Window)(window), unsafe.SliceData(cfilters), (C.int)(len(cfilters)), tmpstring(defaultLocation), (C.bool)(allowMany))
	pinner.Unpin()
}

// Displays a dialog that lets the user choose a new or existing file on their
// filesystem.
//
// This is an asynchronous function; it will return immediately, and the
// result will be passed to the callback.
//
// The callback will be invoked with a null-terminated list of files the user
// chose. The list will be empty if the user canceled the dialog.
//
// Note that the callback may be called from a different thread than the one
// the function was invoked on.
//
// The chosen file may or may not already exist.
//
// On Linux, dialogs may require XDG Portals, which requires DBus, which
// requires an event-handling loop. Apps that do not use SDL to handle events
// should add a call to [PumpEvents] in their main loop.
//
// callback: a function pointer to be invoked when the user selects a
// file and accepts, or cancels the dialog, or an error
// occurs.
//
// window: the window that the dialog should be modal for, may be nil.
// Not all platforms support this option.
//
// filters: a list of filters, may be nil. Not all platforms support
// this option, and platforms that do support it may allow the
// user to ignore the filters.
//
// defaultLocation: the default folder or file to start the dialog at,
// may be empty. Not all platforms support this option.
//
// This function should be called only from the main thread. The
// callback may be invoked from the same thread or from a
// different one, depending on the OS's constraints.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ShowSaveFileDialog
func ShowSaveFileDialog(callback DialogFileCallback, window *Window, filters []DialogFileFilter, defaultLocation string) {
	cfilters := make([]C.SDL_DialogFileFilter, len(filters))
	var pinner runtime.Pinner
	for i, f := range filters {
		cfilters[i].name = tmpstring(f.Name)
		cfilters[i].pattern = tmpstring(f.Pattern)
		pinner.Pin(cfilters[i].name)
		pinner.Pin(cfilters[i].pattern)
	}
	C.wrap_SDL_ShowSaveFileDialog(C.uintptr_t(cgo.NewHandle(callback)), (*C.SDL_Window)(window), unsafe.SliceData(cfilters), (C.int)(len(cfilters)), tmpstring(defaultLocation))
	pinner.Unpin()
}

// Displays a dialog that lets the user select a folder on their filesystem.
//
// This is an asynchronous function; it will return immediately, and the
// result will be passed to the callback.
//
// The callback will be invoked with a null-terminated list of files the user
// chose. The list will be empty if the user canceled the dialog.
//
// Note that the callback may be called from a different thread than the one
// the function was invoked on.
//
// Depending on the platform, the user may be allowed to input paths that
// don't yet exist.
//
// On Linux, dialogs may require XDG Portals, which requires DBus, which
// requires an event-handling loop. Apps that do not use SDL to handle events
// should add a call to [PumpEvents] in their main loop.
//
// callback: a function pointer to be invoked when the user selects a
// file and accepts, or cancels the dialog, or an error
// occurs.
//
// window: the window that the dialog should be modal for, may be nil.
// Not all platforms support this option.
//
// defaultLocation: the default folder or file to start the dialog at,
// may be empty. Not all platforms support this option.
//
// allow_many: if true, the user will be allowed to select multiple
// entries. Not all platforms support this option.
//
// This function should be called only from the main thread. The
// callback may be invoked from the same thread or from a
// different one, depending on the OS's constraints.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ShowOpenFolderDialog
func ShowOpenFolderDialog(callback DialogFileCallback, window *Window, defaultLocation string, allowMany bool) {
	C.wrap_SDL_ShowOpenFolderDialog(C.uintptr_t(cgo.NewHandle(callback)), (*C.SDL_Window)(window), tmpstring(defaultLocation), (C.bool)(allowMany))
}

// Various types of file dialogs.
//
// This is used by [ShowFileDialogWithProperties] to decide what kind of
// dialog to present to the user.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_FileDialogType
type FileDialogType uint32

const (
	FiledialogOpenfile FileDialogType = iota
	FiledialogSavefile
	FiledialogOpenfolder
)

// Create and launch a file dialog with the specified properties.
//
// These are the supported properties:
//
//   - [PropFileDialogFiltersPointer]: a pointer to a list of
//     [DialogFileFilter] structs, which will be used as filters for
//     file-based selections. Ignored if the dialog is an "Open Folder" dialog.
//     If non-NULL, the array of filters must remain valid at least until the
//     callback is invoked.
//   - [PropFileDialogNfiltersNumber]: the number of filters in the
//     array of filters, if it exists.
//   - [PropFileDialogWindowPointer]: the window that the dialog should
//     be modal for.
//   - [PropFileDialogLocationString]: the default folder or file to
//     start the dialog at.
//   - [PropFileDialogManyBoolean]: true to allow the user to select
//     more than one entry.
//   - [PropFileDialogTitleString]: the title for the dialog.
//   - [PropFileDialogAcceptString]: the label that the accept button
//     should have.
//   - [PropFileDialogCancelString]: the label that the cancel button
//     should have.
//
// Note that each platform may or may not support any of the properties.
//
// type: the type of file dialog.
//
// callback: a function pointer to be invoked when the user selects a
// file and accepts, or cancels the dialog, or an error
// occurs.
//
// props: the properties to use.
//
// This function should be called only from the main thread. The
// callback may be invoked from the same thread or from a
// different one, depending on the OS's constraints.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ShowFileDialogWithProperties
func ShowFileDialogWithProperties(typ FileDialogType, callback DialogFileCallback, props PropertiesID) {
	C.wrap_SDL_ShowFileDialogWithProperties((C.SDL_FileDialogType)(typ), C.uintptr_t(cgo.NewHandle(callback)), (C.SDL_PropertiesID)(props))
}

const PropFileDialogFiltersPointer = "SDL.filedialog.filters"
const PropFileDialogNfiltersNumber = "SDL.filedialog.nfilters"
const PropFileDialogWindowPointer = "SDL.filedialog.window"
const PropFileDialogLocationString = "SDL.filedialog.location"
const PropFileDialogManyBoolean = "SDL.filedialog.many"
const PropFileDialogTitleString = "SDL.filedialog.title"
const PropFileDialogAcceptString = "SDL.filedialog.accept"
const PropFileDialogCancelString = "SDL.filedialog.cancel"
