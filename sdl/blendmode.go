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
// #cgo noescape SDL_ComposeCustomBlendMode
// #cgo nocallback SDL_ComposeCustomBlendMode
// #include <SDL3/SDL.h>
import "C"

// # CategoryBlendmode
//
// Blend modes decide how two colors will mix together. There are both
// standard modes for basic needs and a means to create custom modes,
// dictating what sort of math to do on what color components.

// A set of blend modes used in drawing operations.
//
// These predefined blend modes are supported everywhere.
//
// Additional values may be obtained from [ComposeCustomBlendMode].
//
// This datatype is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BlendMode
type BlendMode uint32

const (
	BlendmodeNone               BlendMode = 0x00000000 // no blending: dstRGBA = srcRGBA
	BlendmodeBlend              BlendMode = 0x00000001 // alpha blending: dstRGB = (srcRGB * srcA) + (dstRGB * (1-srcA)), dstA = srcA + (dstA * (1-srcA))
	BlendmodeBlendPremultiplied BlendMode = 0x00000010 // pre-multiplied alpha blending: dstRGBA = srcRGBA + (dstRGBA * (1-srcA))
	BlendmodeAdd                BlendMode = 0x00000002 // additive blending: dstRGB = (srcRGB * srcA) + dstRGB, dstA = dstA
	BlendmodeAddPremultiplied   BlendMode = 0x00000020 // pre-multiplied additive blending: dstRGB = srcRGB + dstRGB, dstA = dstA
	BlendmodeMod                BlendMode = 0x00000004 // color modulate: dstRGB = srcRGB * dstRGB, dstA = dstA
	BlendmodeMul                BlendMode = 0x00000008 // color multiply: dstRGB = (srcRGB * dstRGB) + (dstRGB * (1-srcA)), dstA = dstA
	BlendmodeInvalid            BlendMode = 0x7FFFFFFF
)

// The blend operation used when combining source and destination pixel
// components.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BlendOperation
type BlendOperation uint32

const (
	BlendoperationAdd         BlendOperation = 0x1 // dst + src: supported by all renderers
	BlendoperationSubtract    BlendOperation = 0x2 // src - dst : supported by D3D, OpenGL, OpenGLES, and Vulkan
	BlendoperationRevSubtract BlendOperation = 0x3 // dst - src : supported by D3D, OpenGL, OpenGLES, and Vulkan
	BlendoperationMinimum     BlendOperation = 0x4 // min(dst, src) : supported by D3D, OpenGL, OpenGLES, and Vulkan
	BlendoperationMaximum     BlendOperation = 0x5 // max(dst, src) : supported by D3D, OpenGL, OpenGLES, and Vulkan
)

// The normalized factor used to multiply pixel components.
//
// The blend factors are multiplied with the pixels from a drawing operation
// (src) and the pixels from the render target (dst) before the blend
// operation. The comma-separated factors listed above are always applied in
// the component order red, green, blue, and alpha.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_BlendFactor
type BlendFactor uint32

const (
	BlendfactorZero             BlendFactor = 0x1 // 0, 0, 0, 0
	BlendfactorOne              BlendFactor = 0x2 // 1, 1, 1, 1
	BlendfactorSrcColor         BlendFactor = 0x3 // srcR, srcG, srcB, srcA
	BlendfactorOneMinusSrcColor BlendFactor = 0x4 // 1-srcR, 1-srcG, 1-srcB, 1-srcA
	BlendfactorSrcAlpha         BlendFactor = 0x5 // srcA, srcA, srcA, srcA
	BlendfactorOneMinusSrcAlpha BlendFactor = 0x6 // 1-srcA, 1-srcA, 1-srcA, 1-srcA
	BlendfactorDstColor         BlendFactor = 0x7 // dstR, dstG, dstB, dstA
	BlendfactorOneMinusDstColor BlendFactor = 0x8 // 1-dstR, 1-dstG, 1-dstB, 1-dstA
	BlendfactorDstAlpha         BlendFactor = 0x9 // dstA, dstA, dstA, dstA
	BlendfactorOneMinusDstAlpha BlendFactor = 0xA // 1-dstA, 1-dstA, 1-dstA, 1-dstA
)

