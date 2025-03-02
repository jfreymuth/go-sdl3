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
// #cgo noescape SDL_GetPixelFormatName
// #cgo nocallback SDL_GetPixelFormatName
// #cgo noescape SDL_GetMasksForPixelFormat
// #cgo nocallback SDL_GetMasksForPixelFormat
// #cgo noescape SDL_GetPixelFormatForMasks
// #cgo nocallback SDL_GetPixelFormatForMasks
// #cgo noescape SDL_GetPixelFormatDetails
// #cgo nocallback SDL_GetPixelFormatDetails
// #cgo noescape SDL_CreatePalette
// #cgo nocallback SDL_CreatePalette
// #cgo noescape SDL_SetPaletteColors
// #cgo nocallback SDL_SetPaletteColors
// #cgo noescape SDL_DestroyPalette
// #cgo nocallback SDL_DestroyPalette
// #cgo noescape SDL_MapRGB
// #cgo nocallback SDL_MapRGB
// #cgo noescape SDL_MapRGBA
// #cgo nocallback SDL_MapRGBA
// #cgo noescape SDL_GetRGB
// #cgo nocallback SDL_GetRGB
// #cgo noescape SDL_GetRGBA
// #cgo nocallback SDL_GetRGBA
// #include <SDL3/SDL.h>
import "C"
import "unsafe"

// # CategoryPixels
//
// SDL offers facilities for pixel management.
//
// Largely these facilities deal with pixel _format_: what does this set of
// bits represent?
//
// If you mostly want to think of a pixel as some combination of red, green,
// blue, and maybe alpha intensities, this is all pretty straightforward, and
// in many cases, is enough information to build a perfectly fine game.
//
// However, the actual definition of a pixel is more complex than that:
//
// Pixels are a representation of a color in a particular color space.
//
// The first characteristic of a color space is the color type. SDL
// understands two different color types, RGB and YCbCr, or in SDL also
// referred to as YUV.
//
// RGB colors consist of red, green, and blue channels of color that are added
// together to represent the colors we see on the screen.
//
// https://en.wikipedia.org/wiki/RGB_color_model
//
// YCbCr colors represent colors as a Y luma brightness component and red and
// blue chroma color offsets. This color representation takes advantage of the
// fact that the human eye is more sensitive to brightness than the color in
// an image. The Cb and Cr components are often compressed and have lower
// resolution than the luma component.
//
// https://en.wikipedia.org/wiki/YCbCr
//
// When the color information in YCbCr is compressed, the Y pixels are left at
// full resolution and each Cr and Cb pixel represents an average of the color
// information in a block of Y pixels. The chroma location determines where in
// that block of pixels the color information is coming from.
//
// The color range defines how much of the pixel to use when converting a
// pixel into a color on the display. When the full color range is used, the
// entire numeric range of the pixel bits is significant. When narrow color
// range is used, for historical reasons, the pixel uses only a portion of the
// numeric range to represent colors.
//
// The color primaries and white point are a definition of the colors in the
// color space relative to the standard XYZ color space.
//
// https://en.wikipedia.org/wiki/CIE_1931_color_space
//
// The transfer characteristic, or opto-electrical transfer function (OETF),
// is the way a color is converted from mathematically linear space into a
// non-linear output signals.
//
// https://en.wikipedia.org/wiki/Rec._709#Transfer_characteristics
//
// The matrix coefficients are used to convert between YCbCr and RGB colors.

// A fully opaque 8-bit alpha value.
//
// This macro is available since SDL 3.2.0.
const AlphaOpaque = 255

// A fully opaque floating point alpha value.
//
// This macro is available since SDL 3.2.0.
const AlphaOpaqueFloat = 1.0

// A fully transparent 8-bit alpha value.
//
// This macro is available since SDL 3.2.0.
const AlphaTransparent = 0

// A fully transparent floating point alpha value.
//
// This macro is available since SDL 3.2.0.
const AlphaTransparentFloat = 0.0

// Pixel type.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PixelType
type PixelType uint32

const (
	PixeltypeUnknown PixelType = iota
	PixeltypeIndex1
	PixeltypeIndex4
	PixeltypeIndex8
	PixeltypePacked8
	PixeltypePacked16
	PixeltypePacked32
	PixeltypeArrayU8
	PixeltypeArrayU16
	PixeltypeArrayU32
	PixeltypeArrayF16
	PixeltypeArrayF32
	PixeltypeIndex2
)

// Bitmap pixel order, high bit -> low bit.
//
// Packed component order, high bit -> low bit.
//
// Array component order, low byte -> high byte.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PixelOrder
type PixelOrder uint32

const (
	BitmaporderNone PixelOrder = iota
	Bitmaporder4321
	Bitmaporder1234
)

const (
	PackedorderNone PixelOrder = iota
	PackedorderXRGB
	PackedorderRGBX
	PackedorderARGB
	PackedorderRGBA
	PackedorderXBGR
	PackedorderBGRX
	PackedorderABGR
	PackedorderBGRA
)

const (
	ArrayorderNone PixelOrder = iota
	ArrayorderRGB
	ArrayorderRGBA
	ArrayorderARGB
	ArrayorderBGR
	ArrayorderBGRA
	ArrayorderABGR
)

// Packed component layout.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PackedLayout
type PackedLayout uint32

const (
	PackedlayoutNone PackedLayout = iota
	Packedlayout332
	Packedlayout4444
	Packedlayout1555
	Packedlayout5551
	Packedlayout565
	Packedlayout8888
	Packedlayout2101010
	Packedlayout1010102
)

// A macro for defining custom FourCC pixel formats.
//
// For example, defining [PixelformatYV12] looks like this:
//
//	DefinePixelFourCC('Y', 'V', '1', '2')
//
// A: the first character of the FourCC code.
//
// B: the second character of the FourCC code.
//
// C: the third character of the FourCC code.
//
// D: the fourth character of the FourCC code.
//
// Returns a format value in the style of [PixelFormat].
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DEFINE_PIXELFOURCC
func DefinePixelFourCC(a, b, c, d byte) PixelFormat {
	return PixelFormat(a) | PixelFormat(b)<<8 | PixelFormat(c)<<16 | PixelFormat(d)<<24
}

