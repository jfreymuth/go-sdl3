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
// #cgo noescape SDL_GetNumRenderDrivers
// #cgo nocallback SDL_GetNumRenderDrivers
// #cgo noescape SDL_GetRenderDriver
// #cgo nocallback SDL_GetRenderDriver
// #cgo noescape SDL_CreateWindowAndRenderer
// #cgo nocallback SDL_CreateWindowAndRenderer
// #cgo noescape SDL_CreateRenderer
// #cgo nocallback SDL_CreateRenderer
// #cgo noescape SDL_CreateRendererWithProperties
// #cgo nocallback SDL_CreateRendererWithProperties
// #cgo noescape SDL_CreateSoftwareRenderer
// #cgo nocallback SDL_CreateSoftwareRenderer
// #cgo noescape SDL_GetRenderer
// #cgo nocallback SDL_GetRenderer
// #cgo noescape SDL_GetRenderWindow
// #cgo nocallback SDL_GetRenderWindow
// #cgo noescape SDL_GetRendererName
// #cgo nocallback SDL_GetRendererName
// #cgo noescape SDL_GetRendererProperties
// #cgo nocallback SDL_GetRendererProperties
// #cgo noescape SDL_GetRenderOutputSize
// #cgo nocallback SDL_GetRenderOutputSize
// #cgo noescape SDL_GetCurrentRenderOutputSize
// #cgo nocallback SDL_GetCurrentRenderOutputSize
// #cgo noescape SDL_CreateTexture
// #cgo nocallback SDL_CreateTexture
// #cgo noescape SDL_CreateTextureFromSurface
// #cgo nocallback SDL_CreateTextureFromSurface
// #cgo noescape SDL_CreateTextureWithProperties
// #cgo nocallback SDL_CreateTextureWithProperties
// #cgo noescape SDL_GetTextureProperties
// #cgo nocallback SDL_GetTextureProperties
// #cgo noescape SDL_GetRendererFromTexture
// #cgo nocallback SDL_GetRendererFromTexture
// #cgo noescape SDL_GetTextureSize
// #cgo nocallback SDL_GetTextureSize
// #cgo noescape SDL_SetTextureColorMod
// #cgo nocallback SDL_SetTextureColorMod
// #cgo noescape SDL_SetTextureColorModFloat
// #cgo nocallback SDL_SetTextureColorModFloat
// #cgo noescape SDL_GetTextureColorMod
// #cgo nocallback SDL_GetTextureColorMod
// #cgo noescape SDL_GetTextureColorModFloat
// #cgo nocallback SDL_GetTextureColorModFloat
// #cgo noescape SDL_SetTextureAlphaMod
// #cgo nocallback SDL_SetTextureAlphaMod
// #cgo noescape SDL_SetTextureAlphaModFloat
// #cgo nocallback SDL_SetTextureAlphaModFloat
// #cgo noescape SDL_GetTextureAlphaMod
// #cgo nocallback SDL_GetTextureAlphaMod
// #cgo noescape SDL_GetTextureAlphaModFloat
// #cgo nocallback SDL_GetTextureAlphaModFloat
// #cgo noescape SDL_SetTextureBlendMode
// #cgo nocallback SDL_SetTextureBlendMode
// #cgo noescape SDL_GetTextureBlendMode
// #cgo nocallback SDL_GetTextureBlendMode
// #cgo noescape SDL_SetTextureScaleMode
// #cgo nocallback SDL_SetTextureScaleMode
// #cgo noescape SDL_GetTextureScaleMode
// #cgo nocallback SDL_GetTextureScaleMode
// #cgo noescape SDL_UpdateTexture
// #cgo nocallback SDL_UpdateTexture
// #cgo noescape SDL_UpdateYUVTexture
// #cgo nocallback SDL_UpdateYUVTexture
// #cgo noescape SDL_UpdateNVTexture
// #cgo nocallback SDL_UpdateNVTexture
// #cgo noescape SDL_LockTexture
// #cgo nocallback SDL_LockTexture
// #cgo noescape SDL_LockTextureToSurface
// #cgo nocallback SDL_LockTextureToSurface
// #cgo noescape SDL_UnlockTexture
// #cgo nocallback SDL_UnlockTexture
// #cgo noescape SDL_SetRenderTarget
// #cgo nocallback SDL_SetRenderTarget
// #cgo noescape SDL_GetRenderTarget
// #cgo nocallback SDL_GetRenderTarget
// #cgo noescape SDL_SetRenderLogicalPresentation
// #cgo nocallback SDL_SetRenderLogicalPresentation
// #cgo noescape SDL_GetRenderLogicalPresentation
// #cgo nocallback SDL_GetRenderLogicalPresentation
// #cgo noescape SDL_GetRenderLogicalPresentationRect
// #cgo nocallback SDL_GetRenderLogicalPresentationRect
// #cgo noescape SDL_RenderCoordinatesFromWindow
// #cgo nocallback SDL_RenderCoordinatesFromWindow
// #cgo noescape SDL_RenderCoordinatesToWindow
// #cgo nocallback SDL_RenderCoordinatesToWindow
// #cgo noescape SDL_ConvertEventToRenderCoordinates
// #cgo nocallback SDL_ConvertEventToRenderCoordinates
// #cgo noescape SDL_SetRenderViewport
// #cgo nocallback SDL_SetRenderViewport
// #cgo noescape SDL_GetRenderViewport
// #cgo nocallback SDL_GetRenderViewport
// #cgo noescape SDL_RenderViewportSet
// #cgo nocallback SDL_RenderViewportSet
// #cgo noescape SDL_GetRenderSafeArea
// #cgo nocallback SDL_GetRenderSafeArea
// #cgo noescape SDL_SetRenderClipRect
// #cgo nocallback SDL_SetRenderClipRect
// #cgo noescape SDL_GetRenderClipRect
// #cgo nocallback SDL_GetRenderClipRect
// #cgo noescape SDL_RenderClipEnabled
// #cgo nocallback SDL_RenderClipEnabled
// #cgo noescape SDL_SetRenderScale
// #cgo nocallback SDL_SetRenderScale
// #cgo noescape SDL_GetRenderScale
// #cgo nocallback SDL_GetRenderScale
// #cgo noescape SDL_SetRenderDrawColor
// #cgo nocallback SDL_SetRenderDrawColor
// #cgo noescape SDL_SetRenderDrawColorFloat
// #cgo nocallback SDL_SetRenderDrawColorFloat
// #cgo noescape SDL_GetRenderDrawColor
// #cgo nocallback SDL_GetRenderDrawColor
// #cgo noescape SDL_GetRenderDrawColorFloat
// #cgo nocallback SDL_GetRenderDrawColorFloat
// #cgo noescape SDL_SetRenderColorScale
// #cgo nocallback SDL_SetRenderColorScale
// #cgo noescape SDL_GetRenderColorScale
// #cgo nocallback SDL_GetRenderColorScale
// #cgo noescape SDL_SetRenderDrawBlendMode
// #cgo nocallback SDL_SetRenderDrawBlendMode
// #cgo noescape SDL_GetRenderDrawBlendMode
// #cgo nocallback SDL_GetRenderDrawBlendMode
// #cgo noescape SDL_RenderClear
// #cgo nocallback SDL_RenderClear
// #cgo noescape SDL_RenderPoint
// #cgo nocallback SDL_RenderPoint
// #cgo noescape SDL_RenderPoints
// #cgo nocallback SDL_RenderPoints
// #cgo noescape SDL_RenderLine
// #cgo nocallback SDL_RenderLine
// #cgo noescape SDL_RenderLines
// #cgo nocallback SDL_RenderLines
// #cgo noescape SDL_RenderRect
// #cgo nocallback SDL_RenderRect
// #cgo noescape SDL_RenderRects
// #cgo nocallback SDL_RenderRects
// #cgo noescape SDL_RenderFillRect
// #cgo nocallback SDL_RenderFillRect
// #cgo noescape SDL_RenderFillRects
// #cgo nocallback SDL_RenderFillRects
// #cgo noescape SDL_RenderTexture
// #cgo nocallback SDL_RenderTexture
// #cgo noescape SDL_RenderTextureRotated
// #cgo nocallback SDL_RenderTextureRotated
// #cgo noescape SDL_RenderTextureAffine
// #cgo nocallback SDL_RenderTextureAffine
// #cgo noescape SDL_RenderTextureTiled
// #cgo nocallback SDL_RenderTextureTiled
// #cgo noescape SDL_RenderTexture9Grid
// #cgo nocallback SDL_RenderTexture9Grid
// #cgo noescape SDL_RenderGeometry
// #cgo nocallback SDL_RenderGeometry
// #cgo noescape SDL_RenderGeometryRaw
// #cgo nocallback SDL_RenderGeometryRaw
// #cgo noescape SDL_RenderReadPixels
// #cgo nocallback SDL_RenderReadPixels
// #cgo noescape SDL_RenderPresent
// #cgo nocallback SDL_RenderPresent
// #cgo noescape SDL_DestroyTexture
// #cgo noescape SDL_DestroyRenderer
// #cgo noescape SDL_FlushRenderer
// #cgo nocallback SDL_FlushRenderer
// #cgo noescape SDL_GetRenderMetalLayer
// #cgo nocallback SDL_GetRenderMetalLayer
// #cgo noescape SDL_GetRenderMetalCommandEncoder
// #cgo nocallback SDL_GetRenderMetalCommandEncoder
// #cgo noescape SDL_AddVulkanRenderSemaphores
// #cgo nocallback SDL_AddVulkanRenderSemaphores
// #cgo noescape SDL_SetRenderVSync
// #cgo nocallback SDL_SetRenderVSync
// #cgo noescape SDL_GetRenderVSync
// #cgo nocallback SDL_GetRenderVSync
// #cgo noescape SDL_RenderDebugText
// #cgo nocallback SDL_RenderDebugText
// #include <SDL3/SDL.h>
import "C"
import "unsafe"

// # CategoryRender
//
// Header file for SDL 2D rendering functions.
//
// This API supports the following features:
//
// - single pixel points
// - single pixel lines
// - filled rectangles
// - texture images
// - 2D polygons
//
// The primitives may be drawn in opaque, blended, or additive modes.
//
// The texture images may be drawn in opaque, blended, or additive modes. They
// can have an additional color tint or alpha modulation applied to them, and
// may also be stretched with linear interpolation.
//
// This API is designed to accelerate simple 2D operations. You may want more
// functionality such as polygons and particle effects and in that case you
// should use SDL's OpenGL/Direct3D support, the SDL3 GPU API, or one of the
// many good 3D engines.
//
// These functions must be called from the main thread. See this bug for
// details: https://github.com/libsdl-org/SDL/issues/986

// The name of the software renderer.
//
// This macro is available since SDL 3.2.0.
const SoftwareRenderer = "software"

// Vertex structure.
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Vertex
type Vertex struct {
	Position FPoint // Vertex position, in SDL_Renderer coordinates
	Color    FColor // Vertex color
	TexCoord FPoint // Normalized texture coordinates, if needed
}

// The access pattern allowed for a texture.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_TextureAccess
type TextureAccess uint32

const (
	TextureaccessStatic    TextureAccess = iota // Changes rarely, not lockable
	TextureaccessStreaming                      // Changes frequently, lockable
	TextureaccessTarget                         // Texture can be used as a render target
)

// How the logical size is mapped to the output.
//
// This enum is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RendererLogicalPresentation
type RendererLogicalPresentation uint32

const (
	LogicalPresentationDisabled     RendererLogicalPresentation = iota // There is no logical size in effect
	LogicalPresentationStretch                                         // The rendered content is stretched to the output resolution
	LogicalPresentationLetterbox                                       // The rendered content is fit to the largest dimension and the other dimension is letterboxed with black bars
	LogicalPresentationOverscan                                        // The rendered content is fit to the smallest dimension and the other dimension extends beyond the output bounds
	LogicalPresentationIntegerScale                                    // The rendered content is scaled up by integer multiples to fit the output resolution
)

// A structure representing rendering state
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Renderer
type Renderer C.struct_SDL_Renderer

// An efficient driver-specific representation of pixel data
//
// This struct is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_Texture
type Texture struct {
	internal *C.struct_SDL_Texture
}

// The format of the texture
func (texture *Texture) Format() PixelFormat { return PixelFormat(texture.internal.format) }

// The width of the texture
func (texture *Texture) Width() int { return int(texture.internal.w) }

// The height of the texture
func (texture *Texture) Height() int { return int(texture.internal.h) }

// Get the number of 2D rendering drivers available for the current display.
//
// A render driver is a set of code that handles rendering and texture
// management on a particular display. Normally there is only one, but some
// drivers may have several available with different capabilities.
//
// There may be none if SDL was compiled without render support.
//
// Returns the number of built in render drivers.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetNumRenderDrivers
func GetNumRenderDrivers() int {
	return (int)(C.SDL_GetNumRenderDrivers())
}

