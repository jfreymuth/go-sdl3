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
// #cgo noescape SDL_CreateSurface
// #cgo nocallback SDL_CreateSurface
// #cgo noescape SDL_CreateSurfaceFrom
// #cgo nocallback SDL_CreateSurfaceFrom
// #cgo noescape SDL_DestroySurface
// #cgo noescape SDL_GetSurfaceProperties
// #cgo nocallback SDL_GetSurfaceProperties
// #cgo noescape SDL_SetSurfaceColorspace
// #cgo nocallback SDL_SetSurfaceColorspace
// #cgo noescape SDL_GetSurfaceColorspace
// #cgo nocallback SDL_GetSurfaceColorspace
// #cgo noescape SDL_CreateSurfacePalette
// #cgo nocallback SDL_CreateSurfacePalette
// #cgo noescape SDL_SetSurfacePalette
// #cgo nocallback SDL_SetSurfacePalette
// #cgo noescape SDL_GetSurfacePalette
// #cgo nocallback SDL_GetSurfacePalette
// #cgo noescape SDL_AddSurfaceAlternateImage
// #cgo nocallback SDL_AddSurfaceAlternateImage
// #cgo noescape SDL_SurfaceHasAlternateImages
// #cgo nocallback SDL_SurfaceHasAlternateImages
// #cgo noescape SDL_GetSurfaceImages
// #cgo nocallback SDL_GetSurfaceImages
// #cgo noescape SDL_RemoveSurfaceAlternateImages
// #cgo noescape SDL_LockSurface
// #cgo nocallback SDL_LockSurface
// #cgo noescape SDL_UnlockSurface
// #cgo nocallback SDL_UnlockSurface
// #cgo noescape SDL_LoadBMP_IO
// #cgo noescape SDL_LoadBMP
// #cgo nocallback SDL_LoadBMP
// #cgo noescape SDL_SaveBMP_IO
// #cgo noescape SDL_SaveBMP
// #cgo nocallback SDL_SaveBMP
// #cgo noescape SDL_SetSurfaceRLE
// #cgo nocallback SDL_SetSurfaceRLE
// #cgo noescape SDL_SurfaceHasRLE
// #cgo nocallback SDL_SurfaceHasRLE
// #cgo noescape SDL_SetSurfaceColorKey
// #cgo nocallback SDL_SetSurfaceColorKey
// #cgo noescape SDL_SurfaceHasColorKey
// #cgo nocallback SDL_SurfaceHasColorKey
// #cgo noescape SDL_GetSurfaceColorKey
// #cgo nocallback SDL_GetSurfaceColorKey
// #cgo noescape SDL_SetSurfaceColorMod
// #cgo nocallback SDL_SetSurfaceColorMod
// #cgo noescape SDL_GetSurfaceColorMod
// #cgo nocallback SDL_GetSurfaceColorMod
// #cgo noescape SDL_SetSurfaceAlphaMod
// #cgo nocallback SDL_SetSurfaceAlphaMod
// #cgo noescape SDL_GetSurfaceAlphaMod
// #cgo nocallback SDL_GetSurfaceAlphaMod
// #cgo noescape SDL_SetSurfaceBlendMode
// #cgo nocallback SDL_SetSurfaceBlendMode
// #cgo noescape SDL_GetSurfaceBlendMode
// #cgo nocallback SDL_GetSurfaceBlendMode
// #cgo noescape SDL_SetSurfaceClipRect
// #cgo nocallback SDL_SetSurfaceClipRect
// #cgo noescape SDL_GetSurfaceClipRect
// #cgo nocallback SDL_GetSurfaceClipRect
// #cgo noescape SDL_FlipSurface
// #cgo nocallback SDL_FlipSurface
// #cgo noescape SDL_DuplicateSurface
// #cgo nocallback SDL_DuplicateSurface
// #cgo noescape SDL_ScaleSurface
// #cgo nocallback SDL_ScaleSurface
// #cgo noescape SDL_ConvertSurface
// #cgo nocallback SDL_ConvertSurface
// #cgo noescape SDL_ConvertSurfaceAndColorspace
// #cgo nocallback SDL_ConvertSurfaceAndColorspace
// #cgo noescape SDL_ConvertPixels
// #cgo nocallback SDL_ConvertPixels
// #cgo noescape SDL_ConvertPixelsAndColorspace
// #cgo nocallback SDL_ConvertPixelsAndColorspace
// #cgo noescape SDL_PremultiplyAlpha
// #cgo nocallback SDL_PremultiplyAlpha
// #cgo noescape SDL_PremultiplySurfaceAlpha
// #cgo nocallback SDL_PremultiplySurfaceAlpha
// #cgo noescape SDL_ClearSurface
// #cgo nocallback SDL_ClearSurface
// #cgo noescape SDL_FillSurfaceRect
// #cgo nocallback SDL_FillSurfaceRect
// #cgo noescape SDL_FillSurfaceRects
// #cgo nocallback SDL_FillSurfaceRects
// #cgo noescape SDL_BlitSurface
// #cgo nocallback SDL_BlitSurface
// #cgo noescape SDL_BlitSurfaceUnchecked
// #cgo nocallback SDL_BlitSurfaceUnchecked
// #cgo noescape SDL_BlitSurfaceScaled
// #cgo nocallback SDL_BlitSurfaceScaled
// #cgo noescape SDL_BlitSurfaceUncheckedScaled
// #cgo nocallback SDL_BlitSurfaceUncheckedScaled
// #cgo noescape SDL_BlitSurfaceTiled
// #cgo nocallback SDL_BlitSurfaceTiled
// #cgo noescape SDL_BlitSurfaceTiledWithScale
// #cgo nocallback SDL_BlitSurfaceTiledWithScale
// #cgo noescape SDL_BlitSurface9Grid
// #cgo nocallback SDL_BlitSurface9Grid
// #cgo noescape SDL_MapSurfaceRGB
// #cgo nocallback SDL_MapSurfaceRGB
// #cgo noescape SDL_MapSurfaceRGBA
// #cgo nocallback SDL_MapSurfaceRGBA
// #cgo noescape SDL_ReadSurfacePixel
// #cgo nocallback SDL_ReadSurfacePixel
// #cgo noescape SDL_ReadSurfacePixelFloat
// #cgo nocallback SDL_ReadSurfacePixelFloat
// #cgo noescape SDL_WriteSurfacePixel
// #cgo nocallback SDL_WriteSurfacePixel
// #cgo noescape SDL_WriteSurfacePixelFloat
// #cgo nocallback SDL_WriteSurfacePixelFloat
// #include <SDL3/SDL.h>
import "C"
import "unsafe"

// # CategorySurface
//
// SDL surfaces are buffers of pixels in system RAM. These are useful for
// passing around and manipulating images that are not stored in GPU memory.
//
// [Surface] makes serious efforts to manage images in various formats, and
// provides a reasonable toolbox for transforming the data, including copying
// between surfaces, filling rectangles in the image data, etc.
//
// There is also a simple .bmp loader, [LoadBMP]. SDL itself does not
// provide loaders for various other file formats, but there are several
// excellent external libraries that do, including its own satellite library,
// SDL_image:
//
// https://github.com/libsdl-org/SDL_image

// The flags on a [Surface].
//
// These are generally considered read-only.
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SurfaceFlags
type SurfaceFlags uint32