// A macro for defining custom non-FourCC pixel formats.
//
// For example, defining [PixelformatRGBA8888] looks like this:
//
//	DefinePixelFormat(PixeltypePacked32, PackedorderRGBA, Packedlayout8888, 32, 4)
//
// pixelType: the type of the new format, probably a [PixelType] value.
//
// order: the order of the new format, should match the format.
//
// layout: the layout of the new format, or zero if not applicable.
//
// bits: the number of bits per pixel of the new format.
//
// bytes: the number of bytes per pixel of the new format.
//
// Returns a format value in the style of [PixelFormat].
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DEFINE_PIXELFORMAT
func DefinePixelFormat(pixelType PixelType, order PixelOrder, layout PackedLayout, bits byte, bytes byte) PixelFormat {
	return 1<<28 | PixelFormat(pixelType)<<24 | PixelFormat(order)<<20 | PixelFormat(layout)<<16 | PixelFormat(bits)<<8 | PixelFormat(bytes)
}

// A macro to retrieve the type of a [PixelFormat].
//
// This is usually a value from the [PixelType] enumeration.
//
// f: a [PixelFormat] to check.
//
// Returns the type of f.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PIXELTYPE
func (f PixelFormat) Type() PixelType {
	return PixelType(f>>24) & 0xF
}

// A macro to retrieve the order of a [PixelFormat].
//
// This is usually one of the Bitmaporder*, Packedorder*, or
// Arrayorder* constants, depending on the format type.
//
// f: a [PixelFormat] to check.
//
// Returns the order of f.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PIXELORDER
func (f PixelFormat) Order() PixelOrder {
	return PixelOrder(f>>20) & 0xF
}

// A macro to retrieve the layout of a [PixelFormat].
//
// This is usually a value from the [PackedLayout] enumeration, or zero if a
// layout doesn't make sense for the format type.
//
// f: a [PixelFormat] to check.
//
// Returns the layout of f.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PIXELLAYOUT
func (f PixelFormat) Layout() PackedLayout {
	return PackedLayout(f>>16) & 0xF
}

// A macro to determine a [PixelFormat]'s bits per pixel.
//
// FourCC formats will report zero here, as it rarely makes sense to measure
// them per-pixel.
//
// f: a [PixelFormat] to check.
//
// Returns the bits-per-pixel of f.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BITSPERPIXEL
func (f PixelFormat) BitsPerPixel() byte {
	if f.FourCC() {
		return 0
	}
	return byte(f >> 8)
}

// A macro to determine a [PixelFormat]'s bytes per pixel.
//
// FourCC formats do their best here, but many of them don't have a meaningful
// measurement of bytes per pixel.
//
// f: a [PixelFormat] to check.
//
// Returns the bytes-per-pixel of f.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BYTESPERPIXEL
func (f PixelFormat) BytesPerPixel() byte {
	if f.FourCC() {
		if f == PixelformatYUY2 || f == PixelformatUYVY || f == PixelformatYVYU || f == PixelformatP010 {
			return 2
		}
		return 1
	}
	return byte(f)
}

// A macro to determine if a [PixelFormat] is an indexed format.
//
// f: a [PixelFormat] to check.
//
// Returns true if the format is indexed, false otherwise.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ISPIXELFORMAT_INDEXED
func (f PixelFormat) Indexed() bool {
	return !f.FourCC() && (f.Type() == PixeltypeIndex1 || f.Type() == PixeltypeIndex2 || f.Type() == PixeltypeIndex4 || f.Type() == PixeltypeIndex8)
}

// A macro to determine if a [PixelFormat] is a packed format.
//
// f: a [PixelFormat] to check.
//
// Returns true if the format is packed, false otherwise.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ISPIXELFORMAT_PACKED
func (f PixelFormat) Packed() bool {
	return !f.FourCC() && (f.Type() == PixeltypePacked8 || f.Type() == PixeltypePacked16 || f.Type() == PixeltypePacked32)
}

// A macro to determine if a [PixelFormat] is an array format.
//
// f: a [PixelFormat] to check.
//
// Returns true if the format is an array, false otherwise.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ISPIXELFORMAT_ARRAY
func (f PixelFormat) Array() bool {
	return !f.FourCC() && (f.Type() == PixeltypeArrayU8 || f.Type() == PixeltypeArrayU16 || f.Type() == PixeltypeArrayU32 || f.Type() == PixeltypeArrayF16 || f.Type() == PixeltypeArrayF32)
}

// A macro to determine if a [PixelFormat] is a 10-bit format.
//
// f: a [PixelFormat] to check.
//
// Returns true if the format is 10-bit, false otherwise.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ISPIXELFORMAT_10BIT
func (f PixelFormat) Is10bit() bool {
	return !f.FourCC() && f.Type() == PixeltypePacked32 && f.Layout() == Packedlayout2101010
}

// A macro to determine if a [PixelFormat] is a floating point format.
//
// f: a [PixelFormat] to check.
//
// Returns true if the format is 10-bit, false otherwise.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ISPIXELFORMAT_FLOAT
func (f PixelFormat) Float() bool {
	return !f.FourCC() && (f.Type() == PixeltypeArrayF16 || f.Type() == PixeltypeArrayF32)
}

// A macro to determine if a [PixelFormat] has an alpha channel.
//
// f: a [PixelFormat] to check.
//
// Returns true if the format has alpha, false otherwise.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ISPIXELFORMAT_ALPHA
func (f PixelFormat) Alpha() bool {
	if f.Packed() {
		return f.Order() == PackedorderARGB || f.Order() == PackedorderRGBA || f.Order() == PackedorderABGR || f.Order() == PackedorderBGRA
	}
	if f.Array() {
		return f.Order() == ArrayorderARGB || f.Order() == ArrayorderRGBA || f.Order() == ArrayorderABGR || f.Order() == ArrayorderBGRA
	}
	return false
}