// Use this function to get the name of a built in 2D rendering driver.
//
// The list of rendering drivers is given in the order that they are normally
// initialized by default; the drivers that seem more reasonable to choose
// first (as far as the SDL developers believe) are earlier in the list.
//
// The names of drivers are all simple, low-ASCII identifiers, like "opengl",
// "direct3d12" or "metal". These never have Unicode characters, and are not
// meant to be proper names.
//
// index: the index of the rendering driver; the value ranges from 0 to
// SDL_GetNumRenderDrivers() - 1.
//
// Returns the name of the rendering driver at the requested index, or NULL
// if an invalid index was specified.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRenderDriver
func GetRenderDriver(index int) string {
	return C.GoString(C.SDL_GetRenderDriver((C.int)(index)))
}

// Create a window and default renderer.
//
// title: the title of the window, in UTF-8 encoding.
//
// width: the width of the window.
//
// height: the height of the window.
//
// window_flags: the flags used to create the window (see
// SDL_CreateWindow()).
//
// window: a pointer filled with the window, or NULL on error.
//
// renderer: a pointer filled with the renderer, or NULL on error.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateWindowAndRenderer
func CreateWindowAndRenderer(title string, width int, height int, windowFlags WindowFlags) (*Window, *Renderer, error) {
	var window *C.SDL_Window
	var renderer *C.SDL_Renderer
	if !C.SDL_CreateWindowAndRenderer(tmpstring(title), (C.int)(width), (C.int)(height), (C.SDL_WindowFlags)(windowFlags), &window, &renderer) {
		return nil, nil, getError()
	}
	return (*Window)(window), (*Renderer)(renderer), nil
}

// Create a 2D rendering context for a window.
//
// If you want a specific renderer, you can specify its name here. A list of
// available renderers can be obtained by calling SDL_GetRenderDriver()
// multiple times, with indices from 0 to SDL_GetNumRenderDrivers()-1. If you
// don't need a specific renderer, specify NULL and SDL will attempt to choose
// the best option for you, based on what is available on the user's system.
//
// If `name` is a comma-separated list, SDL will try each name, in the order
// listed, until one succeeds or all of them fail.
//
// By default the rendering size matches the window size in pixels, but you
// can call SDL_SetRenderLogicalPresentation() to change the content size and
// scaling options.
//
// window: the window where rendering is displayed.
//
// name: the name of the rendering driver to initialize, or NULL to let
// SDL choose one.
//
// Returns a valid rendering context or NULL if there was an error; call
// SDL_GetError() for more information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateRenderer
func CreateRenderer(window *Window, name string) (*Renderer, error) {
	r := (*Renderer)(C.SDL_CreateRenderer((*C.SDL_Window)(window), tmpstring(name)))
	if r == nil {
		return nil, getError()
	}
	return r, nil
}

// Create a 2D rendering context for a window, with the specified properties.
//
// These are the supported properties:
//
// - `SDL_PROP_RENDERER_CREATE_NAME_STRING`: the name of the rendering driver
// to use, if a specific one is desired
// - `SDL_PROP_RENDERER_CREATE_WINDOW_POINTER`: the window where rendering is
// displayed, required if this isn't a software renderer using a surface
// - `SDL_PROP_RENDERER_CREATE_SURFACE_POINTER`: the surface where rendering
// is displayed, if you want a software renderer without a window
// - `SDL_PROP_RENDERER_CREATE_OUTPUT_COLORSPACE_NUMBER`: an SDL_Colorspace
// value describing the colorspace for output to the display, defaults to
// SDL_COLORSPACE_SRGB. The direct3d11, direct3d12, and metal renderers
// support SDL_COLORSPACE_SRGB_LINEAR, which is a linear color space and
// supports HDR output. If you select SDL_COLORSPACE_SRGB_LINEAR, drawing
// still uses the sRGB colorspace, but values can go beyond 1.0 and float
// (linear) format textures can be used for HDR content.
// - `SDL_PROP_RENDERER_CREATE_PRESENT_VSYNC_NUMBER`: non-zero if you want
// present synchronized with the refresh rate. This property can take any
// value that is supported by SDL_SetRenderVSync() for the renderer.
//
// With the vulkan renderer:
//
// - `SDL_PROP_RENDERER_CREATE_VULKAN_INSTANCE_POINTER`: the VkInstance to use
// with the renderer, optional.
// - `SDL_PROP_RENDERER_CREATE_VULKAN_SURFACE_NUMBER`: the VkSurfaceKHR to use
// with the renderer, optional.
// - `SDL_PROP_RENDERER_CREATE_VULKAN_PHYSICAL_DEVICE_POINTER`: the
// VkPhysicalDevice to use with the renderer, optional.
// - `SDL_PROP_RENDERER_CREATE_VULKAN_DEVICE_POINTER`: the VkDevice to use
// with the renderer, optional.
// - `SDL_PROP_RENDERER_CREATE_VULKAN_GRAPHICS_QUEUE_FAMILY_INDEX_NUMBER`: the
// queue family index used for rendering.
// - `SDL_PROP_RENDERER_CREATE_VULKAN_PRESENT_QUEUE_FAMILY_INDEX_NUMBER`: the
// queue family index used for presentation.
//
// props: the properties to use.
//
// Returns a valid rendering context or NULL if there was an error; call
// SDL_GetError() for more information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateRendererWithProperties
func CreateRendererWithProperties(props PropertiesID) (*Renderer, error) {
	r := (*Renderer)(C.SDL_CreateRendererWithProperties((C.SDL_PropertiesID)(props)))
	if r == nil {
		return nil, getError()
	}
	return r, nil
}

const PropRendererCreateNameString = "SDL.renderer.create.name"
const PropRendererCreateWindowPointer = "SDL.renderer.create.window"
const PropRendererCreateSurfacePointer = "SDL.renderer.create.surface"
const PropRendererCreateOutputColorspaceNumber = "SDL.renderer.create.output_colorspace"
const PropRendererCreatePresentVsyncNumber = "SDL.renderer.create.present_vsync"
const PropRendererCreateVulkanInstancePointer = "SDL.renderer.create.vulkan.instance"
const PropRendererCreateVulkanSurfaceNumber = "SDL.renderer.create.vulkan.surface"
const PropRendererCreateVulkanPhysicalDevicePointer = "SDL.renderer.create.vulkan.physical_device"
const PropRendererCreateVulkanDevicePointer = "SDL.renderer.create.vulkan.device"
const PropRendererCreateVulkanGraphicsQueueFamilyIndexNumber = "SDL.renderer.create.vulkan.graphics_queue_family_index"
const PropRendererCreateVulkanPresentQueueFamilyIndexNumber = "SDL.renderer.create.vulkan.present_queue_family_index"

// Create a 2D software rendering context for a surface.
//
// Two other API which can be used to create SDL_Renderer:
// SDL_CreateRenderer() and SDL_CreateWindowAndRenderer(). These can _also_
// create a software renderer, but they are intended to be used with an
// SDL_Window as the final destination and not an SDL_Surface.
//
// surface: the SDL_Surface structure representing the surface where
// rendering is done.
//
// Returns a valid rendering context or NULL if there was an error; call
// SDL_GetError() for more information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateSoftwareRenderer
func CreateSoftwareRenderer(surface *Surface) (*Renderer, error) {
	r := (*Renderer)(C.SDL_CreateSoftwareRenderer(surface.internal))
	if r == nil {
		return nil, getError()
	}
	return r, nil
}

// Get the renderer associated with a window.
//
// window: the window to query.
//
// Returns the rendering context on success or NULL on failure; call
// SDL_GetError() for more information.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRenderer
func (window *Window) Renderer() (*Renderer, error) {
	r := (*Renderer)(C.SDL_GetRenderer((*C.SDL_Window)(window)))
	if r == nil {
		return nil, getError()
	}
	return r, nil
}

// Get the window associated with a renderer.
//
// renderer: the renderer to query.
//
// Returns the window on success or NULL on failure; call SDL_GetError() for
// more information.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRenderWindow
func (renderer *Renderer) Window() (*Window, error) {
	w := (*Window)(C.SDL_GetRenderWindow((*C.SDL_Renderer)(renderer)))
	if w == nil {
		return nil, getError()
	}
	return w, nil
}

// Get the name of a renderer.
//
// renderer: the rendering context.
//
// Returns the name of the selected renderer, or NULL on failure; call
// SDL_GetError() for more information.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRendererName
func (renderer *Renderer) Name() (string, error) {
	name := C.SDL_GetRendererName((*C.SDL_Renderer)(renderer))
	if name == nil {
		return "", getError()
	}
	return C.GoString(name), nil
}

// Get the properties associated with a renderer.
//
// The following read-only properties are provided by SDL:
//
// - `SDL_PROP_RENDERER_NAME_STRING`: the name of the rendering driver
// - `SDL_PROP_RENDERER_WINDOW_POINTER`: the window where rendering is
// displayed, if any
// - `SDL_PROP_RENDERER_SURFACE_POINTER`: the surface where rendering is
// displayed, if this is a software renderer without a window
// - `SDL_PROP_RENDERER_VSYNC_NUMBER`: the current vsync setting
// - `SDL_PROP_RENDERER_MAX_TEXTURE_SIZE_NUMBER`: the maximum texture width
// and height
// - `SDL_PROP_RENDERER_TEXTURE_FORMATS_POINTER`: a (const SDL_PixelFormat *)
// array of pixel formats, terminated with SDL_PIXELFORMAT_UNKNOWN,
// representing the available texture formats for this renderer.
// - `SDL_PROP_RENDERER_OUTPUT_COLORSPACE_NUMBER`: an SDL_Colorspace value
// describing the colorspace for output to the display, defaults to
// SDL_COLORSPACE_SRGB.
// - `SDL_PROP_RENDERER_HDR_ENABLED_BOOLEAN`: true if the output colorspace is
// SDL_COLORSPACE_SRGB_LINEAR and the renderer is showing on a display with
// HDR enabled. This property can change dynamically when
// SDL_EVENT_WINDOW_HDR_STATE_CHANGED is sent.
// - `SDL_PROP_RENDERER_SDR_WHITE_POINT_FLOAT`: the value of SDR white in the
// SDL_COLORSPACE_SRGB_LINEAR colorspace. When HDR is enabled, this value is
// automatically multiplied into the color scale. This property can change
// dynamically when SDL_EVENT_WINDOW_HDR_STATE_CHANGED is sent.
// - `SDL_PROP_RENDERER_HDR_HEADROOM_FLOAT`: the additional high dynamic range
// that can be displayed, in terms of the SDR white point. When HDR is not
// enabled, this will be 1.0. This property can change dynamically when
// SDL_EVENT_WINDOW_HDR_STATE_CHANGED is sent.
//
// With the direct3d renderer:
//
// - `SDL_PROP_RENDERER_D3D9_DEVICE_POINTER`: the IDirect3DDevice9 associated
// with the renderer
//
// With the direct3d11 renderer:
//
// - `SDL_PROP_RENDERER_D3D11_DEVICE_POINTER`: the ID3D11Device associated
// with the renderer
// - `SDL_PROP_RENDERER_D3D11_SWAPCHAIN_POINTER`: the IDXGISwapChain1
// associated with the renderer. This may change when the window is resized.
//
// With the direct3d12 renderer:
//
// - `SDL_PROP_RENDERER_D3D12_DEVICE_POINTER`: the ID3D12Device associated
// with the renderer
// - `SDL_PROP_RENDERER_D3D12_SWAPCHAIN_POINTER`: the IDXGISwapChain4
// associated with the renderer.
// - `SDL_PROP_RENDERER_D3D12_COMMAND_QUEUE_POINTER`: the ID3D12CommandQueue
// associated with the renderer
//
// With the vulkan renderer:
//
// - `SDL_PROP_RENDERER_VULKAN_INSTANCE_POINTER`: the VkInstance associated
// with the renderer
// - `SDL_PROP_RENDERER_VULKAN_SURFACE_NUMBER`: the VkSurfaceKHR associated
// with the renderer
// - `SDL_PROP_RENDERER_VULKAN_PHYSICAL_DEVICE_POINTER`: the VkPhysicalDevice
// associated with the renderer
// - `SDL_PROP_RENDERER_VULKAN_DEVICE_POINTER`: the VkDevice associated with
// the renderer
// - `SDL_PROP_RENDERER_VULKAN_GRAPHICS_QUEUE_FAMILY_INDEX_NUMBER`: the queue
// family index used for rendering
// - `SDL_PROP_RENDERER_VULKAN_PRESENT_QUEUE_FAMILY_INDEX_NUMBER`: the queue
// family index used for presentation
// - `SDL_PROP_RENDERER_VULKAN_SWAPCHAIN_IMAGE_COUNT_NUMBER`: the number of
// swapchain images, or potential frames in flight, used by the Vulkan
// renderer
//
// With the gpu renderer:
//
// - `SDL_PROP_RENDERER_GPU_DEVICE_POINTER`: the SDL_GPUDevice associated with
// the renderer
//
// renderer: the rendering context.
//
// Returns a valid property ID on success or 0 on failure; call
// SDL_GetError() for more information.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRendererProperties
func (renderer *Renderer) Properties() (PropertiesID, error) {
	props := (PropertiesID)(C.SDL_GetRendererProperties((*C.SDL_Renderer)(renderer)))
	if props == 0 {
		return 0, getError()
	}
	return props, nil
}