const (
	SurfacePreallocated SurfaceFlags = 0x00000001 //Surface uses preallocated pixel memory
	SurfaceLockNeeded   SurfaceFlags = 0x00000002 //Surface needs to be locked to access pixels
	SurfaceLocked       SurfaceFlags = 0x00000004 //Surface is currently locked
	SurfaceSimdAligned  SurfaceFlags = 0x00000008 //Surface uses pixel memory allocated with SDL_aligned_alloc()
)

// Evaluates to true if the surface needs to be locked before access.
//
// This macro is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MUSTLOCK
func (surface *Surface) MustLock() bool {
	return SurfaceFlags(surface.internal.flags)&SurfaceLockNeeded == SurfaceLockNeeded
}

// The scaling mode.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ScaleMode
type ScaleMode uint32

const (
	ScalemodeNearest ScaleMode = iota // nearest pixel sampling
	ScalemodeLinear                   // linear filtering
)

// The flip mode.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_FlipMode
type FlipMode uint32

const (
	FlipNone       FlipMode = iota // Do not flip
	FlipHorizontal                 // flip horizontally
	FlipVertical                   // flip vertically
)

// A collection of pixels used in software blitting.
//
// Pixels are arranged in memory in rows, with the top row first. Each row
// occupies an amount of memory given by the pitch (sometimes known as the row
// stride in non-SDL APIs).
//
// Within each row, pixels are arranged from left to right until the width is
// reached. Each pixel occupies a number of bits appropriate for its format,
// with most formats representing each pixel as one or more whole bytes (in
// some indexed formats, instead multiple pixels are packed into each byte),
// and a byte order given by the format. After encoding all pixels, any
// remaining bytes to reach the pitch are used as padding to reach a desired
// alignment, and have undefined contents.
//
// When a surface holds YUV format data, the planes are assumed to be
// contiguous without padding between them, e.g. a 32x32 surface in NV12
// format with a pitch of 32 would consist of 32x32 bytes of Y plane followed
// by 32x16 bytes of UV plane.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Surface
type Surface struct {
	internal *C.struct_SDL_Surface
	keep     []byte
}

// The flags of the surface
func (surface *Surface) Flags() SurfaceFlags { return SurfaceFlags(surface.internal.flags) }

// The format of the surface
func (surface *Surface) Format() PixelFormat { return PixelFormat(surface.internal.format) }

// The width of the surface
func (surface *Surface) Width() int { return int(surface.internal.w) }

// The height of the surface
func (surface *Surface) Height() int { return int(surface.internal.h) }

// The distance in bytes between rows of pixels
func (surface *Surface) Pitch() int { return int(surface.internal.pitch) }

// A pointer to the pixels of the surface, the pixels are writeable if non-nil
func (surface *Surface) Pixels() []byte {
	return unsafe.Slice((*byte)(surface.internal.pixels), surface.internal.h*surface.internal.pitch)
}

// Allocate a new surface with a specific pixel format.
//
// The pixels of the new surface are initialized to zero.
//
// width: the width of the surface.
//
// height: the height of the surface.
//
// format: the [PixelFormat] for the new surface's pixel format.
//
// Returns the new [Surface] structure that is created or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateSurface
func CreateSurface(width int, height int, format PixelFormat) (*Surface, error) {
	s := C.SDL_CreateSurface((C.int)(width), (C.int)(height), (C.SDL_PixelFormat)(format))
	if s == nil {
		return nil, getError()
	}
	return &Surface{s, nil}, nil
}

// Allocate a new surface with a specific pixel format and existing pixel
// data.
//
// No copy is made of the pixel data. Pixel data is not managed automatically;
// you must free the surface before you free the pixel data.
//
// Pitch is the offset in bytes from one row of pixels to the next, e.g.
// width*4 for [PixelformatRGBA8888].
//
// You may pass nil for pixels and 0 for pitch to create a surface that you
// will fill in with valid values later.
//
// width: the width of the surface.
//
// height: the height of the surface.
//
// format: the [PixelFormat] for the new surface's pixel format.
//
// pixels: a pointer to existing pixel data.
//
// pitch: the number of bytes between each row, including padding.
//
// Returns the new [Surface] structure that is created or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateSurfaceFrom
func CreateSurfaceFrom(width int, height int, format PixelFormat, pixels []byte, pitch int) (*Surface, error) {
	s := C.SDL_CreateSurfaceFrom((C.int)(width), (C.int)(height), (C.SDL_PixelFormat)(format), unsafe.Pointer(unsafe.SliceData(pixels)), (C.int)(pitch))
	if s == nil {
		return nil, getError()
	}
	return &Surface{s, pixels}, nil
}

// Free a surface.
//
// surface: the [Surface] to free.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DestroySurface
func (surface *Surface) Destroy() {
	if surface.internal.refcount <= 1 {
		surface.keep = nil
	}
	C.SDL_DestroySurface(surface.internal)
}

// Get the properties associated with a surface.
//
// The following properties are understood by SDL:
//
//   - [PropSurfaceSDRWhitePointFloat]: for HDR10 and floating point
//     surfaces, this defines the value of 100% diffuse white, with higher
//     values being displayed in the High Dynamic Range headroom. This defaults
//     to 203 for HDR10 surfaces and 1.0 for floating point surfaces.
//   - [PropSurfaceHDRHeadroomFloat]: for HDR10 and floating point
//     surfaces, this defines the maximum dynamic range used by the content, in
//     terms of the SDR white point. This defaults to 0.0, which disables tone
//     mapping.
//   - [PropSurfaceTonemapOperatorString]: the tone mapping operator
//     used when compressing from a surface with high dynamic range to another
//     with lower dynamic range. Currently this supports "chrome", which uses
//     the same tone mapping that Chrome uses for HDR content, the form "*=N",
//     where N is a floating point scale factor applied in linear space, and
//     "none", which disables tone mapping. This defaults to "chrome".
//
// surface: the [Surface] structure to query.
//
// Returns a valid property ID or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSurfaceProperties
func (surface *Surface) Properties() (PropertiesID, error) {
	props := (PropertiesID)(C.SDL_GetSurfaceProperties(surface.internal))
	if props == 0 {
		return 0, getError()
	}
	return props, nil
}

const PropSurfaceSDRWhitePointFloat = "SDL.surface.SDR_white_point"
const PropSurfaceHDRHeadroomFloat = "SDL.surface.HDR_headroom"
const PropSurfaceTonemapOperatorString = "SDL.surface.tonemap"

// Set the colorspace used by a surface.
//
// Setting the colorspace doesn't change the pixels, only how they are
// interpreted in color operations.
//
// surface: the [Surface] structure to update.
//
// colorspace: a [Colorspace] value describing the surface colorspace.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorspace
func (surface *Surface) SetColorspace(colorspace Colorspace) error {
	if !C.SDL_SetSurfaceColorspace(surface.internal, (C.SDL_Colorspace)(colorspace)) {
		return getError()
	}
	return nil
}

// Get the colorspace used by a surface.
//
// The colorspace defaults to [ColorspaceSRGBLinear] for floating point
// formats, [ColorspaceHDR10] for 10-bit formats, [ColorspaceSRGB] for
// other RGB surfaces and [ColorspaceBT709Full] for YUV textures.
//
// surface: the [Surface] structure to query.
//
// Returns the colorspace used by the surface.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorspace
func (surface *Surface) Colorspace() Colorspace {
	return (Colorspace)(C.SDL_GetSurfaceColorspace(surface.internal))
}

