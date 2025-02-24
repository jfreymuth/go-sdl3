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
// #cgo noescape SDL_GetBasePath
// #cgo nocallback SDL_GetBasePath
// #cgo noescape SDL_GetPrefPath
// #cgo nocallback SDL_GetPrefPath
// #cgo noescape SDL_GetUserFolder
// #cgo nocallback SDL_GetUserFolder
// #include <SDL3/SDL.h>
import "C"
import "unsafe"

// # CategoryFilesystem
//
// SDL offers an API for examining and manipulating the system's filesystem.
// This covers most things one would need to do with directories, except for
// actual file I/O (which is covered by [CategoryIOStream](CategoryIOStream)
// and [CategoryAsyncIO](CategoryAsyncIO) instead).
//
// There are functions to answer necessary path questions:
//
//   - Where is my app's data? SDL_GetBasePath().
//   - Where can I safely write files? SDL_GetPrefPath().
//   - Where are paths like Downloads, Desktop, Music? SDL_GetUserFolder().
//   - What is this thing at this location? SDL_GetPathInfo().
//   - What items live in this folder? SDL_EnumerateDirectory().
//   - What items live in this folder by wildcard? SDL_GlobDirectory().
//   - What is my current working directory? SDL_GetCurrentDirectory().
//
// SDL also offers functions to manipulate the directory tree: renaming,
// removing, copying files.

// Get the directory where the application was run from.
//
// SDL caches the result of this call internally, but the first call to this
// function is not necessarily fast, so plan accordingly.
//
// **macOS and iOS Specific Functionality**: If the application is in a ".app"
// bundle, this function returns the Resource directory (e.g.
// MyApp.app/Contents/Resources/). This behaviour can be overridden by adding
// a property to the Info.plist file. Adding a string key with the name
// SDL_FILESYSTEM_BASE_DIR_TYPE with a supported value will change the
// behaviour.
//
// Supported values for the SDL_FILESYSTEM_BASE_DIR_TYPE property (Given an
// application in /Applications/SDLApp/MyApp.app):
//
//   - `resource`: bundle resource directory (the default). For example:
//     `/Applications/SDLApp/MyApp.app/Contents/Resources`
//   - `bundle`: the Bundle directory. For example:
//     `/Applications/SDLApp/MyApp.app/`
//   - `parent`: the containing directory of the bundle. For example:
//     `/Applications/SDLApp/`
//
// **Nintendo 3DS Specific Functionality**: This function returns "romfs"
// directory of the application as it is uncommon to store resources outside
// the executable. As such it is not a writable directory.
//
// The returned path is guaranteed to end with a path separator ('\\' on
// Windows, '/' on most other platforms).
//
// Returns an absolute path in UTF-8 encoding to the application data
// directory. an error will be returned on error or when the platform
// doesn't implement this functionality.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetBasePath
func GetBasePath() (string, error) {
	path := C.SDL_GetBasePath()
	if path == nil {
		return "", getError()
	}
	return C.GoString(path), nil
}

// Get the user-and-app-specific path where files can be written.
//
// Get the "pref dir". This is meant to be where users can write personal
// files (preferences and save games, etc) that are specific to your
// application. This directory is unique per user, per application.
//
// This function will decide the appropriate location in the native
// filesystem, create the directory if necessary, and return a string of the
// absolute path to the directory in UTF-8 encoding.
//
// On Windows, the string might look like:
//
//	"C:\\Users\\bob\\AppData\\Roaming\\My Company\\My Program Name\\"
//
// On Linux, the string might look like:
//
//	"/home/bob/.local/share/My Program Name/"
//
// On macOS, the string might look like:
//
//	"/Users/bob/Library/Application Support/My Program Name/"
//
// You should assume the path returned by this function is the only safe place
// to write files (and that [GetBasePath], while it might be writable, or
// even the parent of the returned path, isn't where you should be writing
// things).
//
// Both the org and app strings may become part of a directory name, so please
// follow these rules:
//
//   - Try to use the same org string (_including case-sensitivity_) for all
//     your applications that use this function.
//   - Always use a unique app string for each one, and make sure it never
//     changes for an app once you've decided on it.
//   - Unicode characters are legal, as long as they are UTF-8 encoded, but...
//   - ...only use letters, numbers, and spaces. Avoid punctuation like "Game
//     Name 2: Bad Guy's Revenge!" ... "Game Name 2" is sufficient.
//
// The returned path is guaranteed to end with a path separator ('\\' on
// Windows, '/' on most other platforms).
//
// org: the name of your organization.
//
// app: the name of your application.
//
// Returns a string of the user directory in platform-dependent notation,
// or an error if there's a problem (creating directory failed, etc.).
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetPrefPath
func GetPrefPath(org string, app string) (string, error) {
	path := C.SDL_GetPrefPath(tmpstring(org), tmpstring(app))
	if path == nil {
		return "", getError()
	}
	result := C.GoString(path)
	C.SDL_free(unsafe.Pointer(path))
	return result, nil
}