// A macro to determine if a [PixelFormat] is a "FourCC" format.
//
// This covers custom and other unusual formats.
//
// f: a [PixelFormat] to check.
//
// Returns true if the format is a "FourCC" format, false otherwise.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ISPIXELFORMAT_FOURCC
func (f PixelFormat) FourCC() bool {
	return f>>28 != 1
}

// Pixel format.
//
// SDL's pixel formats have the following naming convention:
//
//   - Names with a list of components and a single bit count, such as RGB24 and
//     ABGR32, define a platform-independent encoding into bytes in the order
//     specified. For example, in RGB24 data, each pixel is encoded in 3 bytes
//     (red, green, blue) in that order, and in ABGR32 data, each pixel is
//     encoded in 4 bytes alpha, blue, green, red) in that order. Use these
//     names if the property of a format that is important to you is the order
//     of the bytes in memory or on disk.
//   - Names with a bit count per component, such as ARGB8888 and XRGB1555, are
//     "packed" into an appropriately-sized integer in the platform's native
//     endianness. For example, ARGB8888 is a sequence of 32-bit integers; in
//     each integer, the most significant bits are alpha, and the least
//     significant bits are blue. On a little-endian CPU such as x86, the least
//     significant bits of each integer are arranged first in memory, but on a
//     big-endian CPU such as s390x, the most significant bits are arranged
//     first. Use these names if the property of a format that is important to
//     you is the meaning of each bit position within a native-endianness
//     integer.
//   - In indexed formats such as Index4LSB, each pixel is represented by
//     encoding an index into the palette into the indicated number of bits,
//     with multiple pixels packed into each byte if appropriate. In LSB
//     formats, the first (leftmost) pixel is stored in the least-significant
//     bits of the byte; in MSB formats, it's stored in the most-significant
//     bits. Index8 does not need LSB/MSB variants, because each pixel exactly
//     fills one byte.
//
// The 32-bit byte-array encodings such as RGBA32 are aliases for the
// appropriate 8888 encoding for the current platform. For example, RGBA32 is
// an alias for ABGR8888 on little-endian CPUs like x86, or an alias for
// RGBA8888 on big-endian CPUs.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PixelFormat
type PixelFormat uint32

const (
	PixelformatUnknown      PixelFormat = 0
	PixelformatIndex1LSB    PixelFormat = 0x11100100
	PixelformatIndex1MSB    PixelFormat = 0x11200100
	PixelformatIndex2LSB    PixelFormat = 0x1c100200
	PixelformatIndex2MSB    PixelFormat = 0x1c200200
	PixelformatIndex4LSB    PixelFormat = 0x12100400
	PixelformatIndex4MSB    PixelFormat = 0x12200400
	PixelformatIndex8       PixelFormat = 0x13000801
	PixelformatRGB332       PixelFormat = 0x14110801
	PixelformatXRGB4444     PixelFormat = 0x15120c02
	PixelformatXBGR4444     PixelFormat = 0x15520c02
	PixelformatXRGB1555     PixelFormat = 0x15130f02
	PixelformatXBGR1555     PixelFormat = 0x15530f02
	PixelformatARGB4444     PixelFormat = 0x15321002
	PixelformatRGBA4444     PixelFormat = 0x15421002
	PixelformatABGR4444     PixelFormat = 0x15721002
	PixelformatBGRA4444     PixelFormat = 0x15821002
	PixelformatARGB1555     PixelFormat = 0x15331002
	PixelformatRGBA5551     PixelFormat = 0x15441002
	PixelformatABGR1555     PixelFormat = 0x15731002
	PixelformatBGRA5551     PixelFormat = 0x15841002
	PixelformatRGB565       PixelFormat = 0x15151002
	PixelformatBGR565       PixelFormat = 0x15551002
	PixelformatRGB24        PixelFormat = 0x17101803
	PixelformatBGR24        PixelFormat = 0x17401803
	PixelformatXRGB8888     PixelFormat = 0x16161804
	PixelformatRGBX8888     PixelFormat = 0x16261804
	PixelformatXBGR8888     PixelFormat = 0x16561804
	PixelformatBGRX8888     PixelFormat = 0x16661804
	PixelformatARGB8888     PixelFormat = 0x16362004
	PixelformatRGBA8888     PixelFormat = 0x16462004
	PixelformatABGR8888     PixelFormat = 0x16762004
	PixelformatBGRA8888     PixelFormat = 0x16862004
	PixelformatXRGB2101010  PixelFormat = 0x16172004
	PixelformatXBGR2101010  PixelFormat = 0x16572004
	PixelformatARGB2101010  PixelFormat = 0x16372004
	PixelformatABGR2101010  PixelFormat = 0x16772004
	PixelformatRGB48        PixelFormat = 0x18103006
	PixelformatBGR48        PixelFormat = 0x18403006
	PixelformatRGBA64       PixelFormat = 0x18204008
	PixelformatARGB64       PixelFormat = 0x18304008
	PixelformatBGRA64       PixelFormat = 0x18504008
	PixelformatABGR64       PixelFormat = 0x18604008
	PixelformatRGB48Float   PixelFormat = 0x1a103006
	PixelformatBGR48Float   PixelFormat = 0x1a403006
	PixelformatRGBA64Float  PixelFormat = 0x1a204008
	PixelformatARGB64Float  PixelFormat = 0x1a304008
	PixelformatBGRA64Float  PixelFormat = 0x1a504008
	PixelformatABGR64Float  PixelFormat = 0x1a604008
	PixelformatRGB96Float   PixelFormat = 0x1b10600c
	PixelformatBGR96Float   PixelFormat = 0x1b40600c
	PixelformatRGBA128Float PixelFormat = 0x1b208010
	PixelformatARGB128Float PixelFormat = 0x1b308010
	PixelformatBGRA128Float PixelFormat = 0x1b508010
	PixelformatABGR128Float PixelFormat = 0x1b608010
	PixelformatYV12         PixelFormat = 0x32315659 // Planar mode: Y + V + U  (3 planes)
	PixelformatIYUV         PixelFormat = 0x56555949 // Planar mode: Y + U + V  (3 planes)
	PixelformatYUY2         PixelFormat = 0x32595559 // Packed mode: Y0+U0+Y1+V0 (1 plane)
	PixelformatUYVY         PixelFormat = 0x59565955 // Packed mode: U0+Y0+V0+Y1 (1 plane)
	PixelformatYVYU         PixelFormat = 0x55595659 // Packed mode: Y0+V0+Y1+U0 (1 plane)
	PixelformatNV12         PixelFormat = 0x3231564e // Planar mode: Y + U/V interleaved  (2 planes)
	PixelformatNV21         PixelFormat = 0x3132564e // Planar mode: Y + V/U interleaved  (2 planes)
	PixelformatP010         PixelFormat = 0x30313050 // Planar mode: Y + U/V interleaved  (2 planes)
	PixelformatExternalOES  PixelFormat = 0x2053454f // Android video texture format
)