// Create a palette and associate it with a surface.
//
// This function creates a palette compatible with the provided surface. The
// palette is then returned for you to modify, and the surface will
// automatically use the new palette in future operations. You do not need to
// destroy the returned palette, it will be freed when the reference count
// reaches 0, usually when the surface is destroyed.
//
// Bitmap surfaces (with format [PixelformatIndex1LSB] or
// [PixelformatIndex1MSB]) will have the palette initialized with 0 as
// white and 1 as black. Other surfaces will get a palette initialized with
// white in every entry.
//
// If this function is called for a surface that already has a palette, a new
// palette will be created to replace it.
//
// surface: the [Surface] structure to update.
//
// Returns a new [Palette] structure or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateSurfacePalette
func (surface *Surface) CreatePalette() (*Palette, error) {
	p := C.SDL_CreateSurfacePalette(surface.internal)
	if p == nil {
		return nil, getError()
	}
	return &Palette{p}, nil
}

// Set the palette used by a surface.
//
// A single palette can be shared with many surfaces.
//
// surface: the [Surface] structure to update.
//
// palette: the [Palette] structure to use.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetSurfacePalette
func (surface *Surface) SetPalette(palette *Palette) error {
	if !C.SDL_SetSurfacePalette(surface.internal, palette.internal) {
		return getError()
	}
	return nil
}

// Get the palette used by a surface.
//
// surface: the [Surface] structure to query.
//
// Returns a pointer to the palette used by the surface or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSurfacePalette
func GetSurfacePalette(surface *Surface) *Palette {
	p := C.SDL_GetSurfacePalette(surface.internal)
	if p == nil {
		return nil
	}
	return &Palette{p}
}

// Add an alternate version of a surface.
//
// This function adds an alternate version of this surface, usually used for
// content with high DPI representations like cursors or icons. The size,
// format, and content do not need to match the original surface, and these
// alternate versions will not be updated when the original surface changes.
//
// This function adds a reference to the alternate version, so you should call
// [Surface.Destroy] on the image after this call.
//
// surface: the [Surface] structure to update.
//
// image: a pointer to an alternate [Surface] to associate with this
// surface.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AddSurfaceAlternateImage
func (surface *Surface) AddAlternateImage(image *Surface) error {
	if !C.SDL_AddSurfaceAlternateImage(surface.internal, image.internal) {
		return getError()
	}
	return nil
}

// Return whether a surface has alternate versions available.
//
// surface: the [Surface] structure to query.
//
// Returns true if alternate versions are available or false otherwise.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SurfaceHasAlternateImages
func (surface *Surface) HasAlternateImages() bool {
	return (bool)(C.SDL_SurfaceHasAlternateImages(surface.internal))
}

// Get an array including all versions of a surface.
//
// This returns all versions of a surface, with the surface being queried as
// the first element in the returned array.
//
// Freeing the array of surfaces does not affect the surfaces in the array.
// They are still referenced by the surface being queried and will be cleaned
// up normally.
//
// surface: the [Surface] structure to query.
//
// Returns a slice of [Surface] pointers or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSurfaceImages
func (surface *Surface) GetImages() ([]*Surface, error) {
	var count C.int
	s := C.SDL_GetSurfaceImages(surface.internal, &count)
	if s == nil {
		return nil, getError()
	}
	result := make([]*Surface, count)
	for i, s := range unsafe.Slice(s, count) {
		result[i] = &Surface{s, nil}
	}
	C.SDL_free(unsafe.Pointer(s))
	return result, nil
}

// Remove all alternate versions of a surface.
//
// This function removes a reference from all the alternative versions,
// destroying them if this is the last reference to them.
//
// surface: the [Surface] structure to update.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RemoveSurfaceAlternateImages
func (surface *Surface) RemoveAlternateImages() {
	C.SDL_RemoveSurfaceAlternateImages(surface.internal)
}

// Set up a surface for directly accessing the pixels.
//
// Between calls to [Surface.Lock] / [Surface.Unlock], you can write to
// and read from [Surface.Pixels], using the pixel format stored in
// [Surface.Format]. Once you are done accessing the surface, you should use
// [Surface.Unlock] to release it.
//
// Not all surfaces require locking. If [Surface.MustLock] returns false, then
// you can read and write to the surface at any time, and the pixel format of
// the surface will not change.
//
// surface: the [Surface] structure to be locked.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_LockSurface
func (surface *Surface) Lock() error {
	if !C.SDL_LockSurface(surface.internal) {
		return getError()
	}
	return nil
}

// Release a surface after directly accessing the pixels.
//
// surface: the [Surface] structure to be unlocked.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_UnlockSurface
func (surface *Surface) Unlock() {
	C.SDL_UnlockSurface(surface.internal)
}

// Load a BMP image from a seekable SDL data stream.
//
// The new surface should be freed with [Surface.Destroy]. Not doing so
// will result in a memory leak.
//
// src: the data stream for the surface.
//
// closeio: if true, calls [IOStream.Close] on src before returning, even
// in the case of an error.
//
// Returns a pointer to a new [Surface] structure or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_LoadBMP_IO
func LoadBMP_IO(src *IOStream, closeio bool) (*Surface, error) {
	s := C.SDL_LoadBMP_IO((*C.SDL_IOStream)(src), (C.bool)(closeio))
	if s == nil {
		return nil, getError()
	}
	return &Surface{s, nil}, nil
}

// Load a BMP image from a file.
//
// The new surface should be freed with [Sureface.Destroy]. Not doing so
// will result in a memory leak.
//
// file: the BMP file to load.
//
// Returns a pointer to a new [Surface] structure or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_LoadBMP
func LoadBMP(file string) (*Surface, error) {
	s := C.SDL_LoadBMP(tmpstring(file))
	if s == nil {
		return nil, getError()
	}
	return &Surface{s, nil}, nil
}

// Save a surface to a seekable SDL data stream in BMP format.
//
// Surfaces with a 24-bit, 32-bit and paletted 8-bit format get saved in the
// BMP directly. Other RGB formats with 8-bit or higher get converted to a
// 24-bit surface or, if they have an alpha mask or a colorkey, to a 32-bit
// surface before they are saved. YUV and paletted 1-bit and 4-bit formats are
// not supported.
//
// surface: the [Surface] structure containing the image to be saved.
//
// dst: a data stream to save to.
//
// closeio: if true, calls [IOStream.Close] on dst before returning, even
// in the case of an error.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SaveBMP_IO
func SaveBMP_IO(surface *Surface, dst *IOStream, closeio bool) error {
	if !C.SDL_SaveBMP_IO(surface.internal, (*C.SDL_IOStream)(dst), (C.bool)(closeio)) {
		return getError()
	}
	return nil
}

// Save a surface to a file.
//
// Surfaces with a 24-bit, 32-bit and paletted 8-bit format get saved in the
// BMP directly. Other RGB formats with 8-bit or higher get converted to a
// 24-bit surface or, if they have an alpha mask or a colorkey, to a 32-bit
// surface before they are saved. YUV and paletted 1-bit and 4-bit formats are
// not supported.
//
// surface: the [Surface] structure containing the image to be saved.
//
// file: a file to save to.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SaveBMP
func SaveBMP(surface *Surface, file string) error {
	if !C.SDL_SaveBMP(surface.internal, tmpstring(file)) {
		return getError()
	}
	return nil
}

// Set the RLE acceleration hint for a surface.
//
// If RLE is enabled, color key and alpha blending blits are much faster, but
// the surface must be locked before directly accessing the pixels.
//
// surface: the [Surface] structure to optimize.
//
// enabled: true to enable RLE acceleration, false to disable it.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetSurfaceRLE
func (surface *Surface) SetRLE(enabled bool) error {
	if !C.SDL_SetSurfaceRLE(surface.internal, (C.bool)(enabled)) {
		return getError()
	}
	return nil
}