const PropRendererNameString = "SDL.renderer.name"
const PropRendererWindowPointer = "SDL.renderer.window"
const PropRendererSurfacePointer = "SDL.renderer.surface"
const PropRendererVsyncNumber = "SDL.renderer.vsync"
const PropRendererMaxTextureSizeNumber = "SDL.renderer.max_texture_size"
const PropRendererTextureFormatsPointer = "SDL.renderer.texture_formats"
const PropRendererOutputColorspaceNumber = "SDL.renderer.output_colorspace"
const PropRendererHdrEnabledBoolean = "SDL.renderer.HDR_enabled"
const PropRendererSdrWhitePointFloat = "SDL.renderer.SDR_white_point"
const PropRendererHdrHeadroomFloat = "SDL.renderer.HDR_headroom"
const PropRendererD3D9DevicePointer = "SDL.renderer.d3d9.device"
const PropRendererD3D11DevicePointer = "SDL.renderer.d3d11.device"
const PropRendererD3D11SwapchainPointer = "SDL.renderer.d3d11.swap_chain"
const PropRendererD3D12DevicePointer = "SDL.renderer.d3d12.device"
const PropRendererD3D12SwapchainPointer = "SDL.renderer.d3d12.swap_chain"
const PropRendererD3D12CommandQueuePointer = "SDL.renderer.d3d12.command_queue"
const PropRendererVulkanInstancePointer = "SDL.renderer.vulkan.instance"
const PropRendererVulkanSurfaceNumber = "SDL.renderer.vulkan.surface"
const PropRendererVulkanPhysicalDevicePointer = "SDL.renderer.vulkan.physical_device"
const PropRendererVulkanDevicePointer = "SDL.renderer.vulkan.device"
const PropRendererVulkanGraphicsQueueFamilyIndexNumber = "SDL.renderer.vulkan.graphics_queue_family_index"
const PropRendererVulkanPresentQueueFamilyIndexNumber = "SDL.renderer.vulkan.present_queue_family_index"
const PropRendererVulkanSwapchainImageCountNumber = "SDL.renderer.vulkan.swapchain_image_count"
const PropRendererGPUDevicePointer = "SDL.renderer.gpu.device"

// Get the output size in pixels of a rendering context.
//
// This returns the true output size in pixels, ignoring any render targets or
// logical size and presentation.
//
// renderer: the rendering context.
//
// w: a pointer filled in with the width in pixels.
//
// h: a pointer filled in with the height in pixels.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRenderOutputSize
func (renderer *Renderer) OutputSize() (w, h int, err error) {
	var cw, ch C.int
	if !C.SDL_GetRenderOutputSize((*C.SDL_Renderer)(renderer), &cw, &ch) {
		return 0, 0, getError()
	}
	return int(cw), int(ch), nil
}

// Get the current output size in pixels of a rendering context.
//
// If a rendering target is active, this will return the size of the rendering
// target in pixels, otherwise if a logical size is set, it will return the
// logical size, otherwise it will return the value of
// SDL_GetRenderOutputSize().
//
// renderer: the rendering context.
//
// w: a pointer filled in with the current width.
//
// h: a pointer filled in with the current height.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetCurrentRenderOutputSize
func (renderer *Renderer) CurrentOutputSize() (w, h int, err error) {
	var cw, ch C.int
	if !C.SDL_GetCurrentRenderOutputSize((*C.SDL_Renderer)(renderer), &cw, &ch) {
		return 0, 0, getError()
	}
	return int(cw), int(ch), nil
}

// Create a texture for a rendering context.
//
// The contents of a texture when first created are not defined.
//
// renderer: the rendering context.
//
// format: one of the enumerated values in SDL_PixelFormat.
//
// access: one of the enumerated values in SDL_TextureAccess.
//
// w: the width of the texture in pixels.
//
// h: the height of the texture in pixels.
//
// Returns the created texture or NULL on failure; call SDL_GetError() for
// more information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateTexture
func (renderer *Renderer) CreateTexture(format PixelFormat, access TextureAccess, w int, h int) (*Texture, error) {
	t := C.SDL_CreateTexture((*C.SDL_Renderer)(renderer), (C.SDL_PixelFormat)(format), (C.SDL_TextureAccess)(access), (C.int)(w), (C.int)(h))
	if t == nil {
		return nil, getError()
	}
	return &Texture{t}, nil
}

// Create a texture from an existing surface.
//
// The surface is not modified or freed by this function.
//
// The SDL_TextureAccess hint for the created texture is
// `SDL_TEXTUREACCESS_STATIC`.
//
// The pixel format of the created texture may be different from the pixel
// format of the surface, and can be queried using the
// SDL_PROP_TEXTURE_FORMAT_NUMBER property.
//
// renderer: the rendering context.
//
// surface: the SDL_Surface structure containing pixel data used to fill
// the texture.
//
// Returns the created texture or NULL on failure; call SDL_GetError() for
// more information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateTextureFromSurface
func (renderer *Renderer) CreateTextureFromSurface(surface *Surface) (*Texture, error) {
	t := C.SDL_CreateTextureFromSurface((*C.SDL_Renderer)(renderer), surface.internal)
	if t == nil {
		return nil, getError()
	}
	return &Texture{t}, nil
}

// Create a texture for a rendering context with the specified properties.
//
// These are the supported properties:
//
// - `SDL_PROP_TEXTURE_CREATE_COLORSPACE_NUMBER`: an SDL_Colorspace value
// describing the texture colorspace, defaults to SDL_COLORSPACE_SRGB_LINEAR
// for floating point textures, SDL_COLORSPACE_HDR10 for 10-bit textures,
// SDL_COLORSPACE_SRGB for other RGB textures and SDL_COLORSPACE_JPEG for
// YUV textures.
// - `SDL_PROP_TEXTURE_CREATE_FORMAT_NUMBER`: one of the enumerated values in
// SDL_PixelFormat, defaults to the best RGBA format for the renderer
// - `SDL_PROP_TEXTURE_CREATE_ACCESS_NUMBER`: one of the enumerated values in
// SDL_TextureAccess, defaults to SDL_TEXTUREACCESS_STATIC
// - `SDL_PROP_TEXTURE_CREATE_WIDTH_NUMBER`: the width of the texture in
// pixels, required
// - `SDL_PROP_TEXTURE_CREATE_HEIGHT_NUMBER`: the height of the texture in
// pixels, required
// - `SDL_PROP_TEXTURE_CREATE_SDR_WHITE_POINT_FLOAT`: for HDR10 and floating
// point textures, this defines the value of 100% diffuse white, with higher
// values being displayed in the High Dynamic Range headroom. This defaults
// to 100 for HDR10 textures and 1.0 for floating point textures.
// - `SDL_PROP_TEXTURE_CREATE_HDR_HEADROOM_FLOAT`: for HDR10 and floating
// point textures, this defines the maximum dynamic range used by the
// content, in terms of the SDR white point. This would be equivalent to
// maxCLL / SDL_PROP_TEXTURE_CREATE_SDR_WHITE_POINT_FLOAT for HDR10 content.
// If this is defined, any values outside the range supported by the display
// will be scaled into the available HDR headroom, otherwise they are
// clipped.
//
// With the direct3d11 renderer:
//
// - `SDL_PROP_TEXTURE_CREATE_D3D11_TEXTURE_POINTER`: the ID3D11Texture2D
// associated with the texture, if you want to wrap an existing texture.
// - `SDL_PROP_TEXTURE_CREATE_D3D11_TEXTURE_U_POINTER`: the ID3D11Texture2D
// associated with the U plane of a YUV texture, if you want to wrap an
// existing texture.
// - `SDL_PROP_TEXTURE_CREATE_D3D11_TEXTURE_V_POINTER`: the ID3D11Texture2D
// associated with the V plane of a YUV texture, if you want to wrap an
// existing texture.
//
// With the direct3d12 renderer:
//
// - `SDL_PROP_TEXTURE_CREATE_D3D12_TEXTURE_POINTER`: the ID3D12Resource
// associated with the texture, if you want to wrap an existing texture.
// - `SDL_PROP_TEXTURE_CREATE_D3D12_TEXTURE_U_POINTER`: the ID3D12Resource
// associated with the U plane of a YUV texture, if you want to wrap an
// existing texture.
// - `SDL_PROP_TEXTURE_CREATE_D3D12_TEXTURE_V_POINTER`: the ID3D12Resource
// associated with the V plane of a YUV texture, if you want to wrap an
// existing texture.
//
// With the metal renderer:
//
// - `SDL_PROP_TEXTURE_CREATE_METAL_PIXELBUFFER_POINTER`: the CVPixelBufferRef
// associated with the texture, if you want to create a texture from an
// existing pixel buffer.
//
// With the opengl renderer:
//
// - `SDL_PROP_TEXTURE_CREATE_OPENGL_TEXTURE_NUMBER`: the GLuint texture
// associated with the texture, if you want to wrap an existing texture.
// - `SDL_PROP_TEXTURE_CREATE_OPENGL_TEXTURE_UV_NUMBER`: the GLuint texture
// associated with the UV plane of an NV12 texture, if you want to wrap an
// existing texture.
// - `SDL_PROP_TEXTURE_CREATE_OPENGL_TEXTURE_U_NUMBER`: the GLuint texture
// associated with the U plane of a YUV texture, if you want to wrap an
// existing texture.
// - `SDL_PROP_TEXTURE_CREATE_OPENGL_TEXTURE_V_NUMBER`: the GLuint texture
// associated with the V plane of a YUV texture, if you want to wrap an
// existing texture.
//
// With the opengles2 renderer:
//
// - `SDL_PROP_TEXTURE_CREATE_OPENGLES2_TEXTURE_NUMBER`: the GLuint texture
// associated with the texture, if you want to wrap an existing texture.
// - `SDL_PROP_TEXTURE_CREATE_OPENGLES2_TEXTURE_NUMBER`: the GLuint texture
// associated with the texture, if you want to wrap an existing texture.
// - `SDL_PROP_TEXTURE_CREATE_OPENGLES2_TEXTURE_UV_NUMBER`: the GLuint texture
// associated with the UV plane of an NV12 texture, if you want to wrap an
// existing texture.
// - `SDL_PROP_TEXTURE_CREATE_OPENGLES2_TEXTURE_U_NUMBER`: the GLuint texture
// associated with the U plane of a YUV texture, if you want to wrap an
// existing texture.
// - `SDL_PROP_TEXTURE_CREATE_OPENGLES2_TEXTURE_V_NUMBER`: the GLuint texture
// associated with the V plane of a YUV texture, if you want to wrap an
// existing texture.
//
// With the vulkan renderer:
//
// - `SDL_PROP_TEXTURE_CREATE_VULKAN_TEXTURE_NUMBER`: the VkImage with layout
// VK_IMAGE_LAYOUT_SHADER_READ_ONLY_OPTIMAL associated with the texture, if
// you want to wrap an existing texture.
//
// renderer: the rendering context.
//
// props: the properties to use.
//
// Returns the created texture or NULL on failure; call SDL_GetError() for
// more information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_CreateTextureWithProperties
func (renderer *Renderer) CreateTextureWithProperties(props PropertiesID) (*Texture, error) {
	t := C.SDL_CreateTextureWithProperties((*C.SDL_Renderer)(renderer), (C.SDL_PropertiesID)(props))
	if t == nil {
		return nil, getError()
	}
	return &Texture{t}, nil
}

const PropTextureCreateColorspaceNumber = "SDL.texture.create.colorspace"
const PropTextureCreateFormatNumber = "SDL.texture.create.format"
const PropTextureCreateAccessNumber = "SDL.texture.create.access"
const PropTextureCreateWidthNumber = "SDL.texture.create.width"
const PropTextureCreateHeightNumber = "SDL.texture.create.height"
const PropTextureCreateSdrWhitePointFloat = "SDL.texture.create.SDR_white_point"
const PropTextureCreateHdrHeadroomFloat = "SDL.texture.create.HDR_headroom"
const PropTextureCreateD3D11TexturePointer = "SDL.texture.create.d3d11.texture"
const PropTextureCreateD3D11TextureUPointer = "SDL.texture.create.d3d11.texture_u"
const PropTextureCreateD3D11TextureVPointer = "SDL.texture.create.d3d11.texture_v"
const PropTextureCreateD3D12TexturePointer = "SDL.texture.create.d3d12.texture"
const PropTextureCreateD3D12TextureUPointer = "SDL.texture.create.d3d12.texture_u"
const PropTextureCreateD3D12TextureVPointer = "SDL.texture.create.d3d12.texture_v"
const PropTextureCreateMetalPixelbufferPointer = "SDL.texture.create.metal.pixelbuffer"
const PropTextureCreateOpenglTextureNumber = "SDL.texture.create.opengl.texture"
const PropTextureCreateOpenglTextureUvNumber = "SDL.texture.create.opengl.texture_uv"
const PropTextureCreateOpenglTextureUNumber = "SDL.texture.create.opengl.texture_u"
const PropTextureCreateOpenglTextureVNumber = "SDL.texture.create.opengl.texture_v"
const PropTextureCreateOPENGLES2TextureNumber = "SDL.texture.create.opengles2.texture"
const PropTextureCreateOPENGLES2TextureUvNumber = "SDL.texture.create.opengles2.texture_uv"
const PropTextureCreateOPENGLES2TextureUNumber = "SDL.texture.create.opengles2.texture_u"
const PropTextureCreateOPENGLES2TextureVNumber = "SDL.texture.create.opengles2.texture_v"
const PropTextureCreateVulkanTextureNumber = "SDL.texture.create.vulkan.texture"