// Colorspace color type.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ColorType
type ColorType uint32

const (
	ColorTypeUnknown ColorType = 0
	ColorTypeRGB     ColorType = 1
	ColorTypeYCbCr   ColorType = 2
)

// Colorspace color range, as described by
// https://www.itu.int/rec/R-REC-BT.2100-2-201807-I/en
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ColorRange
type ColorRange uint32

const (
	ColorRangeUnknown ColorRange = 0
	ColorRangeLimited ColorRange = 1 // Narrow range, e.g. 16-235 for 8-bit RGB and luma, and 16-240 for 8-bit chroma
	ColorRangeFull    ColorRange = 2 // Full range, e.g. 0-255 for 8-bit RGB and luma, and 1-255 for 8-bit chroma
)

// Colorspace color primaries, as described by
// https://www.itu.int/rec/T-REC-H.273-201612-S/en
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ColorPrimaries
type ColorPrimaries uint32

const (
	ColorPrimariesUnknown     ColorPrimaries = 0
	ColorPrimariesBT709       ColorPrimaries = 1 // ITU-R BT.709-6
	ColorPrimariesUnspecified ColorPrimaries = 2
	ColorPrimariesBT470M      ColorPrimaries = 4  // ITU-R BT.470-6 System M
	ColorPrimariesBT470BG     ColorPrimaries = 5  // ITU-R BT.470-6 System B, G / ITU-R BT.601-7 625
	ColorPrimariesBT601       ColorPrimaries = 6  // ITU-R BT.601-7 525, SMPTE 170M
	ColorPrimariesSMPTE240    ColorPrimaries = 7  // SMPTE 240M, functionally the same as [ColorPrimariesBT601]
	ColorPrimariesGenericFilm ColorPrimaries = 8  // Generic film (color filters using Illuminant C)
	ColorPrimariesBT2020      ColorPrimaries = 9  // ITU-R BT.2020-2 / ITU-R BT.2100-0
	ColorPrimariesXYZ         ColorPrimaries = 10 // SMPTE ST 428-1
	ColorPrimariesSMPTE431    ColorPrimaries = 11 // SMPTE RP 431-2
	ColorPrimariesSMPTE432    ColorPrimaries = 12 // SMPTE EG 432-1 / DCI P3
	ColorPrimariesEBU3213     ColorPrimaries = 22 // EBU Tech. 3213-E
	ColorPrimariesCustom      ColorPrimaries = 31
)

// Colorspace transfer characteristics.
//
// These are as described by https://www.itu.int/rec/T-REC-H.273-201612-S/en
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_TransferCharacteristics
type TransferCharacteristics uint32

const (
	TransferCharacteristicsUnknown      TransferCharacteristics = 0
	TransferCharacteristicsBT709        TransferCharacteristics = 1 // Rec. ITU-R BT.709-6 / ITU-R BT1361
	TransferCharacteristicsUnspecified  TransferCharacteristics = 2
	TransferCharacteristicsGAMMA22      TransferCharacteristics = 4 // ITU-R BT.470-6 System M / ITU-R BT1700 625 PAL & SECAM
	TransferCharacteristicsGAMMA28      TransferCharacteristics = 5 // ITU-R BT.470-6 System B, G
	TransferCharacteristicsBT601        TransferCharacteristics = 6 // SMPTE ST 170M / ITU-R BT.601-7 525 or 625
	TransferCharacteristicsSMPTE240     TransferCharacteristics = 7 // SMPTE ST 240M
	TransferCharacteristicsLinear       TransferCharacteristics = 8
	TransferCharacteristicsLog100       TransferCharacteristics = 9
	TransferCharacteristicsLog100Sqrt10 TransferCharacteristics = 10
	TransferCharacteristicsIEC61966     TransferCharacteristics = 11 // IEC 61966-2-4
	TransferCharacteristicsBT1361       TransferCharacteristics = 12 // ITU-R BT1361 Extended Colour Gamut
	TransferCharacteristicsSRGB         TransferCharacteristics = 13 // IEC 61966-2-1 (sRGB or sYCC)
	TransferCharacteristicsBT202010bit  TransferCharacteristics = 14 // ITU-R BT2020 for 10-bit system
	TransferCharacteristicsBT202012bit  TransferCharacteristics = 15 // ITU-R BT2020 for 12-bit system
	TransferCharacteristicsPQ           TransferCharacteristics = 16 // SMPTE ST 2084 for 10-, 12-, 14- and 16-bit systems
	TransferCharacteristicsSMPTE428     TransferCharacteristics = 17 // SMPTE ST 428-1
	TransferCharacteristicsHLG          TransferCharacteristics = 18 // ARIB STD-B67, known as "hybrid log-gamma" (HLG)
	TransferCharacteristicsCustom       TransferCharacteristics = 31
)

// Colorspace matrix coefficients.
//
// These are as described by https://www.itu.int/rec/T-REC-H.273-201612-S/en
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MatrixCoefficients
type MatrixCoefficients uint32