// Returns whether the surface is RLE enabled.
//
// surface: the [Surface] structure to query.
//
// Returns true if the surface is RLE enabled, false otherwise.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SurfaceHasRLE
func (surface *Surface) HasRLE() bool {
	return (bool)(C.SDL_SurfaceHasRLE(surface.internal))
}

// Set the color key (transparent pixel) in a surface.
//
// The color key defines a pixel value that will be treated as transparent in
// a blit. For example, one can use this to specify that cyan pixels should be
// considered transparent, and therefore not rendered.
//
// It is a pixel of the format used by the surface, as generated by
// [PixelFormatDetails.MapRGB].
//
// surface: the [Surface] structure to update.
//
// enabled: true to enable color key, false to disable color key.
//
// key: the transparent pixel.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorKey
func (surface *Surface) SetColorKey(enabled bool, key uint32) error {
	if !C.SDL_SetSurfaceColorKey(surface.internal, (C.bool)(enabled), (C.Uint32)(key)) {
		return getError()
	}
	return nil
}

// Returns whether the surface has a color key.
//
// surface: the [Surface] structure to query.
//
// Returns true if the surface has a color key, false otherwise.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SurfaceHasColorKey
func (surface *Surface) HasColorKey() bool {
	return (bool)(C.SDL_SurfaceHasColorKey(surface.internal))
}

// Get the color key (transparent pixel) for a surface.
//
// The color key is a pixel of the format used by the surface, as generated by
// [PixelFormatDetails.MapRGB].
//
// If the surface doesn't have color key enabled this function returns false.
//
// surface: the [Surface] structure to query.
//
// key: the transparent pixel.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorKey
func (surface *Surface) ColorKey() (uint32, error) {
	var key C.Uint32
	if !C.SDL_GetSurfaceColorKey(surface.internal, &key) {
		return 0, getError()
	}
	return uint32(key), nil
}

// Set an additional color value multiplied into blit operations.
//
// When this surface is blitted, during the blit operation each source color
// channel is modulated by the appropriate color value according to the
// following formula:
//
//	srcC = srcC * (color / 255)
//
// surface: the [Surface] structure to update.
//
// r: the red color value multiplied into blit operations.
//
// g: the green color value multiplied into blit operations.
//
// b: the blue color value multiplied into blit operations.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetSurfaceColorMod
func (surface *Surface) SetColorMod(r byte, g byte, b byte) error {
	if !C.SDL_SetSurfaceColorMod(surface.internal, (C.Uint8)(r), (C.Uint8)(g), (C.Uint8)(b)) {
		return getError()
	}
	return nil
}

// Get the additional color value multiplied into blit operations.
//
// surface: the [Surface] structure to query.
//
// r: the current red color value.
//
// g: the current green color value.
//
// b: the current blue color value.
//
// Returns the current color mod or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSurfaceColorMod
func (surface *Surface) ColorMod() (r, g, b byte, err error) {
	var cr, cg, cb C.Uint8
	if !C.SDL_GetSurfaceColorMod(surface.internal, &cr, &cg, &cb) {
		return 0, 0, 0, getError()
	}
	return byte(cr), byte(cg), byte(cb), nil
}

// Set an additional alpha value used in blit operations.
//
// When this surface is blitted, during the blit operation the source alpha
// value is modulated by this alpha value according to the following formula:
//
//	srcA = srcA * (alpha / 255)
//
// surface: the [Surface] structure to update.
//
// alpha: the alpha value multiplied into blit operations.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetSurfaceAlphaMod
func (surface *Surface) SetAlphaMod(alpha byte) error {
	if !C.SDL_SetSurfaceAlphaMod(surface.internal, (C.Uint8)(alpha)) {
		return getError()
	}
	return nil
}

// Get the additional alpha value used in blit operations.
//
// surface: the [Surface] structure to query.
//
// alpha: the current alpha value.
//
// Returns the current alpha value or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSurfaceAlphaMod
func (surface *Surface) AlphaMod() (byte, error) {
	var alpha C.Uint8
	if !C.SDL_GetSurfaceAlphaMod(surface.internal, &alpha) {
		return 0, getError()
	}
	return byte(alpha), nil
}

// Set the blend mode used for blit operations.
//
// To copy a surface to another surface (or texture) without blending with the
// existing data, the blendmode of the SOURCE surface should be set to
// [BlendmodeNone].
//
// surface: the [Surface] structure to update.
//
// blendMode: the [BlendMode] to use for blit blending.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetSurfaceBlendMode
func (surface *Surface) SetBlendMode(blendMode BlendMode) error {
	if !C.SDL_SetSurfaceBlendMode(surface.internal, (C.SDL_BlendMode)(blendMode)) {
		return getError()
	}
	return nil
}

// Get the blend mode used for blit operations.
//
// surface: the [Surface] structure to query.
//
// blendMode: the current [BlendMode].
//
// Returns the current [BlendMode] or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSurfaceBlendMode
func (surface *Surface) BlendMode() (BlendMode, error) {
	var blendMode C.SDL_BlendMode
	if !C.SDL_GetSurfaceBlendMode(surface.internal, &blendMode) {
		return 0, getError()
	}
	return BlendMode(blendMode), nil
}

// Set the clipping rectangle for a surface.
//
// When surface is the destination of a blit, only the area within the clip
// rectangle is drawn into.
//
// Note that blits are automatically clipped to the edges of the source and
// destination surfaces.
//
// surface: the [Surface] structure to be clipped.
//
// rect: the [Rect] structure representing the clipping rectangle, or
// nil to disable clipping.
//
// Returns true if the rectangle intersects the surface, otherwise false and
// blits will be completely clipped.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetSurfaceClipRect
func (surface *Surface) SetClipRect(rect *Rect) bool {
	var crect *C.SDL_Rect
	if rect != nil {
		crect = &C.SDL_Rect{C.int(rect.X), C.int(rect.Y), C.int(rect.W), C.int(rect.H)}
	}
	return (bool)(C.SDL_SetSurfaceClipRect(surface.internal, crect))
}

// Get the clipping rectangle for a surface.
//
// When surface is the destination of a blit, only the area within the clip
// rectangle is drawn into.
//
// surface: the [Surface] structure representing the surface to be
// clipped.
//
// rect: the clipping rectangle for the surface.
//
// Returns the clipping rectangle for the surface or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetSurfaceClipRect
func (surface *Surface) ClipRect() (Rect, error) {
	var rect C.SDL_Rect
	if !C.SDL_GetSurfaceClipRect(surface.internal, &rect) {
		return Rect{}, getError()
	}
	return Rect{int(rect.x), int(rect.y), int(rect.w), int(rect.h)}, nil
}

// Flip a surface vertically or horizontally.
//
// surface: the surface to flip.
//
// flip: the direction to flip.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_FlipSurface
func (surface *Surface) Flip(flip FlipMode) error {
	if !C.SDL_FlipSurface(surface.internal, (C.SDL_FlipMode)(flip)) {
		return getError()
	}
	return nil
}

// Creates a new surface identical to the existing surface.
//
// If the original surface has alternate images, the new surface will have a
// reference to them as well.
//
// The returned surface should be freed with [Surface.Destroy].
//
// surface: the surface to duplicate.
//
// Returns a copy of the surface or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DuplicateSurface
func (surface *Surface) Duplicate() (*Surface, error) {
	s := C.SDL_DuplicateSurface(surface.internal)
	if s == nil {
		return nil, getError()
	}
	return &Surface{s, nil}, nil
}