// Compose a custom blend mode for renderers.
//
// The functions [Renderer.SetDrawBlendMode] and [Texture.SetBlendMode] accept
// the [BlendMode] returned by this function if the renderer supports it.
//
// A blend mode controls how the pixels from a drawing operation (source) get
// combined with the pixels from the render target (destination). First, the
// components of the source and destination pixels get multiplied with their
// blend factors. Then, the blend operation takes the two products and
// calculates the result that will get stored in the render target.
//
// Expressed in pseudocode, it would look like this:
//
//	dstRGB = colorOperation(srcRGB * srcColorFactor, dstRGB * dstColorFactor);
//	dstA = alphaOperation(srcA * srcAlphaFactor, dstA * dstAlphaFactor);
//
// Where the functions `colorOperation(src, dst)` and `alphaOperation(src,
// dst)` can return one of the following:
//
// - `src + dst`
// - `src - dst`
// - `dst - src`
// - `min(src, dst)`
// - `max(src, dst)`
//
// The red, green, and blue components are always multiplied with the first,
// second, and third components of the [BlendFactor], respectively. The
// fourth component is not used.
//
// The alpha component is always multiplied with the fourth component of the
// [BlendFactor]. The other components are not used in the alpha
// calculation.
//
// Support for these blend modes varies for each renderer. To check if a
// specific [BlendMode] is supported, create a renderer and pass it to
// either [Renderer.SetDrawBlendMode] or [Texture.SetBlendMode]. They will
// return with an error if the blend mode is not supported.
//
// This list describes the support of custom blend modes for each renderer.
// All renderers support the four blend modes listed in the [BlendMode]
// enumeration.
//
// - **direct3d**: Supports all operations with all factors. However, some
// factors produce unexpected results with [BlendoperationMinimum] and
// [BlendoperationMaximum].
// - **direct3d11**: Same as Direct3D 9.
// - **opengl**: Supports the [BlendmodeAdd] operation with all
// factors. OpenGL versions 1.1, 1.2, and 1.3 do not work correctly here.
// - **opengles2**: Supports the [BlendmodeAdd],
// [BlendoperationSubtract], [BlendoperationRevSubtract]
// operations with all factors.
// - **psp**: No custom blend mode support.
// - **software**: No custom blend mode support.
//
// Some renderers do not provide an alpha component for the default render
// target. The [BlendfactorDstAlpha] and
// [BlendfactorOneMinusDstAlpha] factors do not have an effect in this
// case.
//
// srcColorFactor: the [BlendFactor] applied to the red, green, and
// blue components of the source pixels.
//
// dstColorFactor: the [BlendFactor] applied to the red, green, and
// blue components of the destination pixels.
//
// colorOperation: the [BlendOperation] used to combine the red,
// green, and blue components of the source and
// destination pixels.
//
// srcAlphaFactor: the [BlendFactor] applied to the alpha component of
// the source pixels.
//
// dstAlphaFactor: the [BlendFactor] applied to the alpha component of
// the destination pixels.
//
// alphaOperation: the [BlendOperation] used to combine the alpha
// component of the source and destination pixels.
//
// Returns an [BlendMode] that represents the chosen factors and
// operations.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ComposeCustomBlendMode
func ComposeCustomBlendMode(srcColorFactor BlendFactor, dstColorFactor BlendFactor, colorOperation BlendOperation, srcAlphaFactor BlendFactor, dstAlphaFactor BlendFactor, alphaOperation BlendOperation) BlendMode {
	return (BlendMode)(C.SDL_ComposeCustomBlendMode((C.SDL_BlendFactor)(srcColorFactor), (C.SDL_BlendFactor)(dstColorFactor), (C.SDL_BlendOperation)(colorOperation), (C.SDL_BlendFactor)(srcAlphaFactor), (C.SDL_BlendFactor)(dstAlphaFactor), (C.SDL_BlendOperation)(alphaOperation)))
}