const (
	MatrixCoefficientsIdentity         MatrixCoefficients = 0
	MatrixCoefficientsBT709            MatrixCoefficients = 1 // ITU-R BT.709-6
	MatrixCoefficientsUnspecified      MatrixCoefficients = 2
	MatrixCoefficientsFCC              MatrixCoefficients = 4 // US FCC Title 47
	MatrixCoefficientsBT470BG          MatrixCoefficients = 5 // ITU-R BT.470-6 System B, G / ITU-R BT.601-7 625, functionally the same as [MatrixCoefficientsBT601]
	MatrixCoefficientsBT601            MatrixCoefficients = 6 // ITU-R BT.601-7 525
	MatrixCoefficientsSMPTE240         MatrixCoefficients = 7 // SMPTE 240M
	MatrixCoefficientsYCgCo            MatrixCoefficients = 8
	MatrixCoefficientsBT2020NCL        MatrixCoefficients = 9  // ITU-R BT.2020-2 non-constant luminance
	MatrixCoefficientsBT2020CL         MatrixCoefficients = 10 // ITU-R BT.2020-2 constant luminance
	MatrixCoefficientsSMPTE2085        MatrixCoefficients = 11 // SMPTE ST 2085
	MatrixCoefficientsChromaDerivedNCL MatrixCoefficients = 12
	MatrixCoefficientsChromaDerivedCL  MatrixCoefficients = 13
	MatrixCoefficientsICTCP            MatrixCoefficients = 14 // ITU-R BT.2100-0 ICTCP
	MatrixCoefficientsCustom           MatrixCoefficients = 31
)

// Colorspace chroma sample location.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ChromaLocation
type ChromaLocation uint32

const (
	ChromaLocationNone    ChromaLocation = 0 // RGB, no chroma sampling
	ChromaLocationLeft    ChromaLocation = 1 // In MPEG-2, MPEG-4, and AVC, Cb and Cr are taken on midpoint of the left-edge of the 2x2 square. In other words, they have the same horizontal location as the top-left pixel, but is shifted one-half pixel down vertically.
	ChromaLocationCenter  ChromaLocation = 2 // In JPEG/JFIF, H.261, and MPEG-1, Cb and Cr are taken at the center of the 2x2 square. In other words, they are offset one-half pixel to the right and one-half pixel down compared to the top-left pixel.
	ChromaLocationTopleft ChromaLocation = 3 // In HEVC for BT.2020 and BT.2100 content (in particular on Blu-rays), Cb and Cr are sampled at the same location as the group's top-left Y pixel ("co-sited", "co-located").
)

// A macro for defining custom [Colorspace] formats.
//
// For example, defining ColorspaceSRGB looks like this:
//
//	SDL_DEFINE_COLORSPACE(ColorTypeRGB,
//		ColorRangeFull,
//		ColorPrimariesBT709,
//		TransferCharacteristicsSRGB,
//		MatrixCoefficientsIdentity,
//		ChromaLocationNone)
//
// colorType: the type of the new format, probably a [ColorType] value.
//
// range: the range of the new format, probably a [ColorRange] value.
//
// primaries: the primaries of the new format, probably a
// [ColorPrimaries] value.
//
// transfer: the transfer characteristics of the new format, probably a
// [TransferCharacteristics] value.
//
// matrix: the matrix coefficients of the new format, probably a
// [MatrixCoefficients] value.
//
// chroma: the chroma sample location of the new format, probably a
// [ChromaLocation] value.
//
// Returns a format value in the style of [Colorspace].
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DEFINE_COLORSPACE
func DefineColorspace(colorType ColorType, colorRange ColorRange, primaries ColorPrimaries, transfer TransferCharacteristics, matrix MatrixCoefficients, chroma ChromaLocation) Colorspace {
	return ((Colorspace)(colorType) << 28) | ((Colorspace)(colorRange) << 24) | ((Colorspace)(chroma) << 20) |
		((Colorspace)(primaries) << 10) | ((Colorspace)(transfer) << 5) | ((Colorspace)(matrix) << 0)
}

// A macro to retrieve the type of a [Colorspace].
//
// c: a [Colorspace] to check.
//
// Returns the [ColorType] for c.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_COLORSPACETYPE
func (c Colorspace) Type() ColorType {
	return (ColorType)(((c) >> 28) & 0x0F)
}

// A macro to retrieve the range of a [Colorspace].
//
// c: a [Colorspace] to check.
//
// Returns the [ColorRange] of c.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_COLORSPACERANGE
func (c Colorspace) Range() ColorRange {
	return (ColorRange)(((c) >> 24) & 0x0F)
}

// A macro to retrieve the chroma sample location of a [Colorspace].
//
// c: a [Colorspace] to check.
//
// Returns the [ChromaLocation] of c.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_COLORSPACECHROMA
func (c Colorspace) Chroma() ChromaLocation {
	return (ChromaLocation)(((c) >> 20) & 0x0F)
}

// A macro to retrieve the primaries of a [Colorspace].
//
// c: a [Colorspace] to check.
//
// Returns the [ColorPrimaries] of c.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_COLORSPACEPRIMARIES
func (c Colorspace) Primaries() ColorPrimaries {
	return (ColorPrimaries)(((c) >> 10) & 0x1F)
}

// A macro to retrieve the transfer characteristics of a [Colorspace].
//
// c: a [Colorspace] to check.
//
// Returns the [TransferCharacteristics] of c.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_COLORSPACETRANSFER
func (c Colorspace) Transfer() TransferCharacteristics {
	return (TransferCharacteristics)(((c) >> 5) & 0x1F)
}

// A macro to retrieve the matrix coefficients of a [Colorspace].
//
// c: a [Colorspace] to check.
//
// Returns the [MatrixCoefficients] of c.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_COLORSPACEMATRIX
func (c Colorspace) Matrix() MatrixCoefficients {
	return (MatrixCoefficients)((c) & 0x1F)
}

// A macro to determine if a [Colorspace] uses BT601 (or BT470BG) matrix
// coefficients.
//
// c: a [Colorspace] to check.
//
// Returns true if BT601 or BT470BG, false otherwise.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ISCOLORSPACE_MATRIX_BT601
func (c Colorspace) MatrixBT601() bool {
	return c.Matrix() == MatrixCoefficientsBT601 || c.Matrix() == MatrixCoefficientsBT470BG
}