// Creates a new surface identical to the existing surface, scaled to the
// desired size.
//
// The returned surface should be freed with [Surface.Destroy].
//
// surface: the surface to duplicate and scale.
//
// width: the width of the new surface.
//
// height: the height of the new surface.
//
// scaleMode: the [ScaleMode] to be used.
//
// Returns a copy of the surface or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ScaleSurface
func (surface *Surface) Scale(width int, height int, scaleMode ScaleMode) (*Surface, error) {
	s := C.SDL_ScaleSurface(surface.internal, (C.int)(width), (C.int)(height), (C.SDL_ScaleMode)(scaleMode))
	if s == nil {
		return nil, getError()
	}
	return &Surface{s, nil}, nil
}

// Copy an existing surface to a new surface of the specified format.
//
// This function is used to optimize images for faster *repeat* blitting. This
// is accomplished by converting the original and storing the result as a new
// surface. The new, optimized surface can then be used as the source for
// future blits, making them faster.
//
// If you are converting to an indexed surface and want to map colors to a
// palette, you can use [Surface.ConvertColorspace] instead.
//
// If the original surface has alternate images, the new surface will have a
// reference to them as well.
//
// surface: the existing [Surface] structure to convert.
//
// format: the new pixel format.
//
// Returns the new [Surface] structure that is created or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ConvertSurface
func (surface *Surface) Convert(format PixelFormat) (*Surface, error) {
	s := C.SDL_ConvertSurface(surface.internal, (C.SDL_PixelFormat)(format))
	if s == nil {
		return nil, getError()
	}
	return &Surface{s, nil}, nil
}

// Copy an existing surface to a new surface of the specified format and
// colorspace.
//
// This function converts an existing surface to a new format and colorspace
// and returns the new surface. This will perform any pixel format and
// colorspace conversion needed.
//
// If the original surface has alternate images, the new surface will have a
// reference to them as well.
//
// surface: the existing [Surface] structure to convert.
//
// format: the new pixel format.
//
// palette: an optional palette to use for indexed formats, may be nil.
//
// colorspace: the new colorspace.
//
// props: a [PropertiesID] with additional color properties, or 0.
//
// Returns the new [Surface] structure that is created or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ConvertSurfaceAndColorspace
func (surface *Surface) ConvertColorspace(format PixelFormat, palette *Palette, colorspace Colorspace, props PropertiesID) (*Surface, error) {
	var p *C.SDL_Palette
	if palette != nil {
		p = palette.internal
	}
	s := C.SDL_ConvertSurfaceAndColorspace(surface.internal, (C.SDL_PixelFormat)(format), p, (C.SDL_Colorspace)(colorspace), (C.SDL_PropertiesID)(props))
	if s == nil {
		return nil, getError()
	}
	return &Surface{s, nil}, nil
}

// Copy a block of pixels of one format to another format.
//
// width: the width of the block to copy, in pixels.
//
// height: the height of the block to copy, in pixels.
//
// srcFormat: a [PixelFormat] value of the src pixels format.
//
// src: a pointer to the source pixels.
//
// srcPitch: the pitch of the source pixels, in bytes.
//
// dstFormat: a [PixelFormat] value of the dst pixels format.
//
// dst: a slice to be filled in with new pixel data.
//
// dstPitch: the pitch of the destination pixels, in bytes.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ConvertPixels
func ConvertPixels(width int, height int, srcFormat PixelFormat, src []byte, srcPitch int, dstFormat PixelFormat, dst []byte, dstPitch int) error {
	if len(src) < height*srcPitch || len(dst) < height*srcPitch {
		return sdlError("ConvertPixels out of bounds")
	}
	if !C.SDL_ConvertPixels((C.int)(width), (C.int)(height), (C.SDL_PixelFormat)(srcFormat), unsafe.Pointer(unsafe.SliceData(src)), (C.int)(srcPitch), (C.SDL_PixelFormat)(dstFormat), unsafe.Pointer(unsafe.SliceData(dst)), (C.int)(dstPitch)) {
		return getError()
	}
	return nil
}

// Copy a block of pixels of one format and colorspace to another format and
// colorspace.
//
// width: the width of the block to copy, in pixels.
//
// height: the height of the block to copy, in pixels.
//
// srcFormat: a [PixelFormat] value of the src pixels format.
//
// srcColorspace: a [Colorspace] value describing the colorspace of
// the src pixels.
//
// srcProperties: a [PropertiesID] with additional source color
// properties, or 0.
//
// src: a pointer to the source pixels.
//
// srcPitch: the pitch of the source pixels, in bytes.
//
// dstFormat: a [PixelFormat] value of the dst pixels format.
//
// dstColorspace: a [Colorspace] value describing the colorspace of
// the dst pixels.
//
// dstProperties: a [PropertiesID] with additional destination color
// properties, or 0.
//
// dst: a slice to be filled in with new pixel data.
//
// dstPitch: the pitch of the destination pixels, in bytes.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ConvertPixelsAndColorspace
func ConvertPixelsAndColorspace(width int, height int, srcFormat PixelFormat, srcColorspace Colorspace, srcProperties PropertiesID, src []byte, srcPitch int, dstFormat PixelFormat, dstColorspace Colorspace, dstProperties PropertiesID, dst []byte, dstPitch int) error {
	if len(src) < height*srcPitch || len(dst) < height*srcPitch {
		return sdlError("ConvertPixelsAndColorspace out of bounds")
	}
	if !C.SDL_ConvertPixelsAndColorspace((C.int)(width), (C.int)(height), (C.SDL_PixelFormat)(srcFormat), (C.SDL_Colorspace)(srcColorspace), (C.SDL_PropertiesID)(srcProperties), unsafe.Pointer(unsafe.SliceData(src)), (C.int)(srcPitch), (C.SDL_PixelFormat)(dstFormat), (C.SDL_Colorspace)(dstColorspace), (C.SDL_PropertiesID)(dstProperties), unsafe.Pointer(unsafe.SliceData(dst)), (C.int)(dstPitch)) {
		return getError()
	}
	return nil
}

// Premultiply the alpha on a block of pixels.
//
// This is safe to use with src == dst, but not for other overlapping areas.
//
// width: the width of the block to convert, in pixels.
//
// height: the height of the block to convert, in pixels.
//
// srcFormat: a [PixelFormat] value of the src pixels format.
//
// src: a pointer to the source pixels.
//
// srcPitch: the pitch of the source pixels, in bytes.
//
// dstFormat: a [PixelFormat] value of the dst pixels format.
//
// dst: a slice to be filled in with premultiplied pixel data.
//
// dstPitch: the pitch of the destination pixels, in bytes.
//
// linear: true to convert from sRGB to linear space for the alpha
// multiplication, false to do multiplication in sRGB space.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PremultiplyAlpha
func PremultiplyAlpha(width int, height int, srcFormat PixelFormat, src []byte, srcPitch int, dstFormat PixelFormat, dst []byte, dstPitch int, linear bool) error {
	if len(src) < height*srcPitch || len(dst) < height*srcPitch {
		return sdlError("PremultiplyAlpha out of bounds")
	}
	if !C.SDL_PremultiplyAlpha((C.int)(width), (C.int)(height), (C.SDL_PixelFormat)(srcFormat), unsafe.Pointer(unsafe.SliceData(src)), (C.int)(srcPitch), (C.SDL_PixelFormat)(dstFormat), unsafe.Pointer(unsafe.SliceData(dst)), (C.int)(dstPitch), (C.bool)(linear)) {
		return getError()
	}
	return nil
}