// Get the properties associated with a texture.
//
// The following read-only properties are provided by SDL:
//
// - `SDL_PROP_TEXTURE_COLORSPACE_NUMBER`: an SDL_Colorspace value describing
// the texture colorspace.
// - `SDL_PROP_TEXTURE_FORMAT_NUMBER`: one of the enumerated values in
// SDL_PixelFormat.
// - `SDL_PROP_TEXTURE_ACCESS_NUMBER`: one of the enumerated values in
// SDL_TextureAccess.
// - `SDL_PROP_TEXTURE_WIDTH_NUMBER`: the width of the texture in pixels.
// - `SDL_PROP_TEXTURE_HEIGHT_NUMBER`: the height of the texture in pixels.
// - `SDL_PROP_TEXTURE_SDR_WHITE_POINT_FLOAT`: for HDR10 and floating point
// textures, this defines the value of 100% diffuse white, with higher
// values being displayed in the High Dynamic Range headroom. This defaults
// to 100 for HDR10 textures and 1.0 for other textures.
// - `SDL_PROP_TEXTURE_HDR_HEADROOM_FLOAT`: for HDR10 and floating point
// textures, this defines the maximum dynamic range used by the content, in
// terms of the SDR white point. If this is defined, any values outside the
// range supported by the display will be scaled into the available HDR
// headroom, otherwise they are clipped. This defaults to 1.0 for SDR
// textures, 4.0 for HDR10 textures, and no default for floating point
// textures.
//
// With the direct3d11 renderer:
//
// - `SDL_PROP_TEXTURE_D3D11_TEXTURE_POINTER`: the ID3D11Texture2D associated
// with the texture
// - `SDL_PROP_TEXTURE_D3D11_TEXTURE_U_POINTER`: the ID3D11Texture2D
// associated with the U plane of a YUV texture
// - `SDL_PROP_TEXTURE_D3D11_TEXTURE_V_POINTER`: the ID3D11Texture2D
// associated with the V plane of a YUV texture
//
// With the direct3d12 renderer:
//
// - `SDL_PROP_TEXTURE_D3D12_TEXTURE_POINTER`: the ID3D12Resource associated
// with the texture
// - `SDL_PROP_TEXTURE_D3D12_TEXTURE_U_POINTER`: the ID3D12Resource associated
// with the U plane of a YUV texture
// - `SDL_PROP_TEXTURE_D3D12_TEXTURE_V_POINTER`: the ID3D12Resource associated
// with the V plane of a YUV texture
//
// With the vulkan renderer:
//
// - `SDL_PROP_TEXTURE_VULKAN_TEXTURE_NUMBER`: the VkImage associated with the
// texture
//
// With the opengl renderer:
//
// - `SDL_PROP_TEXTURE_OPENGL_TEXTURE_NUMBER`: the GLuint texture associated
// with the texture
// - `SDL_PROP_TEXTURE_OPENGL_TEXTURE_UV_NUMBER`: the GLuint texture
// associated with the UV plane of an NV12 texture
// - `SDL_PROP_TEXTURE_OPENGL_TEXTURE_U_NUMBER`: the GLuint texture associated
// with the U plane of a YUV texture
// - `SDL_PROP_TEXTURE_OPENGL_TEXTURE_V_NUMBER`: the GLuint texture associated
// with the V plane of a YUV texture
// - `SDL_PROP_TEXTURE_OPENGL_TEXTURE_TARGET_NUMBER`: the GLenum for the
// texture target (`GL_TEXTURE_2D`, `GL_TEXTURE_RECTANGLE_ARB`, etc)
// - `SDL_PROP_TEXTURE_OPENGL_TEX_W_FLOAT`: the texture coordinate width of
// the texture (0.0 - 1.0)
// - `SDL_PROP_TEXTURE_OPENGL_TEX_H_FLOAT`: the texture coordinate height of
// the texture (0.0 - 1.0)
//
// With the opengles2 renderer:
//
// - `SDL_PROP_TEXTURE_OPENGLES2_TEXTURE_NUMBER`: the GLuint texture
// associated with the texture
// - `SDL_PROP_TEXTURE_OPENGLES2_TEXTURE_UV_NUMBER`: the GLuint texture
// associated with the UV plane of an NV12 texture
// - `SDL_PROP_TEXTURE_OPENGLES2_TEXTURE_U_NUMBER`: the GLuint texture
// associated with the U plane of a YUV texture
// - `SDL_PROP_TEXTURE_OPENGLES2_TEXTURE_V_NUMBER`: the GLuint texture
// associated with the V plane of a YUV texture
// - `SDL_PROP_TEXTURE_OPENGLES2_TEXTURE_TARGET_NUMBER`: the GLenum for the
// texture target (`GL_TEXTURE_2D`, `GL_TEXTURE_EXTERNAL_OES`, etc)
//
// texture: the texture to query.
//
// Returns a valid property ID on success or 0 on failure; call
// SDL_GetError() for more information.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTextureProperties
func (texture *Texture) Properties() (PropertiesID, error) {
	props := (PropertiesID)(C.SDL_GetTextureProperties(texture.internal))
	if props == 0 {
		return 0, getError()
	}
	return props, nil
}

const PropTextureColorspaceNumber = "SDL.texture.colorspace"
const PropTextureFormatNumber = "SDL.texture.format"
const PropTextureAccessNumber = "SDL.texture.access"
const PropTextureWidthNumber = "SDL.texture.width"
const PropTextureHeightNumber = "SDL.texture.height"
const PropTextureSdrWhitePointFloat = "SDL.texture.SDR_white_point"
const PropTextureHdrHeadroomFloat = "SDL.texture.HDR_headroom"
const PropTextureD3D11TexturePointer = "SDL.texture.d3d11.texture"
const PropTextureD3D11TextureUPointer = "SDL.texture.d3d11.texture_u"
const PropTextureD3D11TextureVPointer = "SDL.texture.d3d11.texture_v"
const PropTextureD3D12TexturePointer = "SDL.texture.d3d12.texture"
const PropTextureD3D12TextureUPointer = "SDL.texture.d3d12.texture_u"
const PropTextureD3D12TextureVPointer = "SDL.texture.d3d12.texture_v"
const PropTextureOpenglTextureNumber = "SDL.texture.opengl.texture"
const PropTextureOpenglTextureUvNumber = "SDL.texture.opengl.texture_uv"
const PropTextureOpenglTextureUNumber = "SDL.texture.opengl.texture_u"
const PropTextureOpenglTextureVNumber = "SDL.texture.opengl.texture_v"
const PropTextureOpenglTextureTargetNumber = "SDL.texture.opengl.target"
const PropTextureOpenglTexWFloat = "SDL.texture.opengl.tex_w"
const PropTextureOpenglTexHFloat = "SDL.texture.opengl.tex_h"
const PropTextureOPENGLES2TextureNumber = "SDL.texture.opengles2.texture"
const PropTextureOPENGLES2TextureUvNumber = "SDL.texture.opengles2.texture_uv"
const PropTextureOPENGLES2TextureUNumber = "SDL.texture.opengles2.texture_u"
const PropTextureOPENGLES2TextureVNumber = "SDL.texture.opengles2.texture_v"
const PropTextureOPENGLES2TextureTargetNumber = "SDL.texture.opengles2.target"
const PropTextureVulkanTextureNumber = "SDL.texture.vulkan.texture"

// Get the renderer that created an SDL_Texture.
//
// texture: the texture to query.
//
// Returns a pointer to the SDL_Renderer that created the texture, or NULL on
// failure; call SDL_GetError() for more information.
//
// It is safe to call this function from any thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRendererFromTexture
func (texture *Texture) Renderer() (*Renderer, error) {
	r := (*Renderer)(C.SDL_GetRendererFromTexture(texture.internal))
	if r == nil {
		return nil, getError()
	}
	return r, nil
}

// Get the size of a texture, as floating point values.
//
// texture: the texture to query.
//
// w: a pointer filled in with the width of the texture in pixels. This
// argument can be NULL if you don't need this information.
//
// h: a pointer filled in with the height of the texture in pixels. This
// argument can be NULL if you don't need this information.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTextureSize
func (texture *Texture) Size() (w, h float32, err error) {
	var cw, ch C.float
	if !C.SDL_GetTextureSize(texture.internal, &cw, &ch) {
		return 0, 0, getError()
	}
	return float32(cw), float32(ch), nil
}

// Set an additional color value multiplied into render copy operations.
//
// When this texture is rendered, during the copy operation each source color
// channel is modulated by the appropriate color value according to the
// following formula:
//
// `srcC = srcC * (color / 255)`
//
// Color modulation is not always supported by the renderer; it will return
// false if color modulation is not supported.
//
// texture: the texture to update.
//
// r: the red color value multiplied into copy operations.
//
// g: the green color value multiplied into copy operations.
//
// b: the blue color value multiplied into copy operations.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetTextureColorMod
func (texture *Texture) SetColorMod(r byte, g byte, b byte) error {
	if !C.SDL_SetTextureColorMod(texture.internal, (C.Uint8)(r), (C.Uint8)(g), (C.Uint8)(b)) {
		return getError()
	}
	return nil
}

// Set an additional color value multiplied into render copy operations.
//
// When this texture is rendered, during the copy operation each source color
// channel is modulated by the appropriate color value according to the
// following formula:
//
// `srcC = srcC * color`
//
// Color modulation is not always supported by the renderer; it will return
// false if color modulation is not supported.
//
// texture: the texture to update.
//
// r: the red color value multiplied into copy operations.
//
// g: the green color value multiplied into copy operations.
//
// b: the blue color value multiplied into copy operations.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetTextureColorModFloat
func (texture *Texture) SetColorModFloat(r float32, g float32, b float32) error {
	if !C.SDL_SetTextureColorModFloat(texture.internal, (C.float)(r), (C.float)(g), (C.float)(b)) {
		return getError()
	}
	return nil
}

// Get the additional color value multiplied into render copy operations.
//
// texture: the texture to query.
//
// r: a pointer filled in with the current red color value.
//
// g: a pointer filled in with the current green color value.
//
// b: a pointer filled in with the current blue color value.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTextureColorMod
func (texture *Texture) ColorMod() (r, g, b byte, err error) {
	var cr, cg, cb C.Uint8
	if !C.SDL_GetTextureColorMod(texture.internal, &cr, &cg, &cb) {
		return 0, 0, 0, getError()
	}
	return byte(cr), byte(cg), byte(cb), nil
}

// Get the additional color value multiplied into render copy operations.
//
// texture: the texture to query.
//
// r: a pointer filled in with the current red color value.
//
// g: a pointer filled in with the current green color value.
//
// b: a pointer filled in with the current blue color value.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTextureColorModFloat
func (texture *Texture) ColorModFloat() (r, g, b float32, err error) {
	var cr, cg, cb C.float
	if !C.SDL_GetTextureColorModFloat(texture.internal, &cr, &cg, &cb) {
		return 0, 0, 0, getError()
	}
	return float32(cr), float32(cg), float32(cb), nil
}

// Set an additional alpha value multiplied into render copy operations.
//
// When this texture is rendered, during the copy operation the source alpha
// value is modulated by this alpha value according to the following formula:
//
// `srcA = srcA * (alpha / 255)`
//
// Alpha modulation is not always supported by the renderer; it will return
// false if alpha modulation is not supported.
//
// texture: the texture to update.
//
// alpha: the source alpha value multiplied into copy operations.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetTextureAlphaMod
func (texture *Texture) SetAlphaMod(alpha byte) error {
	if !C.SDL_SetTextureAlphaMod(texture.internal, (C.Uint8)(alpha)) {
		return getError()
	}
	return nil
}

// Set an additional alpha value multiplied into render copy operations.
//
// When this texture is rendered, during the copy operation the source alpha
// value is modulated by this alpha value according to the following formula:
//
// `srcA = srcA * alpha`
//
// Alpha modulation is not always supported by the renderer; it will return
// false if alpha modulation is not supported.
//
// texture: the texture to update.
//
// alpha: the source alpha value multiplied into copy operations.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetTextureAlphaModFloat
func (texture *Texture) SetAlphaModFloat(alpha float32) error {
	if !C.SDL_SetTextureAlphaModFloat(texture.internal, (C.float)(alpha)) {
		return getError()
	}
	return nil
}