// A macro to determine if a [Colorspace] uses BT709 matrix coefficients.
//
// c: a [Colorspace] to check.
//
// Returns true if BT709, false otherwise.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ISCOLORSPACE_MATRIX_BT709
func (c Colorspace) MatrixBT709() bool {
	return c.Matrix() == MatrixCoefficientsBT709
}

// A macro to determine if a [Colorspace] uses BT2020_NCL matrix
// coefficients.
//
// c: a [Colorspace] to check.
//
// Returns true if BT2020_NCL, false otherwise.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ISCOLORSPACE_MATRIX_BT2020_NCL
func (c Colorspace) MatrixBT2020NCL() bool {
	return c.Matrix() == MatrixCoefficientsBT2020NCL
}

// A macro to determine if a [Colorspace] has a limited range.
//
// c: a [Colorspace] to check.
//
// Returns true if limited range, false otherwise.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ISCOLORSPACE_LIMITED_RANGE
func (c Colorspace) LimitedRange() bool {
	return c.Range() != ColorRangeFull
}

// A macro to determine if a [Colorspace] has a full range.
//
// c: a [Colorspace] to check.
//
// Returns true if full range, false otherwise.
//
// It is safe to call this macro from any thread.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ISCOLORSPACE_FULL_RANGE
func (c Colorspace) FullRange() bool {
	return c.Range() == ColorRangeFull
}

// Colorspace definitions.
//
// Since similar colorspaces may vary in their details (matrix, transfer
// function, etc.), this is not an exhaustive list, but rather a
// representative sample of the kinds of colorspaces supported in SDL.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Colorspace
type Colorspace uint32

const (
	ColorspaceUnknown       Colorspace = 0
	ColorspaceSRGB          Colorspace = 0x120005a0     // Equivalent to DXGI_COLOR_SPACE_RGB_FULL_G22_NONE_P709
	ColorspaceSRGBLinear    Colorspace = 0x12000500     // Equivalent to DXGI_COLOR_SPACE_RGB_FULL_G10_NONE_P709
	ColorspaceHDR10         Colorspace = 0x12002600     // Equivalent to DXGI_COLOR_SPACE_RGB_FULL_G2084_NONE_P2020
	ColorspaceJPEG          Colorspace = 0x220004c6     // Equivalent to DXGI_COLOR_SPACE_YCBCR_FULL_G22_NONE_P709_X601
	ColorspaceBT601Limited  Colorspace = 0x211018c6     // Equivalent to DXGI_COLOR_SPACE_YCBCR_STUDIO_G22_LEFT_P601
	ColorspaceBT601Full     Colorspace = 0x221018c6     // Equivalent to DXGI_COLOR_SPACE_YCBCR_STUDIO_G22_LEFT_P601
	ColorspaceBT709Limited  Colorspace = 0x21100421     // Equivalent to DXGI_COLOR_SPACE_YCBCR_STUDIO_G22_LEFT_P709
	ColorspaceBT709Full     Colorspace = 0x22100421     // Equivalent to DXGI_COLOR_SPACE_YCBCR_STUDIO_G22_LEFT_P709
	ColorspaceBT2020Limited Colorspace = 0x21102609     // Equivalent to DXGI_COLOR_SPACE_YCBCR_STUDIO_G22_LEFT_P2020
	ColorspaceBT2020Full    Colorspace = 0x22102609     // Equivalent to DXGI_COLOR_SPACE_YCBCR_FULL_G22_LEFT_P2020
	ColorspaceRGBDefault    Colorspace = ColorspaceSRGB // The default colorspace for RGB surfaces if no colorspace is specified
	ColorspaceYUVDefault    Colorspace = ColorspaceJPEG // The default colorspace for YUV surfaces if no colorspace is specified
)

// A structure that represents a color as RGBA components.
//
// The bits of this structure can be directly reinterpreted as an
// integer-packed color which uses the [PixelformatRGBA32] format
// ([PixelformatABGR8888] on little-endian systems and
// [PixelformatRGBA8888] on big-endian systems).
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Color
type Color struct {
	R byte
	G byte
	B byte
	A byte
}

// The bits of this structure can be directly reinterpreted as a float-packed
// color which uses the [PixelformatRGBA128Float] format
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_FColor
type FColor struct {
	R float32
	G float32
	B float32
	A float32
}

// A set of indexed colors representing a palette.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Palette
type Palette struct {
	internal *C.SDL_Palette
}

// an array of colors
func (palette *Palette) Colors() []Color {
	colors := make([]Color, palette.internal.ncolors)
	for i, c := range unsafe.Slice(palette.internal.colors, len(colors)) {
		colors[i] = Color{byte(c.r), byte(c.g), byte(c.b), byte(c.a)}
	}
	return colors
}

// Details about the format of a pixel.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PixelFormatDetails
type PixelFormatDetails struct {
	Format        PixelFormat
	BitsPerPixel  byte
	BytesPerPixel byte
	RMask         uint32
	GMask         uint32
	BMask         uint32
	AMask         uint32
	RBits         byte
	GBits         byte
	BBits         byte
	ABits         byte
	RShift        byte
	GShift        byte
	BShift        byte
	AShift        byte
}

// Get the human readable name of a pixel format.
//
// f: the pixel format to query.
//
// Returns the human readable name of the specified pixel format or
// "SDL_PIXELFORMAT_UNKNOWN" if the format isn't recognized.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatName
func (format PixelFormat) Name() string {
	return C.GoString(C.SDL_GetPixelFormatName((C.SDL_PixelFormat)(format)))
}