// Premultiply the alpha in a surface.
//
// This is safe to use with src == dst, but not for other overlapping areas.
//
// surface: the surface to modify.
//
// linear: true to convert from sRGB to linear space for the alpha
// multiplication, false to do multiplication in sRGB space.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_PremultiplySurfaceAlpha
func (surface *Surface) PremultiplyAlpha(linear bool) error {
	if !C.SDL_PremultiplySurfaceAlpha(surface.internal, (C.bool)(linear)) {
		return getError()
	}
	return nil
}

// Clear a surface with a specific color, with floating point precision.
//
// This function handles all surface formats, and ignores any clip rectangle.
//
// If the surface is YUV, the color is assumed to be in the sRGB colorspace,
// otherwise the color is assumed to be in the colorspace of the suface.
//
// surface: the [Surface] to clear.
//
// r: the red component of the pixel, normally in the range 0-1.
//
// g: the green component of the pixel, normally in the range 0-1.
//
// b: the blue component of the pixel, normally in the range 0-1.
//
// a: the alpha component of the pixel, normally in the range 0-1.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ClearSurface
func (surface *Surface) Clear(r float32, g float32, b float32, a float32) error {
	if !C.SDL_ClearSurface(surface.internal, (C.float)(r), (C.float)(g), (C.float)(b), (C.float)(a)) {
		return getError()
	}
	return nil
}

// Perform a fast fill of a rectangle with a specific color.
//
// color should be a pixel of the format used by the surface, and can be
// generated by [PixelFormatDetails.MapRGB] or [PixelFormatDetails.MapRGBA].
// If the color value contains an alpha component then the destination is
// simply filled with that alpha information, no blending takes place.
//
// If there is a clip rectangle set on the destination (set via
// [Surface.SetClipRect]), then this function will fill based on the
// intersection of the clip rectangle and rec`.
//
// surface: the [Surface] structure that is the drawing target.
//
// rect: the [Rect] structure representing the rectangle to fill, or
// nil to fill the entire surface.
//
// color: the color to fill with.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_FillSurfaceRect
func (surface *Surface) FillRect(rect *Rect, color uint32) error {
	var crect *C.SDL_Rect
	if rect != nil {
		crect = &C.SDL_Rect{C.int(rect.X), C.int(rect.Y), C.int(rect.W), C.int(rect.H)}
	}
	if !C.SDL_FillSurfaceRect(surface.internal, crect, (C.Uint32)(color)) {
		return getError()
	}
	return nil
}

// Perform a fast fill of a set of rectangles with a specific color.
//
// color should be a pixel of the format used by the surface, and can be
// generated by [PixelFormatDetails.MapRGB] or [PixelFormatDetails.MapRGBA].
// If the color value contains an alpha component then the destination is
// simply filled with that alpha information, no blending takes place.
//
// If there is a clip rectangle set on the destination (set via
// [Surface.SetClipRect]), then this function will fill based on the
// intersection of the clip rectangle and rect.
//
// surface: the [Surface] structure that is the drawing target.
//
// rects: an array of [Rect]s representing the rectangles to fill.
//
// color: the color to fill with.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_FillSurfaceRects
func (surface *Surface) FillRects(dst *Surface, rects []Rect, color uint32) error {
	crects := make([]C.SDL_Rect, len(rects))
	for i, r := range rects {
		crects[i] = C.SDL_Rect{C.int(r.X), C.int(r.Y), C.int(r.W), C.int(r.H)}
	}
	if !C.SDL_FillSurfaceRects(surface.internal, unsafe.SliceData(crects), (C.int)(len(rects)), (C.Uint32)(color)) {
		return getError()
	}
	return nil
}

// Performs a fast blit from the source surface to the destination surface
// with clipping.
//
// If either srcrect or dstrect are nil, the entire surface (src or
// dst) is copied while ensuring clipping to the clip rect of dst.
//
// The final blit rectangles are saved in srcrect and dstrect after all
// clipping is performed.
//
// The blit function should not be called on a locked surface.
//
// The blit semantics for surfaces with and without blending and colorkey are
// defined as follows:
//
//	RGBA->RGB:
//	   Source surface blend mode set to SDL_BLENDMODE_BLEND:
//	    alpha-blend (using the source alpha-channel and per-surface alpha)
//	    SDL_SRCCOLORKEY ignored.
//	  Source surface blend mode set to SDL_BLENDMODE_NONE:
//	    copy RGB.
//	    if SDL_SRCCOLORKEY set, only copy the pixels that do not match the
//	    RGB values of the source color key, ignoring alpha in the
//	    comparison.
//
//	RGB->RGBA:
//	  Source surface blend mode set to SDL_BLENDMODE_BLEND:
//	    alpha-blend (using the source per-surface alpha)
//	  Source surface blend mode set to SDL_BLENDMODE_NONE:
//	    copy RGB, set destination alpha to source per-surface alpha value.
//	  both:
//	    if SDL_SRCCOLORKEY set, only copy the pixels that do not match the
//	    source color key.
//
//	RGBA->RGBA:
//	  Source surface blend mode set to SDL_BLENDMODE_BLEND:
//	    alpha-blend (using the source alpha-channel and per-surface alpha)
//	    SDL_SRCCOLORKEY ignored.
//	  Source surface blend mode set to SDL_BLENDMODE_NONE:
//	    copy all of RGBA to the destination.
//	    if SDL_SRCCOLORKEY set, only copy the pixels that do not match the
//	    RGB values of the source color key, ignoring alpha in the
//	    comparison.
//
//	RGB->RGB:
//	  Source surface blend mode set to SDL_BLENDMODE_BLEND:
//	    alpha-blend (using the source per-surface alpha)
//	  Source surface blend mode set to SDL_BLENDMODE_NONE:
//	    copy RGB.
//	  both:
//	    if SDL_SRCCOLORKEY set, only copy the pixels that do not match the
//	    source color key.
//
// src: the [Surface] structure to be copied from.
//
// srcrect: the [Rect] structure representing the rectangle to be
// copied, or nil to copy the entire surface.
//
// surface: the [Surface] structure that is the blit target.
//
// dstrect: the [Rect] structure representing the x and y position in
// the destination surface, or nil for (0,0). The width and
// height are ignored, and are copied from srcrect. If you
// want a specific width and height, you should use
// [BlitSurfaceScaled].
//
// Returns nil on success or an error on failure.
//
// The same destination surface should not be used from two
// threads at once. It is safe to use the same source surface
// from multiple threads.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BlitSurface
func BlitSurface(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect) error {
	var srcr, dstr *C.SDL_Rect
	if srcrect != nil {
		srcr = &C.SDL_Rect{C.int(srcrect.X), C.int(srcrect.Y), C.int(srcrect.W), C.int(srcrect.H)}
	}
	if srcrect != nil {
		dstr = &C.SDL_Rect{C.int(dstrect.X), C.int(dstrect.Y), C.int(dstrect.W), C.int(dstrect.H)}
	}
	if !C.SDL_BlitSurface(src.internal, srcr, dst.internal, dstr) {
		return getError()
	}
	return nil
}