// Get the additional alpha value multiplied into render copy operations.
//
// texture: the texture to query.
//
// alpha: a pointer filled in with the current alpha value.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTextureAlphaMod
func (texture *Texture) AlphaMod() (byte, error) {
	var alpha C.Uint8
	if !C.SDL_GetTextureAlphaMod(texture.internal, &alpha) {
		return 0, getError()
	}
	return byte(alpha), nil
}

// Get the additional alpha value multiplied into render copy operations.
//
// texture: the texture to query.
//
// alpha: a pointer filled in with the current alpha value.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTextureAlphaModFloat
func (texture *Texture) AlphaModFloat() (float32, error) {
	var alpha C.float
	if !C.SDL_GetTextureAlphaModFloat(texture.internal, &alpha) {
		return 0, getError()
	}
	return float32(alpha), nil
}

// Set the blend mode for a texture, used by SDL_RenderTexture().
//
// If the blend mode is not supported, the closest supported mode is chosen
// and this function returns false.
//
// texture: the texture to update.
//
// blendMode: the SDL_BlendMode to use for texture blending.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetTextureBlendMode
func (texture *Texture) SetBlendMode(blendMode BlendMode) error {
	if !C.SDL_SetTextureBlendMode(texture.internal, (C.SDL_BlendMode)(blendMode)) {
		return getError()
	}
	return nil
}

// Get the blend mode used for texture copy operations.
//
// texture: the texture to query.
//
// blendMode: a pointer filled in with the current SDL_BlendMode.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTextureBlendMode
func (texture *Texture) BlendMode() (BlendMode, error) {
	var blendMode C.SDL_BlendMode
	if !C.SDL_GetTextureBlendMode(texture.internal, &blendMode) {
		return BlendmodeInvalid, getError()
	}
	return BlendMode(blendMode), nil
}

// Set the scale mode used for texture scale operations.
//
// The default texture scale mode is SDL_SCALEMODE_LINEAR.
//
// If the scale mode is not supported, the closest supported mode is chosen.
//
// texture: the texture to update.
//
// scaleMode: the SDL_ScaleMode to use for texture scaling.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetTextureScaleMode
func (texture *Texture) SetScaleMode(scaleMode ScaleMode) error {
	if !C.SDL_SetTextureScaleMode(texture.internal, (C.SDL_ScaleMode)(scaleMode)) {
		return getError()
	}
	return nil
}

// Get the scale mode used for texture scale operations.
//
// texture: the texture to query.
//
// scaleMode: a pointer filled in with the current scale mode.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetTextureScaleMode
func (texture *Texture) ScaleMode() (ScaleMode, error) {
	var scaleMode C.SDL_ScaleMode
	if !C.SDL_GetTextureScaleMode(texture.internal, &scaleMode) {
		return 0, getError()
	}
	return ScaleMode(scaleMode), nil
}

// Update the given texture rectangle with new pixel data.
//
// The pixel data must be in the pixel format of the texture, which can be
// queried using the SDL_PROP_TEXTURE_FORMAT_NUMBER property.
//
// This is a fairly slow function, intended for use with static textures that
// do not change often.
//
// If the texture is intended to be updated often, it is preferred to create
// the texture as streaming and use the locking functions referenced below.
// While this function will work with streaming textures, for optimization
// reasons you may not get the pixels back if you lock the texture afterward.
//
// texture: the texture to update.
//
// rect: an SDL_Rect structure representing the area to update, or NULL
// to update the entire texture.
//
// pixels: the raw pixel data in the format of the texture.
//
// pitch: the number of bytes in a row of pixel data, including padding
// between lines.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_UpdateTexture
func (texture *Texture) Update(rect *Rect, pixels []byte, pitch int32) error {
	// TODO bounds check?
	var crect *C.SDL_Rect
	if rect != nil {
		crect = &C.SDL_Rect{C.int(rect.X), C.int(rect.Y), C.int(rect.W), C.int(rect.H)}
	}
	if !C.SDL_UpdateTexture(texture.internal, crect, unsafe.Pointer(unsafe.SliceData(pixels)), (C.int)(pitch)) {
		return getError()
	}
	return nil
}

// Update a rectangle within a planar YV12 or IYUV texture with new pixel
// data.
//
// You can use SDL_UpdateTexture() as long as your pixel data is a contiguous
// block of Y and U/V planes in the proper order, but this function is
// available if your pixel data is not contiguous.
//
// texture: the texture to update.
//
// rect: a pointer to the rectangle of pixels to update, or NULL to
// update the entire texture.
//
// Yplane: the raw pixel data for the Y plane.
//
// Ypitch: the number of bytes between rows of pixel data for the Y
// plane.
//
// Uplane: the raw pixel data for the U plane.
//
// Upitch: the number of bytes between rows of pixel data for the U
// plane.
//
// Vplane: the raw pixel data for the V plane.
//
// Vpitch: the number of bytes between rows of pixel data for the V
// plane.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_UpdateYUVTexture
func (texture *Texture) UpdateYUV(rect *Rect, Yplane []byte, Ypitch int, Uplane []byte, Upitch int, Vplane []byte, Vpitch int) error {
	// TODO bounds check?
	var crect *C.SDL_Rect
	if rect != nil {
		crect = &C.SDL_Rect{C.int(rect.X), C.int(rect.Y), C.int(rect.W), C.int(rect.H)}
	}
	if !C.SDL_UpdateYUVTexture(texture.internal, crect, (*C.Uint8)(unsafe.SliceData(Yplane)), (C.int)(Ypitch), (*C.Uint8)(unsafe.SliceData(Uplane)), (C.int)(Upitch), (*C.Uint8)(unsafe.SliceData(Vplane)), (C.int)(Vpitch)) {
		return getError()
	}
	return nil
}

// Update a rectangle within a planar NV12 or NV21 texture with new pixels.
//
// You can use SDL_UpdateTexture() as long as your pixel data is a contiguous
// block of NV12/21 planes in the proper order, but this function is available
// if your pixel data is not contiguous.
//
// texture: the texture to update.
//
// rect: a pointer to the rectangle of pixels to update, or NULL to
// update the entire texture.
//
// Yplane: the raw pixel data for the Y plane.
//
// Ypitch: the number of bytes between rows of pixel data for the Y
// plane.
//
// UVplane: the raw pixel data for the UV plane.
//
// UVpitch: the number of bytes between rows of pixel data for the UV
// plane.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_UpdateNVTexture
func (texture *Texture) UpdateNV(rect *Rect, Yplane []byte, Ypitch int, UVplane []byte, UVpitch int) error {
	// TODO bounds check?
	var crect *C.SDL_Rect
	if rect != nil {
		crect = &C.SDL_Rect{C.int(rect.X), C.int(rect.Y), C.int(rect.W), C.int(rect.H)}
	}
	if !C.SDL_UpdateNVTexture(texture.internal, crect, (*C.Uint8)(unsafe.SliceData(Yplane)), (C.int)(Ypitch), (*C.Uint8)(unsafe.SliceData(UVplane)), (C.int)(UVpitch)) {
		return getError()
	}
	return nil
}

// Lock a portion of the texture for **write-only** pixel access.
//
// As an optimization, the pixels made available for editing don't necessarily
// contain the old texture data. This is a write-only operation, and if you
// need to keep a copy of the texture data you should do that at the
// application level.
//
// You must use SDL_UnlockTexture() to unlock the pixels and apply any
// changes.
//
// texture: the texture to lock for access, which was created with
// `SDL_TEXTUREACCESS_STREAMING`.
//
// rect: an SDL_Rect structure representing the area to lock for access;
// NULL to lock the entire texture.
//
// pixels: this is filled in with a pointer to the locked pixels,
// appropriately offset by the locked area.
//
// pitch: this is filled in with the pitch of the locked pixels; the
// pitch is the length of one row in bytes.
//
// Returns true on success or false if the texture is not valid or was not
// created with `SDL_TEXTUREACCESS_STREAMING`; call SDL_GetError()
// for more information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_LockTexture
func (texture *Texture) Lock(rect *Rect) (pixels unsafe.Pointer, pitch int, err error) {
	var crect *C.SDL_Rect
	if rect != nil {
		crect = &C.SDL_Rect{C.int(rect.X), C.int(rect.Y), C.int(rect.W), C.int(rect.H)}
	}
	var cpixels unsafe.Pointer
	var cpitch C.int
	if !C.SDL_LockTexture(texture.internal, crect, &cpixels, &cpitch) {
		return nil, 0, getError()
	}
	return cpixels, int(cpitch), nil
}

// Lock a portion of the texture for **write-only** pixel access, and expose
// it as a SDL surface.
//
// Besides providing an SDL_Surface instead of raw pixel data, this function
// operates like SDL_LockTexture.
//
// As an optimization, the pixels made available for editing don't necessarily
// contain the old texture data. This is a write-only operation, and if you
// need to keep a copy of the texture data you should do that at the
// application level.
//
// You must use SDL_UnlockTexture() to unlock the pixels and apply any
// changes.
//
// The returned surface is freed internally after calling SDL_UnlockTexture()
// or SDL_DestroyTexture(). The caller should not free it.
//
// texture: the texture to lock for access, which must be created with
// `SDL_TEXTUREACCESS_STREAMING`.
//
// rect: a pointer to the rectangle to lock for access. If the rect is
// NULL, the entire texture will be locked.
//
// surface: a pointer to an SDL surface of size **rect**. Don't assume
// any specific pixel content.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_LockTextureToSurface
func (texture *Texture) LockToSurface(rect *Rect) (*Surface, error) {
	var crect *C.SDL_Rect
	if rect != nil {
		crect = &C.SDL_Rect{C.int(rect.X), C.int(rect.Y), C.int(rect.W), C.int(rect.H)}
	}
	var surface *C.SDL_Surface
	if !C.SDL_LockTextureToSurface(texture.internal, crect, &surface) {
		return nil, getError()
	}
	return &Surface{surface, nil}, nil
}

// Unlock a texture, uploading the changes to video memory, if needed.
//
// **Warning**: Please note that SDL_LockTexture() is intended to be
// write-only; it will not guarantee the previous contents of the texture will
// be provided. You must fully initialize any area of a texture that you lock
// before unlocking it, as the pixels might otherwise be uninitialized memory.
//
// Which is to say: locking and immediately unlocking a texture can result in
// corrupted textures, depending on the renderer in use.
//
// texture: a texture locked by SDL_LockTexture().
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_UnlockTexture
func (texture *Texture) Unlock() {
	C.SDL_UnlockTexture(texture.internal)
}

// Set a texture as the current rendering target.
//
// The default render target is the window for which the renderer was created.
// To stop rendering to a texture and render to the window again, call this
// function with a NULL `texture`.
//
// renderer: the rendering context.
//
// texture: the targeted texture, which must be created with the
// `SDL_TEXTUREACCESS_TARGET` flag, or NULL to render to the
// window instead of a texture.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetRenderTarget
func (renderer *Renderer) SetRenderTarget(texture *Texture) error {
	if !C.SDL_SetRenderTarget((*C.SDL_Renderer)(renderer), texture.internal) {
		return getError()
	}
	return nil
}

// Get the current render target.
//
// The default render target is the window for which the renderer was created,
// and is reported a NULL here.
//
// renderer: the rendering context.
//
// Returns the current render target or NULL for the default render target.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRenderTarget
func (renderer *Renderer) RenderTarget() *Texture {
	return &Texture{C.SDL_GetRenderTarget((*C.SDL_Renderer)(renderer))}
}

// Set a device independent resolution and presentation mode for rendering.
//
// This function sets the width and height of the logical rendering output.
// The renderer will act as if the window is always the requested dimensions,
// scaling to the actual window resolution as necessary.
//
// This can be useful for games that expect a fixed size, but would like to
// scale the output to whatever is available, regardless of how a user resizes
// a window, or if the display is high DPI.
//
// You can disable logical coordinates by setting the mode to
// SDL_LOGICAL_PRESENTATION_DISABLED, and in that case you get the full pixel
// resolution of the output window; it is safe to toggle logical presentation
// during the rendering of a frame: perhaps most of the rendering is done to
// specific dimensions but to make fonts look sharp, the app turns off logical
// presentation while drawing text.
//
// Letterboxing will only happen if logical presentation is enabled during
// SDL_RenderPresent; be sure to reenable it first if you were using it.
//
// You can convert coordinates in an event into rendering coordinates using
// SDL_ConvertEventToRenderCoordinates().
//
// renderer: the rendering context.
//
// w: the width of the logical resolution.
//
// h: the height of the logical resolution.
//
// mode: the presentation mode used.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetRenderLogicalPresentation
func (renderer *Renderer) SetLogicalPresentation(w int, h int, mode RendererLogicalPresentation) error {
	if !C.SDL_SetRenderLogicalPresentation((*C.SDL_Renderer)(renderer), (C.int)(w), (C.int)(h), (C.SDL_RendererLogicalPresentation)(mode)) {
		return getError()
	}
	return nil
}