// Convert one of the enumerated pixel formats to a bpp value and RGBA masks.
//
// f: one of the [PixelFormat] values.
//
// bpp: a bits per pixel value; usually 15, 16, or 32.
//
// rMask: the red mask for the format.
//
// gMask: the green mask for the format.
//
// bMask: the blue mask for the format.
//
// aMask: the alpha mask for the format.
//
// Returns an error on failure.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetMasksForPixelFormat
func (format PixelFormat) Masks() (bpp int, rMask, gMask, bMask, aMask uint32, err error) {
	var cbpp C.int
	var Rmask, Gmask, Bmask, Amask C.Uint32
	if !C.SDL_GetMasksForPixelFormat((C.SDL_PixelFormat)(format), &cbpp, &Rmask, &Gmask, &Bmask, &Amask) {
		return 0, 0, 0, 0, 0, getError()
	}
	return int(cbpp), uint32(Rmask), uint32(Gmask), uint32(Bmask), uint32(Amask), nil
}

// Convert a bpp value and RGBA masks to an enumerated pixel format.
//
// This will return [PixelformatUnknown] if the conversion wasn't
// possible.
//
// bpp: a bits per pixel value; usually 15, 16, or 32.
//
// rMask: the red mask for the format.
//
// gMask: the green mask for the format.
//
// bMask: the blue mask for the format.
//
// aMask: the alpha mask for the format.
//
// Returns the [PixelFormat] value corresponding to the format masks, or
// [PixelformatUnknown] if there isn't a match.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatForMasks
func GetPixelFormatForMasks(bpp int, rMask, gMask, bMask, aMask uint32) PixelFormat {
	return (PixelFormat)(C.SDL_GetPixelFormatForMasks((C.int)(bpp), (C.Uint32)(rMask), (C.Uint32)(gMask), (C.Uint32)(bMask), (C.Uint32)(aMask)))
}

// Create a [PixelFormatDetails] structure corresponding to a pixel format.
//
// f: one of the [PixelFormat] values.
//
// Returns a pointer to a [PixelFormatDetails] structure or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetPixelFormatDetails
func (format PixelFormat) Details() (*PixelFormatDetails, error) {
	details := C.SDL_GetPixelFormatDetails((C.SDL_PixelFormat)(format))
	if details == nil {
		return nil, getError()
	}
	return &PixelFormatDetails{
		PixelFormat(details.format),
		byte(details.bits_per_pixel),
		byte(details.bytes_per_pixel),
		uint32(details.Rmask),
		uint32(details.Gmask),
		uint32(details.Bmask),
		uint32(details.Amask),
		byte(details.Rbits),
		byte(details.Gbits),
		byte(details.Bbits),
		byte(details.Abits),
		byte(details.Rshift),
		byte(details.Gshift),
		byte(details.Bshift),
		byte(details.Ashift),
	}, nil
}

// Create a palette structure with the specified number of color entries.
//
// The palette entries are initialized to white.
//
// ncolors: represents the number of color entries in the color palette.
//
// Returns a new [Palette] structure or an error.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreatePalette
func CreatePalette(ncolors int) (*Palette, error) {
	palette := C.SDL_CreatePalette((C.int)(ncolors))
	if palette == nil {
		return nil, getError()
	}
	return &Palette{palette}, nil
}

// Set a range of colors in a palette.
//
// palette: the [Palette] structure to modify.
//
// colors: an array of [Color] structures to copy into the palette.
//
// first: the index of the first palette entry to modify.
//
// Returns nil on success or an error on failure.
//
// It is safe to call this function from any thread, as long as
// the palette is not modified or destroyed in another thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetPaletteColors
func (palette *Palette) SetColors(colors []Color, first int) error {
	if !C.SDL_SetPaletteColors((*C.SDL_Palette)(palette.internal), (*C.SDL_Color)(unsafe.Pointer(unsafe.SliceData(colors))), (C.int)(first), (C.int)(len(colors))) {
		return getError()
	}
	return nil
}

// Free a palette created with [CreatePalette].
//
// palette: the [Palette] structure to be freed.
//
// It is safe to call this function from any thread, as long as
// the palette is not modified or destroyed in another thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DestroyPalette
func (palette *Palette) Destroy() {
	C.SDL_DestroyPalette((*C.SDL_Palette)(palette.internal))
}

// Map an RGB triple to an opaque pixel value for a given pixel format.
//
// This function maps the RGB color value to the specified pixel format and
// returns the pixel value best approximating the given RGB color value for
// the given pixel format.
//
// If the format has a palette (8-bit) the index of the closest matching color
// in the palette will be returned.
//
// If the specified pixel format has an alpha component it will be returned as
// all 1 bits (fully opaque).
//
// If the pixel format bpp (color depth) is less than 32-bpp then the unused
// upper bits of the return value can safely be ignored (e.g., with a 16-bpp
// format the return value can be assigned to a uint16, and similarly a uint8
// for an 8-bpp format).
//
// format: a pointer to [PixelFormatDetails] describing the pixel format.
//
// palette: an optional palette for indexed formats, may be nil.
//
// r: the red component of the pixel in the range 0-255.
//
// g: the green component of the pixel in the range 0-255.
//
// b: the blue component of the pixel in the range 0-255.
//
// Returns a pixel value.
//
// It is safe to call this function from any thread, as long as
// the palette is not modified.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MapRGB
func (format *PixelFormatDetails) MapRGB(palette *Palette, r byte, g byte, b byte) uint32 {
	var p *C.SDL_Palette
	if palette != nil {
		p = palette.internal
	}
	return uint32(C.SDL_MapRGB(&C.SDL_PixelFormatDetails{
		format:          C.SDL_PixelFormat(format.Format),
		bits_per_pixel:  C.Uint8(format.BitsPerPixel),
		bytes_per_pixel: C.Uint8(format.BytesPerPixel),
		Rmask:           C.Uint32(format.RMask),
		Gmask:           C.Uint32(format.GMask),
		Bmask:           C.Uint32(format.BMask),
		Amask:           C.Uint32(format.AMask),
		Rbits:           C.Uint8(format.RBits),
		Gbits:           C.Uint8(format.GBits),
		Bbits:           C.Uint8(format.BBits),
		Abits:           C.Uint8(format.ABits),
		Rshift:          C.Uint8(format.RShift),
		Gshift:          C.Uint8(format.GShift),
		Bshift:          C.Uint8(format.BShift),
		Ashift:          C.Uint8(format.AShift),
	}, p, C.Uint8(r), C.Uint8(g), C.Uint8(b)))
}