// Perform low-level surface blitting only.
//
// This is a semi-private blit function and it performs low-level surface
// blitting, assuming the input rectangles have already been clipped.
//
// src: the [Surface] structure to be copied from.
//
// srcrect: the [Rect] structure representing the rectangle to be
// copied.
//
// surface: the [Surface] structure that is the blit target.
//
// dstrect: the [Rect] structure representing the target rectangle in
// the destination surface.
//
// Returns nil on success or an error on failure.
//
// The same destination surface should not be used from two
// threads at once. It is safe to use the same source surface
// from multiple threads.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceUnchecked
func BlitSurfaceUnchecked(src *Surface, srcrect Rect, dst *Surface, dstrect Rect) error {
	if !C.SDL_BlitSurfaceUnchecked(src.internal,
		&C.SDL_Rect{C.int(srcrect.X), C.int(srcrect.Y), C.int(srcrect.W), C.int(srcrect.H)}, dst.internal,
		&C.SDL_Rect{C.int(dstrect.X), C.int(dstrect.Y), C.int(dstrect.W), C.int(dstrect.H)}) {
		return getError()
	}
	return nil
}

// Perform a scaled blit to a destination surface, which may be of a different
// format.
//
// src: the [Surface] structure to be copied from.
//
// srcrect: the [Rect] structure representing the rectangle to be
// copied, or nil to copy the entire surface.
//
// surface: the [Surface] structure that is the blit target.
//
// dstrect: the [Rect] structure representing the target rectangle in
// the destination surface, or nil to fill the entire
// destination surface.
//
// scaleMode: the [ScaleMode] to be used.
//
// Returns nil on success or an error on failure.
//
// The same destination surface should not be used from two
// threads at once. It is safe to use the same source surface
// from multiple threads.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceScaled
func BlitSurfaceScaled(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect, scaleMode ScaleMode) error {
	var srcr, dstr *C.SDL_Rect
	if srcrect != nil {
		srcr = &C.SDL_Rect{C.int(srcrect.X), C.int(srcrect.Y), C.int(srcrect.W), C.int(srcrect.H)}
	}
	if srcrect != nil {
		dstr = &C.SDL_Rect{C.int(dstrect.X), C.int(dstrect.Y), C.int(dstrect.W), C.int(dstrect.H)}
	}
	if !C.SDL_BlitSurfaceScaled(src.internal, srcr, dst.internal, dstr, (C.SDL_ScaleMode)(scaleMode)) {
		return getError()
	}
	return nil
}

// Perform low-level surface scaled blitting only.
//
// This is a semi-private function and it performs low-level surface blitting,
// assuming the input rectangles have already been clipped.
//
// src: the [Surface] structure to be copied from.
//
// srcrect: the [Rect] structure representing the rectangle to be copied.
//
// surface: the [Surface] structure that is the blit target.
//
// dstrect: the [Rect] structure representing the target rectangle in
// the destination surface.
//
// scaleMode: the [ScaleMode] to be used.
//
// Returns nil on success or an error on failure.
//
// The same destination surface should not be used from two
// threads at once. It is safe to use the same source surface
// from multiple threads.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceUncheckedScaled
func BlitSurfaceUncheckedScaled(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect, scaleMode ScaleMode) error {
	if !C.SDL_BlitSurfaceUncheckedScaled(src.internal,
		&C.SDL_Rect{C.int(srcrect.X), C.int(srcrect.Y), C.int(srcrect.W), C.int(srcrect.H)}, dst.internal,
		&C.SDL_Rect{C.int(dstrect.X), C.int(dstrect.Y), C.int(dstrect.W), C.int(dstrect.H)}, (C.SDL_ScaleMode)(scaleMode)) {
		return getError()
	}
	return nil
}

// Perform a tiled blit to a destination surface, which may be of a different
// format.
//
// The pixels in srcrect will be repeated as many times as needed to
// completely fill dstrect.
//
// src: the [Surface] structure to be copied from.
//
// srcrect: the [Rect] structure representing the rectangle to be
// copied, or nil to copy the entire surface.
//
// surface: the [Surface] structure that is the blit target.
//
// dstrect: the [Rect] structure representing the target rectangle in
// the destination surface, or nil to fill the entire surface.
//
// Returns nil on success or an error on failure.
//
// The same destination surface should not be used from two
// threads at once. It is safe to use the same source surface
// from multiple threads.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceTiled
func BlitSurfaceTiled(src *Surface, srcrect *Rect, dst *Surface, dstrect *Rect) error {
	var srcr, dstr *C.SDL_Rect
	if srcrect != nil {
		srcr = &C.SDL_Rect{C.int(srcrect.X), C.int(srcrect.Y), C.int(srcrect.W), C.int(srcrect.H)}
	}
	if srcrect != nil {
		dstr = &C.SDL_Rect{C.int(dstrect.X), C.int(dstrect.Y), C.int(dstrect.W), C.int(dstrect.H)}
	}
	if !C.SDL_BlitSurfaceTiled(src.internal, srcr, dst.internal, dstr) {
		return getError()
	}
	return nil
}

// Perform a scaled and tiled blit to a destination surface, which may be of a
// different format.
//
// The pixels in srcrect will be scaled and repeated as many times as needed
// to completely fill dstrect.
//
// src: the [Surface] structure to be copied from.
//
// srcrect: the [Rect] structure representing the rectangle to be
// copied, or nil to copy the entire surface.
//
// scale: the scale used to transform srcrect into the destination
// rectangle, e.g. a 32x32 texture with a scale of 2 would fill
// 64x64 tiles.
//
// scaleMode: scale algorithm to be used.
//
// surface: the [Surface] structure that is the blit target.
//
// dstrect: the [Rect] structure representing the target rectangle in
// the destination surface, or nil to fill the entire surface.
//
// Returns nil on success or an error on failure.
//
// The same destination surface should not be used from two
// threads at once. It is safe to use the same source surface
// from multiple threads.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BlitSurfaceTiledWithScale
func BlitSurfaceTiledWithScale(src *Surface, srcrect *Rect, scale float32, scaleMode ScaleMode, dst *Surface, dstrect *Rect) error {
	var srcr, dstr *C.SDL_Rect
	if srcrect != nil {
		srcr = &C.SDL_Rect{C.int(srcrect.X), C.int(srcrect.Y), C.int(srcrect.W), C.int(srcrect.H)}
	}
	if srcrect != nil {
		dstr = &C.SDL_Rect{C.int(dstrect.X), C.int(dstrect.Y), C.int(dstrect.W), C.int(dstrect.H)}
	}
	if !C.SDL_BlitSurfaceTiledWithScale(src.internal, srcr, (C.float)(scale), (C.SDL_ScaleMode)(scaleMode), dst.internal, dstr) {
		return getError()
	}
	return nil
}

