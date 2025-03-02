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
// #cgo noescape SDL_GetPreferredLocales
// #cgo nocallback SDL_GetPreferredLocales
// #include <SDL3/SDL.h>
import "C"
import "unsafe"

// # CategoryLocale
//
// SDL locale services.
//
// This provides a way to get a list of preferred locales (language plus
// country) for the user. There is exactly one function:
// [GetPreferredLocales], which handles all the heavy lifting, and offers
// documentation on all the strange ways humans might have configured their
// language settings.

// A struct to provide locale data.
//
// Locale data is split into a spoken language, like English, and an optional
// country, like Canada. The language will be in ISO-639 format (so English
// would be "en"), and the country, if not empty, will be an ISO-3166 country
// code (so Canada would be "CA").
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Locale
type Locale struct {
	Language string // A language name, like "en" for English.
	Country  string // A country, like "US" for America. Can be "".
}

// Report the user's preferred locale.
//
// Returned language strings are in the format xx, where 'xx' is an ISO-639
// language specifier (such as "en" for English, "de" for German, etc).
// Country strings are in the format YY, where "YY" is an ISO-3166 country
// code (such as "US" for the United States, "CA" for Canada, etc). Country
// might be empty if there's no specific guidance on them (so you might get {
// "en", "US" } for American English, but { "en", "" } means "English
// language, generically"). Language strings are never empty.
//
// Please note that not all of these strings are 2 characters; some are three
// or more.
//
// The returned list of locales are in the order of the user's preference. For
// example, a German citizen that is fluent in US English and knows enough
// Japanese to navigate around Tokyo might have a list like: {"de", "en_US",
// "jp"}. Someone from England might prefer British English (where
// "color" is spelled "colour", etc), but will settle for anything like it:
// {"en_GB", "en"}.
//
// This function returns nil on error, including when the platform does not
// supply this information at all.
//
// This might be a "slow" call that has to query the operating system. It's
// best to ask for this once and save the results. However, this list can
// change, usually because the user has changed a system preference outside of
// your program; SDL will send an [EventLocaleChanged] event in this case,
// if possible, and you can call this function again to get an updated copy of
// preferred locales.
//
// Returns a slice of locale pointers, or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetPreferredLocales
func GetPreferredLocales() ([]Locale, error) {
	var count C.int
	l := C.SDL_GetPreferredLocales(&count)
	if l == nil {
		return nil, getError()
	}
	result := make([]Locale, count)
	for i, l := range unsafe.Slice(l, count) {
		result[i].Language = C.GoString(l.language)
		result[i].Country = C.GoString(l.country)
	}
	C.SDL_free(unsafe.Pointer(l))
	return result, nil
}