// Map an RGBA quadruple to a pixel value for a given pixel format.
//
// This function maps the RGBA color value to the specified pixel format and
// returns the pixel value best approximating the given RGBA color value for
// the given pixel format.
//
// If the specified pixel format has no alpha component the alpha value will
// be ignored (as it will be in formats with a palette).
//
// If the format has a palette (8-bit) the index of the closest matching color
// in the palette will be returned.
//
// If the pixel format bpp (color depth) is less than 32-bpp then the unused
// upper bits of the return value can safely be ignored (e.g., with a 16-bpp
// format the return value can be assigned to a uint16, and similarly a uint8
// for an 8-bpp format).
//
// f: a pointer to [PixelFormatDetails] describing the pixel
// format.
//
// palette: an optional palette for indexed formats, may be nil.
//
// r: the red component of the pixel in the range 0-255.
//
// g: the green component of the pixel in the range 0-255.
//
// b: the blue component of the pixel in the range 0-255.
//
// a: the alpha component of the pixel in the range 0-255.
//
// Returns a pixel value.
//
// It is safe to call this function from any thread, as long as
// the palette is not modified.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MapRGBA
func (format *PixelFormatDetails) MapRGBA(palette *Palette, r byte, g byte, b byte, a byte) uint32 {
	var p *C.SDL_Palette
	if palette != nil {
		p = palette.internal
	}
	return uint32(C.SDL_MapRGBA(&C.SDL_PixelFormatDetails{
		format:          C.SDL_PixelFormat(format.Format),
		bits_per_pixel:  C.Uint8(format.BitsPerPixel),
		bytes_per_pixel: C.Uint8(format.BytesPerPixel),
		Rmask:           C.Uint32(format.RMask),
		Gmask:           C.Uint32(format.GMask),
		Bmask:           C.Uint32(format.BMask),
		Amask:           C.Uint32(format.AMask),
		Rbits:           C.Uint8(format.RBits),
		Gbits:           C.Uint8(format.GBits),
		Bbits:           C.Uint8(format.BBits),
		Abits:           C.Uint8(format.ABits),
		Rshift:          C.Uint8(format.RShift),
		Gshift:          C.Uint8(format.GShift),
		Bshift:          C.Uint8(format.BShift),
		Ashift:          C.Uint8(format.AShift),
	}, p, C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a)))
}

// Get RGB values from a pixel in the specified format.
//
// This function uses the entire 8-bit [0..255] range when converting color
// components from pixel formats with less than 8-bits per RGB component
// (e.g., a completely white pixel in 16-bit RGB565 format would return [0xff,
// 0xff, 0xff] not [0xf8, 0xfc, 0xf8]).
//
// pixel: a pixel value.
//
// format: a pointer to [PixelFormatDetails] describing the pixel format.
//
// palette: an optional palette for indexed formats, may be nil.
//
// r: the red component.
//
// g: the green component.
//
// b: the blue component.
//
// It is safe to call this function from any thread, as long as
// the palette is not modified.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRGB
func (format *PixelFormatDetails) GetRGB(pixel uint32, palette *Palette) (r, b, g byte) {
	var p *C.SDL_Palette
	if palette != nil {
		p = palette.internal
	}
	C.SDL_GetRGB((C.Uint32)(pixel), &C.SDL_PixelFormatDetails{
		format:          C.SDL_PixelFormat(format.Format),
		bits_per_pixel:  C.Uint8(format.BitsPerPixel),
		bytes_per_pixel: C.Uint8(format.BytesPerPixel),
		Rmask:           C.Uint32(format.RMask),
		Gmask:           C.Uint32(format.GMask),
		Bmask:           C.Uint32(format.BMask),
		Amask:           C.Uint32(format.AMask),
		Rbits:           C.Uint8(format.RBits),
		Gbits:           C.Uint8(format.GBits),
		Bbits:           C.Uint8(format.BBits),
		Abits:           C.Uint8(format.ABits),
		Rshift:          C.Uint8(format.RShift),
		Gshift:          C.Uint8(format.GShift),
		Bshift:          C.Uint8(format.BShift),
		Ashift:          C.Uint8(format.AShift),
	}, p, (*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b))
	return
}

// Get RGBA values from a pixel in the specified format.
//
// This function uses the entire 8-bit [0..255] range when converting color
// components from pixel formats with less than 8-bits per RGB component
// (e.g., a completely white pixel in 16-bit RGB565 format would return [0xff,
// 0xff, 0xff] not [0xf8, 0xfc, 0xf8]).
//
// If the surface has no alpha component, the alpha will be returned as 0xff
// (100% opaque).
//
// pixel: a pixel value.
//
// format: a pointer to [PixelFormatDetails] describing the pixel format.
//
// palette: an optional palette for indexed formats, may be nil.
//
// r: the red component.
//
// g: the green component.
//
// b: the blue component.
//
// a: the alpha component.
//
// It is safe to call this function from any thread, as long as
// the palette is not modified.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRGBA
func (format *PixelFormatDetails) GetRGBA(pixel uint32, palette *Palette) (r, g, b, a byte) {
	var p *C.SDL_Palette
	if palette != nil {
		p = palette.internal
	}
	C.SDL_GetRGBA((C.Uint32)(pixel), &C.SDL_PixelFormatDetails{
		C.SDL_PixelFormat(format.Format),
		C.Uint8(format.BitsPerPixel),
		C.Uint8(format.BytesPerPixel),
		[2]C.Uint8{},
		C.Uint32(format.RMask),
		C.Uint32(format.GMask),
		C.Uint32(format.BMask),
		C.Uint32(format.AMask),
		C.Uint8(format.RBits),
		C.Uint8(format.GBits),
		C.Uint8(format.BBits),
		C.Uint8(format.ABits),
		C.Uint8(format.RShift),
		C.Uint8(format.GShift),
		C.Uint8(format.BShift),
		C.Uint8(format.AShift),
	}, p, (*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b), (*C.Uint8)(&a))
	return
}