// Perform a scaled blit using the 9-grid algorithm to a destination surface,
// which may be of a different format.
//
// The pixels in the source surface are split into a 3x3 grid, using the
// different corner sizes for each corner, and the sides and center making up
// the remaining pixels. The corners are then scaled using scale and fit
// into the corners of the destination rectangle. The sides and center are
// then stretched into place to cover the remaining destination rectangle.
//
// src: the [Surface] structure to be copied from.
//
// srcrect: the [Rect] structure representing the rectangle to be used
// for the 9-grid, or nil to use the entire surface.
//
// leftWidth: the width, in pixels, of the left corners in srcrect.
//
// rightWidth: the width, in pixels, of the right corners in srcrect.
//
// topHeight: the height, in pixels, of the top corners in srcrect.
//
// bottomHeight: the height, in pixels, of the bottom corners in
// srcrect.
//
// scale: the scale used to transform the corner of srcrect into the
// corner of dstrect, or 0 for an unscaled blit.
//
// scaleMode: scale algorithm to be used.
//
// surface: the [Surface] structure that is the blit target.
//
// dstrect: the [Rect] structure representing the target rectangle in
// the destination surface, or nil to fill the entire surface.
//
// Returns nil on success or an error on failure.
//
// The same destination surface should not be used from two
// threads at once. It is safe to use the same source surface
// from multiple threads.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BlitSurface9Grid
func BlitSurface9Grid(src *Surface, srcrect *Rect, leftWidth int, rightWidth int, topHeight int, bottomHeight int, scale float32, scaleMode ScaleMode, dst *Surface, dstrect *Rect) error {
	var srcr, dstr *C.SDL_Rect
	if srcrect != nil {
		srcr = &C.SDL_Rect{C.int(srcrect.X), C.int(srcrect.Y), C.int(srcrect.W), C.int(srcrect.H)}
	}
	if srcrect != nil {
		dstr = &C.SDL_Rect{C.int(dstrect.X), C.int(dstrect.Y), C.int(dstrect.W), C.int(dstrect.H)}
	}
	if !C.SDL_BlitSurface9Grid(src.internal, srcr, (C.int)(leftWidth), (C.int)(rightWidth), (C.int)(topHeight), (C.int)(bottomHeight), (C.float)(scale), (C.SDL_ScaleMode)(scaleMode), dst.internal, dstr) {
		return getError()
	}
	return nil
}

// Map an RGB triple to an opaque pixel value for a surface.
//
// This function maps the RGB color value to the specified pixel format and
// returns the pixel value best approximating the given RGB color value for
// the given pixel format.
//
// If the surface has a palette, the index of the closest matching color in
// the palette will be returned.
//
// If the surface pixel format has an alpha component it will be returned as
// all 1 bits (fully opaque).
//
// If the pixel format bpp (color depth) is less than 32-bpp then the unused
// upper bits of the return value can safely be ignored (e.g., with a 16-bpp
// format the return value can be assigned to a uint16, and similarly a uint8
// for an 8-bpp format).
//
// surface: the surface to use for the pixel format and palette.
//
// r: the red component of the pixel in the range 0-255.
//
// g: the green component of the pixel in the range 0-255.
//
// b: the blue component of the pixel in the range 0-255.
//
// Returns a pixel value.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MapSurfaceRGB
func (surface *Surface) MapRGB(r byte, g byte, b byte) uint32 {
	return (uint32)(C.SDL_MapSurfaceRGB(surface.internal, (C.Uint8)(r), (C.Uint8)(g), (C.Uint8)(b)))
}

// Map an RGBA quadruple to a pixel value for a surface.
//
// This function maps the RGBA color value to the specified pixel format and
// returns the pixel value best approximating the given RGBA color value for
// the given pixel format.
//
// If the surface pixel format has no alpha component the alpha value will be
// ignored (as it will be in formats with a palette).
//
// If the surface has a palette, the index of the closest matching color in
// the palette will be returned.
//
// If the pixel format bpp (color depth) is less than 32-bpp then the unused
// upper bits of the return value can safely be ignored (e.g., with a 16-bpp
// format the return value can be assigned to a uint16, and similarly a uint8
// for an 8-bpp format).
//
// surface: the surface to use for the pixel format and palette.
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
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_MapSurfaceRGBA
func (surface *Surface) MapRGBA(r byte, g byte, b byte, a byte) uint32 {
	return (uint32)(C.SDL_MapSurfaceRGBA(surface.internal, (C.Uint8)(r), (C.Uint8)(g), (C.Uint8)(b), (C.Uint8)(a)))
}

// Retrieves a single pixel from a surface.
//
// This function prioritizes correctness over speed: it is suitable for unit
// tests, but is not intended for use in a game engine.
//
// Like [PixelFormatDetails.GetRGBA], this uses the entire 0..255 range when converting color
// components from pixel formats with less than 8 bits per RGB component.
//
// surface: the surface to read.
//
// x: the horizontal coordinate, 0 <= x < width.
//
// y: the vertical coordinate, 0 <= y < height.
//
// r: the red channel, 0-255.
//
// g: the green channel, 0-255.
//
// b: the blue channel, 0-255.
//
// a: the alpha channel, 0-255.
//
// Returns the color values or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReadSurfacePixel
func ReadSurfacePixel(surface *Surface, x int, y int) (r, g, b, a byte, err error) {
	if !C.SDL_ReadSurfacePixel(surface.internal, (C.int)(x), (C.int)(y), (*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b), (*C.Uint8)(&a)) {
		return 0, 0, 0, 0, getError()
	}
	return
}

// Retrieves a single pixel from a surface.
//
// This function prioritizes correctness over speed: it is suitable for unit
// tests, but is not intended for use in a game engine.
//
// surface: the surface to read.
//
// x: the horizontal coordinate, 0 <= x < width.
//
// y: the vertical coordinate, 0 <= y < height.
//
// r: the red channel, normally in the range 0-1.
//
// g: the green channel, normally in the range 0-1.
//
// b: the blue channel, normally in the range 0-1.
//
// a: the alpha channel, normally in the range 0-1.
//
// Returns the color values or an error.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ReadSurfacePixelFloat
func ReadSurfacePixelFloat(surface *Surface, x int, y int) (r, g, b, a float32, err error) {
	if !C.SDL_ReadSurfacePixelFloat(surface.internal, (C.int)(x), (C.int)(y), (*C.float)(&r), (*C.float)(&g), (*C.float)(&b), (*C.float)(&a)) {
		return 0, 0, 0, 0, getError()
	}
	return
}

// Writes a single pixel to a surface.
//
// This function prioritizes correctness over speed: it is suitable for unit
// tests, but is not intended for use in a game engine.
//
// Like [PixelFormatDetails.MapRGBA], this uses the entire 0..255 range when converting color
// components from pixel formats with less than 8 bits per RGB component.
//
// surface: the surface to write.
//
// x: the horizontal coordinate, 0 <= x < width.
//
// y: the vertical coordinate, 0 <= y < height.
//
// r: the red channel value, 0-255.
//
// g: the green channel value, 0-255.
//
// b: the blue channel value, 0-255.
//
// a: the alpha channel value, 0-255.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WriteSurfacePixel
func (surface *Surface) WritePixel(x int, y int, r byte, g byte, b byte, a byte) error {
	if !C.SDL_WriteSurfacePixel(surface.internal, (C.int)(x), (C.int)(y), (C.Uint8)(r), (C.Uint8)(g), (C.Uint8)(b), (C.Uint8)(a)) {
		return getError()
	}
	return nil
}

// Writes a single pixel to a surface.
//
// This function prioritizes correctness over speed: it is suitable for unit
// tests, but is not intended for use in a game engine.
//
// surface: the surface to write.
//
// x: the horizontal coordinate, 0 <= x < width.
//
// y: the vertical coordinate, 0 <= y < height.
//
// r: the red channel value, normally in the range 0-1.
//
// g: the green channel value, normally in the range 0-1.
//
// b: the blue channel value, normally in the range 0-1.
//
// a: the alpha channel value, normally in the range 0-1.
//
// Returns nil on success or an error on failure.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_WriteSurfacePixelFloat
func (surface *Surface) WritePixelFloat(x int, y int, r float32, g float32, b float32, a float32) error {
	if !C.SDL_WriteSurfacePixelFloat(surface.internal, (C.int)(x), (C.int)(y), (C.float)(r), (C.float)(g), (C.float)(b), (C.float)(a)) {
		return getError()
	}
	return nil
}