// The type of the OS-provided default folder for a specific purpose.
//
// Note that the Trash folder isn't included here, because trashing files
// usually involves extra OS-specific functionality to remember the file's
// original location.
//
// The folders supported per platform are:
//
// |             | Windows | macOS/iOS | tvOS | Unix (XDG) | Haiku | Emscripten |
// | ----------- | ------- | --------- | ---- | ---------- | ----- | ---------- |
// | HOME        | X       | X         |      | X          | X     | X          |
// | DESKTOP     | X       | X         |      | X          | X     |            |
// | DOCUMENTS   | X       | X         |      | X          |       |            |
// | DOWNLOADS   | Vista+  | X         |      | X          |       |            |
// | MUSIC       | X       | X         |      | X          |       |            |
// | PICTURES    | X       | X         |      | X          |       |            |
// | PUBLICSHARE |         | X         |      | X          |       |            |
// | SAVEDGAMES  | Vista+  |           |      |            |       |            |
// | SCREENSHOTS | Vista+  |           |      |            |       |            |
// | TEMPLATES   | X       | X         |      | X          |       |            |
// | VIDEOS      | X       | X*        |      | X          |       |            |
//
// Note that on macOS/iOS, the Videos folder is called "Movies".
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Folder
type Folder uint32

const (
	FolderHome        Folder = iota // The folder which contains all of the current user's data, preferences, and documents. It usually contains most of the other folders. If a requested folder does not exist, the home folder can be considered a safe fallback to store a user's documents.
	FolderDesktop                   // The folder of files that are displayed on the desktop. Note that the existence of a desktop folder does not guarantee that the system does show icons on its desktop; certain GNU/Linux distros with a graphical environment may not have desktop icons.
	FolderDocuments                 // User document files, possibly application-specific. This is a good place to save a user's projects.
	FolderDownloads                 // Standard folder for user files downloaded from the internet.
	FolderMusic                     // Music files that can be played using a standard music player (mp3, ogg...).
	FolderPictures                  // Image files that can be displayed using a standard viewer (png, jpg...).
	FolderPublicshare               // Files that are meant to be shared with other users on the same computer.
	FolderSavedgames                // Save files for games.
	FolderScreenshots               // Application screenshots.
	FolderTemplates                 // Template files to be used when the user requests the desktop environment to create a new file in a certain folder, such as "New Text File.txt".  Any file in the Templates folder can be used as a starting point for a new file.
	FolderVideos                    // Video files that can be played using a standard video player (mp4, webm...).
	FolderCount       = iota        // Total number of types in this enum, not a folder type by itself.
)

// Finds the most suitable user folder for a specific purpose.
//
// Many OSes provide certain standard folders for certain purposes, such as
// storing pictures, music or videos for a certain user. This function gives
// the path for many of those special locations.
//
// This function is specifically for _user_ folders, which are meant for the
// user to access and manage. For application-specific folders, meant to hold
// data for the application to manage, see [GetBasePath] and
// [GetPrefPath].
//
// The returned path is guaranteed to end with a path separator ('\\' on
// Windows, '/' on most other platforms).
//
// folder: the type of folder to find.
//
// Returns either a string containing the full path to the
// folder, or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetUserFolder
func GetUserFolder(folder Folder) (string, error) {
	path := C.SDL_GetUserFolder((C.SDL_Folder)(folder))
	if path == nil {
		return "", getError()
	}
	return C.GoString(path), nil
}