// Get device independent resolution and presentation mode for rendering.
//
// This function gets the width and height of the logical rendering output, or
// the output size in pixels if a logical resolution is not enabled.
//
// renderer: the rendering context.
//
// w: an int to be filled with the width.
//
// h: an int to be filled with the height.
//
// mode: the presentation mode used.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRenderLogicalPresentation
func (renderer *Renderer) LogicalPresentation() (w int, h int, mode RendererLogicalPresentation, err error) {
	var cw, ch C.int
	var cmode C.SDL_RendererLogicalPresentation
	if !C.SDL_GetRenderLogicalPresentation((*C.SDL_Renderer)(renderer), &cw, &ch, &cmode) {
		return 0, 0, 0, getError()
	}
	return int(cw), int(ch), RendererLogicalPresentation(cmode), nil
}

// Get the final presentation rectangle for rendering.
//
// This function returns the calculated rectangle used for logical
// presentation, based on the presentation mode and output size. If logical
// presentation is disabled, it will fill the rectangle with the output size,
// in pixels.
//
// renderer: the rendering context.
//
// rect: a pointer filled in with the final presentation rectangle, may
// be NULL.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRenderLogicalPresentationRect
func (renderer *Renderer) LogicalPresentationRect() (FRect, error) {
	var rect C.SDL_FRect
	if !C.SDL_GetRenderLogicalPresentationRect((*C.SDL_Renderer)(renderer), &rect) {
		return FRect{}, getError()
	}
	return FRect{float32(rect.x), float32(rect.y), float32(rect.w), float32(rect.h)}, nil
}

// Get a point in render coordinates when given a point in window coordinates.
//
// This takes into account several states:
//
// - The window dimensions.
// - The logical presentation settings (SDL_SetRenderLogicalPresentation)
// - The scale (SDL_SetRenderScale)
// - The viewport (SDL_SetRenderViewport)
//
// renderer: the rendering context.
//
// window_x: the x coordinate in window coordinates.
//
// window_y: the y coordinate in window coordinates.
//
// x: a pointer filled with the x coordinate in render coordinates.
//
// y: a pointer filled with the y coordinate in render coordinates.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderCoordinatesFromWindow
func (renderer *Renderer) CoordinatesFromWindow(windowX float32, windowY float32) (x, y float32, err error) {
	var cx, cy C.float
	if !C.SDL_RenderCoordinatesFromWindow((*C.SDL_Renderer)(renderer), (C.float)(windowX), (C.float)(windowY), &cx, &cy) {
		return 0, 0, getError()
	}
	return float32(cx), float32(cy), nil
}

// Get a point in window coordinates when given a point in render coordinates.
//
// This takes into account several states:
//
// - The window dimensions.
// - The logical presentation settings (SDL_SetRenderLogicalPresentation)
// - The scale (SDL_SetRenderScale)
// - The viewport (SDL_SetRenderViewport)
//
// renderer: the rendering context.
//
// x: the x coordinate in render coordinates.
//
// y: the y coordinate in render coordinates.
//
// window_x: a pointer filled with the x coordinate in window
// coordinates.
//
// window_y: a pointer filled with the y coordinate in window
// coordinates.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderCoordinatesToWindow
func (renderer *Renderer) CoordinatesToWindow(x float32, y float32) (windowX, windowY float32, err error) {
	var cx, cy C.float
	if !C.SDL_RenderCoordinatesToWindow((*C.SDL_Renderer)(renderer), (C.float)(x), (C.float)(y), &cx, &cy) {
		return 0, 0, getError()
	}
	return float32(cx), float32(cy), nil
}

// Convert the coordinates in an event to render coordinates.
//
// This takes into account several states:
//
// - The window dimensions.
// - The logical presentation settings (SDL_SetRenderLogicalPresentation)
// - The scale (SDL_SetRenderScale)
// - The viewport (SDL_SetRenderViewport)
//
// Various event types are converted with this function: mouse, touch, pen,
// etc.
//
// Touch coordinates are converted from normalized coordinates in the window
// to non-normalized rendering coordinates.
//
// Relative mouse coordinates (xrel and yrel event fields) are _also_
// converted. Applications that do not want these fields converted should use
// SDL_RenderCoordinatesFromWindow() on the specific event fields instead of
// converting the entire event structure.
//
// Once converted, coordinates may be outside the rendering area.
//
// renderer: the rendering context.
//
// event: the event to modify.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_ConvertEventToRenderCoordinates
func (renderer *Renderer) ConvertEventToRenderCoordinates(event *Event) error {
	if !C.SDL_ConvertEventToRenderCoordinates((*C.SDL_Renderer)(renderer), &event.internal) {
		return getError()
	}
	return nil
}

// Set the drawing area for rendering on the current target.
//
// Drawing will clip to this area (separately from any clipping done with
// SDL_SetRenderClipRect), and the top left of the area will become coordinate
// (0, 0) for future drawing commands.
//
// The area's width and height must be >= 0.
//
// renderer: the rendering context.
//
// rect: the SDL_Rect structure representing the drawing area, or NULL
// to set the viewport to the entire target.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetRenderViewport
func (renderer *Renderer) SetViewport(rect *Rect) bool {
	var crect *C.SDL_Rect
	if rect != nil {
		crect = &C.SDL_Rect{C.int(rect.X), C.int(rect.Y), C.int(rect.W), C.int(rect.H)}
	}
	return (bool)(C.SDL_SetRenderViewport((*C.SDL_Renderer)(renderer), crect))
}

// Get the drawing area for the current target.
//
// renderer: the rendering context.
//
// rect: an SDL_Rect structure filled in with the current drawing area.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRenderViewport
func (renderer *Renderer) Viewport() (Rect, error) {
	var rect C.SDL_Rect
	if !C.SDL_GetRenderViewport((*C.SDL_Renderer)(renderer), &rect) {
		return Rect{}, getError()
	}
	return Rect{int(rect.x), int(rect.y), int(rect.w), int(rect.h)}, nil
}

// Return whether an explicit rectangle was set as the viewport.
//
// This is useful if you're saving and restoring the viewport and want to know
// whether you should restore a specific rectangle or NULL. Note that the
// viewport is always reset when changing rendering targets.
//
// renderer: the rendering context.
//
// Returns true if the viewport was set to a specific rectangle, or false if
// it was set to NULL (the entire target).
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderViewportSet
func (renderer *Renderer) ViewportSet() bool {
	return (bool)(C.SDL_RenderViewportSet((*C.SDL_Renderer)(renderer)))
}

// Get the safe area for rendering within the current viewport.
//
// Some devices have portions of the screen which are partially obscured or
// not interactive, possibly due to on-screen controls, curved edges, camera
// notches, TV overscan, etc. This function provides the area of the current
// viewport which is safe to have interactible content. You should continue
// rendering into the rest of the render target, but it should not contain
// visually important or interactible content.
//
// renderer: the rendering context.
//
// rect: a pointer filled in with the area that is safe for interactive
// content.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRenderSafeArea
func (renderer *Renderer) SafeArea() (Rect, error) {
	var rect C.SDL_Rect
	if !C.SDL_GetRenderSafeArea((*C.SDL_Renderer)(renderer), &rect) {
		return Rect{}, getError()
	}
	return Rect{int(rect.x), int(rect.y), int(rect.w), int(rect.h)}, nil
}

// Set the clip rectangle for rendering on the specified target.
//
// renderer: the rendering context.
//
// rect: an SDL_Rect structure representing the clip area, relative to
// the viewport, or NULL to disable clipping.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetRenderClipRect
func (renderer *Renderer) SetClipRect(rect *Rect) error {
	var crect *C.SDL_Rect
	if rect != nil {
		crect = &C.SDL_Rect{C.int(rect.X), C.int(rect.Y), C.int(rect.W), C.int(rect.H)}
	}
	if !C.SDL_SetRenderClipRect((*C.SDL_Renderer)(renderer), crect) {
		return getError()
	}
	return nil
}

// Get the clip rectangle for the current target.
//
// renderer: the rendering context.
//
// rect: an SDL_Rect structure filled in with the current clipping area
// or an empty rectangle if clipping is disabled.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRenderClipRect
func (renderer *Renderer) ClipRect() (Rect, error) {
	var rect C.SDL_Rect
	if !C.SDL_GetRenderClipRect((*C.SDL_Renderer)(renderer), &rect) {
		return Rect{}, getError()
	}
	return Rect{int(rect.x), int(rect.y), int(rect.w), int(rect.h)}, nil
}

// Get whether clipping is enabled on the given renderer.
//
// renderer: the rendering context.
//
// Returns true if clipping is enabled or false if not; call SDL_GetError()
// for more information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderClipEnabled
func (renderer *Renderer) ClipEnabled() bool {
	return (bool)(C.SDL_RenderClipEnabled((*C.SDL_Renderer)(renderer)))
}

// Set the drawing scale for rendering on the current target.
//
// The drawing coordinates are scaled by the x/y scaling factors before they
// are used by the renderer. This allows resolution independent drawing with a
// single coordinate system.
//
// If this results in scaling or subpixel drawing by the rendering backend, it
// will be handled using the appropriate quality hints. For best results use
// integer scaling factors.
//
// renderer: the rendering context.
//
// scaleX: the horizontal scaling factor.
//
// scaleY: the vertical scaling factor.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetRenderScale
func (renderer *Renderer) SetScale(scaleX float32, scaleY float32) error {
	if !C.SDL_SetRenderScale((*C.SDL_Renderer)(renderer), (C.float)(scaleX), (C.float)(scaleY)) {
		return getError()
	}
	return nil
}

// Get the drawing scale for the current target.
//
// renderer: the rendering context.
//
// scaleX: a pointer filled in with the horizontal scaling factor.
//
// scaleY: a pointer filled in with the vertical scaling factor.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRenderScale
func (renderer *Renderer) Scale() (scaleX, scaleY float32, err error) {
	var x, y C.float
	if !C.SDL_GetRenderScale((*C.SDL_Renderer)(renderer), &x, &y) {
		return 0, 0, getError()
	}
	return float32(x), float32(y), nil
}

// Set the color used for drawing operations.
//
// Set the color for drawing or filling rectangles, lines, and points, and for
// SDL_RenderClear().
//
// renderer: the rendering context.
//
// r: the red value used to draw on the rendering target.
//
// g: the green value used to draw on the rendering target.
//
// b: the blue value used to draw on the rendering target.
//
// a: the alpha value used to draw on the rendering target; usually
// `SDL_ALPHA_OPAQUE` (255). Use SDL_SetRenderDrawBlendMode to
// specify how the alpha channel is used.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawColor
func (renderer *Renderer) SetDrawColor(r byte, g byte, b byte, a byte) error {
	if !C.SDL_SetRenderDrawColor((*C.SDL_Renderer)(renderer), (C.Uint8)(r), (C.Uint8)(g), (C.Uint8)(b), (C.Uint8)(a)) {
		return getError()
	}
	return nil
}

// Set the color used for drawing operations (Rect, Line and Clear).
//
// Set the color for drawing or filling rectangles, lines, and points, and for
// SDL_RenderClear().
//
// renderer: the rendering context.
//
// r: the red value used to draw on the rendering target.
//
// g: the green value used to draw on the rendering target.
//
// b: the blue value used to draw on the rendering target.
//
// a: the alpha value used to draw on the rendering target. Use
// SDL_SetRenderDrawBlendMode to specify how the alpha channel is
// used.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawColorFloat
func (renderer *Renderer) SetDrawColorFloat(r float32, g float32, b float32, a float32) error {
	if !C.SDL_SetRenderDrawColorFloat((*C.SDL_Renderer)(renderer), (C.float)(r), (C.float)(g), (C.float)(b), (C.float)(a)) {
		return getError()
	}
	return nil
}

// Get the color used for drawing operations (Rect, Line and Clear).
//
// renderer: the rendering context.
//
// r: a pointer filled in with the red value used to draw on the
// rendering target.
//
// g: a pointer filled in with the green value used to draw on the
// rendering target.
//
// b: a pointer filled in with the blue value used to draw on the
// rendering target.
//
// a: a pointer filled in with the alpha value used to draw on the
// rendering target; usually `SDL_ALPHA_OPAQUE` (255).
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawColor
func (renderer *Renderer) DrawColor() (r, g, b, a byte, err error) {
	var cr, cg, cb, ca C.Uint8
	if !C.SDL_GetRenderDrawColor((*C.SDL_Renderer)(renderer), &cr, &cg, &cb, &ca) {
		return 0, 0, 0, 0, getError()
	}
	return byte(cr), byte(cg), byte(cb), byte(ca), nil
}

