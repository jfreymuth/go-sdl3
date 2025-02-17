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
// #include <SDL3/SDL.h>
import "C"

// # CategoryScancode
//
// Defines keyboard scancodes.
//
// Please refer to the Best Keyboard Practices document for details on what
// this information means and how best to use it.
//
// https://wiki.libsdl.org/SDL3/BestKeyboardPractices

// The SDL keyboard scancode representation.
//
// An SDL scancode is the physical representation of a key on the keyboard,
// independent of language and keyboard mapping.
//
// Values of this type are used to represent keyboard keys, among other places
// in the `scancode` field of the SDL_KeyboardEvent structure.
//
// The values in this enumeration are based on the USB usage page standard:
// https://usb.org/sites/default/files/hut1_5.pdf
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Scancode
type Scancode uint32

const (
	ScancodeUnknown            Scancode = 0
	ScancodeA                  Scancode = 4
	ScancodeB                  Scancode = 5
	ScancodeC                  Scancode = 6
	ScancodeD                  Scancode = 7
	ScancodeE                  Scancode = 8
	ScancodeF                  Scancode = 9
	ScancodeG                  Scancode = 10
	ScancodeH                  Scancode = 11
	ScancodeI                  Scancode = 12
	ScancodeJ                  Scancode = 13
	ScancodeK                  Scancode = 14
	ScancodeL                  Scancode = 15
	ScancodeM                  Scancode = 16
	ScancodeN                  Scancode = 17
	ScancodeO                  Scancode = 18
	ScancodeP                  Scancode = 19
	ScancodeQ                  Scancode = 20
	ScancodeR                  Scancode = 21
	ScancodeS                  Scancode = 22
	ScancodeT                  Scancode = 23
	ScancodeU                  Scancode = 24
	ScancodeV                  Scancode = 25
	ScancodeW                  Scancode = 26
	ScancodeX                  Scancode = 27
	ScancodeY                  Scancode = 28
	ScancodeZ                  Scancode = 29
	Scancode1                  Scancode = 30
	Scancode2                  Scancode = 31
	Scancode3                  Scancode = 32
	Scancode4                  Scancode = 33
	Scancode5                  Scancode = 34
	Scancode6                  Scancode = 35
	Scancode7                  Scancode = 36
	Scancode8                  Scancode = 37
	Scancode9                  Scancode = 38
	Scancode0                  Scancode = 39
	ScancodeReturn             Scancode = 40
	ScancodeEscape             Scancode = 41
	ScancodeBackspace          Scancode = 42
	ScancodeTab                Scancode = 43
	ScancodeSpace              Scancode = 44
	ScancodeMinus              Scancode = 45
	ScancodeEquals             Scancode = 46
	ScancodeLeftbracket        Scancode = 47
	ScancodeRightbracket       Scancode = 48
	ScancodeSemicolon          Scancode = 51
	ScancodeApostrophe         Scancode = 52
	ScancodeComma              Scancode = 54
	ScancodePeriod             Scancode = 55
	ScancodeSlash              Scancode = 56
	ScancodeCapslock           Scancode = 57
	ScancodeF1                 Scancode = 58
	ScancodeF2                 Scancode = 59
	ScancodeF3                 Scancode = 60
	ScancodeF4                 Scancode = 61
	ScancodeF5                 Scancode = 62
	ScancodeF6                 Scancode = 63
	ScancodeF7                 Scancode = 64
	ScancodeF8                 Scancode = 65
	ScancodeF9                 Scancode = 66
	ScancodeF10                Scancode = 67
	ScancodeF11                Scancode = 68
	ScancodeF12                Scancode = 69
	ScancodePrintscreen        Scancode = 70
	ScancodeScrolllock         Scancode = 71
	ScancodePause              Scancode = 72
	ScancodeHome               Scancode = 74
	ScancodePageup             Scancode = 75
	ScancodeDelete             Scancode = 76
	ScancodeEnd                Scancode = 77
	ScancodePagedown           Scancode = 78
	ScancodeRight              Scancode = 79
	ScancodeLeft               Scancode = 80
	ScancodeDown               Scancode = 81
	ScancodeUp                 Scancode = 82
	ScancodeKPDivide           Scancode = 84
	ScancodeKPMultiply         Scancode = 85
	ScancodeKPMinus            Scancode = 86
	ScancodeKPPlus             Scancode = 87
	ScancodeKPEnter            Scancode = 88
	ScancodeKP1                Scancode = 89
	ScancodeKP2                Scancode = 90
	ScancodeKP3                Scancode = 91
	ScancodeKP4                Scancode = 92
	ScancodeKP5                Scancode = 93
	ScancodeKP6                Scancode = 94
	ScancodeKP7                Scancode = 95
	ScancodeKP8                Scancode = 96
	ScancodeKP9                Scancode = 97
	ScancodeKP0                Scancode = 98
	ScancodeKPPeriod           Scancode = 99
	ScancodeApplication        Scancode = 101 // windows contextual menu, compose
	ScancodeKPEquals           Scancode = 103
	ScancodeF13                Scancode = 104
	ScancodeF14                Scancode = 105
	ScancodeF15                Scancode = 106
	ScancodeF16                Scancode = 107
	ScancodeF17                Scancode = 108
	ScancodeF18                Scancode = 109
	ScancodeF19                Scancode = 110
	ScancodeF20                Scancode = 111
	ScancodeF21                Scancode = 112
	ScancodeF22                Scancode = 113
	ScancodeF23                Scancode = 114
	ScancodeF24                Scancode = 115
	ScancodeExecute            Scancode = 116
	ScancodeHelp               Scancode = 117 // AL Integrated Help Center
	ScancodeMenu               Scancode = 118 // Menu (show menu)
	ScancodeSelect             Scancode = 119
	ScancodeStop               Scancode = 120 // AC Stop
	ScancodeAgain              Scancode = 121 // AC Redo/Repeat
	ScancodeUndo               Scancode = 122 // AC Undo
	ScancodeCut                Scancode = 123 // AC Cut
	ScancodeCopy               Scancode = 124 // AC Copy
	ScancodePaste              Scancode = 125 // AC Paste
	ScancodeFind               Scancode = 126 // AC Find
	ScancodeMute               Scancode = 127
	ScancodeVolumeup           Scancode = 128
	ScancodeVolumedown         Scancode = 129
	ScancodeKPComma            Scancode = 133
	ScancodeKPEqualsas400      Scancode = 134
	ScancodeInternational2     Scancode = 136
	ScancodeInternational3     Scancode = 137 // Yen
	ScancodeInternational4     Scancode = 138
	ScancodeInternational5     Scancode = 139
	ScancodeInternational6     Scancode = 140
	ScancodeInternational7     Scancode = 141
	ScancodeInternational8     Scancode = 142
	ScancodeInternational9     Scancode = 143
	ScancodeLang1              Scancode = 144 // Hangul/English toggle
	ScancodeLang2              Scancode = 145 // Hanja conversion
	ScancodeLang3              Scancode = 146 // Katakana
	ScancodeLang4              Scancode = 147 // Hiragana
	ScancodeLang5              Scancode = 148 // Zenkaku/Hankaku
	ScancodeLang6              Scancode = 149 // reserved
	ScancodeLang7              Scancode = 150 // reserved
	ScancodeLang8              Scancode = 151 // reserved
	ScancodeLang9              Scancode = 152 // reserved
	ScancodeAlterase           Scancode = 153 // Erase-Eaze
	ScancodeSysreq             Scancode = 154
	ScancodeCancel             Scancode = 155 // AC Cancel
	ScancodeClear              Scancode = 156
	ScancodePrior              Scancode = 157
	ScancodeReturn2            Scancode = 158
	ScancodeSeparator          Scancode = 159
	ScancodeOut                Scancode = 160
	ScancodeOper               Scancode = 161
	ScancodeClearagain         Scancode = 162
	ScancodeCrsel              Scancode = 163
	ScancodeExsel              Scancode = 164
	ScancodeKP00               Scancode = 176
	ScancodeKP000              Scancode = 177
	ScancodeThousandsseparator Scancode = 178
	ScancodeDecimalseparator   Scancode = 179
	ScancodeCurrencyunit       Scancode = 180
	ScancodeCurrencysubunit    Scancode = 181
	ScancodeKPLeftparen        Scancode = 182
	ScancodeKPRightparen       Scancode = 183
	ScancodeKPLeftbrace        Scancode = 184
	ScancodeKPRightbrace       Scancode = 185
	ScancodeKPTab              Scancode = 186
	ScancodeKPBackspace        Scancode = 187
	ScancodeKPA                Scancode = 188
	ScancodeKPB                Scancode = 189
	ScancodeKPC                Scancode = 190
	ScancodeKPD                Scancode = 191
	ScancodeKPE                Scancode = 192
	ScancodeKPF                Scancode = 193
	ScancodeKPXor              Scancode = 194
	ScancodeKPPower            Scancode = 195
	ScancodeKPPercent          Scancode = 196
	ScancodeKPLess             Scancode = 197
	ScancodeKPGreater          Scancode = 198
	ScancodeKPAmpersand        Scancode = 199
	ScancodeKPDblampersand     Scancode = 200
	ScancodeKPVerticalbar      Scancode = 201
	ScancodeKPDblverticalbar   Scancode = 202
	ScancodeKPColon            Scancode = 203
	ScancodeKPHash             Scancode = 204
	ScancodeKPSpace            Scancode = 205
	ScancodeKPAt               Scancode = 206
	ScancodeKPExclam           Scancode = 207
	ScancodeKPMemstore         Scancode = 208
	ScancodeKPMemrecall        Scancode = 209
	ScancodeKPMemclear         Scancode = 210
	ScancodeKPMemadd           Scancode = 211
	ScancodeKPMemsubtract      Scancode = 212
	ScancodeKPMemmultiply      Scancode = 213
	ScancodeKPMemdivide        Scancode = 214
	ScancodeKPPlusminus        Scancode = 215
	ScancodeKPClear            Scancode = 216
	ScancodeKPClearentry       Scancode = 217
	ScancodeKPBinary           Scancode = 218
	ScancodeKPOctal            Scancode = 219
	ScancodeKPDecimal          Scancode = 220
	ScancodeKPHexadecimal      Scancode = 221
	ScancodeLeftCtrl           Scancode = 224
	ScancodeLeftShift          Scancode = 225
	ScancodeLeftAlt            Scancode = 226 // alt, option
	ScancodeLeftGui            Scancode = 227 // windows, command (apple), meta
	ScancodeRightCtrl          Scancode = 228
	ScancodeRightShift         Scancode = 229
	ScancodeRightAlt           Scancode = 230 // alt gr, option
	ScancodeRightGui           Scancode = 231 // windows, command (apple), meta
	ScancodeSleep              Scancode = 258 // Sleep
	ScancodeWake               Scancode = 259 // Wake
	ScancodeChannelIncrement   Scancode = 260 // Channel Increment
	ScancodeChannelDecrement   Scancode = 261 // Channel Decrement
	ScancodeMediaPlay          Scancode = 262 // Play
	ScancodeMediaPause         Scancode = 263 // Pause
	ScancodeMediaRecord        Scancode = 264 // Record
	ScancodeMediaFastForward   Scancode = 265 // Fast Forward
	ScancodeMediaRewind        Scancode = 266 // Rewind
	ScancodeMediaNextTrack     Scancode = 267 // Next Track
	ScancodeMediaPreviousTrack Scancode = 268 // Previous Track
	ScancodeMediaStop          Scancode = 269 // Stop
	ScancodeMediaEject         Scancode = 270 // Eject
	ScancodeMediaPlayPause     Scancode = 271 // Play / Pause
	ScancodeACNew              Scancode = 273 // AC New
	ScancodeACOpen             Scancode = 274 // AC Open
	ScancodeACClose            Scancode = 275 // AC Close
	ScancodeACExit             Scancode = 276 // AC Exit
	ScancodeACSave             Scancode = 277 // AC Save
	ScancodeACPrint            Scancode = 278 // AC Print
	ScancodeACProperties       Scancode = 279 // AC Properties
	ScancodeACSearch           Scancode = 280 // AC Search
	ScancodeACHome             Scancode = 281 // AC Home
	ScancodeACBack             Scancode = 282 // AC Back
	ScancodeACForward          Scancode = 283 // AC Forward
	ScancodeACStop             Scancode = 284 // AC Stop
	ScancodeACRefresh          Scancode = 285 // AC Refresh
	ScancodeACBookmarks        Scancode = 286 // AC Bookmarks
	ScancodeCall               Scancode = 289 // Used for accepting phone calls.
	ScancodeEndcall            Scancode = 290 // Used for rejecting phone calls.
	ScancodeReserved           Scancode = 400 // 400-500 reserved for dynamic keycodes
	ScancodeCount                       = 512 // not a key, just marks the number of scancodes for array bounds
)