// Get the color used for drawing operations (Rect, Line and Clear).
//
// renderer: the rendering context.
//
// r: a pointer filled in with the red value used to draw on the
// rendering target.
//
// g: a pointer filled in with the green value used to draw on the
// rendering target.
//
// b: a pointer filled in with the blue value used to draw on the
// rendering target.
//
// a: a pointer filled in with the alpha value used to draw on the
// rendering target.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawColorFloat
func (renderer *Renderer) DrawColorFloat() (r, g, b, a float32, err error) {
	var cr, cg, cb, ca C.float
	if !C.SDL_GetRenderDrawColorFloat((*C.SDL_Renderer)(renderer), &cr, &cg, &cb, &ca) {
		return 0, 0, 0, 0, getError()
	}
	return float32(cr), float32(cg), float32(cb), float32(ca), nil
}

// Set the color scale used for render operations.
//
// The color scale is an additional scale multiplied into the pixel color
// value while rendering. This can be used to adjust the brightness of colors
// during HDR rendering, or changing HDR video brightness when playing on an
// SDR display.
//
// The color scale does not affect the alpha channel, only the color
// brightness.
//
// renderer: the rendering context.
//
// scale: the color scale value.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetRenderColorScale
func (renderer *Renderer) SetColorScale(scale float32) error {
	if !C.SDL_SetRenderColorScale((*C.SDL_Renderer)(renderer), (C.float)(scale)) {
		return getError()
	}
	return nil
}

// Get the color scale used for render operations.
//
// renderer: the rendering context.
//
// scale: a pointer filled in with the current color scale value.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRenderColorScale
func (renderer *Renderer) ColorScale() (float32, error) {
	var scale C.float
	if !C.SDL_GetRenderColorScale((*C.SDL_Renderer)(renderer), &scale) {
		return 0, getError()
	}
	return float32(scale), nil
}

// Set the blend mode used for drawing operations (Fill and Line).
//
// If the blend mode is not supported, the closest supported mode is chosen.
//
// renderer: the rendering context.
//
// blendMode: the SDL_BlendMode to use for blending.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetRenderDrawBlendMode
func (renderer *Renderer) SetDrawBlendMode(blendMode BlendMode) error {
	if !C.SDL_SetRenderDrawBlendMode((*C.SDL_Renderer)(renderer), (C.SDL_BlendMode)(blendMode)) {
		return getError()
	}
	return nil
}

// Get the blend mode used for drawing operations.
//
// renderer: the rendering context.
//
// blendMode: a pointer filled in with the current SDL_BlendMode.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRenderDrawBlendMode
func (renderer *Renderer) DrawBlendMode() (BlendMode, error) {
	var blendMode C.SDL_BlendMode
	if !C.SDL_GetRenderDrawBlendMode((*C.SDL_Renderer)(renderer), &blendMode) {
		return 0, getError()
	}
	return BlendMode(blendMode), nil
}

// Clear the current rendering target with the drawing color.
//
// This function clears the entire rendering target, ignoring the viewport and
// the clip rectangle. Note, that clearing will also set/fill all pixels of
// the rendering target to current renderer draw color, so make sure to invoke
// SDL_SetRenderDrawColor() when needed.
//
// renderer: the rendering context.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderClear
func (renderer *Renderer) Clear() error {
	if !C.SDL_RenderClear((*C.SDL_Renderer)(renderer)) {
		return getError()
	}
	return nil
}

// Draw a point on the current rendering target at subpixel precision.
//
// renderer: the renderer which should draw a point.
//
// x: the x coordinate of the point.
//
// y: the y coordinate of the point.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderPoint
func (renderer *Renderer) Point(x, y float32) error {
	if !C.SDL_RenderPoint((*C.SDL_Renderer)(renderer), (C.float)(x), (C.float)(y)) {
		return getError()
	}
	return nil
}

// Draw multiple points on the current rendering target at subpixel precision.
//
// renderer: the renderer which should draw multiple points.
//
// points: the points to draw.
//
// count: the number of points to draw.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderPoints
func (renderer *Renderer) Points(points []FPoint) error {
	if !C.SDL_RenderPoints((*C.SDL_Renderer)(renderer), (*C.SDL_FPoint)(unsafe.Pointer(unsafe.SliceData(points))), (C.int)(len(points))) {
		return getError()
	}
	return nil
}

// Draw a line on the current rendering target at subpixel precision.
//
// renderer: the renderer which should draw a line.
//
// x1: the x coordinate of the start point.
//
// y1: the y coordinate of the start point.
//
// x2: the x coordinate of the end point.
//
// y2: the y coordinate of the end point.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderLine
func (renderer *Renderer) Line(x1, y1, x2, y2 float32) error {
	if !C.SDL_RenderLine((*C.SDL_Renderer)(renderer), (C.float)(x1), (C.float)(y1), (C.float)(x2), (C.float)(y2)) {
		return getError()
	}
	return nil
}

// Draw a series of connected lines on the current rendering target at
// subpixel precision.
//
// renderer: the renderer which should draw multiple lines.
//
// points: the points along the lines.
//
// count: the number of points, drawing count-1 lines.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderLines
func (renderer *Renderer) Lines(points []FPoint) error {
	if !C.SDL_RenderLines((*C.SDL_Renderer)(renderer), (*C.SDL_FPoint)(unsafe.Pointer(unsafe.SliceData(points))), (C.int)(len(points))) {
		return getError()
	}
	return nil
}

// Draw a rectangle on the current rendering target at subpixel precision.
//
// renderer: the renderer which should draw a rectangle.
//
// rect: a pointer to the destination rectangle, or NULL to outline the
// entire rendering target.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderRect
func (renderer *Renderer) Rect(rect *FRect) error {
	if !C.SDL_RenderRect((*C.SDL_Renderer)(renderer), (*C.SDL_FRect)(unsafe.Pointer(rect))) {
		return getError()
	}
	return nil
}

// Draw some number of rectangles on the current rendering target at subpixel
// precision.
//
// renderer: the renderer which should draw multiple rectangles.
//
// rects: a pointer to an array of destination rectangles.
//
// count: the number of rectangles.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderRects
func (renderer *Renderer) Rects(rects []FRect) error {
	if !C.SDL_RenderRects((*C.SDL_Renderer)(renderer), (*C.SDL_FRect)(unsafe.Pointer(unsafe.SliceData(rects))), (C.int)(len(rects))) {
		return getError()
	}
	return nil
}

// Fill a rectangle on the current rendering target with the drawing color at
// subpixel precision.
//
// renderer: the renderer which should fill a rectangle.
//
// rect: a pointer to the destination rectangle, or NULL for the entire
// rendering target.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderFillRect
func (renderer *Renderer) FillRect(rect *FRect) error {
	if !C.SDL_RenderFillRect((*C.SDL_Renderer)(renderer), (*C.SDL_FRect)(unsafe.Pointer(rect))) {
		return getError()
	}
	return nil
}

// Fill some number of rectangles on the current rendering target with the
// drawing color at subpixel precision.
//
// renderer: the renderer which should fill multiple rectangles.
//
// rects: a pointer to an array of destination rectangles.
//
// count: the number of rectangles.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderFillRects
func (renderer *Renderer) FillRects(rects []FRect) error {
	if !C.SDL_RenderFillRects((*C.SDL_Renderer)(renderer), (*C.SDL_FRect)(unsafe.Pointer(unsafe.SliceData(rects))), (C.int)(len(rects))) {
		return getError()
	}
	return nil
}

// Copy a portion of the texture to the current rendering target at subpixel
// precision.
//
// renderer: the renderer which should copy parts of a texture.
//
// texture: the source texture.
//
// srcrect: a pointer to the source rectangle, or NULL for the entire
// texture.
//
// dstrect: a pointer to the destination rectangle, or NULL for the
// entire rendering target.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderTexture
func (renderer *Renderer) Texture(texture *Texture, srcrect *FRect, dstrect *FRect) error {
	if !C.SDL_RenderTexture((*C.SDL_Renderer)(renderer), texture.internal, (*C.SDL_FRect)(unsafe.Pointer(srcrect)), (*C.SDL_FRect)(unsafe.Pointer(dstrect))) {
		return getError()
	}
	return nil
}

// Copy a portion of the source texture to the current rendering target, with
// rotation and flipping, at subpixel precision.
//
// renderer: the renderer which should copy parts of a texture.
//
// texture: the source texture.
//
// srcrect: a pointer to the source rectangle, or NULL for the entire
// texture.
//
// dstrect: a pointer to the destination rectangle, or NULL for the
// entire rendering target.
//
// angle: an angle in degrees that indicates the rotation that will be
// applied to dstrect, rotating it in a clockwise direction.
//
// center: a pointer to a point indicating the point around which
// dstrect will be rotated (if NULL, rotation will be done
// around dstrect.w/2, dstrect.h/2).
//
// flip: an SDL_FlipMode value stating which flipping actions should be
// performed on the texture.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderTextureRotated
func (renderer *Renderer) TextureRotated(texture *Texture, srcrect *FRect, dstrect *FRect, angle float64, center *FPoint, flip FlipMode) error {
	if !C.SDL_RenderTextureRotated((*C.SDL_Renderer)(renderer), texture.internal, (*C.SDL_FRect)(unsafe.Pointer(srcrect)), (*C.SDL_FRect)(unsafe.Pointer(dstrect)), (C.double)(angle), (*C.SDL_FPoint)(unsafe.Pointer(center)), (C.SDL_FlipMode)(flip)) {
		return getError()
	}
	return nil
}

// Copy a portion of the source texture to the current rendering target, with
// affine transform, at subpixel precision.
//
// renderer: the renderer which should copy parts of a texture.
//
// texture: the source texture.
//
// srcrect: a pointer to the source rectangle, or NULL for the entire
// texture.
//
// origin: a pointer to a point indicating where the top-left corner of
// srcrect should be mapped to, or NULL for the rendering
// target's origin.
//
// right: a pointer to a point indicating where the top-right corner of
// srcrect should be mapped to, or NULL for the rendering
// target's top-right corner.
//
// down: a pointer to a point indicating where the bottom-left corner of
// srcrect should be mapped to, or NULL for the rendering target's
// bottom-left corner.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// You may only call this function from the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderTextureAffine
func (renderer *Renderer) TextureAffine(texture *Texture, srcrect *FRect, origin *FPoint, right *FPoint, down *FPoint) error {
	if !C.SDL_RenderTextureAffine((*C.SDL_Renderer)(renderer), texture.internal, (*C.SDL_FRect)(unsafe.Pointer(srcrect)), (*C.SDL_FPoint)(unsafe.Pointer(origin)), (*C.SDL_FPoint)(unsafe.Pointer(right)), (*C.SDL_FPoint)(unsafe.Pointer(down))) {
		return getError()
	}
	return nil
}

// Tile a portion of the texture to the current rendering target at subpixel
// precision.
//
// The pixels in `srcrect` will be repeated as many times as needed to
// completely fill `dstrect`.
//
// renderer: the renderer which should copy parts of a texture.
//
// texture: the source texture.
//
// srcrect: a pointer to the source rectangle, or NULL for the entire
// texture.
//
// scale: the scale used to transform srcrect into the destination
// rectangle, e.g. a 32x32 texture with a scale of 2 would fill
// 64x64 tiles.
//
// dstrect: a pointer to the destination rectangle, or NULL for the
// entire rendering target.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderTextureTiled
func (renderer *Renderer) TextureTiled(texture *Texture, srcrect *FRect, scale float32, dstrect *FRect) error {
	if !C.SDL_RenderTextureTiled((*C.SDL_Renderer)(renderer), texture.internal, (*C.SDL_FRect)(unsafe.Pointer(srcrect)), (C.float)(scale), (*C.SDL_FRect)(unsafe.Pointer(dstrect))) {
		return getError()
	}
	return nil
}

// Perform a scaled copy using the 9-grid algorithm to the current rendering
// target at subpixel precision.
//
// The pixels in the texture are split into a 3x3 grid, using the different
// corner sizes for each corner, and the sides and center making up the
// remaining pixels. The corners are then scaled using `scale` and fit into
// the corners of the destination rectangle. The sides and center are then
// stretched into place to cover the remaining destination rectangle.
//
// renderer: the renderer which should copy parts of a texture.
//
// texture: the source texture.
//
// srcrect: the SDL_Rect structure representing the rectangle to be used
// for the 9-grid, or NULL to use the entire texture.
//
// left_width: the width, in pixels, of the left corners in `srcrect`.
//
// right_width: the width, in pixels, of the right corners in `srcrect`.
//
// top_height: the height, in pixels, of the top corners in `srcrect`.
//
// bottom_height: the height, in pixels, of the bottom corners in
// `srcrect`.
//
// scale: the scale used to transform the corner of `srcrect` into the
// corner of `dstrect`, or 0.0f for an unscaled copy.
//
// dstrect: a pointer to the destination rectangle, or NULL for the
// entire rendering target.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderTexture9Grid
func (renderer *Renderer) Texture9Grid(texture *Texture, srcrect *FRect, leftWidth float32, rightWidth float32, topHeight float32, bottomHeight float32, scale float32, dstrect *FRect) error {
	if !C.SDL_RenderTexture9Grid((*C.SDL_Renderer)(renderer), texture.internal, (*C.SDL_FRect)(unsafe.Pointer(srcrect)), (C.float)(leftWidth), (C.float)(rightWidth), (C.float)(topHeight), (C.float)(bottomHeight), (C.float)(scale), (*C.SDL_FRect)(unsafe.Pointer(dstrect))) {
		return getError()
	}
	return nil
}

// Render a list of triangles, optionally using a texture and indices into the
// vertex array Color and alpha modulation is done per vertex
// (SDL_SetTextureColorMod and SDL_SetTextureAlphaMod are ignored).
//
// renderer: the rendering context.
//
// texture: (optional) The SDL texture to use.
//
// vertices: vertices.
//
// num_vertices: number of vertices.
//
// indices: (optional) An array of integer indices into the 'vertices'
// array, if NULL all vertices will be rendered in sequential
// order.
//
// num_indices: number of indices.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderGeometry
func (renderer *Renderer) Geometry(texture *Texture, vertices []Vertex, indices []int32) error {
	if !C.SDL_RenderGeometry((*C.SDL_Renderer)(renderer), texture.internal, (*C.SDL_Vertex)(unsafe.Pointer(unsafe.SliceData(vertices))), (C.int)(len(vertices)), (*C.int)(unsafe.Pointer(unsafe.SliceData(indices))), (C.int)(len(indices))) {
		return getError()
	}
	return nil
}

// Render a list of triangles, optionally using a texture and indices into the
// vertex arrays Color and alpha modulation is done per vertex
// (SDL_SetTextureColorMod and SDL_SetTextureAlphaMod are ignored).
//
// renderer: the rendering context.
//
// texture: (optional) The SDL texture to use.
//
// xy: vertex positions.
//
// xy_stride: byte size to move from one element to the next element.
//
// color: vertex colors (as SDL_FColor).
//
// color_stride: byte size to move from one element to the next element.
//
// uv: vertex normalized texture coordinates.
//
// uv_stride: byte size to move from one element to the next element.
//
// num_vertices: number of vertices.
//
// indices: (optional) An array of indices into the 'vertices' arrays,
// if NULL all vertices will be rendered in sequential order.
//
// num_indices: number of indices.
//
// size_indices: index size: 1 (byte), 2 (short), 4 (int).
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderGeometryRaw
func (renderer *Renderer) GeometryRaw(texture *Texture, xy *float32, xy_stride int32, color *FColor, color_stride int32, uv *float32, uv_stride int32, num_vertices int32, indices unsafe.Pointer, num_indices int32, size_indices int32) error {
	if !C.SDL_RenderGeometryRaw((*C.SDL_Renderer)(renderer), texture.internal, (*C.float)(xy), (C.int)(xy_stride), (*C.SDL_FColor)(unsafe.Pointer(color)), (C.int)(color_stride), (*C.float)(uv), (C.int)(uv_stride), (C.int)(num_vertices), indices, (C.int)(num_indices), (C.int)(size_indices)) {
		return getError()
	}
	return nil
}

// Read pixels from the current rendering target.
//
// The returned surface should be freed with SDL_DestroySurface()
//
// **WARNING**: This is a very slow operation, and should not be used
// frequently. If you're using this on the main rendering target, it should be
// called after rendering and before SDL_RenderPresent().
//
// renderer: the rendering context.
//
// rect: an SDL_Rect structure representing the area in pixels relative
// to the to current viewport, or NULL for the entire viewport.
//
// Returns a new SDL_Surface on success or NULL on failure; call
// SDL_GetError() for more information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderReadPixels
func (renderer *Renderer) ReadPixels(rect *Rect) (*Surface, error) {
	var crect *C.SDL_Rect
	if rect != nil {
		crect = &C.SDL_Rect{C.int(rect.X), C.int(rect.Y), C.int(rect.W), C.int(rect.H)}
	}
	s := C.SDL_RenderReadPixels((*C.SDL_Renderer)(renderer), crect)
	if s == nil {
		return nil, getError()
	}
	return &Surface{s, nil}, nil
}

// Update the screen with any rendering performed since the previous call.
//
// SDL's rendering functions operate on a backbuffer; that is, calling a
// rendering function such as SDL_RenderLine() does not directly put a line on
// the screen, but rather updates the backbuffer. As such, you compose your
// entire scene and *present* the composed backbuffer to the screen as a
// complete picture.
//
// Therefore, when using SDL's rendering API, one does all drawing intended
// for the frame, and then calls this function once per frame to present the
// final drawing to the user.
//
// The backbuffer should be considered invalidated after each present; do not
// assume that previous contents will exist between frames. You are strongly
// encouraged to call SDL_RenderClear() to initialize the backbuffer before
// starting each new frame's drawing, even if you plan to overwrite every
// pixel.
//
// Please note, that in case of rendering to a texture - there is **no need**
// to call `SDL_RenderPresent` after drawing needed objects to a texture, and
// should not be done; you are only required to change back the rendering
// target to default via `SDL_SetRenderTarget(renderer, NULL)` afterwards, as
// textures by themselves do not have a concept of backbuffers. Calling
// SDL_RenderPresent while rendering to a texture will still update the screen
// with any current drawing that has been done _to the window itself_.
//
// renderer: the rendering context.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderPresent
func (renderer *Renderer) Present() error {
	if !C.SDL_RenderPresent((*C.SDL_Renderer)(renderer)) {
		return getError()
	}
	return nil
}

// Destroy the specified texture.
//
// Passing NULL or an otherwise invalid texture will set the SDL error message
// to "Invalid texture".
//
// texture: the texture to destroy.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DestroyTexture
func (texture *Texture) Destroy() {
	C.SDL_DestroyTexture(texture.internal)
}

// Destroy the rendering context for a window and free all associated
// textures.
//
// This should be called before destroying the associated window.
//
// renderer: the rendering context.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_DestroyRenderer
func (renderer *Renderer) Destroy() {
	C.SDL_DestroyRenderer((*C.SDL_Renderer)(renderer))
}

// Force the rendering context to flush any pending commands and state.
//
// You do not need to (and in fact, shouldn't) call this function unless you
// are planning to call into OpenGL/Direct3D/Metal/whatever directly, in
// addition to using an SDL_Renderer.
//
// This is for a very-specific case: if you are using SDL's render API, and
// you plan to make OpenGL/D3D/whatever calls in addition to SDL render API
// calls. If this applies, you should call this function between calls to
// SDL's render API and the low-level API you're using in cooperation.
//
// In all other cases, you can ignore this function.
//
// This call makes SDL flush any pending rendering work it was queueing up to
// do later in a single batch, and marks any internal cached state as invalid,
// so it'll prepare all its state again later, from scratch.
//
// This means you do not need to save state in your rendering code to protect
// the SDL renderer. However, there lots of arbitrary pieces of Direct3D and
// OpenGL state that can confuse things; you should use your best judgment and
// be prepared to make changes if specific state needs to be protected.
//
// renderer: the rendering context.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_FlushRenderer
func (renderer *Renderer) Flush() error {
	if !C.SDL_FlushRenderer((*C.SDL_Renderer)(renderer)) {
		return getError()
	}
	return nil
}

// Get the CAMetalLayer associated with the given Metal renderer.
//
// This function returns `void *`, so SDL doesn't have to include Metal's
// headers, but it can be safely cast to a `CAMetalLayer *`.
//
// renderer: the renderer to query.
//
// Returns a `CAMetalLayer *` on success, or NULL if the renderer isn't a
// Metal renderer.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRenderMetalLayer
func (renderer *Renderer) MetalLayer() unsafe.Pointer {
	return (unsafe.Pointer)(C.SDL_GetRenderMetalLayer((*C.SDL_Renderer)(renderer)))
}

// Get the Metal command encoder for the current frame.
//
// This function returns `void *`, so SDL doesn't have to include Metal's
// headers, but it can be safely cast to an `id<MTLRenderCommandEncoder>`.
//
// This will return NULL if Metal refuses to give SDL a drawable to render to,
// which might happen if the window is hidden/minimized/offscreen. This
// doesn't apply to command encoders for render targets, just the window's
// backbuffer. Check your return values!
//
// renderer: the renderer to query.
//
// Returns an `id<MTLRenderCommandEncoder>` on success, or NULL if the
// renderer isn't a Metal renderer or there was an error.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRenderMetalCommandEncoder
func (renderer *Renderer) MetalCommandEncoder() unsafe.Pointer {
	return (unsafe.Pointer)(C.SDL_GetRenderMetalCommandEncoder((*C.SDL_Renderer)(renderer)))
}

// Add a set of synchronization semaphores for the current frame.
//
// The Vulkan renderer will wait for `wait_semaphore` before submitting
// rendering commands and signal `signal_semaphore` after rendering commands
// are complete for this frame.
//
// This should be called each frame that you want semaphore synchronization.
// The Vulkan renderer may have multiple frames in flight on the GPU, so you
// should have multiple semaphores that are used for synchronization. Querying
// SDL_PROP_RENDERER_VULKAN_SWAPCHAIN_IMAGE_COUNT_NUMBER will give you the
// maximum number of semaphores you'll need.
//
// renderer: the rendering context.
//
// wait_stage_mask: the VkPipelineStageFlags for the wait.
//
// wait_semaphore: a VkSempahore to wait on before rendering the current
// frame, or 0 if not needed.
//
// signal_semaphore: a VkSempahore that SDL will signal when rendering
// for the current frame is complete, or 0 if not
// needed.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// It is **NOT** safe to call this function from two threads at
// once.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_AddVulkanRenderSemaphores
func (renderer *Renderer) AddVulkanSemaphores(waitStage_mask uint32, waitSemaphore int64, signalSemaphore int64) bool {
	return (bool)(C.SDL_AddVulkanRenderSemaphores((*C.SDL_Renderer)(renderer), (C.Uint32)(waitStage_mask), (C.Sint64)(waitSemaphore), (C.Sint64)(signalSemaphore)))
}

// Toggle VSync of the given renderer.
//
// When a renderer is created, vsync defaults to SDL_RENDERER_VSYNC_DISABLED.
//
// The `vsync` parameter can be 1 to synchronize present with every vertical
// refresh, 2 to synchronize present with every second vertical refresh, etc.,
// SDL_RENDERER_VSYNC_ADAPTIVE for late swap tearing (adaptive vsync), or
// SDL_RENDERER_VSYNC_DISABLED to disable. Not every value is supported by
// every driver, so you should check the return value to see whether the
// requested setting is supported.
//
// renderer: the renderer to toggle.
//
// vsync: the vertical refresh sync interval.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_SetRenderVSync
func (renderer *Renderer) SetVSync(vsync int) error {
	if !C.SDL_SetRenderVSync((*C.SDL_Renderer)(renderer), (C.int)(vsync)) {
		return getError()
	}
	return nil
}

const RendererVSyncDisabled = 0
const RendererVSyncAdaptive = -1

// Get VSync of the given renderer.
//
// renderer: the renderer to toggle.
//
// vsync: an int filled with the current vertical refresh sync interval.
// See SDL_SetRenderVSync() for the meaning of the value.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_GetRenderVSync
func (renderer *Renderer) VSync() (int, error) {
	var vsync C.int
	if !C.SDL_GetRenderVSync((*C.SDL_Renderer)(renderer), &vsync) {
		return 0, getError()
	}
	return int(vsync), nil
}

// The size, in pixels, of a single SDL_RenderDebugText() character.
//
// The font is monospaced and square, so this applies to all characters.
//
// This macro is available since SDL 3.2.0.
const DebugTextFontCharacterSize = 8

// Draw debug text to an SDL_Renderer.
//
// This function will render a string of text to an SDL_Renderer. Note that
// this is a convenience function for debugging, with severe limitations, and
// not intended to be used for production apps and games.
//
// Among these limitations:
//
// - It accepts UTF-8 strings, but will only renders ASCII characters.
// - It has a single, tiny size (8x8 pixels). One can use logical presentation
// or scaling to adjust it, but it will be blurry.
// - It uses a simple, hardcoded bitmap font. It does not allow different font
// selections and it does not support truetype, for proper scaling.
// - It does no word-wrapping and does not treat newline characters as a line
// break. If the text goes out of the window, it's gone.
//
// For serious text rendering, there are several good options, such as
// SDL_ttf, stb_truetype, or other external libraries.
//
// On first use, this will create an internal texture for rendering glyphs.
// This texture will live until the renderer is destroyed.
//
// The text is drawn in the color specified by SDL_SetRenderDrawColor().
//
// renderer: the renderer which should draw a line of text.
//
// x: the x coordinate where the top-left corner of the text will draw.
//
// y: the y coordinate where the top-left corner of the text will draw.
//
// str: the string to render.
//
// Returns true on success or false on failure; call SDL_GetError() for more
// information.
//
// This function should only be called on the main thread.
//
// This function is available since SDL 3.2.0.
//
// https://wiki.libsdl.org/SDL3/SDL_RenderDebugText
func (renderer *Renderer) DebugText(x float32, y float32, str string) error {
	if !C.SDL_RenderDebugText((*C.SDL_Renderer)(renderer), (C.float)(x), (C.float)(y), tmpstring(str)) {
		return getError()
	}
	return nil
}
